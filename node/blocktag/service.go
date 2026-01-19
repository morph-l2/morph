package blocktag

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	"morph-l2/node/types"
)

// BlockTagType represents the type of block tag (safe or finalized)
type BlockTagType int

const (
	TagTypeSafe BlockTagType = iota
	TagTypeFinalized
)

// BlockTagService is responsible for tracking and updating safe/finalized block tags
// based on L1 batch commit tx status.
//
// Key logic:
// - Safe: batch tx is committed to L1 with N block confirmations (configurable)
// - Finalized: batch tx's L1 block is finalized (using L1 finalized block tag)
type BlockTagService struct {
	ctx    context.Context
	cancel context.CancelFunc

	// Current safe and finalized L2 block hashes
	safeL2BlockHash      common.Hash
	finalizedL2BlockHash common.Hash
	// Last notified hashes (to avoid redundant RPC calls)
	lastNotifiedSafeHash      common.Hash
	lastNotifiedFinalizedHash common.Hash

	// Cached batch index for optimization (avoid full binary search each time)
	// Separate caches for safe and finalized since they have different maxBatchIndex
	lastKnownSafeBatchIndex      uint64
	lastKnownFinalizedBatchIndex uint64

	// Clients
	l1Client *ethclient.Client
	l2Client *types.RetryableClient
	rollup   *bindings.Rollup

	// Configuration
	rollupAddress     common.Address
	safeConfirmations uint64 // Number of L1 blocks to wait before considering a batch as safe
	pollInterval      time.Duration

	logger tmlog.Logger
	stop   chan struct{}
}

// NewBlockTagService creates a new BlockTagService
func NewBlockTagService(
	ctx context.Context,
	l2Client *types.RetryableClient,
	config *Config,
	logger tmlog.Logger,
) (*BlockTagService, error) {
	if config.L1Addr == "" {
		return nil, fmt.Errorf("L1 RPC address is required")
	}
	if config.RollupAddress == (common.Address{}) {
		return nil, fmt.Errorf("Rollup contract address is required")
	}

	l1Client, err := ethclient.Dial(config.L1Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to L1: %w", err)
	}

	rollup, err := bindings.NewRollup(config.RollupAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create rollup binding: %w", err)
	}

	ctx, cancel := context.WithCancel(ctx)

	return &BlockTagService{
		ctx:               ctx,
		cancel:            cancel,
		l1Client:          l1Client,
		l2Client:          l2Client,
		rollup:            rollup,
		rollupAddress:     config.RollupAddress,
		safeConfirmations: config.SafeConfirmations,
		pollInterval:      config.PollInterval,
		logger:            logger.With("module", "blocktag"),
		stop:              make(chan struct{}),
	}, nil
}

// Start starts the BlockTagService
func (s *BlockTagService) Start() error {
	s.logger.Info("Starting BlockTagService",
		"safeConfirmations", s.safeConfirmations,
		"pollInterval", s.pollInterval,
	)

	// Initialize by checking current L1 batch status
	if err := s.initialize(); err != nil {
		s.logger.Error("Failed to initialize BlockTagService", "error", err)
		// Don't return error, let the service start and retry
	}

	go s.loop()
	return nil
}

// Stop stops the BlockTagService
func (s *BlockTagService) Stop() {
	s.logger.Info("Stopping BlockTagService")
	s.cancel()
	<-s.stop
	s.l1Client.Close()
	s.logger.Info("BlockTagService stopped")
}

// initialize initializes the service by checking current L1 batch status
func (s *BlockTagService) initialize() error {
	s.logger.Info("Initializing BlockTagService")
	return s.updateBlockTags()
}

// loop is the main loop that polls L1 for batch status updates
func (s *BlockTagService) loop() {
	defer close(s.stop)

	ticker := time.NewTicker(s.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			if err := s.updateBlockTags(); err != nil {
				s.logger.Error("Failed to update block tags", "error", err)
			}
		}
	}
}

