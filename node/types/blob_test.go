package types

import (
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
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

func generateSetCodeTx() *eth.Transaction {
	privKey, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	to := common.BigToAddress(big.NewInt(100))
	data := rand.Bytes(100)
	inner := &eth.SetCodeTx{
		ChainID:   uint256.NewInt(2810),
		Nonce:     1,
		GasFeeCap: uint256.NewInt(1e10),
		GasTipCap: uint256.NewInt(1e8),
		Gas:       500000,
		To:        to,
		Value:     uint256.NewInt(1),
		Data:      data,
		AccessList: []eth.AccessTuple{{
			Address:     address,
			StorageKeys: []common.Hash{common.BigToHash(big.NewInt(2))},
		}},
		AuthList: []eth.SetCodeAuthorization{},
	}
	return eth.NewTx(inner)

}

func TestDecodeTxsFromBytes(t *testing.T) {
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

	setCodeTx := generateSetCodeTx()
	require.NoError(t, err)
	setCodeTxBz, err := setCodeTx.MarshalBinary()
	require.NoError(t, err)

	cks := BatchData{
		txsPayload: append(append(append(transferTxBz, legacyContractTxBz...), contractTxBz...), setCodeTxBz...),
	}
	txs, err := DecodeTxsFromBytes(cks.TxsPayload())
	require.NoError(t, err)
	require.EqualValues(t, 4, txs.Len())
	require.EqualValues(t, transferTx.Hash(), txs[0].Hash())
	require.EqualValues(t, legacyContractTx.Hash(), txs[1].Hash())
	require.EqualValues(t, contractTx.Hash(), txs[2].Hash())
}

func TestDecodeTxsFromBytes_MorphTxV0(t *testing.T) {
	morphV0Bytes := common.FromHex("0x7ff8b7820b02820374835a527f8378d6ff830186a094cfb1186f4e93d60e60a8bdd997427d1f33bc372b80b844a9059cbb000000000000000000000000b055051fb2889be5e9831524f1624941299c49bb0000000000000000000000000000000000000000000000000000000000000064c0068398968080a0953c962c4a4583dadc0ff338166d9f1176a6403a3689b7edcedf583ca401c4cba06ab5c3ef27fa3a8966ae61d994cb658c0d809434ed33dc35f90f9300fcc000c8")
	txs, err := DecodeTxsFromBytes(morphV0Bytes)
	require.NoError(t, err)
	require.EqualValues(t, 1, txs.Len())
	require.EqualValues(t, eth.MorphTxType, txs[0].Type())
}

func TestDecodeTxsFromBytes_MorphTxV1(t *testing.T) {
	morphV1Bytes := common.FromHex("0x7f01f8a1820b5e808326ff9b8345841b8252089425db2115628f08d952e4aacf06b341c8bc04a7f28080c00480a0deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef936d6f72706820686f6f6469207465737420747801a07c085d4f1dafac10ee14bf91dc549c6fb0b7aad797569881aa78f4461eb92903a070038e7e5551422168f4745b90854affd790c0f3a0886a926e53127ea3d0be65")
	txs, err := DecodeTxsFromBytes(morphV1Bytes)
	require.NoError(t, err)
	require.EqualValues(t, 1, txs.Len())
	require.EqualValues(t, eth.MorphTxType, txs[0].Type())
}

func TestDecodeTxsFromBytes_MixedWithMorphTx(t *testing.T) {
	morphV0Bytes := common.FromHex("7ff8b7820b02820374835a527f8378d6ff830186a094cfb1186f4e93d60e60a8bdd997427d1f33bc372b80b844a9059cbb000000000000000000000000b055051fb2889be5e9831524f1624941299c49bb0000000000000000000000000000000000000000000000000000000000000064c0068398968080a0953c962c4a4583dadc0ff338166d9f1176a6403a3689b7edcedf583ca401c4cba06ab5c3ef27fa3a8966ae61d994cb658c0d809434ed33dc35f90f9300fcc000c8")
	morphV1Bytes := common.FromHex("7f01f8a1820b5e808326ff9b8345841b8252089425db2115628f08d952e4aacf06b341c8bc04a7f28080c00480a0deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef936d6f72706820686f6f6469207465737420747801a07c085d4f1dafac10ee14bf91dc549c6fb0b7aad797569881aa78f4461eb92903a070038e7e5551422168f4745b90854affd790c0f3a0886a926e53127ea3d0be65")

	// Generate other tx types
	transferTx, err := generateTransferTx(false)
	require.NoError(t, err)
	transferTxBz, err := transferTx.MarshalBinary()
	require.NoError(t, err)

	legacyTx, err := generateContractTx(true)
	require.NoError(t, err)
	legacyTxBz, err := legacyTx.MarshalBinary()
	require.NoError(t, err)

	// Concatenate: DynamicFee + MorphV0 + Legacy + MorphV1
	var txsBytes []byte
	txsBytes = append(txsBytes, transferTxBz...)
	txsBytes = append(txsBytes, morphV0Bytes...)
	txsBytes = append(txsBytes, legacyTxBz...)
	txsBytes = append(txsBytes, morphV1Bytes...)

	txs, err := DecodeTxsFromBytes(txsBytes)
	require.NoError(t, err)
	require.EqualValues(t, 4, txs.Len())
	require.EqualValues(t, eth.DynamicFeeTxType, txs[0].Type())
	require.EqualValues(t, transferTx.Hash(), txs[0].Hash())
	require.EqualValues(t, eth.MorphTxType, txs[1].Type())
	require.EqualValues(t, eth.LegacyTxType, txs[2].Type())
	require.EqualValues(t, legacyTx.Hash(), txs[2].Hash())
	require.EqualValues(t, eth.MorphTxType, txs[3].Type())
}
