package types

import (
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"math/big"
	"testing"
)

func TestBatchHeaderWithBlobHashes(t *testing.T) {
	expectedBatchHeader := BatchHeader{
		Version:                0,
		BatchIndex:             10,
		L1MessagePopped:        5,
		TotalL1MessagePopped:   20,
		DataHash:               common.BigToHash(big.NewInt(100)),
		ParentBatchHash:        common.BigToHash(big.NewInt(200)),
		SkippedL1MessageBitmap: rand.Bytes(10),
	}
	batchHeaderWithBlob := NewBatchHeaderWithBlobHashes(expectedBatchHeader, nil)
	bytes := batchHeaderWithBlob.MarshalBinary()

	decoded := new(BatchHeaderWithBlobHashes)
	err := decoded.UnmarshalBinary(bytes)
	require.NoError(t, err)
	require.EqualValues(t, expectedBatchHeader.Version, decoded.Version)
	require.EqualValues(t, expectedBatchHeader.BatchIndex, decoded.BatchIndex)
	require.EqualValues(t, expectedBatchHeader.L1MessagePopped, decoded.L1MessagePopped)
	require.EqualValues(t, expectedBatchHeader.TotalL1MessagePopped, decoded.TotalL1MessagePopped)
	require.EqualValues(t, expectedBatchHeader.DataHash, decoded.DataHash)
	require.EqualValues(t, expectedBatchHeader.ParentBatchHash, decoded.ParentBatchHash)
	require.EqualValues(t, expectedBatchHeader.SkippedL1MessageBitmap, decoded.SkippedL1MessageBitmap)
	require.Nil(t, decoded.blobHashes)

	expectedBlobHashes := []common.Hash{common.BigToHash(big.NewInt(1)), common.BigToHash(big.NewInt(2))}
	batchHeaderWithBlob = NewBatchHeaderWithBlobHashes(expectedBatchHeader, expectedBlobHashes)
	bytes = batchHeaderWithBlob.MarshalBinary()

	decoded = new(BatchHeaderWithBlobHashes)
	err = decoded.UnmarshalBinary(bytes)
	require.NoError(t, err)
	require.EqualValues(t, expectedBatchHeader.Version, decoded.Version)
	require.EqualValues(t, expectedBatchHeader.BatchIndex, decoded.BatchIndex)
	require.EqualValues(t, expectedBatchHeader.L1MessagePopped, decoded.L1MessagePopped)
	require.EqualValues(t, expectedBatchHeader.TotalL1MessagePopped, decoded.TotalL1MessagePopped)
	require.EqualValues(t, expectedBatchHeader.DataHash, decoded.DataHash)
	require.EqualValues(t, expectedBatchHeader.ParentBatchHash, decoded.ParentBatchHash)
	require.EqualValues(t, expectedBatchHeader.SkippedL1MessageBitmap, decoded.SkippedL1MessageBitmap)
	require.EqualValues(t, 2, len(decoded.blobHashes))
	require.EqualValues(t, expectedBlobHashes[0], decoded.blobHashes[0])
	require.EqualValues(t, expectedBlobHashes[1], decoded.blobHashes[1])
}
