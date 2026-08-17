package batch

import (
	"bytes"
	"context"
	"encoding/binary"
	"math/big"
	"testing"

	"morph-l2/bindings/bindings"
	"morph-l2/common/blob"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"
)

const (
	transitionFinalized = uint64(10)
	transitionCutover   = uint64(12)
)

type transitionRollupReader struct {
	committed      uint64
	finalized      uint64
	cutover        uint64
	committedHash  map[uint64][32]byte
	batchBlockByID map[uint64]uint64
}

func (r *transitionRollupReader) CommittedBatches(_ *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	return r.committedHash[batchIndex.Uint64()], nil
}

func (r *transitionRollupReader) LastCommittedBatchIndex(_ *bind.CallOpts) (*big.Int, error) {
	return new(big.Int).SetUint64(r.committed), nil
}

func (r *transitionRollupReader) LastFinalizedBatchIndex(_ *bind.CallOpts) (*big.Int, error) {
	return new(big.Int).SetUint64(r.finalized), nil
}

func (r *transitionRollupReader) LegacyCutoverBatchIndex(_ *bind.CallOpts) (*big.Int, error) {
	return new(big.Int).SetUint64(r.cutover), nil
}

func (r *transitionRollupReader) BatchDataStore(_ *bind.CallOpts, batchIndex *big.Int) (struct {
	OriginTimestamp   *big.Int
	FinalizeTimestamp *big.Int
	BlockNumber       *big.Int
	Submitter         common.Address
}, error) {
	return struct {
		OriginTimestamp   *big.Int
		FinalizeTimestamp *big.Int
		BlockNumber       *big.Int
		Submitter         common.Address
	}{
		OriginTimestamp:   new(big.Int),
		FinalizeTimestamp: new(big.Int),
		BlockNumber:       new(big.Int).SetUint64(r.batchBlockByID[batchIndex.Uint64()]),
	}, nil
}

func (r *transitionRollupReader) FilterFinalizeBatch(
	_ *bind.FilterOpts,
	_ []*big.Int,
	_ [][32]byte,
) (*bindings.RollupFinalizeBatchIterator, error) {
	return nil, nil
}

func newTransitionBatchCache(kv SealedBatchKV, rollup RollupBatchReader) *BatchCache {
	return &BatchCache{
		ctx:                context.Background(),
		batchStorage:       NewBatchStorage(kv),
		sealedBatches:      make(map[uint64]*eth.RPCRollupBatch),
		sealedBatchHeaders: make(map[uint64]*BatchHeaderBytes),
		batchData:          NewBatchData(),
		isBatchUpgraded:    func(uint64) bool { return true },
		isBatchV2Upgraded:  func(uint64) bool { return false },
		rollupContract:     rollup,
		maxBlobCount:       1,
	}
}

func makeTransitionHeader(
	index uint64,
	parentHash common.Hash,
	sequencerHash common.Hash,
	blobCommitHash common.Hash,
) BatchHeaderBytes {
	return BatchHeaderV1{
		BatchHeaderV0: BatchHeaderV0{
			BatchIndex:             index,
			L1MessagePopped:        1,
			TotalL1MessagePopped:   index + 100,
			DataHash:               common.BigToHash(new(big.Int).SetUint64(index*10 + 1)),
			BlobVersionedHash:      blobCommitHash,
			PrevStateRoot:          common.BigToHash(new(big.Int).SetUint64(index*10 + 3)),
			PostStateRoot:          common.BigToHash(new(big.Int).SetUint64(index*10 + 4)),
			WithdrawalRoot:         common.BigToHash(new(big.Int).SetUint64(index*10 + 5)),
			SequencerSetVerifyHash: sequencerHash,
			ParentBatchHash:        parentHash,
		},
		LastBlockNumber: index * 10,
	}.Bytes()
}

