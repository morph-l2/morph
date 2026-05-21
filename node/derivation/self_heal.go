package derivation

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"
)

// deriveForSelfHeal implements SPEC-005 §4.3 Path B self-heal: re-derives a
// batch using the engine API V2 path (`engine_newL2BlockV2`), which — unlike
// the normal NewSafeL2Block path used by Path A — accepts blocks whose
// parent is NOT the current chain head. The EL's SetCanonical detects the
// reorg and rewrites the locally divergent unsafe blocks to match the
// sequencer payload that we just pulled from beacon.
//
// Caller invariant: rollupData must come from fetchRollupDataByTxHash
// (Path A's full BatchInfo with blockContexts populated). The local L2
// block at firstBlockNumber - 1 must already match the canonical chain;
// this holds by construction because safe head was advanced past the
// previous batch via Path B / Path A, and self-heal only reorgs blocks
// above the safe head.
//
// Returns the header at lastBlockNumber on success, mirroring
// derive() so the caller can feed it into the shared verifyBatchRoots.
//
// **Temporary dependency**: this method calls
// RetryableClient.NewL2BlockV2, which wraps go-ethereum
// PR #325 (https://github.com/morph-l2/go-ethereum/pull/325). The
// dependency in go.mod is currently bumped to that PR's HEAD commit;
// once PR #325 is merged into main and a release tag is cut, the bump
// can be reverted to the released pseudo-version with no caller change.
func (d *Derivation) deriveForSelfHeal(rollupData *BatchInfo) (*eth.Header, error) {
	if len(rollupData.blockContexts) == 0 {
		return nil, fmt.Errorf("self-heal: empty blockContexts for batch %d", rollupData.batchIndex)
	}
	firstNum := rollupData.firstBlockNumber
	if firstNum == 0 {
		return nil, fmt.Errorf("self-heal: invalid firstBlockNumber 0 for batch %d", rollupData.batchIndex)
	}

	// Anchor: parent of the batch's first block must already exist locally.
	parentHeader, err := d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(firstNum-1)))
	if err != nil {
		return nil, fmt.Errorf("self-heal: read parent header at %d: %w", firstNum-1, err)
	}
	if parentHeader == nil {
		return nil, fmt.Errorf("self-heal: parent header at %d missing", firstNum-1)
	}
	parentHash := parentHeader.Hash()

	var lastHeader *eth.Header
	for _, blockData := range rollupData.blockContexts {
		execData := safeL2DataToExecutable(blockData.SafeL2Data, parentHash)

		callCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		newL2Err := d.l2Client.NewL2BlockV2(callCtx, execData, true /* isSafe */)
		cancel()
		if newL2Err != nil {
			d.logger.Error("self-heal: NewL2BlockV2 failed",
				"batchIndex", rollupData.batchIndex,
				"blockNumber", execData.Number,
				"parent", parentHash.Hex(),
				"error", newL2Err,
			)
			return nil, fmt.Errorf("self-heal: NewL2BlockV2 at %d: %w", execData.Number, newL2Err)
		}

		// Read back to chain the next iteration's parent and to feed
		// verifyBatchRoots at the end.
		h, err := d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(execData.Number)))
		if err != nil {
			return nil, fmt.Errorf("self-heal: read header at %d after NewL2BlockV2: %w", execData.Number, err)
		}
		if h == nil {
			return nil, fmt.Errorf("self-heal: header at %d missing after NewL2BlockV2", execData.Number)
		}
		parentHash = h.Hash()
		lastHeader = h

		d.logger.Info("self-heal: block written via NewL2BlockV2",
			"batchIndex", rollupData.batchIndex,
			"blockNumber", execData.Number,
			"hash", h.Hash().Hex(),
		)
	}
	return lastHeader, nil
}

// safeL2DataToExecutable bridges the SafeL2Data shape (used by Path A's
// NewSafeL2Block) to the ExecutableL2Data shape required by NewL2BlockV2.
//
// Self-heal calls NewL2BlockV2 with isSafe=true, which makes the EL skip
// verifyBlock + ValidateState. Execution-result fields (StateRoot,
// ReceiptRoot, GasUsed, LogsBloom, WithdrawTrieRoot, NextL1MessageIndex,
// Hash) are left zero — the EL fills them itself by re-executing the
// block.
//
// Miner is intentionally left zero. SafeL2Data does not carry it (Path A's
// existing NewSafeL2Block call has no Miner either, and the EL applies
// its configured default sequencer coinbase). If devnet shows state
// divergence after self-heal due to coinbase mismatch, source the
// sequencer address from the staking / consensus layer and plumb it
// through here. TODO: revisit once go-ethereum PR #325 is merged and
// morph-reth gets the matching change.
func safeL2DataToExecutable(s *catalyst.SafeL2Data, parentHash common.Hash) *catalyst.ExecutableL2Data {
	return &catalyst.ExecutableL2Data{
		ParentHash:   parentHash,
		Number:       s.Number,
		GasLimit:     s.GasLimit,
		BaseFee:      s.BaseFee,
		Timestamp:    s.Timestamp,
		Transactions: s.Transactions,
	}
}
