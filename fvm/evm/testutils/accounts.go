package testutils

import (
	"bytes"
	"crypto/ecdsa"
	"io"
	"math/big"
	"sync"
	"testing"

	gethCommon "github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	gethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/onflow/atree"

	"github.com/onflow/flow-go/fvm/evm/emulator"
	"github.com/onflow/flow-go/fvm/evm/emulator/database"
	"github.com/onflow/flow-go/fvm/evm/types"
	"github.com/onflow/flow-go/model/flow"
)

// address:  658bdf435d810c91414ec09147daa6db62406379
const EOATestAccount1KeyHex = "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c"

type EOATestAccount struct {
	address gethCommon.Address
	key     *ecdsa.PrivateKey
	nonce   uint64
	signer  gethTypes.Signer
	lock    sync.Mutex
}

func (a *EOATestAccount) Address() types.Address {
	return types.Address(a.address)
}

func (a *EOATestAccount) PrepareSignAndEncodeTx(
	t testing.TB,
	to gethCommon.Address,
	data []byte,
	amount *big.Int,
	gasLimit uint64,
	gasFee *big.Int,
) []byte {
	tx := a.PrepareAndSignTx(t, to, data, amount, gasLimit, gasFee)
	var b bytes.Buffer
	writer := io.Writer(&b)
	tx.EncodeRLP(writer)
	return b.Bytes()
}

func (a *EOATestAccount) PrepareAndSignTx(
	t testing.TB,
	to gethCommon.Address,
	data []byte,
	amount *big.Int,
	gasLimit uint64,
	gasFee *big.Int,
) *gethTypes.Transaction {

	a.lock.Lock()
	defer a.lock.Unlock()

	tx, err := gethTypes.SignTx(
		gethTypes.NewTransaction(
			a.nonce,
			to,
			amount,
			gasLimit,
			gasFee,
			data),
		a.signer, a.key)
	require.NoError(t, err)
	a.nonce++

	return tx
}

func GetTestEOAAccount(t testing.TB, keyHex string) *EOATestAccount {
	key, _ := gethCrypto.HexToECDSA(keyHex)
	address := gethCrypto.PubkeyToAddress(key.PublicKey)
	signer := emulator.GetDefaultSigner()
	return &EOATestAccount{
		address: address,
		key:     key,
		signer:  signer,
		lock:    sync.Mutex{},
	}
}

func RunWithEOATestAccount(t *testing.T, led atree.Ledger, flowEVMRootAddress flow.Address, f func(*EOATestAccount)) {
	account := GetTestEOAAccount(t, EOATestAccount1KeyHex)

	// fund account
	db, err := database.NewDatabase(led, flowEVMRootAddress)
	require.NoError(t, err)

	e := emulator.NewEmulator(db)
	require.NoError(t, err)

	blk, err := e.NewBlockView(types.NewDefaultBlockContext(2))
	require.NoError(t, err)

	_, err = blk.MintTo(account.Address(), new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000)))
	require.NoError(t, err)

	blk2, err := e.NewReadOnlyBlockView(types.NewDefaultBlockContext(2))
	require.NoError(t, err)

	bal, err := blk2.BalanceOf(account.Address())
	require.NoError(t, err)
	require.Greater(t, bal.Uint64(), uint64(0))

	f(account)
}
