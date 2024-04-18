package state

import (
	"fmt"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/state/protocol"
	"github.com/onflow/flow-go/state/protocol/protocol_state"
	"github.com/onflow/flow-go/storage"
	"github.com/onflow/flow-go/storage/badger/operation"
	"github.com/onflow/flow-go/storage/badger/transaction"
)

// stateMutator is a stateful object to evolve the protocol state. It is instantiated from the parent block's protocol state.
// State-changing operations can be iteratively applied and the stateMutator will internally evolve its in-memory state.
// While the StateMutator does not modify the database, it internally tracks the necessary database updates to persist its
// dependencies (specifically EpochSetup and EpochCommit events). Upon calling `Build` the stateMutator returns the updated
// protocol state, its ID and all database updates necessary for persisting the updated protocol state.
//
// The StateMutator is used by a replica's compliance layer to update protocol state when observing state-changing service in
// blocks. It is used by the primary in the block building process to obtain the correct protocol state for a proposal.
// Specifically, the leader may include state-changing service events in the block payload. The flow protocol prescribes that
// the proposal needs to include the ID of the protocol state, _after_ processing the payload incl. all state-changing events.
// Therefore, the leader instantiates a StateMutator, applies the service events to it and builds the updated protocol state ID.
//
// Not safe for concurrent use.
//
// TODO: Merge methods `EvolveState` and `Build` into one, as they must be always called in this succession (improves API's safety & clarity).
//
//	Temporarily, the stateMutator tracks internally that `EvolveState` is called before `Build` and errors otherwise.
type stateMutator struct {
	headers          storage.Headers
	results          storage.ExecutionResults
	kvStoreSnapshots storage.ProtocolKVStore

	parentState          protocol.KVStoreReader
	kvMutator            protocol_state.KVStoreMutator
	orthoKVStoreMachines []protocol_state.KeyValueStoreStateMachine

	// TODO: temporary shortcut until `EvolveState` and `Build` are merged
	evolveStateCalled bool
	buildCalled       bool
}

var _ protocol.StateMutator = (*stateMutator)(nil)

// newStateMutator creates a new instance of stateMutator.
// stateMutator performs initialization of state machine depending on the operation mode of the protocol.
// No errors are expected during normal operations.
func newStateMutator(
	headers storage.Headers,
	results storage.ExecutionResults,
	kvStoreSnapshots storage.ProtocolKVStore,
	candidateView uint64,
	parentID flow.Identifier,
	parentState protocol_state.KVStoreAPI,
	stateMachineFactories ...protocol_state.KeyValueStoreStateMachineFactory,
) (*stateMutator, error) {
	protocolVersion := parentState.GetProtocolStateVersion()
	if versionUpgrade := parentState.GetVersionUpgrade(); versionUpgrade != nil {
		if candidateView >= versionUpgrade.ActivationView {
			protocolVersion = versionUpgrade.Data
		}
	}

	replicatedState, err := parentState.Replicate(protocolVersion)
	if err != nil {
		return nil, fmt.Errorf("could not replicate parent KV store (version=%d) to protocol version %d: %w",
			parentState.GetProtocolStateVersion(), protocolVersion, err)
	}

	stateMachines := make([]protocol_state.KeyValueStoreStateMachine, 0, len(stateMachineFactories))
	for _, factory := range stateMachineFactories {
		stateMachine, err := factory.Create(candidateView, parentID, parentState, replicatedState)
		if err != nil {
			return nil, fmt.Errorf("could not create state machine: %w", err)
		}
		stateMachines = append(stateMachines, stateMachine)
	}

	return &stateMutator{
		headers:              headers,
		results:              results,
		kvStoreSnapshots:     kvStoreSnapshots,
		orthoKVStoreMachines: stateMachines,
		parentState:          parentState,
		kvMutator:            replicatedState,
		evolveStateCalled:    false,
	}, nil
}

