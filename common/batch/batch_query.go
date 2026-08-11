package batch

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"morph-l2/bindings/bindings"
	"morph-l2/common/blob"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
)

type canonicalCommitEvent struct {
	batchIndex uint64
	batchHash  common.Hash
	raw        ethtypes.Log
}

type canonicalRevertEvent struct {
	batchIndex uint64
	batchHash  common.Hash
	raw        ethtypes.Log
}

func logIdentityLess(a, b ethtypes.Log) bool {
	if a.BlockNumber != b.BlockNumber {
		return a.BlockNumber < b.BlockNumber
	}
	if a.TxIndex != b.TxIndex {
		return a.TxIndex < b.TxIndex
	}
	return a.Index < b.Index
}

// selectCanonicalCommit follows the shared resolver contract from spec 9.3:
// discard commits at/before the last revert and select the last remaining log
// whose hash equals committedBatches[index] at the fixed snapshot.
func selectCanonicalCommit(
	index uint64,
	expectedHash common.Hash,
	commits []canonicalCommitEvent,
	reverts []canonicalRevertEvent,
) (*canonicalCommitEvent, error) {
	if expectedHash == (common.Hash{}) {
		return nil, fmt.Errorf("batch %d has no canonical committed hash at snapshot", index)
	}
	sort.Slice(commits, func(i, j int) bool { return logIdentityLess(commits[i].raw, commits[j].raw) })
	sort.Slice(reverts, func(i, j int) bool { return logIdentityLess(reverts[i].raw, reverts[j].raw) })

	var lastRevert *canonicalRevertEvent
	for i := range reverts {
		if reverts[i].batchIndex == index {
			lastRevert = &reverts[i]
		}
	}
	var selected *canonicalCommitEvent
	for i := range commits {
		candidate := &commits[i]
		if candidate.batchIndex != index || candidate.batchHash != expectedHash {
			continue
		}
		if lastRevert != nil && !logIdentityLess(lastRevert.raw, candidate.raw) {
			continue
		}
		selected = candidate
	}
	if selected == nil {
		return nil, fmt.Errorf("no canonical CommitBatch log for batch %d and hash %s", index, expectedHash)
	}
	return selected, nil
}

// getCanonicalCommittedBatchHeaderByIndex recovers a complete header from the
// last canonical CommitBatch transaction at the already-fixed L1 snapshot. A
// FinalizeBatch transaction is deliberately never used as a header substitute.
func (bc *BatchCache) getCanonicalCommittedBatchHeaderByIndex(index uint64) (*BatchHeaderBytes, error) {
	if bc.l1Snapshot == nil {
		return nil, errors.New("canonical batch header lookup requires a confirmed L1 snapshot")
	}
	expectedRaw, err := bc.rollupContract.CommittedBatches(bc.snapshotCallOpts(), new(big.Int).SetUint64(index))
	if err != nil {
		return nil, fmt.Errorf("read committed batch %d at snapshot: %w", index, err)
	}
	expectedHash := common.Hash(expectedRaw)

	commits, reverts, err := bc.collectCanonicalBatchEvents(index)
	if err != nil {
		return nil, err
	}
	selected, err := selectCanonicalCommit(index, expectedHash, commits, reverts)
	if err != nil {
		return nil, err
	}
	tx, pending, err := bc.l1Client.TransactionByHash(bc.ctx, selected.raw.TxHash)
	if err != nil {
		return nil, fmt.Errorf("load canonical batch %d transaction %s: %w", index, selected.raw.TxHash, err)
	}
	if tx == nil || pending {
		return nil, fmt.Errorf("canonical batch %d transaction %s is missing or pending", index, selected.raw.TxHash)
	}
	header, err := bc.reconstructCommittedBatchHeader(index, expectedHash, tx)
	if err != nil {
		return nil, fmt.Errorf("reconstruct canonical batch %d from transaction %s: %w", index, selected.raw.TxHash, err)
	}
	return &header, nil
}

