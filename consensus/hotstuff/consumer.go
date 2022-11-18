package hotstuff

import (
	"time"

	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
)

// FinalizationConsumer consumes outbound notifications produced by the finalization logic.
// Notifications represent finalization-specific state changes which are potentially relevant
// to the larger node. The notifications are emitted in the order in which the
// finalization algorithm makes the respective steps.
//
// Implementations must:
//   - be concurrency safe
//   - be non-blocking
//   - handle repetition of the same events (with some processing overhead).
type FinalizationConsumer interface {

	// OnBlockIncorporated notifications are produced by the Finalization Logic
	// whenever a block is incorporated into the consensus state.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnBlockIncorporated(*model.Block)

	// OnFinalizedBlock notifications are produced by the Finalization Logic whenever
	// a block has been finalized. They are emitted in the order the blocks are finalized.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnFinalizedBlock(*model.Block)

	// OnDoubleProposeDetected notifications are produced by the Finalization Logic
	// whenever a double block proposal (equivocation) was detected.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnDoubleProposeDetected(*model.Block, *model.Block)
}

// Consumer consumes outbound notifications produced by HotStuff and its components.
// Notifications are consensus-internal state changes which are potentially relevant to
// the larger node in which HotStuff is running. The notifications are emitted
// in the order in which the HotStuff algorithm makes the respective steps.
//
// Implementations must:
//   - be concurrency safe
//   - be non-blocking
//   - handle repetition of the same events (with some processing overhead).
type Consumer interface {
	FinalizationConsumer
	CommunicatorConsumer

	// OnEventProcessed notifications are produced by the EventHandler when it is done processing
	// and hands control back to the EventLoop to wait for the next event.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnEventProcessed()

	// OnStart notifications are produced by the EventHandler when it starts blocks recovery and
	// prepares for handling incoming events from EventLoop.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnStart(currentView uint64)

	// OnReceiveProposal notifications are produced by the EventHandler when it starts processing a block.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnReceiveProposal(currentView uint64, proposal *model.Proposal)

	// OnReceiveQc notifications are produced by the EventHandler when it starts processing a QC.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnReceiveQc(currentView uint64, qc *flow.QuorumCertificate)

	// OnReceiveTc notifications are produced by the EventHandler when it starts processing a TC.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnReceiveTc(currentView uint64, tc *flow.TimeoutCertificate)

	// OnPartialTc notifications are produced by the EventHandler when it starts processing partial TC
	// constructed by local timeout aggregator.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnPartialTc(currentView uint64, partialTc *PartialTcCreated)

	// OnLocalTimeout notifications are produced by the EventHandler when it reacts to expiry of round duration timer.
	// Such a notification indicates that the PaceMaker's timeout was processed by the system.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnLocalTimeout(currentView uint64)

	// OnViewChange notifications are produced by PaceMaker when it moves to a new view
	// based on processing a QC or TC. The arguments specify the oldView (first argument),
	// and the newView to which the PaceMaker transitioned (second argument).
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnViewChange(oldView, newView uint64)

	// OnQcTriggeredViewChange notifications are produced by PaceMaker when it moves to a new view
	// based on processing a QC. The arguments specify the qc (first argument), which triggered
	// the view change, and the newView to which the PaceMaker transitioned (second argument).
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnQcTriggeredViewChange(qc *flow.QuorumCertificate, newView uint64)

	// OnTcTriggeredViewChange notifications are produced by PaceMaker when it moves to a new view
	// based on processing a TC. The arguments specify the tc (first argument), which triggered
	// the view change, and the newView to which the PaceMaker transitioned (second argument).
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnTcTriggeredViewChange(tc *flow.TimeoutCertificate, newView uint64)

	// OnStartingTimeout notifications are produced by PaceMaker. Such a notification indicates that the
	// PaceMaker is now waiting for the system to (receive and) process blocks or votes.
	// The specific timeout type is contained in the TimerInfo.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnStartingTimeout(model.TimerInfo)

	// OnDoubleVotingDetected notifications are produced by the Vote Aggregation logic
	// whenever a double voting (same voter voting for different blocks at the same view) was detected.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnDoubleVotingDetected(*model.Vote, *model.Vote)

	// OnInvalidVoteDetected notifications are produced by the Vote Aggregation logic
	// whenever an invalid vote was detected.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnInvalidVoteDetected(*model.Vote)

	// OnVoteForInvalidBlockDetected notifications are produced by the Vote Aggregation logic
	// whenever vote for invalid proposal was detected.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnVoteForInvalidBlockDetected(vote *model.Vote, invalidProposal *model.Proposal)

	// OnDoubleTimeoutDetected notifications are produced by the Timeout Aggregation logic
	// whenever a double timeout (same replica producing two different timeouts at the same view) was detected.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnDoubleTimeoutDetected(*model.TimeoutObject, *model.TimeoutObject)

	// OnInvalidTimeoutDetected notifications are produced by the Timeout Aggregation logic
	// whenever an invalid timeout was detected.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnInvalidTimeoutDetected(*model.TimeoutObject)
}

