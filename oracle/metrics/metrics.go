package metrics

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
	metricsSubsystem = "staking_oracle"
)

type Metrics struct {
	RewardEpoch metrics.Gauge
	RollupEpoch metrics.Gauge
	BatchEpoch  metrics.Gauge
}

func NewMetrics(namespace string, labelsAndValues ...string) *Metrics {
	var labels []string
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		RewardEpoch: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "reward_epoch",
			Help:      "",
		}, labels).With(labelsAndValues...),
		RollupEpoch: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "rollup_epoch",
			Help:      "",
		}, labels).With(labelsAndValues...),
		BatchEpoch: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "batch_epoch",
			Help:      "",
		}, labels).With(labelsAndValues...),
	}
}

func (m *Metrics) SetRewardEpoch(index uint64) {
	m.RewardEpoch.Set(float64(index))
}

func (m *Metrics) SetRollupEpoch(index uint64) {
	m.RollupEpoch.Set(float64(index))
}

func (m *Metrics) SetBatchEpoch(index uint64) {
	m.BatchEpoch.Set(float64(index))
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
