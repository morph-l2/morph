package types

import (
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
)

type L1Message struct {
	types.L1MessageTx
	L1TxHash common.Hash
}

type L1MessageReader interface {
	GetL1Message(index uint64, txHash common.Hash) (*L1Message, error)
	ReadL1MessagesInRange(start, end uint64) []L1Message
	LatestSynced() uint64
}
