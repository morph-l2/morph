package derivation

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/rpc"
)

// SPEC-005 §3.2 "L1 双通道驱动":
//
// The derivation pipeline must consume two independent L1 cursors:
//   - "safe"      drives safe_head and is allowed to roll back when L1 reorgs out a batch.
//   - "finalized" drives finalized_head; it is monotonic and never rolls back.
//
// The current main loop in derivationBlock() still consumes a single
// `d.confirmations` cursor (rpc.FinalizedBlockNumber by default). The helpers
// below are intentionally not yet wired into the main loop — switching the
// main loop is gated on the SPEC-005 §8 blocking decisions (anchor-window
// depth, sequencer mutex granularity). They are exposed now so that
// downstream tasks #5 / #6 / #7 can build on them without re-establishing
// the L1 access pattern from scratch.

// fetchLatestSafeHeight returns the L1 block number of the latest "safe" head.
//
// "safe" here is the consensus-layer "safe" tag exposed via L1 RPC, not a
// confirmations-derived height. Use this to drive safe_head.
func (d *Derivation) fetchLatestSafeHeight(ctx context.Context) (uint64, error) {
	return d.fetchTaggedHeight(ctx, rpc.SafeBlockNumber, "safe")
}

// fetchLatestFinalizedHeight returns the L1 block number of the latest
// "finalized" head. Use this to drive finalized_head; the result is
// expected to be monotonic across calls.
func (d *Derivation) fetchLatestFinalizedHeight(ctx context.Context) (uint64, error) {
	return d.fetchTaggedHeight(ctx, rpc.FinalizedBlockNumber, "finalized")
}

func (d *Derivation) fetchTaggedHeight(ctx context.Context, tag rpc.BlockNumber, label string) (uint64, error) {
	header, err := d.l1Client.HeaderByNumber(ctx, big.NewInt(int64(tag)))
	if err != nil {
		return 0, fmt.Errorf("get L1 %s head: %w", label, err)
	}
	if header == nil || header.Number == nil {
		return 0, fmt.Errorf("got nil header for L1 %s head", label)
	}
	return header.Number.Uint64(), nil
}
