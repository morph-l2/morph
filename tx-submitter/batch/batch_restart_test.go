package batch

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/db"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/types"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var (
	ErrBatchNotFound = errors.New("batch not found")
)

var (
	rollupAddr = common.HexToAddress("0x0165878a594ca255338adfa4d48449f69242eb8f")

	l1ClientRpc = "http://localhost:9545"
	l2ClientRpc = "http://localhost:8545"
	l1Client, _ = ethclient.Dial(l1ClientRpc)
	l2Client, _ = ethclient.Dial(l2ClientRpc)

	rollupContract *bindings.Rollup

	l2Caller *types.L2Caller
)

func init() {
	var err error
	rollupContract, err = bindings.NewRollup(rollupAddr, l1Client)
	if err != nil {
		panic(err)
	}
	l2Caller, err = types.NewL2Caller([]iface.L2Client{l2Client})
	if err != nil {
		panic(err)
	}
}

func Test_GetFinalizeBatchHeader(t *testing.T) {
	testDir := filepath.Join(t.TempDir(), "testleveldb")
	os.RemoveAll(testDir)
	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})
	testDB, err := db.New(testDir)
	require.NoError(t, err)

	bc := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)
	headerBytes, err := bc.getLastFinalizeBatchHeaderFromRollupByIndex(0)
	require.NoError(t, err)
	t.Log("headerBytes", hex.EncodeToString(headerBytes.Bytes()))
}

func Test_CommitBatchParse(t *testing.T) {
	data, signature, err := getCommitBatchDataByIndex(5357)
	require.NoError(t, err)
	t.Log("data", data)
	t.Log("signature", signature)
	t.Log("data.Version", data.Version)
	t.Log("data.ParentBatchHeader", hex.EncodeToString(data.ParentBatchHeader))
	t.Log("data.LastBlockNumber", data.LastBlockNumber)
	t.Log("data.NumL1Messages", data.NumL1Messages)
	t.Log("data.PrevStateRoot", hex.EncodeToString(data.PrevStateRoot[:]))
	t.Log("data.PostStateRoot", hex.EncodeToString(data.PostStateRoot[:]))
	t.Log("data.WithdrawalRoot", hex.EncodeToString(data.WithdrawalRoot[:]))
}

