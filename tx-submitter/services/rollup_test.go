package services

import (
	"context"
	"math/big"
	"testing"

	"morph-l2/tx-submitter/mock"
	"morph-l2/tx-submitter/utils"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestSendTx(t *testing.T) {
	// Create a new dynamic fee transaction
	_tx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     0,
		GasTipCap: big.NewInt(1000000000),
		GasFeeCap: big.NewInt(1000000000),
		Gas:       1000000,
		To:        &common.Address{},
		Value:     big.NewInt(0),
		Data:      []byte{},
	}
	// new tx
	tx := types.NewTx(_tx)

	// required err
	err := sendTx(nil, 1, tx)
	require.ErrorContains(t, err, utils.ErrExceedFeeLimit.Error())

	// blob tx
	blobTx := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(big.NewInt(1)),
		Nonce:      0,
		GasTipCap:  uint256.MustFromBig(big.NewInt(1000000000)),
		GasFeeCap:  uint256.MustFromBig(big.NewInt(1000000000)),
		Gas:        1000000,
		To:         common.Address{},
		Data:       []byte{},
		BlobFeeCap: uint256.MustFromBig(big.NewInt(1000000000)),
		BlobHashes: []common.Hash{
			{},
		},
	},
	)
	err = sendTx(nil, 1, blobTx)
	require.ErrorContains(t, err, utils.ErrExceedFeeLimit.Error())
}

func TestGetGasTipAndCap(t *testing.T) {
	l1Mock := mock.NewL1ClientWrapper()
	initTip := big.NewInt(1e9)

	baseFee := big.NewInt(1e9)
	excessBlobGas := uint64(1)
	block := types.NewBlockWithHeader(
		&types.Header{
			BaseFee:       baseFee,
			ExcessBlobGas: &excessBlobGas,
		},
	)
	l1Mock.TipCap = initTip
	l1Mock.Block = block
	config := utils.Config{
		MaxTip:     10e9,
		MaxBaseFee: 100e9,
		MinTip:     1e9,
		TipFeeBump: 100,
	}
	r := NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, config, nil, nil, nil, nil, nil)
	tip, feecap, blobfee, err := r.GetGasTipAndCap()
	require.NoError(t, err)
	require.NotNil(t, tip)
	require.NotNil(t, feecap)
	require.NotNil(t, blobfee)
	require.Equal(t, initTip, tip)

	config = utils.Config{
		MaxTip:     10e9,
		MaxBaseFee: 100e9,
		MinTip:     1e9,
		TipFeeBump: 200,
	}
	r = NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, config, nil, nil, nil, nil, nil)
	tip, feecap, blobfee, err = r.GetGasTipAndCap()
	require.NoError(t, err)
	require.NotNil(t, tip)
	require.NotNil(t, feecap)
	require.NotNil(t, blobfee)
	require.Equal(t, tip, initTip.Mul(initTip, big.NewInt(2)))

	config = utils.Config{
		MaxTip:     10e9,
		MaxBaseFee: baseFee.Uint64() - 1,
		MinTip:     1e9,
		TipFeeBump: 200,
	}
	r = NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, config, nil, nil, nil, nil, nil)
	_, _, _, err = r.GetGasTipAndCap()
	require.ErrorContains(t, err, "base fee is too high")

	config = utils.Config{
		MaxTip:     initTip.Uint64() - 1,
		MaxBaseFee: 100e9,
		MinTip:     1e9,
		TipFeeBump: 200,
	}
	r = NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, config, nil, nil, nil, nil, nil)
	_, _, _, err = r.GetGasTipAndCap()
	require.ErrorContains(t, err, "tip is too high")

}

func TestReSubmitTx(t *testing.T) {
	l1Mock := mock.NewL1ClientWrapper()
	initTip := big.NewInt(1e9)

	baseFee := big.NewInt(1e9)
	excessBlobGas := uint64(1)
	block := types.NewBlockWithHeader(
		&types.Header{
			BaseFee:       baseFee,
			ExcessBlobGas: &excessBlobGas,
		},
	)
	l1Mock.TipCap = initTip
	l1Mock.Block = block
	config := utils.Config{
		MaxTip:     10e12,
		MaxBaseFee: 100e9,
		MinTip:     1e10,
		TipFeeBump: 100,
	}

	priv, err := crypto.GenerateKey()
	require.NoError(t, err)

	r := NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, big.NewInt(1), priv, common.Address{}, nil, config, nil, nil, nil, nil, nil)
	_, err = r.ReSubmitTx(false, nil)
	require.ErrorContains(t, err, "nil tx")
	oldTx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: initTip,
	})
	tx, err := r.ReSubmitTx(false, oldTx)
	require.NoError(t, err)
	require.EqualValues(t, config.MinTip, tx.GasTipCap().Uint64())

}