func (bc *BatchCache) collectCanonicalBatchEvents(index uint64) ([]canonicalCommitEvent, []canonicalRevertEvent, error) {
	const blockRange uint64 = 10_000
	var commits []canonicalCommitEvent
	var reverts []canonicalRevertEvent
	for start := uint64(0); ; {
		end := start + blockRange - 1
		if end < start || end > bc.l1Snapshot.number {
			end = bc.l1Snapshot.number
		}
		opts, err := bc.snapshotFilterOpts(start, end)
		if err != nil {
			return nil, nil, err
		}
		commitIter, err := bc.rollupContract.FilterCommitBatch(opts, []*big.Int{new(big.Int).SetUint64(index)}, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("filter CommitBatch %d in L1 range [%d,%d]: %w", index, start, end, err)
		}
		for commitIter.Next() {
			event := commitIter.Event
			if event == nil || event.BatchIndex == nil || !event.BatchIndex.IsUint64() {
				_ = commitIter.Close()
				return nil, nil, fmt.Errorf("malformed CommitBatch event for batch %d", index)
			}
			if event.BatchIndex.Uint64() != index || event.Raw.Removed || event.Raw.BlockNumber < start || event.Raw.BlockNumber > end {
				_ = commitIter.Close()
				return nil, nil, fmt.Errorf("non-canonical CommitBatch event returned for batch %d", index)
			}
			commits = append(commits, canonicalCommitEvent{
				batchIndex: event.BatchIndex.Uint64(),
				batchHash:  common.Hash(event.BatchHash),
				raw:        event.Raw,
			})
		}
		iterErr := commitIter.Error()
		closeErr := commitIter.Close()
		if iterErr != nil {
			return nil, nil, fmt.Errorf("iterate CommitBatch %d in L1 range [%d,%d]: %w", index, start, end, iterErr)
		}
		if closeErr != nil {
			return nil, nil, fmt.Errorf("close CommitBatch iterator for batch %d: %w", index, closeErr)
		}

		opts, err = bc.snapshotFilterOpts(start, end)
		if err != nil {
			return nil, nil, err
		}
		revertIter, err := bc.rollupContract.FilterRevertBatch(opts, []*big.Int{new(big.Int).SetUint64(index)}, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("filter RevertBatch %d in L1 range [%d,%d]: %w", index, start, end, err)
		}
		for revertIter.Next() {
			event := revertIter.Event
			if event == nil || event.BatchIndex == nil || !event.BatchIndex.IsUint64() {
				_ = revertIter.Close()
				return nil, nil, fmt.Errorf("malformed RevertBatch event for batch %d", index)
			}
			if event.BatchIndex.Uint64() != index || event.Raw.Removed || event.Raw.BlockNumber < start || event.Raw.BlockNumber > end {
				_ = revertIter.Close()
				return nil, nil, fmt.Errorf("non-canonical RevertBatch event returned for batch %d", index)
			}
			reverts = append(reverts, canonicalRevertEvent{
				batchIndex: event.BatchIndex.Uint64(),
				batchHash:  common.Hash(event.BatchHash),
				raw:        event.Raw,
			})
		}
		iterErr = revertIter.Error()
		closeErr = revertIter.Close()
		if iterErr != nil {
			return nil, nil, fmt.Errorf("iterate RevertBatch %d in L1 range [%d,%d]: %w", index, start, end, iterErr)
		}
		if closeErr != nil {
			return nil, nil, fmt.Errorf("close RevertBatch iterator for batch %d: %w", index, closeErr)
		}

		if end == bc.l1Snapshot.number {
			break
		}
		start = end + 1
	}
	return commits, reverts, nil
}

