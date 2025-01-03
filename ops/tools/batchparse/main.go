package main

import (
	"context"
	"fmt"
	"log"
	"morph-l2/node/derivation"

	"github.com/morph-l2/go-ethereum/ethclient"
)

var (
	L2RPC             = "http://localhost:8545"
	batchIndex uint64 = 35
)

func main() {
	l2Client, err := ethclient.Dial(L2RPC)
	if err != nil {
		panic(err)
	}
	parentBatch, err := l2Client.GetRollupBatchByIndex(context.Background(), batchIndex-1)
	if err != nil {
		log.Fatalf("failed to get batch, index: %d, err: %v", batchIndex-1, err)
	}
	batch, err := l2Client.GetRollupBatchByIndex(context.Background(), batchIndex)
	if err != nil {
		log.Fatalf("failed to get batch, index: %d, err: %v", batchIndex, err)
	}
	batchInfo := new(derivation.BatchInfo)
	if err = batchInfo.ParseBatch(*batch, &parentBatch.LastBlockNumber); err != nil {
		panic(err)
	}
	fmt.Println("batch index: ", batchIndex)
	fmt.Println("batch blocks: ", batchInfo.BlockNum())
	fmt.Println("batch txs: ", batchInfo.TxNum())
}
