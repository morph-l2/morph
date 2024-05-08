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

func (m *Metrics) Serve(hostname string, port uint64) (*http.Server, error) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := new(http.Server)
	srv.Addr = net.JoinHostPort(hostname, strconv.FormatUint(port, 10))
	srv.Handler = mux
	err := srv.ListenAndServe()
	return srv, err
}
