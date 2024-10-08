package testutils

import (
	cryptoRand "crypto/rand"
	"math/big"
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
	gethCommon "github.com/onflow/go-ethereum/common"
	gethTypes "github.com/onflow/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/fvm/evm/types"
)

func RandomCommonHash(t testing.TB) gethCommon.Hash {
	ret := gethCommon.Hash{}
	_, err := cryptoRand.Read(ret[:gethCommon.HashLength])
	require.NoError(t, err)
	return ret
}

func RandomUint256Int(limit int64) *uint256.Int {
	return uint256.NewInt(uint64(rand.Int63n(limit) + 1))
}

func RandomBigInt(limit int64) *big.Int {
	return big.NewInt(rand.Int63n(limit) + 1)
}

func RandomAddress(t testing.TB) types.Address {
	return types.NewAddress(RandomCommonAddress(t))
}

func RandomCommonAddress(t testing.TB) gethCommon.Address {
	ret := gethCommon.Address{}
	_, err := cryptoRand.Read(ret[:gethCommon.AddressLength])
	require.NoError(t, err)
	return ret
}

func RandomGas(limit int64) uint64 {
	return uint64(rand.Int63n(limit) + 1)
}

func RandomData(t testing.TB) []byte {
	// byte size [1, 100]
	size := rand.Intn(100) + 1
	ret := make([]byte, size)
	_, err := cryptoRand.Read(ret[:])
	require.NoError(t, err)
	return ret
}

func GetRandomLogFixture(t testing.TB) *gethTypes.Log {
	return &gethTypes.Log{
		Address: RandomCommonAddress(t),
		Topics: []gethCommon.Hash{
			RandomCommonHash(t),
			RandomCommonHash(t),
		},
		Data: RandomData(t),
	}
}

func COAOwnershipProofFixture(t testing.TB) *types.COAOwnershipProof {
	return &types.COAOwnershipProof{
		Address:        types.FlowAddress{1, 2, 3},
		CapabilityPath: "path",
		KeyIndices:     types.KeyIndices{1, 2},
		Signatures: types.Signatures{
			types.Signature("sig1"),
			types.Signature("sig2"),
		},
	}
}

func COAOwnershipProofInContextFixture(t testing.TB) *types.COAOwnershipProofInContext {
	signedMsg := RandomCommonHash(t)
	return &types.COAOwnershipProofInContext{
		COAOwnershipProof: *COAOwnershipProofFixture(t),
		SignedData:        types.SignedData(signedMsg[:]),
		EVMAddress:        RandomAddress(t),
	}
}

func RandomResultFixture(t testing.TB) *types.Result {
	contractAddress := RandomAddress(t)
	return &types.Result{
		Index:                   1,
		TxType:                  1,
		TxHash:                  RandomCommonHash(t),
		ReturnedData:            RandomData(t),
		GasConsumed:             RandomGas(1000),
		DeployedContractAddress: &contractAddress,
		Logs: []*gethTypes.Log{
			GetRandomLogFixture(t),
			GetRandomLogFixture(t),
		},
	}
}

func AggregatedPrecompiledCallsFixture(t testing.TB) types.AggregatedPrecompiledCalls {
	return types.AggregatedPrecompiledCalls{
		types.PrecompiledCalls{
			Address:          RandomAddress(t),
			RequiredGasCalls: []uint64{2},
			RunCalls: []types.RunCall{
				{
					Output: RandomData(t),
				},
				{
					Output:   []byte{},
					ErrorMsg: "Some error msg",
				},
			},
		},
	}
}
