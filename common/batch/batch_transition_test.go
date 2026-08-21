package batch

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"testing"

	"morph-l2/bindings/bindings"
	"morph-l2/common/blob"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"
)

const (
	transitionFinalized = uint64(10)
	transitionCutover   = uint64(12)
)

type transitionRollupReader struct {
	committed           uint64
	committedReads      int
	committedSequence   []uint64
	finalized           uint64
	cutover             uint64
	committedHash       map[uint64][32]byte
	committedHashSeries map[uint64][][32]byte
	batchBlockByID      map[uint64]uint64
}

type transitionReplayL2Client struct {
	blocks map[uint64]*ethtypes.Block
}

func (c *transitionReplayL2Client) BlockNumber(context.Context) (uint64, error) {
	var latest uint64
	for number := range c.blocks {
		if number > latest {
			latest = number
		}
	}
	return latest, nil
}

func (c *transitionReplayL2Client) BlockByNumber(_ context.Context, number *big.Int) (*ethtypes.Block, error) {
	block := c.blocks[number.Uint64()]
	if block == nil {
		return nil, fmt.Errorf("test L2 block %d not found", number.Uint64())
	}
	return block, nil
}

func (c *transitionReplayL2Client) Len() int { return 1 }

type transitionReplayL2Gov struct {
	roots               map[uint64]common.Hash
	sequencerSetByBlock map[uint64][]byte
	sequencerCalls      []uint64
}

func (g *transitionReplayL2Gov) GetTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	return [32]byte(g.roots[opts.BlockNumber.Uint64()]), nil
}

func (g *transitionReplayL2Gov) GetSequencerSetBytes(opts *bind.CallOpts) ([]byte, common.Hash, error) {
	blockNumber := opts.BlockNumber.Uint64()
	g.sequencerCalls = append(g.sequencerCalls, blockNumber)
	setBytes := append([]byte(nil), g.sequencerSetByBlock[blockNumber]...)
	return setBytes, crypto.Keccak256Hash(setBytes), nil
}

func (r *transitionRollupReader) CommittedBatches(_ *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	index := batchIndex.Uint64()
	if series := r.committedHashSeries[index]; len(series) != 0 {
		result := series[0]
		r.committedHashSeries[index] = series[1:]
		return result, nil
	}
	return r.committedHash[index], nil
}

func (r *transitionRollupReader) LastCommittedBatchIndex(_ *bind.CallOpts) (*big.Int, error) {
	if len(r.committedSequence) != 0 {
		position := r.committedReads
		if position >= len(r.committedSequence) {
			position = len(r.committedSequence) - 1
		}
		r.committedReads++
		return new(big.Int).SetUint64(r.committedSequence[position]), nil
	}
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

func requireTransitionStorageCleared(t *testing.T, kv SealedBatchKV) {
	t.Helper()
	keys, err := kv.IteratePrefixKeys([]byte(SealedBatchKeyPrefix))
	require.NoError(t, err)
	require.Empty(t, keys)
}

func requireTransitionRecoveryAttempted(t *testing.T, err error) {
	t.Helper()
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrLegacyTransitionCacheRequired)
	require.ErrorContains(t, err, "nil L1 client")
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

func TestInitAndSyncFromDatabaseLegacyTransitionRebuildsUnavailableCache(t *testing.T) {
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
		requireTransitionRecoveryAttempted(t, err)
		requireTransitionStorageCleared(t, kv)
	})

	t.Run("corrupt indices", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
		require.NoError(t, kv.PutBytes([]byte(SealedBatchIndicesKey), []byte("not-json")))

		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		requireTransitionRecoveryAttempted(t, err)
		requireTransitionStorageCleared(t, kv)
	})

	t.Run("missing canonical header", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
		require.NoError(t, kv.Delete(encodeBatchHeaderKey(transitionCutover-1)))

		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		requireTransitionRecoveryAttempted(t, err)
		requireTransitionStorageCleared(t, kv)
	})

	t.Run("explicit rebuild wipes before replay", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})

		err := newTransitionBatchCache(kv, rollup).DeleteBatchStorageAndInitFromRollup()
		requireTransitionRecoveryAttempted(t, err)
		requireTransitionStorageCleared(t, kv)
	})
}

