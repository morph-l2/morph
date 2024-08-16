package services

import (
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
