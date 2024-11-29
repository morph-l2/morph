package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetFailedStatus(t *testing.T) {
	// index 2 failed -> set failed index 2 success
	pt := NewPendingTxs(nil, nil, nil)
	pt.SetPindex(2)
	require.Nil(t, pt.failedIndex)
	pt.SetFailedStatus(2)
	require.NotNil(t, pt.failedIndex)
	require.EqualValues(t, 2, *pt.failedIndex)

	// failed index =2
	// new failed index = 3
	// set failed index failed
	pt = NewPendingTxs(nil, nil, nil)
	failedIndex := uint64(2)
	pt.failedIndex = &failedIndex
	pt.SetFailedStatus(3)
	require.EqualValues(t, 2, *pt.failedIndex)

	// set failed index without pindex -> failed
	pt = NewPendingTxs(nil, nil, nil)
	pt.SetFailedStatus(2)
	require.Nil(t, pt.failedIndex)
}
