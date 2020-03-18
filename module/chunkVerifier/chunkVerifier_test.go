package chunkVerifier

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/dapperlabs/flow-go/engine/execution/computation/virtualmachine"
	"github.com/dapperlabs/flow-go/engine/verification"
	"github.com/dapperlabs/flow-go/language/runtime"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/storage/ledger"
	"github.com/dapperlabs/flow-go/storage/ledger/databases/leveldb"
	"github.com/dapperlabs/flow-go/utils/unittest"
)

type ChunkVerifierTestSuite struct {
	suite.Suite
	verifier ChunkVerifier
}

// Make sure variables are set properly
func (s *ChunkVerifierTestSuite) SetupTest() {
	s.verifier = NewFlowChunkVerifier(newVirtualMachineMock())
}

// TestVerification invokes all the tests in this test suite
func TestVerification(t *testing.T) {
	suite.Run(t, new(ChunkVerifierTestSuite))
}

// TestHappyPath tests verification of the baseline verifiable chunk
func (s *ChunkVerifierTestSuite) TestHappyPath() {
	vch := GetBaselineVerifiableChunk([]byte{})
	err := s.verifier.Verify(vch)
	assert.Nil(s.T(), err)
}

// TestMissingRegisterTouchForUpdate tests verification of the a chunkdatapack missing a register touch (update)
func (s *ChunkVerifierTestSuite) TestMissingRegisterTouchForUpdate() {
	vch := GetBaselineVerifiableChunk([]byte(""))
	// remove the second register touch
	vch.ChunkDataPack.RegisterTouches = vch.ChunkDataPack.RegisterTouches[:1]
	err := s.verifier.Verify(vch)
	assert.NotNil(s.T(), err)
}

// TODO
// TestMissingRegisterTouchForRead tests verification of the a chunkdatapack missing a register touch (read)
// func (s *ChunkVerifierTestSuite) TestMissingRegisterTouchForRead() {
// 	vch := GetBaselineVerifiableChunk([]byte(""))
// 	// remove the second register touch
// 	vch.ChunkDataPack.RegisterTouches = vch.ChunkDataPack.RegisterTouches[1:]
// 	err := s.verifier.Verify(vch)
// 	assert.NotNil(s.T(), err)
// }

func (s *ChunkVerifierTestSuite) TestWrongEndState() {
	vch := GetBaselineVerifiableChunk([]byte("wrongEndState"))
	err := s.verifier.Verify(vch)
	assert.NotNil(s.T(), err)
}

func GetBaselineVerifiableChunk(script []byte) *verification.VerifiableChunk {
	// Collection setup

	coll := unittest.CollectionFixture(5)
	coll.Transactions[3] = &flow.TransactionBody{Script: script}

	guarantee := coll.Guarantee()

	// Block setup
	payload := flow.Payload{
		Identities: unittest.IdentityListFixture(32),
		Guarantees: []*flow.CollectionGuarantee{&guarantee},
	}
	header := unittest.BlockHeaderFixture()
	header.PayloadHash = payload.Hash()
	block := flow.Block{
		Header:  header,
		Payload: payload,
	}

	// registerTouch and State setup
	id1 := make([]byte, 32)
	value1 := []byte{'a'}

	id2 := make([]byte, 32)
	id2[0] = byte(5)
	value2 := []byte{'b'}
	UpdatedValue2 := []byte{'B'}

	ids := make([][]byte, 0)
	values := make([][]byte, 0)
	ids = append(ids, id1, id2)
	values = append(values, value1, value2)

	db, _ := TempLevelDB()
	f, _ := ledger.NewTrieStorage(db)
	startState, _ := f.UpdateRegisters(ids, values)
	regTs, _ := f.GetRegisterTouches(ids, startState)

	ids = [][]byte{id2}
	values = [][]byte{UpdatedValue2}
	endState, _ := f.UpdateRegisters(ids, values)

	// Chunk setup
	chunk := flow.Chunk{
		ChunkBody: flow.ChunkBody{
			CollectionIndex: 0,
			StartState:      startState,
		},
		Index: 0,
	}

	chunkDataPack := flow.ChunkDataPack{
		ChunkID:         chunk.ID(),
		StartState:      startState,
		RegisterTouches: regTs,
	}

	// ExecutionResult setup
	result := flow.ExecutionResult{
		ExecutionResultBody: flow.ExecutionResultBody{
			BlockID: block.ID(),
			Chunks:  flow.ChunkList{&chunk},
		},
	}

	receipt := flow.ExecutionReceipt{
		ExecutionResult: result,
	}

	return &verification.VerifiableChunk{
		ChunkIndex:    chunk.Index,
		EndState:      endState,
		Block:         &block,
		Receipt:       &receipt,
		Collection:    &coll,
		ChunkDataPack: &chunkDataPack,
	}

}

