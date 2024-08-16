package receipt

import (
	"context"

	lru "github.com/hashicorp/golang-lru"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
)

type Fetcher struct {
	l1Client  *ethclient.Client
	rpcClient *rpc.Client

	// cache receipts in bundles per block hash
	// We cache the receipts fetching job to not lose progress when we have to retry the `Fetch` call
	// common.Hash -> *receiptsFetchingJob
	receiptsCache *lru.Cache
	maxBatchSize  int
}

type FetcherConfig struct {
	ReceiptsCacheSize, MaxRequestPerBatch int
}

func defaultFetcherConfig() *FetcherConfig {
	return &FetcherConfig{
		ReceiptsCacheSize:  100,
		MaxRequestPerBatch: 100,
	}
}

func NewFetcher(l1Addr string, config *FetcherConfig) (*Fetcher, error) {
	c, err := rpc.DialContext(context.Background(), l1Addr)
	if err != nil {
		return nil, err
	}
	if config == nil {
		config = defaultFetcherConfig()
	}
	cache, err := lru.New(config.ReceiptsCacheSize)
	if err != nil {
		return nil, err
	}
	l1Client := ethclient.NewClient(c)
	return &Fetcher{
		l1Client:  l1Client,
		rpcClient: c,

		receiptsCache: cache,
		maxBatchSize:  config.MaxRequestPerBatch,
	}, nil
}

func (f *Fetcher) Fetch(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	l1Block, err := f.l1Client.BlockByHash(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	txs := l1Block.Transactions()
	// Try to reuse the receipts fetcher because is caches the results of intermediate calls. This means
	// that if just one of many calls fail, we only retry the failed call rather than all of the calls.
	// The underlying fetcher uses the receipts hash to verify receipt integrity.
	var job *receiptsFetchingJob
	if v, ok := f.receiptsCache.Get(blockHash); ok {
		job = v.(*receiptsFetchingJob)
	} else {
		txHashes := make([]common.Hash, len(txs))
		for i := 0; i < len(txs); i++ {
			txHashes[i] = txs[i].Hash()
		}
		job = NewReceiptsFetchingJob(f.rpcClient, f.maxBatchSize, ToBlockID(l1Block), l1Block.ReceiptHash(), txHashes)
		f.receiptsCache.Add(blockHash, job)
	}
	return job.Fetch(ctx)
}
