package derivation

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/node/db"
)

// detectReorg checks recent L1 blocks for hash mismatches indicating a reorg.
// Returns the L1 height where reorg was first detected, or nil if no reorg.
//
// Optimization: checks the newest saved block first. If it matches, there is
// no reorg (1 RPC call in the common case). Only when the newest block
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

	// Slow path: reorg detected. Scan oldest-to-newest to find the earliest divergence.
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

// handleL1Reorg handles an L1 reorg detected at the given L1 height.
// It only cleans up derivation DB state and resets the derivation L1 height
// so the next derivation loop re-processes from the reorg point.
//
// L1 reorg does NOT directly trigger an L2 rollback — in most cases the same
// batch tx gets re-included in a new L1 block with identical content, so L2
// blocks remain valid. The normal derivation loop will re-fetch batches and
// run verifyBlockContext / verifyBatchRoots; only if those comparisons fail
// will an L2 rollback be triggered through rollbackLocalChain.
func (d *Derivation) handleL1Reorg(reorgAtL1Height uint64) error {
	d.logger.Info("L1 reorg detected, cleaning DB records and restarting derivation from reorg point",
		"reorgAtL1Height", reorgAtL1Height)

	d.db.DeleteDerivationL1BlocksFrom(reorgAtL1Height)

	if reorgAtL1Height > d.startHeight {
		d.db.WriteLatestDerivationL1Height(reorgAtL1Height - 1)
	} else {
		// Reorg at or before startHeight — reset so next loop starts from startHeight.
		if d.startHeight > 0 {
			d.db.WriteLatestDerivationL1Height(d.startHeight - 1)
		} else {
			d.db.WriteLatestDerivationL1Height(0)
		}
	}

	return nil
}

// recordL1Blocks saves L1 block hashes for reorg detection.
// Returns an error if any header fetch fails — the caller must not advance
// derivation height to avoid permanent gaps in L1 block hash tracking.
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
