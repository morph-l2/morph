package derivation

import "fmt"

// L2HeadStage represents the public-facing safety level of an L2 head per SPEC-005.
//
// State semantics:
//   - StageUnsafe:    Block executed locally; not yet anchored to any L1 batch.
//   - StageSafe:      Anchored to an L1 batch found on L1 `safe`; subject to rollback
//                     if the L1 batch reorgs out or batch verification fails.
//   - StageFinalized: Anchored to an L1 batch whose origin is on L1 `finalized`.
//                     Monotonic; never rolls back.
//   - StageHalted:    Unrecoverable inconsistency (e.g. second batch-root mismatch
//                     after rollback, or a rollback target below FinalizedHead).
//                     Derivation refuses to advance until manual intervention.
//
// A node always advertises a single stage per head (one each for safe / finalized);
// halted is global to the derivation pipeline.
type L2HeadStage uint8

const (
	StageUnsafe L2HeadStage = iota
	StageSafe
	StageFinalized
	StageHalted
)

func (s L2HeadStage) String() string {
	switch s {
	case StageUnsafe:
		return "unsafe"
	case StageSafe:
		return "safe"
	case StageFinalized:
		return "finalized"
	case StageHalted:
		return "halted"
	default:
		return fmt.Sprintf("unknown(%d)", uint8(s))
	}
}

// HeadAnchor pairs an L2 head with the L1 origin that justifies its current
// safety stage. Both safe_head and finalized_head are persisted as HeadAnchor
// to allow detecting L1 reorgs that invalidate previously recorded anchors.
type HeadAnchor struct {
	L2Number uint64
	L2Hash   [32]byte
	L1Number uint64
	L1Hash   [32]byte
}

// IsZero reports whether the anchor is uninitialized (e.g. at first node start
// before the first derivation loop has succeeded).
func (a HeadAnchor) IsZero() bool {
	return a.L2Number == 0 && a.L1Number == 0
}
