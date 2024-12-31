package l1checker

import (
	"context"
	"sync"
	"time"

	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/log"
)

const blockTime = time.Second * 12

type IBlockMonitor interface {
	BlockNotIncreasedIn(time.Duration) bool
}

type BlockMonitor struct {
	blockGenerateTime    time.Duration // 12s for Dencun
	latestBlockTime      time.Time
	noGrowthBlockCntTime time.Duration
	client               iface.L1Client
	mu                   sync.Mutex
}

func NewBlockMonitor(notGrowthInBlocks int64, client iface.L1Client) *BlockMonitor {
	return &BlockMonitor{
		blockGenerateTime:    blockTime,
		latestBlockTime:      time.Time{},
		noGrowthBlockCntTime: time.Duration(notGrowthInBlocks) * blockTime,
		client:               client,
	}
}

func (m *BlockMonitor) StartMonitoring() {
	ticker := time.NewTicker(m.blockGenerateTime)
	for ; ; <-ticker.C {
		header, err := m.client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Warn("failed to get block in blockmonitor", "error", err)
			continue
		}
		m.SetLatestBlockTime(time.Unix(int64(header.Time), 0))
	}
}

func (m *BlockMonitor) IsGrowth() bool {
	t := m.GetLatestBlockTime()
	if t.IsZero() {
		log.Warn("latest block time is zero")
		return false
	}
	return time.Since(t) < m.noGrowthBlockCntTime
}

func (m *BlockMonitor) SetLatestBlockTime(t time.Time) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.latestBlockTime = t
}

func (m *BlockMonitor) GetLatestBlockTime() time.Time {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.latestBlockTime
}
