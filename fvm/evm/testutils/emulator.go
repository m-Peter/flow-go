package testutils

import (
	"math/big"

	gethCommon "github.com/onflow/go-ethereum/common"

	gethTypes "github.com/onflow/go-ethereum/core/types"

	"github.com/onflow/flow-go/fvm/evm/types"
)

type TestEmulator struct {
	BalanceOfFunc           func(address types.Address) (*big.Int, error)
	NonceOfFunc             func(address types.Address) (uint64, error)
	CodeOfFunc              func(address types.Address) (types.Code, error)
	CodeHashOfFunc          func(address types.Address) ([]byte, error)
	DirectCallFunc          func(call *types.DirectCall) (*types.Result, error)
	RunTransactionFunc      func(tx *gethTypes.Transaction) (*types.Result, error)
	DryRunTransactionFunc   func(tx *gethTypes.Transaction, address gethCommon.Address) (*types.Result, error)
	BatchRunTransactionFunc func(txs []*gethTypes.Transaction) ([]*types.Result, error)
}

var _ types.Emulator = &TestEmulator{}

// NewBlock returns a new block
func (em *TestEmulator) NewBlockView(_ types.BlockContext) (types.BlockView, error) {
	return em, nil
}

// NewBlock returns a new block view
func (em *TestEmulator) NewReadOnlyBlockView(_ types.BlockContext) (types.ReadOnlyBlockView, error) {
	return em, nil
}

// BalanceOf returns the balance of this address
func (em *TestEmulator) BalanceOf(address types.Address) (*big.Int, error) {
	if em.BalanceOfFunc == nil {
		panic("method not set")
	}
	return em.BalanceOfFunc(address)
}

// NonceOfFunc returns the nonce for this address
func (em *TestEmulator) NonceOf(address types.Address) (uint64, error) {
	if em.NonceOfFunc == nil {
		panic("method not set")
	}
	return em.NonceOfFunc(address)
}

// CodeOf returns the code for this address
func (em *TestEmulator) CodeOf(address types.Address) (types.Code, error) {
	if em.CodeOfFunc == nil {
		panic("method not set")
	}
	return em.CodeOfFunc(address)
}

// CodeHashOf returns the code hash for this address
func (em *TestEmulator) CodeHashOf(address types.Address) ([]byte, error) {
	if em.CodeHashOfFunc == nil {
		panic("method not set")
	}
	return em.CodeHashOfFunc(address)
}

// DirectCall executes a direct call
func (em *TestEmulator) DirectCall(call *types.DirectCall) (*types.Result, error) {
	if em.DirectCallFunc == nil {
		panic("method not set")
	}
	return em.DirectCallFunc(call)
}

// RunTransaction runs a transaction and collect gas fees to the coinbase account
func (em *TestEmulator) RunTransaction(tx *gethTypes.Transaction) (*types.Result, error) {
	if em.RunTransactionFunc == nil {
		panic("method not set")
	}
	return em.RunTransactionFunc(tx)
}

// BatchRunTransactions batch runs transactions and collect gas fees to the coinbase account
func (em *TestEmulator) BatchRunTransactions(txs []*gethTypes.Transaction) ([]*types.Result, error) {
	if em.BatchRunTransactionFunc == nil {
		panic("method not set")
	}
	return em.BatchRunTransactionFunc(txs)
}

// DryRunTransaction simulates transaction execution
func (em *TestEmulator) DryRunTransaction(tx *gethTypes.Transaction, address gethCommon.Address) (*types.Result, error) {
	if em.DryRunTransactionFunc == nil {
		panic("method not set")
	}
	return em.DryRunTransactionFunc(tx, address)
}
