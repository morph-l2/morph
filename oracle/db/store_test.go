package db

import (
	"testing"

	"morph-l2/oracle/types"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestLatestSyncedL1Height(t *testing.T) {
	db := NewMemoryStore()
	err := db.WriteLatestChangeContext(types.ChangeContext{
		L2Sequencers: []types.L2Sequencer{
			{
				TimeStamp:     1627891200,
				EpochInterval: 10,
				Addresses:     []common.Address{common.HexToAddress("0x123")},
			},
			{
				TimeStamp:     1627891260,
				EpochInterval: 10,
				Addresses:     []common.Address{},
			},
		},
		ActiveStakersByTime: []types.ActiveStakers{
			{
				TimeStamp:   1627891200,
				BlockNumber: 100,
				Addresses:   []common.Address{},
			},
			{
				TimeStamp:   1627891260,
				BlockNumber: 101,
				Addresses:   []common.Address{},
			},
		},
		ChangePoints: []types.ChangePoint{
			{
				TimeStamp:     1627891200,
				BlockNumber:   100,
				EpochInterval: 10,
				Submitters:    []common.Address{},
				ChangeType:    1,
			},
			{
				TimeStamp:     1627891260,
				BlockNumber:   101,
				EpochInterval: 10,
				Submitters:    []common.Address{},
				ChangeType:    2,
			},
		},
		L1Synced: 1,
		L2synced: 1,
	})
	require.NoError(t, err)
	changePoints := db.ReadLatestChangePoints()
	t.Log(changePoints)
}
