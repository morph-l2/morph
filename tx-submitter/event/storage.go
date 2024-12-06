package event

import (
	"encoding/json"
	"fmt"
	"sync"

	"morph-l2/tx-submitter/db"
	"morph-l2/tx-submitter/params"
	"morph-l2/tx-submitter/utils"
)

type IEventStorage interface {
	Store() error
	Load() error
}

type EventInfo struct {
	BlockTime      uint64 `json:"block_time"`      // event emit time
	BlockProcessed uint64 `json:"block_processed"` // block processed
}

type EventInfoStorage struct {
	EventInfo
	db *db.Db
	mu sync.Mutex
}

func NewEventInfoStorage(db *db.Db) *EventInfoStorage {
	return &EventInfoStorage{
		db: db,
	}
}

func (e *EventInfoStorage) Store() error {

	// Convert struct to JSON string
	jsonData, err := json.Marshal(e.EventInfo)
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
	evnetInfo, err := e.db.GetString(params.EventInfoKey)
	if err != nil {
		if utils.ErrStringMatch(err, db.ErrKeyNotFound) {
			e.EventInfo = EventInfo{}
			jsonData, err := json.Marshal(e.EventInfo)
			if err != nil {
				return fmt.Errorf("failed to marshal json: %w", err)
			}
			err = e.db.PutString(params.EventInfoKey, string(jsonData))
			if err != nil {
				return fmt.Errorf("failed to init eventinfo to db: %w", err)
			}
			return nil
		}
		return fmt.Errorf("failed to load eventinfo from db: %w", err)
	}

	// parse json data to struct
	err = json.Unmarshal([]byte(evnetInfo), &e.EventInfo)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}
