package hakeeper

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb/v2"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/types"

	hakeeperrpc "morph-l2/node/hakeeper/rpc"
)

const (
	raftTimeout         = 5 * time.Second // default timeout for membership ops and TCP connections
	raftInfiniteTimeout = 0               // wait forever
	raftMaxConnPool     = 10
	raftSnapshots       = 1 // snapshot data is trivial (8 bytes); keep 1 for log compaction

	// advertisedAddr may be a hostname whose DNS record is not yet propagated the
	// instant the pod starts (it resolves a few seconds later). Retry the initial
	// resolve instead of failing fast on that startup race.
	advAddrResolveRetries  = 5               // retries after the first attempt (6 resolves total)
	advAddrResolveInterval = 3 * time.Second // 5 retries × 3s ≈ 15s max wait for DNS to propagate
)

// HAService implements the SequencerHA interface from tendermint/sequencer.
// It also satisfies rpc.ConsensusAdapter so it can be passed directly to the RPC server.
type HAService struct {
	logger         tmlog.Logger
	cfg            *Config
	advertisedAddr string // resolved once in New(), used throughout
	fsm            *BlockFSM
	rpcServer      *hakeeperrpc.Server
	metrics        *Metrics

	// Raft internals (initialized in Start)
	r            *raft.Raft
	transport    *raft.NetworkTransport
	raftObserver *raft.Observer

	leaderReady int32 // atomic: 1 = can produce blocks
	stopCh      chan struct{}
	wg          sync.WaitGroup
}

// Ensure HAService satisfies rpc.ConsensusAdapter at compile time.
var _ hakeeperrpc.ConsensusAdapter = (*HAService)(nil)

// New creates a new HAService.
// Expects cfg to be fully resolved (Resolve + Validate already called).
// Call SetOnBlockApplied before Start().
func New(cfg *Config, logger tmlog.Logger) (*HAService, error) {
	// Register HA metrics on the default registry, tagged with this node's id
	// (matches the morphnode_* namespace used by the executor/syncer metrics).
	m := PrometheusMetrics("morphnode", "server_id", cfg.ServerID)
	return &HAService{
		logger:         logger,
		cfg:            cfg,
		advertisedAddr: cfg.Consensus.AdvertisedAddr, // already resolved
		fsm:            NewBlockFSM(logger, m),
		metrics:        m,
		stopCh:         make(chan struct{}),
	}, nil
}

// SetOnBlockApplied registers the business callback invoked by the FSM on every
// committed log entry. Must be called before Start().
func (h *HAService) SetOnBlockApplied(fn func(*types.BlockV2) error) {
	h.fsm.SetOnBlockApplied(fn)
}

// ── SequencerHA interface ────────────────────────────────────────────────────

// Start initializes Raft and the management RPC server.
// Called by StateV2.OnStart() at upgrade height.
func (h *HAService) Start() error {
	if err := h.initRaft(); err != nil {
		return fmt.Errorf("HAService.Start: %w", err)
	}

	rpcSrv, err := hakeeperrpc.New(h.logger, h.cfg.RPC.ListenAddr, h.cfg.RPC.ListenPort, h, h.cfg.RPC.Token)
	if err != nil {
		h.shutdownRaft()
		return fmt.Errorf("HAService.Start: rpc: %w", err)
	}
	if err := rpcSrv.Start(); err != nil {
		h.shutdownRaft()
		return fmt.Errorf("HAService.Start: rpc start: %w", err)
	}
	h.rpcServer = rpcSrv

	// Subscribe before taking the initial snapshot: any state change that
	// happens between the two is then captured by the observer rather than
	// lost in the gap (which would leave raft_state stale until the next
	// event or the minute-interval safety net).
	h.startRaftMetricsObserver()
	h.refreshRaftMetrics()

	h.wg.Add(1)
	go h.leaderMonitor()

	if !h.cfg.Bootstrap {
		h.wg.Add(1)
		go h.joinLoop()
	}

	h.logger.Info("hakeeper: started", "server_id", h.cfg.ServerID, "bootstrap", h.cfg.Bootstrap)
	return nil
}