func TestInitAndSyncFromDatabaseLegacyTransitionRejectsCutoverHashMismatch(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	rollup.committedHash[transitionCutover] = [32]byte(common.HexToHash("0xdeadbeef"))

	err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
	requireTransitionRecoveryAttempted(t, err)
	requireTransitionStorageCleared(t, kv)
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
		requireTransitionRecoveryAttempted(t, err)
		requireTransitionStorageCleared(t, kv)
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
			name: "blob payload does not match commitment",
			mutate: func(storedBatch *eth.RPCRollupBatch, _ *BatchHeaderBytes) {
				// Keep the field element canonical while changing the blob opening.
				storedBatch.Sidecar.Blobs[0][31] ^= 0x01
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
			requireTransitionRecoveryAttempted(t, err)
			requireTransitionStorageCleared(t, kv)
		})
	}
}

func TestInitAndSyncFromDatabasePostTransitionRejectsInvalidLocalBatch(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover+1, common.Hash{})
	rollup.finalized = transitionCutover
	mutateStoredTransitionRecord(t, kv, transitionCutover+1, func(batch *eth.RPCRollupBatch, _ *BatchHeaderBytes) {
		batch.NumL1Messages++
	})

	err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
	requireTransitionRecoveryAttempted(t, err)
	requireTransitionStorageCleared(t, kv)
}

func TestInitAndSyncFromDatabasePostTransitionRejectsNonZeroCommittedField(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover+1, common.HexToHash("0xabcdef"))
	rollup.finalized = transitionCutover
	rollup.committed = transitionCutover + 1
	committedBatch, err := NewBatchStorage(kv).LoadSealedBatch(transitionCutover + 1)
	require.NoError(t, err)
	rollup.committedHash[transitionCutover+1] = [32]byte(committedBatch.Hash)

	err = newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
	requireTransitionRecoveryAttempted(t, err)
	requireTransitionStorageCleared(t, kv)
}

func TestInitAndSyncFromDatabaseRejectsOrphanKeysAndGettersTrustOnlyValidatedMemory(t *testing.T) {
	writeOrphan := func(t *testing.T, kv SealedBatchKV, source, orphan uint64) {
		t.Helper()
		batchBytes, err := kv.GetBytes(encodeBatchKey(source))
		require.NoError(t, err)
		headerBytes, err := kv.GetBytes(encodeBatchHeaderKey(source))
		require.NoError(t, err)
		require.NoError(t, kv.PutBytes(encodeBatchKey(orphan), batchBytes))
		require.NoError(t, kv.PutBytes(encodeBatchHeaderKey(orphan), headerBytes))
	}

	t.Run("startup rejects orphan key", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
		writeOrphan(t, kv, transitionCutover, transitionCutover+100)

		err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
		requireTransitionRecoveryAttempted(t, err)
		requireTransitionStorageCleared(t, kv)
	})

	t.Run("runtime getters ignore raw orphan key", func(t *testing.T) {
		kv := openTestKV(t)
		rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
		cache := newTransitionBatchCache(kv, rollup)
		require.NoError(t, cache.InitAndSyncFromDatabase())
		orphan := transitionCutover + 100
		writeOrphan(t, kv, transitionCutover, orphan)

		batch, err := cache.Get(orphan)
		require.ErrorIs(t, err, ErrKeyNotFound)
		require.Nil(t, batch)
		header, ok := cache.GetSealedBatchHeader(orphan)
		require.False(t, ok)
		require.Nil(t, header)
	})
}

func TestInitAndSyncFromDatabaseRejectsCacheBehindFinalizedIndex(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	rollup.finalized = transitionCutover + 1
	rollup.committed = transitionCutover + 1
	rollup.committedHash[transitionCutover+1] = [32]byte(common.HexToHash("0x1234"))
	rollup.batchBlockByID[transitionCutover+1] = (transitionCutover + 1) * 10

	err := newTransitionBatchCache(kv, rollup).InitAndSyncFromDatabase()
	requireTransitionRecoveryAttempted(t, err)
	requireTransitionStorageCleared(t, kv)
}

func TestInitAndSyncFromDatabaseSnapshotChangeKeepsValidatedDBForRetry(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	rollup.committedSequence = []uint64{transitionCutover, transitionCutover + 1}

	cache := newTransitionBatchCache(kv, rollup)
	err := cache.InitAndSyncFromDatabase()
	require.ErrorContains(t, err, "rollup recovery snapshot changed")
	require.False(t, cache.isInitialized())
	stored, loadErr := NewBatchStorage(kv).LoadSealedBatch(transitionCutover)
	require.NoError(t, loadErr)
	require.NotNil(t, stored, "a concurrent status change is not cache corruption and must not wipe DB")
}