func TestBatchRestartInit(t *testing.T) {
	testDir := filepath.Join(t.TempDir(), "testleveldb")
	os.RemoveAll(testDir)
	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})
	testDB, err := db.New(testDir)
	require.NoError(t, err)

	sequencerSetBytes, sequencerSetVerifyHash, err := l2Caller.GetSequencerSetBytes(nil)
	require.NoError(t, err)
	t.Log("sequencer set verify hash", hex.EncodeToString(sequencerSetVerifyHash[:]))
	ci, fi := getInfosFromContract()
	t.Log("commit index", ci, " ", "finalize index", fi)
	bc := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)
	startBlockNum, endBlockNum, err := getFirstUnFinalizeBatchBlockNumRange(fi)
	require.NoError(t, err)
	startBlockNum = new(big.Int).Add(startBlockNum, new(big.Int).SetUint64(1))
	t.Log("start block number", startBlockNum, "end block number", endBlockNum)

	// Get the latest finalized batch header
	headerBytes, err := getLastFinalizeBatchHeaderByIndex(fi.Uint64())
	require.NoError(t, err, "failed to get last finalized batch header")
	parentStateRoot, err := headerBytes.PostStateRoot()
	require.NoError(t, err, "failed to get post state root")

	// Initialize BatchCache parent batch information
	// prevStateRoot should be the parent batch's postStateRoot (i.e., the current finalized batch's postStateRoot)
	bc.parentBatchHeader = headerBytes
	bc.prevStateRoot = parentStateRoot // The current batch's prevStateRoot is the parent batch's postStateRoot
	bc.lastPackedBlockHeight, err = headerBytes.LastBlockNumber()
	if err != nil {
		store, err := rollupContract.BatchDataStore(nil, fi)
		require.NoError(t, err)
		bc.lastPackedBlockHeight = store.BlockNumber.Uint64()
	}
	bc.totalL1MessagePopped, err = headerBytes.TotalL1MessagePopped()
	require.NoError(t, err)
	t.Logf("Restored batch header: batchIndex=%d, parentStateRoot=%x (will be used as prevStateRoot for next batch)",
		fi, parentStateRoot[:])

	// Query the first unfinalized batch's block range from rollup contract
	firstUnfinalizedIndex := fi.Uint64() + 1
	t.Logf("First unfinalize batch index: %d, block range: %d - %d", firstUnfinalizedIndex, startBlockNum.Uint64(), endBlockNum.Uint64())

	// Fetch blocks from L2 client in this range and assemble batchHeader
	assembledBatchHeader, err := assembleBatchHeaderFromL2Blocks(bc, startBlockNum.Uint64(), endBlockNum.Uint64(), sequencerSetBytes, l2Client, l2Caller)
	require.NoError(t, err, "failed to assemble batch header from L2 blocks")
	t.Log("assembled batch header success", hex.EncodeToString(assembledBatchHeader.Bytes()))
	// Verify the assembled batchHeader
	assembledBatchIndex, err := assembledBatchHeader.BatchIndex()
	require.NoError(t, err)
	require.Equal(t, firstUnfinalizedIndex, assembledBatchIndex, "assembled batch index should match")
	assembledBatchHash, err := assembledBatchHeader.Hash()
	require.NoError(t, err)

	batchDataInput, batchSignatureInput, err := getCommitBatchDataByIndex(firstUnfinalizedIndex)
	require.NoError(t, err)
	t.Logf("batchDataInput.Version=%d", batchDataInput.Version)
	require.Equal(t, hex.EncodeToString(batchDataInput.ParentBatchHeader), hex.EncodeToString(headerBytes.Bytes()))
	t.Logf("batchDataInput.LastBlockNumber=%d, %d", batchDataInput.LastBlockNumber, endBlockNum)
	l1MsgNum, err := assembledBatchHeader.L1MessagePopped()
	require.NoError(t, err)
	require.Equal(t, uint64(batchDataInput.NumL1Messages), l1MsgNum)
	prevStateRoot, err := assembledBatchHeader.PrevStateRoot()
	require.NoError(t, err)
	require.Equal(t, batchDataInput.PrevStateRoot[:], prevStateRoot.Bytes())
	postStateRoot, err := assembledBatchHeader.PostStateRoot()
	require.NoError(t, err)
	require.Equal(t, batchDataInput.PostStateRoot[:], postStateRoot.Bytes())

	// Compare assembledBatchHeader with the batch header built from commitBatch data
	// Note: batchDataInput and batchSignatureInput can be used to verify data, but need to build a complete batch header
	compareBatchHeaderWithCommitData(t, assembledBatchHeader, batchDataInput, batchSignatureInput, sequencerSetVerifyHash)

	committedBatchHash, err := rollupContract.CommittedBatches(nil, new(big.Int).SetUint64(assembledBatchIndex))
	require.NoError(t, err)
	require.Equal(t, assembledBatchHash, common.Hash(committedBatchHash), "assembled batch hash should match")
	t.Logf("Successfully assembled batch hash: %x", assembledBatchHash)
	t.Logf("Successfully assembled batch header: batchIndex=%d", assembledBatchIndex)
}

