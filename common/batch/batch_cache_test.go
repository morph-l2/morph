package batch

import (
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"

	"morph-l2/bindings/bindings"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func init() {
	var err error
	rollupContract, err = bindings.NewRollup(rollupAddr, l1Client)
	if err != nil {
		panic(err)
	}
	l2Gov, err = NewL2Gov(l2Client)
	if err != nil {
		panic(err)
	}
}

func TestBatchCacheInitServer(t *testing.T) {
	testDB := openTestKV(t)
	cache := NewBatchCache(nil, nil, 2, l1Client, &SingleL2Client{C: l2Client}, rollupContract, l2Gov, testDB)

	var batchCacheSyncMu sync.Mutex

	go func() {
		batchCacheSyncMu.Lock()
		defer batchCacheSyncMu.Unlock()
		for {
			if err := cache.InitAndSyncFromDatabase(); err != nil {
				log.Error("init and sync from database failed, wait for the next try", "error", err)
				time.Sleep(5 * time.Second)
				continue
			}
			break
		}
	}()

	go testLoop(cache.ctx, 5*time.Second, func() {
		batchCacheSyncMu.Lock()
		defer batchCacheSyncMu.Unlock()
		err := cache.AssembleCurrentBatchHeader()
		if err != nil {
			log.Error("Assemble current batch failed, wait for the next try", "error", err)
		}
	})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
}

func TestBatchCacheInit(t *testing.T) {
	testDB := openTestKV(t)
	cache := NewBatchCache(nil, nil, 2, l1Client, &SingleL2Client{C: l2Client}, rollupContract, l2Gov, testDB)
	err := cache.InitAndSyncFromRollup()
	require.NoError(t, err)
}

func TestBatchCacheInitByBlockRange(t *testing.T) {
	testDB := openTestKV(t)
	cache := NewBatchCache(nil, nil, 2, l1Client, &SingleL2Client{C: l2Client}, rollupContract, l2Gov, testDB)
	err := cache.InitFromRollupByRange()
	require.NoError(t, err)
}
