package l1sequencer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
)

const (
	// CacheTTL is the time-to-live for the sequencer verifier cache
	//CacheTTL = 30 * time.Minute
	CacheTTL = 10 * time.Second
)

// SequencerVerifier verifies L1 sequencer status with caching.
// It provides IsSequencer() for checking if an address is the current sequencer.
type SequencerVerifier struct {
	mutex       sync.Mutex
	sequencer   common.Address
	cacheExpiry time.Time

	caller *bindings.L1SequencerCaller
	logger tmlog.Logger
}

// NewSequencerVerifier creates a new SequencerVerifier
func NewSequencerVerifier(caller *bindings.L1SequencerCaller, logger tmlog.Logger) *SequencerVerifier {
	return &SequencerVerifier{
		caller: caller,
		logger: logger.With("module", "l1sequencer_verifier"),
	}
}

// flushCache refreshes the cache (caller must hold the lock)
func (c *SequencerVerifier) flushCache(ctx context.Context) error {
	newSeq, err := c.caller.GetSequencer(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get sequencer from L1: %w", err)
	}

	if c.sequencer != newSeq {
		c.logger.Info("Sequencer address updated",
			"old", c.sequencer.Hex(),
			"new", newSeq.Hex())
	}

	c.sequencer = newSeq
	c.cacheExpiry = time.Now().Add(CacheTTL)
	return nil
}

// IsSequencer checks if the given address is the current sequencer.
// It uses lazy loading: refreshes cache if expired, and retries on miss.
func (c *SequencerVerifier) IsSequencer(ctx context.Context, addr common.Address) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Cache expired, refresh
	if time.Now().After(c.cacheExpiry) {
		if err := c.flushCache(ctx); err != nil {
			return false, err
		}
	}

	// Cache hit
	if c.sequencer == addr {
		return true, nil
	}

	// Cache miss - maybe sequencer just updated, force refresh once
	if err := c.flushCache(ctx); err != nil {
		return false, err
	}

	return c.sequencer == addr, nil
}

// GetSequencer returns the cached sequencer address (refreshes if expired)
func (c *SequencerVerifier) GetSequencer(ctx context.Context) (common.Address, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if time.Now().After(c.cacheExpiry) {
		if err := c.flushCache(ctx); err != nil {
			return common.Address{}, err
		}
	}

	return c.sequencer, nil
}
