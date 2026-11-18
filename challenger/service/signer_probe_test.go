package service

import (
	"context"
	"errors"
	"testing"
	"time"
)

// A failed probe is retried on the short interval, and a success returns to the long one.
func TestRunSignerProbeRetriesFasterAfterFailure(t *testing.T) {
	const interval, retryInterval = time.Hour, time.Millisecond

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := []error{errors.New("down"), errors.New("down"), nil}
	calls := make(chan int, len(results)+1)
	n := 0
	probe := func(context.Context) error {
		n++
		calls <- n
		if n > len(results) {
			return nil
		}
		return results[n-1]
	}

	done := make(chan struct{})
	go func() {
		runSignerProbe(ctx, probe, interval, retryInterval)
		close(done)
	}()

	// Both failures and the recovery arrive on the retry interval.
	for want := 1; want <= len(results); want++ {
		select {
		case got := <-calls:
			if got != want {
				t.Fatalf("probe call %d, want %d", got, want)
			}
		case <-time.After(5 * time.Second):
			t.Fatalf("probe %d was not retried on the short interval", want)
		}
	}

	// After the recovery the next probe waits for the long interval.
	select {
	case got := <-calls:
		t.Fatalf("probe call %d ran without waiting for the long interval", got)
	case <-time.After(50 * time.Millisecond):
	}

	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("runSignerProbe did not stop when the context ended")
	}
}
