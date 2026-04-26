package constants

const (
	// MethodCommitBatch is the method name for committing a batch (with blob when applicable)
	MethodCommitBatch = "commitBatch"
	// MethodCommitState is the method name for recommitting batch state using stored blob hash (no blob in tx)
	MethodCommitState = "commitState"
	// MethodFinalizeBatch is the method name for finalizing a batch
	MethodFinalizeBatch = "finalizeBatch"
)

// IsCommitLikeMethod returns true for commitBatch or commitState (same calldata shape for batch index parsing).
func IsCommitLikeMethod(method string) bool {
	return method == MethodCommitBatch || method == MethodCommitState
}
