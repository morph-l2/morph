package batch

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"morph-l2/tx-submitter/db"

	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"
)

const (
	// Key prefixes for LevelDB storage
	SealedBatchKeyPrefix       = "sealed_batch_"
	SealedBatchHeaderKeyPrefix = "sealed_batch_header_"
	SealedBatchIndicesKey      = "sealed_batch_indices"
)

// BatchStorage handles persistence of sealed batches using JSON encoding
type BatchStorage struct {
	db db.Database
	mu sync.RWMutex
}

// NewBatchStorage creates a new BatchStorage instance
func NewBatchStorage(db db.Database) *BatchStorage {
	return &BatchStorage{
		db: db,
	}
}

// StoreSealedBatch stores a single sealed batch to LevelDB
// Uses JSON encoding for serialization
func (s *BatchStorage) StoreSealedBatch(batchIndex uint64, batch *eth.RPCRollupBatch) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Serialize batch to JSON
	encoded, err := json.Marshal(batch)
	if err != nil {
		return fmt.Errorf("failed to marshal sealed batch %d: %w", batchIndex, err)
	}

	// Store batch data
	key := encodeBatchKey(batchIndex)
	if err := s.db.PutBytes(key, encoded); err != nil {
		return fmt.Errorf("failed to store sealed batch %d: %w", batchIndex, err)
	}

	// Update indices list
	if err := s.updateBatchIndices(batchIndex, true); err != nil {
		log.Warn("Failed to update batch indices", "batch_index", batchIndex, "error", err)
		// Don't fail the operation if indices update fails
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
		if errors.Is(err, db.ErrKeyNotFound) {
			return nil, fmt.Errorf("sealed batch %d not found", batchIndex)
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
		if errors.Is(err, db.ErrKeyNotFound) {
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
		if errors.Is(err, db.ErrKeyNotFound) {
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
			log.Warn("Failed to load sealed batch, skipping",
				"batch_index", idx, "error", err)
			return nil, nil, nil, fmt.Errorf("failed to load batch: %w", err)
		}
		if i > 0 {
			parentBatch := batches[idx-1]
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

// DeleteSealedBatch removes a sealed batch from LevelDB
func (s *BatchStorage) DeleteSealedBatch(batchIndex uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := encodeBatchKey(batchIndex)
	if err := s.db.Delete(key); err != nil {
		return fmt.Errorf("failed to delete sealed batch %d: %w", batchIndex, err)
	}

	// Update indices list
	if err := s.updateBatchIndices(batchIndex, false); err != nil {
		log.Warn("Failed to update batch indices after deletion",
			"batch_index", batchIndex, "error", err)
		// Don't fail the operation if indices update fails
	}

	return nil
}

func (s *BatchStorage) DeleteAllSealedBatches() error {
	s.mu.RLock()
	// Load batch indices
	indices, err := s.loadBatchIndices()
	s.mu.RUnlock()
	if err != nil {
		if errors.Is(err, db.ErrKeyNotFound) {
			// No batches stored yet
			return nil
		}
		return fmt.Errorf("failed to load batch indices: %w", err)
	}

	for _, idx := range indices {
		err = s.DeleteSealedBatch(idx)
		if err != nil {
			log.Error("Failed to delete sealed batch",
				"batch_index", idx, "error", err)
			return err
		}
		err = s.DeleteSealedBatchHeader(idx)
		if err != nil {
			log.Error("Failed to delete sealed batch header",
				"batch_index", idx, "error", err)
			return err
		}
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

// updateBatchIndices updates the list of stored batch indices
// add: true to add index, false to remove
func (s *BatchStorage) updateBatchIndices(batchIndex uint64, add bool) error {
	indices, err := s.loadBatchIndices()
	if err != nil {
		if errors.Is(err, db.ErrKeyNotFound) {
			indices = []uint64{}
		} else {
			return err
		}
	}

	if add {
		// Add index if not already present
		found := false
		for _, idx := range indices {
			if idx == batchIndex {
				found = true
				break
			}
		}
		if !found {
			indices = append(indices, batchIndex)
		}
	} else {
		// Remove index
		newIndices := make([]uint64, 0, len(indices))
		for _, idx := range indices {
			if idx != batchIndex {
				newIndices = append(newIndices, idx)
			}
		}
		indices = newIndices
	}

	return s.saveBatchIndices(indices)
}

// loadBatchIndices loads the list of stored batch indices
func (s *BatchStorage) loadBatchIndices() ([]uint64, error) {
	encoded, err := s.db.GetBytes([]byte(SealedBatchIndicesKey))
	if err != nil {
		return nil, err
	}

	var indices []uint64
	if err := json.Unmarshal(encoded, &indices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal batch indices: %w", err)
	}

	return indices, nil
}

// saveBatchIndices saves the list of batch indices
func (s *BatchStorage) saveBatchIndices(indices []uint64) error {
	encoded, err := json.Marshal(indices)
	if err != nil {
		return fmt.Errorf("failed to marshal batch indices: %w", err)
	}

	return s.db.PutBytes([]byte(SealedBatchIndicesKey), encoded)
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
		if errors.Is(err, db.ErrKeyNotFound) {
			return nil, fmt.Errorf("sealed batch header %d not found", batchIndex)
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
		if errors.Is(err, db.ErrKeyNotFound) {
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
