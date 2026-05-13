package derivation

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	"morph-l2/node/types"
)

// finalizer is the SPEC-005 section 4.7.4 finalized-head subcomponent. It runs as an
// in-process goroutine inside Derivation (not a standalone service): each
// tick it reads L1 finalized -> Rollup.LastCommittedBatchIndex(@finalized),
// takes min with the highest verified batch index recorded by tagAdvancer,
// resolves the corresponding L2 last-block, and forwards to
// tagAdvancer.advanceFinalized.
//
// Cheap relative to derivation main loop: one L1 header + one contract call
// per tick (default 30s).
type finalizer struct {
	ctx      context.Context
	interval time.Duration
	logger   tmlog.Logger

	l1Client    *ethclient.Client
	l2Client    *types.RetryableClient
	rollup      *bindings.Rollup
	tagAdvancer *tagAdvancer

	stopped chan struct{}
}

func newFinalizer(
	ctx context.Context,
	interval time.Duration,
	l1Client *ethclient.Client,
	l2Client *types.RetryableClient,
	rollup *bindings.Rollup,
	tagAdv *tagAdvancer,
	logger tmlog.Logger,
) *finalizer {
	return &finalizer{
		ctx:         ctx,
		interval:    interval,
		l1Client:    l1Client,
		l2Client:    l2Client,
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
	// interval after startup; matches blocktag's `initialize()` behaviour.
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

	// 4. Resolve candidate batch's lastL2Block, then fetch the L2 header.
	lastL2Block, err := f.lookupBatchLastL2Block(candidate)
	if err != nil {
		f.logger.Info("finalizer: lookup batch lastL2Block failed",
			"batchIndex", candidate, "err", err)
		return
	}
	// Defensive: a zero BlockNumber means the contract slot is uninitialised
	// (BatchDataStore returned the zero value). Advancing finalized to genesis
	// would pass the monotonicity check on first call and produce a confusing
	// "finalized at block 0" tag -- skip and retry on next tick.
	if lastL2Block == 0 {
		f.logger.Info("finalizer: batch has zero lastL2Block; skipping",
			"batchIndex", candidate)
		return
	}
	header, err := f.l2Client.HeaderByNumber(f.ctx, big.NewInt(int64(lastL2Block)))
	if err != nil {
		f.logger.Info("finalizer: read L2 header failed",
			"batchIndex", candidate, "l2Block", lastL2Block, "err", err)
		return
	}
	if header == nil {
		return
	}

	f.tagAdvancer.advanceFinalized(f.ctx, candidate, header)
}

// lookupBatchLastL2Block resolves a batch index to its lastL2Block via the
// rollup contract's BatchDataStore mapping (already populated for any
// committed batch). This is the same data source blocktag.service used.
func (f *finalizer) lookupBatchLastL2Block(batchIndex uint64) (uint64, error) {
	bd, err := f.rollup.BatchDataStore(&bind.CallOpts{Context: f.ctx}, new(big.Int).SetUint64(batchIndex))
	if err != nil {
		return 0, err
	}
	if bd.BlockNumber == nil {
		return 0, fmt.Errorf("batch %d has nil BlockNumber in BatchDataStore", batchIndex)
	}
	return bd.BlockNumber.Uint64(), nil
}
