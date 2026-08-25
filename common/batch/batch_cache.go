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
	"sync/atomic"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/common/blob"
)

// BatchCache holds sealed and in-progress rollup batches: it syncs from L1/L2 or local DB,
// packs consecutive L2 blocks into a chunk, seals with blob sidecars, and exposes query/delete APIs.
type BatchCache struct {
	mu       sync.RWMutex
	initMu   sync.Mutex
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

	// Function to determine if batch is upgraded (V0 -> V1)
	isBatchUpgraded func(uint64) bool
	// Function to determine if batch is V2 upgraded (V1 -> V2, multi-blob)
	isBatchV2Upgraded func(uint64) bool

	// Clients and contracts
	l1Client       L1HeaderClient
	l2Clients      L2MultiClient
	rollupContract RollupBatchReader
	l2Gov          L2GovCaller
	// finalizedBatchHeaderLoader is an optional recovery seam used by tests and
	// alternate canonical header sources. Production leaves it nil and loads the
	// header from the L1 finalize transaction.
	finalizedBatchHeaderLoader func(uint64) (*BatchHeaderBytes, error)

	// config
	batchTimeOut  uint64
	blockInterval uint64
	maxBlobCount  int

	// replayL1CommittedBatches is true while InitAndSyncFromRollup is rebuilding committed batches from L2.
	replayL1CommittedBatches atomic.Bool
	// legacyCutoverBatchIndex is the last committed batch whose header contains
	// the historical sequencer-set hash. It is snapshotted from L1 at startup and
	// is consulted only by committed-batch replay; live sealing always writes zero.
	legacyCutoverBatchIndex uint64
}

type batchPackProgressState struct {
	lastLoggedOverallPercent uint64
}

const batchProgressLogStepPercent uint64 = 20

// replayProtocolMaxBlobs is the EIP-4844 per-transaction blob ceiling used when re-sealing
// batches already committed on L1 (independent of max_blob_count flag).
const replayProtocolMaxBlobs = 6

// ErrLegacyTransitionCacheRequired is retained for API compatibility with older
// callers. Startup no longer returns it: a missing or invalid transition cache
// is rebuilt from canonical L1/L2 history.
var ErrLegacyTransitionCacheRequired = errors.New("canonical legacy transition cache required")

// ErrBatchCacheNotInitialized prevents callers from reading persisted or
// in-memory batches before startup validation has completed. In particular, a
// failed legacy transition validation must never leak an unvalidated batch to
// the commit path.
var ErrBatchCacheNotInitialized = errors.New("batch cache is not initialized")

// NewBatchCache creates and initializes a new BatchCache instance
func NewBatchCache(
	isBatchUpgraded func(uint64) bool,
	isBatchV2Upgraded func(uint64) bool,
	maxBlobCount int,
	blockInterval uint64,
	batchTimeOut uint64,
	l1Client L1HeaderClient,
	l2Clients L2MultiClient,
	rollupContract RollupBatchReader,
	l2Gov L2GovCaller,
	ldb SealedBatchKV,
) *BatchCache {
	if isBatchUpgraded == nil {
		// Default implementation: always returns true (use V1 version)
		isBatchUpgraded = func(uint64) bool { return true }
	}
	if isBatchV2Upgraded == nil {
		// Default: V2 not yet activated
		isBatchV2Upgraded = func(uint64) bool { return false }
	}
	if maxBlobCount <= 0 {
		log.Crit("maxBlobCount must be greater than 0")
	}
	ctx := context.Background()
	_, err := l2Clients.BlockNumber(ctx)
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
		isBatchV2Upgraded:                 isBatchV2Upgraded,
		l1Client:                          l1Client,
		l2Clients:                         l2Clients,
		rollupContract:                    rollupContract,
		l2Gov:                             l2Gov,
		batchStorage:                      NewBatchStorage(ldb),
		blockInterval:                     blockInterval,
		batchTimeOut:                      batchTimeOut,
		maxBlobCount:                      maxBlobCount,
	}
}

func (bc *BatchCache) Init() error {
	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("get batch status from rollup failed err: %w", err)
	}
	cutover, _, err := bc.getLegacyTransitionStatus(ci, fi)
	if err != nil {
		return err
	}
	return bc.initFromRollupSnapshot(ci, fi, cutover)
}

// initFromRollupSnapshot installs the exact finalized header as the in-memory
// replay anchor. The anchor intentionally need not be persisted: restart cache
// validation accepts a canonical window beginning at FI+1 and verifies its
// embedded parent header against the on-chain FI hash.
func (bc *BatchCache) initFromRollupSnapshot(ci, fi *big.Int, cutover uint64) error {
	headerBytes, err := bc.getValidatedFinalizedBatchHeader(fi.Uint64())
	if err != nil {
		return err
	}

	// Initialize BatchCache parent batch information.
	// prevStateRoot should be the parent batch's postStateRoot.
	prevStateRoot, err := headerBytes.PostStateRoot()
	if err != nil {
		return fmt.Errorf("get post state root err: %w", err)
	}
	lastPackedBlockHeight, err := headerBytes.LastBlockNumber()
	if err != nil {
		store, err := bc.rollupContract.BatchDataStore(nil, fi)
		if err != nil {
			return err
		}
		lastPackedBlockHeight = store.BlockNumber.Uint64()
	}
	totalL1MessagePopped, err := headerBytes.TotalL1MessagePopped()
	if err != nil {
		return fmt.Errorf("get total l1 message popped err: %w", err)
	}
	bc.mu.Lock()
	bc.parentBatchHeader = headerBytes
	bc.prevStateRoot = prevStateRoot
	bc.lastPackedBlockHeight = lastPackedBlockHeight
	bc.currentBlockNumber = lastPackedBlockHeight
	bc.totalL1MessagePopped = totalL1MessagePopped
	bc.legacyCutoverBatchIndex = cutover
	bc.mu.Unlock()
	log.Info("Start assemble batch", "start batch", fi.Uint64(), "end batch", ci.Uint64())
	return nil
}

func (bc *BatchCache) getValidatedFinalizedBatchHeader(finalized uint64) (*BatchHeaderBytes, error) {
	var (
		headerBytes *BatchHeaderBytes
		err         error
	)
	if bc.finalizedBatchHeaderLoader != nil {
		headerBytes, err = bc.finalizedBatchHeaderLoader(finalized)
	} else {
		if bc.l1Client == nil {
			return nil, errors.New("cannot load finalized batch header: nil L1 client")
		}
		headerBytes, err = bc.getLastFinalizeBatchHeaderFromRollupByIndex(finalized)
	}
	if err != nil {
		return nil, fmt.Errorf("get last finalize batch header err: %w", err)
	}
	if headerBytes == nil {
		return nil, fmt.Errorf("finalized batch header %d is nil", finalized)
	}
	headerIndex, err := headerBytes.BatchIndex()
	if err != nil {
		return nil, fmt.Errorf("get finalized batch header index: %w", err)
	}
	if headerIndex != finalized {
		return nil, fmt.Errorf("finalized batch anchor index mismatch: expected %d, got %d", finalized, headerIndex)
	}
	headerHash, err := headerBytes.Hash()
	if err != nil {
		return nil, fmt.Errorf("hash finalized batch anchor %d: %w", finalized, err)
	}
	committedHash, err := bc.rollupContract.CommittedBatches(nil, new(big.Int).SetUint64(finalized))
	if err != nil {
		return nil, fmt.Errorf("load committed hash for finalized batch %d: %w", finalized, err)
	}
	if headerHash != common.Hash(committedHash) {
		return nil, fmt.Errorf(
			"finalized batch anchor %d hash mismatch: L1 calldata=%s committed=%s",
			finalized, headerHash, common.Hash(committedHash),
		)
	}
	return headerBytes, nil
}

// resetForRollupReplay removes unvalidated runtime state left by a prior failed
// replay attempt. Persisted state is wiped separately by the recovery path.
func (bc *BatchCache) resetForRollupReplay() {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.initDone = false
	bc.sealedBatches = make(map[uint64]*eth.RPCRollupBatch)
	bc.sealedBatchHeaders = make(map[uint64]*BatchHeaderBytes)
	bc.parentBatchHeader = nil
	bc.prevStateRoot = common.Hash{}
	bc.batchData = NewBatchData()
	bc.totalL1MessagePopped = 0
	bc.postStateRoot = common.Hash{}
	bc.withdrawRoot = common.Hash{}
	bc.lastPackedBlockHeight = 0
	bc.currentBlockNumber = 0
	bc.currentBlockHash = common.Hash{}
	bc.legacyCutoverBatchIndex = 0
	bc.ClearCurrent()
}

