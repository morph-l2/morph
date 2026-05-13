package derivation

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/go-kit/kit/metrics/discard"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// fakeTagL2Client implements tagL2Client for unit tests. It records each
// SetBlockTags call so tests can assert on call count and arguments, and
// lets the test set the unsafe upper bound returned by BlockNumber.
type fakeTagL2Client struct {
	unsafe  uint64
	blockNumberErr error
	calls   []setBlockTagsCall
	setErr  error
}

type setBlockTagsCall struct {
	safe      common.Hash
	finalized common.Hash
}

func (f *fakeTagL2Client) BlockNumber(_ context.Context) (uint64, error) {
	if f.blockNumberErr != nil {
		return 0, f.blockNumberErr
	}
	return f.unsafe, nil
}

func (f *fakeTagL2Client) SetBlockTags(_ context.Context, safe common.Hash, finalized common.Hash) error {
	if f.setErr != nil {
		return f.setErr
	}
	f.calls = append(f.calls, setBlockTagsCall{safe: safe, finalized: finalized})
	return nil
}

// newDiscardMetrics returns a *Metrics whose collectors discard all updates.
// Avoids prometheus default-registry double-registration across multiple
// tests in the same process.
func newDiscardMetrics() *Metrics {
	return &Metrics{
		L1SyncHeight:               discard.NewGauge(),
		RollupL2Height:             discard.NewGauge(),
		DeriveL2Height:             discard.NewGauge(),
		BatchStatus:                discard.NewGauge(),
		LatestBatchIndex:           discard.NewGauge(),
		SyncedBatchIndex:           discard.NewGauge(),
		PathBTriggered:             discard.NewCounter(),
		PathBFailed:                discard.NewCounter(),
		SafeAdvanceTotal:           discard.NewCounter(),
		FinalizedAdvanceTotal:      discard.NewCounter(),
		SafeL2BlockNumber:          discard.NewGauge(),
		FinalizedL2BlockNumber:     discard.NewGauge(),
		L1ReorgResetTotal:          discard.NewCounter(),
		TagInvariantViolationTotal: discard.NewCounter(),
	}
}

func newTestTagAdvancer(t *testing.T, unsafe uint64) (*tagAdvancer, *fakeTagL2Client, *Metrics) {
	t.Helper()
	fake := &fakeTagL2Client{unsafe: unsafe}
	m := newDiscardMetrics()
	logger := tmlog.NewNopLogger()
	return newTagAdvancer(fake, m, logger), fake, m
}

func headerAt(num uint64, mark byte) *eth.Header {
	h := &eth.Header{Number: new(big.Int).SetUint64(num)}
	// Mutate ParentHash so different "mark" values produce different block
	// hashes -- header.Hash() mixes everything.
	h.ParentHash = common.BytesToHash([]byte{mark, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return h
}

func TestTagAdvance_Safe_CallsSetBlockTags(t *testing.T) {
	tagAdv, fake, _ := newTestTagAdvancer(t, 100)
	h := headerAt(50, 'a')

	tagAdv.advanceSafe(context.Background(), 7, h)

	if len(fake.calls) != 1 {
		t.Fatalf("expected 1 SetBlockTags call, got %d", len(fake.calls))
	}
	if fake.calls[0].safe != h.Hash() {
		t.Fatalf("safe hash mismatch")
	}
	if tagAdv.SafeMaxBatchIndex() != 7 {
		t.Fatalf("safeMaxBatchIndex got %d, want 7", tagAdv.SafeMaxBatchIndex())
	}
}

func TestTagAdvance_DedupSetBlockTags(t *testing.T) {
	tagAdv, fake, _ := newTestTagAdvancer(t, 100)
	h := headerAt(50, 'a')

	tagAdv.advanceSafe(context.Background(), 7, h)
	tagAdv.advanceSafe(context.Background(), 7, h) // identical state

	if len(fake.calls) != 1 {
		t.Fatalf("expected dedup to suppress 2nd call; got %d total", len(fake.calls))
	}
}

func TestTagAdvance_InvariantSafeGtUnsafe_Skips(t *testing.T) {
	tagAdv, fake, _ := newTestTagAdvancer(t, 30) // unsafe = 30
	h := headerAt(50, 'a')                         // safe wants 50 -- invalid

	tagAdv.advanceSafe(context.Background(), 7, h)

	if len(fake.calls) != 0 {
		t.Fatalf("expected SetBlockTags skipped on invariant violation, got %d calls", len(fake.calls))
	}
}

func TestTagAdvance_InvariantFinalizedGtSafe_Skips(t *testing.T) {
	tagAdv, fake, _ := newTestTagAdvancer(t, 200)

	// safe at 50, finalized would be 80 -> finalized > safe.
	tagAdv.advanceSafe(context.Background(), 5, headerAt(50, 'a'))
	// reset the call recorder so we only inspect the finalized call.
	fake.calls = nil

	tagAdv.advanceFinalized(context.Background(), 6, headerAt(80, 'b'))

	if len(fake.calls) != 0 {
		t.Fatalf("expected SetBlockTags skipped on finalized > safe; got %d calls", len(fake.calls))
	}
}

func TestTagAdvance_FinalizedMonotonic(t *testing.T) {
	tagAdv, fake, _ := newTestTagAdvancer(t, 200)
	tagAdv.advanceSafe(context.Background(), 10, headerAt(120, 'a'))
	fake.calls = nil

	tagAdv.advanceFinalized(context.Background(), 8, headerAt(100, 'b'))
	if got := tagAdv.finalizedL2Number; got != 100 {
		t.Fatalf("finalized first advance: got %d, want 100", got)
	}

	// Second advance with smaller number should be ignored.
	prevHash := tagAdv.finalizedL2Hash
	tagAdv.advanceFinalized(context.Background(), 7, headerAt(80, 'c'))
	if tagAdv.finalizedL2Number != 100 || tagAdv.finalizedL2Hash != prevHash {
		t.Fatalf("finalized regressed: number=%d, hash unchanged=%v",
			tagAdv.finalizedL2Number, tagAdv.finalizedL2Hash == prevHash)
	}
}

func TestTagAdvance_L1ReorgReset(t *testing.T) {
	tagAdv, _, _ := newTestTagAdvancer(t, 200)
	tagAdv.advanceSafe(context.Background(), 10, headerAt(120, 'a'))

	tagAdv.reset(8)

	if tagAdv.safeL2Number != 0 {
		t.Fatalf("safeL2Number not cleared after reset: got %d", tagAdv.safeL2Number)
	}
	if tagAdv.safeL2Hash != (common.Hash{}) {
		t.Fatalf("safeL2Hash not cleared after reset")
	}
	if got := tagAdv.SafeMaxBatchIndex(); got != 8 {
		t.Fatalf("safeMaxBatchIndex after reset: got %d, want 8", got)
	}
	if tagAdv.lastNotifiedSafe != (common.Hash{}) {
		t.Fatalf("lastNotifiedSafe not cleared after reset")
	}
}

func TestTagAdvance_BlockNumberError_SkipsFlush(t *testing.T) {
	tagAdv, fake, _ := newTestTagAdvancer(t, 100)
	fake.blockNumberErr = errors.New("rpc down")

	tagAdv.advanceSafe(context.Background(), 7, headerAt(50, 'a'))

	if len(fake.calls) != 0 {
		t.Fatalf("expected SetBlockTags skipped when BlockNumber fails; got %d", len(fake.calls))
	}
}
