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

func init() {
	prometheus.MustRegister(Challenges)
	for _, result := range []string{"success", "rejected", "busy", "pending", "reverted", "error", "bad_request", "unauthorized"} {
		Challenges.WithLabelValues(result).Add(0)
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
