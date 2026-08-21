package services

import (
	"testing"

	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"

	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/mock"
)

func TestBatchFetcherAcceptsBatchWithoutSignatures(t *testing.T) {
	client := mock.NewL2ClientWrapper()
	client.RollupBatch = &eth.RPCRollupBatch{Version: 1}

	got, err := NewBatchFetcher([]iface.L2Client{client}).GetRollupBatchByIndex(7)
	require.NoError(t, err)
	require.Same(t, client.RollupBatch, got)
}
