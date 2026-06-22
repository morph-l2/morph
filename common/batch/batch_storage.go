package batch

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"
	ldberrors "github.com/syndtr/goleveldb/leveldb/errors"
)

const (
	// Key prefixes for LevelDB storage
	SealedBatchKeyPrefix       = "sealed_batch_"
	SealedBatchHeaderKeyPrefix = "sealed_batch_header_"
	SealedBatchIndicesKey      = "sealed_batch_indices"
)

// BatchStorage handles persistence of sealed batches using JSON encoding
type BatchStorage struct {
	db SealedBatchKV
	mu sync.RWMutex
}

// NewBatchStorage creates a new BatchStorage instance
func NewBatchStorage(db SealedBatchKV) *BatchStorage {
	return &BatchStorage{
		db: db,
	}
}

// StoreSealedBatch stores a single sealed batch to LevelDB
// Uses JSON encoding for serialization.
// Batch data and the indices snapshot are written in one atomic batch so they
// can never get out of sync; indices update failures are no longer swallowed.
func (s *BatchStorage) StoreSealedBatch(batchIndex uint64, batch *eth.RPCRollupBatch) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Serialize batch to JSON
	encoded, err := json.Marshal(batch)
	if err != nil {
		return fmt.Errorf("failed to marshal sealed batch %d: %w", batchIndex, err)
	}

	encodedIndices, err := s.indicesSnapshotWith(batchIndex)
	if err != nil {
		return fmt.Errorf("failed to update batch indices for batch %d: %w", batchIndex, err)
	}

	puts := []KVPair{
		{Key: encodeBatchKey(batchIndex), Value: encoded},
		{Key: []byte(SealedBatchIndicesKey), Value: encodedIndices},
	}
	if err := s.db.WriteBatch(puts, nil); err != nil {
		return fmt.Errorf("failed to store sealed batch %d: %w", batchIndex, err)
	}
	return nil
}

// StoreSealedBatchAndHeader stores the sealed batch, its header and the updated
// indices snapshot in a single atomic write.
func (s *BatchStorage) StoreSealedBatchAndHeader(batchIndex uint64, batch *eth.RPCRollupBatch, header *BatchHeaderBytes) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	encoded, err := json.Marshal(batch)
	if err != nil {
		return fmt.Errorf("failed to marshal sealed batch %d: %w", batchIndex, err)
	}

	encodedIndices, err := s.indicesSnapshotWith(batchIndex)
	if err != nil {
		return fmt.Errorf("failed to update batch indices for batch %d: %w", batchIndex, err)
	}

	puts := []KVPair{
		{Key: encodeBatchKey(batchIndex), Value: encoded},
		{Key: encodeBatchHeaderKey(batchIndex), Value: header.Bytes()},
		{Key: []byte(SealedBatchIndicesKey), Value: encodedIndices},
	}
	if err := s.db.WriteBatch(puts, nil); err != nil {
		return fmt.Errorf("failed to store sealed batch and header %d: %w", batchIndex, err)
	}
	return nil
}

// LoadSealedBatch loads a single sealed batch from LevelDB
func (s *BatchStorage) LoadSealedBatch(batchIndex uint64) (*eth.RPCRollupBatch, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	key := encodeBatchKey(batchIndex)
	encoded, err := s.db.GetBytes(key)
	if err != nil {
		if isKVNotFound(err) {
			return nil, fmt.Errorf("sealed batch %d not found: %w", batchIndex, ErrKeyNotFound)
		}
		return nil, fmt.Errorf("failed to get sealed batch %d: %w", batchIndex, err)
	}

	// Deserialize from JSON
	var batch eth.RPCRollupBatch
	if err := json.Unmarshal(encoded, &batch); err != nil {
		return nil, fmt.Errorf("failed to unmarshal sealed batch %d: %w", batchIndex, err)
	}

	return &batch, nil
}

