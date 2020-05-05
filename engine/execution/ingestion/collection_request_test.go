package ingestion

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/messages"
	"github.com/dapperlabs/flow-go/module/mempool/entity"
	realStorage "github.com/dapperlabs/flow-go/storage"
	"github.com/dapperlabs/flow-go/utils/unittest"
)

type colReqMatcher struct {
	req *messages.CollectionRequest
}

func (c *colReqMatcher) Matches(x interface{}) bool {
	other := x.(*messages.CollectionRequest)
	return c.req.ID == other.ID
}
func (c *colReqMatcher) String() string {
	return fmt.Sprintf("ID %x", c.req.ID)
}

func TestCollectionRequests(t *testing.T) {

	runWithEngine(t, func(ctx testingContext) {

		block := unittest.BlockFixture()
		//To make sure we always have collection if the block fixture changes
		guarantees := unittest.CollectionGuaranteesFixture(3)

		guarantees[0].SignerIDs = []flow.Identifier{collection1Identity.NodeID, collection3Identity.NodeID}
		guarantees[1].SignerIDs = []flow.Identifier{collection2Identity.NodeID}
		guarantees[2].SignerIDs = []flow.Identifier{collection2Identity.NodeID, collection3Identity.NodeID}

		block.Payload.Guarantees = guarantees
		block.Header.PayloadHash = block.Payload.Hash()

		ctx.blocks.EXPECT().Store(gomock.Eq(&block))
		ctx.state.On("AtBlockID", block.ID()).Return(ctx.snapshot).Maybe()

		ctx.collectionConduit.EXPECT().Submit(
			&colReqMatcher{req: &messages.CollectionRequest{ID: guarantees[0].ID(), Nonce: rand.Uint64()}},
			gomock.Eq([]flow.Identifier{collection1Identity.NodeID, collection3Identity.NodeID}),
		)

		ctx.collectionConduit.EXPECT().Submit(
			&colReqMatcher{req: &messages.CollectionRequest{ID: guarantees[1].ID(), Nonce: rand.Uint64()}},
			gomock.Eq(collection2Identity.NodeID),
		)

		ctx.collectionConduit.EXPECT().Submit(
			&colReqMatcher{req: &messages.CollectionRequest{ID: guarantees[2].ID(), Nonce: rand.Uint64()}},
			gomock.Eq([]flow.Identifier{collection2Identity.NodeID, collection3Identity.NodeID}),
		)

		ctx.executionState.On("StateCommitmentByBlockID", block.Header.ParentID).Return(unittest.StateCommitmentFixture(), nil)

		proposal := unittest.ProposalFromBlock(&block)
		err := ctx.engine.ProcessLocal(proposal)

		require.NoError(t, err)
	})
}

func TestNoCollectionRequestsIfParentMissing(t *testing.T) {

	runWithEngine(t, func(ctx testingContext) {

		block := unittest.BlockFixture()
		//To make sure we always have collection if the block fixture changes
		guarantees := unittest.CollectionGuaranteesFixture(3)

		guarantees[0].SignerIDs = []flow.Identifier{collection1Identity.NodeID, collection3Identity.NodeID}
		guarantees[1].SignerIDs = []flow.Identifier{collection2Identity.NodeID}
		guarantees[2].SignerIDs = []flow.Identifier{collection2Identity.NodeID, collection3Identity.NodeID}

		block.Payload.Guarantees = guarantees
		block.Header.PayloadHash = block.Payload.Hash()

		ctx.blocks.EXPECT().Store(gomock.Eq(&block))
		ctx.state.On("AtBlockID", block.ID()).Return(ctx.snapshot).Maybe()

		ctx.collectionConduit.EXPECT().Submit(gomock.Any(), gomock.Any()).Times(0)

		ctx.executionState.On("StateCommitmentByBlockID", block.Header.ParentID).Return(nil, realStorage.ErrNotFound)

		proposal := unittest.ProposalFromBlock(&block)
		err := ctx.engine.ProcessLocal(proposal)

		require.NoError(t, err)
	})
}

func TestValidatingCollectionResponse(t *testing.T) {

	runWithEngine(t, func(ctx testingContext) {

		executableBlock := unittest.ExecutableBlockFixture([][]flow.Identifier{{collection1Identity.NodeID}})
		executableBlock.StartState = unittest.StateCommitmentFixture()

		ctx.blocks.EXPECT().Store(gomock.Eq(executableBlock.Block))

		id := executableBlock.Collections()[0].Guarantee.ID()

		ctx.state.On("AtBlockID", executableBlock.Block.ID()).Return(ctx.snapshot).Maybe()

		ctx.collectionConduit.EXPECT().Submit(
			&colReqMatcher{req: &messages.CollectionRequest{ID: id, Nonce: rand.Uint64()}},
			gomock.Eq(collection1Identity.NodeID),
		).Return(nil)
		ctx.executionState.On("StateCommitmentByBlockID", executableBlock.Block.Header.ParentID).Return(executableBlock.StartState, nil)

		proposal := unittest.ProposalFromBlock(executableBlock.Block)
		err := ctx.engine.ProcessLocal(proposal)
		require.NoError(t, err)

		rightResponse := messages.CollectionResponse{
			Collection: flow.Collection{Transactions: executableBlock.Collections()[0].Transactions},
		}

		// TODO Enable wrong response sending once we have a way to hash collection

		// wrongResponse := provider.CollectionResponse{
		//	Fingerprint:  fingerprint,
		//	Transactions: []flow.TransactionBody{tx},
		// }

		// engine.Submit(collectionIdentity.NodeID, wrongResponse)

		// no interaction with conduit for finished executableBlock
		// </TODO enable>

		//ctx.executionState.On("StateCommitmentByBlockID", executableBlock.Block.Header.ParentID).Return(unittest.StateCommitmentFixture(), realStorage.ErrNotFound)

		ctx.assertSuccessfulBlockComputation(executableBlock, unittest.IdentifierFixture())

		err = ctx.engine.ProcessLocal(&rightResponse)
		require.NoError(t, err)
	})
}

