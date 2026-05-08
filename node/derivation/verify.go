package derivation

import (
	"bytes"
	"fmt"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
)

// rollbackLocalChain rolls back the local L2 chain to the specified block number.
//
// SPEC-005 §3.6 / §5: triggered on block-context mismatch or batch-root mismatch.
// After rollback the caller re-derives the offending batch from L1 calldata.
//
// SPEC-005 §4 (safety considerations) requires the rollback to be atomic w.r.t.
// the sequencer's block-production path: the sequencer must not be able to
// produce a new unsafe block while the rollback is in flight. The atomic
// ordering is:
//
//   1. Acquire the sequencer ↔ derivation mutex (P3 — sequencer_mutex.go).
//   2. Pause sequencer block production (mutex blocks RequestBlockData /
//      DeliverBlock entry points on the L2Node interface; tendermint
//      consensus layer is not modified — see tech-design §3.2.2).
//   3. Pause this derivation loop (already serialized; the caller is the loop).
//   4. Call go-ethereum's hash-matched SetHead (SPEC-005 §8 #4 blocking item).
//   5. Clear derivation cursor for the rolled-back range.
//   6. Clear L1 anchor records for the discarded segment.
//   7. Atomically persist the new safe_head metadata (head_anchor.go).
//   8. Release the mutex.
//
// Boundary: target < finalized_head → halted (SPEC-005 §3.6); enforced before
// invoking the SetHead call. target < genesis → halted.
func (d *Derivation) rollbackLocalChain(targetBlockNumber uint64) error {
	if err := d.checkRollbackBoundary(targetBlockNumber); err != nil {
		return err
	}

	d.logger.Error("L2 chain rollback not yet implemented",
		"targetBlockNumber", targetBlockNumber)

	// TODO(spec-005-rollback): implement steps 1-8 above. Blocked on:
	//   - SPEC-005 §8 #2: sequencer mutex granularity (sequencer_mutex.go).
	//   - SPEC-005 §8 #4: go-ethereum hash-matched SetHead interface (must
	//     refuse to roll back if the supplied (number, hash) does not match
	//     the local canonical chain — see tech-design §3.3).
	//   - node/types/retryable_client.go SetHead wrapper once the upstream
	//     EL method is finalised.
	return fmt.Errorf("rollback not implemented yet, target=%d", targetBlockNumber)
}

// checkRollbackBoundary enforces the SPEC-005 §3.6 boundary: rolling back
// past finalized_head is fatal, regardless of why the caller wanted to.
func (d *Derivation) checkRollbackBoundary(targetBlockNumber uint64) error {
	finalized := d.readFinalizedHead()
	if finalized != nil && targetBlockNumber < finalized.L2Number {
		// SPEC-005 §3.6 / §4.3: enter halted; no recovery short of manual
		// intervention. The caller is expected to set d.halted in response.
		return fmt.Errorf("rollback target %d below finalized_head %d — halted boundary",
			targetBlockNumber, finalized.L2Number)
	}
	return nil
}

// verifyBatchRoots verifies that the local state root and withdrawal root match the L1 batch data.
//
// SPEC-005 §3.4 / §3.2 invariant: this check is **always executed and never
// depends on blob data**. Both `batchInfo.root` (postStateRoot) and
// `batchInfo.withdrawalRoot` are extracted from L1 calldata at parse time
// (see batch_info.go); they reach this function regardless of whether the
// beacon-side blob fetch (Path A) or the local rebuild fallback (Path B,
// SPEC-005 §3.3) has succeeded. Code review must reject any change that
// makes this verification conditional on blob availability.
func (d *Derivation) verifyBatchRoots(batchInfo *BatchInfo, lastHeader *eth.Header) error {
	withdrawalRoot, err := d.L2ToL1MessagePasser.MessageRoot(&bind.CallOpts{
		BlockNumber: lastHeader.Number,
	})
	if err != nil {
		return fmt.Errorf("get withdrawal root failed: %w", err)
	}

	rootMismatch := !bytes.Equal(lastHeader.Root.Bytes(), batchInfo.root.Bytes())
	withdrawalMismatch := !bytes.Equal(withdrawalRoot[:], batchInfo.withdrawalRoot.Bytes())

	if rootMismatch || withdrawalMismatch {
		return fmt.Errorf("root mismatch: stateRoot(l1=%s, local=%s) withdrawalRoot(l1=%s, local=%s)",
			batchInfo.root.Hex(), lastHeader.Root.Hex(),
			batchInfo.withdrawalRoot.Hex(), common.BytesToHash(withdrawalRoot[:]).Hex())
	}
	return nil
}

// verifyBlockContext compares a local L2 block header against the batch block context from L1.
func (d *Derivation) verifyBlockContext(localHeader *eth.Header, blockData *BlockContext) error {
	if localHeader.Time != blockData.Timestamp {
		return fmt.Errorf("timestamp mismatch at block %d: local=%d, batch=%d",
			blockData.Number, localHeader.Time, blockData.Timestamp)
	}
	if localHeader.GasLimit != blockData.GasLimit {
		return fmt.Errorf("gasLimit mismatch at block %d: local=%d, batch=%d",
			blockData.Number, localHeader.GasLimit, blockData.GasLimit)
	}
	switch {
	case blockData.BaseFee != nil && localHeader.BaseFee != nil:
		if localHeader.BaseFee.Cmp(blockData.BaseFee) != 0 {
			return fmt.Errorf("baseFee mismatch at block %d: local=%s, batch=%s",
				blockData.Number, localHeader.BaseFee.String(), blockData.BaseFee.String())
		}
	case blockData.BaseFee == nil && localHeader.BaseFee == nil:
		// Both nil — pre-EIP-1559 or legacy batch format, OK.
	default:
		// One side has BaseFee, the other doesn't — structural inconsistency.
		return fmt.Errorf("baseFee nil mismatch at block %d: local=%v, batch=%v",
			blockData.Number, localHeader.BaseFee, blockData.BaseFee)
	}
	// Batch internal consistency check: txsNum in the block context should match the
	// actual number of transactions assembled in SafeL2Data (L1 messages + L2 txs).
	// This catches batch parsing/corruption issues, not local-vs-L1 divergence.
	// Local-vs-L1 transaction divergence is covered by state root verification
	// in verifyBatchRoots (different txs → different state root).
	if blockData.SafeL2Data != nil {
		batchTxCount := len(blockData.SafeL2Data.Transactions)
		if batchTxCount != int(blockData.txsNum) {
			return fmt.Errorf("batch internal tx count inconsistency at block %d: blockContext.txsNum=%d, safeL2Data.Transactions=%d",
				blockData.Number, blockData.txsNum, batchTxCount)
		}
	}
	return nil
}