func TestInitAndSyncFromDatabaseRejectsSameIndexCommittedTipChange(t *testing.T) {
	kv := openTestKV(t)
	rollup, _ := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	originalTip := rollup.committedHash[transitionCutover]
	rollup.committedHashSeries = map[uint64][][32]byte{
		transitionCutover: {
			originalTip,
			[32]byte(common.HexToHash("0xdeadbeef")),
		},
	}

	cache := newTransitionBatchCache(kv, rollup)
	err := cache.InitAndSyncFromDatabase()
	require.ErrorContains(t, err, "committed tip 12 changed without an index change")
	require.False(t, cache.isInitialized())
	stored, loadErr := NewBatchStorage(kv).LoadSealedBatch(transitionCutover)
	require.NoError(t, loadErr)
	require.NotNil(t, stored, "a changing canonical view must be retried before deciding whether to rebuild")
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
		requireTransitionRecoveryAttempted(t, err)
		requireTransitionStorageCleared(t, kv)
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
	headerFromCache, ok := cache.GetSealedBatchHeader(transitionCutover)
	require.False(t, ok)
	require.Nil(t, headerFromCache)
}

func TestInitFromRollupLegacyTransitionStartsRecovery(t *testing.T) {
	rollup := &transitionRollupReader{
		committed:      transitionCutover,
		finalized:      transitionFinalized,
		cutover:        transitionCutover,
		committedHash:  make(map[uint64][32]byte),
		batchBlockByID: make(map[uint64]uint64),
	}
	err := newTransitionBatchCache(openTestKV(t), rollup).Init()
	requireTransitionRecoveryAttempted(t, err)
}

func TestReplayBuiltTransitionWindowMayStartAfterFinalizedAnchor(t *testing.T) {
	kv := openTestKV(t)
	rollup, expectedHeaders := storeTransitionWindow(t, kv, transitionCutover, common.Hash{})
	storage := NewBatchStorage(kv)
	require.NoError(t, storage.DeleteSealedBatch(transitionFinalized))

	batches, headers, indices, err := storage.LoadAllSealedBatchesAndHeader()
	require.NoError(t, err)
	anchor := expectedHeaders[transitionFinalized]
	require.NoError(t, validateCachedWindowAgainstFinalizedAnchor(
		batches,
		headers,
		transitionFinalized,
		rollup.committed,
		&anchor,
	))
	require.NoError(t, validateLegacyTransitionWindow(
		batches,
		headers,
		indices,
		transitionFinalized,
		rollup.committed,
		transitionCutover,
	))

	mutated := *batches[transitionFinalized+1]
	mutated.ParentBatchHeader = append(hexutil.Bytes(nil), mutated.ParentBatchHeader...)
	mutated.ParentBatchHeader[len(mutated.ParentBatchHeader)-1] ^= 0x01
	batches[transitionFinalized+1] = &mutated
	err = validateCachedWindowAgainstFinalizedAnchor(
		batches,
		headers,
		transitionFinalized,
		rollup.committed,
		&anchor,
	)
	require.ErrorContains(t, err, "exact finalized L1 header")
}

func TestBatchHeaderSequencerHashIsRestrictedToLegacyReplay(t *testing.T) {
	cache := newTransitionBatchCache(openTestKV(t), &transitionRollupReader{})
	parent := makeTransitionHeader(transitionFinalized, common.Hash{}, common.HexToHash("0x01"), blob.EmptyVersionedHash)
	cache.parentBatchHeader = &parent
	cache.totalL1MessagePopped = transitionFinalized + 101
	legacyHash := common.HexToHash("0x1234")

	legacyHeader := cache.createBatchHeaderWithSequencerSetHash(common.HexToHash("0xaa"), nil, legacyHash, 1)
	gotLegacyHash, err := legacyHeader.SequencerSetVerifyHash()
	require.NoError(t, err)
	require.Equal(t, legacyHash, gotLegacyHash)

	liveHeader := cache.createBatchHeader(common.HexToHash("0xbb"), nil, 1)
	gotLiveHash, err := liveHeader.SequencerSetVerifyHash()
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, gotLiveHash)
}

