package batch

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"morph-l2/tx-submitter/db"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/types"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"
)

// BatchCache is a structure for caching and building batch data
// Stores all batch information starting from 0, and has the functionality to pack batches
type BatchCache struct {
	mu       sync.RWMutex
	ctx      context.Context
	initDone bool

	batchStorage *BatchStorage

	// key: batchIndex, value: RPCRollupBatch
	sealedBatches      map[uint64]*eth.RPCRollupBatch
	sealedBatchHeaders map[uint64]*BatchHeaderBytes

	// Currently accumulating batch data (referencing node's BatchingCache)
	// Parent batch information
	parentBatchHeader *BatchHeaderBytes
	prevStateRoot     common.Hash

	// Accumulated batch data
	batchData             *BatchData
	totalL1MessagePopped  uint64
	postStateRoot         common.Hash
	withdrawRoot          common.Hash
	lastPackedBlockHeight uint64

	// Currently processing block data (referencing node's BatchingCache)
	// This data will not be appended to batch until block is confirmed
	currentBlockContext               []byte
	currentTxsPayload                 []byte
	currentL1TxsHashes                []common.Hash
	totalL1MessagePoppedAfterCurBlock uint64
	currentStateRoot                  common.Hash
	currentWithdrawRoot               common.Hash
	currentBlockNumber                uint64
	currentBlockHash                  common.Hash

	// Function to determine if batch is upgraded
	isBatchUpgraded func(uint64) bool

	// Clients and contracts
	l1Client       iface.Client
	l2Clients      iface.L2Clients
	rollupContract iface.IRollup
	l2Caller       *types.L2Caller

	// config
	batchTimeOut  uint64
	blockInterval uint64
}

// NewBatchCache creates and initializes a new BatchCache instance
func NewBatchCache(
	isBatchUpgraded func(uint64) bool,
	l1Client iface.Client,
	l2Clients []iface.L2Client,
	rollupContract iface.IRollup,
	l2Caller *types.L2Caller,
	ldb *db.Db,
) *BatchCache {
	if isBatchUpgraded == nil {
		// Default implementation: always returns true (use V1 version)
		isBatchUpgraded = func(uint64) bool { return true }
	}
	ctx := context.Background()
	ifL2Clients := iface.L2Clients{Clients: l2Clients}
	_, err := ifL2Clients.BlockNumber(ctx)
	if err != nil {
		log.Error("Error getting block number", "err", err)
	}
	return &BatchCache{
		ctx:                               ctx,
		initDone:                          false,
		sealedBatches:                     make(map[uint64]*eth.RPCRollupBatch),
		sealedBatchHeaders:                make(map[uint64]*BatchHeaderBytes),
		parentBatchHeader:                 nil,
		prevStateRoot:                     common.Hash{},
		batchData:                         NewBatchData(),
		totalL1MessagePopped:              0,
		postStateRoot:                     common.Hash{},
		withdrawRoot:                      common.Hash{},
		lastPackedBlockHeight:             0,
		currentBlockContext:               nil,
		currentTxsPayload:                 nil,
		currentL1TxsHashes:                nil,
		totalL1MessagePoppedAfterCurBlock: 0,
		currentStateRoot:                  common.Hash{},
		currentWithdrawRoot:               common.Hash{},
		currentBlockNumber:                0,
		currentBlockHash:                  common.Hash{},
		isBatchUpgraded:                   isBatchUpgraded,
		l1Client:                          l1Client,
		l2Clients:                         iface.L2Clients{Clients: l2Clients},
		rollupContract:                    rollupContract,
		l2Caller:                          l2Caller,
		batchStorage:                      NewBatchStorage(ldb),
	}
}

func (bc *BatchCache) Init() error {
	err := bc.updateBatchConfigFromGov()
	if err != nil {
		return err
	}
	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("get batch status from rollup failed err: %w", err)
	}
	headerBytes, err := bc.getLastFinalizeBatchHeaderFromRollupByIndex(fi.Uint64())
	if err != nil {
		return fmt.Errorf("get last finalize batch header err: %w", err)
	}
	// Initialize BatchCache parent batch information
	// prevStateRoot should be the parent batch's postStateRoot
	bc.parentBatchHeader = headerBytes
	bc.prevStateRoot, err = headerBytes.PostStateRoot()
	if err != nil {
		return fmt.Errorf("get post state root err: %w", err)
	}
	bc.lastPackedBlockHeight, err = headerBytes.LastBlockNumber()
	if err != nil {
		store, err := bc.rollupContract.BatchDataStore(nil, fi)
		if err != nil {
			return err
		}
		bc.lastPackedBlockHeight = store.BlockNumber.Uint64()
	}
	bc.currentBlockNumber = bc.lastPackedBlockHeight
	bc.totalL1MessagePopped, err = headerBytes.TotalL1MessagePopped()
	if err != nil {
		return fmt.Errorf("get total l1 message popped err: %w", err)
	}
	log.Info("Start assemble batch", "start batch", fi.Uint64(), "end batch", ci.Uint64())
	return nil
}

func (bc *BatchCache) InitFromRollupByRange() error {
	if bc.initDone {
		return nil
	}
	err := bc.Init()
	if err != nil {
		return err
	}
	err = bc.assembleUnFinalizeBatchHeaderFromL2Blocks()
	if err != nil {
		return err
	}
	bc.initDone = true
	log.Info("Initialized batch cache success")
	return nil
}

