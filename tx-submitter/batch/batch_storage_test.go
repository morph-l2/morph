package batch

import (
	"testing"

	"github.com/stretchr/testify/require"

	"morph-l2/tx-submitter/iface"
)

func Test_storageBatch(t *testing.T) {
	testDB := setupTestDB(t)
	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)
	err := cache.InitAndSyncFromRollup()
	require.NoError(t, err)

	batches, err := cache.batchStorage.LoadAllSealedBatches()
	require.NoError(t, err)
	require.NotNil(t, batches)
	t.Log("loaded batches count", len(batches))
}
