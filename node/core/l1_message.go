package node

import (
	"fmt"

	"github.com/scroll-tech/go-ethereum/eth/catalyst"

	"github.com/morph-l2/node/types"
	"github.com/scroll-tech/go-ethereum/common"
	eth "github.com/scroll-tech/go-ethereum/core/types"
)

func (e *Executor) updateNextL1MessageIndex(l2Block *catalyst.ExecutableL2Data) {
	e.nextL1MsgIndex = l2Block.NextL1MessageIndex
	e.metrics.NextL1MessageQueueIndex.Set(float64(e.nextL1MsgIndex))

}

// validateL1Messages has the constraints
// 1. all the collected L1 messages are valid(compared to the L1Message on layer1).
// 2. the collected L1 messages are sequenced correctly.
// 3. the L1 messages from the block.Transactions are sorted correctly(queueIndex increases but does not have to be continuous).
// 4. the L1 message from the block.Transactions must be one of the collected L1Messages.
// 5. all the L1 messages from the block.Transactions must precede other normal L2 transactions.
// 6. the block.NextL1MessageIndex MUST be greater the queue index of the last involved L1Message in the block.
func (e *Executor) validateL1Messages(block *catalyst.ExecutableL2Data, collectedL1Msgs []types.L1Message) error {
	nextExpectedIndex := e.nextL1MsgIndex

	// cache: queueIndex -> L1Message Tx Hash(L2)
	cache := make(map[uint64]common.Hash)

	// check the collected L1 messages
	for _, l1Msg := range collectedL1Msgs {

		// constraints 2
		if l1Msg.QueueIndex != nextExpectedIndex {
			return types.ErrInvalidL1MessageOrder
		}

		// constraints 1
		get, err := e.l1MsgReader.GetL1Message(l1Msg.QueueIndex, l1Msg.L1TxHash)
		if err != nil {
			e.logger.Error("error getting L1 message from l1MsgReader", "error", err)
			return types.ErrQueryL1Message
		}
		if get == nil { // has not been synced from L1 yet
			e.logger.Error("the collected L1 message is not valid", "index", l1Msg.QueueIndex, "L1TxHash", l1Msg.L1TxHash.Hex())
			return types.ErrUnknownL1Message
		}
		txHash := eth.NewTx(&l1Msg.L1MessageTx).Hash()
		if txHash != eth.NewTx(&get.L1MessageTx).Hash() {
			e.logger.Error("the collected L1 message is not equals to the actual L1 message", "index", l1Msg.QueueIndex, "L1TxHash", l1Msg.L1TxHash.Hex())
			return types.ErrUnknownL1Message
		}
		cache[l1Msg.QueueIndex] = txHash
		nextExpectedIndex++
	}

	nextExpectedIndex = e.nextL1MsgIndex
	L1SectionOver := false
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
		expectedTxHash, ok := cache[currentTxQueueIndex]
		if !ok {
			collectedCount := len(collectedL1Msgs)
			if collectedCount == 0 {
				e.logger.Error("found the L1Message involved in the block, but no L1Messages collected actually")
			} else {
				e.logger.Error("the included L1Message index exceeds the last collected L1Message index",
					"current index", currentTxQueueIndex,
					"max index of collected L1Messages", collectedL1Msgs[collectedCount-1].QueueIndex,
				)
			}
			return types.ErrUnknownL1Message
		}

		if tx.Hash() != expectedTxHash {
			e.logger.Error("wrong L1Message content", "index", currentTxQueueIndex)
			return types.ErrUnknownL1Message
		}

		nextExpectedIndex = currentTxQueueIndex + 1
	}

	// constraints 6
	if block.NextL1MessageIndex < nextExpectedIndex {
		e.logger.Error("wrong NextL1MessageIndex in the block",
			"index of the last involved L1 tx", nextExpectedIndex-1,
			"block.NextL1MessageIndex", block.NextL1MessageIndex)
		return types.ErrWrongNextL1MessageIndex
	}

	return nil
}

func isL1MessageTxType(rlpEncoded []byte) bool {
	if len(rlpEncoded) == 0 {
		return false
	}
	return rlpEncoded[0] == eth.L1MessageTxType
}
