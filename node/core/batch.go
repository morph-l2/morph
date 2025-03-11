package node

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"morph-l2/node/types"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/bls12381"
	"github.com/tendermint/tendermint/l2node"
	tmtypes "github.com/tendermint/tendermint/types"
)

type BatchingCache struct {
	parentBatchHeader *types.BatchHeaderBytes
	prevStateRoot     common.Hash

	// accumulated batch data
	batchData            *types.BatchData
	totalL1MessagePopped uint64
	postStateRoot        common.Hash
	withdrawRoot         common.Hash

	lastPackedBlockHeight uint64
	// caches sealedBatchHeader according to the above accumulated batch data
	sealedBatchHeader *types.BatchHeaderBytes
	sealedSidecar     *eth.BlobTxSidecar

	currentBlockContext               []byte
	currentTxsPayload                 []byte
	currentTxs                        tmtypes.Txs
	currentL1TxsHashes                []common.Hash
	totalL1MessagePoppedAfterCurBlock uint64
	currentStateRoot                  common.Hash
	currentWithdrawRoot               common.Hash
	currentBlockBytes                 []byte
	currentTxsHash                    []byte
}

func NewBatchingCache() *BatchingCache {
	return &BatchingCache{
		batchData: types.NewBatchData(),
	}
}

func (bc *BatchingCache) IsEmpty() bool {
	return bc.batchData == nil || bc.batchData.IsEmpty()
}

func (bc *BatchingCache) IsCurrentEmpty() bool {
	return len(bc.currentBlockContext) == 0
}

