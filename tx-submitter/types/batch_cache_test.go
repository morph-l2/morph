package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBatchCache(t *testing.T) {
	cache := NewBatchCache()
	cache.Set(1, nil)
	_, ok := cache.Get(1)
	require.True(t, ok)
	cache.Delete(1)
	_, ok = cache.Get(1)
	require.False(t, ok)
	_, ok = cache.Get(2)
	require.False(t, ok)
}
