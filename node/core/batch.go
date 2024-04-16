package node

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/morph-l2/node/types"
	"github.com/scroll-tech/go-ethereum/common"
	eth "github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto/bls12381"
	"github.com/tendermint/tendermint/l2node"
	tmtypes "github.com/tendermint/tendermint/types"
)

type BatchingCache struct {
	parentBatchHeader *types.BatchHeaderWithBlobHashes
	prevStateRoot     common.Hash

	// accumulated batch data
	chunks                *types.Chunks
	totalL1MessagePopped  uint64
	skippedBitmap         []*big.Int
	postStateRoot         common.Hash
	withdrawRoot          common.Hash
	lastPackedBlockHeight uint64
	// caches sealedBatchHeader according to the above accumulated batch data
	sealedBatchHeader *types.BatchHeaderWithBlobHashes
	sealedSidecar     *eth.BlobTxSidecar

	currentBlockContext               []byte
	currentTxsPayload                 []byte
	currentTxs                        tmtypes.Txs
	currentL1TxsHashes                []common.Hash
	totalL1MessagePoppedAfterCurBlock uint64
	skippedBitmapAfterCurBlock        []*big.Int
	currentStateRoot                  common.Hash
	currentWithdrawRoot               common.Hash
	currentBlockBytes                 []byte
	currentTxsHash                    []byte
	currentRowConsumption             eth.RowConsumption
}

func NewBatchingCache() *BatchingCache {
	return &BatchingCache{
		chunks: types.NewChunks(),
	}
}

func (bc *BatchingCache) IsEmpty() bool {
	return bc.chunks == nil || bc.chunks.Size() == 0
}

func (bc *BatchingCache) IsCurrentEmpty() bool {
	return len(bc.currentBlockContext) == 0
}

func (bc *BatchingCache) ClearCurrent() {
	bc.currentTxsPayload = nil
	bc.currentTxs = nil
	bc.currentL1TxsHashes = nil
	bc.currentBlockContext = nil
	bc.skippedBitmapAfterCurBlock = nil
	bc.totalL1MessagePoppedAfterCurBlock = 0
	bc.currentStateRoot = common.Hash{}
	bc.currentWithdrawRoot = common.Hash{}
	bc.currentBlockBytes = nil
	bc.currentTxsHash = nil
}

