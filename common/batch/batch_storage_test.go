package batch

// Regression tests for the restart panic caused by holes in the persisted
// sealed batch indices (nil pointer dereference at the parent batch hash check
// in LoadAllSealedBatchesAndHeader). See .vscode/doing/submitter-issue.md.

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"
)

// makeTestHeader builds a unique, valid V1 batch header (257 bytes, version byte = 1).
func makeTestHeader(idx uint64) BatchHeaderBytes {
	h := make(BatchHeaderBytes, expectedLengthV1)
	h[0] = BatchHeaderVersion1
	h[1] = byte(idx)
	h[2] = byte(idx >> 8)
	return h
}

// storeTestChain stores a chain of sealed batches whose parent headers link
// consecutively, mirroring what SealBatch persists.
func storeTestChain(t *testing.T, s *BatchStorage, indices []uint64) {
	t.Helper()
	for _, idx := range indices {
		header := makeTestHeader(idx)
		hash, err := header.Hash()
		require.NoError(t, err)

		parentHeader := makeTestHeader(idx - 1)
		b := &eth.RPCRollupBatch{
			Hash:              hash,
			ParentBatchHeader: hexutil.Bytes(parentHeader),
		}
		require.NoError(t, s.StoreSealedBatchAndHeader(idx, b, &header))
	}
}

func TestLoadAllSealedBatchesAndHeaderContiguous(t *testing.T) {
	s := NewBatchStorage(openTestKV(t))
	storeTestChain(t, s, []uint64{100, 101, 102, 103})

	batches, headers, indices, err := s.LoadAllSealedBatchesAndHeader()
	require.NoError(t, err)
	require.Len(t, batches, 4)
	require.Len(t, headers, 4)
	require.Equal(t, []uint64{100, 101, 102, 103}, indices)
}

// A hole in the indices (middle batch deleted, as finalize cleanup used to do
// with single-index deletes) must surface as an error so the caller can
// self-heal, instead of panicking on a nil parent batch.
func TestLoadAllSealedBatchesAndHeaderHoleReturnsError(t *testing.T) {
	s := NewBatchStorage(openTestKV(t))
	storeTestChain(t, s, []uint64{100, 101, 102, 103})

	// Inject a hole directly at the KV layer: production DeleteSealedBatch now
	// refuses interior deletes, but a legacy binary or a crash could still leave
	// a gap, so the load path must reject it rather than panic on a nil parent.
	holed, err := json.Marshal([]uint64{100, 101, 103})
	require.NoError(t, err)
	require.NoError(t, s.db.WriteBatch(
		[]KVPair{{Key: []byte(SealedBatchIndicesKey), Value: holed}},
		[][]byte{encodeBatchKey(102), encodeBatchHeaderKey(102)},
	))

	require.NotPanics(t, func() {
		_, _, _, err := s.LoadAllSealedBatchesAndHeader()
		require.Error(t, err)
		require.Contains(t, err.Error(), "not contiguous")
	})
}

func TestDeleteSealedBatchesUpTo(t *testing.T) {
	s := NewBatchStorage(openTestKV(t))
	storeTestChain(t, s, []uint64{100, 101, 102, 103, 104, 105})

	require.NoError(t, s.DeleteSealedBatchesUpTo(103))

	indices, err := s.loadBatchIndices()
	require.NoError(t, err)
	require.Equal(t, []uint64{104, 105}, indices)

	for idx := uint64(100); idx <= 103; idx++ {
		_, err := s.LoadSealedBatch(idx)
		require.True(t, isKVNotFound(err), "batch %d should be deleted, got: %v", idx, err)
		_, err = s.LoadSealedBatchHeader(idx)
		require.True(t, isKVNotFound(err), "header %d should be deleted, got: %v", idx, err)
	}
	for idx := uint64(104); idx <= 105; idx++ {
		_, err := s.LoadSealedBatch(idx)
		require.NoError(t, err, "batch %d should survive", idx)
		_, err = s.LoadSealedBatchHeader(idx)
		require.NoError(t, err, "header %d should survive", idx)
	}

	// The surviving window stays loadable.
	batches, headers, _, err := s.LoadAllSealedBatchesAndHeader()
	require.NoError(t, err)
	require.Len(t, batches, 2)
	require.Len(t, headers, 2)

	// Deleting below the window is a no-op.
	require.NoError(t, s.DeleteSealedBatchesUpTo(50))
	indices, err = s.loadBatchIndices()
	require.NoError(t, err)
	require.Equal(t, []uint64{104, 105}, indices)

	// Delete exactly at the boundary.
	s2 := NewBatchStorage(openTestKV(t))
	storeTestChain(t, s2, []uint64{100, 101})
	require.NoError(t, s2.DeleteSealedBatchesUpTo(101))
	_, err = s2.loadBatchIndices()
	require.True(t, isKVNotFound(err), "all indices should be deleted, got: %v", err)

	// Delete beyond the window.
	s3 := NewBatchStorage(openTestKV(t))
	storeTestChain(t, s3, []uint64{100, 101})
	require.NoError(t, s3.DeleteSealedBatchesUpTo(200))
	_, err = s3.loadBatchIndices()
	require.True(t, isKVNotFound(err), "all indices should be deleted, got: %v", err)
}