func (bc *BatchCache) InitAndSyncFromDatabase() error {
	if bc.initDone {
		return nil
	}
	err := bc.updateBatchConfigFromGov()
	if err != nil {
		return err
	}
	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("get batch status from rollup failed err: %w", err)
	}

	batches, headers, indices, err := bc.batchStorage.LoadAllSealedBatchesAndHeader()
	if err != nil {
		log.Error("Failed to load sealed batch headers from storage", "error", err)
		return bc.DeleteBatchStorageAndInitFromRollup()
	}

	if len(batches) == 0 {
		return bc.InitAndSyncFromRollup()
	}
	maxIndex := indices[0]
	for _, idx := range indices {
		if idx > maxIndex {
			maxIndex = idx
		}
	}
	// check batch hash with the batch that already rollup by submitter
	for i := fi.Uint64(); i <= ci.Uint64(); i++ {
		batchHash, err := bc.rollupContract.CommittedBatches(nil, new(big.Int).SetUint64(i))
		if err != nil {
			return err
		}
		batchStorage, exist := batches[i]
		if !exist || !bytes.Equal(batchHash[:], batchStorage.Hash.Bytes()) {
			// batch not contiguous or batch is invalid
			return bc.DeleteBatchStorageAndInitFromRollup()
		}
	}

	latestHeaderBytes := headers[maxIndex]
	prevStateRoot, err := latestHeaderBytes.PostStateRoot()
	if err != nil {
		log.Error("Get post state root failed", "err", err)
		return bc.DeleteBatchStorageAndInitFromRollup()
	}
	totalL1MessagePopped, err := latestHeaderBytes.TotalL1MessagePopped()
	if err != nil {
		log.Error("Get total l1 message popped failed", "err", err)
		return bc.DeleteBatchStorageAndInitFromRollup()
	}
	lastPackedBlockHeight, err := latestHeaderBytes.LastBlockNumber()
	if err != nil {
		// maybe the latest header is version 0 which do not have blockNum
		latestBatchIndex, err := latestHeaderBytes.BatchIndex()
		if err != nil {
			return fmt.Errorf("get batch index from parent header failed err: %w", err)
		}
		// check batch index range
		if latestBatchIndex < fi.Uint64() || latestBatchIndex > ci.Uint64() {
			// missing batch data, sync from another side
			log.Error("Batch index is out of range",
				"latestBatchIndex", latestBatchIndex,
				"commitIndex", ci.Uint64(), "finalizeIndex", fi.Uint64())
			return bc.DeleteBatchStorageAndInitFromRollup()
		}
		store, err := bc.rollupContract.BatchDataStore(nil, new(big.Int).SetUint64(latestBatchIndex))
		if err != nil {
			log.Error("Failed to load latest batch index from rollup",
				"error", err,
				"batchIndex", latestBatchIndex)
			return bc.DeleteBatchStorageAndInitFromRollup()
		}
		lastPackedBlockHeight = store.BlockNumber.Uint64()
	}
	bc.lastPackedBlockHeight = lastPackedBlockHeight
	bc.sealedBatches = batches
	bc.sealedBatchHeaders = headers
	bc.parentBatchHeader = latestHeaderBytes
	bc.currentBlockNumber = bc.lastPackedBlockHeight
	bc.prevStateRoot = prevStateRoot
	bc.totalL1MessagePopped = totalL1MessagePopped

	bc.initDone = true
	log.Info("Sync sealed batch from database success", "count", len(batches))
	return nil
}

func (bc *BatchCache) InitAndSyncFromRollup() error {
	if bc.initDone {
		return nil
	}
	err := bc.Init()
	if err != nil {
		return err
	}
	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("get batch status from rollup failed err: %w", err)
	}
	log.Info("Start assemble batch",
		"startBatch", fi.Uint64()+1,
		"endBatch", ci.Uint64(),
		"startNum", bc.lastPackedBlockHeight,
		"prevStateRoot", bc.prevStateRoot.String(),
	)
	for i := fi.Uint64() + 1; i <= ci.Uint64(); i++ {
		batchIndex := new(big.Int).SetUint64(i)
		startNum, endNum, err := bc.getBatchBlockRange(batchIndex)
		if err != nil {
			return fmt.Errorf("get batch block range err: %w,start %v, end %v", err, startNum, endNum)
		}
		log.Info("assemble batch block range", "startNum", startNum, "endNum", endNum)
		batchHeaderBytes, err := bc.assembleBatchHeaderFromL2Blocks(startNum, endNum)
		if err != nil {
			return err
		}
		batchHash, err := batchHeaderBytes.Hash()
		if err != nil {
			return fmt.Errorf("get batch hash err: %w", err)
		}
		correct, err := bc.checkBatchHashCorrect(batchIndex, batchHash)
		if err != nil {
			return fmt.Errorf("check batch hash failed, err: %w, batchIndex %v, batchHash %v", err, batchIndex, batchHash.String())
		}
		if !correct {
			return fmt.Errorf("batch hash check failed: batch index %d is incorrect", i)
		}
		log.Info("Assemble batch success", "batch index", i, "last batch index", ci.Uint64())
	}
	bc.initDone = true
	log.Info("Initialized batch cache success")
	return nil
}

func (bc *BatchCache) LatestBatchIndex() (uint64, error) {
	return bc.parentBatchHeader.BatchIndex()
}

func (bc *BatchCache) updateBatchConfigFromGov() error {
	interval, err := bc.l2Caller.BatchBlockInterval(nil)
	if err != nil {
		return err
	}
	timeout, err := bc.l2Caller.BatchTimeout(nil)
	if err != nil {
		return err
	}
	bc.batchTimeOut = timeout.Uint64()
	bc.blockInterval = interval.Uint64()
	log.Info("Update batch config success", "interval", interval.Uint64(), "timeout", timeout.Uint64())
	return nil
}

func (bc *BatchCache) checkBatchHashCorrect(batchIndex *big.Int, batchHash common.Hash) (bool, error) {
	commitBatchHash, err := bc.rollupContract.CommittedBatches(nil, batchIndex)
	if err != nil {
		return false, err
	}
	if !bytes.Equal(commitBatchHash[:], batchHash.Bytes()) {
		log.Error("check commit batch hash failed",
			"index", batchIndex.String(),
			"committed", hex.EncodeToString(commitBatchHash[:]),
			"generate", batchHash.String())
		return false, nil
	}
	return true, nil
}

