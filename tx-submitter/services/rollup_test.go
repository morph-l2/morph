package services

import (
	"context"
	"math/big"
	"testing"

	"morph-l2/tx-submitter/config"
	"morph-l2/tx-submitter/mock"
	"morph-l2/tx-submitter/utils"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"

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
	cfg := config.Config{
		MaxTip:     10e9,
		MaxBaseFee: 100e9,
		MinTip:     1e9,
		TipFeeBump: 100,
	}
	r := NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, cfg, nil, nil, nil, nil, nil)
	tip, feecap, blobfee, err := r.GetGasTipAndCap()
	require.NoError(t, err)
	require.NotNil(t, tip)
	require.NotNil(t, feecap)
	require.NotNil(t, blobfee)
	require.Equal(t, initTip, tip)

	cfg = config.Config{
		MaxTip:     10e9,
		MaxBaseFee: 100e9,
		MinTip:     1e9,
		TipFeeBump: 200,
	}
	r = NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, cfg, nil, nil, nil, nil, nil)
	tip, feecap, blobfee, err = r.GetGasTipAndCap()
	require.NoError(t, err)
	require.NotNil(t, tip)
	require.NotNil(t, feecap)
	require.NotNil(t, blobfee)
	require.Equal(t, tip, initTip.Mul(initTip, big.NewInt(2)))

	cfg = config.Config{
		MaxTip:     10e9,
		MaxBaseFee: baseFee.Uint64() - 1,
		MinTip:     1e9,
		TipFeeBump: 200,
	}
	r = NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, cfg, nil, nil, nil, nil, nil)
	_, _, _, err = r.GetGasTipAndCap()
	require.ErrorContains(t, err, "base fee is too high")

	cfg = config.Config{
		MaxTip:     initTip.Uint64() - 1,
		MaxBaseFee: 100e9,
		MinTip:     1e9,
		TipFeeBump: 200,
	}
	r = NewRollup(context.Background(), nil, nil, l1Mock, nil, nil, nil, nil, nil, common.Address{}, nil, cfg, nil, nil, nil, nil, nil)
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
	config := config.Config{
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