func (bc *BatchCache) isInitialized() bool {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return bc.initDone
}

func (bc *BatchCache) setInitialized() {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.initDone = true
}

func (bc *BatchCache) InitFromRollupByRange() error {
	bc.initMu.Lock()
	defer bc.initMu.Unlock()
	if bc.isInitialized() {
		return nil
	}
	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("get batch status from rollup failed err: %w", err)
	}
	_, transition, err := bc.getLegacyTransitionStatus(ci, fi)
	if err != nil {
		return err
	}
	if transition {
		// Range assembly seals by local capacity and uses the live zero-field
		// format. During the transition, committed batches must instead be
		// replayed by their exact L1 batch ranges with historical sequencer data.
		return bc.initAndSyncFromRollupLocked()
	}
	err = bc.Init()
	if err != nil {
		return err
	}
	err = bc.assembleUnFinalizeBatchHeaderFromL2Blocks()
	if err != nil {
		return err
	}
	bc.setInitialized()
	log.Info("Initialized batch cache success")
	return nil
}

func (bc *BatchCache) InitAndSyncFromDatabase() error {
	bc.initMu.Lock()
	defer bc.initMu.Unlock()
	if bc.isInitialized() {
		return nil
	}
	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("get batch status from rollup failed err: %w", err)
	}
	cutover, transition, err := bc.getLegacyTransitionStatus(ci, fi)
	if err != nil {
		return err
	}

	batches, headers, indices, err := bc.batchStorage.LoadAllSealedBatchesAndHeader()
	if err != nil {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf("load sealed batch headers: %w", err),
		)
	}
	if err := bc.batchStorage.ValidateExactKeySet(indices); err != nil {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf("validate sealed batch storage key set: %w", err),
		)
	}

	if len(batches) == 0 {
		log.Info("Empty sealed batch cache; rebuilding from rollup",
			"finalizedBatchIndex", fi.Uint64(),
			"committedBatchIndex", ci.Uint64(),
			"legacyTransition", transition,
			"legacyCutoverBatchIndex", cutover,
		)
		// An absent/empty indices snapshot can hide orphan batch keys. Always
		// prefix-wipe before an empty-cache rebuild so those keys cannot survive.
		if err := bc.batchStorage.ForceDeleteAllSealedBatches(); err != nil {
			return fmt.Errorf("wipe empty/orphaned sealed batch cache: %w", err)
		}
		return bc.initAndSyncFromRollupLocked()
	}
	if len(indices) == 0 {
		return bc.handleInvalidStoredBatchCache(
			transition,
			errors.New("sealed batches exist without an indices snapshot"),
		)
	}
	maxIndex := indices[0]
	for _, idx := range indices {
		if idx > maxIndex {
			maxIndex = idx
		}
	}
	if maxIndex < fi.Uint64() {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf("latest cached batch %d is behind finalized batch %d", maxIndex, fi.Uint64()),
		)
	}
	var finalizedAnchor *BatchHeaderBytes
	if _, hasFinalizedRecord := batches[fi.Uint64()]; !hasFinalizedRecord {
		anchor, err := bc.getValidatedFinalizedBatchHeader(fi.Uint64())
		if err != nil {
			// This is an L1 availability/canonical-data failure, not evidence that
			// the local cache is corrupt. Leave the cache untouched and retry.
			return fmt.Errorf("validate finalized cache anchor: %w", err)
		}
		if err := validateCachedWindowAgainstFinalizedAnchor(batches, headers, fi.Uint64(), ci.Uint64(), anchor); err != nil {
			return bc.handleInvalidStoredBatchCache(transition, err)
		}
		finalizedAnchor = anchor
	}
	if transition {
		if err := validateLegacyTransitionWindow(batches, headers, indices, fi.Uint64(), ci.Uint64(), cutover); err != nil {
			return bc.handleInvalidStoredBatchCache(transition, fmt.Errorf("invalid transition cache: %w", err))
		}
	}
	// Check every persisted committed batch. A replay-built cache starts at FI+1
	// because FI is loaded as an exact L1 anchor and is not written to LevelDB;
	// migrated caches that still contain FI remain valid and are checked too.
	firstCommitted := fi.Uint64() + 1
	if _, exists := batches[fi.Uint64()]; exists {
		firstCommitted = fi.Uint64()
	}
	for i := firstCommitted; i <= ci.Uint64(); i++ {
		batchHash, err := bc.rollupContract.CommittedBatches(nil, new(big.Int).SetUint64(i))
		if err != nil {
			return err
		}
		batchStorage, exist := batches[i]
		if !exist || !bytes.Equal(batchHash[:], batchStorage.Hash.Bytes()) {
			// batch not contiguous or batch is invalid
			return bc.handleInvalidStoredBatchCache(
				transition,
				fmt.Errorf("batch %d missing or does not match committed hash", i),
			)
		}
		if i > cutover {
			header := headers[i]
			if header == nil {
				return bc.handleInvalidStoredBatchCache(
					transition,
					fmt.Errorf("committed post-cutover batch header %d is missing", i),
				)
			}
			sequencerHash, err := header.SequencerSetVerifyHash()
			if err != nil {
				return bc.handleInvalidStoredBatchCache(
					transition,
					fmt.Errorf("read committed batch %d sequencer-set hash: %w", i, err),
				)
			}
			if sequencerHash != (common.Hash{}) {
				return bc.handleInvalidStoredBatchCache(
					transition,
					fmt.Errorf("committed batch %d is beyond cutover %d but has non-zero sequencer-set hash %s", i, cutover, sequencerHash),
				)
			}
		}
	}
	// Records above CI are locally sealed and may be broadcast after startup, so
	// validate every calldata/blob-derived field before accepting them. Since the
	// cutover is never ahead of CI, all such records must carry the new zero
	// sequencer-set field regardless of whether FI has already crossed K.
	for _, index := range indices {
		if index <= ci.Uint64() {
			continue
		}
		batch := batches[index]
		header := headers[index]
		if batch == nil || header == nil {
			return bc.handleInvalidStoredBatchCache(
				transition,
				fmt.Errorf("local batch %d or its header is missing", index),
			)
		}
		if err := validateLegacyTransitionRecord(index, batch, header, true); err != nil {
			return bc.handleInvalidStoredBatchCache(transition, err)
		}
		sequencerHash, err := header.SequencerSetVerifyHash()
		if err != nil {
			return bc.handleInvalidStoredBatchCache(
				transition,
				fmt.Errorf("read local batch %d sequencer-set hash: %w", index, err),
			)
		}
		if sequencerHash != (common.Hash{}) {
			return bc.handleInvalidStoredBatchCache(
				transition,
				fmt.Errorf("local post-cutover batch %d has non-zero sequencer-set hash %s", index, sequencerHash),
			)
		}
	}

	latestHeaderBytes := headers[maxIndex]
	if latestHeaderBytes == nil {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf("latest cached batch header %d is missing", maxIndex),
		)
	}
	latestHeaderIndex, err := latestHeaderBytes.BatchIndex()
	if err != nil {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf("read latest cached batch header index: %w", err),
		)
	}
	if latestHeaderIndex != maxIndex || latestHeaderIndex < fi.Uint64() {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf(
				"latest cached header index %d does not match storage max %d/finalized %d",
				latestHeaderIndex, maxIndex, fi.Uint64(),
			),
		)
	}
	prevStateRoot, err := latestHeaderBytes.PostStateRoot()
	if err != nil {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf("get latest cached post state root: %w", err),
		)
	}
	totalL1MessagePopped, err := latestHeaderBytes.TotalL1MessagePopped()
	if err != nil {
		return bc.handleInvalidStoredBatchCache(
			transition,
			fmt.Errorf("get latest cached total l1 messages: %w", err),
		)
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
			return bc.handleInvalidStoredBatchCache(
				transition,
				fmt.Errorf(
					"latest cached batch index %d is outside rollup range [%d,%d]",
					latestBatchIndex,
					fi.Uint64(),
					ci.Uint64(),
				),
			)
		}
		store, err := bc.rollupContract.BatchDataStore(nil, new(big.Int).SetUint64(latestBatchIndex))
		if err != nil {
			return bc.handleInvalidStoredBatchCache(
				transition,
				fmt.Errorf("load batch data for cached batch %d: %w", latestBatchIndex, err),
			)
		}
		lastPackedBlockHeight = store.BlockNumber.Uint64()
	}
	if err := bc.ensureRecoverySnapshotUnchanged(ci, fi, cutover); err != nil {
		// A concurrent commit/finalize changes which records are canonical/local.
		// Do not wipe a cache that was valid for the original snapshot; retry the
		// complete validation against the new tuple instead.
		return err
	}
	committedTip := headers[ci.Uint64()]
	if committedTip == nil && ci.Cmp(fi) == 0 {
		committedTip = finalizedAnchor
	}
	if err := bc.validateCommittedTip(ci.Uint64(), committedTip); err != nil {
		// Keep the DB intact on an RPC or same-index canonical-view change. The
		// next startup attempt will validate the whole window against that view
		// and rebuild if the change is stable.
		return fmt.Errorf("revalidate cached committed tip: %w", err)
	}
	bc.mu.Lock()
	bc.lastPackedBlockHeight = lastPackedBlockHeight
	bc.sealedBatches = batches
	bc.sealedBatchHeaders = headers
	bc.parentBatchHeader = latestHeaderBytes
	bc.currentBlockNumber = bc.lastPackedBlockHeight
	bc.prevStateRoot = prevStateRoot
	bc.totalL1MessagePopped = totalL1MessagePopped
	bc.legacyCutoverBatchIndex = cutover

	bc.initDone = true
	bc.mu.Unlock()
	log.Info(
		"Sync sealed batch from database success",
		"count", len(batches),
		"legacyTransition", transition,
		"legacyCutoverBatchIndex", cutover,
	)
	return nil
}

