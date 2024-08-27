package event

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
)

type EventListener struct {
	client          *ethclient.Client
	lastBlock       *big.Int // Last block number processed
	deployBlock     *big.Int // Block number of contract deployment
	storageFileName string
	outputChan      chan<- types.Log
	contract        common.Address
	filterQuery     ethereum.FilterQuery
}

func NewEventListener(storageFileName string, wsurl string, deployedBlock *big.Int, outputChan chan<- types.Log, contractAddress common.Address) (*EventListener, error) {
	client, err := ethclient.Dial(wsurl)
	if err != nil {
		return nil, err
	}

	listener := &EventListener{
		client:          client,
		outputChan:      outputChan,
		contract:        contractAddress,
		deployBlock:     deployedBlock,
		storageFileName: storageFileName,
	}

	err = listener.loadLastBlock()
	if err != nil {
		return nil, fmt.Errorf("failed to load last block: %w", err)
	}
	listener.filterQuery = ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: listener.lastBlock,
	}

	return listener, nil
}

func (l *EventListener) loadLastBlock() error {
	data, err := os.ReadFile(l.storageFileName)
	if err != nil {
		l.lastBlock = l.deployBlock
		log.Warn("Failed to load last block, starting from deploy block",
			"deployBlock", l.deployBlock,
		)
	} else {
		var lastBlock big.Int
		if err := json.Unmarshal(data, &lastBlock); err != nil {
			log.Error("Failed to unmarshal last block: %v", err)
			return fmt.Errorf("failed to unmarshal last block")
		}
		l.lastBlock = &lastBlock
	}

	return nil
}

func (l *EventListener) saveLastBlock() error {
	data, err := json.Marshal(l.lastBlock)
	if err != nil {
		return fmt.Errorf("failed to marshal last block")
	}

	if err := os.WriteFile("last_block.json", data, 0600); err != nil {
		return fmt.Errorf("failed to write last block to file")
	}
	return nil
}

func (l *EventListener) Listen() {
outerLoop:
	for {
		logs := make(chan types.Log)
		sub, err := l.client.SubscribeFilterLogs(context.Background(), l.filterQuery, logs)
		if err != nil {
			log.Error("Failed to subscribe to logs: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for {
			select {
			case lg := <-logs:
				l.outputChan <- lg
				l.lastBlock = big.NewInt(0).SetUint64(lg.BlockNumber)
				err = l.saveLastBlock()
				if err != nil {
					log.Error("Failed to save last block: %v", err)
				}
			case err := <-sub.Err():
				log.Error("Subscription error: %v", err)
				break outerLoop
			}
		}
	}
}