// Build constructs the resulting protocol state, *after* applying all the sealed service events in a block (under construction)
// via `EvolveState(...)`. It returns:
//   - stateID: the hash commitment to the updated Protocol State Snapshot
//   - dbUpdates: database updates necessary for persisting the State Snapshot itself including all data structures
//     that the Snapshot references. In addition, `dbUpdates` also populates the `ProtocolKVStore.ByBlockID`.
//     Therefore, even if there are no changes of the Protocol State, `dbUpdates` still contains deferred storage writes
//     that must be executed to populate the `ByBlockID` index.
//   - err: All error returns indicate potential state corruption and should therefore be treated as fatal.
//
// CAUTION:
//   - For Consensus Participants that are replicas, the calling code must check that the returned `stateID` matches the
//     commitment in the block proposal! If they don't match, the proposal is byzantine and should be slashed.
func (m *stateMutator) Build() (stateID flow.Identifier, dbUpdates *protocol.DeferredBlockPersist, err error) {
	if !m.evolveStateCalled { // TODO: temporary shortcut until `EvolveState` and `Build` are merged
		return flow.ZeroID, protocol.NewDeferredBlockPersist(), irrecoverable.NewExceptionf("cannot build Protocol State without prior call of EvolveState method")
	}
	if m.buildCalled {
		return flow.ZeroID, protocol.NewDeferredBlockPersist(), irrecoverable.NewExceptionf("repeated Build calls are not allowed")
	}
	m.buildCalled = true

	dbUpdates = protocol.NewDeferredBlockPersist()
	for _, stateMachine := range m.orthoKVStoreMachines {
		dbOps, err := stateMachine.Build()
		if err != nil {
			return flow.ZeroID, protocol.NewDeferredBlockPersist(), fmt.Errorf("unexpected exception building state machine's output state: %w", err)
		}
		dbUpdates.AddIndexingOps(dbOps.Pending())
	}
	stateID = m.kvMutator.ID()
	version, data, err := m.kvMutator.VersionedEncode()
	if err != nil {
		return flow.ZeroID, protocol.NewDeferredBlockPersist(), fmt.Errorf("could not encode protocol state: %w", err)
	}

	// Schedule deferred database operations to index the protocol state by the candidate block's ID
	// and persist the new protocol state (if there are any changes)
	dbUpdates.AddIndexingOp(func(blockID flow.Identifier, tx *transaction.Tx) error {
		return m.kvStoreSnapshots.IndexTx(blockID, stateID)(tx)
	})
	dbUpdates.AddDbOp(operation.SkipDuplicatesTx(m.kvStoreSnapshots.StoreTx(stateID, &storage.KeyValueStoreData{
		Version: version,
		Data:    data,
	})))

	return stateID, dbUpdates, nil
}

