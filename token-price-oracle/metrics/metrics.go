package metrics

import (
	"net/http"
	"time"

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

	// UnresolvedTokens tracks how many active tokens no configured feed could price in
	// the last cycle. A cycle that resolves some tokens still counts as successful, so
	// this is the signal that a token is going stale while the oracle looks healthy.
	UnresolvedTokens = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "unresolved_tokens",
			Help: "Number of active tokens no price feed could resolve in the last update cycle",
		},
	)
)

// init registers all metrics
func init() {
	prometheus.MustRegister(UpdateErrors)
	prometheus.MustRegister(AccountBalance)
	prometheus.MustRegister(LastSuccessfulUpdateTimestamp)
	prometheus.MustRegister(UpdatesTotal)
	prometheus.MustRegister(UnresolvedTokens)

	// Initialize metrics with default values to avoid nil pointer issues in alerting systems
	// Set initial timestamp to current time (program start time)
	LastSuccessfulUpdateTimestamp.Set(float64(time.Now().Unix()))
	// Initialize counter labels to ensure they exist from the start
	// Must call Add(0) to actually create the metric, WithLabelValues alone doesn't create it
	UpdatesTotal.WithLabelValues("updated").Add(0)
	UpdatesTotal.WithLabelValues("skipped").Add(0)
	// Initialize error counter labels
	UpdateErrors.WithLabelValues("price").Add(0)
	UpdateErrors.WithLabelValues("unresolved_token").Add(0)
	UnresolvedTokens.Set(0)
	// Note: AccountBalance is NOT initialized here to avoid triggering low balance alerts
	// It will be set with the real value on the first update cycle
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
