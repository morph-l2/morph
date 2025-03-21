package services

import (
	"context"
	"fmt"
	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/eth"
)

type BatchFetcher struct {
	l2Clients []iface.L2Client
}

func NewBatchFetcher(l2Clients []iface.L2Client) *BatchFetcher {
	return &BatchFetcher{
		l2Clients: l2Clients,
	}
}

func (bf *BatchFetcher) GetRollupBatchByIndex(index uint64) (*eth.RPCRollupBatch, error) {
	// Try each L2 client until we get a successful response
	var lastErr error
	for _, client := range bf.l2Clients {
		batch, err := client.GetRollupBatchByIndex(context.Background(), index)
		if err != nil {
			lastErr = err
			continue
		}
		// Validate that batch exists and has signatures before returning
		if batch != nil && len(batch.Signatures) > 0 {
			return batch, nil
		}
	}
	if lastErr != nil {
		return nil, fmt.Errorf("failed to get batch %d from any L2 client: %w", index, lastErr)
	}
	return nil, fmt.Errorf("batch %d not found in any L2 client", index)
}
