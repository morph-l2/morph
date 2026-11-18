package service

import (
	"context"
	"time"

	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/challenger/metrics"
)

// signerProbeInterval is how often the remote signer is probed.
const signerProbeInterval = 10 * time.Minute

// RunSignerProbe checks on a fixed interval that the remote signer can still sign
// a challengeState transaction, starting immediately. The signer is otherwise
// only used when a batch is challenged, so a broken signer would go unnoticed
// until then. Nothing is broadcast. It returns when ctx is done, or at once when
// local signing is used.
func (s *Service) RunSignerProbe(ctx context.Context) {
	signer := s.l1.GetSigner()
	if !s.l1.IsExternalSign() || signer == nil {
		return
	}

	ticker := time.NewTicker(signerProbeInterval)
	defer ticker.Stop()
	for {
		if err := signer.Probe(ctx, s.rollupAddr, s.probeData, s.methodSig); err != nil {
			if ctx.Err() != nil {
				return
			}
			metrics.SignerUp.Set(0)
			log.Error("Remote signer health check failed", "method_sig", s.methodSig, "err", err)
		} else {
			metrics.SignerUp.Set(1)
			metrics.SignerLastSuccess.SetToCurrentTime()
			log.Info("Remote signer health check succeeded", "method_sig", s.methodSig)
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
