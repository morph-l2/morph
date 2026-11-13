package hakeeper

import (
	"testing"

	"github.com/go-kit/kit/metrics"
	"github.com/hashicorp/raft"
	"github.com/morph-l2/go-ethereum/common"
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

func TestNodeInfoReportsSignerAddrs(t *testing.T) {
	h := &HAService{cfg: &Config{ServerID: "node-1"}, advertisedAddr: "10.0.0.1:9400"}

	// No signer configured: empty (not nil) list.
	info := h.NodeInfo()
	if info.ServerID != "node-1" || info.Addr != "10.0.0.1:9400" || info.Leader {
		t.Fatalf("unexpected identity: %+v", info)
	}
	if info.SignerAddrs == nil || len(info.SignerAddrs) != 0 {
		t.Fatalf("SignerAddrs = %v, want empty", info.SignerAddrs)
	}

	// Read on every call, so a changed signer key set is reflected.
	addrs := []common.Address{common.HexToAddress("0xa1"), common.HexToAddress("0xb2")}
	h.SetSignerAddrs(func() []common.Address { return addrs })
	if got := h.NodeInfo().SignerAddrs; len(got) != 2 || got[0] != addrs[0] || got[1] != addrs[1] {
		t.Fatalf("SignerAddrs = %v, want %v", got, addrs)
	}
	addrs = addrs[1:]
	if got := h.NodeInfo().SignerAddrs; len(got) != 1 || got[0] != addrs[0] {
		t.Fatalf("SignerAddrs after change = %v, want %v", got, addrs)
	}
}
