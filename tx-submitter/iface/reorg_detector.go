package iface

import "context"

// IReorgDetector defines the interface for chain reorganization detection
type IReorgDetector interface {
	// DetectReorg checks if a chain reorganization has occurred
	// Returns:
	// - bool: whether a reorg was detected
	// - uint64: the depth of the reorg (number of blocks from head)
	// - error: any error that occurred during detection
	DetectReorg(ctx context.Context) (bool, uint64, error)
}
