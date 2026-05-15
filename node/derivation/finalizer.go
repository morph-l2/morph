package derivation

import (
	"math/big"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/rpc"
)

// SPEC-005 §4.7.4 finalized-head tick.
//
// Originally a separate goroutine (with its own ticker / stopped channel),
// folded into the derivation main loop because the only justification for
// a separate goroutine was "L1 RPC could be slow" -- which was already true
// for the main loop's eth_getLogs / eth_getTransactionByHash calls, so the
// extra goroutine bought nothing and introduced cross-goroutine state writes
// (cursor / tagAdvancer) that complicated the canonicality recovery path.
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
// Cost: 1 L1 RPC + 2 L1 contract calls + 1 L2 RPC per main-loop poll.
// Plus 1 L2 RPC for the rare "local verified beyond L1 finalized" branch.
func (d *Derivation) finalizerTick() {
	// 1. Resolve the L1 finalized header.
	finHeader, err := d.l1Client.HeaderByNumber(d.ctx, big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		d.logger.Info("finalizer: read L1 finalized header failed", "err", err)
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
		Context:     d.ctx,
	}

	committedAtFin, err := d.rollup.LastCommittedBatchIndex(callOpts)
	if err != nil {
		d.logger.Info("finalizer: query LastCommittedBatchIndex@finalized failed",
			"l1Block", finHeader.Number.Uint64(), "err", err)
		return
	}
	if committedAtFin == nil || committedAtFin.Uint64() == 0 {
		// chain not yet committed any batch.
		return
	}

	bd, err := d.rollup.BatchDataStore(callOpts, committedAtFin)
	if err != nil {
		d.logger.Info("finalizer: query BatchDataStore@finalized failed",
			"l1Block", finHeader.Number.Uint64(), "batchIndex", committedAtFin.Uint64(), "err", err)
		return
	}
	if bd.BlockNumber == nil || bd.BlockNumber.Uint64() == 0 {
		// Shouldn't happen for the latest committed batch at L1 finalized
		// (see comment above). If it does, log and skip rather than risk
		// finalizing genesis.
		d.logger.Info("finalizer: BatchDataStore[committedAtFin]@finalized has zero blockNumber; skipping",
			"l1Block", finHeader.Number.Uint64(), "batchIndex", committedAtFin.Uint64())
		return
	}
	l1FinalizedLastBlock := bd.BlockNumber.Uint64()

	// 3. Read local safe head. If derivation hasn't verified anything
	// since process start, there's nothing to anchor finalized to.
	safeHash, safeNum := d.tagAdvancer.Safe()
	if safeNum == 0 {
		return
	}

	// 4. Defensive canonicality check. Re-read the L2 client's header at
	// safeNum and verify it still matches safeHash. On mismatch we rewind
	// the derivation cursor (op-stack-style "reset to a known good parent
	// and re-derive forward" -- shared with the L1 reorg recovery path
	// via rewindAndReset). This catches:
	//   - L2 client state divergence (rare; would surface other bugs too)
	//   - L1 reorg propagation that detectReorg missed (race or bug in the
	//     reorg detection window)
	safeHdr, err := d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(safeNum)))
	if err != nil {
		d.logger.Info("finalizer: read local L2 safe header failed; skipping advance",
			"safeNumber", safeNum, "err", err)
		return
	}
	if safeHdr == nil || safeHdr.Hash() != safeHash {
		actualHash := (common.Hash{}).Hex()
		if safeHdr != nil {
			actualHash = safeHdr.Hash().Hex()
		}
		// Rewind by reorgCheckDepth from the current cursor so the next
		// derivationBlock poll re-fetches recent batches and re-verifies.
		// Persistent breakage will resurface as verifyBatchRoots failure on
		// re-derivation; transient state-client weirdness self-heals.
		var rewindTo uint64
		if cur := d.db.ReadLatestDerivationL1Height(); cur != nil {
			if *cur > d.reorgCheckDepth {
				rewindTo = *cur - d.reorgCheckDepth
			} else {
				rewindTo = d.startHeight
			}
		} else {
			rewindTo = d.startHeight
		}
		d.logger.Error("finalizer: local safe head no longer canonical; rewinding cursor and resetting tag advancer",
			"safeNumber", safeNum,
			"expected", safeHash.Hex(),
			"actual", actualHash,
			"rewindTo", rewindTo)
		d.rewindAndReset(rewindTo)
		return
	}

	// 5. Decide which side to anchor finalized to.
	//
	// In the common case (steady-state operation), L1FinalizedLastBlock >=
	// safeNum because derivation only walks L1-finalized commits and
	// verifies them in-order; both sides advance together with safe
	// trailing slightly. We anchor finalized to the local safe head -- no
	// extra L2 RPC needed, and finalized exactly tracks "what the local
	// node has verified".
	//
	// The other branch (safeNum > L1FinalizedLastBlock) only fires if
	// derivation runs ahead of L1 finalized -- e.g. operator set
	// Confirmations < finalized so derivation processes batches before
	// L1 has finalized them. We then anchor finalized to
	// L1FinalizedLastBlock and pull the L2 header from the local client
	// (we know that block exists locally because L1FinalizedLastBlock <
	// safeNum and we verified up to safeNum).
	if l1FinalizedLastBlock >= safeNum {
		d.tagAdvancer.advanceFinalized(d.ctx, committedAtFin.Uint64(), safeHash, safeNum)
		return
	}

	finalizedHdr, err := d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(l1FinalizedLastBlock)))
	if err != nil {
		d.logger.Info("finalizer: read L2 header at L1FinalizedLastBlock failed",
			"l2Block", l1FinalizedLastBlock, "err", err)
		return
	}
	if finalizedHdr == nil {
		d.logger.Info("finalizer: L2 header at L1FinalizedLastBlock missing locally; skipping",
			"l2Block", l1FinalizedLastBlock)
		return
	}

	d.tagAdvancer.advanceFinalized(d.ctx, committedAtFin.Uint64(), finalizedHdr.Hash(), l1FinalizedLastBlock)
}