// LoadAllSealedBatches loads all sealed batches from LevelDB
// Returns a map of batchIndex -> RPCRollupBatch
func (s *BatchStorage) LoadAllSealedBatches() (map[uint64]*eth.RPCRollupBatch, []uint64, error) {
	s.mu.RLock()
	// Load batch indices
	indices, err := s.loadBatchIndices()
	s.mu.RUnlock()
	if err != nil {
		if isKVNotFound(err) {
			// No batches stored yet
			return make(map[uint64]*eth.RPCRollupBatch), nil, nil
		}
		return nil, nil, fmt.Errorf("failed to load batch indices: %w", err)
	}

	// Load each batch (without holding the lock to avoid deadlock)
	batches := make(map[uint64]*eth.RPCRollupBatch, len(indices))
	for _, idx := range indices {
		batch, err := s.LoadSealedBatch(idx)
		if err != nil {
			log.Warn("Failed to load sealed batch, skipping",
				"batch_index", idx, "error", err)
			continue
		}
		batches[idx] = batch
	}

	return batches, indices, nil
}

// LoadAllSealedBatchesAndHeader loads all sealed batches and batch header from LevelDB
func (s *BatchStorage) LoadAllSealedBatchesAndHeader() (map[uint64]*eth.RPCRollupBatch, map[uint64]*BatchHeaderBytes, []uint64, error) {
	s.mu.RLock()
	// Load batch indices
	indices, err := s.loadBatchIndices()
	s.mu.RUnlock()
	if err != nil {
		if isKVNotFound(err) {
			// No batches stored yet
			return make(map[uint64]*eth.RPCRollupBatch), make(map[uint64]*BatchHeaderBytes), nil, nil
		}
		return nil, nil, nil, fmt.Errorf("failed to load batch indices: %w", err)
	}

	// Load each batch (without holding the lock to avoid deadlock)
	batches := make(map[uint64]*eth.RPCRollupBatch, len(indices))
	for i, idx := range indices {
		batch, err := s.LoadSealedBatch(idx)
		if err != nil {
			log.Warn("Failed to load sealed batch, aborting",
				"batch_index", idx, "error", err)
			return nil, nil, nil, fmt.Errorf("failed to load batch: %w", err)
		}
		if i > 0 {
			// indices is sorted ascending; a hole means some middle batch was
			// deleted (e.g. by a finalize confirmed while other submitters
			// advanced the finalize index). Return an error so the caller can
			// self-heal instead of dereferencing a missing parent batch.
			prevIdx := indices[i-1]
			if idx != prevIdx+1 {
				log.Error("Sealed batch indices are not contiguous",
					"prev_index", prevIdx, "batch_index", idx)
				return nil, nil, nil, fmt.Errorf("sealed batch indices not contiguous: %d -> %d", prevIdx, idx)
			}
			parentBatch := batches[prevIdx]
			if parentBatch == nil {
				log.Error("Parent batch missing", "parent_index", prevIdx, "batch_index", idx)
				return nil, nil, nil, fmt.Errorf("parent batch %d missing for batch %d", prevIdx, idx)
			}
			parentBatchHash, err := BatchHeaderBytes(batch.ParentBatchHeader).Hash()
			if err != nil {
				log.Error("Failed to load parent batch header", "batch_index", idx, "error", err)
				return nil, nil, nil, fmt.Errorf("failed to load batch header: %w", err)
			}
			if !bytes.Equal(parentBatch.Hash.Bytes(), parentBatchHash.Bytes()) {
				log.Error("parent batch hash check failed",
					"batch_index", idx,
					"parent_batch_hash", parentBatch.Hash.String(),
					"pre_batch_hash", parentBatchHash.String())
				return nil, nil, nil, fmt.Errorf("parent batch hash check failed")
			}
		}
		batches[idx] = batch
	}
	// Load each batch header (without holding the lock to avoid deadlock)
	headers := make(map[uint64]*BatchHeaderBytes, len(indices))
	for _, idx := range indices {
		header, err := s.LoadSealedBatchHeader(idx)
		if err != nil {
			log.Warn("Failed to load sealed batch header, skipping",
				"batch_index", idx, "error", err)
			return nil, nil, nil, fmt.Errorf("failed to load batch header bytes: %w", err)
		}
		headers[idx] = header
		headerHash, err := header.Hash()
		if err != nil {
			log.Warn("Failed to hash sealed batch header, skipping",
				"batch_index", idx, "error", err)
			return nil, nil, nil, fmt.Errorf("failed to load batch header bytes: %w", err)
		}
		// check header and batch hash equal
		if !bytes.Equal(headerHash.Bytes(), batches[idx].Hash.Bytes()) {
			log.Error("Sealed batch header bytes do not match",
				"batch_index", idx, "expected", batches[idx].Hash, "actual", headerHash.Bytes())
			return nil, nil, nil, fmt.Errorf("sealed batch header bytes do not match")
		}
	}
	return batches, headers, indices, nil
}