func (bc *BatchCache) reconstructCommittedBatchHeader(
	index uint64,
	expectedHash common.Hash,
	tx *ethtypes.Transaction,
) (BatchHeaderBytes, error) {
	if tx == nil {
		return nil, errors.New("nil canonical commit transaction")
	}
	genesis, matched, err := decodeImportGenesisHeader(tx.Data())
	if err != nil {
		return nil, err
	}
	if matched {
		if err := verifyHeaderIdentity(genesis, index, expectedHash); err != nil {
			return nil, err
		}
		if err := bc.verifyFinalizedMessageCursor(genesis); err != nil {
			return nil, err
		}
		return genesis, nil
	}

	abis, err := DefaultRollupABIs()
	if err != nil {
		return nil, err
	}
	decoded, err := decodeBatchCall(tx.Data(), abis)
	if err != nil {
		return nil, err
	}
	if decoded.input.Version > BatchHeaderVersion2 {
		return nil, fmt.Errorf("canonical commit has unsupported batch version %d", decoded.input.Version)
	}
	if len(decoded.explicitHeader) > 0 {
		if err := verifyHeaderIdentity(decoded.explicitHeader, index, expectedHash); err != nil {
			return nil, fmt.Errorf("invalid explicit header in %s.%s: %w", decoded.era, decoded.method, err)
		}
		if err := bc.verifyFinalizedMessageCursor(decoded.explicitHeader); err != nil {
			return nil, err
		}
		return decoded.explicitHeader, nil
	}

	parent := BatchHeaderBytes(decoded.input.ParentBatchHeader)
	parentIndex, err := parent.BatchIndex()
	if err != nil {
		return nil, fmt.Errorf("decode parent batch header index: %w", err)
	}
	if parentIndex == ^uint64(0) || parentIndex+1 != index {
		return nil, fmt.Errorf("canonical commit parent index mismatch: parent=%d current=%d", parentIndex, index)
	}
	parentHash, err := parent.Hash()
	if err != nil {
		return nil, fmt.Errorf("hash parent batch header: %w", err)
	}
	parentCursor, err := parent.TotalL1MessagePopped()
	if err != nil {
		return nil, fmt.Errorf("decode parent L1 message cursor: %w", err)
	}

	queue, err := bc.l1MessageQueue()
	if err != nil {
		return nil, err
	}
	var (
		dataHash       common.Hash
		l1MessageCount uint64
	)
	switch decoded.era {
	case "pre-submitter", "current":
		l1MessageCount = uint64(decoded.input.NumL1Messages)
		if l1MessageCount > ^uint64(0)-parentCursor {
			return nil, fmt.Errorf("L1 message cursor overflow: parent=%d batch=%d", parentCursor, l1MessageCount)
		}
		data := make([]byte, 10, 10+32*int(l1MessageCount))
		binary.BigEndian.PutUint64(data[:8], decoded.input.LastBlockNumber)
		binary.BigEndian.PutUint16(data[8:10], decoded.input.NumL1Messages)
		for i := uint64(0); i < l1MessageCount; i++ {
			messageHash, err := queue.GetCrossDomainMessage(
				bc.snapshotCallOpts(),
				new(big.Int).SetUint64(parentCursor+i),
			)
			if err != nil {
				return nil, fmt.Errorf("read L1 message %d for canonical batch %d: %w", parentCursor+i, index, err)
			}
			data = append(data, messageHash[:]...)
		}
		dataHash = crypto.Keccak256Hash(data)
	case "before-move-block-context", "legacy":
		data, count, err := bc.reconstructBlockContextData(decoded, parentCursor, queue)
		if err != nil {
			return nil, err
		}
		l1MessageCount = count
		dataHash = crypto.Keccak256Hash(data)
	default:
		return nil, fmt.Errorf("unsupported canonical commit ABI era %q", decoded.era)
	}
	if l1MessageCount > ^uint64(0)-parentCursor {
		return nil, fmt.Errorf("L1 message cursor overflow: parent=%d batch=%d", parentCursor, l1MessageCount)
	}
	totalCursor := parentCursor + l1MessageCount

	blobHashRaw, err := bc.rollupContract.BatchBlobVersionedHashes(
		bc.snapshotCallOpts(),
		new(big.Int).SetUint64(index),
	)
	if err != nil {
		return nil, fmt.Errorf("read blob hash for canonical batch %d: %w", index, err)
	}
	blobCommitHash, err := canonicalBlobCommitHash(tx, decoded, common.Hash(blobHashRaw))
	if err != nil {
		return nil, err
	}
	base := BatchHeaderV0{
		BatchIndex:             index,
		L1MessagePopped:        l1MessageCount,
		TotalL1MessagePopped:   totalCursor,
		DataHash:               dataHash,
		BlobVersionedHash:      blobCommitHash,
		PrevStateRoot:          common.Hash(decoded.input.PrevStateRoot),
		PostStateRoot:          common.Hash(decoded.input.PostStateRoot),
		WithdrawalRoot:         common.Hash(decoded.input.WithdrawalRoot),
		SequencerSetVerifyHash: decoded.sequencerSetHash,
		ParentBatchHash:        parentHash,
	}
	var header BatchHeaderBytes
	if decoded.input.Version >= BatchHeaderVersion1 {
		header = BatchHeaderV1{
			BatchHeaderV0:   base,
			LastBlockNumber: decoded.input.LastBlockNumber,
		}.Bytes()
		if decoded.input.Version == BatchHeaderVersion2 {
			header[0] = BatchHeaderVersion2
		}
	} else {
		header = base.Bytes()
		if decoded.era == "legacy" {
			header = append(header, decoded.skippedBitmap...)
		}
	}
	if err := verifyHeaderIdentity(header, index, expectedHash); err != nil {
		return nil, fmt.Errorf("reconstructed %s.%s header: %w", decoded.era, decoded.method, err)
	}
	if err := bc.verifyFinalizedMessageCursor(header); err != nil {
		return nil, err
	}
	return header, nil
}

