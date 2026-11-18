package hakeeper

import (
	"net"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/go-kit/kit/metrics"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb/v2"
	"github.com/morph-l2/go-ethereum/common"
	ethrpc "github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	hakeeperrpc "morph-l2/node/hakeeper/rpc"
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

// A leader without the designated key looks for a voter that holds it, by
// asking each member's HAKeeper RPC for its signer addresses.
func TestFindSignerHolder(t *testing.T) {
	oldKey, newKey := common.HexToAddress("0xa1"), common.HexToAddress("0xb2")

	// The peer: holds both keys and serves ha_nodeInfo.
	peer := &HAService{cfg: &Config{ServerID: "node-2"}}
	peer.SetSignerAddrs(func() []common.Address { return []common.Address{oldKey, newKey} })
	srv := ethrpc.NewServer()
	if err := srv.RegisterName(hakeeperrpc.RPCNamespace, hakeeperrpc.NewAPIBackend(tmlog.NewNopLogger(), peer)); err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(srv)
	defer ts.Close()
	_, port, _ := net.SplitHostPort(ts.Listener.Addr().String())
	rpcPort, _ := strconv.Atoi(port)

	// The leader uses the same RPC port as the peer, so by default every
	// member's RPC is expected on that port of its Raft host.
	h := &HAService{
		logger: tmlog.NewNopLogger(),
		cfg:    &Config{ServerID: "node-1", RPC: RPCConfig{ListenPort: rpcPort}},
	}
	self := raft.Server{ID: "node-1", Address: "127.0.0.1:9400", Suffrage: raft.Voter}
	voter := raft.Server{ID: "node-2", Address: "127.0.0.1:9400", Suffrage: raft.Voter}
	nonvoter := raft.Server{ID: "node-3", Address: "127.0.0.1:9400", Suffrage: raft.Nonvoter}
	unreachable := raft.Server{ID: "node-4", Address: "no-port", Suffrage: raft.Voter}
	// Its default RPC address is answered by node-2: not node-5's keys.
	impostor := raft.Server{ID: "node-5", Address: "127.0.0.1:9400", Suffrage: raft.Voter}

	// Itself, non-voters, unreachable members and answers from another
	// node are passed over.
	got, ok := h.findSignerHolder([]raft.Server{self, nonvoter, unreachable, impostor, voter}, newKey)
	if !ok || got.ID != voter.ID {
		t.Fatalf("holder = %v,%v; want %s", got.ID, ok, voter.ID)
	}
	if _, ok := h.findSignerHolder([]raft.Server{self, nonvoter, unreachable, impostor}, newKey); ok {
		t.Fatal("no eligible member holds the key: want none")
	}
	if _, ok := h.findSignerHolder([]raft.Server{self, voter}, common.HexToAddress("0xc3")); ok {
		t.Fatal("nobody holds this key: want none")
	}

	// A leader on another RPC port (several nodes on one host) cannot guess
	// the peer's address and needs it spelled out in rpc.peers.
	h.cfg.RPC.ListenPort = rpcPort + 1
	if _, ok := h.findSignerHolder([]raft.Server{self, voter}, newKey); ok {
		t.Fatal("peer is not on the leader's RPC port: want none without rpc.peers")
	}
	h.cfg.RPC.Peers = map[string]string{"node-2": ts.Listener.Addr().String()}
	if got, ok := h.findSignerHolder([]raft.Server{self, voter}, newKey); !ok || got.ID != voter.ID {
		t.Fatalf("with rpc.peers: holder = %v,%v; want %s", got.ID, ok, voter.ID)
	}
}

// Only a node with empty Raft storage asks join_addrs to add it to a cluster.
// One that already has state belongs to the cluster recorded there, and must
// not be added to whichever cluster a join address is part of by now.
func TestJoinOnlyWithEmptyRaftStorage(t *testing.T) {
	start := func(t *testing.T, dir string, bootstrap bool) *HAService {
		cfg := DefaultConfig()
		cfg.ServerID = "node-1"
		cfg.StorageDir = dir
		cfg.Bootstrap = bootstrap
		cfg.JoinAddrs = []string{"127.0.0.1:9401"}
		cfg.Consensus.ListenAddr = "127.0.0.1"
		cfg.Consensus.ListenPort = 0 // any free port
		h := &HAService{
			logger:         tmlog.NewNopLogger(),
			cfg:            cfg,
			advertisedAddr: "127.0.0.1:9400",
			fsm:            NewBlockFSM(tmlog.NewNopLogger(), NopMetrics()),
			metrics:        NopMetrics(),
		}
		if err := h.initRaft(); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(h.shutdownRaft)
		return h
	}

	// Raft storage as a node that has taken part in a cluster leaves it.
	usedDir := func(t *testing.T) string {
		dir := t.TempDir()
		logs, err := boltdb.NewBoltStore(filepath.Join(dir, "raft-log.db"))
		if err != nil {
			t.Fatal(err)
		}
		if err := logs.StoreLog(&raft.Log{Index: 1, Term: 1, Type: raft.LogNoop}); err != nil {
			t.Fatal(err)
		}
		if err := logs.Close(); err != nil {
			t.Fatal(err)
		}
		return dir
	}

	if h := start(t, t.TempDir(), false); !h.shouldJoin() {
		t.Error("empty storage: a follower must join")
	}
	if h := start(t, usedDir(t), false); h.shouldJoin() {
		t.Error("existing raft state: a follower must not join again")
	}
	if h := start(t, t.TempDir(), true); h.shouldJoin() {
		t.Error("a bootstrap node never joins")
	}
}

// rpc.peers comes from the config file; ids are matched case-insensitively
// because the loader lower-cases keys, and may contain dots.
func TestConfigRPCPeers(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ha.toml")
	toml := `
[consensus]
listen_port = 9500

[rpc]
listen_port = 9401

[rpc.peers]
ha1 = "127.0.0.1:9701"
Node-B = "10.0.0.2:9501"
"node-c.example" = "10.0.0.3:9501"
`
	if err := os.WriteFile(path, []byte(toml), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg := DefaultConfig()
	if err := cfg.LoadFile(path); err != nil {
		t.Fatal(err)
	}
	if cfg.Consensus.ListenPort != 9500 || cfg.RPC.ListenPort != 9401 || cfg.RPC.ListenAddr != "0.0.0.0" {
		t.Fatalf("file values / defaults not applied: %+v %+v", cfg.Consensus, cfg.RPC)
	}
	for _, tc := range []struct{ id, raftAddr, want string }{
		{"node-c.example", "10.0.0.3:9400", "10.0.0.3:9501"}, // listed, dotted id
		{"ha1", "127.0.0.1:9700", "127.0.0.1:9701"},          // listed
		{"Node-B", "10.0.0.2:9400", "10.0.0.2:9501"},         // listed, mixed case
		{"ha2", "ha-node-2:9400", "ha-node-2:9401"},          // not listed: raft host + own RPC port
		{"ha3", "[fd00::3]:9400", "[fd00::3]:9401"},          // IPv6
	} {
		got, err := cfg.RPC.PeerAddr(tc.id, tc.raftAddr)
		if err != nil || got != tc.want {
			t.Errorf("PeerAddr(%s, %s) = %q, %v; want %q", tc.id, tc.raftAddr, got, err, tc.want)
		}
	}
	if _, err := cfg.RPC.PeerAddr("ha4", "no-port"); err == nil {
		t.Error("a Raft address without a port must be an error")
	}
}
