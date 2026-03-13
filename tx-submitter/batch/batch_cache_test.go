package batch

import (
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/db"
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

// setupTestDB creates a temporary database for testing
func setupTestDB(t *testing.T) *db.Db {
	testDir := filepath.Join(t.TempDir(), "testleveldb")
	os.RemoveAll(testDir)
	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})

	testDB, err := db.New(testDir)
	require.NoError(t, err)
	return testDB
}

func TestBatchCacheInitServer(t *testing.T) {
	testDB := setupTestDB(t)
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)

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

	go utils.Loop(cache.ctx, 5*time.Second, func() {
		batchCacheSyncMu.Lock()
		defer batchCacheSyncMu.Unlock()
		err := cache.AssembleCurrentBatchHeader()
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
	testDB := setupTestDB(t)
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)
	err := cache.InitAndSyncFromRollup()
	require.NoError(t, err)
}

func TestBatchCacheInitByBlockRange(t *testing.T) {
	testDB := setupTestDB(t)
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)
	err := cache.InitFromRollupByRange()
	require.NoError(t, err)
}
