package metrics

import (
	"net"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const metricsNamespace = "submitter"

type Metrics struct {
	RpcErrors               prometheus.Counter
	WalletBalance           prometheus.Gauge
	LastFinalizedBatchIndex prometheus.Gauge
	LastCommittedBatchIndex prometheus.Gauge
	L2BlockNumber           prometheus.Gauge
	L2BlockNumberRolluped   prometheus.Gauge
}

func NewMetrics() *Metrics {

	return &Metrics{
		RpcErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "submitter_rpc_errors",
			Help:      "Number of RPC errors encountered",
			Namespace: metricsNamespace,
		}),
		LastFinalizedBatchIndex: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "last_finalized_batch_index",
			Help:      "Last finalized batch index",
			Namespace: metricsNamespace,
		}),
		LastCommittedBatchIndex: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "last_committed_batch_index",
			Help:      "Last committed batch index",
			Namespace: metricsNamespace,
		}),
		L2BlockNumber: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "l2_block_number",
			Help:      "L2 block number",
			Namespace: metricsNamespace,
		}),
		L2BlockNumberRolluped: promauto.NewGauge(prometheus.GaugeOpts{

			Name:      "l2_block_number_rolluped",
			Help:      "L2 block number rolluped",
			Namespace: metricsNamespace,
		}),
		WalletBalance: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "wallet_balance",
			Help:      "Wallet balance",
			Namespace: metricsNamespace,
		}),
	}
}

func (m *Metrics) IncRpcErrors() {
	m.RpcErrors.Inc()
}
func (m *Metrics) Serve(hostname string, port uint64) (*http.Server, error) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := new(http.Server)
	srv.Addr = net.JoinHostPort(hostname, strconv.FormatUint(port, 10))
	srv.Handler = mux
	err := srv.ListenAndServe()
	return srv, err
}

func (m *Metrics) SetLastFinalizedBatchIndex(lastFinalized uint64) {
	m.LastFinalizedBatchIndex.Set(float64(lastFinalized))
}

func (m *Metrics) SetLastCommittedBatchIndex(lastCommitted uint64) {
	m.LastCommittedBatchIndex.Set(float64(lastCommitted))
}

func (m *Metrics) SetL2BlockNumber(l2BlockNumber uint64) {
	m.L2BlockNumber.Set(float64(l2BlockNumber))
}

func (m *Metrics) SetLastL2BlockNumberRolluped(l2BlockNumberRolluped uint64) {
	m.L2BlockNumberRolluped.Set(float64(l2BlockNumberRolluped))
}

func (m *Metrics) SetWalletBalance(balance float64) {
	m.WalletBalance.Set(balance)
}
