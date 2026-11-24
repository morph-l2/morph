package service

import (
	"context"
	"fmt"
	"time"

	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/challenger/client"
	"morph-l2/challenger/metrics"
)

// signerProbeInterval is how often the remote signer is probed, as in the prover.
const signerProbeInterval = 10 * time.Minute

// RunSignerProbe checks that the remote signer can still sign a challengeState
// transaction, starting immediately. The signer is otherwise only used when a
// batch is challenged, so a broken signer would go unnoticed until then.
// Nothing is broadcast. It returns when ctx is done, or at once when local
// signing is used.
func (s *Service) RunSignerProbe(ctx context.Context) {
	if !s.l1.IsExternalSign() || s.l1.GetSigner() == nil {
		return
	}
	runSignerProbe(ctx, s.probeSigner, signerProbeInterval)
}

// probeSigner has the remote signer sign a challengeState transaction that
// carries the real deposit, nonce and fees, then drops it before broadcast.
func (s *Service) probeSigner(ctx context.Context) error {
	// The deadline also covers the L1 read, so a hung L1 cannot stall the loop.
	ctx, cancel := context.WithTimeout(ctx, client.ProbeTimeout)
	defer cancel()

	deposit, err := s.view.ChallengeDeposit(ctx)
	if err != nil {
		return fmt.Errorf("challengeDeposit: %w", err)
	}
	return s.l1.GetSigner().Probe(
		ctx,
		s.l1.GetClient(),
		s.rollupAddr,
		deposit,
		s.probeData,
		s.methodSig,
		s.l1.GetMaxGasTipCap(),
		s.l1.GetMaxGasFeeCap(),
	)
}

// runSignerProbe probes once at start and then every interval until ctx is done.
func runSignerProbe(ctx context.Context, probe func(context.Context) error, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		err := probe(ctx)
		if ctx.Err() != nil {
			return
		}
		metrics.RecordSignerProbe(err == nil)
		if err != nil {
			log.Error("Remote signer health check failed", "err", err)
		} else {
			log.Info("Remote signer health check succeeded")
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
