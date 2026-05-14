package types

import (
	"errors"
	"fmt"
	"testing"

	"github.com/morph-l2/go-ethereum"
)

// retryableError must classify ethereum.NotFound as permanent so that
// SPEC-005 Path B fails fast when a target L2 block has not yet been sealed
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
