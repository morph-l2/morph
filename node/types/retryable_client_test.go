package types

import (
	"errors"
	"fmt"
	"testing"

	"github.com/morph-l2/go-ethereum"
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

// retryableError must classify ethereum.NotFound as permanent so that
// SPEC-005 local verify fails fast when a target L2 block has not yet been sealed
// locally (snapshot too old or P2P sync still catching up). Without this
// classification the caller blocks for the full 30-minute backoff budget
// before the gap is surfaced.
func TestRetryableError_NotFoundIsPermanent(t *testing.T) {
	if retryableError(ethereum.NotFound) {
		t.Fatal("ethereum.NotFound must be non-retryable")
	}
	// Wrapped errors must be unwrapped via errors.Is so go-ethereum's
	// fmt.Errorf("...: %w", ethereum.NotFound) wrappers also classify.
	wrapped := fmt.Errorf("BlockByNumber: %w", ethereum.NotFound)
	if retryableError(wrapped) {
		t.Fatal("wrapped ethereum.NotFound must be non-retryable")
	}
}

func TestRetryableError_DiscontinuousBlockIsPermanent(t *testing.T) {
	err := errors.New("discontinuous block number: ...")
	if retryableError(err) {
		t.Fatal("DiscontinuousBlockError must be non-retryable")
	}
}

func TestRetryableError_GenericErrorIsRetryable(t *testing.T) {
	cases := []error{
		errors.New("connection refused"),
		errors.New("EOF"),
		errors.New("i/o timeout"),
		errors.New("502 Bad Gateway"),
	}
	for _, e := range cases {
		if !retryableError(e) {
			t.Errorf("expected retryable for %q", e)
		}
	}
}
