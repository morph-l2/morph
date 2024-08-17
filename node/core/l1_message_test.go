package node

import (
	"math/big"
	"os"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/node/types"
)

func TestValidateL1Messages(t *testing.T) {
	l1Reader := testL1MsgReader{}
	l1Messages := make([]types.L1Message, 10)
	collectedL1TxHashes := make([]common.Hash, 10)
	l1TxBytes := make([][]byte, 10)
	for i := 0; i < 10; i++ {
		to := common.BigToAddress(big.NewInt(1))
		l1Message := types.L1Message{
			L1TxHash: common.BigToHash(big.NewInt(int64(i))),
			L1MessageTx: eth.L1MessageTx{
				QueueIndex: uint64(i),
				Gas:        21000,
				To:         &to,
				Value:      big.NewInt(100),
				Sender:     common.BigToAddress(big.NewInt(int64(i))),
			},
		}
		l1Messages[i] = l1Message
		collectedL1TxHashes[i] = l1Message.L1TxHash
		l1Reader.addL1Message(l1Message)
		txByte, _ := eth.NewTx(&l1Message.L1MessageTx).MarshalBinary()
		l1TxBytes[i] = txByte
	}
	block := &catalyst.ExecutableL2Data{
		NextL1MessageIndex: 10,
		Transactions:       l1TxBytes,
	}

	t.Run("positive case", func(t *testing.T) {
		executor := Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    &l1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}
		require.NoError(t, executor.validateL1Messages(block, collectedL1TxHashes))
	})

	t.Run("constraint 1: unknown L1 message", func(t *testing.T) {
		thisL1Reader := l1Reader.copy()
		executor := Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    thisL1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}
		thisL1Reader.removeL1Message(5)
		err := executor.validateL1Messages(block, collectedL1TxHashes)
		require.EqualError(t, err, types.ErrUnknownL1Message.Error())

		thisL1Reader = l1Reader.copy()
		l1Message := l1Messages[6]
		l1Message.Gas = 30000
		thisL1Reader.addL1Message(l1Message)
		executor.l1MsgReader = thisL1Reader
		err2 := executor.validateL1Messages(block, collectedL1TxHashes)
		require.EqualError(t, err2, types.ErrUnknownL1Message.Error())
	})

	t.Run("constraint 2: L1 messages wrong order", func(t *testing.T) {
		thisL1Reader := l1Reader.copy()
		executor := Executor{
			nextL1MsgIndex: 1,
			l1MsgReader:    thisL1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}
		thisCollectedL1TxHashes := make([]common.Hash, len(collectedL1TxHashes))
		copy(thisCollectedL1TxHashes, collectedL1TxHashes)
		thisCollectedL1TxHashes[2] = common.BigToHash(big.NewInt(100))
		err := executor.validateL1Messages(block, thisCollectedL1TxHashes)
		require.EqualError(t, err, types.ErrIncorrectL1TxHash.Error())

		thisCollectedL1TxHashes = exchangeL1Msg(collectedL1TxHashes, 2, 5)
		executor = Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    thisL1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}
		err = executor.validateL1Messages(block, thisCollectedL1TxHashes)
		require.EqualError(t, err, types.ErrIncorrectL1TxHash.Error())
	})

	t.Run("constraint 3: L1 transactions wrong order", func(t *testing.T) {
		executor := Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    &l1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}

		l1TxBytesCopy := exchangeL1Txs(block.Transactions, 3, 4)
		thisBlock := &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 10,
			Transactions:       l1TxBytesCopy,
		}
		err := executor.validateL1Messages(thisBlock, collectedL1TxHashes)
		require.EqualError(t, err, types.ErrInvalidL1MessageOrder.Error())
	})

	t.Run("constraint 4: l1 transaction validation", func(t *testing.T) {
		executor := Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    &l1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}
		// add a new L1 tx that not belongs to the collected l1 messages
		to := common.BigToAddress(big.NewInt(1))
		addedL1Tx := eth.L1MessageTx{
			QueueIndex: uint64(10),
			Gas:        21000,
			To:         &to,
			Value:      big.NewInt(100),
			Sender:     common.BigToAddress(big.NewInt(int64(10))),
		}
		addedL1TxBytes, _ := eth.NewTx(&addedL1Tx).MarshalBinary()
		l1TxBytesCopy := append(block.Transactions, addedL1TxBytes)
		thisBlock := &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 11,
			Transactions:       l1TxBytesCopy,
		}

		err := executor.validateL1Messages(thisBlock, collectedL1TxHashes)
		require.EqualError(t, err, types.ErrUnknownL1Message.Error())
	})

	t.Run("constraint 5: no l1 tx after l2 tx", func(t *testing.T) {
		executor := Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    &l1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}
		//block.Transactions
		l2tx := eth.NewTx(&eth.LegacyTx{
			Nonce:    1,
			GasPrice: big.NewInt(1000000000),
			Gas:      21000,
			To:       nil,
			Value:    big.NewInt(2000),
			Data:     nil,
		})
		l2txBz, _ := l2tx.MarshalBinary()
		txBytes := make([][]byte, 0)
		txBytes = append(append(append(txBytes, block.Transactions[:2]...), l2txBz), block.Transactions[2:]...)
		thisBlock := &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 10,
			Transactions:       txBytes,
		}

		err := executor.validateL1Messages(thisBlock, collectedL1TxHashes)
		require.ErrorIs(t, err, types.ErrInvalidL1MessageOrder)
	})

	t.Run("constraint 6: testing block.NextL1MessageIndex", func(t *testing.T) {
		to := common.BigToAddress(big.NewInt(1))
		skippedL1WithIndex10 := eth.L1MessageTx{
			QueueIndex: uint64(10),
			Gas:        21000,
			To:         &to,
			Value:      big.NewInt(100),
			Sender:     common.BigToAddress(big.NewInt(int64(10))),
		}
		l1Message10 := types.L1Message{
			L1TxHash:    common.BigToHash(big.NewInt(int64(10))),
			L1MessageTx: skippedL1WithIndex10,
		}
		thisL1Reader := l1Reader.copy()
		thisL1Reader.addL1Message(l1Message10)
		executor := Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    thisL1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}
		block := &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 11,
			Transactions:       l1TxBytes,
			SkippedTxs: []*eth.SkippedTransaction{
				{Tx: eth.NewTx(&skippedL1WithIndex10)},
			},
		}
		collectedL1TxHashesCopy := append(collectedL1TxHashes, common.BigToHash(big.NewInt(int64(10))))
		err := executor.validateL1Messages(block, collectedL1TxHashesCopy)
		require.NoError(t, err)

		block = &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 9,
			Transactions:       l1TxBytes,
		}
		err = executor.validateL1Messages(block, collectedL1TxHashes)
		require.ErrorIs(t, err, types.ErrWrongNextL1MessageIndex)
	})

	t.Run("constraint 7: invalid skipped L1 messages", func(t *testing.T) {
		executor := Executor{
			nextL1MsgIndex: 0,
			l1MsgReader:    &l1Reader,
			logger:         tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		}

		originTxs := block.Transactions[:]
		l1TxBytes := make([][]byte, 0)
		l1TxBytes = append(append(l1TxBytes, originTxs[:2]...), originTxs[3:]...)

		thisBlock := &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 10,
			Transactions:       l1TxBytes,
		}

		err := executor.validateL1Messages(thisBlock, collectedL1TxHashes)
		require.EqualError(t, err, types.ErrInvalidSkippedL1Message.Error())

		skippedTx := new(eth.Transaction)
		err = skippedTx.UnmarshalBinary(originTxs[2])
		require.NoError(t, err)
		thisBlock = &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 10,
			Transactions:       l1TxBytes,
			SkippedTxs: []*eth.SkippedTransaction{
				{Tx: skippedTx},
			},
		}
		err = executor.validateL1Messages(thisBlock, collectedL1TxHashes)
		require.NoError(t, err)

		thisBlock = &catalyst.ExecutableL2Data{
			NextL1MessageIndex: 10,
			Transactions:       l1TxBytes,
			SkippedTxs: []*eth.SkippedTransaction{
				{Tx: skippedTx},
			},
		}
		err = executor.validateL1Messages(thisBlock, collectedL1TxHashes)
		require.NoError(t, err)
	})
}

