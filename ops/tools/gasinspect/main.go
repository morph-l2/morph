package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/params"

	"morph-l2/bindings/bindings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	rpcURL := flag.String("rpc", os.Getenv("L1_RPC_URL"), "L1 JSON-RPC URL (or L1_RPC_URL)")
	txHash := flag.String("tx", "", "Transaction hash to inspect")
	timeout := flag.Duration("timeout", 30*time.Second, "RPC request timeout")
	flag.Parse()
	rawHash, err := hexutil.Decode(*txHash)
	if flag.NArg() != 0 || *rpcURL == "" || err != nil || len(rawHash) != common.HashLength || *timeout <= 0 {
		return errors.New("require -rpc (or L1_RPC_URL), a 32-byte -tx hash, and a positive -timeout")
	}
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	client, err := ethclient.DialContext(ctx, *rpcURL)
	if err != nil {
		return err
	}
	defer client.Close()
	hash := common.BytesToHash(rawHash)
	tx, pending, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		return err
	}
	if pending {
		return errors.New("transaction is pending; gas usage is not available")
	}
	method, err := rollupMethod(tx.Data())
	if err != nil {
		return err
	}
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		return err
	}
	fmt.Printf("tx hash: %s\nmethod: %s\nstatus: %d\ngas used: %d\ncalldata bytes: %d\nEIP-2028 calldata gas: %d\nblob gas used: %d\n",
		hash.Hex(), method, receipt.Status, receipt.GasUsed, len(tx.Data()), dataGasCost(tx.Data()), receipt.BlobGasUsed)
	if receipt.EffectiveGasPrice != nil {
		fmt.Printf("effective gas price (wei): %s\n", receipt.EffectiveGasPrice)
	}
	if receipt.BlobGasPrice != nil {
		fmt.Printf("blob gas price (wei): %s\n", receipt.BlobGasPrice)
	}
	return nil
}

func rollupMethod(data []byte) (string, error) {
	if len(data) < 4 {
		return "", errors.New("transaction has no contract method selector")
	}
	contractABI, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return "", err
	}
	method, err := contractABI.MethodById(data[:4])
	if err != nil {
		return "", fmt.Errorf("selector %x is not in the current Rollup ABI: %w", data[:4], err)
	}
	if _, err := method.Inputs.Unpack(data[4:]); err != nil {
		return "", fmt.Errorf("invalid %s calldata: %w", method.Name, err)
	}
	return method.Name, nil
}

// This is the EIP-2028 byte cost, not execution gas or a post-Pectra gas floor.
func dataGasCost(data []byte) uint64 {
	var cost uint64
	for _, b := range data {
		if b == 0 {
			cost += params.TxDataZeroGas
		} else {
			cost += params.TxDataNonZeroGasEIP2028
		}
	}
	return cost
}
