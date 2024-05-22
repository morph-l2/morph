package types

import (
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
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
	require.EqualValues(t, expectedBatchHeader.ParentBatchHash, decoded.ParentBatchHash)
	require.EqualValues(t, expectedBatchHeader.SkippedL1MessageBitmap, decoded.SkippedL1MessageBitmap)
}
