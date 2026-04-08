package derivation

import (
	"bytes"
	"fmt"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
)

// rollbackLocalChain rolls back the local L2 chain to the specified block number.
// This is only triggered when batch data comparison fails — i.e. the local L2 block
// does not match the L1 batch data (block context mismatch or state root mismatch).
// After rollback, the caller re-derives blocks using L1 batch data as source of truth.
func (d *Derivation) rollbackLocalChain(targetBlockNumber uint64) error {
	d.logger.Error("L2 chain rollback not yet implemented",
		"targetBlockNumber", targetBlockNumber)

	// TODO: Implement actual rollback via geth SetHead engine API:
	//  1. Expose SetL2Head(number uint64) in go-ethereum/eth/catalyst/l2_api.go
	//  2. Add SetHead method to go-ethereum/ethclient/authclient
	//  3. Add SetHead method to node/types/retryable_client.go
	//  4. Call d.l2Client.SetHead(d.ctx, targetBlockNumber)
	return fmt.Errorf("rollback not implemented yet, target=%d", targetBlockNumber)
}

// verifyBatchRoots verifies that the local state root and withdrawal root match the L1 batch data.
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
		// Check if should skip validation during upgrade transition
		if d.switchTime > 0 {
			beforeSwitch := lastHeader.Time < d.switchTime
			if (beforeSwitch && !d.useZktrie) || (!beforeSwitch && d.useZktrie) {
				d.logger.Info("Root validation skipped during upgrade transition",
					"originStateRootHash", batchInfo.root,
					"deriveStateRootHash", lastHeader.Root.Hex(),
					"blockTimestamp", lastHeader.Time,
					"switchTime", d.switchTime,
					"useZktrie", d.useZktrie,
				)
				return nil
			}
		}
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