func (bc *BatchingCache) ClearCurrent() {
	bc.currentTxsPayload = nil
	bc.currentTxs = nil
	bc.currentL1TxsHashes = nil
	bc.currentBlockContext = nil
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
func (e *Executor) CalculateCapWithProposalBlock(currentBlockBytes []byte, currentTxs tmtypes.Txs, get l2node.GetFromBatchStartFunc) (bool, error) {
	e.logger.Info("CalculateCapWithProposalBlock request", "block size", len(currentBlockBytes), "txs size", len(currentTxs))
	if e.batchingCache.IsEmpty() {
		parentBatchHeaderBytes, blocks, transactions, err := get()
		if err != nil {
			return false, err
		}

		var parentBatchHeader types.BatchHeaderBytes
		if len(parentBatchHeaderBytes) == 0 {
			genesisHeader, err := e.l2Client.HeaderByNumber(context.Background(), big.NewInt(0))
			if err != nil {
				return false, err
			}
			genesisBatchHeader, err := GenesisBatchHeader(genesisHeader)
			if err != nil {
				return false, err
			}
			parentBatchHeader = genesisBatchHeader.Bytes()
		} else {
			parentBatchHeader = parentBatchHeaderBytes
		}

		var txsPayload []byte
		var l1TxHashes []common.Hash
		var lastHeightBeforeCurrentBatch uint64
		var lastBlockStateRoot common.Hash
		var lastBlockWithdrawRoot common.Hash
		var l2TxNum int

		totalL1MessagePopped, err := parentBatchHeader.TotalL1MessagePopped()
		if err != nil {
			e.logger.Error("failed to get totalL1MessagePopped from parentBatchHeader", "error", err)
			return false, err
		}

		for i, blockBz := range blocks {
			wBlock := new(types.WrappedBlock)
			if err = wBlock.UnmarshalBinary(blockBz); err != nil {
				return false, err
			}

			if i == 0 {
				lastHeightBeforeCurrentBatch = wBlock.Number - 1
			}

			if i == len(blocks)-1 { // last block
				lastBlockStateRoot = wBlock.StateRoot
				lastBlockWithdrawRoot = wBlock.WithdrawTrieRoot
			}

			totalL1MessagePoppedBefore := totalL1MessagePopped
			txsPayload, l1TxHashes, totalL1MessagePopped, l2TxNum, err = ParsingTxs(transactions[i], totalL1MessagePoppedBefore)
			if err != nil {
				return false, err
			}
			l1TxNum := int(totalL1MessagePopped - totalL1MessagePoppedBefore)
			e.logger.Info("fetched block", "block height", wBlock.Number, "involved transaction count", len(transactions[i]), "l2 tx num", l2TxNum, "l1 tx num", l1TxNum)
			blockContext := wBlock.BlockContextBytes(l2TxNum+l1TxNum, l1TxNum)
			e.batchingCache.batchData.Append(blockContext, txsPayload, l1TxHashes)
			e.batchingCache.totalL1MessagePopped = totalL1MessagePopped
			e.batchingCache.lastPackedBlockHeight = wBlock.Number
		}

		// make sure passed block is the next block of the last packed block
		curHeight, err := types.HeightFromBlockBytes(currentBlockBytes)
		if err != nil {
			return false, err
		}
		if curHeight != e.batchingCache.lastPackedBlockHeight+1 {
			return false, fmt.Errorf("wrong propose height passed. lastPackedBlockHeight: %d, passed height: %d", e.batchingCache.lastPackedBlockHeight, curHeight)
		}

		e.batchingCache.parentBatchHeader = &parentBatchHeader
		header, err := e.l2Client.HeaderByNumber(context.Background(), big.NewInt(int64(lastHeightBeforeCurrentBatch)))
		if err != nil {
			return false, err
		}
		e.batchingCache.prevStateRoot = header.Root
		e.batchingCache.postStateRoot = lastBlockStateRoot
		e.batchingCache.withdrawRoot = lastBlockWithdrawRoot

		// initialize latest batch index
		index, _ := e.batchingCache.parentBatchHeader.BatchIndex()
		e.metrics.BatchIndex.Set(float64(index))
	}

	block, err := types.WrappedBlockFromBytes(currentBlockBytes)
	if err != nil {
		return false, err
	}
	height := block.Number
	if height <= e.batchingCache.lastPackedBlockHeight {
		return false, fmt.Errorf("wrong propose height passed. lastPackedBlockHeight: %d, passed height: %d", e.batchingCache.lastPackedBlockHeight, height)
	} else if height > e.batchingCache.lastPackedBlockHeight+1 { // skipped some blocks, cache is dirty. need rebuild the cache
		e.batchingCache = NewBatchingCache() // clean the cache, recall the function
		e.logger.Info("the proposed block height is discontinuous from the block height in the cache, start to clean the cache and recall the function",
			"proposed block height", height,
			"batchingCache.lastPackedBlockHeight", e.batchingCache.lastPackedBlockHeight)
		return e.CalculateCapWithProposalBlock(currentBlockBytes, currentTxs, get)
	}

	if err := e.setCurrentBlock(currentBlockBytes, currentTxs); err != nil {
		return false, err
	}

	return e.batchingCache.batchData.WillExceedCompressedSizeLimit(e.batchingCache.currentBlockContext, e.batchingCache.currentTxsPayload)
}

func (e *Executor) AppendBlsData(height int64, batchHash []byte, data l2node.BlsData) error {
	if len(batchHash) != 32 {
		return fmt.Errorf("wrong batchHash length. expected: 32, actual: %d", len(batchHash))
	}
	blsSig, err := e.ConvertBlsData(data)
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

	curHeight, err := types.HeightFromBlockBytes(currentBlockBytes)
	if err != nil {
		return err
	}
	if e.batchingCache.batchData == nil {
		e.batchingCache.batchData = types.NewBatchData()
	}
	e.batchingCache.batchData.Append(e.batchingCache.currentBlockContext, e.batchingCache.currentTxsPayload, e.batchingCache.currentL1TxsHashes)
	e.batchingCache.totalL1MessagePopped = e.batchingCache.totalL1MessagePoppedAfterCurBlock
	e.batchingCache.withdrawRoot = e.batchingCache.currentWithdrawRoot
	e.batchingCache.postStateRoot = e.batchingCache.currentStateRoot
	e.batchingCache.lastPackedBlockHeight = curHeight
	e.batchingCache.ClearCurrent()

	e.logger.Info("Packed current block into the batch")
	return nil
}

func (e *Executor) BatchHash(batchHeaderBytes []byte) ([]byte, error) {
	return crypto.Keccak256Hash(batchHeaderBytes).Bytes(), nil
}

func (e *Executor) setCurrentBlock(currentBlockBytes []byte, currentTxs tmtypes.Txs) error {
	currentTxsPayload, curL1TxsHashes, totalL1MessagePopped, l2TxNum, err := ParsingTxs(currentTxs, e.batchingCache.totalL1MessagePopped)
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
	e.batchingCache.currentStateRoot = curBlock.StateRoot
	e.batchingCache.currentWithdrawRoot = curBlock.WithdrawTrieRoot
	e.batchingCache.currentBlockBytes = currentBlockBytes
	e.batchingCache.currentTxsHash = currentTxs.Hash()
	return nil
}

func ParsingTxs(transactions tmtypes.Txs, totalL1MessagePoppedBefore uint64) (txsPayload []byte, l1TxHashes []common.Hash, totalL1MessagePopped uint64, l2TxNum int, err error) {
	// the next queue index that we need to process
	nextIndex := totalL1MessagePoppedBefore

	for i, txBz := range transactions {
		var tx eth.Transaction
		if err = tx.UnmarshalBinary(txBz); err != nil {
			return nil, nil, 0, 0, fmt.Errorf("transaction %d is not valid: %v", i, err)
		}

		if isL1MessageTxType(txBz) {
			l1TxHashes = append(l1TxHashes, tx.Hash())

			currentIndex := tx.L1MessageQueueIndex()

			if currentIndex < nextIndex {
				return nil, nil, 0, 0, fmt.Errorf("unexpected batch payload, expected queue index: %d, got: %d. transaction hash: %v", nextIndex, currentIndex, tx.Hash())
			}

			nextIndex = currentIndex + 1
			continue
		}
		l2TxNum++
		txsPayload = append(txsPayload, txBz...)
	}

	totalL1MessagePopped = nextIndex
	return
}

func GenesisBatchHeader(genesisHeader *eth.Header) (types.BatchHeaderV0, error) {
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
	batchData := types.NewBatchData()
	batchData.Append(blockContext, nil, nil)

	return types.BatchHeaderV0{
		BatchIndex:           0,
		L1MessagePopped:      0,
		TotalL1MessagePopped: 0,
		DataHash:             batchData.DataHash(),
		BlobVersionedHash:    types.EmptyVersionedHash,
		PostStateRoot:        genesisHeader.Root,
		ParentBatchHash:      common.Hash{},
	}, nil
}

func (e *Executor) ConvertBlsDatas(blsDatas []l2node.BlsData) (ret []eth.BatchSignature, err error) {
	for _, blsData := range blsDatas {
		bs, err := e.ConvertBlsData(blsData)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *bs)
	}
	return
}

func (e *Executor) ConvertBlsData(blsData l2node.BlsData) (*eth.BatchSignature, error) {
	var signer [32]byte
	copy(signer[:], blsData.Signer)
	val, found := e.valsByTmKey[signer]
	if !found {
		return nil, fmt.Errorf("found invalid validator: %x", blsData.Signer)
	}

	bs := eth.BatchSignature{
		Signer:       val.address,
		SignerPubKey: new(bls12381.G2).EncodePoint(val.blsPubKey.Key),
		Signature:    blsData.Signature,
	}
	return &bs, nil
}

func (e *Executor) isBatchUpgraded(blockTime uint64) bool {
	return blockTime >= e.UpgradeBatchTime
}
