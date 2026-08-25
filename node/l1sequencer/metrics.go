package l1sequencer

import (
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// metricsSubsystem groups the L1 tracker metrics under morphnode_l1tracker_*.
const metricsSubsystem = "l1tracker"

// Metrics holds the L1 health-tracker metrics. The tracker owns and reports
// these itself from its poll loop, so no other component needs to read its
// state. We only care about lag at second granularity (the halt/warn
// thresholds are minute-scale), so the value is whole seconds.
type Metrics struct {
	// LagSeconds is how far behind wall-clock the node's view of L1 is: now
	// minus the latest observed L1 head time (or since the first failed poll
	// during an RPC outage). Crossing the halt threshold stops block
	// production and sync.
	LagSeconds metrics.Gauge
}

// PrometheusMetrics registers the L1 tracker metrics on the default Prometheus
// registry under the given namespace (use "morphnode").
func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	var labels []string
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		LagSeconds: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "lag_seconds",
			Help:      "Seconds the node's view of L1 has fallen behind wall-clock.",
		}, labels).With(labelsAndValues...),
	}
}

// NopMetrics returns metrics that discard all observations (no registration).
func NopMetrics() *Metrics {
	return &Metrics{
		LagSeconds: discard.NewGauge(),
	}
}

// SetLag records the L1 lag, truncated to whole seconds.
func (m *Metrics) SetLag(d time.Duration) { m.LagSeconds.Set(float64(d / time.Second)) }
