package execution_indexer

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go/cmd/util/ledger/migrations"
	"github.com/onflow/flow-go/engine/execution/scripts"
	"github.com/onflow/flow-go/fvm/storage/snapshot"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/storage"
)

var _ module.ExecutionStateIndexer = &Indexer{}
var _ scripts.ScriptExecutionState = &Indexer{}

type Indexer struct {
	registers   storage.Registers
	headers     storage.Headers
	last        uint64                          // todo persist
	commitments map[uint64]flow.StateCommitment // todo persist
}

func New(registers storage.Registers, headers storage.Headers) *Indexer {
	return &Indexer{
		registers:   registers,
		headers:     headers,
		last:        0,
		commitments: make(map[uint64]flow.StateCommitment),
	}
}

func (i *Indexer) NewStorageSnapshot(commitment flow.StateCommitment) snapshot.StorageSnapshot {
	var height uint64
	for h, commit := range i.commitments {
		if commit == commitment {
			height = h
		}
	}

	// TODO what if commitment is not found

	reader := func(id flow.RegisterID) (flow.RegisterValue, error) {
		entry, err := i.registers.Get(id, height)
		if err != nil {
			return nil, err
		}

		return entry.Value, nil
	}

	return snapshot.NewReadFuncStorageSnapshot(reader)
}

func (i *Indexer) StateCommitmentByBlockID(
	ctx context.Context,
	identifier flow.Identifier,
) (flow.StateCommitment, error) {
	h, err := i.headers.ByBlockID(identifier)
	if err != nil {
		return flow.DummyStateCommitment, fmt.Errorf("could not find block by ID %s for state commitment: %w", identifier.String(), err)
	}

	commit, ok := i.commitments[h.Height]
	if !ok {
		return flow.DummyStateCommitment, fmt.Errorf("could not find the state commitment by height %d", h.Height)
	}

	return commit, nil
}

func (i *Indexer) HasState(commitment flow.StateCommitment) bool {
	for _, c := range i.commitments {
		if c == commitment {
			return true
		}
	}

	return false
}

func (i *Indexer) Last() (uint64, error) {
	return i.last, nil
}

func (i *Indexer) StoreLast(last uint64) error {
	i.last = last
	return nil
}

func (i *Indexer) HeightByBlockID(ID flow.Identifier) (uint64, error) {
	header, err := i.headers.ByBlockID(ID)
	if err != nil {
		return 0, err
	}

	return header.Height, nil
}

func (i *Indexer) Commitment(height uint64) (flow.StateCommitment, error) {
	val, ok := i.commitments[height]
	if !ok {
		return flow.DummyStateCommitment, fmt.Errorf("could not find commitment at height %d", height)
	}

	return val, nil
}

func (i *Indexer) StoreCommitment(commitment flow.StateCommitment, height uint64) error {
	i.commitments[height] = commitment
	return nil
}

func (i *Indexer) Values(IDs flow.RegisterIDs, height uint64) ([]flow.RegisterValue, error) {
	values := make([]flow.RegisterValue, len(IDs))

	for j, id := range IDs {
		entry, err := i.registers.Get(id, height)
		if err != nil {
			return nil, err
		}

		values[j] = entry.Value
	}

	return values, nil
}

func (i *Indexer) StorePayloads(payloads []*ledger.Payload, height uint64) error {
	regEntries := make(flow.RegisterEntries, len(payloads))

	for j, payload := range payloads {
		k, err := payload.Key()
		if err != nil {
			return err
		}

		id, err := migrations.KeyToRegisterID(k)
		if err != nil {
			return err
		}

		regEntries[j] = flow.RegisterEntry{
			Key:   id,
			Value: payload.Value(),
		}
	}

	return i.registers.Store(regEntries, height)
}