// updateBlockTags updates the safe and finalized block tags based on L1 batch tx status
func (s *BlockTagService) updateBlockTags() error {
	l2Head, err := s.l2Client.BlockNumber(s.ctx)
	if err != nil {
		return fmt.Errorf("failed to get L2 head: %w", err)
	}

	var safeBlockNum uint64
	var safeBlockHash common.Hash

	// Update safe block
	safeBlockNum, safeBlockHash, err = s.getL2BlockForTag(TagTypeSafe, l2Head)
	if err != nil {
		s.logger.Error("Failed to get safe L2 block", "error", err)
	} else if safeBlockHash != (common.Hash{}) {
		s.setSafeL2Block(safeBlockHash)
	}

	// Update finalized block
	finalizedBlockNum, finalizedBlockHash, err := s.getL2BlockForTag(TagTypeFinalized, l2Head)
	if err != nil {
		s.logger.Error("Failed to get finalized L2 block", "error", err)
	} else if finalizedBlockHash != (common.Hash{}) {
		// If finalized > safe, update safe to finalized (finalized is a stronger state)
		if finalizedBlockNum > safeBlockNum {
			safeBlockHash = finalizedBlockHash
			s.setSafeL2Block(safeBlockHash)
		}
		s.setFinalizedL2Block(finalizedBlockHash)
	}

	// Notify geth
	if err := s.notifyGeth(); err != nil {
		s.logger.Error("Failed to notify geth of block tags", "error", err)
	}

	s.logger.Debug("Block tags updated",
		"l2Head", l2Head,
		"safeL2BlockHash", s.safeL2BlockHash.Hex(),
		"finalizedL2BlockHash", s.finalizedL2BlockHash.Hex(),
	)

	return nil
}

// getL2BlockForTag gets the L2 block number and hash based on the L1 block tag
// Also validates state root matches between L1 batch and L2 block
func (s *BlockTagService) getL2BlockForTag(tagType BlockTagType, l2Head uint64) (uint64, common.Hash, error) {
	var l1BlockTag rpc.BlockNumber

	switch tagType {
	case TagTypeSafe:
		latestL1, err := s.l1Client.BlockNumber(s.ctx)
		if err != nil {
			return 0, common.Hash{}, fmt.Errorf("failed to get L1 latest block: %w", err)
		}
		if latestL1 <= s.safeConfirmations {
			return 0, common.Hash{}, nil
		}
		l1BlockTag = rpc.BlockNumber(latestL1 - s.safeConfirmations)

	case TagTypeFinalized:
		l1BlockTag = rpc.FinalizedBlockNumber

	default:
		return 0, common.Hash{}, fmt.Errorf("unknown tag type: %d", tagType)
	}

	// Query rollup contract at specified L1 block
	lastCommittedBatchIndex, err := s.getLastCommittedBatchAtBlock(l1BlockTag)
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("failed to get last committed batch: %w", err)
	}
	if lastCommittedBatchIndex == 0 {
		return 0, common.Hash{}, nil
	}

	// Find the largest completed batch (lastL2Block <= l2Head)
	// This works for both synced and syncing scenarios
	targetBatchIndex, targetBatchLastBlockNum, err := s.findCompletedBatchForL2Block(tagType, l2Head, lastCommittedBatchIndex)
	if err != nil {
		s.logger.Debug("No completed batch found", "l2Head", l2Head, "error", err)
		return 0, common.Hash{}, nil
	}

	// Validate state root.
	// Skip validation for already finalized batches, as their state roots may have been
	// deleted from the L1 contract after finalization
	lastFinalizedBatchIndex, err := s.rollup.LastFinalizedBatchIndex(nil)
	if err != nil {
		s.logger.Info("Failed to get last finalized batch index, skipping state root validation", "error", err)
		return 0, common.Hash{}, nil
	}
	if targetBatchIndex < lastFinalizedBatchIndex.Uint64() {
		// Batch data may have been deleted after finalization, cannot validate
		// Return error so caller skips this batch and keeps previous safe/finalized value
		// TODO: optimize this by using a different approach to get the state root
		s.logger.Info("batch %d already finalized (lastFinalized=%d), state root may be deleted",
			targetBatchIndex, lastFinalizedBatchIndex.Uint64())
		return 0, common.Hash{}, nil
	}
	if err := s.validateBatchStateRoot(targetBatchIndex, targetBatchLastBlockNum); err != nil {
		s.logger.Error("State root validation failed",
			"tagType", tagType,
			"batchIndex", targetBatchIndex,
			"l2Block", targetBatchLastBlockNum,
			"error", err,
		)
		return 0, common.Hash{}, err
	}

	// Get L2 block header for hash
	l2Header, err := s.l2Client.HeaderByNumber(s.ctx, big.NewInt(int64(targetBatchLastBlockNum)))
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("failed to get L2 block header: %w", err)
	}

	l2BlockHash := l2Header.Hash()

	s.logger.Debug("Got L2 block for tag",
		"tagType", tagType,
		"l1BlockTag", l1BlockTag,
		"batchIndex", targetBatchIndex,
		"l2Block", targetBatchLastBlockNum,
		"l2BlockHash", l2BlockHash.Hex(),
	)

	return targetBatchLastBlockNum, l2BlockHash, nil
}

