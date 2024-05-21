package oracle

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	"morph-l2/oracle/config"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	tmhttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultRewardEpoch = time.Hour / time.Second * 24
	defaultPrecision   = 10 ^ 8
	defaultSleepTime   = 30 * time.Second
)

func Main() func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cfg, err := config.NewConfig(ctx)
		if err != nil {
			return err
		}
		log.Info("Initializing staking-oracle")
		o, err := NewOracle(&cfg)
		if err != nil {
			log.Error("Unable to create staking-oracle", "error", err)
			return err
		}
		log.Info("Starting staking-oracle")
		o.Start()
		log.Info("Staking oracle started")

		<-(chan struct{})(nil)
		log.Info("staking oracle stoped")
		return nil
	}
}

type Oracle struct {
	ctx                 context.Context
	l1Client            *ethclient.Client
	l2Client            *ethclient.Client
	l2Staking           *bindings.L2Staking
	sequencer           *bindings.Sequencer
	gov                 *bindings.Gov
	rollup              *bindings.Rollup
	record              *bindings.Record
	TmClient            *tmhttp.HTTP
	rewardEpoch         time.Duration
	cfg                 *config.Config
	privKey             *ecdsa.PrivateKey
	isFinalized         bool
	rollupEpochMaxBlock uint64
}

func NewOracle(cfg *config.Config) (*Oracle, error) {
	var logHandler log.Handler
	output := io.Writer(os.Stderr)
	if cfg.LogFilename != "" {
		f, err := os.OpenFile(cfg.LogFilename, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
		if err != nil {
			return nil, fmt.Errorf("wrong log.filename set: %d", err)
		}
		f.Close()

		if cfg.LogFileMaxSize < 1 {
			return nil, fmt.Errorf("wrong log.maxsize set: %d", cfg.LogFileMaxSize)
		}

		if cfg.LogFileMaxAge < 1 {
			return nil, fmt.Errorf("wrong log.maxage set: %d", cfg.LogFileMaxAge)
		}
		logFile := &lumberjack.Logger{
			Filename: cfg.LogFilename,
			MaxSize:  cfg.LogFileMaxSize, // megabytes
			MaxAge:   cfg.LogFileMaxAge,  // days
			Compress: cfg.LogCompress,
		}
		output = io.MultiWriter(output, logFile)
	}

	if cfg.LogTerminal {
		logHandler = log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	} else {
		logHandler = log.StreamHandler(output, log.JSONFormat())
	}

	logLevel, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))
	l1Client, err := ethclient.Dial(cfg.L1EthRpc)
	if err != nil {
		panic(err)
	}
	l2Client, err := ethclient.Dial(cfg.L2EthRpc)
	if err != nil {
		panic(err)
	}
	tmClient, err := tmhttp.New(cfg.TendermintRpc, cfg.WsEndpoint)
	if err != nil {
		panic(err)
	}

	rollup, err := bindings.NewRollup(cfg.RollupAddr, l1Client)
	if err != nil {
		panic(err)
	}
	l2Staking, err := bindings.NewL2Staking(predeploys.L2StakingAddr, l2Client)
	if err != nil {
		panic(err)
	}
	record, err := bindings.NewRecord(predeploys.RecordAddr, l2Client)
	if err != nil {
		panic(err)
	}
	sequencer, err := bindings.NewSequencer(predeploys.SequencerAddr, l2Client)
	if err != nil {
		panic(err)
	}
	gov, err := bindings.NewGov(predeploys.GovAddr, l2Client)
	if err != nil {
		panic(err)
	}
	hex := strings.TrimPrefix(cfg.PrivKey, "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}

	return &Oracle{
		l1Client:            l1Client,
		l2Client:            l2Client,
		rollup:              rollup,
		l2Staking:           l2Staking,
		record:              record,
		sequencer:           sequencer,
		gov:                 gov,
		TmClient:            tmClient,
		cfg:                 cfg,
		rewardEpoch:         defaultRewardEpoch,
		privKey:             privKey,
		ctx:                 context.TODO(),
		rollupEpochMaxBlock: cfg.MaxSize,
	}, nil
}

func (o *Oracle) Start() {
	go func() {
		o.setStartBlock()
		for {
			if err := o.syncRewardEpoch(); err != nil {
				log.Error("syncReward Epoch failed", "error", err)
				time.Sleep(30 * time.Second)
			}
		}
	}()

	go func() {
		for {
			if err := o.submitRecord(); err != nil {
				log.Error("reward submission batch failed", "error", err)
				time.Sleep(30 * time.Second)
			}
		}
	}()

	go func() {
		for {
			if err := o.recordRollupEpoch(); err != nil {
				log.Error("record rollup epoch failed", "error", err)
				time.Sleep(30 * time.Second)
			}
		}
	}()
}

func (o *Oracle) waitReceiptWithCtx(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("timeout")
		case <-t.C:
			receipt, err := o.l2Client.TransactionReceipt(o.ctx, txHash)
			if errors.Is(err, ethereum.NotFound) {
				continue
			}
			if err != nil {
				return nil, err
			}
			if receipt != nil {
				t.Stop()
				return receipt, nil
			}
		}
	}
}
