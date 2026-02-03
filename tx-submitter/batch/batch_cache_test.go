package batch

import (
	"github.com/morph-l2/go-ethereum/log"
	"morph-l2/tx-submitter/utils"
	"testing"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/types"

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

func TestBatchCacheInitSer(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller)

	go utils.Loop(cache.ctx, 5*time.Second, func() {
		err := cache.InitAndSyncFromRollup()
		if err != nil {
			log.Error("init and sync from rollup failed, wait for the next try", "error", err)
		}
		err = cache.AssembleCurrentBatchHeader()
		if err != nil {
			log.Error("Assemble current batch failed, wait for the next try", "error", err)
		}
	})
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
