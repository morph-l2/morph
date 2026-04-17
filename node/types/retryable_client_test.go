package types

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRetryableError(t *testing.T) {
	cases := []struct {
		name      string
		err       error
		retryable bool
	}{
		{
			name:      "nil-safe (transient connection refused)",
			err:       errors.New("dial tcp 127.0.0.1:8551: connect: connection refused"),
			retryable: true,
		},
		{
			name:      "miner closed (transient)",
			err:       errors.New(MinerClosed),
			retryable: true,
		},
		{
			name:      "discontinuous block (permanent)",
			err:       fmt.Errorf("cannot new block with %s 11, expected 12", DiscontinuousBlockError),
			retryable: false,
		},
		{
			name:      "wrong block number (permanent)",
			err:       fmt.Errorf("%s: expected 5, got 9", WrongBlockNumberError),
			retryable: false,
		},
		{
			name:      "parent not found (permanent)",
			err:       fmt.Errorf("%s: 0xdeadbeef", ParentNotFoundError),
			retryable: false,
		},
		{
			name:      "block hash mismatch (permanent, security)",
			err:       fmt.Errorf("%s: declared 0xaaa, computed 0xbbb", BlockHashMismatchError),
			retryable: false,
		},
		{
			name:      "invalid NextL1MsgIndex (permanent, security)",
			err:       fmt.Errorf("%s at #100 0xabc: header=99, computed=42", InvalidNextL1MsgIndexError),
			retryable: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.retryable, retryableError(tc.err))
		})
	}
}