// EvolveState updates the overall Protocol State based on information from the candidate block
// (potentially still under construction). Information that may change the state is:
//   - the candidate block's view (already provided at construction time)
//   - Service Events sealed in the candidate block
//
// We only mutate the `StateMutator`'s internal in-memory copy of the protocol state, without
// changing the parent state (i.e. the state we started from).
//
// In a nutshell, we proceed as follows:
//   - If there are any, we arrange the sealed service events in chronologically order. This
//     can be achieved by ordering the sealed execution results by increasing block height.
//     Within each execution result, the service events are in chronological order.
//   - We call `KeyValueStoreStateMachine.EvolveState(..)` for each of the
//     `OrthogonalStoreStateMachine`s and provide the sealed service events as input. EvolveState
//     is always called on each state machine, even if there are no service events, because
//     reaching or exceeding a certain view can trigger a state change (e.g. Epoch Fallback Mode).
//   - We collect the deferred database updates necessary to persist each of the updated sub-states
//     including all of their dependencies and respective indices. The subsequent `Build` step will
//     add further db updates. Executing the deferred database updates is the responsibility of
//     the calling code.
//
// SAFETY REQUIREMENT:
// The StateMutator assumes that the proposal has passed the following correctness checks!
//   - The seals in the payload continuously follow the ancestry of this fork. Specifically,
//     there are no gaps in the seals.
//   - The seals guarantee correctness of the sealed execution result, including the contained
//     service events. This is actively checked by the verification node, whose aggregated
//     approvals in the form of a seal attest to the correctness of the sealed execution result
//     (specifically the Service Events contained in the result and their order).
//   - `EvolveState` must be called before `Build`
//
// Consensus nodes actively verify protocol compliance for any block proposal they receive,
// including integrity of each seal individually as well as the seals continuously following
// the fork. Light clients only process certified blocks, which guarantees that consensus nodes
// already ran those checks and found the proposal to be valid.
//
// Details on SERVICE EVENTS:
// Consider a chain where a service event is emitted during execution of block A. Block B contains
// an execution receipt `RA` for A. Block C contains a seal `SA` for A's execution result.
//
//	A <- .. <- B(RA) <- .. <- C(SA)
//
// Service Events are included within execution results, which are stored opaquely as part of the
// block payload (block B in our example). We only validate, process and persist the typed service
// event to storage once we process C, the block containing the seal for block A. This is because
// we rely on the sealing subsystem to validate correctness of the service event before processing
// it. Consequently, any change to the protocol state introduced by a service event emitted during
// execution of block A would only become visible when querying C or its descendants.
//
// Error returns:
// [TLDR] All error returns indicate potential state corruption and should therefore be treated as fatal.
//   - Per convention, the input seals from the block payload have already been confirmed to be protocol compliant.
//     Hence, the service events in the sealed execution results represent the honest execution path.
//     Therefore, the sealed service events should encode a valid evolution of the protocol state -- provided
//     the system smart contracts are correct.
//   - As we can rule out byzantine attacks as the source of failures, the only remaining sources of problems
//     can be (a) bugs in the system smart contracts or (b) bugs in the node implementation. A service event
//     not representing a valid state transition despite all consistency checks passing is interpreted as
//     case (a) and _should be_ handled internally by the respective state machine. Otherwise, any bug or
//     unforeseen edge cases in the system smart contracts would in consensus halt, due to errors while
//     evolving the protocol state.
//   - A consistency or sanity check failing within the StateMutator is likely the symptom of an internal bug
//     in the node software or state corruption, i.e. case (b). This is the only scenario where the error return
//     of this function is not nil. If such an exception is returned, continuing is not an option.
func (m *stateMutator) EvolveState(seals []*flow.Seal) error {
	if m.evolveStateCalled { // TODO: temporary shortcut until `EvolveState` and `Build` are merged
		return irrecoverable.NewExceptionf("repeated calls of EvolveState are not allowed")
	}
	m.evolveStateCalled = true

	// block payload may not specify seals in order, so order them by block height before processing
	orderedSeals, err := protocol.OrderedSeals(seals, m.headers)
	if err != nil {
		// Per API contract, the input seals must have already passed verification, which necessitates
		// successful ordering. Hence, calling protocol.OrderedSeals with the same inputs that succeeded
		// earlier now failed. In all cases, this is an exception.
		return irrecoverable.NewExceptionf("ordering already validated seals unexpectedly failed: %w", err)
	}
	results := make([]*flow.ExecutionResult, 0, len(orderedSeals))
	for _, seal := range orderedSeals {
		result, err := m.results.ByID(seal.ResultID)
		if err != nil {
			return fmt.Errorf("could not get result (id=%x) for seal (id=%x): %w", seal.ResultID, seal.ID(), err)
		}
		results = append(results, result)
	}

	// order all service events in one list
	orderedUpdates := make([]flow.ServiceEvent, 0)
	for _, result := range results {
		for _, event := range result.ServiceEvents {
			orderedUpdates = append(orderedUpdates, event)
		}
	}

	for _, stateMachine := range m.orthoKVStoreMachines {
		// only exceptions should be propagated
		err := stateMachine.EvolveState(orderedUpdates)
		if err != nil {
			return fmt.Errorf("could not process protocol state change for candidate block: %w", err)
		}
	}

	return nil
}