func canonicalBlobCommitHash(
	tx *ethtypes.Transaction,
	decoded *decodedBatchCall,
	stored common.Hash,
) (common.Hash, error) {
	hashes := tx.BlobHashes()
	if len(hashes) == 0 {
		if decoded.era == "current" && decoded.method == "commitBatch" {
			return common.Hash{}, errors.New("current canonical commitBatch has no fresh blob hash")
		}
		if stored != (common.Hash{}) {
			return stored, nil
		}
		if decoded.method == "commitState" || decoded.input.Version >= BatchHeaderVersion2 {
			return common.Hash{}, fmt.Errorf("%s.%s has no recoverable stored blob hash", decoded.era, decoded.method)
		}
		// Frozen V0/V1 commitBatch transactions mined before blobs were
		// available used the protocol's non-zero empty versioned hash.
		return blob.EmptyVersionedHash, nil
	}

	var fromTransaction common.Hash
	if decoded.input.Version >= BatchHeaderVersion2 {
		fromTransaction = aggregateBlobHashes(hashes)
	} else {
		fromTransaction = hashes[0]
	}
	if stored != (common.Hash{}) && stored != fromTransaction {
		return common.Hash{}, fmt.Errorf("canonical batch blob hash mismatch: stored=%s transaction=%s", stored, fromTransaction)
	}
	return fromTransaction, nil
}

func decodeImportGenesisHeader(data []byte) (BatchHeaderBytes, bool, error) {
	rollupABI, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, false, err
	}
	method, ok := rollupABI.Methods["importGenesisBatch"]
	if !ok || len(data) < 4 || !bytes.Equal(data[:4], method.ID) {
		return nil, false, nil
	}
	args, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return nil, true, fmt.Errorf("decode importGenesisBatch: %w", err)
	}
	if len(args) != 1 {
		return nil, true, fmt.Errorf("decode importGenesisBatch: expected one argument, got %d", len(args))
	}
	header, ok := args[0].([]byte)
	if !ok {
		return nil, true, errors.New("decode importGenesisBatch: invalid header argument")
	}
	return BatchHeaderBytes(bytes.Clone(header)), true, nil
}

func verifyHeaderIdentity(header BatchHeaderBytes, index uint64, expectedHash common.Hash) error {
	gotIndex, err := header.BatchIndex()
	if err != nil {
		return fmt.Errorf("decode batch header index: %w", err)
	}
	if gotIndex != index {
		return fmt.Errorf("batch header index mismatch: event=%d header=%d", index, gotIndex)
	}
	gotHash, err := header.Hash()
	if err != nil {
		return fmt.Errorf("hash batch header: %w", err)
	}
	if gotHash != expectedHash {
		return fmt.Errorf("batch header hash mismatch: getter=%s reconstructed=%s", expectedHash, gotHash)
	}
	return nil
}

func (bc *BatchCache) l1MessageQueue() (*bindings.L1MessageQueueWithGasPriceOracleCaller, error) {
	address, err := bc.rollupContract.MessageQueue(bc.snapshotCallOpts())
	if err != nil {
		return nil, fmt.Errorf("read Rollup messageQueue at snapshot: %w", err)
	}
	if address == (common.Address{}) {
		return nil, errors.New("Rollup messageQueue is zero at snapshot")
	}
	queue, err := bindings.NewL1MessageQueueWithGasPriceOracleCaller(address, bc.l1Client)
	if err != nil {
		return nil, fmt.Errorf("bind L1 message queue %s: %w", address, err)
	}
	return queue, nil
}

func (bc *BatchCache) verifyFinalizedMessageCursor(header BatchHeaderBytes) error {
	total, err := header.TotalL1MessagePopped()
	if err != nil {
		return fmt.Errorf("decode finalized L1 message cursor: %w", err)
	}
	queue, err := bc.l1MessageQueue()
	if err != nil {
		return err
	}
	pending, err := queue.PendingQueueIndex(bc.snapshotCallOpts())
	if err != nil {
		return fmt.Errorf("read finalized L1 message queue cursor: %w", err)
	}
	if pending == nil || !pending.IsUint64() || pending.Uint64() != total {
		return fmt.Errorf("finalized L1 message cursor mismatch: header=%d queue=%v", total, pending)
	}
	next, err := queue.NextCrossDomainMessageIndex(bc.snapshotCallOpts())
	if err != nil {
		return fmt.Errorf("read next L1 message queue index: %w", err)
	}
	if next == nil || !next.IsUint64() || next.Uint64() < total {
		return fmt.Errorf("L1 message queue shorter than finalized cursor: header=%d next=%v", total, next)
	}
	return nil
}