// CalculateCapWithProposalBlock calculate the transaction payload size and chunks count with the proposed block.
// It queries the blocks from the last batch point to now, in order to seal a new batch by SealBatch with these blocks.
// It stores the proposed block as the `currentBlockContext`, which is used by PackCurrentBlock to pack it to batch.
// It can be called by multiple times during the same height consensus process.
func (e *Executor) CalculateCapWithProposalBlock(currentBlockBytes []byte, currentTxs tmtypes.Txs, get l2node.GetFromBatchStartFunc) (bool, int64, error) {
	e.logger.Info("CalculateCapWithProposalBlock request", "block size", len(currentBlockBytes), "txs size", len(currentTxs))
	if e.batchingCache.IsEmpty() {
		parentBatchHeaderBytes, blocks, transactions, err := get()
		if err != nil {
			return false, 0, err
		}

		parentBatchHeader := new(types.BatchHeaderWithBlobHashes)
		if len(parentBatchHeaderBytes) == 0 {
			genesisHeader, err := e.l2Client.HeaderByNumber(context.Background(), big.NewInt(0))
			if err != nil {
				return false, 0, err
			}
			genesisBatchHeader, err := GenesisBatchHeader(genesisHeader)
			if err != nil {
				return false, 0, err
			}
			parentBatchHeader = &types.BatchHeaderWithBlobHashes{
				BatchHeader: genesisBatchHeader,
			}
		} else {
			if err = parentBatchHeader.UnmarshalBinary(parentBatchHeaderBytes); err != nil {
				return false, 0, err
			}
		}

		// skipped L1 message bitmap, an array of 256-bit bitmaps
		var skippedBitmap []*big.Int
		var txsPayload []byte
		var l1TxHashes []common.Hash
		var totalL1MessagePopped = parentBatchHeader.TotalL1MessagePopped
		var lastHeightBeforeCurrentBatch uint64
		var l2TxNum int

		for i, blockBz := range blocks {
			wBlock := new(types.WrappedBlock)
			if err = wBlock.UnmarshalBinary(blockBz); err != nil {
				return false, 0, err
			}

			if i == 0 {
				lastHeightBeforeCurrentBatch = wBlock.Number - 1
			}

			totalL1MessagePoppedBefore := totalL1MessagePopped
			txsPayload, l1TxHashes, totalL1MessagePopped, skippedBitmap, l2TxNum, err = ParsingTxs(transactions[i], parentBatchHeader.TotalL1MessagePopped, totalL1MessagePoppedBefore, skippedBitmap)
			if err != nil {
				return false, 0, err
			}
			l1TxNum := int(totalL1MessagePopped - totalL1MessagePoppedBefore) // include skipped L1 messages
			e.logger.Info("fetched block", "block height", wBlock.Number, "involved transaction count", len(transactions[i]), "l2 tx num", l2TxNum, "l1 tx num", l1TxNum)
			blockContext := wBlock.BlockContextBytes(l2TxNum+l1TxNum, l1TxNum)
			e.batchingCache.chunks.Append(blockContext, txsPayload, l1TxHashes, wBlock.RowConsumption)
			e.batchingCache.totalL1MessagePopped = totalL1MessagePopped
			e.batchingCache.lastPackedBlockHeight = wBlock.Number
		}

		// make sure passed block is the next block of the last packed block
		curHeight, err := heightFromBCBytes(currentBlockBytes)
		if err != nil {
			return false, 0, err
		}
		if curHeight != e.batchingCache.lastPackedBlockHeight+1 {
			return false, 0, fmt.Errorf("wrong propose height passed. lastPackedBlockHeight: %d, passed height: %d", e.batchingCache.lastPackedBlockHeight, curHeight)
		}

		e.batchingCache.parentBatchHeader = parentBatchHeader
		e.batchingCache.skippedBitmap = skippedBitmap
		header, err := e.l2Client.HeaderByNumber(context.Background(), big.NewInt(int64(lastHeightBeforeCurrentBatch)))
		if err != nil {
			return false, 0, err
		}
		e.batchingCache.prevStateRoot = header.Root
	}

	height, err := heightFromBCBytes(currentBlockBytes)
	if err != nil {
		return false, 0, err
	}
	if height <= e.batchingCache.lastPackedBlockHeight {
		return false, 0, fmt.Errorf("wrong propose height passed. lastPackedBlockHeight: %d, passed height: %d", e.batchingCache.lastPackedBlockHeight, height)
	} else if height > e.batchingCache.lastPackedBlockHeight+1 { // skipped some blocks, cache is dirty. need rebuild the cache
		e.batchingCache = NewBatchingCache() // clean the cache, recall the function
		e.logger.Info("the proposed block height is discontinuous from the block height in the cache, start to clean the cache and recall the function",
			"proposed block height", height,
			"batchingCache.lastPackedBlockHeight", e.batchingCache.lastPackedBlockHeight)
		return e.CalculateCapWithProposalBlock(currentBlockBytes, currentTxs, get)
	}

	if err := e.setCurrentBlock(currentBlockBytes, currentTxs); err != nil {
		return false, 0, err
	}

	chunkNum := e.batchingCache.chunks.ChunkNum()
	var exceeded bool
	var extraSize int
	// if current block will be filled in a new chunk
	chunkAppended, zeroNum := e.batchingCache.chunks.IsChunksAppendedWithNewBlock(e.batchingCache.currentRowConsumption)
	if chunkAppended {
		chunkNum += 1
		// if current block takes up a new chunk,
		// then both the added zero bytes and the first 4bytes within this chunk are all need to be counted.
		extraSize += zeroNum + 4
	}
	blobSizeWithCurBlock := e.batchingCache.chunks.CurrentPayloadForBlobSize() +
		len(e.batchingCache.currentTxsPayload) + extraSize
	if blobSizeWithCurBlock > types.MaxBlobTxPayloadSize {
		exceeded = true
	}
	e.logger.Info("CalculateCapWithProposalBlock response", "blobSizeWithCurBlock", blobSizeWithCurBlock, "exceeded", exceeded)
	return exceeded, int64(chunkNum), nil
}

