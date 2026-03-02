package services

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/log"
)

// ReorgDetector tracks chain reorganizations
type ReorgDetector struct {
	mu sync.RWMutex

	// Track the last N block hashes and numbers
	blockHistory []blockInfo
	maxHistory   int

	l1Client iface.Client
	metrics  iface.IMetrics
}

type blockInfo struct {
	number uint64
	hash   common.Hash
}

func NewReorgDetector(l1Client iface.Client, metrics iface.IMetrics) *ReorgDetector {
	return &ReorgDetector{
		blockHistory: make([]blockInfo, 0),
		maxHistory:   5, // Track last 50 blocks
		l1Client:     l1Client,
		metrics:      metrics,
	}
}

// DetectReorg checks if a reorg has occurred by comparing current chain with tracked history
func (r *ReorgDetector) DetectReorg(ctx context.Context) (bool, uint64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.blockHistory) == 0 {
		// First run, initialize history
		return false, 0, r.updateHistory(ctx)
	}

	// Get latest block
	latestBlock, err := r.l1Client.BlockByNumber(ctx, nil)
	if err != nil {
		return false, 0, fmt.Errorf("failed to get latest block: %w", err)
	}
	if latestBlock == nil {
		return false, 0, fmt.Errorf("latest block is nil")
	}

	// Check each block in history to find reorg point
	reorgDepth := uint64(0)
	for i, info := range r.blockHistory {
		block, err := r.l1Client.BlockByNumber(ctx, new(big.Int).SetUint64(info.number))
		if err != nil {
			return false, 0, fmt.Errorf("failed to get block %d: %w", info.number, err)
		}
		if block == nil {
			return false, 0, fmt.Errorf("block %d is nil", info.number)
		}

		if block.Hash() != info.hash {
			// Reorg detected
			reorgDepth = latestBlock.NumberU64() - info.number
			log.Warn("Chain reorganization detected",
				"depth", reorgDepth,
				"old_hash", info.hash,
				"new_hash", block.Hash(),
				"block_number", info.number)

			// Update metrics
			r.metrics.IncReorgs()
			r.metrics.SetReorgDepth(float64(reorgDepth))

			// Truncate history before reorg point and rebuild
			r.blockHistory = r.blockHistory[:i]
			err = r.updateHistory(ctx)
			return true, reorgDepth, err
		}
	}

	// No reorg, just update history
	return false, 0, r.updateHistory(ctx)
}

// updateHistory updates the tracked block history
func (r *ReorgDetector) updateHistory(ctx context.Context) error {
	latest, err := r.l1Client.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get latest block: %w", err)
	}
	if latest == nil {
		return fmt.Errorf("latest block is nil")
	}

	// Add new blocks to history
	currentNum := latest.NumberU64()
	startNum := currentNum
	if len(r.blockHistory) > 0 {
		startNum = r.blockHistory[len(r.blockHistory)-1].number + 1
	}

	for num := startNum; num <= currentNum; num++ {
		block, err := r.l1Client.BlockByNumber(ctx, new(big.Int).SetUint64(num))
		if err != nil {
			return fmt.Errorf("failed to get block %d: %w", num, err)
		}
		if block == nil {
			return fmt.Errorf("block %d is nil", num)
		}

		r.blockHistory = append(r.blockHistory, blockInfo{
			number: num,
			hash:   block.Hash(),
		})
	}

	// Trim history if too long
	if len(r.blockHistory) > r.maxHistory {
		r.blockHistory = r.blockHistory[len(r.blockHistory)-r.maxHistory:]
	}

	return nil
}