func (bc *BatchCache) reconstructBlockContextData(
	decoded *decodedBatchCall,
	parentCursor uint64,
	queue *bindings.L1MessageQueueWithGasPriceOracleCaller,
) ([]byte, uint64, error) {
	contexts := decoded.blockContexts
	if len(contexts) < 2 {
		return nil, 0, fmt.Errorf("%s block contexts are shorter than count prefix", decoded.era)
	}
	blockCount := int(binary.BigEndian.Uint16(contexts[:2]))
	if blockCount == 0 || len(contexts) != 2+blockCount*60 {
		return nil, 0, fmt.Errorf("%s block context length mismatch: blocks=%d bytes=%d", decoded.era, blockCount, len(contexts))
	}
	data := make([]byte, 0, blockCount*58)
	var messageCount uint64
	for i := 0; i < blockCount; i++ {
		contextBytes := contexts[2+i*60 : 2+(i+1)*60]
		data = append(data, contextBytes[:58]...)
		messageCount += uint64(binary.BigEndian.Uint16(contextBytes[58:60]))
	}
	if decoded.era == "legacy" {
		expectedBitmapBytes := int((messageCount + 255) / 256 * 32)
		if len(decoded.skippedBitmap) != expectedBitmapBytes {
			return nil, 0, fmt.Errorf("legacy skipped bitmap length mismatch: messages=%d expected=%d actual=%d", messageCount, expectedBitmapBytes, len(decoded.skippedBitmap))
		}
	}
	if messageCount > ^uint64(0)-parentCursor {
		return nil, 0, fmt.Errorf("historical L1 message cursor overflow: parent=%d batch=%d", parentCursor, messageCount)
	}
	for i := uint64(0); i < messageCount; i++ {
		if decoded.era == "legacy" && legacyMessageSkipped(decoded.skippedBitmap, i) {
			continue
		}
		messageHash, err := queue.GetCrossDomainMessage(
			bc.snapshotCallOpts(),
			new(big.Int).SetUint64(parentCursor+i),
		)
		if err != nil {
			return nil, 0, fmt.Errorf("read historical L1 message %d: %w", parentCursor+i, err)
		}
		data = append(data, messageHash[:]...)
	}
	return data, messageCount, nil
}

func legacyMessageSkipped(bitmap []byte, index uint64) bool {
	wordStart := int(index/256) * 32
	bitInWord := index % 256
	byteIndex := wordStart + 31 - int(bitInWord/8)
	return byteIndex >= 0 && byteIndex < len(bitmap) && bitmap[byteIndex]&(1<<uint(bitInWord%8)) != 0
}

var (
	ErrCalldataTooShort     = errors.New("rollup calldata is shorter than a selector")
	ErrUnknownBatchSelector = errors.New("unknown rollup batch selector")
)

// RollupABIs contains every ABI era needed by historical and current readers.
// Historical ABIs intentionally remain local frozen descriptions: regenerating
// the current binding must never mutate the decoder for already-mined calldata.
type RollupABIs struct {
	BeforeMoveBlockCtx abi.ABI
	Legacy             abi.ABI
	PreSubmitter       abi.ABI
	Current            abi.ABI
}

type batchDataInputStruct struct {
	Version           uint8    `json:"version"`
	ParentBatchHeader []byte   `json:"parentBatchHeader"`
	LastBlockNumber   uint64   `json:"lastBlockNumber"`
	NumL1Messages     uint16   `json:"numL1Messages"`
	PrevStateRoot     [32]byte `json:"prevStateRoot"`
	PostStateRoot     [32]byte `json:"postStateRoot"`
	WithdrawalRoot    [32]byte `json:"withdrawalRoot"`
}

type beforeMoveBatchDataInput struct {
	Version           uint8    `json:"version"`
	ParentBatchHeader []byte   `json:"parentBatchHeader"`
	BlockContexts     []byte   `json:"blockContexts"`
	PrevStateRoot     [32]byte `json:"prevStateRoot"`
	PostStateRoot     [32]byte `json:"postStateRoot"`
	WithdrawalRoot    [32]byte `json:"withdrawalRoot"`
}

type legacyBatchDataInput struct {
	Version                uint8    `json:"version"`
	ParentBatchHeader      []byte   `json:"parentBatchHeader"`
	BlockContexts          []byte   `json:"blockContexts"`
	SkippedL1MessageBitmap []byte   `json:"skippedL1MessageBitmap"`
	PrevStateRoot          [32]byte `json:"prevStateRoot"`
	PostStateRoot          [32]byte `json:"postStateRoot"`
	WithdrawalRoot         [32]byte `json:"withdrawalRoot"`
}

