package batch

import (
	"context"
	"fmt"
	"morph-l2/tx-submitter/utils"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"

	"morph-l2/tx-submitter/db"
	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

var (
	// Fill this with hex-encoded batch header bytes, e.g. "0x00....".
	// This test will use it as the genesis parent header to initialize cache.
	globalGenesisBatchHeaderHex  = "0x00000000000000000000000000000000000000000000000000d81a073a4abd227068a2a334f4a41b3abba26144dc866a78ed28e2ae90f86f5a010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c4440140000000000000000000000000000000000000000000000000000000000000000290233e7a85533655c301d3e1043f03acd5427c73d1bbcbf8784db3f3974327f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	globalGenesisBatchHeader     *BatchHeaderBytes
	globalGenesisBatchHeaderErr  error
	globalGenesisBatchHeaderOnce sync.Once

	// Global overrides for cache batch config in tests (instead of updateBatchConfigFromGov).
	globalBatchTimeoutForTest  uint64 = 10000000
	globalBlockIntervalForTest uint64 = 10000
)

func ensureGlobalGenesisBatchHeader() error {
	globalGenesisBatchHeaderOnce.Do(func() {
		if globalGenesisBatchHeaderHex == "" {
			globalGenesisBatchHeaderErr = fmt.Errorf("globalGenesisBatchHeaderHex is empty")
			return
		}
		raw, err := hexutil.Decode(globalGenesisBatchHeaderHex)
		if err != nil {
			globalGenesisBatchHeaderErr = fmt.Errorf("decode globalGenesisBatchHeaderHex failed: %w", err)
			return
		}
		header := BatchHeaderBytes(raw)
		if err := header.validate(); err != nil {
			globalGenesisBatchHeaderErr = fmt.Errorf("invalid global genesis batch header: %w", err)
			return
		}
		globalGenesisBatchHeader = &header
	})
	return globalGenesisBatchHeaderErr
}

// initCacheWithGlobalGenesisHeader initializes cache base fields from the
// globally cached genesis batch header, instead of loading through Init().
func initCacheWithGlobalGenesisHeader(cache *BatchCache) error {
	if err := ensureGlobalGenesisBatchHeader(); err != nil {
		return err
	}
	if globalGenesisBatchHeader == nil {
		return db.ErrKeyNotFound
	}
	// Use global test knobs instead of querying gov config from chain.
	cache.batchTimeOut = globalBatchTimeoutForTest
	cache.blockInterval = globalBlockIntervalForTest
	headerCopy := make(BatchHeaderBytes, len(*globalGenesisBatchHeader))
	copy(headerCopy, *globalGenesisBatchHeader)
	cache.parentBatchHeader = &headerCopy

	prevStateRoot, err := cache.parentBatchHeader.PostStateRoot()
	if err != nil {
		return err
	}
	cache.prevStateRoot = prevStateRoot

	totalL1MessagePopped, err := cache.parentBatchHeader.TotalL1MessagePopped()
	if err != nil {
		return err
	}
	cache.totalL1MessagePopped = totalL1MessagePopped

	lastPackedBlockHeight, err := cache.parentBatchHeader.LastBlockNumber()
	if err != nil {
		lastPackedBlockHeight = 0
	}
	cache.lastPackedBlockHeight = lastPackedBlockHeight
	cache.currentBlockNumber = lastPackedBlockHeight
	cache.initDone = true

	return nil
}

func TestBatchCacheInitWithGlobalGenesisHeader(t *testing.T) {
	testDB := setupTestDB(t)
	a := func(uint64) bool { return true }
	cache := NewBatchCache(nil, a, 3, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)

	var batchCacheSyncMu sync.Mutex
	done := make(chan error, 1)
	go func() {
		batchCacheSyncMu.Lock()
		defer batchCacheSyncMu.Unlock()
		for {
			if err := initCacheWithGlobalGenesisHeader(cache); err != nil {
				log.Error("init with global genesis header failed, wait for next try", "error", err)
				time.Sleep(3 * time.Second)
				continue
			}
			done <- nil
			return
		}
	}()

	select {
	case err := <-done:
		require.NoError(t, err)
	case <-time.After(20 * time.Second):
		t.Fatal("timeout waiting for cache init with global genesis header")
	}

	require.True(t, cache.initDone)
	require.NotNil(t, cache.parentBatchHeader)
	version, err := cache.parentBatchHeader.Version()
	require.NoError(t, err)
	require.Equal(t, uint8(BatchHeaderVersion0), version)
	require.Equal(t, cache.lastPackedBlockHeight, cache.currentBlockNumber)
	_, err = cache.l2Clients.BlockNumber(context.Background())
	require.NoError(t, err)

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
