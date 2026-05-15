package derivation

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/node/db"
)

// SPEC-005 §4.7.6 L1 reorg detection.
//
// derivation persists the (number, hash) of every L1 block it has scanned for
// commit batch logs (via recordL1Blocks at the end of each successful poll).
// The next poll cycle calls detectReorg first; if any of the last
// reorgCheckDepth saved blocks no longer matches the live L1 hash, the
// earliest divergence height is returned and handleL1Reorg rewinds the
// derivation cursor + clears stale records.
//
// This is always-on regardless of the --derivation.confirmations setting.
// When confirmations=finalized (default), L1 finalized doesn't reorg by
// Ethereum consensus assumption, so detectReorg's fast path always returns
// (no reorg) at one L1 RPC per poll. When confirmations is configured below
// finalized (e.g. safe), detection becomes load-bearing without any code
// path divergence.
//
// L1 reorg does NOT directly trigger an L2 chain rollback in this PR. The
// L2 rollback executor (verifyBlockContext + halted state machine +
// rollbackLocalChain) is out of SPEC-005 scope (§3 non-goals). When a
// reorg replaces a committed batch with different content, derivation will
// re-derive on the next poll: if the L2 blocks come out identical (the
// common case -- same calldata, deterministic decoder), nothing further
// happens; if they differ, verifyBatchRoots fails and derivation halts at
// that batch with an error log, requiring operator intervention to re-sync.

// detectReorg checks recent L1 blocks for hash mismatches indicating a reorg.
// Returns the earliest L1 height where a mismatch was found, or nil if
// none.
//
// Optimisation: checks the newest saved block first. If it matches, there
// is no reorg (1 RPC call in the common case). Only when the newest block
// mismatches does it do a full oldest-to-newest scan to find the earliest
// divergence point.
func (d *Derivation) detectReorg(ctx context.Context) (*uint64, error) {
	latestDerivation := d.db.ReadLatestDerivationL1Height()
	if latestDerivation == nil {
		return nil, nil
	}

	checkFrom := d.startHeight
	if *latestDerivation > d.reorgCheckDepth && (*latestDerivation-d.reorgCheckDepth) > checkFrom {
		checkFrom = *latestDerivation - d.reorgCheckDepth
	}

	savedBlocks := d.db.ReadDerivationL1BlockRange(checkFrom, *latestDerivation)
	if len(savedBlocks) == 0 {
		return nil, nil
	}

	// Fast path: check the newest block first. If it matches, no reorg occurred.
	newest := savedBlocks[len(savedBlocks)-1]
	newestHeader, err := d.l1Client.HeaderByNumber(ctx, big.NewInt(int64(newest.Number)))
	if err != nil {
		return nil, fmt.Errorf("failed to get L1 header at %d: %w", newest.Number, err)
	}
	if newestHeader.Hash() == common.BytesToHash(newest.Hash[:]) {
		return nil, nil
	}

	// Slow path: reorg detected. Scan oldest-to-newest to find the earliest
	// divergence so handleL1Reorg can rewind only the affected window.
	for i := 0; i < len(savedBlocks); i++ {
		block := savedBlocks[i]
		header, err := d.l1Client.HeaderByNumber(ctx, big.NewInt(int64(block.Number)))
		if err != nil {
			return nil, fmt.Errorf("failed to get L1 header at %d: %w", block.Number, err)
		}
		savedHash := common.BytesToHash(block.Hash[:])
		if header.Hash() != savedHash {
			d.logger.Info("L1 block hash mismatch detected",
				"height", block.Number,
				"savedHash", savedHash.Hex(),
				"currentHash", header.Hash().Hex(),
			)
			return &block.Number, nil
		}
	}
	return nil, nil
}

// handleL1Reorg responds to a reorg detected at the given L1 height. It is a
// thin wrapper around rewindAndReset chosen so the call site reads as the
// reorg-handling phase of derivationBlock.
func (d *Derivation) handleL1Reorg(reorgAtL1Height uint64) error {
	d.logger.Info("L1 reorg detected, cleaning DB records and restarting derivation from reorg point",
		"reorgAtL1Height", reorgAtL1Height)
	d.rewindAndReset(reorgAtL1Height)
	return nil
}

// rewindAndReset rewinds the derivation L1 cursor to (rewindToL1Height - 1),
// clears any saved L1 block hashes at or above rewindToL1Height, and resets
// the tag advancer's safe head. Used by:
//
//   - handleL1Reorg, after detectReorg finds an L1 hash divergence
//   - finalizer's canonicality check, when the local L2 client's safe block
//     hash no longer matches what tagAdvancer recorded
//
// Both situations are recovered by the same op-stack-style "reset to a known
// good parent and re-derive forward" pattern: the next derivationBlock poll
// re-fetches L1 commit batch logs from the rewound cursor, re-runs Path A or
// Path B verification, and re-populates safe via advanceSafe. Persistent
// problems surface naturally when verifyBatchRoots fails on re-derivation.
//
// L2 chain rollback is intentionally NOT performed here -- the same commit
// tx typically gets re-included with identical content, so L2 blocks remain
// valid. If they don't, derivation halts at the offending batch with an
// error log, requiring operator intervention (SPEC-005 §3 non-goal).
//
// finalized is intentionally NOT cleared -- L1 finalized is monotonic, so
// the previous finalized value remains valid.
func (d *Derivation) rewindAndReset(rewindToL1Height uint64) {
	if rewindToL1Height < d.startHeight {
		rewindToL1Height = d.startHeight
	}

	d.db.DeleteDerivationL1BlocksFrom(rewindToL1Height)

	if rewindToL1Height > 0 {
		d.db.WriteLatestDerivationL1Height(rewindToL1Height - 1)
	} else {
		d.db.WriteLatestDerivationL1Height(0)
	}

	if d.tagAdvancer != nil {
		safeMax := d.tagAdvancer.SafeMaxBatchIndex()
		if safeMax > 0 {
			d.tagAdvancer.reset(safeMax - 1)
		} else {
			d.tagAdvancer.reset(0)
		}
	}
}

// recordL1Blocks saves L1 block hashes for reorg detection, called at the
// end of a successful poll cycle. Returns an error if any header fetch
// fails -- the caller must NOT advance the derivation cursor in that case
// to avoid permanent gaps in the L1 hash record (which would defeat
// detection).
func (d *Derivation) recordL1Blocks(ctx context.Context, from, to uint64) error {
	for h := from; h <= to; h++ {
		header, err := d.l1Client.HeaderByNumber(ctx, big.NewInt(int64(h)))
		if err != nil {
			return fmt.Errorf("failed to get L1 header at %d: %w", h, err)
		}

		var hashBytes [32]byte
		copy(hashBytes[:], header.Hash().Bytes())

		d.db.WriteDerivationL1Block(&db.DerivationL1Block{
			Number: h,
			Hash:   hashBytes,
		})
	}
	return nil
}
