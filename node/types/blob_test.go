package types

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"morph-l2/node/zstd"
	"testing"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
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

func TestDecodeRawTxPayload(t *testing.T) {
	// empty blob
	blob := new(kzg4844.Blob)
	rawTxPayload, err := DecodeRawTxPayload(blob)
	require.NoError(t, err)
	require.EqualValues(t, 0, len(rawTxPayload))

	txsPayload10 := rand.Bytes(10)
	txsPayload27 := rand.Bytes(27)
	txsPayload58 := rand.Bytes(58)
	txsPayload59 := rand.Bytes(59)
	len10 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len10, 10)
	len27 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len27, 27)
	len58 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len58, 58)
	len59 := make([]byte, 4)
	binary.LittleEndian.PutUint32(len59, 59)
	sealedTxPayload := bytes.NewBuffer(make([]byte, 31)) // empty chunk
	sealedTxPayload.Write(len10)                         // 10bytes chunk
	sealedTxPayload.Write(txsPayload10)
	sealedTxPayload.Write(make([]byte, 31-14))
	sealedTxPayload.Write(len27) // 27bytes chunk
	sealedTxPayload.Write(txsPayload27)
	sealedTxPayload.Write(make([]byte, 31)) // empty chunk
	sealedTxPayload.Write(len58)            // 58bytes chunk
	sealedTxPayload.Write(txsPayload58)
	sealedTxPayload.Write(len59) // 59bytes chunk
	sealedTxPayload.Write(txsPayload59)
	sealedTxPayload.Write(make([]byte, 30))

	expectedRawTxPayload := bytes.NewBuffer(txsPayload10)
	expectedRawTxPayload.Write(txsPayload27)
	expectedRawTxPayload.Write(txsPayload58)
	expectedRawTxPayload.Write(txsPayload59)

	blob, err = MakeBlobCanonical(sealedTxPayload.Bytes())
	require.NoError(t, err)
	decodedRawTxPayload, err := DecodeRawTxPayload(blob)
	require.NoError(t, err)
	require.EqualValues(t, expectedRawTxPayload.Bytes(), decodedRawTxPayload)
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
	compressedBlobBytes, err := zstd.CompressBatchBytes(cks.ConstructBlobPayload())
	require.NoError(t, err)
	b, err := MakeBlobCanonical(compressedBlobBytes)
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
	compressedBlobBytes, err := zstd.CompressBatchBytes(cks.ConstructBlobPayload())
	require.NoError(t, err)
	b, err := MakeBlobCanonical(compressedBlobBytes)
	require.NoError(t, err)
	txs, err := DecodeTxsFromBlob(b)
	require.NoError(t, err)
	require.EqualValues(t, 3, txs.Len())
	require.EqualValues(t, transferTx.Hash(), txs[0].Hash())
	require.EqualValues(t, legacyContractTx.Hash(), txs[1].Hash())
	require.EqualValues(t, contractTx.Hash(), txs[2].Hash())
}
