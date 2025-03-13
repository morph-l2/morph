package mock

import (
	"context"
	"sync"

	"morph-l2/tx-submitter/iface"
)

// MockReorgDetector implements the IReorgDetector interface for testing
type MockReorgDetector struct {
	mu sync.RWMutex

	// Mock return values
	detectReorgReturn struct {
		hasReorg   bool
		reorgDepth uint64
		err        error
	}
}

// Ensure MockReorgDetector implements IReorgDetector
var _ iface.IReorgDetector = (*MockReorgDetector)(nil)

// NewMockReorgDetector creates a new instance of MockReorgDetector
func NewMockReorgDetector() *MockReorgDetector {
	return &MockReorgDetector{}
}

// DetectReorg implements IReorgDetector.DetectReorg
func (m *MockReorgDetector) DetectReorg(ctx context.Context) (bool, uint64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.detectReorgReturn.hasReorg, m.detectReorgReturn.reorgDepth, m.detectReorgReturn.err
}

// SetDetectReorgReturn sets the return values for the DetectReorg method
func (m *MockReorgDetector) SetDetectReorgReturn(hasReorg bool, reorgDepth uint64, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.detectReorgReturn.hasReorg = hasReorg
	m.detectReorgReturn.reorgDepth = reorgDepth
	m.detectReorgReturn.err = err
}