// compareAndReportBatchHeaders compares two batch headers and reports all mismatched fields
func compareAndReportBatchHeaders(t *testing.T, batchHeader1 *BatchHeaderBytes, batchHeader2 *BatchHeaderBytes, name1, name2 string) {
	var mismatches []string

	// Compare BatchIndex
	index1, err1 := batchHeader1.BatchIndex()
	index2, err2 := batchHeader2.BatchIndex()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get BatchIndex: err1=%v, err2=%v", err1, err2)
		return
	}
	if index1 != index2 {
		mismatches = append(mismatches, fmt.Sprintf("BatchIndex: %s=%d, %s=%d", name1, index1, name2, index2))
	} else {
		t.Logf("✓ BatchIndex: %d (match)", index1)
	}

	// Compare L1MessagePopped
	l1Msg1, err1 := batchHeader1.L1MessagePopped()
	l1Msg2, err2 := batchHeader2.L1MessagePopped()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get L1MessagePopped: err1=%v, err2=%v", err1, err2)
		return
	}
	if l1Msg1 != l1Msg2 {
		mismatches = append(mismatches, fmt.Sprintf("L1MessagePopped: %s=%d, %s=%d", name1, l1Msg1, name2, l1Msg2))
	} else {
		t.Logf("✓ L1MessagePopped: %d (match)", l1Msg1)
	}

	// Compare TotalL1MessagePopped
	totalL1Msg1, err1 := batchHeader1.TotalL1MessagePopped()
	totalL1Msg2, err2 := batchHeader2.TotalL1MessagePopped()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get TotalL1MessagePopped: err1=%v, err2=%v", err1, err2)
		return
	}
	if totalL1Msg1 != totalL1Msg2 {
		mismatches = append(mismatches, fmt.Sprintf("TotalL1MessagePopped: %s=%d, %s=%d", name1, totalL1Msg1, name2, totalL1Msg2))
	} else {
		t.Logf("✓ TotalL1MessagePopped: %d (match)", totalL1Msg1)
	}

	// Compare DataHash
	dataHash1, err1 := batchHeader1.DataHash()
	dataHash2, err2 := batchHeader2.DataHash()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get DataHash: err1=%v, err2=%v", err1, err2)
		return
	}
	if dataHash1 != dataHash2 {
		mismatches = append(mismatches, fmt.Sprintf("DataHash: %s=%x, %s=%x", name1, dataHash1, name2, dataHash2))
	} else {
		t.Logf("✓ DataHash: %x (match)", dataHash1)
	}

	// Compare BlobVersionedHash
	blobHash1, err1 := batchHeader1.BlobVersionedHash()
	blobHash2, err2 := batchHeader2.BlobVersionedHash()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get BlobVersionedHash: err1=%v, err2=%v", err1, err2)
		return
	}
	if blobHash1 != blobHash2 {
		mismatches = append(mismatches, fmt.Sprintf("BlobVersionedHash: %s=%x, %s=%x", name1, blobHash1, name2, blobHash2))
	} else {
		t.Logf("✓ BlobVersionedHash: %x (match)", blobHash1)
	}

	// 比较 PrevStateRoot
	prevStateRoot1, err1 := batchHeader1.PrevStateRoot()
	prevStateRoot2, err2 := batchHeader2.PrevStateRoot()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get PrevStateRoot: err1=%v, err2=%v", err1, err2)
		return
	}
	if prevStateRoot1 != prevStateRoot2 {
		mismatches = append(mismatches, fmt.Sprintf("PrevStateRoot: %s=%x, %s=%x", name1, prevStateRoot1, name2, prevStateRoot2))
	} else {
		t.Logf("✓ PrevStateRoot: %x (match)", prevStateRoot1)
	}

	// 比较 PostStateRoot
	postStateRoot1, err1 := batchHeader1.PostStateRoot()
	postStateRoot2, err2 := batchHeader2.PostStateRoot()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get PostStateRoot: err1=%v, err2=%v", err1, err2)
		return
	}
	if postStateRoot1 != postStateRoot2 {
		mismatches = append(mismatches, fmt.Sprintf("PostStateRoot: %s=%x, %s=%x", name1, postStateRoot1, name2, postStateRoot2))
	} else {
		t.Logf("✓ PostStateRoot: %x (match)", postStateRoot1)
	}

	// 比较 WithdrawalRoot
	withdrawRoot1, err1 := batchHeader1.WithdrawalRoot()
	withdrawRoot2, err2 := batchHeader2.WithdrawalRoot()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get WithdrawalRoot: err1=%v, err2=%v", err1, err2)
		return
	}
	if withdrawRoot1 != withdrawRoot2 {
		mismatches = append(mismatches, fmt.Sprintf("WithdrawalRoot: %s=%x, %s=%x", name1, withdrawRoot1, name2, withdrawRoot2))
	} else {
		t.Logf("✓ WithdrawalRoot: %x (match)", withdrawRoot1)
	}

	// 比较 SequencerSetVerifyHash
	seqHash1, err1 := batchHeader1.SequencerSetVerifyHash()
	seqHash2, err2 := batchHeader2.SequencerSetVerifyHash()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get SequencerSetVerifyHash: err1=%v, err2=%v", err1, err2)
		return
	}
	if seqHash1 != seqHash2 {
		mismatches = append(mismatches, fmt.Sprintf("SequencerSetVerifyHash: %s=%x, %s=%x", name1, seqHash1, name2, seqHash2))
	} else {
		t.Logf("✓ SequencerSetVerifyHash: %x (match)", seqHash1)
	}

	// Compare ParentBatchHash
	parentHash1, err1 := batchHeader1.ParentBatchHash()
	parentHash2, err2 := batchHeader2.ParentBatchHash()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to get ParentBatchHash: err1=%v, err2=%v", err1, err2)
		return
	}
	if parentHash1 != parentHash2 {
		mismatches = append(mismatches, fmt.Sprintf("ParentBatchHash: %s=%x, %s=%x", name1, parentHash1, name2, parentHash2))
	} else {
		t.Logf("✓ ParentBatchHash: %x (match)", parentHash1)
	}

	// Compare LastBlockNumber (if supported)
	lastBlock1, err1 := batchHeader1.LastBlockNumber()
	lastBlock2, err2 := batchHeader2.LastBlockNumber()
	if err1 == nil && err2 == nil {
		if lastBlock1 != lastBlock2 {
			mismatches = append(mismatches, fmt.Sprintf("LastBlockNumber: %s=%d, %s=%d", name1, lastBlock1, name2, lastBlock2))
		} else {
			t.Logf("✓ LastBlockNumber: %d (match)", lastBlock1)
		}
	}

	// Report mismatched fields
	if len(mismatches) > 0 {
		t.Errorf("\n❌ Found %d mismatched fields between %s and %s:", len(mismatches), name1, name2)
		for _, mismatch := range mismatches {
			t.Errorf("  - %s", mismatch)
		}
	} else {
		t.Logf("\n✅ All fields match between %s and %s", name1, name2)
	}
}

