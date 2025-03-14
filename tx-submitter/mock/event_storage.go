package mock

import (
	"morph-l2/tx-submitter/event"
)

// MockEventInfoStorage is a mock implementation of the EventInfoStorage
type MockEventInfoStorage struct {
	event.EventInfoStorage
}

// NewMockEventInfoStorage creates a new mock EventInfoStorage
func NewMockEventInfoStorage() *event.EventInfoStorage {
	mockDB := NewMockDB()
	return event.NewEventInfoStorage(mockDB)
}

// Store implements the Store method
func (m *MockEventInfoStorage) Store() error {
	return nil
}

// Load implements the Load method
func (m *MockEventInfoStorage) Load() error {
	return nil
}