// validateBatchStateRoot validates that the state root of batch's lastL2Block matches L1
func (s *BlockTagService) validateBatchStateRoot(batchIndex uint64, batchLastBlockNum uint64) error {
	// Get L2 block header
	l2Header, err := s.l2Client.HeaderByNumber(s.ctx, big.NewInt(int64(batchLastBlockNum)))
	if err != nil {
		return fmt.Errorf("failed to get L2 block header for block %d: %w", batchLastBlockNum, err)
	}

	// Get state root from L1 committed batch
	stateRoot, err := s.rollup.CommittedStateRoots(nil, big.NewInt(int64(batchIndex)))
	if err != nil {
		return fmt.Errorf("failed to get state root from L1: %w", err)
	}

	// Compare state roots
	l1StateRoot := common.BytesToHash(stateRoot[:])
	if l1StateRoot != l2Header.Root {
		return fmt.Errorf("state root mismatch for batch %d: L1=%s, L2=%s", batchIndex, l1StateRoot.Hex(), l2Header.Root.Hex())
	}

	return nil
}

// findCompletedBatchForL2Block finds the largest batch where lastL2Block <= l2BlockNum.
// Uses cached index for optimization: first call binary search, subsequent calls search forward.
// Separate caches for safe and finalized to avoid conflicts.
func (s *BlockTagService) findCompletedBatchForL2Block(tagType BlockTagType, l2HeaderNum uint64, lastCommittedBatchIndex uint64) (uint64, uint64, error) {
	return s.findCompletedBatchForL2BlockWithDepth(tagType, l2HeaderNum, lastCommittedBatchIndex, 0)
}

// findCompletedBatchForL2BlockWithDepth is the internal implementation with recursion depth limit.
// maxDepth is set to 1 to allow one retry after cache reset.
func (s *BlockTagService) findCompletedBatchForL2BlockWithDepth(tagType BlockTagType, l2HeaderNum uint64, lastCommittedBatchIndex uint64, depth int) (uint64, uint64, error) {
	const maxDepth = 2

	if lastCommittedBatchIndex == 0 {
		return 0, 0, fmt.Errorf("no batches available")
	}

	// Get cached index based on tag type
	startIdx := s.getCachedBatchIndex(tagType)
	if startIdx == 0 || startIdx > lastCommittedBatchIndex {
		// First time or cache invalid: use binary search to find starting point
		startIdx = s.binarySearchBatch(l2HeaderNum, lastCommittedBatchIndex)
		if startIdx == 0 {
			return 0, 0, fmt.Errorf("no completed batch found for L2 block %d", l2HeaderNum)
		}
	}

	// Search forward from startIdx
	var resultIdx, resultLastL2Block uint64
	for idx := startIdx; idx <= lastCommittedBatchIndex; idx++ {
		batchData, err := s.rollup.BatchDataStore(nil, big.NewInt(int64(idx)))
		if err != nil {
			return 0, 0, fmt.Errorf("failed to get batch data for index %d: %w", idx, err)
		}

		lastL2Block := batchData.BlockNumber.Uint64()
		if lastL2Block <= l2HeaderNum {
			resultIdx = idx
			resultLastL2Block = lastL2Block
			s.setCachedBatchIndex(tagType, idx)
		} else {
			break
		}
	}

	// Handle L2 reorg: if cache was too new, reset and use binary search
	if resultIdx == 0 {
		if depth >= maxDepth {
			return 0, 0, fmt.Errorf("no completed batch found for L2 block %d after retry", l2HeaderNum)
		}
		s.setCachedBatchIndex(tagType, 0)
		return s.findCompletedBatchForL2BlockWithDepth(tagType, l2HeaderNum, lastCommittedBatchIndex, depth+1)
	}

	return resultIdx, resultLastL2Block, nil
}

