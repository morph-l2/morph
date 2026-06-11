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
	require.NoError(t, s.DeleteSealedBatch(102))

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
		require.Error(t, err, "batch %d should be deleted", idx)
		_, err = s.LoadSealedBatchHeader(idx)
		require.Error(t, err, "header %d should be deleted", idx)
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
	require.Error(t, err)
	_, err = s.LoadSealedBatchHeader(100)
	require.Error(t, err)

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
}
