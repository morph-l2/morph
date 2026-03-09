package batchprocessor

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	"morph-l2/node/derivation"
	"morph-l2/node/types"
)

// BatchProcessor sequentially processes committed batches from L1, verifies
// each via BatchVerifier, and updates safe/finalized block tags on L2.
//
// Unlike BlockTagService (which it replaces), it walks every batch in order
// so no intermediate batch is skipped.
type BatchProcessor struct {
	ctx    context.Context
	cancel context.CancelFunc

	l1Client *ethclient.Client
	l2Client *types.RetryableClient
	rollup   *bindings.Rollup

	batchVerifier *derivation.BatchVerifier

	rollupAddress     common.Address
	safeConfirmations uint64
	pollInterval      time.Duration

	// Sequential cursors: the last batch index we verified for each tag.
	// On each tick we attempt to advance from cursor+1 up to the on-chain head.
	lastSafeBatchIndex      uint64
	lastFinalizedBatchIndex uint64

	// L1 block of the last CommitBatch event we found. Used as the search
	// start for subsequent FilterLogs calls so we never miss events during
	// catch-up (replaces the fixed 1000-block window).
	lastCommitL1Block uint64

	// Tag state
	safeL2BlockHash           common.Hash
	finalizedL2BlockHash      common.Hash
	lastNotifiedSafeHash      common.Hash
	lastNotifiedFinalizedHash common.Hash

	logger tmlog.Logger
	stop   chan struct{}
}

func NewBatchProcessor(
	ctx context.Context,
	l2Client *types.RetryableClient,
	config *Config,
	bv *derivation.BatchVerifier,
	logger tmlog.Logger,
) (*BatchProcessor, error) {
	if config.L1Addr == "" {
		return nil, fmt.Errorf("L1 RPC address is required")
	}

	l1Client, err := ethclient.Dial(config.L1Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to L1: %w", err)
	}

	rollup, err := bindings.NewRollup(config.RollupAddress, l1Client)
	if err != nil {
		l1Client.Close()
		return nil, fmt.Errorf("failed to create rollup binding: %w", err)
	}

	ctx, cancel := context.WithCancel(ctx)

	return &BatchProcessor{
		ctx:               ctx,
		cancel:            cancel,
		l1Client:          l1Client,
		l2Client:          l2Client,
		rollup:            rollup,
		batchVerifier:     bv,
		rollupAddress:     config.RollupAddress,
		safeConfirmations: config.SafeConfirmations,
		pollInterval:      config.PollInterval,
		logger:            logger.With("module", "batchprocessor"),
		stop:              make(chan struct{}),
	}, nil
}

func (bp *BatchProcessor) Start() error {
	if err := bp.initCursors(); err != nil {
		bp.logger.Error("failed to init cursors, starting from 0", "error", err)
	}

	go bp.loop()
	bp.logger.Info("BatchProcessor started",
		"safeCursor", bp.lastSafeBatchIndex,
		"finalizedCursor", bp.lastFinalizedBatchIndex,
		"pollInterval", bp.pollInterval)
	return nil
}

func (bp *BatchProcessor) Stop() {
	bp.cancel()
	<-bp.stop

	bp.l1Client.Close()
	if bp.batchVerifier != nil {
		bp.batchVerifier.Close()
	}
	bp.logger.Info("BatchProcessor stopped")
}

// initCursors sets both cursors to the on-chain LastFinalizedBatchIndex so we
// skip already finalized (and therefore already validated) history.
// It also resolves the corresponding L2 block hashes so that geth receives
// correct safe/finalized tags immediately after restart.
func (bp *BatchProcessor) initCursors() error {
	lastFinalized, err := bp.rollup.LastFinalizedBatchIndex(nil)
	if err != nil || lastFinalized == nil {
		return fmt.Errorf("query LastFinalizedBatchIndex: %w", err)
	}
	idx := lastFinalized.Uint64()
	if idx == 0 {
		return nil
	}

	bp.lastSafeBatchIndex = idx
	bp.lastFinalizedBatchIndex = idx

	// Resolve the L2 block hash for the cursor batch so that notifyGeth can
	// send correct tags right away instead of waiting for the next new batch.
	batchData, err := bp.rollup.BatchDataStore(nil, new(big.Int).SetUint64(idx))
	if err != nil {
		bp.logger.Error("initCursors: failed to query BatchDataStore, tags will be empty until next batch", "error", err)
		return nil
	}
	lastBlock := batchData.BlockNumber.Uint64()
	header, err := bp.l2Client.HeaderByNumber(bp.ctx, new(big.Int).SetUint64(lastBlock))
	if err != nil {
		bp.logger.Error("initCursors: failed to get L2 header, tags will be empty until next batch",
			"l2Block", lastBlock, "error", err)
		return nil
	}
	blockHash := header.Hash()
	bp.safeL2BlockHash = blockHash
	bp.finalizedL2BlockHash = blockHash

	bp.logger.Info("cursors initialized from on-chain finalized index",
		"batchIndex", idx, "l2Block", lastBlock, "blockHash", blockHash.Hex())
	return nil
}