// SealBatch seals the accumulated blocks into a batch
// It should be called after CalculateBatchSizeWithProposalBlock which ensure the accumulated blocks is correct.
func (e *Executor) SealBatch() ([]byte, []byte, error) {
	if e.batchingCache.IsEmpty() {
		return nil, nil, errors.New("failed to seal batch. No data found in batch cache")
	}

	// compute skipped bitmap
	skippedL1MessageBitmapBytes := make([]byte, len(e.batchingCache.skippedBitmap)*32)
	for ii, num := range e.batchingCache.skippedBitmap {
		bz := num.Bytes()
		padding := 32 - len(bz)
		copy(skippedL1MessageBitmapBytes[32*ii+padding:], bz)
	}

	batchHeader := types.BatchHeader{
		Version:                0,
		BatchIndex:             e.batchingCache.parentBatchHeader.BatchIndex + 1,
		L1MessagePopped:        e.batchingCache.totalL1MessagePopped - e.batchingCache.parentBatchHeader.TotalL1MessagePopped,
		TotalL1MessagePopped:   e.batchingCache.totalL1MessagePopped,
		DataHash:               e.batchingCache.chunks.DataHash(),
		ParentBatchHash:        e.batchingCache.parentBatchHeader.BatchHash(),
		SkippedL1MessageBitmap: skippedL1MessageBitmapBytes,
	}
	txPayload := e.batchingCache.chunks.SealTxPayloadForBlob()
	sidecar, err := types.MakeBlobTxSidecarWithTxPayload(txPayload)
	if err != nil {
		return nil, nil, err
	}
	var blobHashes []common.Hash
	if sidecar != nil && len(sidecar.Blobs) > 0 {
		blobHashes = sidecar.BlobHashes()
	}
	e.batchingCache.sealedBatchHeader = types.NewBatchHeaderWithBlobHashes(batchHeader, blobHashes)
	e.batchingCache.sealedSidecar = sidecar
	batchHash := e.batchingCache.sealedBatchHeader.BatchHash()
	e.logger.Info("Sealed batch header", "batchHash", batchHash.Hex())
	e.logger.Info(fmt.Sprintf("===batchIndex: %d \n===L1MessagePopped: %d \n===TotalL1MessagePopped: %d \n===dataHash: %x \n===blockNum: %d \n===ParentBatchHash: %x \n===SkippedL1MessageBitmap: %x \n",
		batchHeader.BatchIndex,
		batchHeader.L1MessagePopped,
		batchHeader.TotalL1MessagePopped,
		batchHeader.DataHash,
		e.batchingCache.chunks.BlockNum(),
		batchHeader.ParentBatchHash,
		batchHeader.SkippedL1MessageBitmap))
	chunksBytes, _ := e.batchingCache.chunks.Encode()
	for i, chunk := range chunksBytes {
		e.logger.Info(fmt.Sprintf("===chunk%d: %x \n", i, chunk))
	}
	e.logger.Info(fmt.Sprintf("===txs: %x \n", txPayload))
	return batchHash[:], e.batchingCache.sealedBatchHeader.MarshalBinary(), nil
}