func validateCachedWindowAgainstFinalizedAnchor(
	batches map[uint64]*eth.RPCRollupBatch,
	headers map[uint64]*BatchHeaderBytes,
	finalized uint64,
	committed uint64,
	anchor *BatchHeaderBytes,
) error {
	if anchor == nil {
		return errors.New("finalized batch anchor is nil")
	}
	anchorIndex, err := anchor.BatchIndex()
	if err != nil {
		return fmt.Errorf("read finalized batch anchor index: %w", err)
	}
	if anchorIndex != finalized {
		return fmt.Errorf("finalized anchor index mismatch: expected %d, got %d", finalized, anchorIndex)
	}
	anchorHash, err := anchor.Hash()
	if err != nil {
		return fmt.Errorf("hash finalized batch anchor %d: %w", finalized, err)
	}

	childIndex := finalized + 1
	child := batches[childIndex]
	if child == nil {
		return fmt.Errorf(
			"non-empty cache without finalized batch %d must begin at canonical child %d (committed=%d)",
			finalized, childIndex, committed,
		)
	}
	if !bytes.Equal(child.ParentBatchHeader, anchor.Bytes()) {
		return fmt.Errorf("batch %d does not embed the exact finalized L1 header %d", childIndex, finalized)
	}
	childHeader := headers[childIndex]
	if childHeader == nil {
		return fmt.Errorf("batch header %d is missing", childIndex)
	}
	childParentHash, err := childHeader.ParentBatchHash()
	if err != nil {
		return fmt.Errorf("read batch %d parent hash: %w", childIndex, err)
	}
	if childParentHash != anchorHash {
		return fmt.Errorf(
			"batch %d parent hash mismatch: header=%s finalized-anchor=%s",
			childIndex, childParentHash, anchorHash,
		)
	}
	return nil
}

func (bc *BatchCache) InitAndSyncFromRollup() error {
	bc.initMu.Lock()
	defer bc.initMu.Unlock()
	return bc.initAndSyncFromRollupLocked()
}

func (bc *BatchCache) initAndSyncFromRollupLocked() (retErr error) {
	if bc.isInitialized() {
		return nil
	}
	bc.replayL1CommittedBatches.Store(true)
	defer bc.replayL1CommittedBatches.Store(false)

	ci, fi, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("get batch status from rollup failed err: %w", err)
	}
	cutover, transition, err := bc.getLegacyTransitionStatus(ci, fi)
	if err != nil {
		return err
	}
	bc.resetForRollupReplay()
	defer func() {
		if retErr == nil {
			return
		}
		// A replay publishes only batches whose index and L1 hash have already
		// been verified, but an incomplete prefix is not a valid startup cache.
		// Remove it from both runtime and storage before the next retry.
		bc.resetForRollupReplay()
		if err := bc.batchStorage.ForceDeleteAllSealedBatches(); err != nil {
			retErr = fmt.Errorf("%v; cleanup failed: %w", retErr, err)
		}
	}()
	if err := bc.initFromRollupSnapshot(ci, fi, cutover); err != nil {
		return err
	}
	log.Info("Start assemble batch",
		"startBatch", fi.Uint64()+1,
		"endBatch", ci.Uint64(),
		"startNum", bc.lastPackedBlockHeight,
		"prevStateRoot", bc.prevStateRoot.String(),
		"legacyTransition", transition,
		"legacyCutoverBatchIndex", cutover,
	)
	nextBatch := fi.Uint64() + 1
	targetCI := new(big.Int).Set(ci)
	observedFI := new(big.Int).Set(fi)
	for {
		for i := nextBatch; i <= targetCI.Uint64(); i++ {
			batchIndex := new(big.Int).SetUint64(i)
			startNum, endNum, err := bc.getBatchBlockRange(batchIndex)
			if err != nil {
				return fmt.Errorf("get batch block range err: %w,start %v, end %v", err, startNum, endNum)
			}
			log.Info("assemble batch block range", "startNum", startNum, "endNum", endNum)
			replayIdx := i
			header, err := bc.assembleBatchHeaderFromL2Blocks(startNum, endNum, &replayIdx)
			if err != nil {
				return err
			}
			// sealBatch has already checked generated index and committedBatches
			// before publishing. Log the exact format branch and reconciliation.
			sequencerHash, err := header.SequencerSetVerifyHash()
			if err != nil {
				return fmt.Errorf("read replay batch %d sequencer-set hash: %w", i, err)
			}
			batchHash, err := header.Hash()
			if err != nil {
				return fmt.Errorf("hash replay batch %d: %w", i, err)
			}
			log.Info("Replayed committed batch",
				"batchIndex", i,
				"legacySequencerSet", i <= cutover,
				"sequencerSetVerifyHash", sequencerHash,
				"generatedBatchHash", batchHash,
				"committedBatchHash", batchHash,
				"targetCommittedBatchIndex", targetCI,
			)
		}

		latestCI, latestFI, err := bc.getBatchStatusFromContract()
		if err != nil {
			return fmt.Errorf("recheck batch status after replay: %w", err)
		}
		latestCutover, _, err := bc.getLegacyTransitionStatus(latestCI, latestFI)
		if err != nil {
			return fmt.Errorf("recheck legacy cutover after replay: %w", err)
		}
		if latestCutover != cutover || latestCI.Cmp(targetCI) < 0 || latestFI.Cmp(observedFI) < 0 {
			return fmt.Errorf(
				"rollup recovery state regressed or cutover changed: previous=(FI=%s CI=%s K=%d) latest=(FI=%s CI=%s K=%d)",
				observedFI, targetCI, cutover, latestFI, latestCI, latestCutover,
			)
		}
		if latestCI.Cmp(targetCI) == 0 {
			bc.mu.RLock()
			committedTip := bc.parentBatchHeader
			bc.mu.RUnlock()
			if err := bc.validateCommittedTip(targetCI.Uint64(), committedTip); err != nil {
				return fmt.Errorf("revalidate replayed committed tip: %w", err)
			}
			break
		}
		log.Info("Committed batch index advanced during recovery; continuing incremental replay",
			"fromBatch", targetCI.Uint64()+1,
			"toBatch", latestCI.Uint64(),
			"latestFinalizedBatchIndex", latestFI,
		)
		nextBatch = targetCI.Uint64() + 1
		targetCI.Set(latestCI)
		observedFI.Set(latestFI)
	}
	bc.setInitialized()
	log.Info("Initialized batch cache success")
	return nil
}

func (bc *BatchCache) ensureRecoverySnapshotUnchanged(ci, fi *big.Int, cutover uint64) error {
	latestCI, latestFI, err := bc.getBatchStatusFromContract()
	if err != nil {
		return fmt.Errorf("recheck batch status after recovery: %w", err)
	}
	latestCutover, _, err := bc.getLegacyTransitionStatus(latestCI, latestFI)
	if err != nil {
		return fmt.Errorf("recheck legacy cutover after recovery: %w", err)
	}
	if latestCI.Cmp(ci) != 0 || latestFI.Cmp(fi) != 0 || latestCutover != cutover {
		return fmt.Errorf(
			"rollup recovery snapshot changed: start=(FI=%s CI=%s K=%d) end=(FI=%s CI=%s K=%d)",
			fi, ci, cutover, latestFI, latestCI, latestCutover,
		)
	}
	return nil
}