// Legacy snapshots may have been persisted unsorted; loading must normalize.
func TestLoadBatchIndicesSortsLegacySnapshot(t *testing.T) {
	kv := openTestKV(t)
	s := NewBatchStorage(kv)

	encoded, err := json.Marshal([]uint64{103, 100, 102, 101})
	require.NoError(t, err)
	require.NoError(t, kv.PutBytes([]byte(SealedBatchIndicesKey), encoded))

	indices, err := s.loadBatchIndices()
	require.NoError(t, err)
	require.Equal(t, []uint64{100, 101, 102, 103}, indices)
}

func TestDeleteAllSealedBatches(t *testing.T) {
	s := NewBatchStorage(openTestKV(t))
	storeTestChain(t, s, []uint64{100, 101, 102})

	require.NoError(t, s.DeleteAllSealedBatches())

	_, err := s.loadBatchIndices()
	require.True(t, isKVNotFound(err))
	_, err = s.LoadSealedBatch(100)
	require.True(t, isKVNotFound(err), "batch should be deleted, got: %v", err)
	_, err = s.LoadSealedBatchHeader(100)
	require.True(t, isKVNotFound(err), "header should be deleted, got: %v", err)

	// Idempotent on empty storage.
	require.NoError(t, s.DeleteAllSealedBatches())
}

// failingIndicesKV simulates a KV whose indices key is unreadable; store/delete
// must propagate the error instead of swallowing it.
type failingIndicesKV struct {
	SealedBatchKV
}

var errIndicesUnavailable = errors.New("indices unavailable")

func (f *failingIndicesKV) GetBytes(key []byte) ([]byte, error) {
	if string(key) == SealedBatchIndicesKey {
		return nil, errIndicesUnavailable
	}
	return f.SealedBatchKV.GetBytes(key)
}

func TestStoreSealedBatchPropagatesIndicesError(t *testing.T) {
	s := NewBatchStorage(&failingIndicesKV{SealedBatchKV: openTestKV(t)})

	header := makeTestHeader(100)
	hash, err := header.Hash()
	require.NoError(t, err)
	b := &eth.RPCRollupBatch{Hash: hash}

	err = s.StoreSealedBatch(100, b)
	require.ErrorIs(t, err, errIndicesUnavailable)

	err = s.StoreSealedBatchAndHeader(100, b, &header)
	require.ErrorIs(t, err, errIndicesUnavailable)

	err = s.DeleteSealedBatch(100)
	require.ErrorIs(t, err, errIndicesUnavailable)

	err = s.DeleteSealedBatchesUpTo(100)
	require.ErrorIs(t, err, errIndicesUnavailable)

	err = s.DeleteAllSealedBatches()
	require.ErrorIs(t, err, errIndicesUnavailable)
}

type failingWriteBatchKV struct {
	SealedBatchKV
	fail bool
	err  error
}

var errWriteBatchUnavailable = errors.New("write batch unavailable")

func (f *failingWriteBatchKV) WriteBatch(puts []KVPair, deletes [][]byte) error {
	if f.fail {
		return f.err
	}
	return f.SealedBatchKV.WriteBatch(puts, deletes)
}

func TestSealBatchPropagatesStoreErrorAndKeepsState(t *testing.T) {
	kv := &failingWriteBatchKV{
		SealedBatchKV: openTestKV(t),
		fail:          true,
		err:           errWriteBatchUnavailable,
	}
	parentHeader := makeTestHeader(1)
	bc := &BatchCache{
		batchStorage:         NewBatchStorage(kv),
		sealedBatches:        make(map[uint64]*eth.RPCRollupBatch),
		sealedBatchHeaders:   make(map[uint64]*BatchHeaderBytes),
		parentBatchHeader:    &parentHeader,
		batchData:            NewBatchData(),
		maxBlobCount:         1,
		isBatchUpgraded:      func(uint64) bool { return false },
		isBatchV2Upgraded:    func(uint64) bool { return false },
		totalL1MessagePopped: 0,
	}
	bc.batchData.Append(make([]byte, 60), []byte{1, 2, 3}, nil)
	originalBatchData := bc.batchData
	originalParent := bc.parentBatchHeader

	_, _, _, err := bc.SealBatch([]byte{1}, 1, nil)
	require.ErrorIs(t, err, errWriteBatchUnavailable)
	require.Empty(t, bc.sealedBatches)
	require.Empty(t, bc.sealedBatchHeaders)
	require.Same(t, originalBatchData, bc.batchData)
	require.False(t, bc.batchData.IsEmpty())
	require.Same(t, originalParent, bc.parentBatchHeader)
}

