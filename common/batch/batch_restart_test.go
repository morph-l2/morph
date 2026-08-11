package batch

import (
	"encoding/hex"
	"testing"

	"morph-l2/bindings/bindings"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var (
	rollupAddr = common.HexToAddress("0x0165878a594ca255338adfa4d48449f69242eb8f")

	l1ClientRpc = "http://localhost:9545"
	l2ClientRpc = "http://localhost:8545"
	l1Client, _ = ethclient.Dial(l1ClientRpc)
	l2Client, _ = ethclient.Dial(l2ClientRpc)

	rollupContract *bindings.Rollup
	l2Gov          *L2Gov
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

func TestGetCanonicalCommittedBatchHeader(t *testing.T) {
	requireBatchIntegration(t)
	cache := NewBatchCache(
		nil,
		nil,
		2,
		testBatchConfig,
		l1Client,
		&SingleL2Client{C: l2Client},
		rollupContract,
		l2Gov,
		openTestKV(t),
	)
	var header *BatchHeaderBytes
	err := cache.withConfirmedL1Snapshot(func() error {
		var lookupErr error
		header, lookupErr = cache.getCanonicalCommittedBatchHeaderByIndex(0)
		return lookupErr
	})
	require.NoError(t, err)
	t.Log("headerBytes", hex.EncodeToString(header.Bytes()))
}

// This opt-in integration test exercises the same canonical bootstrap entry
// used by tx-submitter; it intentionally does not recover from FinalizeBatch
// calldata or choose the first CommitBatch log.
func TestBatchRestartInit(t *testing.T) {
	requireBatchIntegration(t)
	cache := NewBatchCache(
		nil,
		nil,
		2,
		testBatchConfig,
		l1Client,
		&SingleL2Client{C: l2Client},
		rollupContract,
		l2Gov,
		openTestKV(t),
	)
	require.NoError(t, cache.InitAndSyncFromRollup())
	require.True(t, cache.isInitialized())
}
