package common

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"

	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
)

type DeployContractBackend interface {
	bind.DeployBackend
	bind.ContractBackend
	ethereum.ChainReader
}

func GetLatestConfirmedBlockNumber(ctx context.Context, l1Client *ethclient.Client, confirmations rpc.BlockNumber) (uint64, error) {
	// confirmation based on "safe" or "finalized" block tag
	if confirmations == rpc.SafeBlockNumber || confirmations == rpc.FinalizedBlockNumber {
		tag := big.NewInt(int64(confirmations))
		header, err := l1Client.HeaderByNumber(ctx, tag)
		if err != nil {
			return 0, err
		}
		if !header.Number.IsInt64() {
			return 0, fmt.Errorf("received invalid block confirm: %v", header.Number)
		}
		return header.Number.Uint64(), nil
	}

	// confirmation based on latest block number
	if confirmations == rpc.LatestBlockNumber {
		number, err := l1Client.BlockNumber(ctx)
		if err != nil {
			return 0, err
		}
		return number, nil
	}

	// confirmation based on a certain number of blocks
	if confirmations.Int64() >= 0 {
		number, err := l1Client.BlockNumber(ctx)
		if err != nil {
			return 0, err
		}
		confirmations := uint64(confirmations.Int64())
		if number >= confirmations {
			return number - confirmations, nil
		}
		return 0, nil
	}

	return 0, fmt.Errorf("unknown confirmation type: %v", confirmations)
}
