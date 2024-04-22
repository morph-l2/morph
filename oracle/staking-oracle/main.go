package main

import (
	"context"
	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum"
	"time"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"
)

type Oracle struct {
	ctx                 context.Context
	l1Client            DeployContractBackend
	l2Client            *ethclient.Client
	staking             bindings.Staking
	cancel              context.CancelFunc
	pollInterval        time.Duration
	logProgressInterval time.Duration
	stop                chan struct{}
}

type DeployContractBackend interface {
	bind.DeployBackend
	bind.ContractBackend
	ethereum.ChainReader
	ethereum.TransactionReader
}

func NewOracle() {

}