func (bc *BatchCache) getBatchStatusFromContract() (*big.Int, *big.Int, error) {
	latestCommitBatchIndex, err := bc.rollupContract.LastCommittedBatchIndex(nil)
	if err != nil {
		return nil, nil, err
	}
	lastFinalizedBatchIndex, err := bc.rollupContract.LastFinalizedBatchIndex(nil)
	if err != nil {
		return nil, nil, err
	}
	return latestCommitBatchIndex, lastFinalizedBatchIndex, nil
}

func (bc *BatchCache) getBatchBlockRange(batchIndex *big.Int) (uint64, uint64, error) {
	preIndex := new(big.Int).Sub(batchIndex, big.NewInt(1))
	preBatchStorage, err := bc.rollupContract.BatchDataStore(nil, preIndex)
	if err != nil {
		return 0, 0, err
	}
	batchStorage, err := bc.rollupContract.BatchDataStore(nil, batchIndex)
	if err != nil {
		return 0, 0, err
	}
	return preBatchStorage.BlockNumber.Uint64() + 1, batchStorage.BlockNumber.Uint64(), nil
}

func (bc *BatchCache) getUnFinalizeBlockRange() (uint64, uint64, *big.Int, error) {
	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return 0, 0, nil, err
	}
	finalizeBatchStorage, err := bc.rollupContract.BatchDataStore(nil, fi)
	if err != nil {
		return 0, 0, nil, err
	}
	startNum := finalizeBatchStorage.BlockNumber.Uint64() + 1
	endNum, err := bc.l2Clients.BlockNumber(context.Background())
	if err != nil {
		return 0, 0, nil, err
	}
	return startNum, endNum, ci, nil
}

// IsEmpty checks if current batch data is empty
func (bc *BatchCache) IsEmpty() bool {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return bc.batchData == nil || bc.batchData.IsEmpty()
}

// IsCurrentEmpty checks if current block data is empty
func (bc *BatchCache) IsCurrentEmpty() bool {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return len(bc.currentBlockContext) == 0
}

// ClearCurrent clears current block data
// Note: lock must be held before calling this method
func (bc *BatchCache) ClearCurrent() {
	bc.currentTxsPayload = nil
	bc.currentL1TxsHashes = nil
	bc.currentBlockContext = nil
	bc.totalL1MessagePoppedAfterCurBlock = 0
	bc.currentStateRoot = common.Hash{}
	bc.currentWithdrawRoot = common.Hash{}
}

// GetSealedBatch gets sealed batch information
func (bc *BatchCache) GetSealedBatch(batchIndex uint64) (*eth.RPCRollupBatch, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	batch, ok := bc.sealedBatches[batchIndex]
	return batch, ok
}

// GetSealedBatchHeader gets sealed batch header information
func (bc *BatchCache) GetSealedBatchHeader(batchIndex uint64) (*BatchHeaderBytes, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	header, ok := bc.sealedBatchHeaders[batchIndex]
	if !ok {
		// Check again after acquiring write lock
		header, ok = bc.sealedBatchHeaders[batchIndex]
		if !ok {
			loadedHeader, err := bc.batchStorage.LoadSealedBatchHeader(batchIndex)
			if err != nil {
				return nil, false
			}
			return loadedHeader, true
		}
		return header, true
	}
	return header, ok
}

// GetLatestSealedBatchIndex gets the latest sealed batch index
func (bc *BatchCache) GetLatestSealedBatchIndex() uint64 {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	var maxIndex uint64 = 0
	for index := range bc.sealedBatches {
		if index > maxIndex {
			maxIndex = index
		}
	}
	return maxIndex
}

// CalculateCapWithProposalBlock calculates batch capacity after including the specified block
func (bc *BatchCache) CalculateCapWithProposalBlock(blockNumber uint64, withdrawRoot common.Hash) (bool, error) {
	if len(bc.l2Clients.Clients) == 0 {
		return false, fmt.Errorf("l2 client is nil")
	}

	// Fetch complete block from L2 client (including transactions)
	block, err := bc.l2Clients.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		return false, fmt.Errorf("failed to fetch block %d: %w", blockNumber, err)
	}

	if block == nil {
		return false, fmt.Errorf("block is nil for block %d", blockNumber)
	}

	header := block.Header()

	// Verify block number matches
	if header.Number.Uint64() != blockNumber {
		return false, fmt.Errorf("block number mismatch: expected %d, got %d", blockNumber, header.Number.Uint64())
	}

	bc.mu.Lock()
	defer bc.mu.Unlock()
	// Verify block number continuity
	if blockNumber <= bc.lastPackedBlockHeight {
		if blockNumber != 0 || bc.lastPackedBlockHeight != 0 {
			return false, fmt.Errorf("wrong block number: lastPackedBlockHeight=%d, proposed=%d", bc.lastPackedBlockHeight, blockNumber)
		}
	}
	if blockNumber > bc.lastPackedBlockHeight+1 {
		// Some blocks were skipped, need to clear cache
		return false, fmt.Errorf("discontinuous block number: lastPackedBlockHeight=%d, proposed=%d", bc.lastPackedBlockHeight, blockNumber)
	}

	// Ensure BatchData is initialized
	if bc.batchData == nil {
		bc.batchData = NewBatchData()
	}

	// Parse transactions, distinguish L1 and L2 transactions
	txsPayload, l1TxHashes, newTotalL1MessagePopped, l2TxNum, err := parsingTxs(block.Transactions(), bc.totalL1MessagePopped)
	if err != nil {
		return false, fmt.Errorf("failed to parse transactions: %w", err)
	}

	l1TxNum := int(newTotalL1MessagePopped - bc.totalL1MessagePopped)
	txsNum := l2TxNum + l1TxNum

	// Build BlockContext (60 bytes)
	blockContext := buildBlockContext(header, txsNum, l1TxNum)

	// Store to current, do not immediately append to batch
	bc.currentBlockContext = blockContext
	bc.currentTxsPayload = txsPayload
	bc.currentL1TxsHashes = l1TxHashes
	bc.totalL1MessagePoppedAfterCurBlock = newTotalL1MessagePopped
	bc.currentStateRoot = header.Root
	bc.currentBlockNumber = blockNumber
	bc.currentBlockHash = block.Hash()
	bc.currentWithdrawRoot = withdrawRoot

	// Check capacity: if compressed size would exceed limit after adding current block
	var exceeded bool
	if bc.isBatchUpgraded(header.Time) {
		exceeded, err = bc.batchData.WillExceedCompressedSizeLimit(blockContext, txsPayload)
	} else {
		exceeded, err = bc.batchData.EstimateCompressedSizeWithNewPayload(txsPayload)
	}
	if err != nil {
		return false, fmt.Errorf("failed to estimate compressed size: %w", err)
	}

	return exceeded, nil
}

