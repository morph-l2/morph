package types

import (
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/core/types"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

func TestL1Message(t *testing.T) {
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
	bytes, err := rlp.EncodeToBytes(&msg)
	require.NoError(t, err)

	var actual L1Message
	err = rlp.DecodeBytes(bytes, &actual)
	require.NoError(t, err)

	require.EqualValues(t, msg.QueueIndex, actual.QueueIndex)
	require.EqualValues(t, msg.Gas, actual.Gas)
	require.EqualValues(t, msg.To.Bytes(), actual.To.Bytes())
	require.EqualValues(t, msg.Value, actual.Value)
	require.EqualValues(t, msg.Data, actual.Data)
	require.EqualValues(t, msg.Sender, actual.Sender)
	require.EqualValues(t, msg.L1TxHash, actual.L1TxHash)
}
