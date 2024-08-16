package types

import (
	"math/big"
	"testing"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestBLSMessage(t *testing.T) {
	c := BLSMessage{
		ParentHash: common.BigToHash(big.NewInt(1000)),
		Miner:      common.BigToAddress(big.NewInt(2000)),
		Number:     1,
		GasLimit:   1000000000,
		BaseFee:    big.NewInt(3e9),
		Timestamp:  uint64(time.Now().Unix()),
	}
	b, err := c.MarshalBinary()
	require.NoError(t, err)
	require.NotNil(t, b)

	actualC := new(BLSMessage)
	err = actualC.UnmarshalBinary(b)
	require.NoError(t, err)
	require.EqualValues(t, actualC.ParentHash, c.ParentHash)
	require.EqualValues(t, actualC.Miner, c.Miner)
	require.EqualValues(t, actualC.Number, c.Number)
	require.EqualValues(t, actualC.GasLimit, c.GasLimit)
	require.EqualValues(t, actualC.BaseFee, c.BaseFee)
	require.EqualValues(t, actualC.Timestamp, c.Timestamp)
}

func TestNonBLSMessage(t *testing.T) {
	to := common.BigToAddress(big.NewInt(101))
	msg := L1Message{
		L1MessageTx: types.L1MessageTx{
			QueueIndex: 200,
			Gas:        500000,
			To:         &to,
			Value:      big.NewInt(3e9),
			Data:       []byte("0x1a2b3c"),
			Sender:     common.BigToAddress(big.NewInt(202)),
		},
		L1TxHash: common.BigToHash(big.NewInt(1111)),
	}
	c := NonBLSMessage{
		StateRoot:   common.BigToHash(big.NewInt(1111)),
		GasUsed:     50000000,
		ReceiptRoot: common.BigToHash(big.NewInt(2222)),
		LogsBloom:   []byte("0x1a2b3c4d"),
		L1Messages:  []L1Message{msg},
	}
	b, err := c.MarshalBinary()
	require.NoError(t, err)
	require.NotNil(t, b)

	actualC := new(NonBLSMessage)
	err = actualC.UnmarshalBinary(b)
	require.NoError(t, err)
	require.EqualValues(t, actualC.StateRoot, c.StateRoot)
	require.EqualValues(t, actualC.GasUsed, c.GasUsed)
	require.EqualValues(t, actualC.ReceiptRoot, c.ReceiptRoot)
	require.EqualValues(t, actualC.LogsBloom, c.LogsBloom)
	require.EqualValues(t, len(c.L1Messages), len(actualC.L1Messages))
	require.EqualValues(t, c.L1Messages[0].QueueIndex, actualC.L1Messages[0].QueueIndex)
	require.EqualValues(t, c.L1Messages[0].Gas, actualC.L1Messages[0].Gas)
	require.EqualValues(t, c.L1Messages[0].To.Bytes(), actualC.L1Messages[0].To.Bytes())
	require.EqualValues(t, c.L1Messages[0].Value, actualC.L1Messages[0].Value)
	require.EqualValues(t, c.L1Messages[0].Data, actualC.L1Messages[0].Data)
	require.EqualValues(t, c.L1Messages[0].Sender, actualC.L1Messages[0].Sender)
	require.EqualValues(t, c.L1Messages[0].L1TxHash, actualC.L1Messages[0].L1TxHash)
}
