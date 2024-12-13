package services

import (
	"context"
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"morph-l2/tx-submitter/utils"
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

func TestRoughRollupGasEstimate(t *testing.T) {

	cfg := &utils.Config{
		RollupTxGasBase:       2,
		RollupTxGasPerL1Msg:   3,
		RollupTxGasPerL2Block: 4,
	}
	r := NewRollup(context.Background(), nil, nil, nil, nil, nil, nil, nil, nil, common.Address{}, nil, *cfg, nil, nil, nil, nil, nil)
	estimateGas := r.RoughRollupGasEstimate(0, 0)
	require.EqualValues(t, 2, estimateGas)
	estimateGas = r.RoughRollupGasEstimate(1, 0)
	require.EqualValues(t, 5, estimateGas)
	estimateGas = r.RoughRollupGasEstimate(0, 1)
	require.EqualValues(t, 6, estimateGas)
	estimateGas = r.RoughRollupGasEstimate(1, 1)
	require.EqualValues(t, 9, estimateGas)

	//
	// utils.ParseL1MessageCnt()
}