// storeTransitionWindow persists a canonical, contiguous [F,last] snapshot and
// configures the fake rollup's committed hashes through K. Indices above K model
// batches sealed locally but not yet committed at the upgrade boundary.
func storeTransitionWindow(
	t *testing.T,
	kv SealedBatchKV,
	last uint64,
	postCutoverSequencerHash common.Hash,
) (*transitionRollupReader, map[uint64]BatchHeaderBytes) {
	t.Helper()

	rollup := &transitionRollupReader{
		committed:      transitionCutover,
		finalized:      transitionFinalized,
		cutover:        transitionCutover,
		committedHash:  make(map[uint64][32]byte),
		batchBlockByID: make(map[uint64]uint64),
	}
	storage := NewBatchStorage(kv)
	headers := make(map[uint64]BatchHeaderBytes)
	legacySequencerHash := common.HexToHash("0x123456")
	parentHeader := makeTransitionHeader(
		transitionFinalized-1,
		common.Hash{},
		legacySequencerHash,
		common.BigToHash(new(big.Int).SetUint64((transitionFinalized-1)*10+2)),
	)
	parentHash, err := parentHeader.Hash()
	require.NoError(t, err)

	for index := transitionFinalized; index <= last; index++ {
		sequencerHash := legacySequencerHash
		blobCommitHash := common.BigToHash(new(big.Int).SetUint64(index*10 + 2))
		var sidecar ethtypes.BlobTxSidecar
		if index > transitionCutover {
			sequencerHash = postCutoverSequencerHash
			generatedSidecar, sidecarErr := blob.MakeBlobTxSidecar(nil, 1)
			require.NoError(t, sidecarErr)
			sidecar = *generatedSidecar
			blobHashes := sidecar.BlobHashes()
			require.Len(t, blobHashes, 1)
			blobCommitHash = blobHashes[0]
		}
		header := makeTransitionHeader(index, parentHash, sequencerHash, blobCommitHash)
		hash, hashErr := header.Hash()
		require.NoError(t, hashErr)
		batch := &eth.RPCRollupBatch{
			Version:           uint(BatchHeaderVersion1),
			Hash:              hash,
			ParentBatchHeader: hexutil.Bytes(parentHeader),
			PrevStateRoot:     common.BigToHash(new(big.Int).SetUint64(index*10 + 3)),
			PostStateRoot:     common.BigToHash(new(big.Int).SetUint64(index*10 + 4)),
			WithdrawRoot:      common.BigToHash(new(big.Int).SetUint64(index*10 + 5)),
			LastBlockNumber:   index * 10,
			NumL1Messages:     1,
			Sidecar:           sidecar,
		}
		require.NoError(t, storage.StoreSealedBatchAndHeader(index, batch, &header))
		headers[index] = header
		if index <= transitionCutover {
			rollup.committedHash[index] = [32]byte(hash)
		}
		rollup.batchBlockByID[index] = index * 10
		parentHeader = header
		parentHash = hash
	}
	return rollup, headers
}

func requireTransitionStoragePreserved(t *testing.T, kv SealedBatchKV, index uint64) {
	t.Helper()
	_, err := kv.GetBytes(encodeBatchKey(index))
	require.NoError(t, err)
	_, err = kv.GetBytes([]byte(SealedBatchIndicesKey))
	require.NoError(t, err)
}

func mutateStoredTransitionRecord(
	t *testing.T,
	kv SealedBatchKV,
	index uint64,
	mutate func(*eth.RPCRollupBatch, *BatchHeaderBytes),
) {
	t.Helper()
	storage := NewBatchStorage(kv)
	storedBatch, err := storage.LoadSealedBatch(index)
	require.NoError(t, err)
	storedHeader, err := storage.LoadSealedBatchHeader(index)
	require.NoError(t, err)
	mutate(storedBatch, storedHeader)
	storedBatch.Hash, err = storedHeader.Hash()
	require.NoError(t, err)
	require.NoError(t, storage.StoreSealedBatchAndHeader(index, storedBatch, storedHeader))
}