// DeleteSealedBatch removes a sealed batch (data + header) from LevelDB.
// Data, header and the indices snapshot are removed in one atomic write;
// indices update failures are no longer swallowed.
func (s *BatchStorage) DeleteSealedBatch(batchIndex uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	encodedIndices, err := s.indicesSnapshotWithout(batchIndex)
	if err != nil {
		return fmt.Errorf("failed to update batch indices for batch %d: %w", batchIndex, err)
	}

	puts := []KVPair{{Key: []byte(SealedBatchIndicesKey), Value: encodedIndices}}
	deletes := [][]byte{encodeBatchKey(batchIndex), encodeBatchHeaderKey(batchIndex)}
	if err := s.db.WriteBatch(puts, deletes); err != nil {
		return fmt.Errorf("failed to delete sealed batch %d: %w", batchIndex, err)
	}
	return nil
}

// DeleteSealedBatchesUpTo removes every sealed batch (data + header) with
// index <= maxIndex in a single atomic write. Range-based cleanup keeps the
// surviving indices a contiguous window: single-index deletes punch holes when
// the finalize target jumps (multiple submitters, finalizing several batches at
// once), and such holes crash the startup load path.
func (s *BatchStorage) DeleteSealedBatchesUpTo(maxIndex uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	indices, err := s.loadBatchIndices()
	if err != nil {
		if isKVNotFound(err) {
			// No batches stored yet
			return nil
		}
		return fmt.Errorf("failed to load batch indices: %w", err)
	}

	kept := make([]uint64, 0, len(indices))
	var deletes [][]byte
	for _, idx := range indices {
		if idx <= maxIndex {
			deletes = append(deletes, encodeBatchKey(idx), encodeBatchHeaderKey(idx))
		} else {
			kept = append(kept, idx)
		}
	}
	if len(deletes) == 0 {
		return nil
	}

	encodedIndices, err := json.Marshal(kept)
	if err != nil {
		return fmt.Errorf("failed to marshal batch indices: %w", err)
	}
	var puts []KVPair
	if len(kept) == 0 {
		deletes = append(deletes, []byte(SealedBatchIndicesKey))
	} else {
		puts = []KVPair{{Key: []byte(SealedBatchIndicesKey), Value: encodedIndices}}
	}
	if err := s.db.WriteBatch(puts, deletes); err != nil {
		return fmt.Errorf("failed to delete sealed batches up to %d: %w", maxIndex, err)
	}
	return nil
}

func (s *BatchStorage) DeleteAllSealedBatches() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	indices, err := s.loadBatchIndices()
	if err != nil {
		if isKVNotFound(err) {
			// No batches stored yet
			return nil
		}
		return fmt.Errorf("failed to load batch indices: %w", err)
	}

	deletes := make([][]byte, 0, len(indices)*2+1)
	for _, idx := range indices {
		deletes = append(deletes, encodeBatchKey(idx), encodeBatchHeaderKey(idx))
	}
	deletes = append(deletes, []byte(SealedBatchIndicesKey))
	if err := s.db.WriteBatch(nil, deletes); err != nil {
		return fmt.Errorf("failed to delete all sealed batches: %w", err)
	}
	return nil
}

// encodeBatchKey encodes batch index to a byte key
func encodeBatchKey(batchIndex uint64) []byte {
	key := make([]byte, len(SealedBatchKeyPrefix)+8)
	copy(key, SealedBatchKeyPrefix)
	binary.BigEndian.PutUint64(key[len(SealedBatchKeyPrefix):], batchIndex)
	return key
}

// indicesSnapshotWith returns the marshaled indices snapshot with batchIndex added.
func (s *BatchStorage) indicesSnapshotWith(batchIndex uint64) ([]byte, error) {
	indices, err := s.loadBatchIndices()
	if err != nil {
		if isKVNotFound(err) {
			indices = []uint64{}
		} else {
			return nil, err
		}
	}

	found := false
	for _, idx := range indices {
		if idx == batchIndex {
			found = true
			break
		}
	}
	if !found {
		indices = append(indices, batchIndex)
		sort.Slice(indices, func(i, j int) bool { return indices[i] < indices[j] })
	}
	return json.Marshal(indices)
}