// Stop gracefully shuts down the HAService.
// Order: close stopCh → shutdown Raft (unblocks Barrier) → wg.Wait → stop RPC.
func (h *HAService) Stop() {
	close(h.stopCh)
	if h.raftObserver != nil && h.r != nil {
		h.r.DeregisterObserver(h.raftObserver)
	}
	h.shutdownRaft()
	h.wg.Wait()
	if h.rpcServer != nil {
		h.rpcServer.Stop()
	}
	h.logger.Info("hakeeper: stopped")
}

// IsLeader returns true only when this node is the Raft leader AND the
// post-election Barrier has completed (leaderReady == 1).
func (h *HAService) IsLeader() bool {
	if h.r == nil {
		return false
	}
	return h.r.State() == raft.Leader && atomic.LoadInt32(&h.leaderReady) == 1
}

// Join tries each address in JoinAddrs until one succeeds in adding this node to the cluster.
func (h *HAService) Join() error {
	var lastErr error
	for _, addr := range h.cfg.JoinAddrs {
		if err := h.tryJoin(addr); err != nil {
			lastErr = err
			h.logger.Error("hakeeper: join attempt failed", "addr", addr, "err", err)
			continue
		}
		return nil
	}
	return fmt.Errorf("Join: all addresses failed, last error: %w", lastErr)
}

func (h *HAService) tryJoin(addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), raftTimeout)
	defer cancel()

	client, err := hakeeperrpc.DialAPIClient(ctx, addr, h.cfg.RPC.Token)
	if err != nil {
		return fmt.Errorf("dial %s: %w", addr, err)
	}
	defer client.Close()

	membership, err := client.ClusterMembership(ctx)
	if err != nil {
		return fmt.Errorf("get membership from %s: %w", addr, err)
	}

	// If this node is already a member (e.g. after a restart), skip AddServerAsVoter.
	for _, srv := range membership.Servers {
		if srv.ID == h.cfg.ServerID {
			h.logger.Info("hakeeper: already a cluster member, skipping join", "id", h.cfg.ServerID)
			return nil
		}
	}

	if err := client.AddServerAsVoter(ctx, h.cfg.ServerID, h.advertisedAddr, membership.Version); err != nil {
		return err
	}
	h.refreshClusterMembers()
	return nil
}

// Commit replicates a signed block via Raft.
// Three-level response: quorum error → return; leader FSM error → panic; ok → nil.
func (h *HAService) Commit(block *types.BlockV2) error {
	t0 := time.Now()

	data, err := encodeBlock(block)
	if err != nil {
		return fmt.Errorf("Commit: encode: %w", err)
	}
	encodeDur := time.Since(t0)
	h.metrics.ObserveCommitDuration("encode", encodeDur)

	t1 := time.Now()
	f := h.r.Apply(data, raftInfiniteTimeout)
	if err := f.Error(); err != nil {
		return err
	}
	raftDur := time.Since(t1)
	h.metrics.ObserveCommitDuration("raft", raftDur)

	if resp := f.Response(); resp != nil {
		if err, ok := resp.(error); ok {
			panic(fmt.Sprintf("hakeeper: leader FSM.Apply failed: %v", err))
		}
	}

	totalDur := time.Since(t0)
	h.logger.Debug("[PERF] HAService.Commit",
		"height", block.Number,
		"encode_ms", float64(encodeDur.Microseconds())/1000.0,
		"raft_ms", float64(raftDur.Microseconds())/1000.0,
		"total_ms", float64(totalDur.Microseconds())/1000.0,
		"dataBytes", len(data),
		"txCount", len(block.Transactions),
	)

	return nil
}

// Subscribe returns the channel delivering blocks after FSM.Apply.
func (h *HAService) Subscribe() <-chan *types.BlockV2 {
	return h.fsm.blockCh
}

// ── rpc.ConsensusAdapter interface ──────────────────────────────────────────

func (h *HAService) Leader() bool {
	return h.r != nil && h.r.State() == raft.Leader
}

func (h *HAService) LeaderWithID() *hakeeperrpc.ServerInfo {
	if h.r == nil {
		return nil
	}
	addr, id := h.r.LeaderWithID()
	if id == "" {
		return nil
	}
	return &hakeeperrpc.ServerInfo{ID: string(id), Addr: string(addr), Suffrage: hakeeperrpc.Voter}
}

