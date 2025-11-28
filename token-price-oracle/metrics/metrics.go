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
)

// init registers all metrics
func init() {
	prometheus.MustRegister(UpdateErrors)
	prometheus.MustRegister(AccountBalance)
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