// indicesSnapshotWithout returns the marshaled indices snapshot with batchIndex removed.
func (s *BatchStorage) indicesSnapshotWithout(batchIndex uint64) ([]byte, error) {
	indices, err := s.loadBatchIndices()
	if err != nil {
		if isKVNotFound(err) {
			indices = []uint64{}
		} else {
			return nil, err
		}
	}

	newIndices := make([]uint64, 0, len(indices))
	for _, idx := range indices {
		if idx != batchIndex {
			newIndices = append(newIndices, idx)
		}
	}
	return json.Marshal(newIndices)
}

// loadBatchIndices loads the list of stored batch indices, sorted ascending.
func (s *BatchStorage) loadBatchIndices() ([]uint64, error) {
	encoded, err := s.db.GetBytes([]byte(SealedBatchIndicesKey))
	if err != nil {
		return nil, err
	}

	var indices []uint64
	if err := json.Unmarshal(encoded, &indices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal batch indices: %w", err)
	}

	// Keep ordering deterministic regardless of how the snapshot was produced.
	sort.Slice(indices, func(i, j int) bool { return indices[i] < indices[j] })
	return indices, nil
}

// StoreSealedBatchHeader stores a single sealed batch header to LevelDB
func (s *BatchStorage) StoreSealedBatchHeader(batchIndex uint64, header *BatchHeaderBytes) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Store batch header data
	key := encodeBatchHeaderKey(batchIndex)
	if err := s.db.PutBytes(key, header.Bytes()); err != nil {
		return fmt.Errorf("failed to store sealed batch header %d: %w", batchIndex, err)
	}

	return nil
}

// LoadSealedBatchHeader loads a single sealed batch header from LevelDB
func (s *BatchStorage) LoadSealedBatchHeader(batchIndex uint64) (*BatchHeaderBytes, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	key := encodeBatchHeaderKey(batchIndex)
	headerBytes, err := s.db.GetBytes(key)
	if err != nil {
		if isKVNotFound(err) {
			return nil, fmt.Errorf("sealed batch header %d not found: %w", batchIndex, ErrKeyNotFound)
		}
		return nil, fmt.Errorf("failed to get sealed batch header %d: %w", batchIndex, err)
	}

	header := BatchHeaderBytes(headerBytes)
	return &header, nil
}

// LoadAllSealedBatchHeaders loads all sealed batch headers from LevelDB
// Returns a map of batchIndex -> BatchHeaderBytes
func (s *BatchStorage) LoadAllSealedBatchHeaders() (map[uint64]*BatchHeaderBytes, error) {
	s.mu.RLock()
	// Load batch indices
	indices, err := s.loadBatchIndices()
	s.mu.RUnlock()
	if err != nil {
		if isKVNotFound(err) {
			// No batches stored yet
			return make(map[uint64]*BatchHeaderBytes), nil
		}
		return nil, fmt.Errorf("failed to load batch indices: %w", err)
	}

	// Load each batch header (without holding the lock to avoid deadlock)
	headers := make(map[uint64]*BatchHeaderBytes, len(indices))
	for _, idx := range indices {
		header, err := s.LoadSealedBatchHeader(idx)
		if err != nil {
			log.Warn("Failed to load sealed batch header, skipping",
				"batch_index", idx, "error", err)
			continue
		}
		headers[idx] = header
	}

	return headers, nil
}

// DeleteSealedBatchHeader removes a sealed batch header from LevelDB
func (s *BatchStorage) DeleteSealedBatchHeader(batchIndex uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := encodeBatchHeaderKey(batchIndex)
	if err := s.db.Delete(key); err != nil {
		return fmt.Errorf("failed to delete sealed batch header %d: %w", batchIndex, err)
	}

	return nil
}

// encodeBatchHeaderKey encodes batch index to a byte key for batch header
func encodeBatchHeaderKey(batchIndex uint64) []byte {
	key := make([]byte, len(SealedBatchHeaderKeyPrefix)+8)
	copy(key, SealedBatchHeaderKeyPrefix)
	binary.BigEndian.PutUint64(key[len(SealedBatchHeaderKeyPrefix):], batchIndex)
	return key
}

func isKVNotFound(err error) bool {
	return errors.Is(err, ErrKeyNotFound) || errors.Is(err, ldberrors.ErrNotFound)
}