func (s *BlockTagService) getCachedBatchIndex(tagType BlockTagType) uint64 {
	if tagType == TagTypeSafe {
		return s.lastKnownSafeBatchIndex
	}
	return s.lastKnownFinalizedBatchIndex
}

func (s *BlockTagService) setCachedBatchIndex(tagType BlockTagType, idx uint64) {
	if tagType == TagTypeSafe {
		s.lastKnownSafeBatchIndex = idx
	} else {
		s.lastKnownFinalizedBatchIndex = idx
	}
}

// binarySearchBatch finds the largest batch index where lastL2BlockInBatch <= l2HeaderNum
func (s *BlockTagService) binarySearchBatch(l2HeaderNum uint64, maxBatchIndex uint64) uint64 {
	low, high := uint64(1), maxBatchIndex
	var result uint64

	for low <= high {
		mid := (low + high) / 2
		batchData, err := s.rollup.BatchDataStore(nil, big.NewInt(int64(mid)))
		if err != nil {
			return result // Return best result so far on error
		}

		if batchData.BlockNumber.Uint64() <= l2HeaderNum {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

// getLastCommittedBatchAtBlock queries the rollup contract at a specific L1 block
func (s *BlockTagService) getLastCommittedBatchAtBlock(l1BlockTag rpc.BlockNumber) (uint64, error) {
	var blockNum *big.Int
	if l1BlockTag == rpc.FinalizedBlockNumber {
		blockNum = big.NewInt(int64(rpc.FinalizedBlockNumber))
	} else if l1BlockTag >= 0 {
		blockNum = big.NewInt(int64(l1BlockTag))
	}

	lastCommitted, err := s.rollup.LastCommittedBatchIndex(&bind.CallOpts{
		BlockNumber: blockNum,
		Context:     s.ctx,
	})
	if err != nil {
		return 0, err
	}

	return lastCommitted.Uint64(), nil
}

// setSafeL2Block sets the safe L2 block hash
func (s *BlockTagService) setSafeL2Block(blockHash common.Hash) {
	if blockHash != s.safeL2BlockHash {
		s.safeL2BlockHash = blockHash
		s.logger.Info("Updated safe L2 block", "hash", blockHash.Hex())
	}
}

// setFinalizedL2Block sets the finalized L2 block hash
func (s *BlockTagService) setFinalizedL2Block(blockHash common.Hash) {
	if blockHash != s.finalizedL2BlockHash {
		s.finalizedL2BlockHash = blockHash
		s.logger.Info("Updated finalized L2 block", "hash", blockHash.Hex())
	}
}

// notifyGeth notifies geth of the new block tags via RPC
// Only calls RPC if there are changes since last notification
func (s *BlockTagService) notifyGeth() error {
	safeBlockHash := s.safeL2BlockHash
	finalizedBlockHash := s.finalizedL2BlockHash

	// Skip if no changes
	if safeBlockHash == s.lastNotifiedSafeHash && finalizedBlockHash == s.lastNotifiedFinalizedHash {
		return nil
	}

	// Skip if both are empty
	if safeBlockHash == (common.Hash{}) && finalizedBlockHash == (common.Hash{}) {
		return nil
	}

	if err := s.l2Client.SetBlockTags(s.ctx, safeBlockHash, finalizedBlockHash); err != nil {
		return err
	}

	// Update last notified hashes
	s.lastNotifiedSafeHash = safeBlockHash
	s.lastNotifiedFinalizedHash = finalizedBlockHash
	return nil
}