func (h *HAService) AddVoter(id, addr string, version uint64) error {
	if err := h.r.AddVoter(raft.ServerID(id), raft.ServerAddress(addr), version, raftTimeout).Error(); err != nil {
		return err
	}
	h.refreshClusterMembers()
	return nil
}

func (h *HAService) AddNonVoter(id, addr string, version uint64) error {
	if err := h.r.AddNonvoter(raft.ServerID(id), raft.ServerAddress(addr), version, raftTimeout).Error(); err != nil {
		return err
	}
	h.refreshClusterMembers()
	return nil
}

func (h *HAService) DemoteVoter(id string, version uint64) error {
	if err := h.r.DemoteVoter(raft.ServerID(id), version, raftTimeout).Error(); err != nil {
		return err
	}
	h.refreshClusterMembers()
	return nil
}

func (h *HAService) RemoveServer(id string, version uint64) error {
	if err := h.r.RemoveServer(raft.ServerID(id), version, raftTimeout).Error(); err != nil {
		return err
	}
	h.refreshClusterMembers()
	return nil
}

func (h *HAService) TransferLeader() error {
	if err := h.r.LeadershipTransfer().Error(); err != nil && err != raft.ErrNotLeader {
		return err
	}
	return nil
}

func (h *HAService) TransferLeaderTo(id, addr string) error {
	return h.r.LeadershipTransferToServer(raft.ServerID(id), raft.ServerAddress(addr)).Error()
}

func (h *HAService) ClusterMembership() (*hakeeperrpc.ClusterMembership, error) {
	future := h.r.GetConfiguration()
	if err := future.Error(); err != nil {
		return nil, err
	}
	var servers []hakeeperrpc.ServerInfo
	for _, srv := range future.Configuration().Servers {
		servers = append(servers, hakeeperrpc.ServerInfo{
			ID:       string(srv.ID),
			Addr:     string(srv.Address),
			Suffrage: hakeeperrpc.ServerSuffrage(srv.Suffrage),
		})
	}
	_, leaderID := h.r.LeaderWithID()
	return &hakeeperrpc.ClusterMembership{
		Servers:  servers,
		LeaderID: string(leaderID),
		Version:  future.Index(),
	}, nil
}

func (h *HAService) ServerID() string { return h.cfg.ServerID }

func (h *HAService) Addr() string { return h.advertisedAddr }

// ── internal ─────────────────────────────────────────────────────────────────

