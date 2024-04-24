package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"io"
	"math/big"
	"os"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/morph/oracle/backoff"
	"github.com/morph-l2/node/derivation"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	tmhttp "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"gopkg.in/natefinch/lumberjack.v2"
)

var defaultRewardEpoch = time.Hour * 24

type Config struct {
	RollupAddr    common.Address
	L2StakingAddr common.Address
	L1Rpc         string
	L2Rpc         string
	TendermintRpc string
	WsEndpoint    string
	MaxSize       uint64

	LogFilename    string
	LogFileMaxSize int
	LogFileMaxAge  int
	LogCompress    bool
	LogTerminal    bool
	LogLevel       string
}

type Oracle struct {
	ctx                 context.Context
	l1Client            DeployContractBackend
	l2Client            *ethclient.Client
	l2Staking           *bindings.L2Staking
	rollup              *bindings.Rollup
	rollupAddr          common.Address
	record              *bindings.Record
	TmClient            *tmhttp.HTTP
	cancel              context.CancelFunc
	pollInterval        time.Duration
	rewardEpoch         time.Duration
	logProgressInterval time.Duration
	stop                chan struct{}
	cfg                 *Config
}

type DeployContractBackend interface {
	bind.DeployBackend
	bind.ContractBackend
	ethereum.ChainReader
	ethereum.TransactionReader
}

func NewOracle(cfg *Config) (*Oracle, error) {
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
	l1Client, err := ethclient.Dial(cfg.L1Rpc)
	if err != nil {
		panic(err)
	}
	l2Client, err := ethclient.Dial(cfg.L2Rpc)
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
	l2Staking, err := bindings.NewL2Staking(cfg.L2StakingAddr, l2Client)

	return &Oracle{
		l1Client:  l1Client,
		l2Client:  l2Client,
		rollup:    rollup,
		l2Staking: l2Staking,
		TmClient:  tmClient,
		cfg:       cfg,
	}, nil
}

func (o *Oracle) loop() {
	// block node startup during initial sync and print some helpful logs

	go func() {
		t := time.NewTicker(o.pollInterval)
		defer t.Stop()

		for {
			// don't wait for ticker during startup
			o.syncRewardEpochs()
			select {
			case <-o.ctx.Done():
				log.Error("derivation node Unexpected exit")
				close(o.stop)
				return
			case <-t.C:
				continue
			}
		}
	}()
}

func (o *Oracle) syncRewardEpochs() error {
	finalizedBlockTime, err := o.getFinalizedBlockTime()
	if err != nil {
		return err
	}
	syncedBlockTime := o.getSyncedBlockTime()
	if finalizedBlockTime > syncedBlockTime {
		epochNum := int64(finalizedBlockTime) - int64(syncedBlockTime)/o.rewardEpoch.Nanoseconds()
		for i := 0; i < int(epochNum); i++ {
			o.syncRewardEpoch()
		}
	}
	return nil
}

func (o *Oracle) getSyncedBlockTime() uint64 {
	// TODO
	return 0
}

func (o *Oracle) getFinalizedBlockTime() (uint64, error) {
	latestFinalized, err := o.rollup.LastFinalizedBatchIndex(nil)
	if err != nil {
		return 0, err
	}
	batch, err := o.l2Client.GetRollupBatchByIndex(context.Background(), latestFinalized.Uint64())
	if err != nil {
		return 0, err
	}
	if batch == nil {
		// TODO
		return 0, fmt.Errorf("batch not found")
	}
	var batchData derivation.BatchInfo
	if err = batchData.ParseBatch(*batch); err != nil {
		return 0, fmt.Errorf("parse batch error:%v", err)
	}
	lastBlockNumber := batchData.LastBlockNumber()
	header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(lastBlockNumber)))
	if err != nil {
		return 0, err
	}
	return header.Time, nil
}

func (o *Oracle) syncRewardEpoch() {

	recordRewardEpochInfo, err := o.getRewardEpochs()
	if err != nil {
		return
	}
	tx, err := o.record.RecordRewardEpochs(nil, []bindings.IRecordRewardEpochInfo{*recordRewardEpochInfo})
	receipt, err := o.l2Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {

	}
	if receipt.Status != types.ReceiptStatusSuccessful {

	}
}

func (o *Oracle) getRewardEpochs() (*bindings.IRecordRewardEpochInfo, error) {
	// reward
	// TODO
	nextRewardEpochIndex, err := o.record.NextRewardEpochIndex(nil)
	if err != nil {
		return nil, err
	}
	startHeight := o.getNextHeight()
	nextTime, err := o.getNextTime(nextRewardEpochIndex)
	if err != nil {
		return nil, err
	}
	height := startHeight
	sequencersBlockCount := make(map[common.Address]int64)
	for {
		tmHeader, err := o.L2HeaderByNumberWithRetry(height)
		if err != nil {
			return nil, err
		}
		if tmHeader.Time.Unix() > nextTime.Int64() {
			break
		}
		sequencer := o.getSequencer(tmHeader.ProposerAddress)
		sequencersBlockCount[sequencer] += 1
	}
	var sequencers []common.Address
	var blockCounts []*big.Int
	for seq, count := range sequencersBlockCount {
		sequencers = append(sequencers, seq)
		blockCounts = append(blockCounts, big.NewInt(count))
	}
	rewardEpochInfo := bindings.IRecordRewardEpochInfo{
		Index:           nextRewardEpochIndex,
		BlockCount:      big.NewInt(height - startHeight + 1),
		Sequencers:      sequencers,
		SequencerBlocks: blockCounts,
		// TODO
		SequencerRatios:      nil,
		SequencerCommissions: nil,
	}
	return &rewardEpochInfo, nil
}

func (o *Oracle) getHeader(height int64) (*tmtypes.Header, error) {
	headerResp, err := o.TmClient.Header(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	return headerResp.Header, nil
}

// L2HeaderByNumberWithRetry retries getting headers.
func (o *Oracle) L2HeaderByNumberWithRetry(height int64) (*tmtypes.Header, error) {
	var res *tmtypes.Header
	err := backoff.DoCtx(o.ctx, 3, backoff.Exponential(), func() error {
		var err error
		headerResp, err := o.TmClient.Header(context.Background(), &height)
		if err != nil {
			return err
		}
		res = headerResp.Header
		return nil
	})
	return res, err
}

func (o *Oracle) getSequencer(proposerAddress tmtypes.Address) common.Address {
	// TODO
	return common.Address{}
}

func (o *Oracle) getNextHeight() int64 {
	// TODO
	return 0
}

func (o *Oracle) getNextTime(nextRewardEpochIndex *big.Int) (*big.Int, error) {
	// TODO check start index 0 or 1
	startTime, err := o.l2Staking.REWARDSTARTTIME(nil)
	if err != nil {
		return nil, err
	}
	return startTime.Add(startTime, nextRewardEpochIndex.Add(nextRewardEpochIndex, big.NewInt(int64(o.rewardEpoch)))), nil
}
