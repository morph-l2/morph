package derivation

import (
	"context"
	"fmt"
	"math/big"
)

// SPEC-005 §3.1 step 4 — advance finalized_head from the L1 "finalized" channel.
//
// The main loop drives the L1 "safe" channel; finalized_head lags behind
// safe_head by the L1 finalization delay (~13 minutes on mainnet). On every
// loop iteration we:
//
//  1. Read the latest safe_head (the most recent batch we've verified).
//  2. Fetch the latest L1 finalized height.
//  3. If safe_head.L1Number ≤ L1 finalized, the cheap path applies: every
//     batch up through safe_head is finalized, so finalized_head can ratchet
//     directly to safe_head.
//  4. Otherwise (the steady-state case where safe is ahead of finalized), walk
//     the rollup commit logs in the (currentFinalized.L1Number, L1 finalized]
//     range to find the latest finalizable batch, and ratchet finalized_head
//     to its L2 anchor.
//
// finalized_head is monotonic per spec §3.1; any regression attempt halts the
// node.
//
// This function is best-effort: any L1 RPC error is logged and the call
// returns. The next iteration retries.
func (d *Derivation) advanceFinalizedHead(ctx context.Context) {
	safe := d.readSafeHead()
	if safe == nil {
		return
	}

	finalizedL1, err := d.fetchLatestFinalizedHeight(ctx)
	if err != nil {
		d.logger.Debug("fetch L1 finalized failed during finalized_head advance", "err", err)
		return
	}

	current := d.readFinalizedHead()

	// Cheap path: safe_head's L1 anchor is already finalized.
	if safe.L1Number <= finalizedL1 {
		d.commitFinalizedHead(*safe, current)
		return
	}

	// Steady-state path: walk commit logs in the (currentFinalized, L1 finalized] window.
	var fromL1 uint64
	if current != nil {
		fromL1 = current.L1Number + 1
	} else {
		fromL1 = d.startHeight
	}
	if fromL1 > finalizedL1 {
		return
	}

	candidate, err := d.scanFinalizableBatch(ctx, fromL1, finalizedL1)
	if err != nil {
		d.logger.Debug("scan finalizable batch failed", "err", err, "from", fromL1, "to", finalizedL1)
		return
	}
	if candidate == nil {
		return
	}
	d.commitFinalizedHead(*candidate, current)
}

// scanFinalizableBatch returns the L2 anchor of the latest commit-batch event
// in the inclusive [from, to] L1 block window, or nil if no commit lands in
// the window. Decode-light: only consumes calldata, never fetches blob data.
func (d *Derivation) scanFinalizableBatch(ctx context.Context, from, to uint64) (*HeadAnchor, error) {
	logs, err := d.fetchRollupLog(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("fetch rollup logs: %w", err)
	}
	if len(logs) == 0 {
		return nil, nil
	}
	lg := logs[len(logs)-1]

	// Decode-light: pull the L2 last block out of calldata. No blob fetch.
	tx, pending, err := d.l1Client.TransactionByHash(ctx, lg.TxHash)
	if err != nil {
		return nil, fmt.Errorf("tx by hash %s: %w", lg.TxHash.Hex(), err)
	}
	if pending {
		return nil, nil
	}
	rb, err := d.UnPackData(tx.Data())
	if err != nil {
		return nil, fmt.Errorf("unpack commit batch: %w", err)
	}
	// Older calldata variants don't carry LastBlockNumber on the wire; in that
	// case skip finalized advance for this iteration. The next iteration after
	// safe_head ratchets past this L1 height will pick it up via the cheap
	// path above.
	if rb.LastBlockNumber == 0 {
		return nil, nil
	}

	l2Header, err := d.l2Client.HeaderByNumber(ctx, big.NewInt(int64(rb.LastBlockNumber)))
	if err != nil {
		return nil, fmt.Errorf("L2 header by number %d: %w", rb.LastBlockNumber, err)
	}
	if l2Header == nil {
		return nil, fmt.Errorf("nil L2 header for block %d", rb.LastBlockNumber)
	}

	anchor := HeadAnchor{
		L2Number: l2Header.Number.Uint64(),
		L1Number: lg.BlockNumber,
	}
	copy(anchor.L2Hash[:], l2Header.Hash().Bytes())
	copy(anchor.L1Hash[:], lg.BlockHash.Bytes())
	return &anchor, nil
}

// commitFinalizedHead writes a new finalized_head, enforcing the SPEC-005
// §3.1 monotonicity invariant: finalized_head must never regress. Any attempt
// to do so halts the node.
func (d *Derivation) commitFinalizedHead(candidate HeadAnchor, current *HeadAnchor) {
	if current != nil {
		if candidate.L2Number == current.L2Number {
			return
		}
		if candidate.L2Number < current.L2Number {
			d.logger.Error("CRITICAL: finalized_head regression attempt; halting derivation",
				"current", current.L2Number, "candidate", candidate.L2Number)
			d.halted = true
			d.metrics.SetHalted()
			return
		}
	}
	d.writeFinalizedHead(candidate)
	d.metrics.SetFinalizedHeadL2Number(candidate.L2Number)
	d.logger.Info("finalized_head advanced",
		"l2Number", candidate.L2Number, "l1Number", candidate.L1Number)
}
