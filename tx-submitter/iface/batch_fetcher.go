package iface

import "github.com/morph-l2/go-ethereum/eth"

// BatchFetcher defines the interface for fetching batch data from nodes
type BatchFetcher interface {
	GetRollupBatchByIndex(index uint64) (*eth.RPCRollupBatch, error)
}