func TestDeleteUntilKeepsMemoryWhenStorageDeleteFails(t *testing.T) {
	kv := &failingWriteBatchKV{
		SealedBatchKV: openTestKV(t),
		err:           errWriteBatchUnavailable,
	}
	s := NewBatchStorage(kv)
	storeTestChain(t, s, []uint64{100, 101, 102})

	sealedBatches := make(map[uint64]*eth.RPCRollupBatch)
	sealedBatchHeaders := make(map[uint64]*BatchHeaderBytes)
	for _, idx := range []uint64{100, 101, 102} {
		header := makeTestHeader(idx)
		hash, err := header.Hash()
		require.NoError(t, err)
		sealedBatches[idx] = &eth.RPCRollupBatch{Hash: hash}
		sealedBatchHeaders[idx] = &header
	}
	bc := &BatchCache{
		batchStorage:       s,
		sealedBatches:      sealedBatches,
		sealedBatchHeaders: sealedBatchHeaders,
	}

	kv.fail = true
	err := bc.DeleteUntil(101)
	require.ErrorIs(t, err, errWriteBatchUnavailable)
	require.Contains(t, bc.sealedBatches, uint64(100))
	require.Contains(t, bc.sealedBatches, uint64(101))
	require.Contains(t, bc.sealedBatchHeaders, uint64(100))
	require.Contains(t, bc.sealedBatchHeaders, uint64(101))

	kv.fail = false
	indices, err := s.loadBatchIndices()
	require.NoError(t, err)
	require.Equal(t, []uint64{100, 101, 102}, indices)
}

// Single-index deletes may only remove a window boundary; an interior delete
// would punch a hole into sealed_batch_indices and crash the next restart load,
// so it must be refused. Finalize cleanup uses the range-based DeleteSealedBatchesUpTo.
func TestDeleteSealedBatchRejectsInteriorIndex(t *testing.T) {
	s := NewBatchStorage(openTestKV(t))
	storeTestChain(t, s, []uint64{100, 101, 102})

	// Interior index is refused and storage is left unchanged.
	err := s.DeleteSealedBatch(101)
	require.Error(t, err)
	require.Contains(t, err.Error(), "interior")
	indices, err := s.loadBatchIndices()
	require.NoError(t, err)
	require.Equal(t, []uint64{100, 101, 102}, indices)

	// Boundaries (lowest, then highest) are allowed and keep the window contiguous.
	require.NoError(t, s.DeleteSealedBatch(100))
	require.NoError(t, s.DeleteSealedBatch(102))
	indices, err = s.loadBatchIndices()
	require.NoError(t, err)
	require.Equal(t, []uint64{101}, indices)

	// The final remaining index can be removed (single-delete leaves an empty snapshot).
	require.NoError(t, s.DeleteSealedBatch(101))
	indices, err = s.loadBatchIndices()
	require.NoError(t, err)
	require.Empty(t, indices)
}

// When sealed_batch_indices is corrupt the normal wipe cannot enumerate keys, so
// self-heal would stall forever. The prefix-scan force wipe must still clear data,
// header and the indices key so the node can rebuild from rollup.
func TestForceDeleteAllSealedBatchesWipesCorruptIndices(t *testing.T) {
	kv := openTestKV(t)
	s := NewBatchStorage(kv)
	storeTestChain(t, s, []uint64{100, 101, 102})

	// Corrupt the indices snapshot.
	require.NoError(t, kv.PutBytes([]byte(SealedBatchIndicesKey), []byte("{not json")))

	// Normal wipe can no longer enumerate which keys to delete.
	require.Error(t, s.DeleteAllSealedBatches())

	// Force wipe scans the shared prefix and removes everything.
	require.NoError(t, s.ForceDeleteAllSealedBatches())

	keys, err := kv.IteratePrefixKeys([]byte(SealedBatchKeyPrefix))
	require.NoError(t, err)
	require.Empty(t, keys)

	for _, idx := range []uint64{100, 101, 102} {
		_, err := s.LoadSealedBatch(idx)
		require.True(t, isKVNotFound(err), "batch %d should be wiped, got: %v", idx, err)
	}

	// Idempotent on empty storage.
	require.NoError(t, s.ForceDeleteAllSealedBatches())
}