// CommitBatch commit the sealed batch. It does nothing if no batch header is sealed.
// It is supposed to be called when the current block is confirmed.
func (e *Executor) CommitBatch(currentBlockBytes []byte, currentTxs tmtypes.Txs, blsDatas []l2node.BlsData) error {
	if e.batchingCache.IsEmpty() || e.batchingCache.sealedBatchHeader == nil { // nothing to commit
		return nil
	}

	// reconstruct current block context
	// it is possible that the confirmed current block is different from the existing cached current block context
	if !bytes.Equal(currentBlockBytes, e.batchingCache.currentBlockBytes) ||
		!bytes.Equal(currentTxs.Hash(), e.batchingCache.currentTxsHash) {
		e.logger.Info("current block is changed, reconstructing current context...")
		if err := e.setCurrentBlock(currentBlockBytes, currentTxs); err != nil {
			return err
		}
	}

	chunksBytes, err := e.batchingCache.chunks.Encode()
	if err != nil {
		return err
	}

	curHeight, err := heightFromBCBytes(e.batchingCache.currentBlockBytes)
	if err != nil {
		return err
	}

	var batchSigs []eth.BatchSignature
	if !e.devSequencer {
		batchSigs, err = e.ConvertBlsDatas(blsDatas, curHeight)
		if err != nil {
			return err
		}
	}

	if err = e.l2Client.CommitBatch(context.Background(), &eth.RollupBatch{
		Version:                0,
		Index:                  e.batchingCache.parentBatchHeader.BatchIndex + 1,
		Hash:                   e.batchingCache.sealedBatchHeader.BatchHash(),
		ParentBatchHeader:      e.batchingCache.parentBatchHeader.Encode(),
		Chunks:                 chunksBytes,
		SkippedL1MessageBitmap: e.batchingCache.sealedBatchHeader.SkippedL1MessageBitmap,
		PrevStateRoot:          e.batchingCache.prevStateRoot,
		PostStateRoot:          e.batchingCache.postStateRoot,
		WithdrawRoot:           e.batchingCache.withdrawRoot,
		Sidecar:                e.batchingCache.sealedSidecar,
	}, batchSigs); err != nil {
		return err
	}

	// commit sealed batch header; move current block into the next batch
	e.batchingCache.parentBatchHeader = e.batchingCache.sealedBatchHeader
	e.batchingCache.prevStateRoot = e.batchingCache.postStateRoot
	e.batchingCache.sealedBatchHeader = nil
	e.batchingCache.sealedSidecar = nil

	_, _, totalL1MessagePopped, skippedBitmap, _, err := ParsingTxs(e.batchingCache.currentTxs, e.batchingCache.totalL1MessagePopped, e.batchingCache.totalL1MessagePopped, nil)
	if err != nil {
		return err
	}
	e.batchingCache.totalL1MessagePopped = totalL1MessagePopped
	e.batchingCache.skippedBitmap = skippedBitmap
	e.batchingCache.postStateRoot = e.batchingCache.currentStateRoot
	e.batchingCache.withdrawRoot = e.batchingCache.currentWithdrawRoot
	e.batchingCache.lastPackedBlockHeight = curHeight
	e.batchingCache.chunks = types.NewChunks()
	e.batchingCache.chunks.Append(e.batchingCache.currentBlockContext, e.batchingCache.currentTxsPayload, e.batchingCache.currentL1TxsHashes, e.batchingCache.currentRowConsumption)
	e.batchingCache.ClearCurrent()

	e.logger.Info("Committed batch")
	return nil
}

func (e *Executor) AppendBlsData(height int64, batchHash []byte, data l2node.BlsData) error {
	if len(batchHash) != 32 {
		return fmt.Errorf("wrong batchHash length. expected: 32, actual: %d", len(batchHash))
	}
	blsSig, err := e.ConvertBlsData(data, uint64(height))
	if err != nil {
		return err
	}
	var hash common.Hash
	copy(hash[:], batchHash)
	return e.l2Client.AppendBlsSignature(context.Background(), hash, *blsSig)
}

// PackCurrentBlock pack the current block data in batchingCache into the batch
// It is supposed to be called when the current block is confirmed.
func (e *Executor) PackCurrentBlock(currentBlockBytes []byte, currentTxs tmtypes.Txs) error {
	// It is ok here to return nil, as `CalculateBatchSizeWithProposalBlock` will search historic blocks belongs to the batch being sealed.
	if e.batchingCache.IsCurrentEmpty() {
		return nil // nothing to pack
	}

	// reconstruct current block context
	// it is possible that the confirmed current block is different from the existing cached current block context
	if !bytes.Equal(currentBlockBytes, e.batchingCache.currentBlockBytes) ||
		!bytes.Equal(currentTxs.Hash(), e.batchingCache.currentTxsHash) {
		e.logger.Info("current block is changed, reconstructing current context...")
		if err := e.setCurrentBlock(currentBlockBytes, currentTxs); err != nil {
			return err
		}
	}

	curHeight, err := heightFromBCBytes(currentBlockBytes)
	if err != nil {
		return err
	}
	if e.batchingCache.chunks == nil {
		e.batchingCache.chunks = types.NewChunks()
	}
	e.batchingCache.chunks.Append(e.batchingCache.currentBlockContext, e.batchingCache.currentTxsPayload, e.batchingCache.currentL1TxsHashes, e.batchingCache.currentRowConsumption)
	e.batchingCache.skippedBitmap = e.batchingCache.skippedBitmapAfterCurBlock
	e.batchingCache.totalL1MessagePopped = e.batchingCache.totalL1MessagePoppedAfterCurBlock
	e.batchingCache.withdrawRoot = e.batchingCache.currentWithdrawRoot
	e.batchingCache.postStateRoot = e.batchingCache.currentStateRoot
	e.batchingCache.lastPackedBlockHeight = curHeight
	e.batchingCache.ClearCurrent()

	e.logger.Info("Packed current block into the batch")
	return nil
}

