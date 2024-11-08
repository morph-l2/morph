package l1checker

import (
	"context"
	"morph-l2/tx-submitter/iface"
	"time"

	"github.com/morph-l2/go-ethereum/log"
)

type IBlockMonitor interface {
	BlockNotIncreasedIn(time.Duration) bool
}

type BlockMonitor struct {
	blockGenerateTime    time.Duration //12s for Dencun
	latestBlockTime      time.Time
	noGrowthBlockCntTime time.Duration
	client               iface.L1Client
}

func NewBlockMonitor(notGrowthInBlocks int64, client iface.L1Client) *BlockMonitor {
	return &BlockMonitor{
		blockGenerateTime:    time.Second * 12,
		latestBlockTime:      time.Time{},
		noGrowthBlockCntTime: time.Second * time.Duration(notGrowthInBlocks) * 12,
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
		m.latestBlockTime = time.Unix(int64(header.Time), 0)
	}
}

func (m *BlockMonitor) IsGrowth() bool {
	if m.latestBlockTime.IsZero() {
		return false
	}
	return time.Since(m.latestBlockTime) > m.noGrowthBlockCntTime
}