// initRaft creates the Raft instance. Called once from Start().
// On failure, all opened resources are cleaned up via a single deferred closure.
func (h *HAService) initRaft() (retErr error) {
	if err := os.MkdirAll(h.cfg.StorageDir, 0o750); err != nil {
		return fmt.Errorf("mkdir %q: %w", h.cfg.StorageDir, err)
	}

	var (
		logStore    *boltdb.BoltStore
		stableStore *boltdb.BoltStore
		transport   *raft.NetworkTransport
		r           *raft.Raft
	)
	defer func() {
		if retErr != nil {
			if r != nil {
				r.Shutdown()
			}
			if transport != nil {
				_ = transport.Close()
			}
			if stableStore != nil {
				_ = stableStore.Close()
			}
			if logStore != nil {
				_ = logStore.Close()
			}
		}
	}()

	var err error
	logStore, err = boltdb.NewBoltStore(filepath.Join(h.cfg.StorageDir, "raft-log.db"))
	if err != nil {
		return fmt.Errorf("log store: %w", err)
	}
	stableStore, err = boltdb.NewBoltStore(filepath.Join(h.cfg.StorageDir, "raft-stable.db"))
	if err != nil {
		return fmt.Errorf("stable store: %w", err)
	}

	raftLogLevel := hclog.Info
	if h.cfg.Debug {
		raftLogLevel = hclog.Debug
	}
	raftLogger := hclog.New(&hclog.LoggerOptions{
		Name:   "raft",
		Level:  raftLogLevel,
		Output: os.Stderr,
	})

	snapshotStore, err := raft.NewFileSnapshotStoreWithLogger(h.cfg.StorageDir, raftSnapshots, raftLogger)
	if err != nil {
		return fmt.Errorf("snapshot store: %w", err)
	}

	rc := raft.DefaultConfig()
	rc.LocalID = raft.ServerID(h.cfg.ServerID)
	rc.SnapshotInterval = h.cfg.Snapshot.Interval
	rc.SnapshotThreshold = h.cfg.Snapshot.Threshold
	rc.TrailingLogs = h.cfg.Snapshot.TrailingLogs
	rc.HeartbeatTimeout = h.cfg.Timeout.Heartbeat
	rc.LeaderLeaseTimeout = h.cfg.Timeout.LeaderLease
	rc.Logger = raftLogger

	// Resolve advertised addr to *net.TCPAddr for the transport layer (required by hashicorp/raft).
	// Note: the resolved IP is only used by the transport's LocalAddr(). The ServerAddress
	// stored in Raft cluster config (BootstrapCluster/AddServerAsVoter) uses the raw
	// h.advertisedAddr which may be a hostname — Raft's Dial() re-resolves DNS each time.
	var tcpAdvAddr *net.TCPAddr
	for attempt := 0; ; attempt++ {
		tcpAdvAddr, err = net.ResolveTCPAddr("tcp", h.advertisedAddr)
		if err == nil {
			break
		}
		if attempt >= advAddrResolveRetries {
			// DNS still hasn't propagated after the full retry budget. Some consensus-switch
			// paths only log StartSequencerRoutines errors and keep running, which would leave
			// this node alive but with Raft never started. Panic instead so the process exits
			// non-zero and the orchestrator (k8s) restarts the pod for another DNS attempt.
			panic(fmt.Errorf("hakeeper: resolve advertised addr %q after %d retries: %w", h.advertisedAddr, advAddrResolveRetries, err))
		}
		h.logger.Info("hakeeper: advertised addr not resolvable yet, retrying",
			"addr", h.advertisedAddr, "retry", attempt+1, "max", advAddrResolveRetries, "err", err)
		time.Sleep(advAddrResolveInterval)
	}

	bindAddr := fmt.Sprintf("%s:%d", h.cfg.Consensus.ListenAddr, h.cfg.Consensus.ListenPort)
	transport, err = raft.NewTCPTransportWithLogger(bindAddr, tcpAdvAddr, raftMaxConnPool, raftTimeout, raftLogger)
	if err != nil {
		return fmt.Errorf("TCP transport: %w", err)
	}

	r, err = raft.NewRaft(rc, h.fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return fmt.Errorf("raft.NewRaft: %w", err)
	}

	if h.cfg.Bootstrap {
		f := r.BootstrapCluster(raft.Configuration{Servers: []raft.Server{
			{ID: raft.ServerID(h.cfg.ServerID), Address: raft.ServerAddress(h.advertisedAddr), Suffrage: raft.Voter},
		}})
		if err := f.Error(); err != nil && !errors.Is(err, raft.ErrCantBootstrap) {
			return fmt.Errorf("bootstrap: %w", err)
		}
	}

	h.r = r
	h.transport = transport

	h.logger.Info("hakeeper: raft initialized", "bind", bindAddr)
	return nil
}

func (h *HAService) shutdownRaft() {
	if h.r != nil {
		if err := h.r.Shutdown().Error(); err != nil {
			h.logger.Error("hakeeper: raft shutdown error", "err", err)
		}
	}
}

// joinLoop retries Join() with exponential backoff (2s → 30s) until success or stop.
func (h *HAService) joinLoop() {
	defer h.wg.Done()
	backoff := 2 * time.Second
	for {
		select {
		case <-h.stopCh:
			return
		case <-time.After(backoff):
			if err := h.Join(); err != nil {
				h.logger.Error("hakeeper: join failed, retrying", "backoff", backoff, "err", err)
				if backoff < 30*time.Second {
					backoff *= 2
				}
				continue
			}
			h.logger.Info("hakeeper: joined cluster")
			h.refreshClusterMembers()
			return
		}
	}
}