func TestNoBlockExecutedUntilAllCollectionsArePosted(t *testing.T) {

	runWithEngine(t, func(ctx testingContext) {

		executableBlock := unittest.ExecutableBlockFixture(
			[][]flow.Identifier{
				{collection1Identity.NodeID},
				{collection1Identity.NodeID},
				{collection1Identity.NodeID},
			},
		)

		for _, col := range executableBlock.Block.Payload.Guarantees {
			ctx.collectionConduit.EXPECT().Submit(
				&colReqMatcher{req: &messages.CollectionRequest{ID: col.ID(), Nonce: rand.Uint64()}},
				gomock.Eq(collection1Identity.NodeID),
			)
		}

		ctx.state.On("AtBlockID", executableBlock.ID()).Return(ctx.snapshot)

		ctx.blocks.EXPECT().Store(gomock.Eq(executableBlock.Block))
		ctx.executionState.On("StateCommitmentByBlockID", executableBlock.Block.Header.ParentID).Return(unittest.StateCommitmentFixture(), nil)

		proposal := unittest.ProposalFromBlock(executableBlock.Block)
		err := ctx.engine.ProcessLocal(proposal)
		require.NoError(t, err)

		// Expected no calls so test should fail if any occurs
		rightResponse := messages.CollectionResponse{
			Collection: flow.Collection{Transactions: executableBlock.Collections()[1].Transactions},
		}

		err = ctx.engine.ProcessLocal(&rightResponse)

		require.NoError(t, err)
	})
}

func TestCollectionSharedByMultipleBlocks(t *testing.T) {

	runWithEngine(t, func(ctx testingContext) {

		blockA := unittest.BlockFixture()
		blockB := unittest.BlockFixture()

		collection := unittest.CollectionFixture(1)
		guarantee := collection.Guarantee()
		guarantee.SignerIDs = []flow.Identifier{collection1Identity.NodeID}

		blockA.Payload.Guarantees = []*flow.CollectionGuarantee{&guarantee}
		blockA.Header.PayloadHash = blockA.Payload.Hash()

		blockB.Payload.Guarantees = []*flow.CollectionGuarantee{&guarantee}
		blockB.Header.PayloadHash = blockB.Payload.Hash()

		ctx.blocks.EXPECT().Store(gomock.Eq(&blockA))
		ctx.blocks.EXPECT().Store(gomock.Eq(&blockB))
		ctx.state.On("AtBlockID", blockA.ID()).Return(ctx.snapshot).Maybe()
		ctx.state.On("AtBlockID", blockB.ID()).Return(ctx.snapshot).Maybe()

		ctx.collectionConduit.EXPECT().Submit(
			&colReqMatcher{req: &messages.CollectionRequest{ID: guarantee.ID(), Nonce: rand.Uint64()}},
			gomock.Any(),
		).Times(1)

		ctx.executionState.On("StateCommitmentByBlockID", blockA.Header.ParentID).Return(unittest.StateCommitmentFixture(), nil)
		ctx.executionState.On("StateCommitmentByBlockID", blockB.Header.ParentID).Return(unittest.StateCommitmentFixture(), nil)

		proposalA := unittest.ProposalFromBlock(&blockA)
		err := ctx.engine.ProcessLocal(proposalA)
		require.NoError(t, err)

		proposalB := unittest.ProposalFromBlock(&blockB)
		err = ctx.engine.ProcessLocal(proposalB)
		require.NoError(t, err)

		ctx.computationManager.On("ComputeBlock", mock.MatchedBy(func(b *entity.ExecutableBlock) bool {
			return b.Block.ID() == blockA.ID()
		}), mock.Anything).Return(nil, nil)

		ctx.computationManager.On("ComputeBlock", mock.MatchedBy(func(b *entity.ExecutableBlock) bool {
			return b.Block.ID() == blockB.ID()
		}), mock.Anything).Return(nil, nil)

		rightResponse := messages.CollectionResponse{
			Collection: flow.Collection{Transactions: collection.Transactions},
		}

		err = ctx.engine.ProcessLocal(&rightResponse)

		require.NoError(t, err)
	})
}
