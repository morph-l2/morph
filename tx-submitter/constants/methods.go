package constants

const (
	// MethodCommitBatch is the method name for committing a batch (with blob when applicable)
	MethodCommitBatch = "commitBatch"
	// MethodFinalizeBatch is the method name for finalizing a batch
	MethodFinalizeBatch = "finalizeBatch"
)

func IsOwnedCommitMethod(method string) bool {
	return method == MethodCommitBatch
}
