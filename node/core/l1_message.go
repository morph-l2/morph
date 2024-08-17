package node

import (
	"fmt"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"

	"morph-l2/node/types"
)

func (e *Executor) updateNextL1MessageIndex(l2Block *catalyst.ExecutableL2Data) {
	e.nextL1MsgIndex = l2Block.NextL1MessageIndex
	e.metrics.NextL1MessageQueueIndex.Set(float64(e.nextL1MsgIndex))

}

// validateL1Messages has the constraints
// 1. all the collected L1 messages belonged to the given L1TxHash are exist.
// 2. the collected L1 messages belonged to the given L1TxHash are sequenced correctly.
// 3. the L1 messages from the block.Transactions are sorted correctly(queueIndex increases but does not have to be continuous).
// 4. the L1 message from the block.Transactions must be one of the collected L1Messages.
// 5. all the L1 messages from the block.Transactions must precede other normal L2 transactions.
// 6. the block.NextL1MessageIndex MUST be greater the queue index of the last involved L1Message in the block.
// 7. the skipped transactions from ExecutableL2Data extracted Must be the same as the ones from Layer1.
func (e *Executor) validateL1Messages(block *catalyst.ExecutableL2Data, collectedL1TxHashes []common.Hash) error {
	nextExpectedIndex := e.nextL1MsgIndex

	// cache: queueIndex -> L1MessageTx
	cache := make(map[uint64]*eth.Transaction)

	// constraints 1 & 2
	// build the collected L1 messages
	for _, l1TxHash := range collectedL1TxHashes {
		get, err := e.l1MsgReader.GetL1Message(nextExpectedIndex, l1TxHash)
		if err != nil {
			e.logger.Error("error getting L1 message from l1MsgReader", "error", err)
			return types.ErrQueryL1Message
		}
		if get == nil { // has not been synced from L1 yet
			e.logger.Error("the collected L1 tx hash is not valid", "L1TxHash", l1TxHash.Hex(), "expected corresponding index", nextExpectedIndex)
			return types.ErrUnknownL1Message
		} else if get.L1TxHash != l1TxHash {
			e.logger.Error("unexpected l1TxHash for the expected index", "expected index", nextExpectedIndex, "expected l1TxHash", get.L1TxHash.Hex(), "actual l1TxHash", l1TxHash.Hex())
			return types.ErrIncorrectL1TxHash
		} else if get.QueueIndex != nextExpectedIndex {
			e.logger.Error("unexpected index for the given l1TxHash", "given l1TxHash", l1TxHash.Hex(), "expected index", nextExpectedIndex, "actual index", get.QueueIndex)
			return types.ErrIncorrectL1TxHash
		}
		cache[get.QueueIndex] = eth.NewTx(&get.L1MessageTx)
		nextExpectedIndex++
	}

	nextExpectedIndex = e.nextL1MsgIndex
	L1SectionOver := false
	var skipped eth.Transactions
	// check the L1 messages from block.Transactions
	for i, txBz := range block.Transactions {
		if !isL1MessageTxType(txBz) {
			L1SectionOver = true
			continue
		}
		// constraints 5
		// check that L1 messages are before L2 transactions
		if L1SectionOver {
			return types.ErrInvalidL1MessageOrder
		}

		var tx eth.Transaction
		if err := tx.UnmarshalBinary(txBz); err != nil {
			return fmt.Errorf("transaction %d is not valid: %v", i, err)
		}
		currentTxQueueIndex := tx.L1MessageQueueIndex()

		// constraints 3
		if currentTxQueueIndex < nextExpectedIndex {
			return types.ErrInvalidL1MessageOrder
		}

		// constraints 4
		l1MessageTx, ok := cache[currentTxQueueIndex]
		if !ok {
			collectedCount := len(collectedL1TxHashes)
			if collectedCount == 0 {
				e.logger.Error("found the L1Message involved in the block, but no L1Messages collected actually")
			} else {
				e.logger.Error("the included L1Message index exceeds the last collected L1Message index",
					"current index", currentTxQueueIndex,
				)
			}
			return types.ErrUnknownL1Message
		}
		expectedTxHash := l1MessageTx.Hash()
		if tx.Hash() != expectedTxHash {
			e.logger.Error("wrong L1Message content", "index", currentTxQueueIndex)
			return types.ErrUnknownL1Message
		}
		for queueIndex := nextExpectedIndex; queueIndex < currentTxQueueIndex; queueIndex++ {
			skippedTx, ok := cache[queueIndex]
			if !ok {
				e.logger.Error("lost skipped L1Message collected", "queueIndex", queueIndex)
				return types.ErrInvalidSkippedL1Message
			}
			skipped = append(skipped, skippedTx)
		}
		nextExpectedIndex = currentTxQueueIndex + 1
	}

	// constraints 6
	if block.NextL1MessageIndex < nextExpectedIndex {
		e.logger.Error("wrong NextL1MessageIndex in the block",
			"indexOfLastInvolvedL1Tx", nextExpectedIndex-1,
			"block.NextL1MessageIndex", block.NextL1MessageIndex)
		return types.ErrWrongNextL1MessageIndex
	}
	for queueIndex := nextExpectedIndex; queueIndex < block.NextL1MessageIndex; queueIndex++ {
		skippedTx, ok := cache[queueIndex]
		if !ok {
			e.logger.Error("lost skipped L1Message collected", "queueIndex", queueIndex)
			return types.ErrInvalidL1Message
		}
		skipped = append(skipped, skippedTx)
	}

	// constraints 7
	if len(skipped) != len(block.SkippedTxs) {
		e.logger.Error("found wrong number of skipped txs", "expected skippedTx num", len(skipped), "actual", len(block.SkippedTxs))
		return types.ErrInvalidSkippedL1Message
	}
	for i, skippedTx := range skipped {
		if skippedTx.Hash() != block.SkippedTxs[i].Tx.Hash() {
			e.logger.Error("found wrong skipped tx hash", "expected skippedTx hash", skippedTx.Hash().Hex(), "actual", block.SkippedTxs[i].Tx.Hash())
			return types.ErrInvalidSkippedL1Message
		}
	}
	return nil
}

func isL1MessageTxType(rlpEncoded []byte) bool {
	if len(rlpEncoded) == 0 {
		return false
	}
	return rlpEncoded[0] == eth.L1MessageTxType
}