// validateCommittedTip closes the same-index reorg/revert window left by an
// FI/CI/K tuple check. Since each batch hash commits to its parent, matching the
// current L1 tip authenticates the complete reconstructed prefix.
func (bc *BatchCache) validateCommittedTip(index uint64, header *BatchHeaderBytes) error {
	if header == nil {
		return fmt.Errorf("committed tip header %d is missing", index)
	}
	headerIndex, err := header.BatchIndex()
	if err != nil {
		return fmt.Errorf("read committed tip header index: %w", err)
	}
	if headerIndex != index {
		return fmt.Errorf("committed tip index mismatch: expected %d, got %d", index, headerIndex)
	}
	headerHash, err := header.Hash()
	if err != nil {
		return fmt.Errorf("hash committed tip %d: %w", index, err)
	}
	committedHash, err := bc.rollupContract.CommittedBatches(nil, new(big.Int).SetUint64(index))
	if err != nil {
		return fmt.Errorf("load current committed tip %d: %w", index, err)
	}
	if headerHash != common.Hash(committedHash) {
		return fmt.Errorf(
			"committed tip %d changed without an index change: cached=%s current=%s",
			index, headerHash, common.Hash(committedHash),
		)
	}
	return nil
}

func (bc *BatchCache) LatestBatchIndex() (uint64, error) {
	return bc.parentBatchHeader.BatchIndex()
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
	if latestCommitBatchIndex == nil || lastFinalizedBatchIndex == nil {
		return nil, nil, errors.New("rollup returned nil batch status")
	}
	if latestCommitBatchIndex.Sign() < 0 || lastFinalizedBatchIndex.Sign() < 0 {
		return nil, nil, fmt.Errorf(
			"rollup returned negative batch status: committed=%s finalized=%s",
			latestCommitBatchIndex, lastFinalizedBatchIndex,
		)
	}
	if lastFinalizedBatchIndex.Cmp(latestCommitBatchIndex) > 0 {
		return nil, nil, fmt.Errorf(
			"last finalized batch %s is ahead of last committed batch %s",
			lastFinalizedBatchIndex, latestCommitBatchIndex,
		)
	}
	return latestCommitBatchIndex, lastFinalizedBatchIndex, nil
}

func (bc *BatchCache) getLegacyTransitionStatus(ci, fi *big.Int) (uint64, bool, error) {
	cutover, err := bc.rollupContract.LegacyCutoverBatchIndex(nil)
	if err != nil {
		return 0, false, fmt.Errorf("get legacy cutover batch index: %w", err)
	}
	if cutover == nil || cutover.Sign() < 0 {
		return 0, false, errors.New("invalid legacy cutover batch index")
	}
	if cutover.Cmp(ci) > 0 {
		return 0, false, fmt.Errorf(
			"legacy cutover batch index %s is ahead of last committed batch %s",
			cutover,
			ci,
		)
	}
	return cutover.Uint64(), fi.Cmp(cutover) < 0, nil
}

func (bc *BatchCache) handleInvalidStoredBatchCache(_ bool, cause error) error {
	log.Error("Invalid sealed batch cache; rebuilding from rollup", "error", cause)
	return bc.deleteBatchStorageAndInitFromRollupLocked()
}

// validateLegacyTransitionWindow applies the additional invariants needed while
// a pre-upgrade batch remains unfinalized. The finalized batch itself is an L1
// anchor and may be absent from LevelDB; in that case the persisted canonical
// window must begin at FI+1, whose embedded parent header is validated below.
func validateLegacyTransitionWindow(
	batches map[uint64]*eth.RPCRollupBatch,
	headers map[uint64]*BatchHeaderBytes,
	indices []uint64,
	finalized uint64,
	committed uint64,
	cutover uint64,
) error {
	if finalized >= cutover {
		return fmt.Errorf("legacy transition is not active: finalized=%d cutover=%d", finalized, cutover)
	}
	if cutover > committed {
		return fmt.Errorf("legacy cutover %d is ahead of committed batch %d", cutover, committed)
	}

	canonicalStart := finalized
	if batches[finalized] == nil {
		canonicalStart = finalized + 1
	}
	for idx := canonicalStart; ; idx++ {
		batch := batches[idx]
		if batch == nil {
			return fmt.Errorf("canonical batch %d is missing", idx)
		}
		header := headers[idx]
		if header == nil {
			return fmt.Errorf("canonical batch header %d is missing", idx)
		}
		if err := validateLegacyTransitionRecord(idx, batch, header, false); err != nil {
			return err
		}
		if idx == committed {
			break
		}
	}

	if headers[cutover] == nil {
		return fmt.Errorf("legacy cutover header %d is missing", cutover)
	}

	for _, idx := range indices {
		if idx <= cutover {
			continue
		}
		batch := batches[idx]
		if batch == nil {
			return fmt.Errorf("post-cutover batch %d is missing", idx)
		}
		header := headers[idx]
		if header == nil {
			return fmt.Errorf("post-cutover batch header %d is missing", idx)
		}
		if err := validateLegacyTransitionRecord(idx, batch, header, idx > committed); err != nil {
			return err
		}
		sequencerHash, err := header.SequencerSetVerifyHash()
		if err != nil {
			return fmt.Errorf("read batch %d sequencer-set hash: %w", idx, err)
		}
		if sequencerHash != (common.Hash{}) {
			return fmt.Errorf(
				"batch %d is beyond legacy cutover %d but has non-zero sequencer-set hash %s",
				idx,
				cutover,
				sequencerHash,
			)
		}
	}
	return nil
}