// compareBatchHeaderWithCommitData compares the assembled batch header with information extracted from commitBatch data
func compareBatchHeaderWithCommitData(t *testing.T, assembledBatchHeader *BatchHeaderBytes, batchDataInput *bindings.IRollupBatchDataInput, batchSignatureInput *bindings.IRollupBatchSignatureInput, sequencerSetVerifyHash common.Hash) {
	t.Logf("\n=== Comparing assembled batch header with commitBatch data ===")

	// Compare Version
	version, err := assembledBatchHeader.Version()
	require.NoError(t, err)
	if version != batchDataInput.Version {
		t.Errorf("❌ Version mismatch: assembled=%d, commitBatch=%d", version, batchDataInput.Version)
	} else {
		t.Logf("✓ Version: %d (match)", version)
	}

	// Compare ParentBatchHeader
	// Note: We should use batch index instead of version, but we need to get batch index from assembledBatchHeader
	batchIndex, err := assembledBatchHeader.BatchIndex()
	if err == nil && batchIndex > 0 {
		parentBatchHeader, err := getLastFinalizeBatchHeaderByIndex(batchIndex - 1)
		if err == nil {
			parentBytes := parentBatchHeader.Bytes()
			if !bytes.Equal(parentBytes, batchDataInput.ParentBatchHeader) {
				t.Errorf("❌ ParentBatchHeader mismatch: assembled=%x, commitBatch=%x", parentBytes[:min(32, len(parentBytes))], batchDataInput.ParentBatchHeader[:min(32, len(batchDataInput.ParentBatchHeader))])
			} else {
				t.Logf("✓ ParentBatchHeader: match")
			}
		}
	}

	// Compare LastBlockNumber
	lastBlock, err := assembledBatchHeader.LastBlockNumber()
	if err == nil {
		if lastBlock != batchDataInput.LastBlockNumber {
			t.Errorf("❌ LastBlockNumber mismatch: assembled=%d, commitBatch=%d", lastBlock, batchDataInput.LastBlockNumber)
		} else {
			t.Logf("✓ LastBlockNumber: %d (match)", lastBlock)
		}
	}

	// Compare NumL1Messages
	l1MsgPopped, err := assembledBatchHeader.L1MessagePopped()
	require.NoError(t, err)
	if l1MsgPopped != uint64(batchDataInput.NumL1Messages) {
		t.Errorf("❌ NumL1Messages mismatch: assembled=%d, commitBatch=%d", l1MsgPopped, batchDataInput.NumL1Messages)
	} else {
		t.Logf("✓ NumL1Messages: %d (match)", l1MsgPopped)
	}

	// 比较 PrevStateRoot
	prevStateRoot, err := assembledBatchHeader.PrevStateRoot()
	require.NoError(t, err)
	prevStateRootFromCommit := common.BytesToHash(batchDataInput.PrevStateRoot[:])
	if prevStateRoot != prevStateRootFromCommit {
		t.Errorf("❌ PrevStateRoot mismatch: assembled=%x, commitBatch=%x", prevStateRoot, prevStateRootFromCommit)
	} else {
		t.Logf("✓ PrevStateRoot: %x (match)", prevStateRoot)
	}

	// 比较 PostStateRoot
	postStateRoot, err := assembledBatchHeader.PostStateRoot()
	require.NoError(t, err)
	postStateRootFromCommit := common.BytesToHash(batchDataInput.PostStateRoot[:])
	if postStateRoot != postStateRootFromCommit {
		t.Errorf("❌ PostStateRoot mismatch: assembled=%x, commitBatch=%x", postStateRoot, postStateRootFromCommit)
	} else {
		t.Logf("✓ PostStateRoot: %x (match)", postStateRoot)
	}

	// 比较 WithdrawalRoot
	withdrawRoot, err := assembledBatchHeader.WithdrawalRoot()
	require.NoError(t, err)
	withdrawRootFromCommit := common.BytesToHash(batchDataInput.WithdrawalRoot[:])
	if withdrawRoot != withdrawRootFromCommit {
		t.Errorf("❌ WithdrawalRoot mismatch: assembled=%x, commitBatch=%x", withdrawRoot, withdrawRootFromCommit)
	} else {
		t.Logf("✓ WithdrawalRoot: %x (match)", withdrawRoot)
	}

	// 比较 SequencerSetVerifyHash
	sequencerSetsHash := crypto.Keccak256Hash(batchSignatureInput.SequencerSets)
	seqHash, err := assembledBatchHeader.SequencerSetVerifyHash()
	require.NoError(t, err)
	if seqHash != sequencerSetsHash {
		t.Errorf("❌ SequencerSetVerifyHash mismatch: assembled=%x, from SequencerSets=%x", seqHash, sequencerSetsHash)
	} else {
		t.Logf("✓ SequencerSetVerifyHash: %x (match)", seqHash)
	}

	if seqHash != sequencerSetVerifyHash {
		t.Errorf("❌ SequencerSetVerifyHash mismatch with provided hash: assembled=%x, provided=%x", seqHash, sequencerSetVerifyHash)
	} else {
		t.Logf("✓ SequencerSetVerifyHash matches provided hash: %x", sequencerSetVerifyHash)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getBatchHeaderFromGeth(index uint64) (*BatchHeaderBytes, error) {
	batch, err := l2Client.GetRollupBatchByIndex(context.Background(), index+1)
	if err != nil {
		return nil, err
	}
	batchHeaderBytes := BatchHeaderBytes(batch.ParentBatchHeader[:])
	return &batchHeaderBytes, nil
}

// getLastFinalizeBatchHeaderByIndex gets the batch header with the specified index from the rollup contract's FinalizeBatch event
// The finalizeBatch function only receives one parameter: batchHeader bytes, so it can be parsed directly from the transaction
// Query is limited to 10000 block heights, starting from the latest height and querying backwards until data is found
func getLastFinalizeBatchHeaderByIndex(index uint64) (*BatchHeaderBytes, error) {
	// Get the current latest block height
	latestBlock, err := l1Client.BlockNumber(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block number: %w", err)
	}

	const blockRange = uint64(10000) // Query 10000 blocks each time
	var endBlock uint64 = latestBlock
	var startBlock uint64

	// Start from the latest height, query backwards 10000 blocks each time until data is found
	for endBlock > 0 {
		// Calculate the start block for this query
		if endBlock >= blockRange {
			startBlock = endBlock - blockRange + 1
		} else {
			startBlock = 0
		}

		// Set query options
		filterOpts := &bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}

		// Query the FinalizeBatch event with the corresponding index from the rollup contract
		finalizeEventIter, err := rollupContract.FilterFinalizeBatch(filterOpts, []*big.Int{new(big.Int).SetUint64(index)}, nil)
		if err != nil {
			// If query fails, continue querying backwards
			if endBlock < blockRange {
				break // Already queried to block 0, exit loop
			}
			endBlock = startBlock - 1
			continue
		}

		// Iterate through query results
		for finalizeEventIter.Next() {
			event := finalizeEventIter.Event
			// Get transaction hash from event
			txHash := event.Raw.TxHash

			// Get transaction details
			tx, _, err := l1Client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				continue // If getting transaction fails, try next event
			}

			// Parse finalizeBatch transaction data to get batchHeader
			batchHeader, err := parseFinalizeBatchTxData(tx.Data())
			if err != nil {
				continue // If parsing fails, try next event
			}

			// Verify if batch index matches
			batchIndex, err := batchHeader.BatchIndex()
			if err != nil {
				continue
			}
			if batchIndex == index {
				finalizeEventIter.Close()
				return &batchHeader, nil
			}
		}
		finalizeEventIter.Close()

		// Continue querying backwards
		if endBlock < blockRange {
			break // Already queried to block 0, exit loop
		}
		endBlock = startBlock - 1
	}

	return nil, ErrBatchNotFound
}

