package oracle

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestSortBlock(t *testing.T) {
	arr := []int{5, 2, 3, 4, 2, 4, 5, 1, 3, 4, 1, 0, 9, 8}
	sortedArr := removeDuplicatesAndSort(arr)
	sort.Ints(sortedArr)
	fmt.Println(sortedArr)
}

func TestXXXXXX(t *testing.T) {
	o := testNewOracleClient(t)
	err := o.recordRollupEpoch()
	require.NoError(t, err)
}