var _ types.L1MessageReader = (*testL1MsgReader)(nil)

type testL1MsgReader struct {
	storedL1Msgs map[uint64]*types.L1Message
}

func (r *testL1MsgReader) copy() *testL1MsgReader {
	copiedMap := make(map[uint64]*types.L1Message)
	for k, v := range r.storedL1Msgs {
		newV := types.L1Message{
			L1MessageTx: eth.L1MessageTx{
				QueueIndex: v.QueueIndex,
				Gas:        v.Gas,
				To:         v.To,
				Value:      v.Value,
				Data:       v.Data,
				Sender:     v.Sender,
			},
			L1TxHash: v.L1TxHash,
		}
		copiedMap[k] = &newV
	}
	return &testL1MsgReader{
		storedL1Msgs: copiedMap,
	}
}

func (r *testL1MsgReader) addL1Message(l1Message types.L1Message) {
	if r.storedL1Msgs == nil {
		r.storedL1Msgs = make(map[uint64]*types.L1Message)
	}
	r.storedL1Msgs[l1Message.QueueIndex] = &l1Message
}

func (r *testL1MsgReader) removeL1Message(index uint64) {
	if r.storedL1Msgs != nil {
		delete(r.storedL1Msgs, index)
	}
}

func (r *testL1MsgReader) GetL1Message(index uint64, txHash common.Hash) (*types.L1Message, error) {
	return r.storedL1Msgs[index], nil
}

func (r *testL1MsgReader) ReadL1MessagesInRange(start, end uint64) []types.L1Message {
	return nil
}

func (r *testL1MsgReader) LatestSynced() uint64 {
	return 0
}

func exchangeL1Msg(origin []common.Hash, a, b int) []common.Hash {
	after := make([]common.Hash, len(origin))
	for i, msg := range origin {
		switch i {
		case a:
			after[b] = msg
		case b:
			after[a] = msg
		default:
			after[i] = msg
		}
	}
	return after
}

func exchangeL1Txs(origin [][]byte, a, b int) [][]byte {
	after := make([][]byte, len(origin))
	for i, msg := range origin {
		switch i {
		case a:
			after[b] = msg
		case b:
			after[a] = msg
		default:
			after[i] = msg
		}
	}
	return after
}
