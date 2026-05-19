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

	// SPEC-005 section 4.6 Path B counters. PathBTriggered increments once per batch
	// processed under VerifyModePathB; PathBFailed is the unlabelled total.
	// PathBFailedByKind carries a "kind" label so dashboards / alerts can split
	// failures by category (versioned hash mismatch vs local block missing vs
	// encoding error vs ...). Increment both via IncPathBFailedKind so the
	// total stays in sync with the sum across kinds.
	PathBTriggered    metrics.Counter
	PathBFailed       metrics.Counter
	PathBFailedByKind metrics.Counter

	// SPEC-005 §4.2 Hybrid fallback counter. Increments once per batch where
	// hybrid mode demoted to Path A because Path B reported local_block_missing.
	// Steady-state value should approach zero on followers that keep up via P2P;
	// a sustained non-zero rate indicates the local L2 sync lags the L1 commit
	// feed. Distinct from path_b_failed_by_kind_total{kind="local_block_missing"}
	// in that the latter still counts hybrid's underlying B attempt.
	HybridFallbackTotal metrics.Counter

	// SPEC-005 section 4.7 Tag management metrics. Replace the (previously absent)
	// blocktag instrumentation; on-call alerts should now key off these.
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
		PathBTriggered: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "path_b_triggered_total",
			Help:      "Number of batches verified via SPEC-005 Path B (local-rebuild).",
		}, labels).With(labelsAndValues...),
		PathBFailed: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "path_b_failed_total",
			Help:      "Path B failures: local block missing, encoding error, or versioned hash mismatch.",
		}, labels).With(labelsAndValues...),
		PathBFailedByKind: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "path_b_failed_by_kind_total",
			Help:      "Path B failures broken down by kind label (versioned_hash_mismatch, local_block_missing, ...).",
		}, append(append([]string(nil), labels...), "kind")).With(labelsAndValues...),
		HybridFallbackTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "verify_hybrid_fallback_total",
			Help:      "Times hybrid verify fell back to Path A because the local L2 chain hadn't synced the batch range yet (sync-lag backup).",
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

func (m *Metrics) IncPathBTriggered() {
	m.PathBTriggered.Add(1)
}

func (m *Metrics) IncHybridFallback() {
	m.HybridFallbackTotal.Add(1)
}

func (m *Metrics) IncPathBFailed() {
	m.PathBFailed.Add(1)
}

// IncPathBFailedKind increments both the unlabelled PathBFailed total and the
// PathBFailedByKind counter scoped to the given kind. Call sites in
// verify_path_b.go use this so the kind label and the total stay aligned.
func (m *Metrics) IncPathBFailedKind(kind string) {
	m.PathBFailed.Add(1)
	m.PathBFailedByKind.With("kind", kind).Add(1)
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
