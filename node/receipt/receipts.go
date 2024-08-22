package receipt

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/rpc"
	"github.com/morph-l2/go-ethereum/trie"
)

type BlockID struct {
	Hash   common.Hash `json:"hash"`
	Number uint64      `json:"number"`
}

func ToBlockID(b *types.Block) BlockID {
	return BlockID{
		Hash:   b.Hash(),
		Number: b.NumberU64(),
	}
}

func (id BlockID) String() string {
	return fmt.Sprintf("%s:%d", id.Hash.String(), id.Number)
}

// validateReceipts validates that the receipt contents are valid.
// Warning: contractAddress is not verified, since it is a more expensive operation for data we do not use.
// See go-ethereum/crypto.CreateAddress to verify contract deployment address data based on sender and tx nonce.
func validateReceipts(block BlockID, receiptHash common.Hash, txHashes []common.Hash, receipts []*types.Receipt) error {
	if len(receipts) != len(txHashes) {
		return fmt.Errorf("got %d receipts but expected %d", len(receipts), len(txHashes))
	}
	if len(txHashes) == 0 {
		if receiptHash != types.EmptyRootHash {
			return fmt.Errorf("no transactions, but got non-empty receipt trie root: %s", receiptHash)
		}
	}
	// We don't trust the RPC to provide consistent cached receipt info that we use for critical rollup derivation work.
	// Let's check everything quickly.
	logIndex := uint(0)
	cumulativeGas := uint64(0)
	for i, r := range receipts {
		if r == nil { // on reorgs or other cases the receipts may disappear before they can be retrieved.
			return fmt.Errorf("receipt of tx %d returns nil on retrieval", i)
		}
		if r.TransactionIndex != uint(i) {
			return fmt.Errorf("receipt %d has unexpected tx index %d", i, r.TransactionIndex)
		}
		if r.BlockNumber == nil {
			return fmt.Errorf("receipt %d has unexpected nil block number, expected %d", i, block.Number)
		}
		if r.BlockNumber.Uint64() != block.Number {
			return fmt.Errorf("receipt %d has unexpected block number %d, expected %d", i, r.BlockNumber, block.Number)
		}
		if r.BlockHash != block.Hash {
			return fmt.Errorf("receipt %d has unexpected block hash %s, expected %s", i, r.BlockHash, block.Hash)
		}
		if expected := r.CumulativeGasUsed - cumulativeGas; r.GasUsed != expected {
			return fmt.Errorf("receipt %d has invalid gas used metadata: %d, expected %d", i, r.GasUsed, expected)
		}
		for j, log := range r.Logs {
			if log.Index != logIndex {
				return fmt.Errorf("log %d (%d of tx %d) has unexpected log index %d", logIndex, j, i, log.Index)
			}
			if log.TxIndex != uint(i) {
				return fmt.Errorf("log %d has unexpected tx index %d", log.Index, log.TxIndex)
			}
			if log.BlockHash != block.Hash {
				return fmt.Errorf("log %d of block %s has unexpected block hash %s", log.Index, block.Hash, log.BlockHash)
			}
			if log.BlockNumber != block.Number {
				return fmt.Errorf("log %d of block %d has unexpected block number %d", log.Index, block.Number, log.BlockNumber)
			}
			if log.TxHash != txHashes[i] {
				return fmt.Errorf("log %d of tx %s has unexpected tx hash %s", log.Index, txHashes[i], log.TxHash)
			}
			if log.Removed {
				return fmt.Errorf("canonical log (%d) must never be removed due to reorg", log.Index)
			}
			logIndex++
		}
		cumulativeGas = r.CumulativeGasUsed
		// Note: 3 non-consensus L1 receipt fields are ignored:
		// PostState - not part of L1 ethereum anymore since EIP 658 (part of Byzantium)
		// ContractAddress - we do not care about contract deployments
		// And Optimism L1 fee meta-data in the receipt is ignored as well
	}

	// Sanity-check: external L1-RPC sources are notorious for not returning all receipts,
	// or returning them out-of-order. Verify the receipts against the expected receipt-hash.
	hasher := trie.NewStackTrie(nil)
	computed := types.DeriveSha(types.Receipts(receipts), hasher)
	if receiptHash != computed {
		return fmt.Errorf("failed to fetch list of receipts: expected receipt root %s but computed %s from retrieved receipts", receiptHash, computed)
	}
	return nil
}

// receiptsFetchingJob runs the receipt fetching for a specific block,
// and can re-run and adapt based on the fetching method preferences and errors communicated with the requester.
type receiptsFetchingJob struct {
	m sync.Mutex

	client       *rpc.Client
	maxBatchSize int

	block       BlockID
	receiptHash common.Hash
	txHashes    []common.Hash

	fetcher *IterativeBatchCall[common.Hash, *types.Receipt]

	result types.Receipts
}

func NewReceiptsFetchingJob(
	client *rpc.Client,
	maxBatchSize int,
	block BlockID,
	receiptHash common.Hash,
	txHashes []common.Hash,
) *receiptsFetchingJob {
	return &receiptsFetchingJob{
		client:       client,
		maxBatchSize: maxBatchSize,
		block:        block,
		receiptHash:  receiptHash,
		txHashes:     txHashes,
	}
}

// runFetcher retrieves the result by continuing previous batched receipt fetching work,
// and starting this work if necessary.
func (job *receiptsFetchingJob) runFetcher(ctx context.Context) error {
	if job.fetcher == nil {
		// start new work
		job.fetcher = NewIterativeBatchCall[common.Hash, *types.Receipt](
			job.txHashes,
			makeReceiptRequest,
			job.client.BatchCallContext,
			job.maxBatchSize,
		)
	}
	// Fetch all receipts
	for {
		if err := job.fetcher.Fetch(ctx); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	result, err := job.fetcher.Result()
	if err != nil { // errors if results are not available yet, should never happen.
		return err
	}
	if err := validateReceipts(job.block, job.receiptHash, job.txHashes, result); err != nil {
		job.fetcher.Reset() // if results are fetched but invalid, try restart all the fetching to try and get valid data.
		return err
	}
	// Remember the result, and don't keep the fetcher and tx hashes around for longer than needed
	job.result = result
	job.fetcher = nil
	job.txHashes = nil
	return nil
}

// // receiptsWrapper is a decoding type util. Alchemy in particular wraps the receipts array result.
// type receiptsWrapper struct {
// 	Receipts []*types.Receipt `json:"receipts"`
// }

func makeReceiptRequest(txHash common.Hash) (*types.Receipt, rpc.BatchElem) {
	out := new(types.Receipt)
	return out, rpc.BatchElem{
		Method: "eth_getTransactionReceipt",
		Args:   []any{txHash},
		Result: &out, // receipt may become nil, double pointer is intentional
	}
}

// Fetch makes the job fetch the receipts, and returns the results, if any.
// An error may be returned if the fetching is not successfully completed,
// and fetching may be continued/re-attempted by calling Fetch again.
// The job caches the result, so repeated Fetches add no additional cost.
// Fetch is safe to be called concurrently, and will lock to avoid duplicate work or internal inconsistency.
func (job *receiptsFetchingJob) Fetch(ctx context.Context) (types.Receipts, error) {
	job.m.Lock()
	defer job.m.Unlock()

	if job.result != nil {
		return job.result, nil
	}

	if err := job.runFetcher(ctx); err != nil {
		return nil, err
	}

	return job.result, nil
}
