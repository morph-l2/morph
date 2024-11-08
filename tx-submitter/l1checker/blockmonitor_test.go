package l1checker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIsGrowth(t *testing.T) {

	blockCnt := int64(2)
	monitor := NewBlockMonitor(blockCnt, nil)
	monitor.latestBlockTime = time.Time{}
	require.Equal(t, false, monitor.IsGrowth())

	monitor.latestBlockTime = time.Now()
	require.Equal(t, false, monitor.IsGrowth())

	monitor.latestBlockTime = time.Now().Add(-monitor.noGrowthBlockCntTime)
	require.Equal(t, true, monitor.IsGrowth())

}
