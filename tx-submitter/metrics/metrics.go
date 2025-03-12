package metrics

import (
	"net"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics represents the metrics collection for the tx-submitter
type Metrics struct {
	WalletBalance         prometheus.Gauge
	RpcErrors             prometheus.Counter
	RollupCostSum         prometheus.Gauge
	FinalizeCostSum       prometheus.Gauge
	RollupCost            prometheus.Gauge
	FinalizeCost          prometheus.Gauge
	CollectedL1FeeSum     prometheus.Gauge
	IndexerBlockProcessed prometheus.Gauge
	reorgs                prometheus.Counter
	reorgDepthVal         uint64
	reorgCountVal         uint64
	confirmedTxs          *prometheus.CounterVec
}

// NewMetrics creates a new Metrics instance
func NewMetrics() *Metrics {
	m := &Metrics{
		WalletBalance: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_wallet_balance",
			Help: "Current balance of the submitter wallet in ETH",
		}),
		RpcErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "tx_submitter_rpc_errors_total",
			Help: "Total number of RPC errors encountered",
		}),
		RollupCostSum: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_rollup_cost_sum",
			Help: "Total cost of rollup transactions in ETH",
		}),
		FinalizeCostSum: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_finalize_cost_sum",
			Help: "Total cost of finalize transactions in ETH",
		}),
		RollupCost: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_rollup_cost",
			Help: "Current rollup transaction cost in ETH",
		}),
		FinalizeCost: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_finalize_cost",
			Help: "Current finalize transaction cost in ETH",
		}),
		CollectedL1FeeSum: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_collected_l1_fee_sum",
			Help: "Total L1 fees collected in ETH",
		}),
		IndexerBlockProcessed: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_indexer_block_processed",
			Help: "Latest block number processed by the indexer",
		}),
		reorgs: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "tx_submitter_reorgs_total",
			Help: "Total number of chain reorganizations detected",
		}),
		confirmedTxs: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "tx_submitter_confirmed_txs_total",
				Help: "Total number of confirmed transactions by type",
			},
			[]string{"type"},
		),
	}

	return m
}

// SetWalletBalance sets the wallet balance metric
func (m *Metrics) SetWalletBalance(balance float64) {
	m.WalletBalance.Set(balance)
}

// IncRpcErrors increments the RPC error counter
func (m *Metrics) IncRpcErrors() {
	m.RpcErrors.Inc()
}

// SetRollupCost sets the rollup cost metrics
func (m *Metrics) SetRollupCost(cost float64) {
	m.RollupCostSum.Set(cost)
	m.RollupCost.Set(cost)
}

// SetFinalizeCost sets the finalize cost metrics
func (m *Metrics) SetFinalizeCost(cost float64) {
	m.FinalizeCostSum.Set(cost)
	m.FinalizeCost.Set(cost)
}

// SetCollectedL1Fee sets the collected L1 fee metric
func (m *Metrics) SetCollectedL1Fee(cost float64) {
	m.CollectedL1FeeSum.Set(cost)
}

// SetIndexerBlockProcessed sets the indexer block processed metric
func (m *Metrics) SetIndexerBlockProcessed(blockNumber uint64) {
	m.IndexerBlockProcessed.Set(float64(blockNumber))
}

// IncReorgs increments the reorg counter
func (m *Metrics) IncReorgs() {
	atomic.AddUint64(&m.reorgCountVal, 1)
	m.reorgs.Inc()
}

// SetReorgDepth sets the reorg depth metric
func (m *Metrics) SetReorgDepth(depth float64) {
	atomic.StoreUint64(&m.reorgDepthVal, uint64(depth))
}

// GetReorgDepth gets the current reorg depth
func (m *Metrics) GetReorgDepth() float64 {
	return float64(atomic.LoadUint64(&m.reorgDepthVal))
}

// GetReorgCount gets the current reorg count
func (m *Metrics) GetReorgCount() float64 {
	return float64(atomic.LoadUint64(&m.reorgCountVal))
}

// IncTxConfirmed increments the confirmed transaction counter for a given type
func (m *Metrics) IncTxConfirmed(txType string) {
	m.confirmedTxs.WithLabelValues(txType).Inc()
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

// UnregisterMetrics unregisters all metrics from the default registry
func (m *Metrics) UnregisterMetrics() {
	prometheus.Unregister(m.WalletBalance)
	prometheus.Unregister(m.RpcErrors)
	prometheus.Unregister(m.RollupCostSum)
	prometheus.Unregister(m.FinalizeCostSum)
	prometheus.Unregister(m.RollupCost)
	prometheus.Unregister(m.FinalizeCost)
	prometheus.Unregister(m.CollectedL1FeeSum)
	prometheus.Unregister(m.IndexerBlockProcessed)
	prometheus.Unregister(m.reorgs)
	prometheus.Unregister(m.confirmedTxs)
}
