package event

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type IEventStorage interface {
	Store() error
	Load() error
}

var (
	defaultFilename = "event_info.json"
)

type EventInfo struct {
	BlockTime      uint64 `json:"block_time"`      // event emit time
	BlockProcessed uint64 `json:"block_processed"` // block processed
}

type EventInfoStorage struct {
	EventInfo
	Filename string `json:"filename"` // filename
	lock     sync.Mutex
}

func NewEventInfoStorage(filename string) *EventInfoStorage {
	if filename == "" {
		filename = defaultFilename
	}
	return &EventInfoStorage{
		Filename: filename,
	}
}

func (e *EventInfoStorage) Store() error {

	// Convert struct to JSON string
	jsonData, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("failed to convert struct to JSON: %w", err)
	}

	e.lock.Lock()
	defer e.lock.Unlock()
	// Write JSON data to file
	err = os.WriteFile(e.Filename, jsonData, 0600)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
func (e *EventInfoStorage) Load() error {
	e.lock.Lock()
	defer e.lock.Unlock()
	jsonData, err := os.ReadFile(e.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			e.EventInfo = EventInfo{}
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	// parse json data to struct
	err = json.Unmarshal(jsonData, e)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}