func TestEmptyTransitionDatabaseRebuildsAndRestartsFromFIPlusOne(t *testing.T) {
	const (
		finalized = uint64(10)
		cutover   = uint64(12)
		committed = uint64(13)
	)

	rollup := &transitionRollupReader{
		committed:      committed,
		finalized:      finalized,
		cutover:        cutover,
		committedHash:  make(map[uint64][32]byte),
		batchBlockByID: make(map[uint64]uint64),
	}
	for index := finalized; index <= committed; index++ {
		rollup.batchBlockByID[index] = index - finalized
	}

	anchor := makeTransitionHeader(
		finalized,
		common.HexToHash("0x1000"),
		common.HexToHash("0x2000"),
		blob.EmptyVersionedHash,
	)
	binary.BigEndian.PutUint64(anchor[len(anchor)-8:], 0)
	anchorHash, err := anchor.Hash()
	require.NoError(t, err)
	rollup.committedHash[finalized] = [32]byte(anchorHash)

	blocks := make(map[uint64]*ethtypes.Block)
	roots := make(map[uint64]common.Hash)
	sequencerSets := make(map[uint64][]byte)
	for index := finalized + 1; index <= committed; index++ {
		blockNumber := index - finalized
		blocks[blockNumber] = ethtypes.NewBlockWithHeader(&ethtypes.Header{
			Number: new(big.Int).SetUint64(blockNumber),
			Time:   1_000 + index,
			Root:   common.BigToHash(new(big.Int).SetUint64(index*100 + 1)),
		})
		roots[blockNumber] = common.BigToHash(new(big.Int).SetUint64(index*100 + 2))
		if index <= cutover {
			sequencerSets[blockNumber] = []byte(fmt.Sprintf("historical-sequencer-set-%d", index))
		}
	}
	l2Client := &transitionReplayL2Client{blocks: blocks}

	// Build the canonical committed hashes with the same deterministic packing
	// primitives. This is independent of the recovery loop and lets the test
	// exercise its pre-persistence committedBatches checks.
	expected := newTransitionBatchCache(openTestKV(t), rollup)
	expected.l2Clients = l2Client
	expected.parentBatchHeader = &anchor
	expected.prevStateRoot, err = anchor.PostStateRoot()
	require.NoError(t, err)
	expected.totalL1MessagePopped, err = anchor.TotalL1MessagePopped()
	require.NoError(t, err)
	expected.lastPackedBlockHeight = 0
	expected.legacyCutoverBatchIndex = cutover
	expectedHeaders := make(map[uint64]BatchHeaderBytes)
	for index := finalized + 1; index <= committed; index++ {
		blockNumber := index - finalized
		expected.batchData = NewBatchData()
		expected.ClearCurrent()
		_, err = expected.CalculateCapWithProposalBlock(blockNumber, roots[blockNumber])
		require.NoError(t, err)
		require.NoError(t, expected.PackCurrentBlock(blockNumber))
		replayIndex := index
		blobCap := expected.sealEffectiveBlobCount(blocks[blockNumber].Time(), &replayIndex)
		payload, dataHash, blobCap, err := expected.handleBatchSealing(blocks[blockNumber].Time(), blobCap, &replayIndex)
		require.NoError(t, err)
		sidecar, err := blob.MakeBlobTxSidecar(payload, blobCap)
		require.NoError(t, err)
		sequencerHash := common.Hash{}
		if index <= cutover {
			sequencerHash = crypto.Keccak256Hash(sequencerSets[blockNumber])
		}
		header := expected.createBatchHeaderWithSequencerSetHash(dataHash, sidecar, sequencerHash, blocks[blockNumber].Time())
		hash, err := header.Hash()
		require.NoError(t, err)
		rollup.committedHash[index] = [32]byte(hash)
		expectedHeaders[index] = header
		headerCopy := append(BatchHeaderBytes(nil), header...)
		expected.parentBatchHeader = &headerCopy
		expected.prevStateRoot = expected.postStateRoot
	}

	badRollup := *rollup
	badRollup.committedHash = make(map[uint64][32]byte, len(rollup.committedHash))
	for index, hash := range rollup.committedHash {
		badRollup.committedHash[index] = hash
	}
	badRollup.committedHash[cutover] = [32]byte(common.HexToHash("0xdeadbeef"))
	badKV := openTestKV(t)
	badCache := newTransitionBatchCache(badKV, &badRollup)
	badCache.l2Clients = l2Client
	badCache.l2Gov = &transitionReplayL2Gov{roots: roots, sequencerSetByBlock: sequencerSets}
	badCache.finalizedBatchHeaderLoader = func(index uint64) (*BatchHeaderBytes, error) {
		copyOfAnchor := append(BatchHeaderBytes(nil), anchor...)
		return &copyOfAnchor, nil
	}
	err = badCache.InitAndSyncFromDatabase()
	require.ErrorContains(t, err, "replay batch hash mismatch")
	require.False(t, badCache.isInitialized())
	_, ok := badCache.GetSealedBatchHeader(finalized + 1)
	require.False(t, ok, "partial replay headers must remain unpublished")
	requireTransitionStorageCleared(t, badKV)

	advancingRollup := *rollup
	advancingRollup.committedSequence = []uint64{cutover, cutover, committed, committed}
	advancingRollup.committedReads = 0
	advancingKV := openTestKV(t)
	advancingGov := &transitionReplayL2Gov{roots: roots, sequencerSetByBlock: sequencerSets}
	advancingCache := newTransitionBatchCache(advancingKV, &advancingRollup)
	advancingCache.l2Clients = l2Client
	advancingCache.l2Gov = advancingGov
	advancingCache.finalizedBatchHeaderLoader = func(index uint64) (*BatchHeaderBytes, error) {
		copyOfAnchor := append(BatchHeaderBytes(nil), anchor...)
		return &copyOfAnchor, nil
	}
	require.NoError(t, advancingCache.InitAndSyncFromDatabase())
	latestAdvancedIndex, err := advancingCache.LatestBatchIndex()
	require.NoError(t, err)
	require.Equal(t, committed, latestAdvancedIndex, "CI growth during replay must be caught up incrementally")
	require.Equal(t, []uint64{1, 2}, advancingGov.sequencerCalls)

	kv := openTestKV(t)
	l2Gov := &transitionReplayL2Gov{roots: roots, sequencerSetByBlock: sequencerSets}
	cache := newTransitionBatchCache(kv, rollup)
	cache.l2Clients = l2Client
	cache.l2Gov = l2Gov
	cache.finalizedBatchHeaderLoader = func(index uint64) (*BatchHeaderBytes, error) {
		require.Equal(t, finalized, index)
		copyOfAnchor := append(BatchHeaderBytes(nil), anchor...)
		return &copyOfAnchor, nil
	}

	require.NoError(t, cache.InitAndSyncFromDatabase())
	require.True(t, cache.isInitialized())
	require.Equal(t, []uint64{1, 2}, l2Gov.sequencerCalls, "only batches at or before K read historical sequencer state")
	for index := finalized + 1; index <= committed; index++ {
		header, ok := cache.GetSealedBatchHeader(index)
		require.True(t, ok)
		require.Equal(t, expectedHeaders[index].Bytes(), header.Bytes())
		sequencerHash, err := header.SequencerSetVerifyHash()
		require.NoError(t, err)
		batch, ok := cache.GetSealedBatch(index)
		require.True(t, ok)
		if index <= cutover {
			require.Equal(t, crypto.Keccak256Hash(sequencerSets[index-finalized]), sequencerHash)
			require.Equal(t, sequencerSets[index-finalized], []byte(batch.CurrentSequencerSetBytes))
		} else {
			require.Equal(t, common.Hash{}, sequencerHash)
			require.Empty(t, batch.CurrentSequencerSetBytes)
		}
	}

	// FI is intentionally absent on disk. A second process must accept the
	// canonical [FI+1, CI] window and take the DB fast path without touching L2.
	restartGov := &transitionReplayL2Gov{roots: roots, sequencerSetByBlock: sequencerSets}
	restarted := newTransitionBatchCache(kv, rollup)
	restarted.l2Clients = l2Client
	restarted.l2Gov = restartGov
	restarted.finalizedBatchHeaderLoader = cache.finalizedBatchHeaderLoader
	require.NoError(t, restarted.InitAndSyncFromDatabase())
	require.True(t, restarted.isInitialized())
	require.Empty(t, restartGov.sequencerCalls)
	latestIndex, err := restarted.LatestBatchIndex()
	require.NoError(t, err)
	require.Equal(t, committed, latestIndex)
	_, err = NewBatchStorage(kv).LoadSealedBatch(finalized)
	require.ErrorIs(t, err, ErrKeyNotFound)
}
