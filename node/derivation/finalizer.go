package derivation

import (
	"context"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	"morph-l2/node/types"
)

// finalizer is the SPEC-005 section 4.7.4 finalized-head subcomponent. It runs
// as an in-process goroutine inside Derivation (not a standalone service):
// each tick it computes the new finalized L2 head from L1 state and the
// local safe head, then forwards to tagAdvancer.advanceFinalized.
//
// The lookup is intentionally driven by L2 block numbers (not batch
// indices) so it doesn't depend on Rollup.BatchDataStore being populated
// for arbitrarily-old batches. The contract clears
//
//	delete batchDataStore[_batchIndex - 1];
//
// on every finalize, so an older batchIndex returns zero -- but the
// LATEST committed batch index (queried at the L1 finalized block) is
// always populated, since at that block its delete has not yet happened.
// Pinning both contract calls to the L1 finalized block makes the read
// reliable, and from there the math becomes a number comparison against
// the local safe head.
//
// Cheap: 1 L1 RPC + 2 L1 contract calls + 1 L2 RPC per tick (default 30s).
// Plus 1 L2 RPC for the rare "local verified beyond L1 finalized" branch.
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

	// 2. Pin the rollup queries to the L1 finalized block. At that block,
	// `lastCommittedBatchIndex` always references a batch whose
	// `batchDataStore` slot is still populated: the on-chain GC only
	// deletes `batchIndex - 1` on each finalizeBatch call, so for any
	// batchIndex >= lastFinalizedBatchIndex@thatBlock the slot is intact
	// at that block's state. Using the same `BlockNumber: finHeader.Number`
	// for both calls is what makes the lookup reliable.
	callOpts := &bind.CallOpts{
		BlockNumber: finHeader.Number,
		Context:     f.ctx,
	}

	committedAtFin, err := f.rollup.LastCommittedBatchIndex(callOpts)
	if err != nil {
		f.logger.Info("finalizer: query LastCommittedBatchIndex@finalized failed",
			"l1Block", finHeader.Number.Uint64(), "err", err)
		return
	}
	if committedAtFin == nil || committedAtFin.Uint64() == 0 {
		// chain not yet committed any batch.
		return
	}

	bd, err := f.rollup.BatchDataStore(callOpts, committedAtFin)
	if err != nil {
		f.logger.Info("finalizer: query BatchDataStore@finalized failed",
			"l1Block", finHeader.Number.Uint64(), "batchIndex", committedAtFin.Uint64(), "err", err)
		return
	}
	if bd.BlockNumber == nil || bd.BlockNumber.Uint64() == 0 {
		// Shouldn't happen for the latest committed batch at L1 finalized
		// (see comment above). If it does, log and skip rather than risk
		// finalizing genesis.
		f.logger.Info("finalizer: BatchDataStore[committedAtFin]@finalized has zero blockNumber; skipping",
			"l1Block", finHeader.Number.Uint64(), "batchIndex", committedAtFin.Uint64())
		return
	}
	l1FinalizedLastBlock := bd.BlockNumber.Uint64()

	// 3. Read local safe head. If derivation hasn't verified anything
	// since process start, there's nothing to anchor finalized to.
	safeHash, safeNum := f.tagAdvancer.Safe()
	if safeNum == 0 {
		return
	}

	// 4. Defensive canonicality check. Re-read the L2 client's header at
	// safeNum and verify it still matches safeHash. This catches:
	//   - L2 client state divergence (rare; would surface other bugs too)
	//   - L1 reorg propagation when Confirmations < finalized (currently
	//     not the default, but is configurable; once L1 reorg detection
	//     lands and Confirmations is upgraded, this check is the first
	//     line of defense between advanceSafe and the next L1-reorg reset)
	// On mismatch we don't advance finalized AND we reset the tag
	// advancer's safe state so derivation re-verifies before we trust it
	// again.
	safeHdr, err := f.l2Client.HeaderByNumber(f.ctx, big.NewInt(int64(safeNum)))
	if err != nil {
		f.logger.Info("finalizer: read local L2 safe header failed; skipping advance",
			"safeNumber", safeNum, "err", err)
		return
	}
	if safeHdr == nil || safeHdr.Hash() != safeHash {
		actualHash := (common.Hash{}).Hex()
		if safeHdr != nil {
			actualHash = safeHdr.Hash().Hex()
		}
		f.logger.Error("finalizer: local safe head no longer canonical; skipping advance and resetting tag advancer",
			"safeNumber", safeNum,
			"expected", safeHash.Hex(),
			"actual", actualHash)
		// Reset back to one batch before the current safe; derivation will
		// re-verify and re-call advanceSafe with the now-canonical header.
		safeMaxBatch := f.tagAdvancer.SafeMaxBatchIndex()
		if safeMaxBatch > 0 {
			f.tagAdvancer.reset(safeMaxBatch - 1)
		} else {
			f.tagAdvancer.reset(0)
		}
		return
	}

	// 5. Decide which side to anchor finalized to.
	//
	// In the common case (steady-state operation with default
	// Confirmations=finalized), L1FinalizedLastBlock >= safeNum because
	// derivation only walks L1-finalized commits and verifies them
	// in-order; both sides advance together with safe trailing slightly.
	// We anchor finalized to the local safe head -- no extra L2 RPC
	// needed, and finalized exactly tracks "what the local node has
	// verified".
	//
	// The other branch (safeNum > L1FinalizedLastBlock) only fires if
	// derivation runs ahead of L1 finalized -- e.g. operator set
	// Confirmations < finalized so derivation processes batches before
	// L1 has finalized them. We then anchor finalized to
	// L1FinalizedLastBlock and pull the L2 header from the local client
	// (we know that block exists locally because L1FinalizedLastBlock <
	// safeNum and we verified up to safeNum).
	if l1FinalizedLastBlock >= safeNum {
		f.tagAdvancer.advanceFinalized(f.ctx, committedAtFin.Uint64(), safeHash, safeNum)
		return
	}

	finalizedHdr, err := f.l2Client.HeaderByNumber(f.ctx, big.NewInt(int64(l1FinalizedLastBlock)))
	if err != nil {
		f.logger.Info("finalizer: read L2 header at L1FinalizedLastBlock failed",
			"l2Block", l1FinalizedLastBlock, "err", err)
		return
	}
	if finalizedHdr == nil {
		f.logger.Info("finalizer: L2 header at L1FinalizedLastBlock missing locally; skipping",
			"l2Block", l1FinalizedLastBlock)
		return
	}

	f.tagAdvancer.advanceFinalized(f.ctx, committedAtFin.Uint64(), finalizedHdr.Hash(), l1FinalizedLastBlock)
}