const (
	beforeMoveRollupABIJSON   = `[{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"blockContexts","type":"bytes"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"components":[{"name":"signedSequencersBitmap","type":"uint256"},{"name":"sequencerSets","type":"bytes"},{"name":"signature","type":"bytes"}],"name":"batchSignatureInput","type":"tuple"}],"name":"commitBatch","outputs":[],"stateMutability":"payable","type":"function"}]`
	legacyRollupABIJSON       = `[{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"blockContexts","type":"bytes"},{"name":"skippedL1MessageBitmap","type":"bytes"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"components":[{"name":"signedSequencersBitmap","type":"uint256"},{"name":"sequencerSets","type":"bytes"},{"name":"signature","type":"bytes"}],"name":"batchSignatureInput","type":"tuple"}],"name":"commitBatch","outputs":[],"stateMutability":"payable","type":"function"}]`
	preSubmitterRollupABIJSON = `[{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"components":[{"name":"signedSequencersBitmap","type":"uint256"},{"name":"sequencerSets","type":"bytes"},{"name":"signature","type":"bytes"}],"name":"batchSignatureInput","type":"tuple"}],"name":"commitBatch","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"components":[{"name":"signedSequencersBitmap","type":"uint256"},{"name":"sequencerSets","type":"bytes"},{"name":"signature","type":"bytes"}],"name":"batchSignatureInput","type":"tuple"}],"name":"commitState","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"components":[{"name":"signedSequencersBitmap","type":"uint256"},{"name":"sequencerSets","type":"bytes"},{"name":"signature","type":"bytes"}],"name":"batchSignatureInput","type":"tuple"},{"name":"batchHeader","type":"bytes"},{"name":"batchProof","type":"bytes"}],"name":"commitBatchWithProof","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	currentRollupABIJSON      = `[{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"}],"name":"commitBatch","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"}],"name":"commitState","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"components":[{"name":"version","type":"uint8"},{"name":"parentBatchHeader","type":"bytes"},{"name":"lastBlockNumber","type":"uint64"},{"name":"numL1Messages","type":"uint16"},{"name":"prevStateRoot","type":"bytes32"},{"name":"postStateRoot","type":"bytes32"},{"name":"withdrawalRoot","type":"bytes32"}],"name":"batchDataInput","type":"tuple"},{"name":"batchHeader","type":"bytes"},{"name":"batchProof","type":"bytes"}],"name":"commitBatchWithProof","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
)

func parseRollupABI(raw string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(raw))
}

// DefaultRollupABIs loads the frozen historical ABI fragments and the current
// post-submitter ABI fragment.
func DefaultRollupABIs() (RollupABIs, error) {
	beforeMove, err := parseRollupABI(beforeMoveRollupABIJSON)
	if err != nil {
		return RollupABIs{}, fmt.Errorf("parse before-move ABI: %w", err)
	}
	legacy, err := parseRollupABI(legacyRollupABIJSON)
	if err != nil {
		return RollupABIs{}, fmt.Errorf("parse legacy ABI: %w", err)
	}
	preSubmitter, err := parseRollupABI(preSubmitterRollupABIJSON)
	if err != nil {
		return RollupABIs{}, fmt.Errorf("parse pre-submitter ABI: %w", err)
	}
	currentBinding, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return RollupABIs{}, fmt.Errorf("load current Rollup binding ABI: %w", err)
	}
	if currentBinding == nil {
		return RollupABIs{}, errors.New("load current Rollup binding ABI: nil ABI")
	}
	// Keep the protocol's fixed post-cutover selectors explicit and fail fast if
	// generated bindings drift from the reviewed ABI surface.
	frozenCurrent, err := parseRollupABI(currentRollupABIJSON)
	if err != nil {
		return RollupABIs{}, fmt.Errorf("parse frozen current ABI: %w", err)
	}
	for _, name := range []string{"commitBatch", "commitState", "commitBatchWithProof"} {
		generated, generatedOK := currentBinding.Methods[name]
		frozen, frozenOK := frozenCurrent.Methods[name]
		if !generatedOK || !frozenOK || !bytes.Equal(generated.ID, frozen.ID) {
			return RollupABIs{}, fmt.Errorf("current Rollup binding selector drift for %s", name)
		}
	}
	return RollupABIs{
		BeforeMoveBlockCtx: beforeMove,
		Legacy:             legacy,
		PreSubmitter:       preSubmitter,
		Current:            *currentBinding,
	}, nil
}