func (bp *BatchProcessor) loop() {
	defer close(bp.stop)

	ticker := time.NewTicker(bp.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-bp.ctx.Done():
			return
		case <-ticker.C:
			if err := bp.processTick(); err != nil {
				bp.logger.Error("tick failed", "error", err)
			}
		}
	}
}

func (bp *BatchProcessor) processTick() error {
	l2Head, err := bp.l2Client.BlockNumber(bp.ctx)
	if err != nil {
		return fmt.Errorf("get L2 head: %w", err)
	}

	// Cache L1 head once per tick to avoid redundant RPC calls when processing
	// multiple batches in a single tick (catch-up scenario).
	currentL1, err := bp.l1Client.BlockNumber(bp.ctx)
	if err != nil {
		return fmt.Errorf("get L1 head: %w", err)
	}

	// --- safe ---
	safeL1Head := bp.calcSafeL1Head(currentL1)
	if safeL1Head > 0 {
		safeOnChainHead, err := bp.getLastCommittedBatchAtBlock(rpc.BlockNumber(safeL1Head))
		if err != nil {
			bp.logger.Error("get safe committed head failed", "error", err)
		} else {
			bp.advanceSafe(l2Head, safeOnChainHead, currentL1)
		}
	}

	// --- finalized ---
	finalizedOnChainHead, err := bp.getLastCommittedBatchAtBlock(rpc.FinalizedBlockNumber)
	if err != nil {
		bp.logger.Error("get finalized committed head failed", "error", err)
	} else {
		bp.advanceFinalized(l2Head, finalizedOnChainHead, currentL1)
	}

	if err := bp.notifyGeth(); err != nil {
		bp.logger.Error("notify geth failed", "error", err)
	}

	bp.logger.Debug("tick done",
		"l2Head", l2Head,
		"safeCursor", bp.lastSafeBatchIndex,
		"finalizedCursor", bp.lastFinalizedBatchIndex,
		"safeHash", bp.safeL2BlockHash.Hex(),
		"finalizedHash", bp.finalizedL2BlockHash.Hex())

	return nil
}

// advanceSafe tries to move the safe cursor forward one batch at a time.
func (bp *BatchProcessor) advanceSafe(l2Head, onChainHead, currentL1 uint64) {
	for idx := bp.lastSafeBatchIndex + 1; idx <= onChainHead; idx++ {
		lastBlock, hash, err := bp.processOneBatch(idx, l2Head, currentL1)
		if err != nil {
			bp.logger.Error("safe batch processing failed", "batchIndex", idx, "error", err)
			return
		}
		if lastBlock == 0 {
			return // batch not yet completed on L2
		}
		bp.lastSafeBatchIndex = idx
		bp.safeL2BlockHash = hash
		bp.logger.Info("safe cursor advanced", "batchIndex", idx, "l2Block", lastBlock)
	}
}

// advanceFinalized tries to move the finalized cursor forward. It can never
// exceed the safe cursor because finalized <= safe.
func (bp *BatchProcessor) advanceFinalized(l2Head, onChainHead, currentL1 uint64) {
	limit := onChainHead
	if bp.lastSafeBatchIndex < limit {
		limit = bp.lastSafeBatchIndex
	}
	for idx := bp.lastFinalizedBatchIndex + 1; idx <= limit; idx++ {
		lastBlock, hash, err := bp.processOneBatch(idx, l2Head, currentL1)
		if err != nil {
			bp.logger.Error("finalized batch processing failed", "batchIndex", idx, "error", err)
			return
		}
		if lastBlock == 0 {
			return
		}
		bp.lastFinalizedBatchIndex = idx
		bp.finalizedL2BlockHash = hash

		// finalized implies safe
		if bp.safeL2BlockHash == (common.Hash{}) {
			bp.safeL2BlockHash = hash
		}
		bp.logger.Info("finalized cursor advanced", "batchIndex", idx, "l2Block", lastBlock)
	}
}

