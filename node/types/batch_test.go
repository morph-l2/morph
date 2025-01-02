package types

import (
	"math/big"
	"morph-l2/bindings/bindings"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestBatchHeader(t *testing.T) {
	expectedBatchHeaderV0 := BatchHeaderV0{
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
	}
	batchHeaderBytes := BatchHeaderBytes(expectedBatchHeaderV0.Bytes())

	version, err := batchHeaderBytes.Version()
	require.NoError(t, err)
	batchIndex, err := batchHeaderBytes.BatchIndex()
	require.NoError(t, err)
	l1MessagePopped, err := batchHeaderBytes.L1MessagePopped()
	require.NoError(t, err)
	totalL1MessagePopped, err := batchHeaderBytes.TotalL1MessagePopped()
	require.NoError(t, err)
	dataHash, err := batchHeaderBytes.DataHash()
	require.NoError(t, err)
	blobVersionedHash, err := batchHeaderBytes.BlobVersionedHash()
	require.NoError(t, err)
	prevStateRoot, err := batchHeaderBytes.PrevStateRoot()
	require.NoError(t, err)
	postStateRoot, err := batchHeaderBytes.PostStateRoot()
	require.NoError(t, err)
	withdrawalRoot, err := batchHeaderBytes.WithdrawalRoot()
	require.NoError(t, err)
	sequencerSetVerifyHash, err := batchHeaderBytes.SequencerSetVerifyHash()
	require.NoError(t, err)
	parentBatchHash, err := batchHeaderBytes.ParentBatchHash()
	require.NoError(t, err)

	require.EqualValues(t, 0, version)
	require.EqualValues(t, expectedBatchHeaderV0.BatchIndex, batchIndex)
	require.EqualValues(t, expectedBatchHeaderV0.L1MessagePopped, l1MessagePopped)
	require.EqualValues(t, expectedBatchHeaderV0.TotalL1MessagePopped, totalL1MessagePopped)
	require.EqualValues(t, expectedBatchHeaderV0.DataHash, dataHash)
	require.EqualValues(t, expectedBatchHeaderV0.BlobVersionedHash, blobVersionedHash)
	require.EqualValues(t, expectedBatchHeaderV0.PrevStateRoot, prevStateRoot)
	require.EqualValues(t, expectedBatchHeaderV0.PostStateRoot, postStateRoot)
	require.EqualValues(t, expectedBatchHeaderV0.WithdrawalRoot, withdrawalRoot)
	require.EqualValues(t, expectedBatchHeaderV0.SequencerSetVerifyHash, sequencerSetVerifyHash)
	require.EqualValues(t, expectedBatchHeaderV0.ParentBatchHash, parentBatchHash)

	expectedBatchHeaderV1 := BatchHeaderV1{
		BatchHeaderV0:   expectedBatchHeaderV0,
		LastBlockNumber: 1000,
	}.Bytes()
	version, err = expectedBatchHeaderV1.Version()
	require.NoError(t, err)
	lastBlockNumber, err := expectedBatchHeaderV1.LastBlockNumber()
	require.NoError(t, err)
	require.EqualValues(t, 1, version)
	require.EqualValues(t, 1000, lastBlockNumber)
}

func TestMethodID(t *testing.T) {
	beforeSkipABI, err := LegacyRollupMetaData.GetAbi()
	require.NoError(t, err)
	beforeMoveBlockCtxABI, err := BeforeMoveBlockCtxABI.GetAbi()
	require.NoError(t, err)
	currentABI, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	require.NotEqualValues(t, beforeSkipABI.Methods["commitBatch"].ID, beforeMoveBlockCtxABI.Methods["commitBatch"].ID, currentABI.Methods["commitBatch"].ID)
}
