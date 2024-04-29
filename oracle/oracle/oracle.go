package oracle

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/big"
	"os"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/morph/oracle/backoff"
	"github.com/morph-l2/morph/oracle/config"
	"github.com/morph-l2/node/derivation"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmhttp "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultRewardEpoch = time.Hour * 24
	defaultPrecision   = 10 ^ 8
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

		return nil
	}
}

type Oracle struct {
	ctx                 context.Context
	l1Client            *ethclient.Client
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
	cfg                 *config.Config
	sequencerMap        map[string]common.Address
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
	l2Staking, err := bindings.NewL2Staking(cfg.L2StakingAddr, l2Client)

	return &Oracle{
		l1Client:    l1Client,
		l2Client:    l2Client,
		rollup:      rollup,
		l2Staking:   l2Staking,
		TmClient:    tmClient,
		cfg:         cfg,
		rewardEpoch: defaultRewardEpoch,
	}, nil
}

func (o *Oracle) Start() {
	go func() {
		for {
			if err := o.syncRewardEpoch(); err != nil {
				log.Error("syncReward Epoch error")
				time.Sleep(30 * time.Second)
			}
		}
	}()
}

func (o *Oracle) getFinalizedBlockTimeAndNumber() (uint64, *big.Int, error) {
	latestFinalized, err := o.rollup.LastFinalizedBatchIndex(nil)
	if err != nil {
		return 0, nil, err
	}
	batch, err := o.l2Client.GetRollupBatchByIndex(context.Background(), latestFinalized.Uint64())
	if err != nil {
		return 0, nil, err
	}
	if batch == nil {
		return 0, nil, fmt.Errorf("batch not found")
	}
	var batchData derivation.BatchInfo
	if err = batchData.ParseBatch(*batch); err != nil {
		return 0, nil, fmt.Errorf("parse batch error:%v", err)
	}
	lastBlockNumber := batchData.LastBlockNumber()
	header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(lastBlockNumber)))
	if err != nil {
		return 0, nil, err
	}
	return header.Time, header.Number, nil
}

func (o *Oracle) syncRewardEpoch() error {
	_, finalizedBlock, err := o.getFinalizedBlockTimeAndNumber()
	startRewardEpochIndex, err := o.record.NextRewardEpochIndex(nil)
	if err != nil {
		return err
	}
	startHeight, err := o.getNextHeight()
	if startHeight.Cmp(finalizedBlock) > 0 {
		time.Sleep(30 * time.Second)
		return nil
	}
	recordRewardEpochInfo, err := o.getRewardEpochs(startRewardEpochIndex, startHeight)
	if err != nil {
		return err
	}
	tx, err := o.record.RecordRewardEpochs(nil, []bindings.IRecordRewardEpochInfo{*recordRewardEpochInfo})
	receipt, err := o.l2Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("record reward epochs not success")
	}
	return nil
}

func (o *Oracle) getRewardEpochs(startRewardEpochIndex, startHeight *big.Int) (*bindings.IRecordRewardEpochInfo, error) {
	endTime, err := o.getEndTime(startHeight, startRewardEpochIndex)
	if err != nil {
		return nil, err
	}
	height := startHeight
	sequencersBlockCount := make(map[common.Address]int64)
	for {
		tmHeader, err := o.L2HeaderByNumberWithRetry(height.Int64())
		if err != nil {
			return nil, fmt.Errorf("get l2 header error:%v", err)
		}
		if tmHeader.Time.Unix() > endTime.Int64() {
			break
		}
		sequencer, err := o.getSequencer(tmHeader.ProposerAddress, height)
		if err != nil {
			return nil, fmt.Errorf("get sequencer error:%v", err)
		}
		sequencersBlockCount[sequencer] += 1
		height = new(big.Int).Add(height, big.NewInt(1))
	}
	var sequencers []common.Address
	var seqBlockCounts, sequencerRatios, sequencerCommissions []*big.Int
	for seq, count := range sequencersBlockCount {
		sequencers = append(sequencers, seq)
		seqBlockCounts = append(seqBlockCounts, big.NewInt(count))
	}
	blockCount := height.Add(height.Sub(height, startHeight), big.NewInt(1))
	precision := big.NewInt(defaultPrecision)
	residue := big.NewInt(defaultPrecision)
	maxRatio := big.NewInt(0)
	var maxRatioIndex int
	for i := 0; i < len(sequencers); i++ {
		ratio := new(big.Int).Div(new(big.Int).Mul(seqBlockCounts[i], precision), blockCount)
		sequencerRatios = append(sequencerRatios, ratio)
		residue = new(big.Int).Sub(residue, ratio)
		if ratio.Cmp(maxRatio) > 0 {
			maxRatio = ratio
			maxRatioIndex = i
		}
		commission, err := o.getSequencerCommission(new(big.Int).Sub(startHeight, big.NewInt(1)), sequencers[i])
		if err != nil {
			return nil, fmt.Errorf("get sequencer commission error:%v", err)
		}
		sequencerCommissions = append(sequencerCommissions, commission)
	}
	sequencerRatios[maxRatioIndex] = new(big.Int).Add(sequencerRatios[maxRatioIndex], residue)
	rewardEpochInfo := bindings.IRecordRewardEpochInfo{
		Index:                startRewardEpochIndex,
		BlockCount:           height.Add(height.Sub(height, startHeight), big.NewInt(1)),
		Sequencers:           sequencers,
		SequencerBlocks:      seqBlockCounts,
		SequencerRatios:      sequencerRatios,
		SequencerCommissions: sequencerCommissions,
	}
	return &rewardEpochInfo, nil
}

func (o *Oracle) getSequencerCommission(blockNumber *big.Int, address common.Address) (*big.Int, error) {
	if blockNumber.Uint64() < o.cfg.StartBlock {
		return big.NewInt(0), nil
	}
	return o.l2Staking.Commissions(&bind.CallOpts{
		BlockNumber: blockNumber,
	}, address)
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

func (o *Oracle) getSequencer(proposerAddress tmtypes.Address, blockNumber *big.Int) (common.Address, error) {
	stakers, err := o.l2Staking.GetStakers(&bind.CallOpts{
		BlockNumber: new(big.Int).Sub(blockNumber, big.NewInt(1)),
	})
	if err != nil {
		return common.Address{}, err
	}
	for _, staker := range stakers {
		if bytes.Compare(proposerAddress, ed25519.PubKey(staker.TmKey[:]).Address().Bytes()) == 0 {
			return staker.Addr, nil
		}
	}
	return common.Address{}, fmt.Errorf("sequencer not found")
}

func (o *Oracle) getNextHeight() (*big.Int, error) {
	return o.record.LatestRewardEpochBlock(nil)
}

func (o *Oracle) getEndTime(blockNumber *big.Int, nextRewardEpochIndex *big.Int) (*big.Int, error) {
	startTime, err := o.l2Staking.RewardStartTime(&bind.CallOpts{
		BlockNumber: blockNumber,
	})
	if err != nil {
		return nil, err
	}
	internal := new(big.Int).Add(nextRewardEpochIndex, big.NewInt(int64(o.rewardEpoch)))
	epochStart := new(big.Int).Add(startTime, internal)
	epochEnd := new(big.Int).Add(epochStart, internal)
	return epochEnd, nil
}
