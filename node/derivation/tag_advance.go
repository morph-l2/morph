package derivation

import (
	"context"
	"sync"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// tagL2Client narrows the dependency on types.RetryableClient to the two
// methods the tag advancer actually calls. Keeping this local makes
// tagAdvancer trivially mockable from tests without dragging in an authclient
// stack.
type tagL2Client interface {
	BlockNumber(ctx context.Context) (uint64, error)
	SetBlockTags(ctx context.Context, safe common.Hash, finalized common.Hash) error
}

// tagAdvancer is the SPEC-005 section 4.7 single source of truth for safe and
// finalized L2 head propagation. It replaces the previous standalone
// polling service: derivation main loop drives `advanceSafe` per
// verified batch; the in-process finalizer subcomponent drives
// `advanceFinalized`. Both paths converge on `flushTags` which enforces the
// `finalized <= safe <= unsafe` invariant before calling the existing
// `RetryableClient.SetBlockTags` engine RPC.
//
// In-memory only by design: SPEC-005 section 4.7.7 -- restart starts from zero and
// derivation refills naturally as it walks its cursor.
type tagAdvancer struct {
	mu sync.Mutex

	l2Client tagL2Client
	metrics  *Metrics
	logger   tmlog.Logger

	// safe head -- last verified batch's lastL2Block.
	safeL2Hash        common.Hash
	safeL2Number      uint64
	safeMaxBatchIndex uint64

	// finalized head -- L1 finalized derived verified batch's lastL2Block.
	finalizedL2Hash   common.Hash
	finalizedL2Number uint64

	// Suppress redundant SetBlockTags RPCs (mirrors blocktag's
	// lastNotifiedSafeHash / lastNotifiedFinalizedHash semantics).
	lastNotifiedSafe      common.Hash
	lastNotifiedFinalized common.Hash
}

func newTagAdvancer(l2Client tagL2Client, metrics *Metrics, logger tmlog.Logger) *tagAdvancer {
	return &tagAdvancer{
		l2Client: l2Client,
		metrics:  metrics,
		logger:   logger.With("component", "tag-advancer"),
	}
}

// advanceSafe is called by the derivation main loop after a batch passes both
// content verification (Path A or Path B) and verifyBatchRoots. It records the
// new safe head and flushes via SetBlockTags.
func (t *tagAdvancer) advanceSafe(ctx context.Context, batchIndex uint64, lastHeader *eth.Header) {
	if lastHeader == nil {
		return
	}
	t.mu.Lock()
	t.safeL2Hash = lastHeader.Hash()
	t.safeL2Number = lastHeader.Number.Uint64()
	if batchIndex > t.safeMaxBatchIndex {
		t.safeMaxBatchIndex = batchIndex
	}
	t.metrics.IncSafeAdvance()
	t.metrics.SetSafeL2BlockNumber(t.safeL2Number)
	t.mu.Unlock()

	t.flushTags(ctx)
}

// advanceFinalized is called by the finalizer subcomponent each tick if the
// L1 finalized block produces a new finalized batch <= safeMaxBatchIndex.
// finalized never moves backwards; if a lower number is provided we log and
// keep the previous value (SPEC-005 section 4.7.4 monotonicity check).
func (t *tagAdvancer) advanceFinalized(ctx context.Context, batchIndex uint64, lastHeader *eth.Header) {
	if lastHeader == nil {
		return
	}
	t.mu.Lock()
	newNumber := lastHeader.Number.Uint64()
	if t.finalizedL2Number != 0 && newNumber < t.finalizedL2Number {
		t.logger.Error("finalized monotonicity violated; ignoring",
			"prev", t.finalizedL2Number, "next", newNumber)
		t.mu.Unlock()
		return
	}
	if newNumber == t.finalizedL2Number && lastHeader.Hash() == t.finalizedL2Hash {
		t.mu.Unlock()
		return
	}
	t.finalizedL2Hash = lastHeader.Hash()
	t.finalizedL2Number = newNumber
	t.metrics.IncFinalizedAdvance()
	t.metrics.SetFinalizedL2BlockNumber(t.finalizedL2Number)
	t.mu.Unlock()

	_ = batchIndex // currently logged by the finalizer; reserved for future telemetry
	t.flushTags(ctx)
}

// SafeMaxBatchIndex returns the highest verified batch index recorded so far,
// for the finalizer to take min(L1 finalized batch, safe).
func (t *tagAdvancer) SafeMaxBatchIndex() uint64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.safeMaxBatchIndex
}

// reset clears safe head when the derivation main loop detects an L1 reorg
// and rewinds its cursor. finalized is intentionally NOT reset -- see
// SPEC-005 section 4.7.6: L1 finalized is assumed monotonic, and finalizer.tick will
// re-evaluate on the next iteration.
func (t *tagAdvancer) reset(toBatchIndex uint64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.safeL2Hash = common.Hash{}
	t.safeL2Number = 0
	t.safeMaxBatchIndex = toBatchIndex
	t.lastNotifiedSafe = common.Hash{}
	t.metrics.IncL1ReorgReset()
	t.metrics.SetSafeL2BlockNumber(0)
	t.logger.Info("tag advancer reset on L1 reorg", "to_batch_index", toBatchIndex)
}

// flushTags enforces the finalized <= safe <= unsafe invariant and calls
// SetBlockTags exactly once per state change. On invariant violation we log
// error and skip -- no panic, no halt -- matching op-node's
// tryUpdateEngineInternal behaviour.
func (t *tagAdvancer) flushTags(ctx context.Context) {
	unsafeNum, err := t.l2Client.BlockNumber(ctx)
	if err != nil {
		t.logger.Info("flushTags: read L2 latest failed", "err", err)
		return
	}

	t.mu.Lock()
	safeHash := t.safeL2Hash
	safeNum := t.safeL2Number
	finalizedHash := t.finalizedL2Hash
	finalizedNum := t.finalizedL2Number
	notifiedSafe := t.lastNotifiedSafe
	notifiedFinalized := t.lastNotifiedFinalized
	t.mu.Unlock()

	if finalizedNum > safeNum {
		t.metrics.IncTagInvariantViolation()
		t.logger.Error("invariant violation: finalized > safe",
			"finalized", finalizedNum, "safe", safeNum)
		return
	}
	if safeNum > unsafeNum {
		t.metrics.IncTagInvariantViolation()
		t.logger.Error("invariant violation: safe > unsafe",
			"safe", safeNum, "unsafe", unsafeNum)
		return
	}

	if safeHash == notifiedSafe && finalizedHash == notifiedFinalized {
		return
	}
	if safeHash == (common.Hash{}) && finalizedHash == (common.Hash{}) {
		return
	}

	if err := t.l2Client.SetBlockTags(ctx, safeHash, finalizedHash); err != nil {
		t.logger.Error("SetBlockTags failed", "err", err)
		return
	}

	t.mu.Lock()
	t.lastNotifiedSafe = safeHash
	t.lastNotifiedFinalized = finalizedHash
	t.mu.Unlock()
}
