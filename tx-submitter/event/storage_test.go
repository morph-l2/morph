package event

import (
	"morph-l2/tx-submitter/db"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventInfoStorage(t *testing.T) {

	db, err := db.New("./testleveldb")
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