// PackCurrentBlock packs current block data into batch
// References node's PackCurrentBlock
// Parameters:
//   - blockNumber: block number to pack (for verification)
//
// Returns:
//   - error: returns error if packing fails
//
// Note: This method should be called after block is confirmed, appending data from currentBlockContext to batch
func (bc *BatchCache) PackCurrentBlock(blockNumber uint64) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// If the current block is empty, return directly
	if len(bc.currentBlockContext) == 0 {
		return nil // nothing to pack
	}

	// Verify block number matches
	if bc.currentBlockNumber != blockNumber {
		return fmt.Errorf("block number mismatch: expected %d, got %d", blockNumber, bc.currentBlockNumber)
	}

	// Ensure BatchData is initialized
	if bc.batchData == nil {
		bc.batchData = NewBatchData()
	}

	// Append current block data to batch
	bc.batchData.Append(bc.currentBlockContext, bc.currentTxsPayload, bc.currentL1TxsHashes)

	// Update accumulated state
	bc.totalL1MessagePopped = bc.totalL1MessagePoppedAfterCurBlock
	bc.withdrawRoot = bc.currentWithdrawRoot
	bc.postStateRoot = bc.currentStateRoot
	bc.lastPackedBlockHeight = blockNumber

	// Clear current block data
	bc.ClearCurrent()

	return nil
}

// FetchAndCacheHeader fetches complete block from L2 client for specified block number, parses transactions and stores to current
// Note: This method has been replaced by CalculateCapWithProposalBlock and PackCurrentBlock
// Kept for backward compatibility, but recommend using new methods
func (bc *BatchCache) FetchAndCacheHeader(blockNumber uint64, withdrawRoot common.Hash) (*ethtypes.Header, error) {
	// Use new method
	_, err := bc.CalculateCapWithProposalBlock(blockNumber, withdrawRoot)
	if err != nil {
		return nil, err
	}

	// Pack immediately (backward compatible behavior)
	if err := bc.PackCurrentBlock(blockNumber); err != nil {
		return nil, err
	}

	bc.mu.RLock()
	defer bc.mu.RUnlock()

	// Return header (need to re-fetch because current has been cleared)
	block, err := bc.l2Clients.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
}

// SealBatch seals the currently accumulated batch, generates batch header and stores to sealedBatches
// Parameters:
//   - sequencerSetVerifyHash: sequencer set verification hash (obtained from L1 contract)
//   - blockTimestamp: current block timestamp (used to determine batch version)
//
// Returns:
//   - batchIndex: sealed batch index
//   - batchHash: batch hash
//   - reachedExpectedSize: whether the sealed data size reaches expected value (compressed payload size close to or reaches MaxBlobBytesSize)
//   - error: returns error if sealing fails
//
// Note: Sealed batch will be stored in BatchCache's sealedBatches, not sent anywhere
func (bc *BatchCache) SealBatch(sequencerSets []byte, blockTimestamp uint64) (uint64, BatchHeaderBytes, bool, error) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// Ensure batch data is not empty
	if bc.batchData == nil || bc.batchData.IsEmpty() {
		return 0, BatchHeaderBytes{}, false, errors.New("failed to seal batch: batch cache is empty")
	}

	// Compress data and calculate dataHash
	compressedPayload, batchDataHash, err := bc.handleBatchSealing(blockTimestamp)
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to handle batch sealing: %w", err)
	}

	// Check if sealed data size reaches expected value
	// Expected value: compressed payload size close to or reaches MaxBlobBytesSize
	// Use 90% as threshold, i.e., if compressed size >= MaxBlobBytesSize * 0.9, consider it reached expected
	threshold := float64(MaxBlobBytesSize) * 0.9
	expectedSizeThreshold := uint64(threshold)
	reachedExpectedSize := uint64(len(compressedPayload)) >= expectedSizeThreshold

	// Generate blob sidecar
	sidecar, err := MakeBlobTxSidecar(compressedPayload)
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to create blob sidecar: %w", err)
	}

	// Create batch header
	batchHeader := bc.createBatchHeader(batchDataHash, sidecar, crypto.Keccak256Hash(sequencerSets), blockTimestamp)

	// Calculate batch hash
	batchHash, err := batchHeader.Hash()
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to hash batch header: %w", err)
	}

	// Get batch index
	batchIndex, err := batchHeader.BatchIndex()
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to get batch index: %w", err)
	}

	// Build parent batch header bytes
	var parentBatchHeaderBytes hexutil.Bytes
	if bc.parentBatchHeader != nil {
		parentBatchHeaderBytes = hexutil.Bytes(*bc.parentBatchHeader)
	}

	// Get the version from batch header
	version, err := batchHeader.Version()
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to get batch version: %w", err)
	}

	// Build block contexts from batch data (encode block contexts)
	blockContextsData, err := bc.batchData.Encode()
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to encode batch data: %w", err)
	}
	blockContexts := hexutil.Bytes(blockContextsData)

	// Convert sequencerSetVerifyHash to bytes
	currentSequencerSetBytes := hexutil.Bytes(sequencerSets)

	// Get L1 message count from batch data
	numL1Messages := bc.batchData.l1TxNum

	// Store sealed batch information as RPCRollupBatch
	sealedBatch := &eth.RPCRollupBatch{
		Version:                  uint(version),
		Hash:                     batchHash,
		ParentBatchHeader:        parentBatchHeaderBytes,
		BlockContexts:            blockContexts,
		CurrentSequencerSetBytes: currentSequencerSetBytes,
		PrevStateRoot:            bc.prevStateRoot,
		PostStateRoot:            bc.postStateRoot,
		WithdrawRoot:             bc.withdrawRoot,
		LastBlockNumber:          bc.lastPackedBlockHeight,
		NumL1Messages:            numL1Messages,
		Sidecar:                  *sidecar,
		Signatures:               []eth.RPCBatchSignature{},
		CollectedL1Fee:           nil,
	}
	bc.sealedBatches[batchIndex] = sealedBatch
	// Store batch header copy
	batchHeaderCopy := make(BatchHeaderBytes, len(batchHeader))
	copy(batchHeaderCopy, batchHeader)
	bc.sealedBatchHeaders[batchIndex] = &batchHeaderCopy

	err = bc.batchStorage.StoreSealedBatch(batchIndex, sealedBatch)
	if err != nil {
		log.Error("failed to store sealed batch", "err", err)
	}
	err = bc.batchStorage.StoreSealedBatchHeader(batchIndex, &batchHeaderCopy)
	if err != nil {
		log.Error("failed to store sealed batch header", "err", err)
	}
	// Update parent batch information for next batch
	bc.parentBatchHeader = &batchHeader
	bc.prevStateRoot = bc.postStateRoot

	bc.logSealedBatch(batchHeader, batchHash)

	// Reset currently accumulated batch data
	bc.batchData = NewBatchData()

	return batchIndex, batchHeader, reachedExpectedSize, nil
}

