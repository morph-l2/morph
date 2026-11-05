package iface

import (
	"context"
	"errors"
	"testing"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

type fakeL1 struct {
	L1Client
	receipt *types.Receipt
	err     error
	calls   int
}

func (f *fakeL1) TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error) {
	f.calls++
	return f.receipt, f.err
}

func TestL1ClientsFailover(t *testing.T) {
	want := &types.Receipt{BlockNumber: nil, Status: 1}
	bad := &fakeL1{err: errors.New("rpc down")}
	empty := &fakeL1{err: ethereum.NotFound}
	nilRes := &fakeL1{}
	good := &fakeL1{receipt: want}

	c, err := NewL1Clients([]L1Client{bad, empty, nilRes, good})
	require.NoError(t, err)
	got, err := c.TransactionReceipt(context.Background(), common.Hash{})
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, 1, bad.calls)
	require.Equal(t, 1, empty.calls)
	require.Equal(t, 1, nilRes.calls)

	// all NotFound -> NotFound preserved
	c, _ = NewL1Clients([]L1Client{&fakeL1{err: ethereum.NotFound}, &fakeL1{err: ethereum.NotFound}})
	_, err = c.TransactionReceipt(context.Background(), common.Hash{})
	require.ErrorIs(t, err, ethereum.NotFound)

	_, err = NewL1Clients(nil)
	require.Error(t, err)
}
