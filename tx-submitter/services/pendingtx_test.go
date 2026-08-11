package services

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/stretchr/testify/require"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/mock"
)

func TestPendingTxsRecoverRejectsNonOwnedCalldata(t *testing.T) {
	current, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)

	selectors := map[string]string{
		"pre submitter commitBatch":          "428868b5",
		"pre submitter commitState":          "1e8825be",
		"pre submitter commitBatchWithProof": "4e8f1d67",
		"unknown":                            "deadbeef",
	}
	selectors["current commitState"] = hex.EncodeToString(current.Methods["commitState"].ID)
	selectors["current commitBatchWithProof"] = hex.EncodeToString(current.Methods["commitBatchWithProof"].ID)

	for name, selector := range selectors {
		t.Run(name, func(t *testing.T) {
			data, err := hex.DecodeString(selector)
			require.NoError(t, err)
			tx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{
				ChainID: big.NewInt(1),
				To:      &common.Address{},
				Data:    data,
			})
			pending := NewPendingTxs(mock.NewMockJournal())
			err = pending.Recover([]*ethtypes.Transaction{tx}, current)
			require.ErrorContains(t, err, "legacy or unknown journal calldata")
			require.Zero(t, pending.Len())
		})
	}
}

func TestPendingTxsRecoverAcceptsCurrentFreshBlobCommit(t *testing.T) {
	current, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	require.Equal(t, "41f756da", hex.EncodeToString(current.Methods["commitBatch"].ID))

	parentHeader := make([]byte, 9)
	binary.BigEndian.PutUint64(parentHeader[1:], 7)
	calldata, err := current.Pack("commitBatch", bindings.IRollupBatchDataInput{
		ParentBatchHeader: parentHeader,
	})
	require.NoError(t, err)

	tx := ethtypes.NewTx(&ethtypes.BlobTx{
		ChainID:    uint256.NewInt(1),
		To:         common.Address{},
		Data:       calldata,
		BlobFeeCap: uint256.NewInt(1),
		BlobHashes: []common.Hash{{1}},
		Sidecar: &ethtypes.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{{1}},
			Commitments: []kzg4844.Commitment{{1}},
		},
	})
	pending := NewPendingTxs(mock.NewMockJournal())
	require.NoError(t, pending.Recover([]*ethtypes.Transaction{tx}, current))
	require.Equal(t, 1, pending.Len())
	require.Equal(t, uint64(8), pending.GetPindex())
}
