package hakeeper

import (
	"github.com/hashicorp/raft"
)

// raftMetricsObservationFilter selects only the observations the metrics layer
// cares about: local role transitions (RaftState) and peer add/remove
// (PeerObservation).
func raftMetricsObservationFilter(o *raft.Observation) bool {
	switch o.Data.(type) {
	case raft.RaftState, raft.PeerObservation:
		return true
	default:
		return false
	}
}

// startRaftMetricsObserver subscribes to Raft observations and drains them in a
// dedicated goroutine. The channel is non-blocking (best-effort): if the buffer
// fills, an observation may be dropped, leaving its gauge stale until the next
// event of that kind (raft_state self-corrects on the next role change).
func (h *HAService) startRaftMetricsObserver() {
	ch := make(chan raft.Observation, 16)
	observer := raft.NewObserver(ch, false, raftMetricsObservationFilter)
	h.raftObserver = observer
	h.r.RegisterObserver(observer)

	h.wg.Add(1)
	go h.raftMetricsObserverLoop(ch)
}

func (h *HAService) raftMetricsObserverLoop(ch <-chan raft.Observation) {
	defer h.wg.Done()
	for {
		select {
		case <-h.stopCh:
			return
		case observation := <-ch:
			h.handleRaftObservation(observation)
		}
	}
}

func (h *HAService) handleRaftObservation(observation raft.Observation) {
	switch state := observation.Data.(type) {
	case raft.RaftState:
		h.metrics.SetRaftState(state)
	case raft.PeerObservation:
		h.refreshClusterMembers()
	}
}

// refreshRaftMetrics seeds both gauges with a one-shot snapshot. Called once at
// startup (right after the observer is registered) to set initial values.
func (h *HAService) refreshRaftMetrics() {
	if h.r == nil {
		return
	}
	h.metrics.SetRaftState(h.r.State())
	h.refreshClusterMembers()
}

func (h *HAService) refreshClusterMembers() {
	if h.r == nil {
		return
	}
	future := h.r.GetConfiguration()
	if err := future.Error(); err != nil {
		h.logger.Error("hakeeper: refresh cluster members metric failed", "err", err)
		return
	}
	h.metrics.SetClusterMembers(len(future.Configuration().Servers))
}
