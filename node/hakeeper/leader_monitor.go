package hakeeper

import "sync/atomic"

// leaderMonitor watches the Raft leader channel.
// On becoming leader: run Barrier to ensure FSM is caught up, then set leaderReady=1.
// On losing leadership: immediately set leaderReady=0.
func (h *HAService) leaderMonitor() {
	defer h.wg.Done()

	for {
		select {
		case <-h.stopCh:
			return
		case isLeader, ok := <-h.r.LeaderCh():
			if !ok {
				return
			}
			if isLeader {
				h.logger.Info("hakeeper: became leader, running Barrier")
				if err := h.r.Barrier(raftInfiniteTimeout).Error(); err != nil {
					h.logger.Error("hakeeper: Barrier failed, leaderReady not set", "err", err)
					continue
				}
				atomic.StoreInt32(&h.leaderReady, 1)
				h.logger.Info("hakeeper: leader ready")
			} else {
				atomic.StoreInt32(&h.leaderReady, 0)
				h.logger.Info("hakeeper: lost leadership")
			}
		}
	}
}
