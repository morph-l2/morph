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
	RpcErrors                           prometheus.Counter
	WalletBalance                       prometheus.Gauge
	LastFinalizedBatchIndex             prometheus.Gauge
	LastCommittedBatchIndex             prometheus.Gauge
	LastFinalizedCommitedBatchIndexDiff prometheus.Gauge
	L2BlockNumber                       prometheus.Gauge
	L2BlockNumberRolluped               prometheus.Gauge
	LastRollupedBlocknumberDiff         prometheus.Gauge
}

func NewMetrics() *Metrics {

	return &Metrics{
		RpcErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "submitter_rpc_errors",
			Help:      "Number of RPC errors encountered",
			Namespace: metricsNamespace,
		}),

		WalletBalance: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "submitter_wallet_balance",
			Help:      "Wallet balance",
			Namespace: metricsNamespace,
		}),
	}
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

func (m *Metrics) IncRpcErrors() {
	m.RpcErrors.Inc()
}

func (m *Metrics) SetWalletBalance(balance float64) {
	m.WalletBalance.Set(balance)
}