// QCCreatedConsumer consumes outbound notifications produced by HotStuff and its components.
// Notifications are consensus-internal state changes which are potentially relevant to
// the larger node in which HotStuff is running. The notifications are emitted
// in the order in which the HotStuff algorithm makes the respective steps.
//
// Implementations must:
//   - be concurrency safe
//   - be non-blocking
//   - handle repetition of the same events (with some processing overhead).
type QCCreatedConsumer interface {
	// OnQcConstructedFromVotes notifications are produced by the VoteAggregator
	// component, whenever it constructs a QC from votes.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnQcConstructedFromVotes(*flow.QuorumCertificate)
}

// TimeoutCollectorConsumer consumes outbound notifications produced by HotStuff's timeout aggregation
// component. These events are primarily intended for the HotStuff-internal state machine (EventHandler),
// but might also be relevant to the larger node in which HotStuff is running.
//
// Caution: the events are not strictly ordered by increasing views!
// The notifications are emitted by concurrent processing logic. Over larger time scales, the
// emitted events are for statistically increasing views. However, on short time scales there
// are _no_ monotonicity guarantees w.r.t. the events' views.
//
// Implementations must:
//   - be concurrency safe
//   - be non-blocking
//   - handle repetition of the same events (with some processing overhead).
type TimeoutCollectorConsumer interface {
	// OnTcConstructedFromTimeouts notifications are produced by the TimeoutProcessor
	// component, whenever it constructs a TC based on TimeoutObjects from a
	// supermajority of consensus participants.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnTcConstructedFromTimeouts(certificate *flow.TimeoutCertificate)

	// OnPartialTcCreated notifications are produced by the TimeoutProcessor
	// component, whenever it collected TimeoutObjects from a superminority
	// of consensus participants for a specific view. Along with the view, it
	// reports the newest QC and TC (for previous view) discovered in process of
	// timeout collection. Per convention, the newest QC is never nil, while
	// the TC for the previous view might be nil.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnPartialTcCreated(view uint64, newestQC *flow.QuorumCertificate, lastViewTC *flow.TimeoutCertificate)

	// OnNewQcDiscovered notifications are produced by the TimeoutCollector
	// component, whenever it discovers new QC included in timeout object.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnNewQcDiscovered(certificate *flow.QuorumCertificate)

	// OnNewTcDiscovered notifications are produced by the TimeoutCollector
	// component, whenever it discovers new TC included in timeout object.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnNewTcDiscovered(certificate *flow.TimeoutCertificate)
}

// CommunicatorConsumer consumes outbound notifications produced by HotStuff and it's components.
// Notifications allow the HotStuff core algorithm to communicate with the other actors of the consensus process.
// Implementations must:
//   - be concurrency safe
//   - be non-blocking
//   - handle repetition of the same events (with some processing overhead).
type CommunicatorConsumer interface {
	// OnOwnVote notifies about intent to send a vote for the given parameters to the specified recipient.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnOwnVote(blockID flow.Identifier, view uint64, sigData []byte, recipientID flow.Identifier)

	// OnOwnTimeout notifies about intent to broadcast the given timeout object(TO) to all actors of the consensus process.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnOwnTimeout(timeout *model.TimeoutObject)

	// OnOwnProposal notifies about intent to broadcast the given block proposal to all actors of
	// the consensus process.
	// delay is to hold the proposal before broadcasting it. Useful to control the block production rate.
	// Prerequisites:
	// Implementation must be concurrency safe; Non-blocking;
	// and must handle repetition of the same events (with some processing overhead).
	OnOwnProposal(proposal *flow.Header, targetPublicationTime time.Time)
}
