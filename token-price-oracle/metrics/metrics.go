package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	// UpdateErrors update error count
	UpdateErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "update_errors_total",
			Help: "Total number of update errors",
		},
		[]string{"type"}, // error type: basefee, scalar, price
	)

	// AccountBalance tracks account balance in ETH
	AccountBalance = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "account_balance_eth",
			Help: "Account balance in ETH",
		},
	)

	// LastSuccessfulUpdateTimestamp records the Unix timestamp of the last successful update cycle
	// A successful update includes: prices updated on-chain OR prices skipped (below threshold)
	// This helps monitor if the oracle is running normally
	LastSuccessfulUpdateTimestamp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "last_successful_update_timestamp",
			Help: "Unix timestamp of the last successful price update cycle (includes both updates and skips)",
		},
	)

	// UpdatesTotal counts total number of successful update cycles
	UpdatesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "updates_total",
			Help: "Total number of successful update cycles",
		},
		[]string{"type"}, // type: "updated" or "skipped"
	)
)

// init registers all metrics
func init() {
	prometheus.MustRegister(UpdateErrors)
	prometheus.MustRegister(AccountBalance)
	prometheus.MustRegister(LastSuccessfulUpdateTimestamp)
	prometheus.MustRegister(UpdatesTotal)
}

// StartMetricsServer starts metrics HTTP server
func StartMetricsServer(address string) error {
	logrus.WithField("address", address).Info("Starting metrics server")

	http.Handle("/metrics", promhttp.Handler())

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return http.ListenAndServe(address, nil)
}
