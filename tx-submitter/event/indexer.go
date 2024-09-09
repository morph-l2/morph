package event

import (
	"context"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
)

type EventIndexer struct {
	client      *ethclient.Client
	storePath   string   // store path
	deployBlock *big.Int // Block number of contract deployment
	filterQuery ethereum.FilterQuery
	indexStep   uint64 // index step
}

func NewEventIndexer(storePath string, client *ethclient.Client, deployedBlock *big.Int, filter ethereum.FilterQuery, indexStep uint64) *EventIndexer {
	return &EventIndexer{
		storePath:   storePath,
		client:      client,
		deployBlock: deployedBlock,
		filterQuery: filter,
		indexStep:   indexStep,
	}
}

func (l *EventIndexer) Index() {
	log.Info("event indexer started")
	storage := NewEventInfoStorage(l.storePath)
	err := storage.Load()
	if err != nil {
		log.Crit("failed to load storage", "error", err, "file_name", storage.Filename)
	}
	if storage.BlockProcessed == 0 {
		storage.BlockProcessed = l.deployBlock.Uint64()
	}

	// Create a ticker that triggers every minute
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	for range ticker.C {

		// Get the current block number
		currentBlock, err := l.client.BlockNumber(context.Background())
		if err != nil {
			log.Error("failed to get current block number", "error", err)
			continue
		}

		if currentBlock <= storage.BlockProcessed {
			log.Info("no new block to index", "current_block", currentBlock, "last_processed_block", storage.BlockProcessed)
			continue
		}

		// Perform indexing operation
		indexedEventInfo, err := l.index(l.client, big.NewInt(int64(storage.BlockProcessed)), big.NewInt(int64(currentBlock)))
		if err != nil {
			log.Error("indexing operation failed", "error", err)
			continue
		}

		if indexedEventInfo != nil {
			storage.EventInfo = *indexedEventInfo
		} else {
			storage.EventInfo = EventInfo{
				BlockProcessed: currentBlock,
			}
		}
		// Update storage
		err = storage.Store()
		if err != nil {
			log.Error("failed to update storage", "error", err, "file_name", storage.Filename)
		} else {
			log.Info("storage updated", "processed_block", storage.EventInfo.BlockProcessed, "block_time", storage.EventInfo.BlockTime)
		}

	}

}

// filter logs from from_block to to_block
// if the range is too large, will divide the range
func (ei *EventIndexer) index(client *ethclient.Client, fromBlock, toBlock *big.Int) (*EventInfo, error) {

	endBlock := toBlock.Uint64()
	startBlock := endBlock - ei.indexStep
	lastProcessedBlock := fromBlock.Uint64()
	logFilter := ei.GetFilter()

	// Find the last unprocessed log
	for endBlock >= lastProcessedBlock && endBlock > startBlock {
		logFilter.FromBlock = big.NewInt(int64(startBlock))
		logFilter.ToBlock = big.NewInt(int64(endBlock))
		historicalLogs, err := client.FilterLogs(context.Background(), logFilter)
		if err != nil {
			log.Error("Failed to get historical logs",
				"error", err,
				"start_block", startBlock,
				"end_block", endBlock,
			)
			continue
		} else {
			if len(historicalLogs) > 0 {
				latestLog := historicalLogs[len(historicalLogs)-1]
				var block *types.Block
				for {
					block, err = client.BlockByNumber(context.Background(), new(big.Int).SetUint64(latestLog.BlockNumber))
					if err != nil {
						log.Error("Failed to get block", "error", err)
						time.Sleep(time.Second * 5)
						continue
					} else {
						return &EventInfo{
							BlockProcessed: toBlock.Uint64(),
							BlockTime:      block.Time(),
						}, nil
					}
				}
			}
			// update query range
			endBlock = startBlock
			if endBlock <= lastProcessedBlock+ei.indexStep {
				startBlock = lastProcessedBlock + 1
			} else {
				startBlock = endBlock - ei.indexStep
			}

		}
	}
	return nil, nil
}

func (l *EventIndexer) GetFilter() ethereum.FilterQuery {
	return l.filterQuery
}

func (l *EventIndexer) GetStorePath() string {
	return l.storePath
}