// processOneBatch fetches, optionally verifies, and returns the last L2 block
// number + hash for a given batch. Returns (0, empty, nil) if the batch's last
// L2 block is not yet available locally (node still syncing).
func (bp *BatchProcessor) processOneBatch(batchIndex, l2Head, currentL1 uint64) (uint64, common.Hash, error) {
	batchData, err := bp.rollup.BatchDataStore(nil, new(big.Int).SetUint64(batchIndex))
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("query BatchDataStore(%d): %w", batchIndex, err)
	}

	lastBlock := batchData.BlockNumber.Uint64()
	if lastBlock > l2Head {
		return 0, common.Hash{}, nil // not synced yet
	}

	if bp.batchVerifier != nil {
		if err := bp.verifyBatch(batchIndex, currentL1); err != nil {
			// TODO: decide on verification failure behavior
			return 0, common.Hash{}, fmt.Errorf("batch %d verification failed: %w", batchIndex, err)
		}
	}

	header, err := bp.l2Client.HeaderByNumber(bp.ctx, new(big.Int).SetUint64(lastBlock))
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("get L2 header %d: %w", lastBlock, err)
	}

	return lastBlock, header.Hash(), nil
}

// verifyBatch locates the CommitBatch L1 tx and runs full verification.
func (bp *BatchProcessor) verifyBatch(batchIndex, currentL1 uint64) error {
	txHash, err := bp.fetchCommitBatchTxHash(batchIndex, currentL1)
	if err != nil {
		return fmt.Errorf("fetch CommitBatch tx: %w", err)
	}

	roots, err := bp.batchVerifier.FetchBatchRoots(bp.ctx, txHash, batchIndex)
	if err != nil {
		return fmt.Errorf("FetchBatchRoots: %w", err)
	}

	var batchInfo *derivation.BatchInfo
	batchInfo, err = bp.batchVerifier.FetchBatchData(bp.ctx, txHash)
	if err != nil {
		bp.logger.Error("FetchBatchData failed, skipping tx-level verification", "error", err)
		batchInfo = nil
	}

	if err := bp.batchVerifier.VerifyBatch(bp.ctx, bp.l2Client, roots, batchInfo); err != nil {
		return fmt.Errorf("VerifyBatch: %w", err)
	}
	return nil
}

// fetchCommitBatchTxHash finds the CommitBatch event on L1 for a given batch
// index. It searches forward from lastCommitL1Block (the L1 block of the
// previous CommitBatch event) so the window naturally covers catch-up scenarios
// instead of being limited to a fixed range.
func (bp *BatchProcessor) fetchCommitBatchTxHash(batchIndex, currentL1 uint64) (common.Hash, error) {
	fromBlock := bp.lastCommitL1Block

	batchIndexBig := new(big.Int).SetUint64(batchIndex)
	logs, err := bp.l1Client.FilterLogs(bp.ctx, ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(currentL1),
		Addresses: []common.Address{bp.rollupAddress},
		Topics: [][]common.Hash{
			{derivation.RollupEventTopicHash},
			{common.BigToHash(batchIndexBig)},
		},
	})
	if err != nil {
		return common.Hash{}, fmt.Errorf("FilterLogs: %w", err)
	}
	if len(logs) == 0 {
		return common.Hash{}, fmt.Errorf("no CommitBatch event found for batch %d in L1 blocks [%d, %d]", batchIndex, fromBlock, currentL1)
	}

	bp.lastCommitL1Block = logs[0].BlockNumber
	return logs[0].TxHash, nil
}

func (bp *BatchProcessor) calcSafeL1Head(currentL1 uint64) uint64 {
	if currentL1 <= bp.safeConfirmations {
		return 0
	}
	return currentL1 - bp.safeConfirmations
}

func (bp *BatchProcessor) getLastCommittedBatchAtBlock(l1BlockTag rpc.BlockNumber) (uint64, error) {
	var blockNum *big.Int
	if l1BlockTag == rpc.FinalizedBlockNumber {
		blockNum = big.NewInt(int64(rpc.FinalizedBlockNumber))
	} else if l1BlockTag >= 0 {
		blockNum = big.NewInt(int64(l1BlockTag))
	}

	lastCommitted, err := bp.rollup.LastCommittedBatchIndex(&bind.CallOpts{
		BlockNumber: blockNum,
		Context:     bp.ctx,
	})
	if err != nil {
		return 0, err
	}
	return lastCommitted.Uint64(), nil
}

func (bp *BatchProcessor) notifyGeth() error {
	safe := bp.safeL2BlockHash
	finalized := bp.finalizedL2BlockHash

	if safe == bp.lastNotifiedSafeHash && finalized == bp.lastNotifiedFinalizedHash {
		return nil
	}
	if safe == (common.Hash{}) && finalized == (common.Hash{}) {
		return nil
	}

	if err := bp.l2Client.SetBlockTags(bp.ctx, safe, finalized); err != nil {
		return err
	}

	bp.lastNotifiedSafeHash = safe
	bp.lastNotifiedFinalizedHash = finalized
	return nil
}
