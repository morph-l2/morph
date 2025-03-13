package event

import (
	"encoding/json"
	"fmt"
	"sync"

	"morph-l2/tx-submitter/db"
	"morph-l2/tx-submitter/params"
)

type IEventStorage interface {
	Store() error
	Load() error
	BlockProcessed() uint64
	SetBlockProcessed(blockNum uint64)
	BlockTime() uint64
	SetBlockTime(blockTime uint64)
	EventInfo() *EventInfo
}

type EventInfo struct {
	BlockTime      uint64 `json:"block_time"`      // event emit time
	BlockProcessed uint64 `json:"block_processed"` // block processed
}

type EventInfoStorage struct {
	eventInfo EventInfo
	db        db.Database
	mu        sync.RWMutex
}

func NewEventInfoStorage(db db.Database) *EventInfoStorage {
	return &EventInfoStorage{
		db: db,
	}
}

func (e *EventInfoStorage) Store() error {
	// Convert struct to JSON string
	jsonData, err := json.Marshal(e.eventInfo)
	if err != nil {
		return fmt.Errorf("failed to convert struct to JSON: %w", err)
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	// Write JSON data to file
	err = e.db.PutString(params.EventInfoKey, string(jsonData))
	if err != nil {
		return fmt.Errorf("failed to write to db: err=%w, data=%v", err, jsonData)
	}
	return nil
}

func (e *EventInfoStorage) Load() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	jsonStr, err := e.db.GetString(params.EventInfoKey)
	if err != nil {
		if err == db.ErrKeyNotFound {
			// Initialize with default values if not found
			e.eventInfo = EventInfo{}
			return nil
		}
		return fmt.Errorf("failed to read from db: %w", err)
	}

	err = json.Unmarshal([]byte(jsonStr), &e.eventInfo)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
}

func (e *EventInfoStorage) BlockProcessed() uint64 {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.eventInfo.BlockProcessed
}

func (e *EventInfoStorage) SetBlockProcessed(blockNum uint64) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.eventInfo.BlockProcessed = blockNum
}

func (e *EventInfoStorage) BlockTime() uint64 {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.eventInfo.BlockTime
}

func (e *EventInfoStorage) SetBlockTime(blockTime uint64) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.eventInfo.BlockTime = blockTime
}

func (e *EventInfoStorage) EventInfo() EventInfo {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.eventInfo
}
