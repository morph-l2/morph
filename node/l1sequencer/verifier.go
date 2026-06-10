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
	"github.com/tendermint/tendermint/upgrade"

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
//
// Startup is fail-fast: if the initial syncHistory() fails or returns empty
// history, the global upgrade.UpgradeBlockHeight is never set and stays at its
// -1 sentinel. Running with UpgradeBlockHeight=-1 is unsafe: IsUpgraded()
// returns false for every height, so the PBFT state machine can run past the
// true upgrade height, and on a block-synced fullnode restart the handshake's
// sequencer-mode replay exemption fails (ErrAppBlockHeightTooHigh) while
// reconstructLastCommit runs against a nil commit and panics. We therefore
// refuse to start rather than boot into that state; the operator's supervisor
// restarts us once L1 is reachable.
func NewSequencerVerifier(caller *bindings.L1SequencerCaller, logger tmlog.Logger) (*SequencerVerifier, error) {
	ctx, cancel := context.WithCancel(context.Background())
	v := &SequencerVerifier{
		caller: caller,
		logger: logger.With("module", "l1sequencer_verifier"),
		cancel: cancel,
	}
	if err := v.syncHistory(); err != nil {
		cancel()
		return nil, fmt.Errorf("refusing to start with UpgradeBlockHeight=-1: initial sequencer history sync from L1 failed: %w", err)
	}
	if upgrade.UpgradeBlockHeight < 0 {
		cancel()
		return nil, fmt.Errorf("refusing to start with UpgradeBlockHeight=-1: L1 returned empty sequencer history; upgrade height unknown")
	}
	v.logCurrentState()
	go v.refreshLoop(ctx)
	return v, nil
}

// logCurrentState prints a one-line snapshot of the loaded contract state at
// startup: the active sequencer and the height from which verification applies.
// Existing logs list every record but not which one is currently in effect.
func (c *SequencerVerifier) logCurrentState() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.history) == 0 {
		c.logger.Info("Sequencer contract state: no records loaded; verification inactive until L1 history is available")
		return
	}
	current := c.history[len(c.history)-1]
	c.logger.Info("Sequencer contract state loaded",
		"totalRecords", len(c.history),
		"verificationStartHeight", c.history[0].StartL2Block,
		"currentSequencer", current.SequencerAddr.Hex(),
		"currentSequencerStartHeight", current.StartL2Block)
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
	// Set upgrade height from L1 contract on first successful load
	if prev == 0 && len(c.history) > 0 {
		height := int64(c.history[0].StartL2Block)
		upgrade.SetUpgradeBlockHeight(height)
		c.logger.Info("Upgrade height set from L1 contract", "height", height)
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
