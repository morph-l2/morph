package derivation

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	geth "github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/eth/catalyst"

	"morph-l2/node/types"
	"morph-l2/node/zstd"
)

const (
	// BlockContextLegacyLength is the length of a legacy block context without coinbase
	BlockContextLegacyLength = 60
)

type BlockContext struct {
	Number    uint64 `json:"number"`
	Timestamp uint64 `json:"timestamp"`
	BaseFee   *big.Int
	GasLimit  uint64
	txsNum    uint16
	l1MsgNum  uint16
	coinbase  common.Address

	SafeL2Data *catalyst.SafeL2Data
	L2TxHashes []byte
	TxHashes   []byte
}

func (b *BlockContext) Decode(bc []byte) error {
	wb := new(types.WrappedBlock)
	txsNum, l1MsgNum, err := wb.DecodeBlockContext(bc)
	if err != nil {
		return err
	}
	b.BaseFee = wb.BaseFee
	b.GasLimit = wb.GasLimit
	b.txsNum = txsNum
	b.l1MsgNum = l1MsgNum
	b.coinbase = wb.Miner
	return nil
}

func decodeCoinbase(bc []byte) (common.Address, error) {
	return types.DecodeCoinbase(bc)
}

type BatchInfo struct {
	batchIndex       uint64
	blockNum         uint64
	txNum            uint64
	version          uint64
	blockContexts    []*BlockContext
	l1BlockNumber    uint64
	txHash           common.Hash
	nonce            uint64
	lastBlockNumber  uint64
	firstBlockNumber uint64

	root                       common.Hash
	withdrawalRoot             common.Hash
	parentTotalL1MessagePopped uint64
}

func (bi *BatchInfo) FirstBlockNumber() uint64 {
	return bi.firstBlockNumber
}

func (bi *BatchInfo) LastBlockNumber() uint64 {
	return bi.lastBlockNumber
}

func (bi *BatchInfo) BlockNum() uint64 {
	return bi.blockNum
}

func (bi *BatchInfo) TxNum() uint64 {
	return bi.txNum
}

