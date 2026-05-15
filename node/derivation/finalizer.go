package derivation

import (
	"context"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
)

// finalizer is the SPEC-005 section 4.7.4 finalized-head subcomponent. It runs as an
// in-process goroutine inside Derivation (not a standalone service): each
// tick it reads L1 finalized -> Rollup.LastCommittedBatchIndex(@finalized),
// takes min with the highest verified batch index recorded by tagAdvancer,
// resolves the corresponding L2 header from tagAdvancer's local
// verified-batch map, and forwards to tagAdvancer.advanceFinalized.
//
// The local map replaces what used to be a Rollup.BatchDataStore lookup. The
// contract clears storage of older batches as part of its on-chain GC, so any
// candidate older than the very latest committed batch returned zero on hoodi
// (BatchDataStore(17389) and (17796) zero, only (17797) populated). With the
// in-memory mapping the lookup is independent of contract retention; if a
// candidate is missing from the map (e.g. derivation hasn't re-verified it
// since restart) we log info and retry next tick.
//
// Cheap relative to derivation main loop: one L1 header + one contract call
// per tick (default 30s).
type finalizer struct {
	ctx      context.Context
	interval time.Duration
	logger   tmlog.Logger

	l1Client    *ethclient.Client
	rollup      *bindings.Rollup
	tagAdvancer *tagAdvancer

	stopped chan struct{}
}

func newFinalizer(
	ctx context.Context,
	interval time.Duration,
	l1Client *ethclient.Client,
	rollup *bindings.Rollup,
	tagAdv *tagAdvancer,
	logger tmlog.Logger,
) *finalizer {
	return &finalizer{
		ctx:         ctx,
		interval:    interval,
		l1Client:    l1Client,
		rollup:      rollup,
		tagAdvancer: tagAdv,
		logger:      logger.With("component", "finalizer"),
		stopped:     make(chan struct{}),
	}
}

func (f *finalizer) run() {
	defer close(f.stopped)

	t := time.NewTicker(f.interval)
	defer t.Stop()

	// Run once immediately so the first tag flush doesn't wait a full
	// interval after startup; matches blocktag's `initialize()` behavior.
	f.tick()

	for {
		select {
		case <-f.ctx.Done():
			return
		case <-t.C:
			f.tick()
		}
	}
}

func (f *finalizer) tick() {
	// 1. Resolve the L1 finalized header.
	finHeader, err := f.l1Client.HeaderByNumber(f.ctx, big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		f.logger.Info("finalizer: read L1 finalized header failed", "err", err)
		return
	}
	if finHeader == nil {
		return
	}

	// 2. Query Rollup.LastCommittedBatchIndex pinned at that L1 block.
	maxCommittedAtFin, err := f.rollup.LastCommittedBatchIndex(&bind.CallOpts{
		BlockNumber: finHeader.Number,
		Context:     f.ctx,
	})
	if err != nil {
		f.logger.Info("finalizer: query LastCommittedBatchIndex@finalized failed",
			"l1Block", finHeader.Number.Uint64(), "err", err)
		return
	}
	if maxCommittedAtFin == nil {
		return
	}

	// 3. Take min with the highest verified batch index recorded by tagAdvancer.
	verifiedMax := f.tagAdvancer.SafeMaxBatchIndex()
	if verifiedMax == 0 {
		// derivation hasn't yet verified any batch this run; nothing to finalize.
		return
	}
	candidate := maxCommittedAtFin.Uint64()
	if verifiedMax < candidate {
		candidate = verifiedMax
	}
	if candidate == 0 {
		return
	}

	// 4. Resolve candidate -> lastL2Block header via tagAdvancer's local map.
	// Missing entry is expected during the catch-up window after restart
	// (derivation hasn't re-verified that index yet) and resolves on the next
	// tick once derivation walks past it.
	header, ok := f.tagAdvancer.LookupVerifiedBatchHeader(candidate)
	if !ok {
		f.logger.Info("finalizer: verified batch header not found in local map; will retry",
			"batchIndex", candidate, "verifiedMax", verifiedMax,
			"maxCommittedAtFin", maxCommittedAtFin.Uint64())
		return
	}

	f.tagAdvancer.advanceFinalized(f.ctx, candidate, header)
}
