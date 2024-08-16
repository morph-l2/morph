package main

import (
	"context"
	"fmt"
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
	batch, err := l2Client.GetRollupBatchByIndex(context.Background(), batchIndex)
	if err != nil {
		panic(err)
	}
	batchInfo := new(derivation.BatchInfo)
	if err = batchInfo.ParseBatch(*batch); err != nil {
		panic(err)
	}
	fmt.Println("batch index: ", batchIndex)
	fmt.Println("batch blocks: ", batchInfo.BlockNum())
	fmt.Println("batch txs: ", batchInfo.TxNum())
}
