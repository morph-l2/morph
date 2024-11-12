package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Get_Put(t *testing.T) {
	db, err := New("test")
	require.NoError(t, err)

	expect := 1.123456789012345
	// expectStr := "1.123456789012345"

	err = db.PutFloat("test", expect)
	require.NoError(t, err)
	v, err := db.GetFloat("test")
	require.NoError(t, err)
	require.Equal(t, expect, v)
}
