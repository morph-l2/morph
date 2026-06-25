package hakeeper

import (
	"testing"

	"github.com/go-kit/kit/metrics"
	"github.com/hashicorp/raft"
)

type recordingGauge struct {
	value float64
}

func (g *recordingGauge) With(labelValues ...string) metrics.Gauge { return g }
func (g *recordingGauge) Set(value float64)                        { g.value = value }
func (g *recordingGauge) Add(delta float64)                        { g.value += delta }

func TestRaftMetricsObservationFilter(t *testing.T) {
	if !raftMetricsObservationFilter(&raft.Observation{Data: raft.Leader}) {
		t.Fatal("RaftState observations should be included")
	}
	if !raftMetricsObservationFilter(&raft.Observation{Data: raft.PeerObservation{}}) {
		t.Fatal("PeerObservation observations should be included")
	}
	if raftMetricsObservationFilter(&raft.Observation{Data: "ignored"}) {
		t.Fatal("unrelated observations should be filtered out")
	}
}

func TestHandleRaftObservationUpdatesRaftStateImmediately(t *testing.T) {
	stateGauge := &recordingGauge{}
	h := &HAService{
		metrics: NopMetrics(),
	}
	h.metrics.RaftState = stateGauge

	h.handleRaftObservation(raft.Observation{Data: raft.Candidate})
	if stateGauge.value != float64(raft.Candidate) {
		t.Fatalf("raft_state metric = %v, want %v", stateGauge.value, raft.Candidate)
	}

	h.handleRaftObservation(raft.Observation{Data: raft.Leader})
	if stateGauge.value != float64(raft.Leader) {
		t.Fatalf("raft_state metric = %v, want %v", stateGauge.value, raft.Leader)
	}
}