type batchCallCandidate struct {
	era     string
	method  abi.Method
	current bool
}

type historicalBatchSignatureInput struct {
	SignedSequencersBitmap *big.Int `json:"signedSequencersBitmap"`
	SequencerSets          []byte   `json:"sequencerSets"`
	Signature              []byte   `json:"signature"`
}

type decodedBatchCall struct {
	era              string
	method           string
	input            *bindings.IRollupBatchDataInput
	blockContexts    []byte
	skippedBitmap    []byte
	sequencerSetHash common.Hash
	explicitHeader   BatchHeaderBytes
}

func supportedBatchCalls(abis RollupABIs) ([]batchCallCandidate, error) {
	var calls []batchCallCandidate
	add := func(era string, contractABI abi.ABI, current bool, methods ...string) error {
		for _, name := range methods {
			method, ok := contractABI.Methods[name]
			if !ok {
				return fmt.Errorf("%s ABI does not contain %s", era, name)
			}
			calls = append(calls, batchCallCandidate{era: era, method: method, current: current})
		}
		return nil
	}
	if err := add("before-move-block-context", abis.BeforeMoveBlockCtx, false, "commitBatch"); err != nil {
		return nil, err
	}
	if err := add("legacy", abis.Legacy, false, "commitBatch"); err != nil {
		return nil, err
	}
	if err := add("pre-submitter", abis.PreSubmitter, true, "commitBatch", "commitState", "commitBatchWithProof"); err != nil {
		return nil, err
	}
	if err := add("current", abis.Current, true, "commitBatch", "commitState", "commitBatchWithProof"); err != nil {
		return nil, err
	}

	seen := make(map[[4]byte]string, len(calls))
	for _, call := range calls {
		var selector [4]byte
		copy(selector[:], call.method.ID)
		name := call.era + "." + call.method.Name
		if previous, ok := seen[selector]; ok {
			return nil, fmt.Errorf("batch selector collision between %s and %s", previous, name)
		}
		seen[selector] = name
	}
	return calls, nil
}

func decodeBatchCall(data []byte, abis RollupABIs) (*decodedBatchCall, error) {
	if len(data) < 4 {
		return nil, ErrCalldataTooShort
	}
	calls, err := supportedBatchCalls(abis)
	if err != nil {
		return nil, err
	}
	for _, call := range calls {
		if !bytes.Equal(data[:4], call.method.ID) {
			continue
		}
		args, err := call.method.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, fmt.Errorf("decode %s.%s: %w", call.era, call.method.Name, err)
		}
		if len(args) == 0 {
			return nil, fmt.Errorf("decode %s.%s: missing BatchDataInput", call.era, call.method.Name)
		}
		result := &decodedBatchCall{era: call.era, method: call.method.Name}
		if call.current {
			result.input, err = convertCurrentBatchDataInput(args[0])
		} else {
			result.input, err = convertHistoricalBatchDataInput(call.era, args[0])
			switch call.era {
			case "before-move-block-context":
				converted, ok := abi.ConvertType(args[0], new(beforeMoveBatchDataInput)).(*beforeMoveBatchDataInput)
				if !ok || converted == nil {
					return nil, errors.New("failed to preserve before-move block contexts")
				}
				result.blockContexts = bytes.Clone(converted.BlockContexts)
			case "legacy":
				converted, ok := abi.ConvertType(args[0], new(legacyBatchDataInput)).(*legacyBatchDataInput)
				if !ok || converted == nil {
					return nil, errors.New("failed to preserve legacy block contexts")
				}
				result.blockContexts = bytes.Clone(converted.BlockContexts)
				result.skippedBitmap = bytes.Clone(converted.SkippedL1MessageBitmap)
			}
		}
		if err != nil {
			return nil, err
		}

		if call.era != "current" {
			if len(args) < 2 {
				return nil, fmt.Errorf("decode %s.%s: missing historical signature tuple", call.era, call.method.Name)
			}
			signature, ok := abi.ConvertType(args[1], new(historicalBatchSignatureInput)).(*historicalBatchSignatureInput)
			if !ok || signature == nil {
				return nil, fmt.Errorf("decode %s.%s: invalid historical signature tuple", call.era, call.method.Name)
			}
			result.sequencerSetHash = crypto.Keccak256Hash(signature.SequencerSets)
		}
		if call.method.Name == "commitBatchWithProof" {
			headerArg := 1
			if call.era != "current" {
				headerArg = 2
			}
			if len(args) <= headerArg {
				return nil, fmt.Errorf("decode %s.%s: missing explicit batch header", call.era, call.method.Name)
			}
			header, ok := args[headerArg].([]byte)
			if !ok {
				return nil, fmt.Errorf("decode %s.%s: invalid explicit batch header type", call.era, call.method.Name)
			}
			result.explicitHeader = BatchHeaderBytes(bytes.Clone(header))
		}
		return result, nil
	}
	return nil, fmt.Errorf("%w: 0x%x", ErrUnknownBatchSelector, data[:4])
}