// CheckBatchSizeReached checks if the specified batch's data size reaches expected value
// Parameters:
//   - batchIndex: batch index to check
//
// Returns:
//   - reached: returns true if batch exists and compressed payload size reaches expected value (>= MaxBlobBytesSize * 0.9)
//   - found: whether batch exists
func (bc *BatchCache) CheckBatchSizeReached(batchIndex uint64) (reached bool, found bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	sealedBatch, ok := bc.sealedBatches[batchIndex]
	if !ok {
		return false, false
	}

	// Expected value: compressed payload size >= MaxBlobBytesSize * 0.9
	// We need to estimate the compressed size from the block contexts
	// For now, we'll use a simple heuristic based on block contexts size
	threshold := float64(MaxBlobBytesSize) * 0.9
	expectedSizeThreshold := uint64(threshold)

	// Estimate compressed size from block contexts (rough approximation)
	blockContextsSize := uint64(len(sealedBatch.BlockContexts))
	// Use a compression ratio estimate (zstd typically achieves 2-3x compression)
	estimatedCompressedSize := blockContextsSize / 2
	reached = estimatedCompressedSize >= expectedSizeThreshold

	return reached, true
}

// handleBatchSealing determines which version to use for compression and calculates data hash
func (bc *BatchCache) handleBatchSealing(blockTimestamp uint64) ([]byte, common.Hash, error) {
	var (
		compressedPayload []byte
		batchDataHash     common.Hash
		err               error
	)

	// Check if upgraded version should be used
	if bc.isBatchUpgraded(blockTimestamp) {
		compressedPayload, err = CompressBatchBytes(bc.batchData.TxsPayloadV2())
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to compress upgraded payload: %w", err)
		}

		if len(compressedPayload) <= MaxBlobBytesSize {
			batchDataHash, err = bc.batchData.DataHashV2()
			if err != nil {
				return nil, common.Hash{}, fmt.Errorf("failed to calculate upgraded data hash: %w", err)
			}
			return compressedPayload, batchDataHash, nil
		}
	}

	// Fall back to the old version
	compressedPayload, err = CompressBatchBytes(bc.batchData.TxsPayload())
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("failed to compress payload: %w", err)
	}
	batchDataHash = bc.batchData.DataHash()

	return compressedPayload, batchDataHash, nil
}

// createBatchHeader creates BatchHeader
func (bc *BatchCache) createBatchHeader(dataHash common.Hash, sidecar *ethtypes.BlobTxSidecar, sequencerSetVerifyHash common.Hash, blockTimestamp uint64) BatchHeaderBytes {
	blobHashes := []common.Hash{EmptyVersionedHash}
	if sidecar != nil && len(sidecar.Blobs) > 0 {
		blobHashes = sidecar.BlobHashes()
	}

	var parentBatchHeaderTotalL1 uint64
	var parentBatchIndex uint64
	var parentBatchHash common.Hash

	if bc.parentBatchHeader != nil {
		parentBatchHeaderTotalL1, _ = bc.parentBatchHeader.TotalL1MessagePopped()
		parentBatchIndex, _ = bc.parentBatchHeader.BatchIndex()
		parentBatchHash, _ = bc.parentBatchHeader.Hash()
	}

	l1MessagePopped := bc.totalL1MessagePopped - parentBatchHeaderTotalL1

	batchHeaderV0 := BatchHeaderV0{
		BatchIndex:             parentBatchIndex + 1,
		L1MessagePopped:        l1MessagePopped,
		TotalL1MessagePopped:   bc.totalL1MessagePopped,
		DataHash:               dataHash,
		BlobVersionedHash:      blobHashes[0],
		PrevStateRoot:          bc.prevStateRoot,
		PostStateRoot:          bc.postStateRoot,
		WithdrawalRoot:         bc.withdrawRoot,
		SequencerSetVerifyHash: sequencerSetVerifyHash,
		ParentBatchHash:        parentBatchHash,
	}

	if bc.isBatchUpgraded(blockTimestamp) {
		batchHeaderV1 := BatchHeaderV1{
			BatchHeaderV0:   batchHeaderV0,
			LastBlockNumber: bc.lastPackedBlockHeight,
		}
		return batchHeaderV1.Bytes()
	}

	return batchHeaderV0.Bytes()
}

