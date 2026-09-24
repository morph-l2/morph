// Package version exposes morphnode's build/version metadata as a Prometheus
// info-style gauge, mirroring morph-reth's reth_info{} gauge: a static gauge
// permanently set to 1, whose Prometheus labels (not its value) carry the
// actual information. Scraping tools use this to build a "which build/version
// is running where" view (e.g. `max by (version) (morphnode_info{})`).
package version

import (
	"runtime"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Info is a static Prometheus gauge carrying morphnode's build metadata as
// labels. Its value is always 1; only the labels are meaningful.
type Info struct {
	gauge metrics.Gauge
}

// PrometheusMetrics registers the morphnode_info gauge on the default
// Prometheus registry under the given namespace (no subsystem: the resulting
// metric name is "<namespace>_info", e.g. "morphnode_info"). version,
// gitCommit and buildTime should be the values injected via -ldflags at build
// time (see cmd/node/version.go); goVersion/os/arch are read from the runtime.
func PrometheusMetrics(namespace, version, gitCommit, buildTime string) *Info {
	labels := []string{"version", "git_commit", "build_time", "go_version", "os", "arch"}
	g := prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "info",
		Help:      "Static info metric (always 1) carrying morphnode build/version metadata as labels.",
	}, labels).With(
		"version", version,
		"git_commit", gitCommit,
		"build_time", buildTime,
		"go_version", runtime.Version(),
		"os", runtime.GOOS,
		"arch", runtime.GOARCH,
	)
	return &Info{gauge: g}
}

// NopMetrics returns an Info that discards its observation (no registration).
// Mirrors the NopMetrics helpers used by the other metrics packages (e.g.
// hakeeper.NopMetrics) for tests that don't need a live Prometheus registry.
func NopMetrics() *Info {
	return &Info{gauge: discard.NewGauge()}
}

// Register sets the gauge to 1, publishing the labels attached at
// construction time. Call once at startup.
func (i *Info) Register() { i.gauge.Set(1) }
