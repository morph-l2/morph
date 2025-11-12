package updater

import "context"

// Updater defines the interface for all update tasks
type Updater interface {
	// Name returns the name of the updater
	Name() string
	
	// Start starts the updater goroutine
	Start(ctx context.Context) error
	
	// Stop gracefully stops the updater
	Stop() error
}