// parsingTxs parses transactions, distinguishes L1 and L2 transactions
func parsingTxs(transactions []*ethtypes.Transaction, totalL1MessagePoppedBefore uint64) (
	txsPayload []byte,
	l1TxHashes []common.Hash,
	totalL1MessagePopped uint64,
	l2TxNum int,
	err error,
) {
	nextIndex := totalL1MessagePoppedBefore

	for i, tx := range transactions {
		if isL1MessageTxType(tx) {
			l1TxHashes = append(l1TxHashes, tx.Hash())
			currentIndex := tx.L1MessageQueueIndex()

			if currentIndex < nextIndex {
				return nil, nil, 0, 0, fmt.Errorf(
					"unexpected batch payload, expected queue index: %d, got: %d. transaction hash: %v",
					nextIndex, currentIndex, tx.Hash(),
				)
			}

			nextIndex = currentIndex + 1
			continue
		}

		l2TxNum++
		txBytes, err := tx.MarshalBinary()
		if err != nil {
			return nil, nil, 0, 0, fmt.Errorf("failed to marshal transaction %d: %w", i, err)
		}
		txsPayload = append(txsPayload, txBytes...)
	}

	totalL1MessagePopped = nextIndex
	return
}

// isL1MessageTxType checks if transaction is L1 message transaction type
func isL1MessageTxType(tx *ethtypes.Transaction) bool {
	return tx.Type() == ethtypes.L1MessageTxType
}

// buildBlockContext builds BlockContext from block header (60 bytes)
// Format: Number(8) || Timestamp(8) || BaseFee(32) || GasLimit(8) || numTxs(2) || numL1Messages(2)
func buildBlockContext(header *ethtypes.Header, txsNum, l1MsgNum int) []byte {
	blsBytes := make([]byte, 60)

	// Number (8 bytes)
	binary.BigEndian.PutUint64(blsBytes[:8], header.Number.Uint64())

	// Timestamp (8 bytes)
	binary.BigEndian.PutUint64(blsBytes[8:16], header.Time)

	// BaseFee (32 bytes)
	if header.BaseFee != nil {
		copy(blsBytes[16:48], header.BaseFee.FillBytes(make([]byte, 32)))
	} else {
		copy(blsBytes[16:48], make([]byte, 32))
	}

	// GasLimit (8 bytes)
	binary.BigEndian.PutUint64(blsBytes[48:56], header.GasLimit)

	// numTxs (2 bytes)
	binary.BigEndian.PutUint16(blsBytes[56:58], uint16(txsNum))

	// numL1Messages (2 bytes)
	binary.BigEndian.PutUint16(blsBytes[58:60], uint16(l1MsgNum))

	return blsBytes
}

func (bc *BatchCache) assembleBatchHeaderFromL2Blocks(
	startBlockNum, endBlockNum uint64,
) (*BatchHeaderBytes, error) {
	ctx := context.Background()
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	// Fetch blocks from L2 client in the specified range and accumulate to batch
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNum)
		root, err := bc.l2Caller.GetTreeRoot(callOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to get withdraw root at block %d: %w", blockNum, err)
		}

		// Check capacity and store to current
		_, err = bc.CalculateCapWithProposalBlock(blockNum, root)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate cap with block %d: %w", blockNum, err)
		}

		// Pack current block (confirm and append to batch)
		if err = bc.PackCurrentBlock(blockNum); err != nil {
			return nil, fmt.Errorf("failed to pack block %d: %w", blockNum, err)
		}
	}

	sequencerSet, _, err := bc.l2Caller.GetSequencerSetBytes(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get sequencer set verify hash at block %d: %w", callOpts.BlockNumber.Uint64(), err)
	}
	// Get the last block's timestamp for packing
	lastBlock, err := bc.l2Clients.BlockByNumber(ctx, big.NewInt(int64(endBlockNum)))
	if err != nil {
		return nil, fmt.Errorf("failed to get last block %d: %w", endBlockNum, err)
	}
	blockTimestamp := lastBlock.Time()

	// Seal batch and generate batchHeader
	batchIndex, batchHeader, reachedExpectedSize, err := bc.SealBatch(sequencerSet, blockTimestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to seal batch: %w", err)
	}

	batchHeaderHash, err := batchHeader.Hash()
	if err != nil {
		return nil, fmt.Errorf("failed to hash batch header: %w", err)
	}
	log.Info("seal batch success", "batchIndex", batchIndex, "batchHash", batchHeaderHash.String(), "reachedExpectedSize", reachedExpectedSize)
	return &batchHeader, nil
}

