package hakeeper

import (
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/hashicorp/raft"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// metricsSubsystem groups all HA-keeper metrics under morphnode_hakeeper_*.
const metricsSubsystem = "hakeeper"

// commitBuckets / fsmApplyBuckets are millisecond exponential buckets
// (1ms .. ~8.2s) suitable for the sub-second-to-seconds commit / apply paths.
var (
	commitBuckets   = stdprometheus.ExponentialBuckets(1, 2, 14)
	fsmApplyBuckets = stdprometheus.ExponentialBuckets(1, 2, 14)
)

// Metrics holds the Raft HA-cluster metrics. Every series carries a server_id
// label (the local Raft node id) so a cluster dashboard can aggregate across
// members. State gauges are sampled periodically by HAService; the duration
// and counter series are updated inline on the commit / FSM-apply paths.
//
// All values are integer-valued: durations are recorded in whole milliseconds
// (seconds would round sub-second work to 0). Use the typed helper methods so
// call sites never deal with float64 directly.
type Metrics struct {
	// RaftState is the local Raft role as an enum:
	// Follower=0, Candidate=1, Leader=2, Shutdown=3. Across the cluster exactly
	// one member should report Leader(2).
	RaftState metrics.Gauge

	// ClusterMembers is the number of servers in the current Raft configuration.
	ClusterMembers metrics.Gauge

	// CommitDurationMilliseconds measures the leader-side block commit latency
	// in milliseconds, labeled by step (encode / raft).
	CommitDurationMilliseconds metrics.Histogram

	// FSMApplyDurationMilliseconds measures the FSM business-callback (geth apply
	// + signature persist) duration in milliseconds per committed log entry.
	FSMApplyDurationMilliseconds metrics.Histogram

	// FSMBlockChannelDropsTotal counts blocks dropped because the FSM->broadcast
	// channel was full (subscriber too slow).
	FSMBlockChannelDropsTotal metrics.Counter
}

// PrometheusMetrics registers the HA metrics on the default Prometheus
// registry under the given namespace. Pass "server_id", <id> in
// labelsAndValues so every series is tagged with the local node id.
func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	var labels []string
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		RaftState: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "raft_state",
			Help:      "Local Raft role: Follower=0, Candidate=1, Leader=2, Shutdown=3.",
		}, labels).With(labelsAndValues...),
		ClusterMembers: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "cluster_members",
			Help:      "Number of servers in the current Raft configuration.",
		}, labels).With(labelsAndValues...),
		CommitDurationMilliseconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "commit_duration_milliseconds",
			Help:      "Leader-side block commit latency in milliseconds, labeled by step (encode / raft).",
			Buckets:   commitBuckets,
		}, append(labels, "step")).With(labelsAndValues...),
		FSMApplyDurationMilliseconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "fsm_apply_duration_milliseconds",
			Help:      "FSM business-callback (geth apply + signature persist) duration in milliseconds per committed log entry.",
			Buckets:   fsmApplyBuckets,
		}, labels).With(labelsAndValues...),
		FSMBlockChannelDropsTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: metricsSubsystem,
			Name:      "fsm_block_channel_drops_total",
			Help:      "Blocks dropped because the FSM->broadcast channel was full.",
		}, labels).With(labelsAndValues...),
	}
}

// NopMetrics returns metrics that discard all observations (no registration).
func NopMetrics() *Metrics {
	return &Metrics{
		RaftState:                    discard.NewGauge(),
		ClusterMembers:               discard.NewGauge(),
		CommitDurationMilliseconds:   discard.NewHistogram(),
		FSMApplyDurationMilliseconds: discard.NewHistogram(),
		FSMBlockChannelDropsTotal:    discard.NewCounter(),
	}
}

// ---- Typed helpers (keep float64 conversions out of call sites) ----

// SetRaftState records the local Raft role enum.
func (m *Metrics) SetRaftState(s raft.RaftState) { m.RaftState.Set(float64(s)) }

// SetClusterMembers records the current Raft configuration size.
func (m *Metrics) SetClusterMembers(n int) { m.ClusterMembers.Set(float64(n)) }

// ObserveCommitDuration records a leader-side commit step latency in ms.
func (m *Metrics) ObserveCommitDuration(step string, d time.Duration) {
	m.CommitDurationMilliseconds.With("step", step).Observe(float64(d.Milliseconds()))
}

// ObserveFSMApplyDuration records the FSM business-callback latency in ms.
func (m *Metrics) ObserveFSMApplyDuration(d time.Duration) {
	m.FSMApplyDurationMilliseconds.Observe(float64(d.Milliseconds()))
}

// IncFSMBlockChannelDrops counts one dropped block (FSM->broadcast full).
func (m *Metrics) IncFSMBlockChannelDrops() { m.FSMBlockChannelDropsTotal.Add(1) }
