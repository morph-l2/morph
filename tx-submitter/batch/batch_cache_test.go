package batch

import (
	"os"
	"os/signal"
	"testing"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/types"
	"morph-l2/tx-submitter/utils"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/stretchr/testify/require"
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

func TestBatchCacheInitServer(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller)

	go utils.Loop(cache.ctx, 5*time.Second, func() {
		err := cache.InitFromRollupByRange()
		if err != nil {
			log.Error("init and sync from rollup failed, wait for the next try", "error", err)
		}
		cache.batchTimeOut = 60
		err = cache.AssembleCurrentBatchHeader()
		if err != nil {
			log.Error("Assemble current batch failed, wait for the next try", "error", err)
		}
	})

	// Catch CTRL-C to ensure a graceful shutdown.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Wait until the interrupt signal is received from an OS signal.
	<-interrupt
}

func TestBatchCacheInit(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller)
	err := cache.InitAndSyncFromRollup()
	require.NoError(t, err)
}

func TestBatchCacheInitByBlockRange(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller)
	err := cache.InitFromRollupByRange()
	require.NoError(t, err)
}

func TestBatchCacheInitByBlockRange1(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller)
	err := cache.Init()
	require.NoError(t, err)
	batch, err := cache.assembleBatchHeaderFromL2Blocks(0, 18)
	require.NoError(t, err)
	hash, err := batch.Hash()
	require.NoError(t, err)
	t.Log("0-18 batch hash", hash.String())
}

func TestBatchCacheInitByBlockRange2(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller)
	err := cache.Init()
	require.NoError(t, err)
	batch, err := cache.assembleBatchHeaderFromL2Blocks(1, 18)
	require.NoError(t, err)
	hash, err := batch.Hash()
	require.NoError(t, err)
	t.Log("1-18 batch hash", hash.String())
}
