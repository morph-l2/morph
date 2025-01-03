package node

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"morph-l2/node/types"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/tendermint/tendermint/l2node"
	tmtypes "github.com/tendermint/tendermint/types"
)

// CommitBatch commits the sealed batch. It does nothing if no batch header is sealed.
// It should be called when the current block is confirmed.
func (e *Executor) CommitBatch(currentBlockBytes []byte, currentTxs tmtypes.Txs, blsDatas []l2node.BlsData) error {
	// If no batch data is available, do nothing
	if e.batchingCache.IsEmpty() || e.batchingCache.sealedBatchHeader == nil {
		return nil
	}

	// Reconstruct current block context if needed
	if !bytes.Equal(currentBlockBytes, e.batchingCache.currentBlockBytes) || !bytes.Equal(currentTxs.Hash(), e.batchingCache.currentTxsHash) {
		e.logger.Info("Current block has changed. Reconstructing current context...")
		if err := e.setCurrentBlock(currentBlockBytes, currentTxs); err != nil {
			return fmt.Errorf("failed to set current block: %w", err)
		}
	}

	// Get current block height
	curHeight, err := types.HeightFromBlockBytes(e.batchingCache.currentBlockBytes)
	if err != nil {
		return fmt.Errorf("failed to parse current block height: %w", err)
	}

	// Convert BlsData to batch signatures (if applicable)
	var batchSigs []eth.BatchSignature
	if !e.devSequencer {
		batchSigs, err = e.ConvertBlsDatas(blsDatas)
		if err != nil {
			return fmt.Errorf("failed to convert BLS data: %w", err)
		}
	}

	// Get the sequencer set at current height - 1
	callOpts := &bind.CallOpts{BlockNumber: big.NewInt(int64(curHeight - 1))}
	sequencerSetBytes, err := e.sequencerCaller.GetSequencerSetBytes(callOpts)
	if err != nil {
		e.logger.Error("Failed to GetSequencerSetBytes", "blockHeight", curHeight-1, "error", err)
		return fmt.Errorf("failed to get sequencer set bytes: %w", err)
	}

	// Encode batch data and commit batch to L2 client
	blockContexts, err := e.batchingCache.batchData.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode block contexts: %w", err)
	}

	parentBatchIndex, _ := e.batchingCache.parentBatchHeader.BatchIndex()
	hash, _ := e.batchingCache.sealedBatchHeader.Hash()
	l1MessagePopped, _ := e.batchingCache.sealedBatchHeader.L1MessagePopped()
	// Construct the batch and commit it
	if err = e.l2Client.CommitBatch(context.Background(), &eth.RollupBatch{
		Version:                  0,
		Index:                    parentBatchIndex + 1,
		Hash:                     hash,
		ParentBatchHeader:        *e.batchingCache.parentBatchHeader,
		CurrentSequencerSetBytes: sequencerSetBytes,
		BlockContexts:            blockContexts,
		PrevStateRoot:            e.batchingCache.prevStateRoot,
		PostStateRoot:            e.batchingCache.postStateRoot,
		WithdrawRoot:             e.batchingCache.withdrawRoot,
		Sidecar:                  e.batchingCache.sealedSidecar,
		LastBlockNumber:          e.batchingCache.lastPackedBlockHeight,
		NumL1Messages:            uint16(l1MessagePopped),
	}, batchSigs); err != nil {
		return fmt.Errorf("failed to commit batch to L2 client: %w", err)
	}

	// Update batch index metric
	e.metrics.BatchIndex.Set(float64(parentBatchIndex + 1))

	// Commit the batch and reset the cache for the next batch
	e.commitSealedBatch(curHeight)

	e.logger.Info("Committed batch", "batchIndex", parentBatchIndex+1)
	return nil
}

// commitSealedBatch commits the sealed batch and resets cache for the next batch.
func (e *Executor) commitSealedBatch(curHeight uint64) {
	e.batchingCache.parentBatchHeader = e.batchingCache.sealedBatchHeader
	e.batchingCache.prevStateRoot = e.batchingCache.postStateRoot
	e.batchingCache.sealedBatchHeader = nil
	e.batchingCache.sealedSidecar = nil

	e.batchingCache.totalL1MessagePopped = e.batchingCache.totalL1MessagePoppedAfterCurBlock
	e.batchingCache.postStateRoot = e.batchingCache.currentStateRoot
	e.batchingCache.withdrawRoot = e.batchingCache.currentWithdrawRoot
	e.batchingCache.lastPackedBlockHeight = curHeight

	// Reset batch data and current context
	e.batchingCache.batchData = types.NewBatchData()
	e.batchingCache.batchData.Append(e.batchingCache.currentBlockContext, e.batchingCache.currentTxsPayload, e.batchingCache.currentL1TxsHashes)
	e.batchingCache.ClearCurrent()
}