func TestInitAndSyncFromDatabaseLegacyTransitionUsesExactCutoverParent(t *testing.T) {
	kv := openTestKV(t)
	rollup, headers := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	cache := newTransitionBatchCache(kv, rollup)

	require.NoError(t, cache.InitAndSyncFromDatabase())
	require.True(t, cache.isInitialized())
	require.Equal(t, headers[transitionCutover].Bytes(), cache.parentBatchHeader.Bytes())
	for index := transitionFinalized; index <= transitionCutover; index++ {
		_, exists := cache.GetSealedBatchHeader(index)
		require.True(t, exists, "finalizer must retain canonical header %d", index)
	}
	cutoverSequencerHash, err := cache.parentBatchHeader.SequencerSetVerifyHash()
	require.NoError(t, err)
	require.NotEqual(t, common.Hash{}, cutoverSequencerHash)

	cache.batchData.Append(make([]byte, 60), []byte{1, 2, 3}, nil)
	cache.postStateRoot = common.HexToHash("0x2001")
	cache.withdrawRoot = common.HexToHash("0x2002")
	cache.lastPackedBlockHeight++

	index, newHeader, _, err := cache.SealBatch(1, nil)
	require.NoError(t, err)
	require.Equal(t, transitionCutover+1, index)
	newBatch, exists := cache.GetSealedBatch(index)
	require.True(t, exists)
	require.True(t, bytes.Equal(headers[transitionCutover], newBatch.ParentBatchHeader), "K+1 must embed the exact canonical K header")
	sequencerHash, err := newHeader.SequencerSetVerifyHash()
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, sequencerHash, "K+1 must use the post-upgrade zero sequencer hash")
}

func TestInitAndSyncFromDatabaseLegacyTransitionAcceptsAdvancedFinalizedWindow(t *testing.T) {
	kv := openTestKV(t)
	rollup, headers := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	rollup.finalized = transitionFinalized + 1
	require.NoError(t, NewBatchStorage(kv).DeleteSealedBatch(transitionFinalized))
	cache := newTransitionBatchCache(kv, rollup)

	require.NoError(t, cache.InitAndSyncFromDatabase())
	require.Equal(t, headers[transitionCutover].Bytes(), cache.parentBatchHeader.Bytes())
	for index := rollup.finalized; index <= transitionCutover; index++ {
		_, exists := cache.GetSealedBatchHeader(index)
		require.True(t, exists, "finalizer must retain pending canonical header %d", index)
	}
}

func TestInitAndSyncFromDatabaseLegacyTransitionFailsClosedOnUnavailableCache(t *testing.T) {
	t.Run("empty database", func(t *testing.T) {
		kv := openTestKV(t)
		rollup := &transitionRollupReader{
			committed:      transitionCutover,
			finalized:      transitionFinalized,
			cutover:        transitionCutover,
			committedHash:  make(map[uint64][32]byte),
			batchBlockByID: make(map[uint64]uint64),
		}

		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
		keys, scanErr := kv.IteratePrefixKeys([]byte(SealedBatchKeyPrefix))
		require.NoError(t, scanErr)
		require.Empty(t, keys)
	})

	t.Run("corrupt indices", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
		require.NoError(t, kv.PutBytes([]byte(SealedBatchIndicesKey), []byte("not-json")))

		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
		requireTransitionStoragePreserved(t, kv, transitionCutover)
	})

	t.Run("missing canonical header", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
		require.NoError(t, kv.Delete(encodeBatchHeaderKey(transitionCutover-1)))

		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
		requireTransitionStoragePreserved(t, kv, transitionCutover)
	})

	t.Run("explicit rebuild refuses to wipe", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})

		err := newTransitionBatchCache(kv, rollup).DeleteBatchStorageAndInitFromRollup()
		require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
		requireTransitionStoragePreserved(t, kv, transitionCutover)
	})
}

func TestInitAndSyncFromDatabaseLegacyTransitionRejectsCutoverHashMismatch(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	rollup.committedHash[transitionCutover] = [32]byte(common.HexToHash("0xdeadbeef"))

	err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
	require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
	requireTransitionStoragePreserved(t, kv, transitionCutover)
}

func TestInitAndSyncFromDatabaseLegacyTransitionHandlesLocalPostCutoverBatch(t *testing.T) {
	t.Run("zero sequencer hash is accepted", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, headers := storeTransitionWindow(t, kv, transitionCutover+1, common.Hash{})
		cache := newTransitionBatchCache(kv, rollup)

		require.NoError(t, cache.InitAndSyncFromDatabase())
		require.Equal(t, headers[transitionCutover+1].Bytes(), cache.parentBatchHeader.Bytes())
	})

	t.Run("non-zero legacy sequencer hash is rejected", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover+1, common.HexToHash("0xabcdef"))

		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
		requireTransitionStoragePreserved(t, kv, transitionCutover+1)
	})
}