func convertCurrentBatchDataInput(value interface{}) (*bindings.IRollupBatchDataInput, error) {
	converted, ok := abi.ConvertType(value, new(batchDataInputStruct)).(*batchDataInputStruct)
	if !ok || converted == nil {
		return nil, errors.New("failed to convert current BatchDataInput tuple")
	}
	return &bindings.IRollupBatchDataInput{
		Version:           converted.Version,
		ParentBatchHeader: bytes.Clone(converted.ParentBatchHeader),
		LastBlockNumber:   converted.LastBlockNumber,
		NumL1Messages:     converted.NumL1Messages,
		PrevStateRoot:     converted.PrevStateRoot,
		PostStateRoot:     converted.PostStateRoot,
		WithdrawalRoot:    converted.WithdrawalRoot,
	}, nil
}

func convertHistoricalBatchDataInput(era string, value interface{}) (*bindings.IRollupBatchDataInput, error) {
	result := new(bindings.IRollupBatchDataInput)
	switch era {
	case "before-move-block-context":
		converted, ok := abi.ConvertType(value, new(beforeMoveBatchDataInput)).(*beforeMoveBatchDataInput)
		if !ok || converted == nil {
			return nil, errors.New("failed to convert before-move BatchDataInput tuple")
		}
		result.Version = converted.Version
		result.ParentBatchHeader = bytes.Clone(converted.ParentBatchHeader)
		result.PrevStateRoot = converted.PrevStateRoot
		result.PostStateRoot = converted.PostStateRoot
		result.WithdrawalRoot = converted.WithdrawalRoot
	case "legacy":
		converted, ok := abi.ConvertType(value, new(legacyBatchDataInput)).(*legacyBatchDataInput)
		if !ok || converted == nil {
			return nil, errors.New("failed to convert legacy BatchDataInput tuple")
		}
		result.Version = converted.Version
		result.ParentBatchHeader = bytes.Clone(converted.ParentBatchHeader)
		result.PrevStateRoot = converted.PrevStateRoot
		result.PostStateRoot = converted.PostStateRoot
		result.WithdrawalRoot = converted.WithdrawalRoot
	default:
		return nil, fmt.Errorf("unsupported historical ABI era %q", era)
	}
	return result, nil
}

// DecodeBatchDataInput dispatches strictly by the four-byte selector and only
// exposes the first BatchDataInput argument. Historical signature tuples are
// consumed by their frozen ABI during unpacking but never reach business code.
func DecodeBatchDataInput(data []byte, abis RollupABIs) (*bindings.IRollupBatchDataInput, error) {
	if len(data) < 4 {
		return nil, ErrCalldataTooShort
	}
	calls, err := supportedBatchCalls(abis)
	if err != nil {
		return nil, err
	}
	for _, call := range calls {
		if !bytes.Equal(data[:4], call.method.ID) {
			continue
		}
		args, err := call.method.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, fmt.Errorf("decode %s.%s: %w", call.era, call.method.Name, err)
		}
		if len(args) == 0 {
			return nil, fmt.Errorf("decode %s.%s: missing BatchDataInput", call.era, call.method.Name)
		}
		if call.current {
			return convertCurrentBatchDataInput(args[0])
		}
		return convertHistoricalBatchDataInput(call.era, args[0])
	}
	return nil, fmt.Errorf("%w: 0x%x", ErrUnknownBatchSelector, data[:4])
}

// DecodeBatchDataInputDefault decodes with the repository's frozen ABI set.
func DecodeBatchDataInputDefault(data []byte) (*bindings.IRollupBatchDataInput, error) {
	abis, err := DefaultRollupABIs()
	if err != nil {
		return nil, err
	}
	return DecodeBatchDataInput(data, abis)
}

// parseCommitBatchTxData is retained as a package-local compatibility wrapper
// for restart helpers while callers migrate to DecodeBatchDataInputDefault.
func parseCommitBatchTxData(txData []byte) (*bindings.IRollupBatchDataInput, error) {
	return DecodeBatchDataInputDefault(txData)
}
