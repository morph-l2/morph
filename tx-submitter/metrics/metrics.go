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
	RpcErrors             prometheus.Counter
	WalletBalance         prometheus.Gauge
	RollupCost            prometheus.Counter
	FinalizeCost          prometheus.Counter
	L1FeeCollection       prometheus.Counter
	IndexerBlockProcessed prometheus.Gauge
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
		RollupCost: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "submitter_rollup_cost",
			Help:      "Rollup cost",
			Namespace: metricsNamespace,
		}),
		FinalizeCost: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "submitter_finalize_cost",
			Help:      "Finalize cost",
			Namespace: metricsNamespace,
		}),
		IndexerBlockProcessed: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "submitter_indexer_block_processed",
			Help:      "Indexer block processed",
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

func (m *Metrics) SetWalletBalance(balance float64) {
	m.WalletBalance.Set(balance)
}

func (m *Metrics) IncRpcErrors() {
	m.RpcErrors.Inc()
}

func (m *Metrics) AddRollupCost(cost float64) {
	m.RollupCost.Add(cost)
}

func (m *Metrics) AddFinalizeCost(cost float64) {
	m.FinalizeCost.Add(cost)
}

func (m *Metrics) AddCollectedL1Fee(cost float64) {
	m.L1FeeCollection.Add(cost)
}

func (m *Metrics) SetIndexerBlockProcessed(blockNumber uint64) {
	m.IndexerBlockProcessed.Set(float64(blockNumber))
}