func TestInitAndSyncFromDatabaseLegacyTransitionRejectsMutatedLocalBroadcastFields(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*eth.RPCRollupBatch, *BatchHeaderBytes)
	}{
		{
			name: "L1 message count",
			mutate: func(storedBatch *eth.RPCRollupBatch, _ *BatchHeaderBytes) {
				storedBatch.NumL1Messages++
			},
		},
		{
			name: "cumulative L1 messages",
			mutate: func(_ *eth.RPCRollupBatch, storedHeader *BatchHeaderBytes) {
				total := binary.BigEndian.Uint64((*storedHeader)[17:25])
				binary.BigEndian.PutUint64((*storedHeader)[17:25], total+1)
			},
		},
		{
			name: "single blob commitment",
			mutate: func(storedBatch *eth.RPCRollupBatch, _ *BatchHeaderBytes) {
				storedBatch.Sidecar.Commitments[0][0] ^= 0x01
			},
		},
		{
			name: "V1 multiple blobs",
			mutate: func(storedBatch *eth.RPCRollupBatch, _ *BatchHeaderBytes) {
				storedBatch.Sidecar.Blobs = append(storedBatch.Sidecar.Blobs, storedBatch.Sidecar.Blobs[0])
				storedBatch.Sidecar.Commitments = append(storedBatch.Sidecar.Commitments, storedBatch.Sidecar.Commitments[0])
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			kv := openTestKV(t)
			rollup, _ := storeTransitionWindow(t, kv, transitionCutover+1, common.Hash{})
			mutateStoredTransitionRecord(t, kv, transitionCutover+1, test.mutate)

			err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
			require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
			requireTransitionStoragePreserved(t, kv, transitionCutover+1)
		})
	}
}

func TestInitAndSyncFromDatabaseLegacyTransitionValidatesV2AggregateBlobHash(t *testing.T) {
	makeV2 := func(useAggregate bool) (SealedBatchKV, *transitionRollupReader) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover+1, common.Hash{})
		mutateStoredTransitionRecord(t, kv, transitionCutover+1, func(storedBatch *eth.RPCRollupBatch, storedHeader *BatchHeaderBytes) {
			storedBatch.Sidecar.Blobs = append(storedBatch.Sidecar.Blobs, storedBatch.Sidecar.Blobs[0])
			storedBatch.Sidecar.Commitments = append(storedBatch.Sidecar.Commitments, storedBatch.Sidecar.Commitments[0])
			(*storedHeader)[0] = BatchHeaderVersion2
			storedBatch.Version = uint(BatchHeaderVersion2)
			blobHashes := blob.BlobHashes(storedBatch.Sidecar.Blobs, storedBatch.Sidecar.Commitments)
			commitHash := blobHashes[0]
			if useAggregate {
				commitHash = aggregateBlobHashes(blobHashes)
			}
			copy((*storedHeader)[57:89], commitHash[:])
		})
		return kv, rollup
	}

	t.Run("aggregate is accepted", func(t *testing.T) {
		kv, rollup := makeV2(true)
		require.NoError(t, newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase())
	})

	t.Run("first blob hash is rejected", func(t *testing.T) {
		kv, rollup := makeV2(false)
		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
		requireTransitionStoragePreserved(t, kv, transitionCutover+1)
	})
}

func TestBatchCacheGetFailsClosedBeforeInitialization(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	cache := newTransitionBatchCache(kv, rollup)

	storedBatch, err := NewBatchStorage(kv).LoadSealedBatch(transitionCutover)
	require.NoError(t, err)
	require.NotNil(t, storedBatch, "fixture must contain the batch Get would otherwise load")
	batchFromCache, err := cache.Get(transitionCutover)
	require.ErrorIs(t, err, ErrBatchCacheNotInitialized)
	require.Nil(t, batchFromCache)
}

func TestInitFromRollupLegacyTransitionFailsClosed(t *testing.T) {
	rollup := &transitionRollupReader{
		committed:      transitionCutover,
		finalized:      transitionFinalized,
		cutover:        transitionCutover,
		committedHash:  make(map[uint64][32]byte),
		batchBlockByID: make(map[uint64]uint64),
	}
	err := newTransitionBatchCache(openTestKV(t), rollup).Init()
	require.ErrorIs(t, err, ErrLegacyTransitionCacheRequired)
}
