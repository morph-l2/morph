package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func gaugeValue(t *testing.T, name string) (float64, bool) {
	t.Helper()
	families, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		t.Fatal(err)
	}
	for _, family := range families {
		if family.GetName() == name {
			return family.GetMetric()[0].GetGauge().GetValue(), true
		}
	}
	return 0, false
}

func TestRecordSignerProbe(t *testing.T) {
	const up = "challenger_ext_signer_up"
	const last = "challenger_ext_signer_last_success_timestamp_seconds"

	RecordSignerProbe(false)
	if v, ok := gaugeValue(t, up); !ok || v != 0 {
		t.Fatalf("%s = %v (exported %v), want 0", up, v, ok)
	}
	if v, _ := gaugeValue(t, last); v != 0 {
		t.Fatalf("%s = %v after a failed probe, want 0", last, v)
	}

	RecordSignerProbe(true)
	if v, _ := gaugeValue(t, up); v != 1 {
		t.Fatalf("%s = %v, want 1", up, v)
	}
	success, _ := gaugeValue(t, last)
	if success == 0 {
		t.Fatalf("%s not set after a successful probe", last)
	}

	RecordSignerProbe(false)
	if v, _ := gaugeValue(t, up); v != 0 {
		t.Fatalf("%s = %v, want 0", up, v)
	}
	if v, _ := gaugeValue(t, last); v != success {
		t.Fatalf("%s changed on a failed probe: %v -> %v", last, success, v)
	}
}