// ParseBatch This method is externally referenced for parsing Batch
func (bi *BatchInfo) ParseBatch(batch geth.RPCRollupBatch, morph204Time uint64) error {
	if len(batch.Sidecar.Blobs) == 0 {
		return fmt.Errorf("blobs length can not be zero")
	}

	// Parse parent batch header
	parentBatchHeader := types.BatchHeaderBytes(batch.ParentBatchHeader)
	parentBatchIndex, err := parentBatchHeader.BatchIndex()
	if err != nil {
		return fmt.Errorf("decode batch header index error: %v", err)
	}

	totalL1MessagePopped, err := parentBatchHeader.TotalL1MessagePopped()
	if err != nil {
		return fmt.Errorf("decode batch header totalL1MessagePopped error: %v", err)
	}

	// Initialize batch info fields
	bi.parentTotalL1MessagePopped = totalL1MessagePopped
	bi.root = batch.PostStateRoot
	bi.batchIndex = parentBatchIndex + 1
	bi.withdrawalRoot = batch.WithdrawRoot
	bi.version = uint64(batch.Version)

	tq := newTxQueue()
	var rawBlockContextsAndTxs hexutil.Bytes

	// Handle version upgrade scenario
	blobData, err := types.RetrieveBlobBytes(&batch.Sidecar.Blobs[0])
	if err != nil {
		return fmt.Errorf("retrieve blob bytes error: %v", err)
	}

	batchBytes, err := zstd.DecompressBatchBytes(blobData)
	if err != nil {
		return fmt.Errorf("decompress batch bytes error: %v", err)
	}

	// Calculate block count based on version
	var blockCount uint64
	if batch.Version > 0 {
		rawBlockContextsAndTxs = batchBytes
		// Ensure we have enough data for block context
		if len(batchBytes) < 60 {
			return fmt.Errorf("insufficient batch bytes for block context, got %d bytes", len(batchBytes))
		}

		var startBlock BlockContext
		// coinbase does not enter batch at this time
		if err := startBlock.Decode(batchBytes[:60]); err != nil {
			return fmt.Errorf("decode chunk block context error: %v", err)
		}

		blockCount = batch.LastBlockNumber - startBlock.Number + 1
	} else {
		// First 2 bytes contain the block count
		if len(batch.BlockContexts) < 2 {
			return fmt.Errorf("insufficient block contexts data: %d bytes", len(batch.BlockContexts))
		}

		blockCount = uint64(binary.BigEndian.Uint16(batch.BlockContexts[:2]))
		rawBlockContextsAndTxs = append(rawBlockContextsAndTxs, batch.BlockContexts[2:]...)
		rawBlockContextsAndTxs = append(rawBlockContextsAndTxs, batchBytes...)
	}

	var txsNum uint64
	blockContexts := make([]*BlockContext, int(blockCount))

	reader := bytes.NewReader(rawBlockContextsAndTxs)
	// Process block contexts
	for i := 0; i < int(blockCount); i++ {
		var block BlockContext
		bcBytes := make([]byte, BlockContextLegacyLength)
		_, err = reader.Read(bcBytes)
		if err != nil {
			return fmt.Errorf("read block context numberAndTimeBytes error:%s", err.Error())
		}
		if err := block.Decode(bcBytes); err != nil {
			return fmt.Errorf("decode number and timestamp error: %v", err)
		}
		var coinbase common.Address
		// handle coinbase
		if morph204Time == 0 || block.Timestamp >= morph204Time {
			coinbaseBytes := make([]byte, common.AddressLength)
			_, err = reader.Read(coinbaseBytes)
			if err != nil {
				return fmt.Errorf("read skipped block context  error:%s", err.Error())
			}

			coinbase, err = decodeCoinbase(coinbaseBytes)
			if err != nil {
				return err
			}
		}

		// Set boundary block numbers
		if i == 0 {
			bi.firstBlockNumber = block.Number
		}
		if i == int(blockCount)-1 {
			bi.lastBlockNumber = block.Number
		}

		// Setup SafeL2Data
		var safeL2Data catalyst.SafeL2Data
		safeL2Data.Number = block.Number
		safeL2Data.GasLimit = block.GasLimit
		safeL2Data.BaseFee = block.BaseFee
		safeL2Data.Timestamp = block.Timestamp
		// TODO coinbase
		fmt.Println(coinbase)

		// Handle zero BaseFee case
		if block.BaseFee != nil && block.BaseFee.Cmp(big.NewInt(0)) == 0 {
			safeL2Data.BaseFee = nil
		}

		// Validate transaction numbers
		if block.txsNum < block.l1MsgNum {
			return fmt.Errorf("txsNum must be greater than or equal to l1MsgNum, txsNum: %v, l1MsgNum: %v",
				block.txsNum, block.l1MsgNum)
		}

		block.SafeL2Data = &safeL2Data
		blockContexts[i] = &block
	}

	// Read transaction data
	txsData, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("read transaction data error: %s", err.Error())
	}

	// Decode transactions
	data, err := types.DecodeTxsFromBytes(txsData)
	if err != nil {
		return fmt.Errorf("decode transactions error: %v", err)
	}

	// Process transactions
	tq.enqueue(data)

	for i := 0; i < int(blockCount); i++ {
		// Skip if index is out of bounds
		if i >= len(blockContexts) {
			return fmt.Errorf("block context index out of bounds: %d >= %d", i, len(blockContexts))
		}

		txCount := int(blockContexts[i].txsNum) - int(blockContexts[i].l1MsgNum)
		txs, err := tq.dequeue(txCount)
		if err != nil {
			return fmt.Errorf("decode transaction payload error: %v", err)
		}

		txsNum += uint64(blockContexts[i].txsNum)
		// l1 transactions will be inserted later in front of L2 transactions
		blockContexts[i].SafeL2Data.Transactions = encodeTransactions(txs)
	}

	bi.txNum += txsNum
	bi.blockContexts = blockContexts

	return nil
}

func encodeTransactions(txs []*eth.Transaction) [][]byte {
	var enc = make([][]byte, len(txs))
	for i, tx := range txs {
		enc[i], _ = tx.MarshalBinary()
	}
	return enc
}

type txQueue struct {
	txs     eth.Transactions
	pointer int
}

func newTxQueue() *txQueue {
	var txs eth.Transactions
	return &txQueue{
		txs: txs,
	}
}

func (q *txQueue) enqueue(txs eth.Transactions) {
	q.txs = append(q.txs, txs...)
}

func (q *txQueue) dequeue(txNum int) (eth.Transactions, error) {
	maxLen := q.txs.Len() - q.pointer
	if maxLen < txNum {
		return nil, fmt.Errorf("invalid txNum,must small than %v", maxLen)
	}
	txs := q.txs[q.pointer : q.pointer+txNum]
	q.pointer += txNum
	return txs, nil
}