func (bc *BatchCache) assembleUnFinalizeBatchHeaderFromL2Blocks() error {
	ctx := context.Background()
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	startBlockNum, endBlockNum, ci, err := bc.getUnFinalizeBlockRange()
	if err != nil {
		return err
	}

	// Get start block once to avoid repeated queries
	startBlock, err := bc.l2Clients.BlockByNumber(ctx, big.NewInt(int64(startBlockNum)))
	if err != nil {
		return fmt.Errorf("failed to get start block %d: %w", startBlockNum, err)
	}
	startBlockTime := startBlock.Time()

	// Fetch blocks from L2 client in the specified range and accumulate to batch
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNum)
		root, err := bc.l2Caller.GetTreeRoot(callOpts)
		if err != nil {
			return fmt.Errorf("failed to get withdraw root at block %d: %w", blockNum, err)
		}

		// Check capacity and store to current
		exceeded, err := bc.CalculateCapWithProposalBlock(blockNum, root)
		if err != nil {
			return fmt.Errorf("failed to calculate cap with block %d: %w", blockNum, err)
		}

		// Get the current block to check timeout after packing
		nowBlock, err := bc.l2Clients.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
		if err != nil {
			return fmt.Errorf("failed to get block %d: %w", blockNum, err)
		}
		nowBlockTime := nowBlock.Time()

		// Check timeout: if elapsed time >= batchTimeOut, must seal batch immediately
		// This ensures batch is sealed before exceeding the maximum timeout configured in gov contract
		timeout := false
		if bc.batchTimeOut > 0 {
			elapsedTime := nowBlockTime - startBlockTime
			if elapsedTime >= bc.batchTimeOut {
				timeout = true
				log.Info("Batch timeout reached, must seal batch", "startBlock", startBlockNum, "currentBlock", blockNum,
					"elapsedTime", elapsedTime, "batchTimeOut", bc.batchTimeOut)
			}
		}

		// Check if we need to seal batch due to capacity, block interval, or timeout
		// check ensures batch is sealed before exceeding the maximum timeout
		if exceeded || (bc.blockInterval > 0 && (blockNum-startBlockNum+1) == bc.blockInterval) || timeout {
			log.Info("block exceeds limit", "start", startBlockNum, "to", blockNum-1, "exceeded", exceeded, "timeout", timeout)
			batchHash, reachedExpectedSize, batchIndex, err := bc.SealBatchAndCheck(callOpts, ci)
			if err != nil {
				return err
			}
			batch, ok := bc.GetSealedBatch(batchIndex)
			if !ok {
				return fmt.Errorf("batch %d not found in cache", batchIndex)
			}
			startBlockNum = batch.LastBlockNumber + 1
			startBlock, err = bc.l2Clients.BlockByNumber(ctx, big.NewInt(int64(startBlockNum)))
			if err != nil {
				return fmt.Errorf("failed to get start block %d: %w", startBlockNum, err)
			}
			startBlockTime = startBlock.Time()
			index, err := bc.parentBatchHeader.BatchIndex()
			if err != nil {
				return err
			}
			log.Info("seal batch success", "batchIndex", index, "batchHash", batchHash.String(), "reachedExpectedSize", reachedExpectedSize)
		}

		// Pack current block (confirm and append to batch)
		if err = bc.PackCurrentBlock(blockNum); err != nil {
			return fmt.Errorf("failed to pack block %d: %w", blockNum, err)
		}
	}
	return nil
}

func (bc *BatchCache) SealBatchAndCheck(callOpts *bind.CallOpts, ci *big.Int) (common.Hash, bool, uint64, error) {
	sequencerSetBytes, _, err := bc.l2Caller.GetSequencerSetBytes(callOpts)
	if err != nil {
		return common.Hash{}, false, 0, err
	}
	lastBlock, err := bc.l2Clients.BlockByNumber(context.Background(), big.NewInt(int64(bc.lastPackedBlockHeight)))
	if err != nil {
		return common.Hash{}, false, 0, fmt.Errorf("failed to get last block %d: %w", bc.lastPackedBlockHeight, err)
	}
	blockTimestamp := lastBlock.Time()
	// Seal batch and generate batchHeader
	batchIndex, batchHeaderBytes, reachedExpectedSize, err := bc.SealBatch(sequencerSetBytes, blockTimestamp)
	if err != nil {
		return common.Hash{}, false, 0, fmt.Errorf("failed to seal batch: %w", err)
	}
	sealedBatch, found := bc.GetSealedBatch(batchIndex)
	if !found {
		return common.Hash{}, false, 0, fmt.Errorf("sealed batch not found for index %d", batchIndex)
	}
	if batchIndex <= ci.Uint64() {
		// batch already committed, check batch hash
		correct, err := bc.checkBatchHashCorrect(new(big.Int).SetUint64(batchIndex), sealedBatch.Hash)
		if err != nil {
			return common.Hash{}, false, 0, err
		}
		if !correct {
			log.Error("batch hash does not match sealed batch", "batchIndex", batchIndex, "sealedBatchHash", sealedBatch.Hash.String())
			return common.Hash{}, false, 0, fmt.Errorf("batch hash does not match sealed batch")
		}
	}
	batchHash, err := batchHeaderBytes.Hash()
	if err != nil {
		return common.Hash{}, false, 0, err
	}
	return batchHash, reachedExpectedSize, batchIndex, nil
}

// Get gets sealed batch information by batch index
// Returns the sealed batch info and a boolean indicating if the batch was found
func (bc *BatchCache) Get(batchIndex uint64) (*eth.RPCRollupBatch, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	batch, ok := bc.sealedBatches[batchIndex]
	var err error
	if !ok {
		batch, err = bc.batchStorage.LoadSealedBatch(batchIndex)
		if err != nil {
			return nil, err
		}
	}
	return batch, nil
}

// Delete deletes a sealed batch from the cache by batch index
// Returns a boolean indicating if the batch was found and deleted
func (bc *BatchCache) Delete(batchIndex uint64) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	_, exists := bc.sealedBatches[batchIndex]
	if exists {
		delete(bc.sealedBatches, batchIndex)
	}
	_, headerExists := bc.sealedBatchHeaders[batchIndex]
	if headerExists {
		delete(bc.sealedBatchHeaders, batchIndex)
	}
	err := bc.batchStorage.DeleteSealedBatch(batchIndex)
	if err != nil {
		return err
	}
	return nil
}

// logSealedBatch logs the details of the sealed batch for debugging purposes.
func (bc *BatchCache) logSealedBatch(batchHeader BatchHeaderBytes, batchHash common.Hash) {
	log.Info("Sealed batch header", "batchHash", batchHash.Hex())
	batchIndex, _ := batchHeader.BatchIndex()
	l1MessagePopped, _ := batchHeader.L1MessagePopped()
	totalL1MessagePopped, _ := batchHeader.TotalL1MessagePopped()
	dataHash, _ := batchHeader.DataHash()
	parentBatchHash, _ := batchHeader.ParentBatchHash()
	log.Info(fmt.Sprintf("===batchIndex: %d \n===L1MessagePopped: %d \n===TotalL1MessagePopped: %d \n===dataHash: %x \n===blockCount: %d \n===ParentBatchHash: %x \n",
		batchIndex,
		l1MessagePopped,
		totalL1MessagePopped,
		dataHash,
		bc.batchData.BlockNum(),
		parentBatchHash))
}

