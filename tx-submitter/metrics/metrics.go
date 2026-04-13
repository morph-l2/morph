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
	WalletBalance           prometheus.Gauge
	RpcErrors               prometheus.Counter
	RollupCostSum           prometheus.Gauge
	FinalizeCostSum         prometheus.Gauge
	RollupCost              prometheus.Gauge
	FinalizeCost            prometheus.Gauge
	CollectedL1FeeSum       prometheus.Gauge
	IndexerBlockProcessed   prometheus.Gauge
	LastCommittedBatch      prometheus.Gauge
	LastFinalizedBatch      prometheus.Gauge
	HasPendingFinalizeBatch prometheus.Gauge
	LastCacheBatchIndex     prometheus.Gauge
	reorgs                  prometheus.Counter
	reorgDepthVal           uint64
	reorgCountVal           uint64
	confirmedTxs            *prometheus.CounterVec
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
		LastCommittedBatch: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_last_committed_batch",
			Help: "Latest batch committed by the submitter",
		}),
		LastFinalizedBatch: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_last_finalized_batch",
			Help: "Latest batch finalized by the submitter",
		}),
		LastCacheBatchIndex: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_last_batch_index",
			Help: "Latest batch index by the submitter",
		}),
		HasPendingFinalizeBatch: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "tx_submitter_has_pending_finalize_batch",
			Help: "Whether there are batches pending finalization (1 = yes, 0 = no)",
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

	// Register metrics with Prometheus
	// We use Register instead of MustRegister to avoid panics if metrics are already registered
	_ = prometheus.Register(m.WalletBalance)
	_ = prometheus.Register(m.RpcErrors)
	_ = prometheus.Register(m.RollupCostSum)
	_ = prometheus.Register(m.FinalizeCostSum)
	_ = prometheus.Register(m.RollupCost)
	_ = prometheus.Register(m.FinalizeCost)
	_ = prometheus.Register(m.CollectedL1FeeSum)
	_ = prometheus.Register(m.IndexerBlockProcessed)
	_ = prometheus.Register(m.LastCommittedBatch)
	_ = prometheus.Register(m.LastFinalizedBatch)
	_ = prometheus.Register(m.LastCacheBatchIndex)
	_ = prometheus.Register(m.HasPendingFinalizeBatch)
	_ = prometheus.Register(m.reorgs)
	_ = prometheus.Register(m.confirmedTxs)

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

// SetLastCommittedBatch sets the last committed batch index metric
func (m *Metrics) SetLastCommittedBatch(index uint64) {
	m.LastCommittedBatch.Set(float64(index))
}

// SetLastFinalizedBatch sets the last finalized batch index metric
func (m *Metrics) SetLastFinalizedBatch(index uint64) {
	m.LastFinalizedBatch.Set(float64(index))
}

// SetLastCacheBatchIndex sets the last batch index metric
func (m *Metrics) SetLastCacheBatchIndex(index uint64) {
	m.LastCacheBatchIndex.Set(float64(index))
}

// SetHasPendingFinalizeBatch sets whether there are batches pending finalization
// hasPending should be true if there are pending batches, false otherwise
func (m *Metrics) SetHasPendingFinalizeBatch(hasPending bool) {
	if hasPending {
		m.HasPendingFinalizeBatch.Set(1)
	} else {
		m.HasPendingFinalizeBatch.Set(0)
	}
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
	prometheus.Unregister(m.LastCommittedBatch)
	prometheus.Unregister(m.LastFinalizedBatch)
	prometheus.Unregister(m.HasPendingFinalizeBatch)
	prometheus.Unregister(m.reorgs)
	prometheus.Unregister(m.confirmedTxs)
}
