package validator

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
)

type Validator struct {
	cli             DeployContractBackend
	privateKey      *ecdsa.PrivateKey
	l1ChainID       *big.Int
	contract        *bindings.Rollup
	challengeEnable bool
	logger          tmlog.Logger
}

type DeployContractBackend interface {
	bind.DeployBackend
	bind.ContractBackend
}

func NewValidator(cfg *Config, rollup *bindings.Rollup, logger tmlog.Logger) (*Validator, error) {
	cli, err := ethclient.Dial(cfg.l1RPC)
	if err != nil {
		return nil, fmt.Errorf("dial l1 node error:%v", err)
	}
	return &Validator{
		cli:        cli,
		contract:   rollup,
		privateKey: cfg.PrivateKey,
		l1ChainID:  cfg.L1ChainID,
		logger:     logger,
	}, nil
}

func (v *Validator) SetLogger() {
	v.logger = v.logger.With("module", "validator")
}

func (v *Validator) ChallengeEnable() bool {
	return v.challengeEnable
}

func (v *Validator) ChallengeState(batchIndex uint64) error {
	if !v.ChallengeEnable() {
		return fmt.Errorf("The challenge is not enabled,please set challengeEnable is true")
	}
	opts, err := bind.NewKeyedTransactorWithChainID(v.privateKey, v.l1ChainID)
	if err != nil {
		return err
	}
	gasPrice, err := v.cli.SuggestGasPrice(opts.Context)
	if err != nil {
		return err
	}
	opts.GasPrice = gasPrice
	opts.NoSend = true
	//publicKey := v.privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }

	//receiver := crypto.PubkeyToAddress(*publicKeyECDSA)
	tx, err := v.contract.ChallengeState(opts, batchIndex)
	if err != nil {
		return err
	}
	log.Info("send ChallengeState transaction ", "txHash", tx.Hash().Hex())
	if err := v.cli.SendTransaction(context.Background(), tx); err != nil {
		return err
	}
	// Wait for the receipt
	receipt, err := waitForReceipt(v.cli, tx)
	if err != nil {
		return err
	}
	log.Info("Validator has already started the challenge", "hash", tx.Hash().Hex(),
		"gas-used", receipt.GasUsed, "blocknumber", receipt.BlockNumber)
	return nil
}

func waitForReceipt(backend DeployContractBackend, tx *ethtypes.Transaction) (*ethtypes.Receipt, error) {
	t := time.NewTicker(300 * time.Millisecond)
	receipt := new(ethtypes.Receipt)
	var err error
	for range t.C {
		receipt, err = backend.TransactionReceipt(context.Background(), tx.Hash())
		if errors.Is(err, ethereum.NotFound) {
			continue
		}
		if err != nil {
			return nil, err
		}
		if receipt != nil {
			t.Stop()
			break
		}
	}
	return receipt, nil
}
