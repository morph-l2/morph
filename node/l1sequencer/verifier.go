package l1sequencer

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
)

const refreshInterval = 5 * time.Minute

// sequencerCursor caches the current sequencer interval for O(1) lookup.
type sequencerCursor struct {
	from  uint64
	to    uint64 // exclusive; math.MaxUint64 = no upper bound
	addr  common.Address
	valid bool
}

// SequencerVerifier verifies L1 sequencer status.
// Implements tendermint SequencerVerifier interface.
//
// History is loaded from L1 at construction and refreshed every 5 minutes.
// All L1 reads use the finalized block tag to avoid ingesting reorged data.
type SequencerVerifier struct {
	mu     sync.Mutex
	history []bindings.L1SequencerHistoryRecord
	cursor  sequencerCursor

	caller *bindings.L1SequencerCaller
	logger tmlog.Logger
	cancel context.CancelFunc
}

// NewSequencerVerifier creates a new SequencerVerifier, loads the full sequencer
// history from L1 (finalized), and starts a background refresh goroutine.
// Call Stop to terminate the background loop.
func NewSequencerVerifier(caller *bindings.L1SequencerCaller, logger tmlog.Logger) *SequencerVerifier {
	ctx, cancel := context.WithCancel(context.Background())
	v := &SequencerVerifier{
		caller: caller,
		logger: logger.With("module", "l1sequencer_verifier"),
		cancel: cancel,
	}
	if err := v.syncHistory(); err != nil {
		v.logger.Error("Failed to load sequencer history from L1", "err", err)
	}
	go v.refreshLoop(ctx)
	return v
}

// Stop terminates the background refresh loop.
func (c *SequencerVerifier) Stop() {
	c.cancel()
}

// syncHistory fetches the full sequencer history from L1 (finalized tag) and
// replaces the local cache if anything changed.
func (c *SequencerVerifier) syncHistory() error {
	raw, err := c.caller.GetSequencerHistory(&bind.CallOpts{
		BlockNumber: big.NewInt(int64(rpc.FinalizedBlockNumber)),
	})
	if err != nil {
		return fmt.Errorf("GetSequencerHistory: %w", err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if len(raw) == len(c.history) {
		return nil // no change
	}

	prev := len(c.history)
	c.history = raw
	// Only invalidate cursor if it was pointing at the last record (to == MaxUint64),
	// because new records change that interval's upper bound.
	// Existing records never change, so earlier cursor intervals remain valid.
	if c.cursor.valid && c.cursor.to == math.MaxUint64 {
		c.cursor.valid = false
	}

	// Log new records
	for i := prev; i < len(c.history); i++ {
		c.logger.Info("Sequencer record",
			"startL2Block", c.history[i].StartL2Block,
			"address", c.history[i].SequencerAddr.Hex())
	}
	c.logger.Info("Sequencer history synced", "total", len(c.history), "new", len(c.history)-prev)
	return nil
}

// refreshLoop polls L1 until ctx is cancelled.
// Uses exponential backoff (10s -> 20s -> ... -> 5min) while history is empty,
// then switches to the normal 5-minute interval once loaded.
func (c *SequencerVerifier) refreshLoop(ctx context.Context) {
	const minRetry = 10 * time.Second

	interval := refreshInterval
	c.mu.Lock()
	empty := len(c.history) == 0
	c.mu.Unlock()
	if empty {
		interval = minRetry
	}

	timer := time.NewTimer(interval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := c.syncHistory(); err != nil {
				c.logger.Error("Failed to refresh sequencer history", "err", err)
			}

			c.mu.Lock()
			empty = len(c.history) == 0
			c.mu.Unlock()

			if empty {
				// Exponential backoff, capped at refreshInterval
				interval = interval * 2
				if interval > refreshInterval {
					interval = refreshInterval
				}
			} else {
				interval = refreshInterval
			}
			timer.Reset(interval)
		}
	}
}

// SequencerAtHeight returns the sequencer address at the given L2 height.
func (c *SequencerVerifier) SequencerAtHeight(l2Height uint64) (common.Address, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.history) == 0 {
		return common.Address{}, false
	}

	if c.cursor.valid && l2Height >= c.cursor.from && l2Height < c.cursor.to {
		return c.cursor.addr, true
	}

	idx := sort.Search(len(c.history), func(i int) bool {
		return c.history[i].StartL2Block > l2Height
	}) - 1
	if idx < 0 {
		return common.Address{}, false
	}

	c.cursor.from = c.history[idx].StartL2Block
	if idx+1 < len(c.history) {
		c.cursor.to = c.history[idx+1].StartL2Block
	} else {
		c.cursor.to = math.MaxUint64
	}
	c.cursor.addr = c.history[idx].SequencerAddr
	c.cursor.valid = true
	return c.cursor.addr, true
}

// ============================================================================
// Interface implementation
// ============================================================================

// IsSequencerAt checks if addr was the sequencer at the given L2 height.
func (c *SequencerVerifier) IsSequencerAt(addr common.Address, l2Height uint64) (bool, error) {
	histAddr, found := c.SequencerAtHeight(l2Height)
	if !found {
		return false, fmt.Errorf("no sequencer record for height %d", l2Height)
	}
	return addr == histAddr, nil
}

// VerificationStartHeight returns history[0].StartL2Block (= contract activeHeight).
// Returns math.MaxUint64 if history is empty.
func (c *SequencerVerifier) VerificationStartHeight() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.history) == 0 {
		return math.MaxUint64
	}
	return c.history[0].StartL2Block
}
