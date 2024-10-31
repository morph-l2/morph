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
	RollupCostSum         prometheus.Counter
	FinalizeCostSum       prometheus.Counter
	RollupCost            prometheus.Gauge
	FinalizeCost          prometheus.Gauge
	CollectedL1FeeSum     prometheus.Counter
	CollectedL1Fee        prometheus.Gauge
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
		RollupCostSum: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "submitter_rollup_cost_sum",
			Help:      "Rollup cost",
			Namespace: metricsNamespace,
		}),
		FinalizeCostSum: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "submitter_finalize_cost_sum",
			Help:      "Finalize cost",
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
		CollectedL1Fee: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "submitter_collected_l1_fee",
			Help:      "Collected L1 fee for every batch",
			Namespace: metricsNamespace,
		}),
		CollectedL1FeeSum: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "submitter_collected_l1_fee_sum",
			Help:      "Collected L1 fee for all batches commited ",
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

func (m *Metrics) SetRollupCost(cost float64) {
	m.RollupCostSum.Add(cost)
	m.RollupCost.Set(cost)
}

func (m *Metrics) SetFinalizeCost(cost float64) {
	m.FinalizeCostSum.Add(cost)
	m.FinalizeCost.Set(cost)
}

func (m *Metrics) SetCollectedL1Fee(cost float64) {
	m.CollectedL1FeeSum.Add(cost)
	m.CollectedL1Fee.Set(cost)
}

func (m *Metrics) SetIndexerBlockProcessed(blockNumber uint64) {
	m.IndexerBlockProcessed.Set(float64(blockNumber))
}
