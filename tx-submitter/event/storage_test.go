package event

import (
	"morph-l2/tx-submitter/db"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventInfoStorage(t *testing.T) {
	// Create a unique test directory
	testDir := filepath.Join(t.TempDir(), "testleveldb")

	// Cleanup before test (in case it exists)
	os.RemoveAll(testDir)

	// Cleanup after test
	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})

	db, err := db.New(testDir)
	require.NoError(t, err)
	storage := NewEventInfoStorage(db)
	err = storage.Load()
	require.NoError(t, err)

	storage.SetBlockTime(100)
	storage.SetBlockProcessed(100)
	err = storage.Store()
	require.NoError(t, err)

	storage2 := NewEventInfoStorage(db)
	err = storage2.Load()
	require.NoError(t, err)
	require.Equal(t, storage.BlockTime(), storage2.BlockTime())
	require.Equal(t, storage.BlockProcessed(), storage2.BlockProcessed())
}
