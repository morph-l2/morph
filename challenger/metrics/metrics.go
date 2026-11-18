package metrics

import (
	"net/http"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Challenges counts challenge HTTP outcomes.
var Challenges = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "challenger_requests_total",
		Help: "Challenge requests by result",
	},
	[]string{"result"},
)

// Comparisons counts root comparisons by verdict.
var Comparisons = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "challenger_batch_comparisons_total",
		Help: "Batch root comparisons by verdict",
	},
	[]string{"verdict"},
)

// signerUp is 1 when the last remote signer probe succeeded and 0 when it failed.
var signerUp = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "challenger_ext_signer_up",
		Help: "Whether the last remote signer probe succeeded (1) or failed (0)",
	},
)

// signerLastSuccess is the unix time of the last successful remote signer probe.
var signerLastSuccess = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "challenger_ext_signer_last_success_timestamp_seconds",
		Help: "Unix time of the last successful remote signer probe",
	},
)

// RecordSignerProbe records the result of one remote signer probe.
func RecordSignerProbe(ok bool) {
	if !ok {
		signerUp.Set(0)
		return
	}
	signerUp.Set(1)
	signerLastSuccess.SetToCurrentTime()
}

func init() {
	prometheus.MustRegister(Challenges)
	prometheus.MustRegister(Comparisons)
	prometheus.MustRegister(signerUp)
	prometheus.MustRegister(signerLastSuccess)
	for _, result := range []string{"success", "rejected", "busy", "pending", "reverted", "error", "bad_request", "unauthorized"} {
		Challenges.WithLabelValues(result).Add(0)
	}
	for _, verdict := range []string{"match", "mismatch", "inconclusive"} {
		Comparisons.WithLabelValues(verdict).Add(0)
	}
}

// StartServer serves Prometheus metrics and a health check.
func StartServer(address string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})
	log.Info("Starting metrics server", "address", address)
	return http.ListenAndServe(address, mux)
}