func getInfosFromContract() (*big.Int, *big.Int) {
	latestCommitBatchIndex, _ := rollupContract.LastCommittedBatchIndex(nil)
	lastFinalizedBatchIndex, _ := rollupContract.LastFinalizedBatchIndex(nil)
	return latestCommitBatchIndex, lastFinalizedBatchIndex
}

func getFirstUnFinalizeBatchBlockNumRange(lastFinalizedBatchIndex *big.Int) (*big.Int, *big.Int, error) {
	fis, err := rollupContract.BatchDataStore(nil, lastFinalizedBatchIndex)
	if err != nil {
		return nil, nil, err
	}
	ufis, err := rollupContract.BatchDataStore(nil, new(big.Int).SetUint64(lastFinalizedBatchIndex.Uint64()+1))
	if err != nil {
		return nil, nil, err
	}

	return fis.BlockNumber, ufis.BlockNumber, nil
}

// getCommitBatchDataByIndex gets batchDataInput and batchSignatureInput with the specified index from the rollup contract's CommitBatch event
// Reference the implementation of getLastFinalizeBatchHeaderByIndex
// Query is limited to 10000 block heights, starting from the latest height and querying backwards until data is found
func getCommitBatchDataByIndex(index uint64) (*bindings.IRollupBatchDataInput, *bindings.IRollupBatchSignatureInput, error) {
	// Get the current latest block height
	latestBlock, err := l1Client.BlockNumber(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get latest block number: %w", err)
	}

	const blockRange = uint64(10000) // Query 10000 blocks each time
	var endBlock uint64 = latestBlock
	var startBlock uint64

	// Start from the latest height, query backwards 10000 blocks each time until data is found
	for endBlock > 0 {
		// Calculate the start block for this query
		if endBlock >= blockRange {
			startBlock = endBlock - blockRange + 1
		} else {
			startBlock = 0
		}

		// Set query options
		filterOpts := &bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}

		// Query the CommitBatch event with the corresponding index from the rollup contract
		commitEventIter, err := rollupContract.FilterCommitBatch(filterOpts, []*big.Int{new(big.Int).SetUint64(index)}, nil)
		if err != nil {
			// If query fails, continue querying backwards
			if endBlock < blockRange {
				break // Already queried to block 0, exit loop
			}
			endBlock = startBlock - 1
			continue
		}

		// Iterate through query results
		for commitEventIter.Next() {
			event := commitEventIter.Event
			// Get transaction hash from event
			txHash := event.Raw.TxHash

			// Get transaction details
			tx, _, err := l1Client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to get transaction by hash: %w", err)
			}

			// Parse commitBatch transaction data to get batchDataInput and batchSignatureInput
			batchDataInput, batchSignatureInput, err := parseCommitBatchTxData(tx.Data())
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse commit batch data: %w", err)
			}

			// Verify if batch index matches (by checking batchIndex in parentBatchHeader)
			if len(batchDataInput.ParentBatchHeader) > 0 {
				parentHeader := BatchHeaderBytes(batchDataInput.ParentBatchHeader)
				parentBatchIndex, err := parentHeader.BatchIndex()
				if err == nil && parentBatchIndex+1 == index {
					commitEventIter.Close()
					return batchDataInput, batchSignatureInput, nil
				}
			}
		}
		commitEventIter.Close()

		// Continue querying backwards
		if endBlock < blockRange {
			break // Already queried to block 0, exit loop
		}
		endBlock = startBlock - 1
	}

	return nil, nil, ErrBatchNotFound
}

