package derivation

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	metricsSubsystem = "derivation"
	stateNormal      = 0
	stateException   = 1
)

type Metrics struct {
	L1SyncHeight       metrics.Gauge
	RollupL2Height     metrics.Gauge
	DeriveL2Height     metrics.Gauge
	BatchStatus        metrics.Gauge
	LatestBatchIndex   metrics.Gauge
	SyncedBatchIndex   metrics.Gauge
	ReorgCount         metrics.Counter
	RollbackCount      metrics.Counter
	BlockMismatchCount metrics.Counter
	Halted             metrics.Gauge
}

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	var labels []string
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		L1SyncHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "l1_sync_height",
			Help:      "",
		}, labels).With(labelsAndValues...),
		RollupL2Height: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "rollup_l2_height",
			Help:      "",
		}, labels).With(labelsAndValues...),
		DeriveL2Height: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "derive_l2_height",
			Help:      "",
		}, labels).With(labelsAndValues...),
		BatchStatus: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "batch_root_exception",
			Help:      "",
		}, labels).With(labelsAndValues...),
		LatestBatchIndex: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "latest_batch_index",
			Help:      "",
		}, labels).With(labelsAndValues...),
		SyncedBatchIndex: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "synced_batch_index",
			Help:      "",
		}, labels).With(labelsAndValues...),
		ReorgCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "l1_reorg_detected_total",
			Help:      "Total number of L1 reorgs detected",
		}, labels).With(labelsAndValues...),
		RollbackCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "l2_rollback_total",
			Help:      "Total number of L2 chain rollbacks triggered by batch mismatch",
		}, labels).With(labelsAndValues...),
		BlockMismatchCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "block_mismatch_total",
			Help:      "Total number of block context mismatches detected during verification",
		}, labels).With(labelsAndValues...),
		Halted: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "halted",
			Help:      "Set to 1 when derivation is halted due to unrecoverable batch mismatch requiring manual intervention",
		}, labels).With(labelsAndValues...),
	}
}

func (m *Metrics) SetL1SyncHeight(height uint64) {
	m.L1SyncHeight.Set(float64(height))
}

func (m *Metrics) SetL2DeriveHeight(height uint64) {
	m.DeriveL2Height.Set(float64(height))
}

func (m *Metrics) SetRollupL2Height(height uint64) {
	m.RollupL2Height.Set(float64(height))
}

func (m *Metrics) SetBatchStatus(status uint64) {
	m.BatchStatus.Set(float64(status))
}

func (m *Metrics) SetLatestBatchIndex(batchIndex uint64) {
	m.LatestBatchIndex.Set(float64(batchIndex))
}

func (m *Metrics) SetSyncedBatchIndex(batchIndex uint64) {
	m.SyncedBatchIndex.Set(float64(batchIndex))
}

func (m *Metrics) IncReorgCount() {
	m.ReorgCount.Add(1)
}

func (m *Metrics) IncRollbackCount() {
	m.RollbackCount.Add(1)
}

func (m *Metrics) IncBlockMismatchCount() {
	m.BlockMismatchCount.Add(1)
}

func (m *Metrics) SetHalted() {
	m.Halted.Set(1)
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
