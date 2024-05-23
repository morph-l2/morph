package types

import (
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	eth "github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestBlobFromSealedTxPayload(t *testing.T) {
	sealedTxPayload := rand.Bytes(31)
	_, err := MakeBlobCanonical(sealedTxPayload)
	require.NoError(t, err)

	sealedTxPayload = rand.Bytes(MaxBlobBytesSize)
	_, err = MakeBlobCanonical(sealedTxPayload)
	require.NoError(t, err)

	sealedTxPayload = make([]byte, 0)
	blob, err := MakeBlobCanonical(sealedTxPayload)
	require.NoError(t, err)
	require.EqualValues(t, make([]byte, 4096*32), blob)

	blob, err = MakeBlobCanonical(nil)
	require.NoError(t, err)
	require.EqualValues(t, make([]byte, 4096*32), blob)
}

func generateTransferTx(isLegacy bool) (*eth.Transaction, error) {
	privKey, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(2810))
	to := common.BigToAddress(big.NewInt(100))
	var inner eth.TxData
	if isLegacy {
		inner = &eth.LegacyTx{
			Nonce:    1,
			GasPrice: big.NewInt(1),
			Gas:      21000,
			To:       &to,
			Value:    big.NewInt(1),
		}
	} else {
		inner = &eth.DynamicFeeTx{
			ChainID:   big.NewInt(2810),
			Nonce:     1,
			GasFeeCap: big.NewInt(1),
			GasTipCap: big.NewInt(1),
			Gas:       21000,
			To:        &to,
			Value:     big.NewInt(1),
		}
	}
	transferTx := eth.NewTx(inner)
	return auth.Signer(address, transferTx)
}

func generateContractTx(isLegacy bool) (*eth.Transaction, error) {
	privKey, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(2810))
	to := common.BigToAddress(big.NewInt(100))
	data := rand.Bytes(100)
	var inner eth.TxData
	if isLegacy {
		inner = &eth.LegacyTx{
			Nonce:    1,
			GasPrice: big.NewInt(1e10),
			Gas:      500000,
			To:       &to,
			Value:    big.NewInt(1),
			Data:     data,
		}
	} else {
		inner = &eth.DynamicFeeTx{
			ChainID:   big.NewInt(2810),
			Nonce:     1,
			GasFeeCap: big.NewInt(1e10),
			GasTipCap: big.NewInt(1e8),
			Gas:       500000,
			To:        &to,
			Value:     big.NewInt(1),
			Data:      data,
			AccessList: []eth.AccessTuple{{
				Address:     address,
				StorageKeys: []common.Hash{common.BigToHash(big.NewInt(2))},
			}},
		}
	}
	contractTx := eth.NewTx(inner)
	return auth.Signer(address, contractTx)
}

func TestDecodeLegacyTxsFromBlob(t *testing.T) {
	legacyTransferTx, err := generateTransferTx(true)
	require.NoError(t, err)
	legacyTransferTxBz, err := legacyTransferTx.MarshalBinary()
	require.NoError(t, err)

	legacyContractTx, err := generateContractTx(true)
	require.NoError(t, err)
	legacyContractTxBz, err := legacyContractTx.MarshalBinary()
	require.NoError(t, err)

	cks := Chunks{
		data: []*Chunk{
			{txsPayload: legacyTransferTxBz}, {}, {txsPayload: legacyContractTxBz},
		},
	}
	b, err := MakeBlobCanonical(cks.ConstructBlobPayload())
	require.NoError(t, err)
	txs, err := DecodeTxsFromBlob(b)
	require.NoError(t, err)
	require.EqualValues(t, 2, txs.Len())
	require.EqualValues(t, legacyTransferTx.Hash(), txs[0].Hash())
	require.EqualValues(t, legacyContractTx.Hash(), txs[1].Hash())
}

func TestDecodeTxsFromBlob(t *testing.T) {
	transferTx, err := generateTransferTx(false)
	require.NoError(t, err)
	transferTxBz, err := transferTx.MarshalBinary()
	require.NoError(t, err)

	legacyContractTx, err := generateContractTx(true)
	require.NoError(t, err)
	legacyContractTxBz, err := legacyContractTx.MarshalBinary()
	require.NoError(t, err)

	contractTx, err := generateContractTx(false)
	require.NoError(t, err)
	contractTxBz, err := contractTx.MarshalBinary()
	require.NoError(t, err)

	cks := Chunks{
		data: []*Chunk{
			{txsPayload: transferTxBz}, {}, {txsPayload: legacyContractTxBz}, {}, {txsPayload: contractTxBz},
		},
	}
	b, err := MakeBlobCanonical(cks.ConstructBlobPayload())
	require.NoError(t, err)
	txs, err := DecodeTxsFromBlob(b)
	require.NoError(t, err)
	require.EqualValues(t, 3, txs.Len())
	require.EqualValues(t, transferTx.Hash(), txs[0].Hash())
	require.EqualValues(t, legacyContractTx.Hash(), txs[1].Hash())
	require.EqualValues(t, contractTx.Hash(), txs[2].Hash())
}

func TestCalculateChunkNum(t *testing.T) {
	chunkNum := CalculateChunkNum(0)
	require.EqualValues(t, 15, chunkNum)
	chunkNum = CalculateChunkNum(10)
	require.EqualValues(t, 15, chunkNum)
	chunkNum = CalculateChunkNum(15)
	require.EqualValues(t, 15, chunkNum)

	chunkNum = CalculateChunkNum(16)
	require.EqualValues(t, 20, chunkNum)
	chunkNum = CalculateChunkNum(18)
	require.EqualValues(t, 20, chunkNum)
	chunkNum = CalculateChunkNum(20)
	require.EqualValues(t, 20, chunkNum)

	chunkNum = CalculateChunkNum(21)
	require.EqualValues(t, 25, chunkNum)
	chunkNum = CalculateChunkNum(25)
	require.EqualValues(t, 25, chunkNum)
	chunkNum = CalculateChunkNum(26)
	require.EqualValues(t, 30, chunkNum)
}
