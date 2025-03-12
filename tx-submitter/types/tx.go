package types

import (
	"github.com/morph-l2/go-ethereum/core/types"
)

// TxRecord represents a transaction record with metadata
type TxRecord struct {
	Tx         *types.Transaction
	SendTime   uint64
	QueryTimes uint64
	Confirmed  bool // Track if transaction has been confirmed in a block
} 