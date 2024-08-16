package types

import (
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestBatchHeader(t *testing.T) {
	expectedBatchHeader := BatchHeader{
		Version:                0,
		BatchIndex:             10,
		L1MessagePopped:        5,
		TotalL1MessagePopped:   20,
		DataHash:               common.BigToHash(big.NewInt(100)),
		BlobVersionedHash:      EmptyVersionedHash,
		PrevStateRoot:          common.BigToHash(big.NewInt(101)),
		PostStateRoot:          common.BigToHash(big.NewInt(102)),
		WithdrawalRoot:         common.BigToHash(big.NewInt(103)),
		SequencerSetVerifyHash: common.BigToHash(big.NewInt(104)),
		ParentBatchHash:        common.BigToHash(big.NewInt(200)),
		SkippedL1MessageBitmap: rand.Bytes(10),
	}
	bytes := expectedBatchHeader.Encode()

	decoded, err := DecodeBatchHeader(bytes)
	require.NoError(t, err)
	require.EqualValues(t, expectedBatchHeader.Version, decoded.Version)
	require.EqualValues(t, expectedBatchHeader.BatchIndex, decoded.BatchIndex)
	require.EqualValues(t, expectedBatchHeader.L1MessagePopped, decoded.L1MessagePopped)
	require.EqualValues(t, expectedBatchHeader.TotalL1MessagePopped, decoded.TotalL1MessagePopped)
	require.EqualValues(t, expectedBatchHeader.DataHash, decoded.DataHash)
	require.EqualValues(t, expectedBatchHeader.BlobVersionedHash, decoded.BlobVersionedHash)
	require.EqualValues(t, expectedBatchHeader.PrevStateRoot, decoded.PrevStateRoot)
	require.EqualValues(t, expectedBatchHeader.PostStateRoot, decoded.PostStateRoot)
	require.EqualValues(t, expectedBatchHeader.WithdrawalRoot, decoded.WithdrawalRoot)
	require.EqualValues(t, expectedBatchHeader.SequencerSetVerifyHash, decoded.SequencerSetVerifyHash)
	require.EqualValues(t, expectedBatchHeader.ParentBatchHash, decoded.ParentBatchHash)
	require.EqualValues(t, expectedBatchHeader.SkippedL1MessageBitmap, decoded.SkippedL1MessageBitmap)
}