func (e *Executor) BatchHash(batchHeaderBytes []byte) ([]byte, error) {
	batchHeaderWithBlobHashes := new(types.BatchHeaderWithBlobHashes)
	if err := batchHeaderWithBlobHashes.UnmarshalBinary(batchHeaderBytes); err != nil {
		return nil, err
	}
	return batchHeaderWithBlobHashes.BatchHash().Bytes(), nil
}

func (e *Executor) setCurrentBlock(currentBlockBytes []byte, currentTxs tmtypes.Txs) error {
	currentTxsPayload, curL1TxsHashes, totalL1MessagePopped, skippedBitmap, l2TxNum, err := ParsingTxs(currentTxs, e.batchingCache.parentBatchHeader.TotalL1MessagePopped, e.batchingCache.totalL1MessagePopped, e.batchingCache.skippedBitmap)
	if err != nil {
		return err
	}
	var curBlock = new(types.WrappedBlock)
	if err = curBlock.UnmarshalBinary(currentBlockBytes); err != nil {
		return err
	}
	l1TxNum := int(totalL1MessagePopped - e.batchingCache.totalL1MessagePopped)
	currentBlockContext := curBlock.BlockContextBytes(l2TxNum+l1TxNum, l1TxNum)
	e.batchingCache.currentBlockContext = currentBlockContext
	e.batchingCache.currentTxsPayload = currentTxsPayload
	e.batchingCache.currentTxs = currentTxs
	e.batchingCache.currentL1TxsHashes = curL1TxsHashes
	e.batchingCache.totalL1MessagePoppedAfterCurBlock = totalL1MessagePopped
	e.batchingCache.skippedBitmapAfterCurBlock = skippedBitmap
	e.batchingCache.currentStateRoot = curBlock.StateRoot
	e.batchingCache.currentWithdrawRoot = curBlock.WithdrawTrieRoot
	e.batchingCache.currentBlockBytes = currentBlockBytes
	e.batchingCache.currentTxsHash = currentTxs.Hash()
	e.batchingCache.currentRowConsumption = curBlock.RowConsumption
	return nil
}

func ParsingTxs(transactions tmtypes.Txs, totalL1MessagePoppedBeforeTheBatch, totalL1MessagePoppedBefore uint64, skippedBitmapBefore []*big.Int) (txsPayload []byte, l1TxHashes []common.Hash, totalL1MessagePopped uint64, skippedBitmap []*big.Int, l2TxNum int, err error) {
	// the first queue index that belongs to this batch
	baseIndex := totalL1MessagePoppedBeforeTheBatch
	// the next queue index that we need to process
	nextIndex := totalL1MessagePoppedBefore

	skippedBitmap = make([]*big.Int, len(skippedBitmapBefore))
	for i, bm := range skippedBitmapBefore {
		skippedBitmap[i] = new(big.Int).SetBytes(bm.Bytes())
	}

	for i, txBz := range transactions {
		var tx eth.Transaction
		if err = tx.UnmarshalBinary(txBz); err != nil {
			return nil, nil, 0, nil, 0, fmt.Errorf("transaction %d is not valid: %v", i, err)
		}

		if isL1MessageTxType(txBz) {
			l1TxHashes = append(l1TxHashes, tx.Hash())

			currentIndex := tx.L1MessageQueueIndex()

			if currentIndex < nextIndex {
				return nil, nil, 0, nil, 0, fmt.Errorf("unexpected batch payload, expected queue index: %d, got: %d. transaction hash: %v", nextIndex, currentIndex, tx.Hash())
			}

			// mark skipped messages
			for skippedIndex := nextIndex; skippedIndex < currentIndex; skippedIndex++ {
				quo := int((skippedIndex - baseIndex) / 256)
				rem := int((skippedIndex - baseIndex) % 256)
				for len(skippedBitmap) <= quo {
					bitmap := big.NewInt(0)
					skippedBitmap = append(skippedBitmap, bitmap)
				}
				skippedBitmap[quo].SetBit(skippedBitmap[quo], rem, 1)
			}

			// process included message
			quo := int((currentIndex - baseIndex) / 256)
			for len(skippedBitmap) <= quo {
				bitmap := big.NewInt(0)
				skippedBitmap = append(skippedBitmap, bitmap)
			}

			nextIndex = currentIndex + 1
			continue
		}
		l2TxNum++

		var txLen [4]byte
		binary.BigEndian.PutUint32(txLen[:], uint32(len(txBz)))
		txsPayload = append(txsPayload, txLen[:]...)
		txsPayload = append(txsPayload, txBz...)
	}

	totalL1MessagePopped = nextIndex
	return
}