// assembleBatchHeaderFromL2Blocks fetches blocks from L2 client in the specified range and assembles batchHeader
// Parameters:
//   - bc: BatchCache instance (parentBatchHeader and prevStateRoot already initialized)
//   - startBlockNum: starting block number
//   - endBlockNum: ending block number
//   - sequencerSetVerifyHash: sequencer set verification hash
//   - l2Client: L2 client
//
// Returns:
//   - batchHeader: assembled batchHeader
//   - error: returns error if assembly fails
func assembleBatchHeaderFromL2Blocks(
	bc *BatchCache,
	startBlockNum, endBlockNum uint64,
	sequencerBytes []byte,
	l2Client iface.L2Client,
	l2Caller *types.L2Caller,
) (*BatchHeaderBytes, error) {
	ctx := context.Background()

	// Fetch blocks from L2 client in the specified range and accumulate to batch
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		root, err := l2Caller.GetTreeRoot(&bind.CallOpts{
			Context:     ctx,
			BlockNumber: new(big.Int).SetUint64(blockNum),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get withdraw root at block %d: %w", blockNum, err)
		}
		// Check capacity and store to current
		exceeded, err := bc.CalculateCapWithProposalBlock(blockNum, root)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate cap with block %d: %w", blockNum, err)
		}

		// Pack current block (confirm and append to batch)
		if err = bc.PackCurrentBlock(blockNum); err != nil {
			return nil, fmt.Errorf("failed to pack block %d: %w", blockNum, err)
		}

		// If capacity exceeds limit, can stop early (optional)
		_ = exceeded // Checked but not used in this test
	}

	// Get the last block's timestamp for packing
	lastBlock, err := l2Client.BlockByNumber(ctx, big.NewInt(int64(endBlockNum)))
	if err != nil {
		return nil, fmt.Errorf("failed to get last block %d: %w", endBlockNum, err)
	}
	blockTimestamp := lastBlock.Time()

	// Seal batch and generate batchHeader
	batchIndex, batchHeaderBytes, _, err := bc.SealBatch(sequencerBytes, blockTimestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to seal batch: %w", err)
	}

	// Get the sealed batch header
	_, found := bc.GetSealedBatch(batchIndex)
	if !found {
		return nil, fmt.Errorf("sealed batch not found for index %d", batchIndex)
	}

	return &batchHeaderBytes, nil
}
