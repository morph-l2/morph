package derivation

import (
	"bytes"
	"fmt"
	"math/big"

	node "github.com/morph-l2/node/core"
	"github.com/morph-l2/node/types"
	"github.com/scroll-tech/go-ethereum/common"
	eth "github.com/scroll-tech/go-ethereum/core/types"
	geth "github.com/scroll-tech/go-ethereum/eth"
	"github.com/scroll-tech/go-ethereum/eth/catalyst"
)

type Chunk struct {
	blockContextes []*BlockContext
	blockNum       int
	Raw            *types.Chunk
}

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
	chunks           []*Chunk
	l1BlockNumber    uint64
	txHash           common.Hash
	nonce            uint64
	lastBlockNumber  uint64
	firstBlockNumber uint64

	root                   common.Hash
	skippedL1MessageBitmap *big.Int
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
	//var rollupData BatchInfo
	bi.root = batch.PostStateRoot
	bi.skippedL1MessageBitmap = new(big.Int).SetBytes(batch.SkippedL1MessageBitmap[:])
	bi.version = uint64(batch.Version)
	var txPayload []byte
	for _, blob := range batch.Sidecar.Blobs {
		blobCopy := blob
		data, err := types.DecodeRawTxPayload(&blobCopy)
		if err != nil {
			return err
		}
		txPayload = append(txPayload, data...)
	}
	txReader := bytes.NewReader(txPayload)
	for cbIndex, chunkByte := range batch.Chunks {
		chunk := new(types.Chunk)
		if err := chunk.Decode(chunkByte); err != nil {
			return fmt.Errorf("parse chunk error:%v", err)
		}
		bi.blockNum += uint64(chunk.BlockNum())
		var ck Chunk
		var txsNum uint64
		var l1MsgNum uint64
		for i := 0; i < chunk.BlockNum(); i++ {
			var block BlockContext
			if err := block.Decode(chunk.BlockContext()[i*60 : i*60+60]); err != nil {
				return fmt.Errorf("decode chunk block context error:%v", err)
			}
			if cbIndex == 0 && i == 0 {
				bi.firstBlockNumber = block.Number
			}
			if cbIndex == len(batch.Chunks)-1 && i == chunk.BlockNum()-1 {
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

			txs, err := node.DecodeTxsPayload(txReader, int(block.txsNum)-int(block.l1MsgNum))
			if err != nil {
				return fmt.Errorf("decode txsPayload error:%v", err)
			}
			txsNum += uint64(block.txsNum)
			l1MsgNum += uint64(block.l1MsgNum)
			// l1 transactions will be inserted later in front of L2 transactions
			safeL2Data.Transactions = encodeTransactions(txs)
			block.SafeL2Data = &safeL2Data
			ck.blockContextes = append(ck.blockContextes, &block)
		}
		bi.txNum += txsNum
		ck.Raw = chunk
		bi.chunks = append(bi.chunks, &ck)
	}
	return nil
}

func encodeTransactions(txs []*eth.Transaction) [][]byte {
	var enc = make([][]byte, len(txs))
	for i, tx := range txs {
		enc[i], _ = tx.MarshalBinary()
	}
	return enc
}

func decodeTransactions(txsBytes [][]byte) []*eth.Transaction {
	var txs []*eth.Transaction
	for _, txByte := range txsBytes {
		var tx eth.Transaction
		if err := tx.UnmarshalBinary(txByte); err != nil {
			panic(err)
		}
		txs = append(txs, &tx)
	}
	return txs
}
