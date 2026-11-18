package service

import (
	"context"
	"time"

	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/challenger/metrics"
)

const (
	// signerProbeInterval is how often a healthy remote signer is probed.
	signerProbeInterval = 10 * time.Minute
	// signerProbeRetryInterval is how often it is probed after a failed probe.
	signerProbeRetryInterval = 30 * time.Second
)

// RunSignerProbe checks that the remote signer can still sign a challengeState
// transaction, starting immediately. The signer is otherwise only used when a
// batch is challenged, so a broken signer would go unnoticed until then.
// Nothing is broadcast. It returns when ctx is done, or at once when local
// signing is used.
func (s *Service) RunSignerProbe(ctx context.Context) {
	signer := s.l1.GetSigner()
	if !s.l1.IsExternalSign() || signer == nil {
		return
	}
	probe := func(ctx context.Context) error {
		return signer.Probe(ctx, s.rollupAddr, s.probeData, s.methodSig)
	}
	runSignerProbe(ctx, probe, signerProbeInterval, signerProbeRetryInterval)
}

// runSignerProbe probes every interval while the signer is healthy and every
// retryInterval after a failure, until a probe succeeds again.
func runSignerProbe(ctx context.Context, probe func(context.Context) error, interval, retryInterval time.Duration) {
	var (
		downSince time.Time
		failures  int
	)
	for {
		err := probe(ctx)
		if ctx.Err() != nil {
			return
		}

		wait := interval
		switch {
		case err != nil:
			if failures == 0 {
				downSince = time.Now()
			}
			failures++
			wait = retryInterval
			metrics.RecordSignerProbe(false)
			log.Error("Remote signer health check failed",
				"err", err, "failed_probes", failures, "retry_in", retryInterval)
		case failures > 0:
			metrics.RecordSignerProbe(true)
			log.Warn("!!! Remote signer RECOVERED, signing works again !!!",
				"down_for", time.Since(downSince).Round(time.Second), "failed_probes", failures)
			failures = 0
		default:
			metrics.RecordSignerProbe(true)
			log.Info("Remote signer health check succeeded")
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(wait):
		}
	}
}
