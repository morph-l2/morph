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

type BlockContext struct {
	Number    uint64 `json:"number"`
	Timestamp uint64 `json:"timestamp"`
	BaseFee   *big.Int
	GasLimit  uint64
	txsNum    uint16
	l1MsgNum  uint16

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
	b.Number = wb.Number
	b.Timestamp = wb.Timestamp
	b.BaseFee = wb.BaseFee
	b.GasLimit = wb.GasLimit
	b.txsNum = txsNum
	b.l1MsgNum = l1MsgNum
	return nil
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
func (bi *BatchInfo) ParseBatch(batch geth.RPCRollupBatch) error {
	parentBatchHeader := types.BatchHeaderBytes(batch.ParentBatchHeader)
	parentBatchIndex, err := parentBatchHeader.BatchIndex()
	if err != nil {
		return fmt.Errorf("decode batch header index error:%v", err)
	}
	totalL1MessagePopped, err := parentBatchHeader.TotalL1MessagePopped()
	if err != nil {
		return fmt.Errorf("decode batch header totalL1MessagePopped error:%v", err)
	}
	bi.parentTotalL1MessagePopped = totalL1MessagePopped
	bi.root = batch.PostStateRoot
	bi.batchIndex = parentBatchIndex + 1
	bi.withdrawalRoot = batch.WithdrawRoot
	bi.version = uint64(batch.Version)
	tq := newTxQueue()
	var rawBlockContexts hexutil.Bytes
	var txsData []byte
	var blockCount uint64
	if batch.Version < 1 {
		parentVersion, err := parentBatchHeader.Version()
		if err != nil {
			return fmt.Errorf("decode batch header version error:%v", err)
		}
		if parentVersion == 0 {
			if len(batch.Sidecar.Blobs) == 0 {
				return fmt.Errorf("blobs lenght can not be zero")
			}
			blobData, err := types.RetrieveBlobBytes(&batch.Sidecar.Blobs[0])
			if err != nil {
				return err
			}
			batchBytes, err := zstd.DecompressBatchBytes(blobData)
			if err != nil {
				return fmt.Errorf("decompress batch bytes error:%v", err)
			}
			var startBlock BlockContext
			if err := startBlock.Decode(batchBytes[:60]); err != nil {
				return fmt.Errorf("decode chunk block context error:%v", err)
			}
			blockCount = batch.LastBlockNumber - startBlock.Number + 1
		} else {
			parentBatchBlock, err := parentBatchHeader.LastBlockNumber()
			if err != nil {
				return fmt.Errorf("decode batch header lastBlockNumber error:%v", err)
			}
			blockCount = batch.LastBlockNumber - parentBatchBlock
		}

	}
	// If BlockContexts is not nil, the block context should not be included in the blob.
	// Therefore, the required length must be zero.
	length := blockCount * 60
	for _, blob := range batch.Sidecar.Blobs {
		blobCopy := blob
		blobData, err := types.RetrieveBlobBytes(&blobCopy)
		if err != nil {
			return err
		}
		batchBytes, err := zstd.DecompressBatchBytes(blobData)
		if err != nil {
			return err
		}
		reader := bytes.NewReader(batchBytes)
		if batch.BlockContexts == nil {
			if len(batchBytes) < int(length) {
				rawBlockContexts = append(rawBlockContexts, batchBytes...)
				length -= uint64(len(batchBytes))
				reader.Reset(nil)
			} else {
				bcBytes := make([]byte, length)
				_, err = reader.Read(bcBytes)
				if err != nil {
					return fmt.Errorf("read block context error:%s", err.Error())
				}
				rawBlockContexts = append(rawBlockContexts, bcBytes...)
				length = 0
			}
		}
		data, err := io.ReadAll(reader)
		if err != nil {
			return fmt.Errorf("read txBytes error:%s", err.Error())
		}
		txsData = append(txsData, data...)
	}
	if batch.BlockContexts != nil {
		blockCount = uint64(binary.BigEndian.Uint16(batch.BlockContexts[:2]))
		rawBlockContexts = batch.BlockContexts[2 : 60*blockCount+2]
	}
	data, err := types.DecodeTxsFromBytes(txsData)
	if err != nil {
		return err
	}
	tq.enqueue(data)
	var txsNum uint64
	var l1MsgNum uint64
	blockContexts := make([]*BlockContext, int(blockCount))
	for i := 0; i < int(blockCount); i++ {
		var block BlockContext
		if err := block.Decode(rawBlockContexts[i*60 : i*60+60]); err != nil {
			return fmt.Errorf("decode chunk block context error:%v", err)
		}
		if i == 0 {
			bi.firstBlockNumber = block.Number
		}
		if i == int(blockCount)-1 {
			bi.lastBlockNumber = block.Number
		}
		var safeL2Data catalyst.SafeL2Data
		safeL2Data.Number = block.Number
		safeL2Data.GasLimit = block.GasLimit
		safeL2Data.BaseFee = block.BaseFee
		safeL2Data.Timestamp = block.Timestamp
		if block.BaseFee != nil && block.BaseFee.Cmp(big.NewInt(0)) == 0 {
			safeL2Data.BaseFee = nil
		}
		if block.txsNum < block.l1MsgNum {
			return fmt.Errorf("txsNum must be or equal to or greater than l1MsgNum,txsNum:%v,l1MsgNum:%v", block.txsNum, block.l1MsgNum)
		}
		var txs []*eth.Transaction
		var err error
		if len(batch.Sidecar.Blobs) != 0 {
			txs, err = tq.dequeue(int(block.txsNum) - int(block.l1MsgNum))
			if err != nil {
				return fmt.Errorf("decode txsPayload error:%v", err)
			}
		}
		txsNum += uint64(block.txsNum)
		l1MsgNum += uint64(block.l1MsgNum)
		// l1 transactions will be inserted later in front of L2 transactions
		safeL2Data.Transactions = encodeTransactions(txs)
		block.SafeL2Data = &safeL2Data

		blockContexts[i] = &block
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