// ChunkMissingRegTouch *verification.VerifiableChunk
// ChunkWrongEndState   *verification.VerifiableChunk
// TooManyInputs

func TempLevelDB() (*leveldb.LevelDB, error) {
	dir, err := ioutil.TempDir("", "flow-test-db")
	if err != nil {
		return nil, err
	}
	kvdbPath := filepath.Join(dir, "kvdb")
	tdbPath := filepath.Join(dir, "tdb")
	db, err := leveldb.NewLevelDB(kvdbPath, tdbPath)
	return db, err
}

type blockContextMock struct {
	vm     *virtualMachineMock
	header *flow.Header
}

func (bc *blockContextMock) ExecuteTransaction(
	ledger virtualmachine.Ledger,
	tx *flow.TransactionBody,
	options ...virtualmachine.TransactionContextOption,
) (*virtualmachine.TransactionResult, error) {
	var txRes virtualmachine.TransactionResult
	switch string(tx.Script) {
	case "wrongEndState":
		id1 := make([]byte, 32)
		UpdatedValue1 := []byte{'F'}
		// add updates to the ledger
		ledger.Set(id1, UpdatedValue1)
		txRes = virtualmachine.TransactionResult{
			TransactionID: unittest.IdentifierFixture(),
			Events:        []runtime.Event{},
			Logs:          []string{"log1", "log2"}, // []string
			Error:         nil,                      // inside the runtime (e.g. div by zero, access account)
			GasUsed:       0,
		}
	default:
		id1 := make([]byte, 32)
		id2 := make([]byte, 32)
		id2[0] = byte(5)
		UpdatedValue2 := []byte{'B'}
		_, _ = ledger.Get(id1)
		ledger.Set(id2, UpdatedValue2)
		txRes = virtualmachine.TransactionResult{
			TransactionID: unittest.IdentifierFixture(),
			Events:        []runtime.Event{},
			Logs:          []string{"log1", "log2"}, // []string
			Error:         nil,                      // inside the runtime (e.g. div by zero, access account)
			GasUsed:       0,
		}
	}
	return &txRes, nil
}

func (bc *blockContextMock) ExecuteScript(
	ledger virtualmachine.Ledger,
	script []byte,
) (*virtualmachine.ScriptResult, error) {
	return nil, nil
}

// virtualMachineMock is a mocked virtualMachine
type virtualMachineMock struct {
}

func newVirtualMachineMock() *virtualMachineMock {
	// TODO set execution outcome
	return &virtualMachineMock{}
}

func (vm *virtualMachineMock) NewBlockContext(header *flow.Header) virtualmachine.BlockContext {
	return &blockContextMock{
		vm:     vm,
		header: header,
	}
}

// func (vm *virtualMachineMock) executeTransaction(
// 	script []byte,
// 	runtimeInterface runtime.Interface,
// 	location runtime.Location,
// ) error {
// 	return nil
// }

// func (vm *virtualMachineMock) executeScript(
// 	script []byte,
// 	runtimeInterface runtime.Interface,
// 	location runtime.Location,
// ) (runtime.Value, error) {
// 	return nil, nil
// }
