package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	// L1BaseFee L1 base fee (Gwei)
	L1BaseFee = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "l1_base_fee",
		Help: "L1 base fee in Gwei",
	})

	// L1BaseFeeOnL2 L1 base fee recorded on L2 (Gwei)
	L1BaseFeeOnL2 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "l1_base_fee_on_l2",
		Help: "L1 base fee on L2 in Gwei",
	})

	// L1BlobBaseFeeOnL2 L1 blob base fee recorded on L2 (Gwei)
	L1BlobBaseFeeOnL2 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "l1_blob_base_fee_on_l2",
		Help: "L1 blob base fee on L2 in Gwei",
	})

	// CommitScalar commit scalar value
	CommitScalar = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "commit_scalar",
		Help: "Commit scalar value",
	})

	// BlobScalar blob scalar value
	BlobScalar = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "blob_scalar",
		Help: "Blob scalar value",
	})

	// TxnPerBatch transactions per batch
	TxnPerBatch = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "txn_per_batch",
		Help: "Transactions per batch",
	})

	// GasOracleOwnerBalance Oracle account balance (ETH)
	GasOracleOwnerBalance = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gas_oracle_owner_balance",
		Help: "Gas oracle owner balance in ETH",
	})

	// L1RPCStatus L1 RPC status (0=ok, 1=warning, 2=error)
	L1RPCStatus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "l1_rpc_status",
		Help: "L1 RPC status (0=ok, 1=warning, 2=error)",
	})

	// BaseFeeUpdateCount base fee update count
	BaseFeeUpdateCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "base_fee_update_count",
		Help: "Total number of base fee updates",
	})

	// ScalarUpdateCount scalar update count
	ScalarUpdateCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "scalar_update_count",
		Help: "Total number of scalar updates",
	})

	// UpdateErrors update error count
	UpdateErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "update_errors_total",
			Help: "Total number of update errors",
		},
		[]string{"type"}, // error type: basefee, scalar
	)
)

// init registers all metrics
func init() {
	prometheus.MustRegister(L1BaseFee)
	prometheus.MustRegister(L1BaseFeeOnL2)
	prometheus.MustRegister(L1BlobBaseFeeOnL2)
	prometheus.MustRegister(CommitScalar)
	prometheus.MustRegister(BlobScalar)
	prometheus.MustRegister(TxnPerBatch)
	prometheus.MustRegister(GasOracleOwnerBalance)
	prometheus.MustRegister(L1RPCStatus)
	prometheus.MustRegister(BaseFeeUpdateCount)
	prometheus.MustRegister(ScalarUpdateCount)
	prometheus.MustRegister(UpdateErrors)
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