func TestCancelTx(t *testing.T) {
	// Setup mock L1 client
	l1Mock := mock.NewL1ClientWrapper()
	initTip := big.NewInt(1e9)
	baseFee := big.NewInt(1e9)
	excessBlobGas := uint64(1)
	block := types.NewBlockWithHeader(
		&types.Header{
			BaseFee:       baseFee,
			ExcessBlobGas: &excessBlobGas,
		},
	)
	l1Mock.TipCap = initTip
	l1Mock.Block = block

	// Setup config
	config := utils.Config{
		MaxTip:     10e12,
		MaxBaseFee: 100e9,
		MinTip:     1e9,
		TipFeeBump: 120, // 20% bump
	}

	// Setup private key
	priv, err := crypto.GenerateKey()
	require.NoError(t, err)

	// Create rollup instance
	r := NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, big.NewInt(1), priv, common.Address{}, nil, config, nil, nil, nil, nil, nil)

	// Test 1: Cancel nil transaction
	_, err = r.CancelTx(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "nil tx")

	// Test 2: Cancel DynamicFeeTx
	originalDynamicTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     1,
		GasTipCap: big.NewInt(2e9),  // 2 Gwei
		GasFeeCap: big.NewInt(10e9), // 10 Gwei
		Gas:       100000,
		To:        &common.Address{},
		Value:     big.NewInt(0),
		Data:      []byte{1, 2, 3, 4}, // Some non-empty data
	})

	cancelTx, err := r.CancelTx(originalDynamicTx)
	require.NoError(t, err)
	require.NotNil(t, cancelTx)

	// Verify the cancel transaction
	require.Equal(t, originalDynamicTx.Nonce(), cancelTx.Nonce())
	require.Equal(t, originalDynamicTx.Gas(), cancelTx.Gas())
	require.Equal(t, 0, len(cancelTx.Data()))
	require.True(t, cancelTx.GasTipCap().Cmp(originalDynamicTx.GasTipCap()) >= 0)
	require.True(t, cancelTx.GasFeeCap().Cmp(originalDynamicTx.GasFeeCap()) >= 0)

	// Test 3: Cancel BlobTx
	blobTx := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(big.NewInt(1)),
		Nonce:      2,
		GasTipCap:  uint256.MustFromBig(big.NewInt(2e9)),
		GasFeeCap:  uint256.MustFromBig(big.NewInt(10e9)),
		Gas:        200000,
		To:         common.Address{},
		Value:      uint256.NewInt(0),
		Data:       []byte{1, 2, 3, 4},
		BlobFeeCap: uint256.MustFromBig(big.NewInt(5e9)),
		BlobHashes: []common.Hash{{1}},
		Sidecar: &types.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{{1}},
			Commitments: []kzg4844.Commitment{{1}},
			Proofs:      []kzg4844.Proof{{1}},
		},
	})

	cancelBlobTx, err := r.CancelTx(blobTx)
	require.NoError(t, err)
	require.NotNil(t, cancelBlobTx)

	// Verify the cancel blob transaction
	require.Equal(t, blobTx.Nonce(), cancelBlobTx.Nonce())
	require.Equal(t, blobTx.Gas(), cancelBlobTx.Gas())
	require.Equal(t, 0, len(cancelBlobTx.Data()))
	require.True(t, cancelBlobTx.GasTipCap().Cmp(blobTx.GasTipCap()) >= 0)
	require.True(t, cancelBlobTx.GasFeeCap().Cmp(blobTx.GasFeeCap()) >= 0)
	require.True(t, cancelBlobTx.BlobGasFeeCap().Cmp(blobTx.BlobGasFeeCap()) >= 0)
	require.Equal(t, 1, len(cancelBlobTx.BlobHashes()))
	require.Equal(t, 1, len(cancelBlobTx.BlobTxSidecar().Blobs))
}