func validateLegacyTransitionRecord(
	index uint64,
	batch *eth.RPCRollupBatch,
	header *BatchHeaderBytes,
	validateBroadcastFields bool,
) error {
	headerIndex, err := header.BatchIndex()
	if err != nil {
		return fmt.Errorf("read batch header %d index: %w", index, err)
	}
	if headerIndex != index {
		return fmt.Errorf("batch header key %d contains index %d", index, headerIndex)
	}

	version, err := header.Version()
	if err != nil {
		return fmt.Errorf("read batch header %d version: %w", index, err)
	}
	if batch.Version != uint(version) {
		return fmt.Errorf("batch %d version mismatch: record=%d header=%d", index, batch.Version, version)
	}

	// Batch zero is the genesis header and has no preceding batch header.
	if index > 0 {
		parentHeader := BatchHeaderBytes(batch.ParentBatchHeader)
		parentHash, err := parentHeader.Hash()
		if err != nil {
			return fmt.Errorf("hash batch %d parent header: %w", index, err)
		}
		headerParentHash, err := header.ParentBatchHash()
		if err != nil {
			return fmt.Errorf("read batch header %d parent hash: %w", index, err)
		}
		if headerParentHash != parentHash {
			return fmt.Errorf(
				"batch header %d parent hash mismatch: record=%s header=%s",
				index,
				parentHash,
				headerParentHash,
			)
		}
		parentIndex, err := parentHeader.BatchIndex()
		if err != nil {
			return fmt.Errorf("read batch %d parent index: %w", index, err)
		}
		if parentIndex != index-1 {
			return fmt.Errorf("batch %d parent header contains index %d", index, parentIndex)
		}
	}

	prevStateRoot, err := header.PrevStateRoot()
	if err != nil {
		return fmt.Errorf("read batch header %d previous state root: %w", index, err)
	}
	postStateRoot, err := header.PostStateRoot()
	if err != nil {
		return fmt.Errorf("read batch header %d post state root: %w", index, err)
	}
	withdrawalRoot, err := header.WithdrawalRoot()
	if err != nil {
		return fmt.Errorf("read batch header %d withdrawal root: %w", index, err)
	}
	if batch.PrevStateRoot != prevStateRoot || batch.PostStateRoot != postStateRoot || batch.WithdrawRoot != withdrawalRoot {
		return fmt.Errorf("batch %d state roots do not match its header", index)
	}

	if version >= BatchHeaderVersion1 {
		lastBlock, err := header.LastBlockNumber()
		if err != nil {
			return fmt.Errorf("read batch header %d last block: %w", index, err)
		}
		if batch.LastBlockNumber != lastBlock {
			return fmt.Errorf(
				"batch %d last block mismatch: record=%d header=%d",
				index,
				batch.LastBlockNumber,
				lastBlock,
			)
		}
	}

	// Committed historical records are authenticated by their on-chain header
	// hash and may come from old databases that did not retain a complete
	// sidecar. Only locally sealed records beyond currentCommitted can be
	// broadcast in the future, so validate all calldata/blob-derived fields for
	// those records before accepting them as a restart parent.
	if validateBroadcastFields {
		l1MessagePopped, err := header.L1MessagePopped()
		if err != nil {
			return fmt.Errorf("read batch header %d L1 message count: %w", index, err)
		}
		if uint64(batch.NumL1Messages) != l1MessagePopped {
			return fmt.Errorf(
				"batch %d L1 message count mismatch: record=%d header=%d",
				index,
				batch.NumL1Messages,
				l1MessagePopped,
			)
		}

		parentHeader := BatchHeaderBytes(batch.ParentBatchHeader)
		parentTotal, err := parentHeader.TotalL1MessagePopped()
		if err != nil {
			return fmt.Errorf("read batch %d parent total L1 messages: %w", index, err)
		}
		total, err := header.TotalL1MessagePopped()
		if err != nil {
			return fmt.Errorf("read batch header %d total L1 messages: %w", index, err)
		}
		messageCount := uint64(batch.NumL1Messages)
		if parentTotal > ^uint64(0)-messageCount || parentTotal+messageCount != total {
			return fmt.Errorf(
				"batch %d cumulative L1 message mismatch: parent=%d batch=%d header=%d",
				index,
				parentTotal,
				messageCount,
				total,
			)
		}

		if len(batch.Sidecar.Blobs) == 0 || len(batch.Sidecar.Blobs) != len(batch.Sidecar.Commitments) {
			return fmt.Errorf(
				"batch %d invalid sidecar cardinality: blobs=%d commitments=%d",
				index,
				len(batch.Sidecar.Blobs),
				len(batch.Sidecar.Commitments),
			)
		}
		for i := range batch.Sidecar.Blobs {
			commitment, err := kzg4844.BlobToCommitment(&batch.Sidecar.Blobs[i])
			if err != nil {
				return fmt.Errorf("batch %d blob %d is invalid: %w", index, i, err)
			}
			if commitment != batch.Sidecar.Commitments[i] {
				return fmt.Errorf("batch %d blob %d does not match its stored commitment", index, i)
			}
		}
		blobHashes := blob.BlobHashes(batch.Sidecar.Blobs, batch.Sidecar.Commitments)
		var sidecarCommitHash common.Hash
		if version >= BatchHeaderVersion2 {
			sidecarCommitHash = aggregateBlobHashes(blobHashes)
		} else {
			if len(blobHashes) != 1 {
				return fmt.Errorf("batch %d header version %d requires exactly one blob, got %d", index, version, len(blobHashes))
			}
			sidecarCommitHash = blobHashes[0]
		}
		headerCommitHash, err := header.BlobCommitHash()
		if err != nil {
			return fmt.Errorf("read batch header %d blob commitment: %w", index, err)
		}
		if sidecarCommitHash != headerCommitHash {
			return fmt.Errorf(
				"batch %d sidecar commitment mismatch: sidecar=%s header=%s",
				index,
				sidecarCommitHash,
				headerCommitHash,
			)
		}
	}
	return nil
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
	if !bc.initDone {
		return nil, false
	}
	header, ok := bc.sealedBatchHeaders[batchIndex]
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
	if bc.l2Clients.Len() == 0 {
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
	txsPayload, l1TxHashes, newTotalL1MessagePopped, l2TxNum, err := ParsingTxs(block.Transactions(), bc.totalL1MessagePopped)
	if err != nil {
		return false, fmt.Errorf("failed to parse transactions: %w", err)
	}

	l1TxNum := int(newTotalL1MessagePopped - bc.totalL1MessagePopped)
	txsNum := l2TxNum + l1TxNum

	// Build BlockContext (60 bytes)
	blockContext := BuildBlockContext(header, txsNum, l1TxNum)

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
	effectiveBlobCount := bc.effectiveMaxBlobCount(header.Time)
	log.Debug("batch capacity check",
		"proposedBlock", blockNumber,
		"blockTime", header.Time,
		"compressedLimitBytes", effectiveBlobCount*blob.MaxBlobBytesSize,
		"effectiveBlobCount", effectiveBlobCount,
		"configuredMaxBlobCount", bc.maxBlobCount,
		"v2Upgraded", bc.isBatchV2Upgraded(header.Time),
	)
	var exceeded bool
	if bc.isBatchUpgraded(header.Time) {
		exceeded, err = bc.batchData.WillExceedCompressedSizeLimit(blockContext, txsPayload, effectiveBlobCount)
	} else {
		exceeded, err = bc.batchData.EstimateCompressedSizeWithNewPayload(txsPayload, effectiveBlobCount)
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
//   - blockTimestamp: current block timestamp (used to determine batch version)
//
// Returns:
//   - batchIndex: sealed batch index
//   - batchHash: batch hash
//   - reachedExpectedSize: whether the sealed data size reaches expected value (compressed payload size close to or reaches MaxBlobBytesSize)
//   - error: returns error if sealing fails
//
// Note: Sealed batch will be stored in BatchCache's sealedBatches, not sent anywhere
//
// replayCommittedBatchIndex, when non-nil, is the rollup batch index being re-sealed while syncing
// from L1 (InitAndSyncFromRollup). After V2 multi-blob, blob capacity is capped at replayProtocolMaxBlobs
// (6), not max_blob_count, without querying L1 CommitBatch logs.
func (bc *BatchCache) SealBatch(blockTimestamp uint64, replayCommittedBatchIndex *uint64) (uint64, BatchHeaderBytes, bool, error) {
	var (
		sequencerSetBytes []byte
		sequencerSetHash  common.Hash
	)
	if replayCommittedBatchIndex != nil {
		bc.mu.RLock()
		lastBlock := bc.lastPackedBlockHeight
		bc.mu.RUnlock()
		var err error
		sequencerSetBytes, sequencerSetHash, err = bc.getReplaySequencerSet(&bind.CallOpts{
			Context:     bc.ctx,
			BlockNumber: new(big.Int).SetUint64(lastBlock),
		}, replayCommittedBatchIndex)
		if err != nil {
			return 0, BatchHeaderBytes{}, false, err
		}
	}
	return bc.sealBatch(blockTimestamp, replayCommittedBatchIndex, sequencerSetBytes, sequencerSetHash)
}

// sealBatch is the common sealing implementation. sequencerSetBytes/hash are
// accepted only for replay of canonical batches at or before the legacy
// cutover. The live SealBatch path (nil replay index) always supplies an empty
// set and therefore always produces a zero sequencer-set field.
func (bc *BatchCache) sealBatch(
	blockTimestamp uint64,
	replayCommittedBatchIndex *uint64,
	sequencerSetBytes []byte,
	sequencerSetVerifyHash common.Hash,
) (uint64, BatchHeaderBytes, bool, error) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// Ensure batch data is not empty
	if bc.batchData == nil || bc.batchData.IsEmpty() {
		return 0, BatchHeaderBytes{}, false, errors.New("failed to seal batch: batch cache is empty")
	}

	sealBlobCap := bc.sealEffectiveBlobCount(blockTimestamp, replayCommittedBatchIndex)

	// Compress data and calculate dataHash
	compressedPayload, batchDataHash, sealBlobCap, err := bc.handleBatchSealing(blockTimestamp, sealBlobCap, replayCommittedBatchIndex)
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to handle batch sealing: %w", err)
	}

	// Check if sealed data size reaches expected value
	// Expected value: compressed payload size close to or reaches total blob capacity
	// Use 90% as threshold, i.e., if compressed size >= totalCapacity * 0.9, consider it reached expected
	effectiveBlobCount := sealBlobCap
	totalBlobCapacity := effectiveBlobCount * blob.MaxBlobBytesSize
	threshold := float64(totalBlobCapacity) * 0.9
	expectedSizeThreshold := uint64(threshold)
	reachedExpectedSize := uint64(len(compressedPayload)) >= expectedSizeThreshold

	// Generate blob sidecar
	sidecar, err := blob.MakeBlobTxSidecar(compressedPayload, effectiveBlobCount)
	if err != nil {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to create blob sidecar: %w", err)
	}
	log.Info("Sealing batch payload stats",
		"compressedPayloadBytes", len(compressedPayload),
		"effectiveBlobCount", effectiveBlobCount,
		"configuredMaxBlobCount", bc.maxBlobCount,
		"replayCommittedBatchIndex", replayCommittedBatchIndex,
		"v2Upgraded", bc.isBatchV2Upgraded(blockTimestamp),
		"sidecarBlobCount", len(sidecar.Blobs),
		"sidecarCapacityBytes", effectiveBlobCount*blob.MaxBlobBytesSize,
	)

	if replayCommittedBatchIndex == nil {
		if len(sequencerSetBytes) != 0 || sequencerSetVerifyHash != (common.Hash{}) {
			return 0, BatchHeaderBytes{}, false, errors.New("live batch cannot contain a sequencer-set hash")
		}
	} else if *replayCommittedBatchIndex <= bc.legacyCutoverBatchIndex {
		calculated := crypto.Keccak256Hash(sequencerSetBytes)
		if len(sequencerSetBytes) == 0 || calculated != sequencerSetVerifyHash {
			return 0, BatchHeaderBytes{}, false, fmt.Errorf(
				"replay batch %d has invalid legacy sequencer set: bytes=%d contract=%s calculated=%s",
				*replayCommittedBatchIndex, len(sequencerSetBytes), sequencerSetVerifyHash, calculated,
			)
		}
	} else if len(sequencerSetBytes) != 0 || sequencerSetVerifyHash != (common.Hash{}) {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf(
			"replay batch %d is beyond legacy cutover %d but has a sequencer-set hash",
			*replayCommittedBatchIndex, bc.legacyCutoverBatchIndex,
		)
	}

	// Create batch header. Only canonical legacy replay can supply a non-zero
	// sequencer-set hash; normal live sealing reaches this function with zero.
	batchHeader := bc.createBatchHeaderWithSequencerSetHash(
		batchDataHash,
		sidecar,
		sequencerSetVerifyHash,
		blockTimestamp,
	)

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
	if replayCommittedBatchIndex != nil && batchIndex != *replayCommittedBatchIndex {
		return 0, BatchHeaderBytes{}, false, fmt.Errorf(
			"replay batch index mismatch: expected %d, generated %d",
			*replayCommittedBatchIndex, batchIndex,
		)
	}
	if replayCommittedBatchIndex != nil {
		correct, err := bc.checkBatchHashCorrect(new(big.Int).SetUint64(batchIndex), batchHash)
		if err != nil {
			return 0, BatchHeaderBytes{}, false, fmt.Errorf("check replay batch %d hash: %w", batchIndex, err)
		}
		if !correct {
			return 0, BatchHeaderBytes{}, false, fmt.Errorf(
				"replay batch hash mismatch: batch index %d, generated %s",
				batchIndex, batchHash,
			)
		}
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

	// Get L1 message count from batch data
	numL1Messages := bc.batchData.l1TxNum

	// Store sealed batch information as RPCRollupBatch
	sealedBatch := &eth.RPCRollupBatch{
		Version:                  uint(version),
		Hash:                     batchHash,
		ParentBatchHeader:        parentBatchHeaderBytes,
		BlockContexts:            blockContexts,
		CurrentSequencerSetBytes: hexutil.Bytes(sequencerSetBytes),
		PrevStateRoot:            bc.prevStateRoot,
		PostStateRoot:            bc.postStateRoot,
		WithdrawRoot:             bc.withdrawRoot,
		LastBlockNumber:          bc.lastPackedBlockHeight,
		NumL1Messages:            numL1Messages,
		Sidecar:                  *sidecar,
	}
	bc.sealedBatches[batchIndex] = sealedBatch
	// Store batch header copy
	batchHeaderCopy := make(BatchHeaderBytes, len(batchHeader))
	copy(batchHeaderCopy, batchHeader)
	bc.sealedBatchHeaders[batchIndex] = &batchHeaderCopy

	// Persist batch data, header and indices in one atomic write so the stored
	// snapshot can never be partially updated.
	err = bc.batchStorage.StoreSealedBatchAndHeader(batchIndex, sealedBatch, &batchHeaderCopy)
	if err != nil {
		log.Error("failed to store sealed batch and header", "batch_index", batchIndex, "err", err)
		delete(bc.sealedBatches, batchIndex)
		delete(bc.sealedBatchHeaders, batchIndex)
		return 0, BatchHeaderBytes{}, false, fmt.Errorf("failed to store sealed batch and header for batch %d: %w", batchIndex, err)
	}
	// Update parent batch information for next batch
	bc.parentBatchHeader = &batchHeaderCopy
	bc.prevStateRoot = bc.postStateRoot

	// Save block count before resetting batch data for logging
	blockCount := bc.batchData.BlockNum()
	bc.logSealedBatch(batchHeader, batchHash, blockCount, len(sidecar.Blobs))

	// Reset currently accumulated batch data
	bc.batchData = NewBatchData()

	return batchIndex, batchHeader, reachedExpectedSize, nil
}

// handleBatchSealing determines which version to use for compression and calculates data hash.
// The returned sealBlobCap may be raised during L1 batch replay so the compressed payload fits.
func (bc *BatchCache) handleBatchSealing(blockTimestamp uint64, sealBlobCap int, replayCommittedBatchIndex *uint64) ([]byte, common.Hash, int, error) {
	var (
		compressedPayload []byte
		batchDataHash     common.Hash
		err               error
	)

	// Check if upgraded version should be used
	if bc.isBatchUpgraded(blockTimestamp) {
		compressedPayload, err = blob.CompressBatchBytes(bc.batchData.TxsPayloadV2())
		if err != nil {
			return nil, common.Hash{}, sealBlobCap, fmt.Errorf("failed to compress upgraded payload: %w", err)
		}

		replayRaise := replayCommittedBatchIndex != nil || bc.replayL1CommittedBatches.Load()

		if replayRaise {
			needed := (len(compressedPayload) + blob.MaxBlobBytesSize - 1) / blob.MaxBlobBytesSize
			if needed > replayProtocolMaxBlobs {
				if replayCommittedBatchIndex != nil {
					return nil, common.Hash{}, sealBlobCap, fmt.Errorf(
						"replay batch %d: compressed payload needs %d blobs (protocol max %d)",
						*replayCommittedBatchIndex, needed, replayProtocolMaxBlobs,
					)
				}
				return nil, common.Hash{}, sealBlobCap, fmt.Errorf(
					"replay (L1 sync): compressed payload needs %d blobs (protocol max %d)", needed, replayProtocolMaxBlobs,
				)
			}
			if needed > sealBlobCap {
				if replayCommittedBatchIndex != nil {
					log.Info("replay: raising seal blob cap to fit compressed V2 payload",
						"batchIndex", *replayCommittedBatchIndex,
						"fromBlobs", sealBlobCap, "toBlobs", needed,
						"compressedBytes", len(compressedPayload))
				} else {
					log.Info("replay: raising seal blob cap to fit compressed V2 payload",
						"fromBlobs", sealBlobCap, "toBlobs", needed,
						"compressedBytes", len(compressedPayload))
				}
				sealBlobCap = needed
			}
		}

		if len(compressedPayload) <= sealBlobCap*blob.MaxBlobBytesSize {
			batchDataHash, err = bc.batchData.DataHashV2()
			if err != nil {
				return nil, common.Hash{}, sealBlobCap, fmt.Errorf("failed to calculate upgraded data hash: %w", err)
			}
			return compressedPayload, batchDataHash, sealBlobCap, nil
		}
		if bc.isBatchV2Upgraded(blockTimestamp) {
			return nil, common.Hash{}, sealBlobCap, fmt.Errorf(
				"compressed V2 batch size %d exceeds capacity for %d blobs (%d bytes)",
				len(compressedPayload), sealBlobCap, sealBlobCap*blob.MaxBlobBytesSize,
			)
		}
	}

	// Fall back to the old version
	compressedPayload, err = blob.CompressBatchBytes(bc.batchData.TxsPayload())
	if err != nil {
		return nil, common.Hash{}, sealBlobCap, fmt.Errorf("failed to compress payload: %w", err)
	}

	replayRaise := replayCommittedBatchIndex != nil || bc.replayL1CommittedBatches.Load()

	if replayRaise {
		needed := (len(compressedPayload) + blob.MaxBlobBytesSize - 1) / blob.MaxBlobBytesSize
		if needed > replayProtocolMaxBlobs {
			if replayCommittedBatchIndex != nil {
				return nil, common.Hash{}, sealBlobCap, fmt.Errorf(
					"replay batch %d: legacy compressed payload needs %d blobs (protocol max %d)",
					*replayCommittedBatchIndex, needed, replayProtocolMaxBlobs,
				)
			}
			return nil, common.Hash{}, sealBlobCap, fmt.Errorf(
				"replay (L1 sync): legacy compressed payload needs %d blobs (protocol max %d)", needed, replayProtocolMaxBlobs,
			)
		}
		if needed > sealBlobCap {
			if replayCommittedBatchIndex != nil {
				log.Info("replay: raising seal blob cap to fit legacy compressed payload",
					"batchIndex", *replayCommittedBatchIndex,
					"fromBlobs", sealBlobCap, "toBlobs", needed,
					"compressedBytes", len(compressedPayload))
			} else {
				log.Info("replay: raising seal blob cap to fit legacy compressed payload",
					"fromBlobs", sealBlobCap, "toBlobs", needed,
					"compressedBytes", len(compressedPayload))
			}
			sealBlobCap = needed
		}
	}

	if len(compressedPayload) > sealBlobCap*blob.MaxBlobBytesSize {
		return nil, common.Hash{}, sealBlobCap, fmt.Errorf(
			"compressed batch size %d exceeds capacity for %d blobs (%d bytes)",
			len(compressedPayload), sealBlobCap, sealBlobCap*blob.MaxBlobBytesSize,
		)
	}

	batchDataHash = bc.batchData.DataHash()

	return compressedPayload, batchDataHash, sealBlobCap, nil
}

// createBatchHeader creates BatchHeader
func (bc *BatchCache) createBatchHeader(dataHash common.Hash, sidecar *ethtypes.BlobTxSidecar, blockTimestamp uint64) BatchHeaderBytes {
	return bc.createBatchHeaderWithSequencerSetHash(dataHash, sidecar, common.Hash{}, blockTimestamp)
}

func (bc *BatchCache) createBatchHeaderWithSequencerSetHash(
	dataHash common.Hash,
	sidecar *ethtypes.BlobTxSidecar,
	sequencerSetVerifyHash common.Hash,
	blockTimestamp uint64,
) BatchHeaderBytes {
	blobHashes := []common.Hash{blob.EmptyVersionedHash}
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
		// V2 is activated: use V1-format header (257 bytes) with version byte 2.
		// Store keccak256(concat all blob hashes) at offset 57 as the aggregated blob hash.
		if bc.isBatchV2Upgraded(blockTimestamp) {
			batchHeaderV1.BlobVersionedHash = aggregateBlobHashes(blobHashes)
			h := batchHeaderV1.Bytes()
			h[0] = BatchHeaderVersion2
			return h
		}
		return batchHeaderV1.Bytes()
	}

	return batchHeaderV0.Bytes()
}

// ParsingTxs encodes a block's transactions into the on-chain payload format
// used by the batch builder: L2 transactions are RLP-marshalled and concatenated
// in order; L1 message transactions are excluded from the payload but their
// hashes and queue indices are tracked separately.
//
// Exported for derivation local verify (SPEC-005), which must rebuild blob bytes from
// local L2 blocks using the same encoding the sequencer applied at seal time.
func ParsingTxs(transactions []*ethtypes.Transaction, totalL1MessagePoppedBefore uint64) (
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

			if currentIndex != nextIndex {
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

// aggregateBlobHashes computes keccak256 of the concatenation of all blob hash bytes.
func aggregateBlobHashes(hashes []common.Hash) common.Hash {
	var concat []byte
	for _, h := range hashes {
		concat = append(concat, h[:]...)
	}
	return crypto.Keccak256Hash(concat)
}

// effectiveMaxBlobCount returns the allowed blob count for the given block timestamp.
// V2 multi-blob is only permitted when isBatchV2Upgraded returns true; otherwise cap at 1.
func (bc *BatchCache) effectiveMaxBlobCount(blockTimestamp uint64) int {
	if bc.isBatchV2Upgraded(blockTimestamp) {
		return bc.maxBlobCount
	}
	return 1
}

// sealEffectiveBlobCount is the blob count used for sealing.
// Live packing uses effectiveMaxBlobCount (max_blob_count flag).
// Replaying an L1-committed batch after V2 multi-blob uses replayProtocolMaxBlobs (6), independent of
// max_blob_count and without L1 log queries; handleBatchSealing tightens from compressed size (still ≤6).
func (bc *BatchCache) sealEffectiveBlobCount(blockTimestamp uint64, replayCommittedBatchIndex *uint64) int {
	base := bc.effectiveMaxBlobCount(blockTimestamp)
	if replayCommittedBatchIndex == nil {
		return base
	}
	if !bc.isBatchV2Upgraded(blockTimestamp) {
		return base
	}
	return replayProtocolMaxBlobs
}

// BuildBlockContext serialises a block header + tx counts into the 60-byte
// BlockContext blob the batch builder writes for each block.
// Format: Number(8) || Timestamp(8) || BaseFee(32) || GasLimit(8) || numTxs(2) || numL1Messages(2)
//
// Exported for derivation local verify (SPEC-005); see ParsingTxs.
func BuildBlockContext(header *ethtypes.Header, txsNum, l1MsgNum int) []byte {
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
	replayCommittedBatchIndex *uint64,
) (*BatchHeaderBytes, error) {
	// Fresh accumulation for this chain batch; a failed prior SealBatch must not double-pack blocks.
	bc.mu.Lock()
	bc.batchData = NewBatchData()
	bc.ClearCurrent()
	bc.mu.Unlock()

	ctx := context.Background()
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	// Fetch blocks from L2 client in the specified range and accumulate to batch
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNum)
		root, err := bc.l2Gov.GetTreeRoot(callOpts)
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

	var (
		sequencerSetBytes      []byte
		sequencerSetVerifyHash common.Hash
	)
	sequencerSetBytes, sequencerSetVerifyHash, err := bc.getReplaySequencerSet(callOpts, replayCommittedBatchIndex)
	if err != nil {
		return nil, err
	}

	// Get the last block's timestamp for packing
	lastBlock, err := bc.l2Clients.BlockByNumber(ctx, big.NewInt(int64(endBlockNum)))
	if err != nil {
		return nil, fmt.Errorf("failed to get last block %d: %w", endBlockNum, err)
	}
	blockTimestamp := lastBlock.Time()

	// Seal batch and generate batchHeader
	batchIndex, batchHeader, reachedExpectedSize, err := bc.sealBatch(
		blockTimestamp,
		replayCommittedBatchIndex,
		sequencerSetBytes,
		sequencerSetVerifyHash,
	)
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

func (bc *BatchCache) getReplaySequencerSet(
	callOpts *bind.CallOpts,
	replayCommittedBatchIndex *uint64,
) ([]byte, common.Hash, error) {
	if replayCommittedBatchIndex == nil || *replayCommittedBatchIndex > bc.legacyCutoverBatchIndex {
		return nil, common.Hash{}, nil
	}
	if bc.l2Gov == nil {
		return nil, common.Hash{}, fmt.Errorf(
			"replay batch %d requires historical L2 sequencer state, but L2 gov caller is nil",
			*replayCommittedBatchIndex,
		)
	}
	sequencerSetBytes, sequencerSetVerifyHash, err := bc.l2Gov.GetSequencerSetBytes(callOpts)
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf(
			"failed to get sequencer set at block %s for replay batch %d: %w",
			callOpts.BlockNumber, *replayCommittedBatchIndex, err,
		)
	}
	calculated := crypto.Keccak256Hash(sequencerSetBytes)
	if calculated != sequencerSetVerifyHash {
		return nil, common.Hash{}, fmt.Errorf(
			"sequencer set hash mismatch at block %s for replay batch %d: contract=%s calculated=%s",
			callOpts.BlockNumber, *replayCommittedBatchIndex, sequencerSetVerifyHash, calculated,
		)
	}
	return sequencerSetBytes, sequencerSetVerifyHash, nil
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
	progressState := batchPackProgressState{}

	// Fetch blocks from L2 client in the specified range and accumulate to batch
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNum)
		root, err := bc.l2Gov.GetTreeRoot(callOpts)
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
		bc.logBatchPackingProgress(startBlockNum, blockNum, startBlockTime, nowBlockTime, &progressState)

		// Check timeout: if elapsed time >= batchTimeOut, must seal batch immediately
		// This ensures batch is sealed before exceeding the configured maximum timeout.
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
			batchHash, reachedExpectedSize, batchIndex, err := bc.SealBatchAndCheck(ci)
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
			progressState = batchPackProgressState{}
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

func (bc *BatchCache) SealBatchAndCheck(ci *big.Int) (common.Hash, bool, uint64, error) {
	lastBlock, err := bc.l2Clients.BlockByNumber(context.Background(), big.NewInt(int64(bc.lastPackedBlockHeight)))
	if err != nil {
		return common.Hash{}, false, 0, fmt.Errorf("failed to get last block %d: %w", bc.lastPackedBlockHeight, err)
	}
	blockTimestamp := lastBlock.Time()
	// Seal batch and generate batchHeader
	batchIndex, batchHeaderBytes, reachedExpectedSize, err := bc.SealBatch(blockTimestamp, nil)
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
	if !bc.initDone {
		return nil, ErrBatchCacheNotInitialized
	}
	batch, ok := bc.sealedBatches[batchIndex]
	if !ok {
		return nil, fmt.Errorf("sealed batch %d not found in validated cache: %w", batchIndex, ErrKeyNotFound)
	}
	return batch, nil
}

// Delete removes a single sealed batch (data + header) from the cache and storage.
//
// Must NOT be used for finalize cleanup: it delegates to BatchStorage.DeleteSealedBatch,
// which rejects interior-index deletes because they punch holes into the persisted
// indices and crash the next restart load. Finalize cleanup must use DeleteUntil.
func (bc *BatchCache) Delete(batchIndex uint64) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	if err := bc.batchStorage.DeleteSealedBatch(batchIndex); err != nil {
		return err
	}
	_, exists := bc.sealedBatches[batchIndex]
	if exists {
		delete(bc.sealedBatches, batchIndex)
	}
	_, headerExists := bc.sealedBatchHeaders[batchIndex]
	if headerExists {
		delete(bc.sealedBatchHeaders, batchIndex)
	}
	return nil
}

// DeleteUntil removes every sealed batch and header with index <= maxIndex from
// both the in-memory maps and persistent storage. Finalize cleanup must use this
// range-based form: the finalize target can jump (multiple submitters, several
// batches finalized at once), so deleting a single index leaves stale lower
// indices behind and punches holes into the persisted indices snapshot, which
// breaks the contiguity assumption of the startup load path.
func (bc *BatchCache) DeleteUntil(maxIndex uint64) error {
	bc.initMu.Lock()
	defer bc.initMu.Unlock()
	if err := bc.batchStorage.DeleteSealedBatchesUpTo(maxIndex); err != nil {
		return err
	}

	bc.mu.Lock()
	defer bc.mu.Unlock()
	for idx := range bc.sealedBatches {
		if idx <= maxIndex {
			delete(bc.sealedBatches, idx)
		}
	}
	for idx := range bc.sealedBatchHeaders {
		if idx <= maxIndex {
			delete(bc.sealedBatchHeaders, idx)
		}
	}
	return nil
}

// logSealedBatch logs the details of the sealed batch for debugging purposes.
func (bc *BatchCache) logSealedBatch(batchHeader BatchHeaderBytes, batchHash common.Hash, blockCount uint16, blobCount int) {
	version, err := batchHeader.Version()
	if err != nil {
		version = 0
	}
	blobCommitHash, blobErr := batchHeader.BlobCommitHash()
	if blobErr != nil {
		log.Warn("Sealed batch: blob commit hash unavailable", "batchHash", batchHash.Hex(), "version", version, "err", blobErr)
	}
	log.Info("Sealed batch header",
		"batchHash", batchHash.Hex(),
		"version", version,
		"blobCommitHash", blobCommitHash.Hex(),
		"blobCount", blobCount,
	)
	batchIndex, _ := batchHeader.BatchIndex()
	l1MessagePopped, _ := batchHeader.L1MessagePopped()
	totalL1MessagePopped, _ := batchHeader.TotalL1MessagePopped()
	dataHash, _ := batchHeader.DataHash()
	parentBatchHash, _ := batchHeader.ParentBatchHash()
	blobFieldLabel := "BlobVersionedHash"
	if version >= BatchHeaderVersion2 {
		blobFieldLabel = "BlobHashesHash"
	}
	log.Info(fmt.Sprintf("===version: %d \n===batchIndex: %d \n===L1MessagePopped: %d \n===TotalL1MessagePopped: %d \n===dataHash: %x \n===%s: %x \n===blockCount: %d \n===blobCount: %d \n===ParentBatchHash: %x \n",
		version,
		batchIndex,
		l1MessagePopped,
		totalL1MessagePopped,
		dataHash,
		blobFieldLabel,
		blobCommitHash,
		blockCount,
		blobCount,
		parentBatchHash))
}

func (bc *BatchCache) AssembleCurrentBatchHeader() error {
	if !bc.isInitialized() {
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
	if bc.parentBatchHeader == nil {
		return fmt.Errorf("parent batch header is nil, cannot assemble batch")
	}
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
	progressState := batchPackProgressState{}

	// Fetch blocks from L2 client in the specified range and accumulate to batch
	for blockNum := currentBlockNum + 1; blockNum <= endBlockNum; blockNum++ {
		callOpts.BlockNumber = new(big.Int).SetUint64(blockNum)
		root, err := bc.l2Gov.GetTreeRoot(callOpts)
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
		bc.logBatchPackingProgress(startBlockNum, blockNum, startBlockTime, nowBlockTime, &progressState)

		// Check timeout: if elapsed time >= batchTimeOut, must seal batch immediately
		// This ensures batch is sealed before exceeding the configured maximum timeout.
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
			lastBlock, err := bc.l2Clients.BlockByNumber(context.Background(), big.NewInt(int64(bc.lastPackedBlockHeight)))
			if err != nil {
				return fmt.Errorf("failed to get last block %d: %w", bc.lastPackedBlockHeight, err)
			}
			blockTimestamp := lastBlock.Time()
			batchIndex, _, _, err := bc.SealBatch(blockTimestamp, nil)
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
			progressState = batchPackProgressState{}
		}

		// Pack current block (confirm and append to batch)
		if err = bc.PackCurrentBlock(blockNum); err != nil {
			return fmt.Errorf("failed to pack block %d: %w", blockNum, err)
		}
	}
	return nil
}

func (bc *BatchCache) logBatchPackingProgress(startBlockNum, currentBlockNum, startBlockTime, currentBlockTime uint64, state *batchPackProgressState) {
	if state == nil || currentBlockNum < startBlockNum {
		return
	}

	elapsedTime := uint64(0)
	if currentBlockTime >= startBlockTime {
		elapsedTime = currentBlockTime - startBlockTime
	}

	packedBlocks := currentBlockNum - startBlockNum + 1
	effectiveBlobCount := bc.effectiveMaxBlobCount(currentBlockTime)
	totalBlobCapacity := uint64(effectiveBlobCount * blob.MaxBlobBytesSize)
	payloadBytes := uint64(0)
	if totalBlobCapacity > 0 {
		payloadBytes = bc.estimatedBatchPayloadBytesWithCurrent(currentBlockTime)
	}

	timePercent := uint64(0)
	if bc.batchTimeOut > 0 {
		timePercent = progressPercent(elapsedTime, bc.batchTimeOut)
	}

	blockPercent := uint64(0)
	if bc.blockInterval > 0 {
		blockPercent = progressPercent(packedBlocks, bc.blockInterval)
	}

	blobPercent := uint64(0)
	if totalBlobCapacity > 0 {
		blobPercent = progressPercent(payloadBytes, totalBlobCapacity)
	}

	overallPercent := maxUint64(timePercent, blockPercent, blobPercent)
	// Throttle progress logs to reduce noisy output.
	overallStep := (overallPercent / batchProgressLogStepPercent) * batchProgressLogStepPercent
	if overallStep <= state.lastLoggedOverallPercent {
		return
	}
	state.lastLoggedOverallPercent = overallStep

	log.Info("Batch packing progress",
		"loadedBlockHeight", currentBlockNum,
		"overallPercent", overallPercent,
		"timePercent", timePercent,
		"blockPercent", blockPercent,
		"blobPercent", blobPercent,
	)
}

func (bc *BatchCache) estimatedBatchPayloadBytesWithCurrent(blockTimestamp uint64) uint64 {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	var (
		existingBlockContextLen int
		existingTxPayloadLen    int
	)
	if bc.batchData != nil {
		existingBlockContextLen = len(bc.batchData.blockContexts)
		existingTxPayloadLen = len(bc.batchData.txsPayload)
	}

	if bc.isBatchUpgraded(blockTimestamp) {
		return uint64(existingBlockContextLen + len(bc.currentBlockContext) + existingTxPayloadLen + len(bc.currentTxsPayload))
	}

	return uint64(existingTxPayloadLen + len(bc.currentTxsPayload))
}

func progressPercent(current, total uint64) uint64 {
	if total == 0 {
		return 0
	}
	p := current * 100 / total
	if p > 100 {
		return 100
	}
	return p
}

func maxUint64(values ...uint64) uint64 {
	var max uint64
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func (bc *BatchCache) DeleteBatchStorageAndInitFromRollup() error {
	bc.initMu.Lock()
	defer bc.initMu.Unlock()
	return bc.deleteBatchStorageAndInitFromRollupLocked()
}

func (bc *BatchCache) deleteBatchStorageAndInitFromRollupLocked() error {
	// Recovery always scans the prefix directly. A parseable indices snapshot
	// can still omit orphan data/header keys, which DeleteAllSealedBatches would
	// leave behind and could re-surface on a later repair.
	if err := bc.batchStorage.ForceDeleteAllSealedBatches(); err != nil {
		return fmt.Errorf("force wipe sealed batches failed: %w", err)
	}
	// Rebuild the complete committed window from the exact finalized L1 anchor.
	// Runtime state is reset inside initAndSyncFromRollupLocked so a failed prior
	// attempt cannot become the parent of this replay.
	return bc.initAndSyncFromRollupLocked()
}
