package event

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventInfoStorage(t *testing.T) {
	filename := "hello.json"
	storage := NewEventInfoStorage(filename)
	err := storage.Load()
	require.NoError(t, err)

	storage.BlockTime = 100
	storage.BlockProcessed = 100
	err = storage.Store()
	require.NoError(t, err)

	storage2 := NewEventInfoStorage(filename)
	err = storage2.Load()
	require.NoError(t, err)
	require.Equal(t, storage.BlockTime, storage2.BlockTime)
	require.Equal(t, storage.BlockProcessed, storage2.BlockProcessed)
}