func DecodeTxsPayload(reader *bytes.Reader, txNum int) ([]*eth.Transaction, error) {
	var txs []*eth.Transaction
	for i := 0; i < txNum; i++ {
		var txLen uint32
		if err := binary.Read(reader, binary.BigEndian, &txLen); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		txBz := make([]byte, txLen)
		if err := binary.Read(reader, binary.BigEndian, &txBz); err != nil {
			return nil, fmt.Errorf("redad txBz error:%v", err)
		}
		var tx eth.Transaction
		if err := tx.UnmarshalBinary(txBz); err != nil {
			return nil, fmt.Errorf("transaction is not valid: %v", err)
		}
		txs = append(txs, &tx)
	}
	return txs, nil
}

func GenesisBatchHeader(genesisHeader *eth.Header) (types.BatchHeader, error) {
	wb := types.WrappedBlock{
		ParentHash:  genesisHeader.ParentHash,
		Miner:       genesisHeader.Coinbase,
		Number:      genesisHeader.Number.Uint64(),
		GasLimit:    genesisHeader.GasLimit,
		BaseFee:     genesisHeader.BaseFee,
		Timestamp:   genesisHeader.Time,
		StateRoot:   genesisHeader.Root,
		GasUsed:     genesisHeader.GasUsed,
		ReceiptRoot: genesisHeader.ReceiptHash,
	}
	blockContext := wb.BlockContextBytes(0, 0)
	chunks := types.NewChunks()
	chunks.Append(blockContext, nil, nil, nil)

	return types.BatchHeader{
		Version:              0,
		BatchIndex:           0,
		L1MessagePopped:      0,
		TotalL1MessagePopped: 0,
		DataHash:             chunks.DataHash(),
		ParentBatchHash:      common.Hash{},
	}, nil
}

func (e *Executor) ConvertBlsDatas(blsDatas []l2node.BlsData, height uint64) (ret []eth.BatchSignature, err error) {
	for _, blsData := range blsDatas {
		bs, err := e.ConvertBlsData(blsData, height)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *bs)
	}
	return
}

func (e *Executor) ConvertBlsData(blsData l2node.BlsData, height uint64) (*eth.BatchSignature, error) {
	seqKey := e.getBlsPubKeyByTmKey(blsData.Signer, &height)
	if seqKey == nil {
		return nil, fmt.Errorf("found invalid validator: %x", blsData.Signer)
	}

	var curVersion uint64
	if e.currentSequencerSet != nil {
		curVersion = e.currentSequencerSet.version
	}

	bs := eth.BatchSignature{
		Version:      curVersion,
		Signer:       seqKey.index,
		SignerPubKey: new(bls12381.G2).EncodePoint(seqKey.blsPubKey.Key),
		Signature:    blsData.Signature,
	}
	return &bs, nil
}

func heightFromBCBytes(blockBytes []byte) (uint64, error) {
	var curBlock = new(types.WrappedBlock)
	if err := curBlock.UnmarshalBinary(blockBytes); err != nil {
		return 0, err
	}
	return curBlock.Number, nil
}
