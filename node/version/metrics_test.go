package version

import (
	"testing"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
)

// TestPrometheusMetrics_RegistersInfoGaugeSetToOne verifies that Register
// publishes a single gauge, always valued 1, carrying the build metadata
// passed to PrometheusMetrics as labels.
func TestPrometheusMetrics_RegistersInfoGaugeSetToOne(t *testing.T) {
	// Scoped per-test namespace: dodges "duplicate metrics collector
	// registration" panics against the shared default registry, matching the
	// pattern used in derivation/base_client_test.go.
	namespace := "morphnode_test_" + t.Name()

	info := PrometheusMetrics(namespace, "v1.2.3", "abc1234", "2026-09-22T00:00:00Z")
	info.Register()

	mfs, err := stdprometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	var found *dto.Metric
	for _, mf := range mfs {
		if mf.GetName() == namespace+"_info" {
			require.Len(t, mf.GetMetric(), 1)
			found = mf.GetMetric()[0]
			break
		}
	}
	require.NotNil(t, found, "expected %s_info gauge to be registered", namespace)
	require.Equal(t, float64(1), found.GetGauge().GetValue())

	labels := map[string]string{}
	for _, l := range found.GetLabel() {
		labels[l.GetName()] = l.GetValue()
	}
	require.Equal(t, "v1.2.3", labels["version"])
	require.Equal(t, "abc1234", labels["git_commit"])
	require.Equal(t, "2026-09-22T00:00:00Z", labels["build_time"])
	require.NotEmpty(t, labels["go_version"])
	require.NotEmpty(t, labels["os"])
	require.NotEmpty(t, labels["arch"])
}

// TestNopMetrics_DoesNotPanic verifies NopMetrics().Register() is a safe
// no-op, for callers (e.g. tests) that don't want a live Prometheus registry.
func TestNopMetrics_DoesNotPanic(t *testing.T) {
	require.NotPanics(t, func() { NopMetrics().Register() })
}
