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
	L1SyncHeight     metrics.Gauge
	RollupL2Height   metrics.Gauge
	DeriveL2Height   metrics.Gauge
	BatchStatus      metrics.Gauge
	LatestBatchIndex metrics.Gauge
	SyncedBatchIndex metrics.Gauge

	// LocalVerifyTriggered increments once per batch processed under
	// VerifyModeLocal -- presence/absence on dashboards confirms the local
	// verifier is running. Failure tracking is intentionally not split into
	// separate counters; failures surface as Error logs and propagate as
	// ErrBatchVerifyDivergence to BatchStatus=stateException.
	LocalVerifyTriggered metrics.Counter

	// Tag management metrics. SafeL2BlockNumber / FinalizedL2BlockNumber are
	// the canonical "where is the chain now" gauges; the counters track
	// transitions for rate-based alerts.
	SafeAdvanceTotal           metrics.Counter
	FinalizedAdvanceTotal      metrics.Counter
	SafeL2BlockNumber          metrics.Gauge
	FinalizedL2BlockNumber     metrics.Gauge
	L1ReorgResetTotal          metrics.Counter
	TagInvariantViolationTotal metrics.Counter
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
		LocalVerifyTriggered: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "local_verify_triggered_total",
			Help:      "Number of batches processed by the local-rebuild verifier.",
		}, labels).With(labelsAndValues...),
		SafeAdvanceTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "safe_advance_total",
			Help:      "Times derivation advanced the safe L2 head after a verified batch.",
		}, labels).With(labelsAndValues...),
		FinalizedAdvanceTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "finalized_advance_total",
			Help:      "Times the finalizer advanced the finalized L2 head from L1 finalized state.",
		}, labels).With(labelsAndValues...),
		SafeL2BlockNumber: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "safe_l2_block_number",
			Help:      "Current in-memory safe L2 block number (mirror of derivation tag advancer).",
		}, labels).With(labelsAndValues...),
		FinalizedL2BlockNumber: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "finalized_l2_block_number",
			Help:      "Current in-memory finalized L2 block number (mirror of derivation tag advancer).",
		}, labels).With(labelsAndValues...),
		L1ReorgResetTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "l1_reorg_reset_total",
			Help:      "Times an L1 reorg triggered a tag advancer reset (safe cleared, refilled by re-derivation).",
		}, labels).With(labelsAndValues...),
		TagInvariantViolationTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "tag_invariant_violation_total",
			Help:      "Times the finalized <= safe <= unsafe invariant failed; SetBlockTags is skipped on each occurrence.",
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

func (m *Metrics) IncLocalVerifyTriggered() {
	m.LocalVerifyTriggered.Add(1)
}

func (m *Metrics) IncSafeAdvance() {
	m.SafeAdvanceTotal.Add(1)
}

func (m *Metrics) IncFinalizedAdvance() {
	m.FinalizedAdvanceTotal.Add(1)
}

func (m *Metrics) SetSafeL2BlockNumber(n uint64) {
	m.SafeL2BlockNumber.Set(float64(n))
}

func (m *Metrics) SetFinalizedL2BlockNumber(n uint64) {
	m.FinalizedL2BlockNumber.Set(float64(n))
}

func (m *Metrics) IncL1ReorgReset() {
	m.L1ReorgResetTotal.Add(1)
}

func (m *Metrics) IncTagInvariantViolation() {
	m.TagInvariantViolationTotal.Add(1)
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
