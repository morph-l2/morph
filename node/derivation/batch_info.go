package derivation

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	geth "github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/eth/catalyst"

	commonbatch "morph-l2/common/batch"
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

	// blobHashes is the ordered list of EIP-4844 blob versioned hashes
	// declared by the L1 commitBatch tx. Path B uses this to compare
	// against locally-rebuilt versioned hashes (SPEC-005 section 4).
	blobHashes []common.Hash

	// hasCalldataBlockContexts records whether the L1 commitBatch tx
	// carried BlockContexts in calldata (legacy ABI) versus relying on
	// the blob payload to encode them at the head (new ABI with
	// LastBlockNumber + NumL1Messages). This is the only correct
	// discriminator for Path B's blob payload format:
	//   - true  -> blob = TxsPayload (V1 encoding, txs only)
	//   - false -> blob = TxsPayloadV2 (V2 encoding, blockContexts || txs)
	// `batch.Version` byte is NOT a valid discriminator because the
	// sequencer's createBatchHeader sets it from
	// (isBatchUpgraded, isBatchV2Upgraded) while handleBatchSealing
	// chooses encoding from (isBatchUpgraded, V2-fits-in-cap), so
	// version=1 batches frequently carry V2-encoded blobs in the
	// V1->V2 transition window. Path A already keys off
	// `batch.BlockContexts != nil` (see ParseBatch); Path B mirrors
	// that with this flag.
	hasCalldataBlockContexts bool
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

// ParseBatchMetadataOnly populates BatchInfo using only L1 calldata --
// it does NOT touch the blob sidecar and does NOT decode any transactions.
//
// Used by Path B (SPEC-005), which verifies the batch by rebuilding the
// blob locally rather than downloading and decoding it. Fields populated:
// batchIndex, version, root, withdrawalRoot, parentTotalL1MessagePopped,
// firstBlockNumber, lastBlockNumber. blockContexts / SafeL2Data / blobs
// are intentionally left empty; callers in Path B must not call derive().
//
// blobHashes is populated separately by the caller from tx.BlobHashes().
func (bi *BatchInfo) ParseBatchMetadataOnly(batch geth.RPCRollupBatch) error {
	parentBatchHeader := commonbatch.BatchHeaderBytes(batch.ParentBatchHeader)
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
	bi.lastBlockNumber = batch.LastBlockNumber
	// New commitBatch ABI (rollupABI / commitBatchWithProof) leaves
	// batch.BlockContexts nil; legacy ABIs (beforeMoveBlockCtxABI,
	// legacyRollupABI) populate it from calldata. UnPackData reflects this
	// directly. See the field doc on BatchInfo for why version byte cannot
	// be used here.
	bi.hasCalldataBlockContexts = len(batch.BlockContexts) > 0

	// Derive firstBlockNumber from parent batch's LastBlockNumber + 1.
	// V0 -> V1 transition leaves parent LastBlockNumber unset; in that
	// case fall back to decoding the first BlockContext from calldata.
	parentVersion, err := parentBatchHeader.Version()
	if err != nil {
		return fmt.Errorf("decode parent batch header version error:%v", err)
	}
	if parentVersion == 0 {
		if len(batch.BlockContexts) < 2+60 {
			return fmt.Errorf("calldata block contexts too short for first block context: have %d, need %d", len(batch.BlockContexts), 2+60)
		}
		var firstBlock BlockContext
		if err := firstBlock.Decode(batch.BlockContexts[2 : 2+60]); err != nil {
			return fmt.Errorf("decode first block context error:%v", err)
		}
		bi.firstBlockNumber = firstBlock.Number
	} else {
		parentLast, err := parentBatchHeader.LastBlockNumber()
		if err != nil {
			return fmt.Errorf("decode parent batch header lastBlockNumber error:%v", err)
		}
		bi.firstBlockNumber = parentLast + 1
	}
	return nil
}

// ParseBatch This method is externally referenced for parsing Batch
func (bi *BatchInfo) ParseBatch(batch geth.RPCRollupBatch) error {
	if len(batch.Sidecar.Blobs) == 0 {
		return fmt.Errorf("blobs length can not be zero")
	}
	parentBatchHeader := commonbatch.BatchHeaderBytes(batch.ParentBatchHeader)
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

	// Multi-blob batches (V2+) are produced by zstd-compressing the entire
	// batch payload as a single stream and then splitting the compressed
	// bytes across N blobs in submission order. To recover the payload we
	// must concatenate all blob bodies first and decompress once; per-blob
	// decompression would fail on the second blob since it is not a
	// standalone zstd stream.
	compressed := make([]byte, 0, len(batch.Sidecar.Blobs)*commonbatch.MaxBlobBytesSize)
	for i := range batch.Sidecar.Blobs {
		blobCopy := batch.Sidecar.Blobs[i]
		blobData, err := commonbatch.RetrieveBlobBytes(&blobCopy)
		if err != nil {
			return err
		}
		compressed = append(compressed, blobData...)
	}
	batchBytes, err := zstd.DecompressBatchBytes(compressed)
	if err != nil {
		return fmt.Errorf("decompress batch bytes error:%v", err)
	}

	var blockCount uint64
	if batch.Version > 0 {
		parentVersion, err := parentBatchHeader.Version()
		if err != nil {
			return fmt.Errorf("decode batch header version error:%v", err)
		}
		if parentVersion == 0 {
			// V0 -> V1+ transition: parent header carries no LastBlockNumber,
			// so derive blockCount from the first block context embedded at
			// the start of the decompressed batch.
			if len(batchBytes) < 60 {
				return fmt.Errorf("decompressed batch too short for start block context: have %d, need 60", len(batchBytes))
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

	var rawBlockContexts hexutil.Bytes
	var txsData []byte
	if batch.BlockContexts != nil {
		// Block contexts come from calldata; the entire decompressed stream
		// is tx payload data. ABI-decoded `bytes` can be a non-nil zero/short
		// slice, so guard the 2-byte block-count prefix read explicitly.
		if len(batch.BlockContexts) < 2 {
			return fmt.Errorf("calldata block contexts too short for block count prefix: have %d, need 2", len(batch.BlockContexts))
		}
		blockCount = uint64(binary.BigEndian.Uint16(batch.BlockContexts[:2]))
		if uint64(len(batch.BlockContexts)) < 2+60*blockCount {
			return fmt.Errorf("calldata block contexts too short: have %d, need %d", len(batch.BlockContexts), 2+60*blockCount)
		}
		rawBlockContexts = batch.BlockContexts[2 : 60*blockCount+2]
		txsData = batchBytes
	} else {
		// Block contexts are at the head of the decompressed stream,
		// immediately followed by the tx payload bytes.
		bcLen := blockCount * 60
		if uint64(len(batchBytes)) < bcLen {
			return fmt.Errorf("decompressed batch too short for block contexts: have %d, need %d", len(batchBytes), bcLen)
		}
		rawBlockContexts = batchBytes[:bcLen]
		txsData = batchBytes[bcLen:]
	}

	data, err := commonbatch.DecodeTxsFromBytes(txsData)
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
		txs, err = tq.dequeue(int(block.txsNum) - int(block.l1MsgNum))
		if err != nil {
			return fmt.Errorf("decode txsPayload error:%v", err)
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
