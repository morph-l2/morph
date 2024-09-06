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

func NewEventIndexer(storePath string, client *ethclient.Client, deployedBlock *big.Int, filter ethereum.FilterQuery) *EventIndexer {
	return &EventIndexer{
		storePath:   storePath,
		client:      client,
		deployBlock: deployedBlock,
		filterQuery: filter,
	}
}

func (l *EventIndexer) Index() {

	storage := NewEventInfoStorage(l.storePath)
	err := storage.Load()
	if err != nil {
		log.Crit("Failed to load storage", "error", err, "file_name", storage.Filename)
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
			log.Error("Failed to get current block number", "error", err)
			continue
		}

		// Perform indexing operation
		indexedEventInfo, err := l.index(l.client, big.NewInt(int64(storage.BlockProcessed)), big.NewInt(int64(currentBlock)))
		if err != nil {
			log.Error("Indexing operation failed", "error", err)
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
			log.Error("Failed to update storage", "error", err, "file_name", storage.Filename)
		} else {
			log.Info("Storage updated", "processed_block", indexedEventInfo.BlockProcessed, "block_time", indexedEventInfo.BlockTime)
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
	for endBlock >= lastProcessedBlock {
		logFilter.FromBlock = big.NewInt(int64(startBlock))
		logFilter.ToBlock = big.NewInt(int64(endBlock))
		historicalLogs, err := client.FilterLogs(context.Background(), logFilter)
		if err != nil {
			log.Error("Failed to get historical logs", "error", err)
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
			startBlock = endBlock - ei.indexStep
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