func (bc *BatchCache) AssembleCurrentBatchHeader() error {
	if !bc.initDone {
		return errors.New("batch has not been initialized, should wait")
	}
	callOpts := &bind.CallOpts{
		Context: bc.ctx,
	}
	endBlockNum, err := bc.l2Clients.BlockNumber(bc.ctx)
	if err != nil {
		return err
	}
	if endBlockNum < bc.currentBlockNumber {
		return fmt.Errorf("has reorg, should check block status current %v, now %v", bc.currentBlockNumber, endBlockNum)
	}
	startBlockNum := uint64(0)
	version, _ := bc.parentBatchHeader.Version()
	if version < 1 {
		parentIndex, err := bc.parentBatchHeader.BatchIndex()
		if err != nil {
			log.Error("failed to get block index", "err", err)
			return err
		}
		store, err := bc.rollupContract.BatchDataStore(nil, new(big.Int).SetUint64(parentIndex))
		if err != nil {
			log.Error("failed to get batch store", "err", err)
			return err
		}
		startBlockNum = store.BlockNumber.Uint64()
	} else {
		startBlockNum, err = bc.parentBatchHeader.LastBlockNumber()
		if err != nil {
			log.Error("failed to get block number", "err", err)
			return err
		}
	}
	currentBlockNum := bc.currentBlockNumber
	if currentBlockNum < startBlockNum {
		log.Error("invalid block number", "currentBlockNum", currentBlockNum, "startBlockNum", startBlockNum)
		return fmt.Errorf("invalid block number")
	}
	startBlockNum++
	// Get start block once to avoid repeated queries
	startBlock, err := bc.l2Clients.BlockByNumber(bc.ctx, big.NewInt(int64(startBlockNum)))
	if err != nil {
		return fmt.Errorf("failed to get start block %d: %w", startBlockNum, err)
	}
	startBlockTime := startBlock.Time()

	// Fetch blocks from L2 client in the specified range and accumulate to batch
	for blockNum := currentBlockNum + 1; blockNum <= endBlockNum; blockNum++ {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNum)
		root, err := bc.l2Caller.GetTreeRoot(callOpts)
		if err != nil {
			return fmt.Errorf("failed to get withdraw root at block %d: %w", blockNum, err)
		}

		// Check capacity and store to current
		exceeded, err := bc.CalculateCapWithProposalBlock(blockNum, root)
		if err != nil {
			return fmt.Errorf("failed to calculate cap with block %d: %w", blockNum, err)
		}

		// Get the current block to check timeout after packing
		nowBlock, err := bc.l2Clients.BlockByNumber(bc.ctx, big.NewInt(int64(blockNum)))
		if err != nil {
			return fmt.Errorf("failed to get block %d: %w", blockNum, err)
		}
		nowBlockTime := nowBlock.Time()

		// Check timeout: if elapsed time >= batchTimeOut, must seal batch immediately
		// This ensures batch is sealed before exceeding the maximum timeout configured in gov contract
		timeout := false
		if bc.batchTimeOut > 0 {
			elapsedTime := nowBlockTime - startBlockTime
			if elapsedTime >= bc.batchTimeOut {
				timeout = true
				log.Info("Batch timeout reached, must seal batch", "startBlock", startBlockNum, "currentBlock", blockNum,
					"elapsedTime", elapsedTime, "batchTimeOut", bc.batchTimeOut)
			}
		}

		// Check if we need to seal batch due to capacity, block interval, or timeout
		// check ensures batch is sealed before exceeding the maximum timeout
		if exceeded || (bc.blockInterval > 0 && (blockNum-startBlockNum+1) == bc.blockInterval) || timeout {
			log.Info("block exceeds limit", "start", startBlockNum, "to", blockNum, "exceeded", exceeded, "timeout", timeout)
			sequencerSetBytes, _, err := bc.l2Caller.GetSequencerSetBytes(callOpts)
			if err != nil {
				return fmt.Errorf("failed to get sequencer set verify hash at block %d: %w", callOpts.BlockNumber.Uint64(), err)
			}
			lastBlock, err := bc.l2Clients.BlockByNumber(context.Background(), big.NewInt(int64(bc.lastPackedBlockHeight)))
			if err != nil {
				return fmt.Errorf("failed to get last block %d: %w", bc.lastPackedBlockHeight, err)
			}
			blockTimestamp := lastBlock.Time()
			batchIndex, _, _, err := bc.SealBatch(sequencerSetBytes, blockTimestamp)
			if err != nil {
				return fmt.Errorf("failed to seal batch: %w", err)
			}
			batch, ok := bc.GetSealedBatch(batchIndex)
			if !ok {
				return fmt.Errorf("batch %d not found in cache", batchIndex)
			}
			startBlockNum = batch.LastBlockNumber + 1
			startBlock, err = bc.l2Clients.BlockByNumber(bc.ctx, big.NewInt(int64(startBlockNum)))
			if err != nil {
				return fmt.Errorf("failed to get start block %d: %w", startBlockNum, err)
			}
			startBlockTime = startBlock.Time()
		}

		// Pack current block (confirm and append to batch)
		if err = bc.PackCurrentBlock(blockNum); err != nil {
			return fmt.Errorf("failed to pack block %d: %w", blockNum, err)
		}
	}
	return nil
}

func (bc *BatchCache) DeleteBatchStorageAndInitFromRollup() error {
	// should delete invalid batch data and batch header bytes
	err := bc.batchStorage.DeleteAllSealedBatches()
	if err != nil {
		return err
	}
	// batch not contiguous or batch is invalid
	return bc.InitAndSyncFromRollup()
}
