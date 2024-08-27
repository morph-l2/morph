package db

import (
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"morph-l2/node/types"
)

func TestLatestSyncedL1Height(t *testing.T) {
	db := NewMemoryStore()
	db.WriteLatestSyncedL1Height(100)
	require.EqualValues(t, 100, *db.ReadLatestSyncedL1Height())
	db.WriteLatestSyncedL1Height(101)
	require.EqualValues(t, 101, *db.ReadLatestSyncedL1Height())
}

func TestSyncedL1Messages(t *testing.T) {
	db := NewMemoryStore()

	to := common.BigToAddress(big.NewInt(101))
	msgs := make([]types.L1Message, 0)
	for i := 0; i < 200; i++ {
		msg := types.L1Message{
			L1MessageTx: eth.L1MessageTx{
				QueueIndex: uint64(i),
				Gas:        500000 + uint64(i),
				To:         &to,
				Value:      big.NewInt(3e9),
				Data:       []byte("0x1a2b3c"),
				Sender:     common.BigToAddress(big.NewInt(202)),
			},
			L1TxHash: common.BigToHash(big.NewInt(1111)),
		}
		msgs = append(msgs, msg)
	}
	err := db.WriteSyncedL1Messages(msgs, 20000)
	require.NoError(t, err)

	rangeMsgs := db.ReadL1MessagesInRange(100, 150)
	for i, msg := range rangeMsgs {
		require.EqualValues(t, uint64(i+100), msg.QueueIndex)
	}
	require.EqualValues(t, 51, len(rangeMsgs))
	require.EqualValues(t, 20000, *db.ReadLatestSyncedL1Height())

	msg := db.ReadL1MessageByIndex(190)
	require.EqualValues(t, 190, msg.QueueIndex)

	msg = db.ReadL1MessageByIndex(200)
	require.Nil(t, msg)
}
