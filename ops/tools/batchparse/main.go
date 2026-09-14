package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/morph-l2/go-ethereum/ethclient"

	"morph-l2/node/derivation"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	rpcURL := flag.String("rpc", os.Getenv("L2_RPC_URL"), "L2 JSON-RPC URL (or L2_RPC_URL)")
	batchIndex := flag.Uint64("batch", 0, "Batch index to parse (must be specified)")
	timeout := flag.Duration("timeout", 30*time.Second, "RPC request timeout")
	flag.Parse()
	batchSet := false
	flag.Visit(func(f *flag.Flag) { batchSet = batchSet || f.Name == "batch" })
	if flag.NArg() != 0 || *rpcURL == "" || !batchSet || *timeout <= 0 {
		return errors.New("require -rpc (or L2_RPC_URL), -batch, and a positive -timeout")
	}
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	client, err := ethclient.DialContext(ctx, *rpcURL)
	if err != nil {
		return err
	}
	defer client.Close()
	batch, err := client.GetRollupBatchByIndex(ctx, *batchIndex)
	if err != nil {
		return err
	}
	if batch == nil {
		return fmt.Errorf("batch %d is not available", *batchIndex)
	}
	batchInfo := new(derivation.BatchInfo)
	if err := batchInfo.ParseBatch(*batch); err != nil {
		return err
	}
	fmt.Println("batch index:", *batchIndex)
	fmt.Println("batch blocks:", batchInfo.BlockNum())
	fmt.Println("batch txs:", batchInfo.TxNum())
	return nil
}
