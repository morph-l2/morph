package oracle

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/scroll-tech/go-ethereum"
	"io"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmhttp "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	"morph-l2/node/derivation"
	"morph-l2/oracle/backoff"
	"morph-l2/oracle/config"
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

		return nil
	}
}

type Oracle struct {
	ctx                 context.Context
	l1Client            *ethclient.Client
	l2Client            *ethclient.Client
	l2Staking           *bindings.L2Staking
	rollup              *bindings.Rollup
	record              *bindings.Record
	TmClient            *tmhttp.HTTP
	cancel              context.CancelFunc
	pollInterval        time.Duration
	rewardEpoch         time.Duration
	logProgressInterval time.Duration
	stop                chan struct{}
	cfg                 *config.Config
	sequencerMap        map[string]common.Address
	lastRewardStartTime int64
	privKey             *ecdsa.PrivateKey
	isFinalized         bool
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
	hex := strings.TrimPrefix(cfg.PrivKey, "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}

	return &Oracle{
		l1Client:    l1Client,
		l2Client:    l2Client,
		rollup:      rollup,
		l2Staking:   l2Staking,
		record:      record,
		TmClient:    tmClient,
		cfg:         cfg,
		rewardEpoch: defaultRewardEpoch,
		privKey:     privKey,
		ctx:         context.TODO(),
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

func (o *Oracle) getBlockTimeAndNumber(isFinalized bool) (uint64, *big.Int, error) {
	var lastBlockNumber *big.Int
	if isFinalized {
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
		lastBlockNumber = big.NewInt(int64(batchData.LastBlockNumber()))
	}

	header, err := o.l2Client.HeaderByNumber(o.ctx, lastBlockNumber)
	if err != nil {
		return 0, nil, err
	}
	return header.Time, header.Number, nil
}

func (o *Oracle) syncRewardEpoch() error {
	_, finalizedBlock, err := o.getBlockTimeAndNumber(o.isFinalized)
	if err != nil {
		return fmt.Errorf("get block time and number error:%v", err)
	}
	startRewardEpochIndex, err := o.record.NextRewardEpochIndex(nil)
	if err != nil {
		return err
	}
	startHeight, err := o.getNextHeight()
	if err != nil {
		return err
	}
	if startHeight.Cmp(finalizedBlock) > 0 {
		time.Sleep(30 * time.Second)
		return nil
	}
	recordRewardEpochInfo, err := o.getRewardEpochs(startRewardEpochIndex, startHeight)
	if err != nil {
		return err
	}
	chainId, err := o.l2Client.ChainID(o.ctx)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, chainId)
	tx, err := o.record.RecordRewardEpochs(opts, []bindings.IRecordRewardEpochInfo{*recordRewardEpochInfo})
	if err != nil {
		return fmt.Errorf("record reward epochs error:%v", err)
	}
	log.Info("send record reward tx success", "txHash", tx.Hash().Hex(), "nonce", tx.Nonce())
	receipt, err := o.l2Client.TransactionReceipt(o.ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("receipt record reward epochs error:%v", err)
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
	log.Info("new epoch fetching...", "startHeight", startHeight, "startRewardEpochIndex", startRewardEpochIndex, "endTime", endTime)
	height := startHeight
	sequencersBlockCount := make(map[common.Address]int64)
	for {
		_, finalizedBlock, err := o.getBlockTimeAndNumber(o.isFinalized)
		if err != nil {
			continue
		}
		if height.Cmp(finalizedBlock) > 0 {
			log.Info("finalized block small than syncing block,wait...", "finalizedBlock", finalizedBlock, "syncingBlock", height)
			time.Sleep(time.Second * 30)
			continue
		}
		tmHeader, err := o.L2HeaderByNumberWithRetry(height.Int64())
		if err != nil {
			return nil, fmt.Errorf("get l2 header error:%v", err)
		}
		if tmHeader.Time.Unix() > endTime.Int64() {
			break
		}
		log.Info("get new header", "headerNumber", tmHeader.Height, "headerTime", tmHeader.Time)
		sequencer, err := o.getSequencer(tmHeader.ProposerAddress, height)
		if err != nil {
			return nil, fmt.Errorf("get sequencer error:%v", err)
		}
		sequencersBlockCount[sequencer] += 1

		height = new(big.Int).Add(height, big.NewInt(1))
	}

	var sequencers []common.Address
	var seqBlockCounts, sequencerRatios, sequencerCommissions []*big.Int
	var blC int64
	for seq, count := range sequencersBlockCount {
		blC += count
		sequencers = append(sequencers, seq)
		seqBlockCounts = append(seqBlockCounts, big.NewInt(count))
	}
	blockCount := new(big.Int).Sub(height, startHeight)
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
		BlockCount:           blockCount,
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
	latest, err := o.record.LatestRewardEpochBlock(nil)
	if err != nil {
		return latest, err
	}
	return new(big.Int).Add(latest, big.NewInt(1)), nil
}

func (o *Oracle) getEndTime(blockNumber *big.Int, nextRewardEpochIndex *big.Int) (*big.Int, error) {
	startTime, err := o.l2Staking.RewardStartTime(&bind.CallOpts{
		BlockNumber: blockNumber,
	})
	if err != nil {
		return nil, err
	}
	internal := new(big.Int).Mul(nextRewardEpochIndex, big.NewInt(int64(o.rewardEpoch)))
	epochStart := new(big.Int).Add(startTime, internal)
	epochEnd := new(big.Int).Add(epochStart, big.NewInt(int64(o.rewardEpoch)))
	return epochEnd, nil
}

func (o *Oracle) findStartBlock(start, end uint64, timeStamp int64) (int64, error) {
	headerStart, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(start)))
	if err != nil {
		return 0, err
	}
	headerEnd, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(end)))
	if err != nil {
		return 0, err
	}
	if end < start {
		return 0, fmt.Errorf("invalid start or end,start:%v,end:%v", start, end)
	}
	if int64(headerStart.Time) > timeStamp || int64(headerEnd.Time) < timeStamp {
		return 0, fmt.Errorf("this timestamp is not within the given block range")
	}

	s := sort.Search(int(end)+1-int(start), func(i int) bool {
		header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(i)+int64(start)))
		if err != nil {
			log.Error("get header by number failed", "error", err)
			return false
		}
		return int64(header.Time) >= timeStamp
	})
	if s == int(end)+1-int(start) {
		log.Error("start block not found")
	}
	target, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(start)+int64(s)))
	if err != nil {
		return 0, err
	}
	preHeader, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(start)+int64(s)-1))
	if err != nil {
		return 0, err
	}
	if !(int64(preHeader.Time) < timeStamp && int64(target.Time) >= timeStamp) {
		return 0, fmt.Errorf("invalid start block")
	}
	log.Info("find start block success", "preHeader_time", preHeader.Time, "timestamp", timeStamp, "target_time", target.Time)
	return int64(start) + int64(s), nil
}

func (o *Oracle) setStartBlock() {
	start := o.cfg.StartBlock
	for {
		header, err := o.l2Client.HeaderByNumber(o.ctx, nil)
		if err != nil {
			panic(err)
		}
		rewardStart, err := o.l2Staking.RewardStart(&bind.CallOpts{
			BlockNumber: header.Number,
		})
		if err != nil {
			panic(err)
		}
		if rewardStart {
			log.Info("reward start")
			break
		}
		start = header.Number.Uint64()
		log.Info("wait reward start...")
		time.Sleep(defaultSleepTime)
		continue
	}

	for {
		header, err := o.l2Client.HeaderByNumber(o.ctx, nil)
		if err != nil {

		}
		startTime, err := o.l2Staking.RewardStartTime(&bind.CallOpts{
			BlockNumber: header.Number,
		})
		if err != nil {

		}
		latestRewardEpochBlock, err := o.record.LatestRewardEpochBlock(nil)
		if latestRewardEpochBlock.Uint64() != 0 {
			break
		}
		if header.Time < startTime.Uint64() {
			start = header.Number.Uint64()
			time.Sleep(defaultSleepTime)
			continue
		}
		log.Info("start find start block", "start_block", start, "end_block", header.Number.Uint64())
		startBlock, err := o.findStartBlock(start, header.Number.Uint64(), startTime.Int64())
		if err != nil {
			log.Error("find start block failed", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}

		chainID, err := o.l2Client.ChainID(o.ctx)
		if err != nil {
			log.Error("get chain id failed", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}
		opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, chainID)
		nonce, err := o.l2Client.NonceAt(context.Background(), crypto.PubkeyToAddress(o.privKey.PublicKey), nil)
		if err != nil {
			return
		}
		opts.NoSend = true
		opts.Nonce = big.NewInt(int64(nonce))
		tx, err := o.record.SetLatestRewardEpochBlock(opts, big.NewInt(startBlock))
		if err != nil {
			log.Error("set latestReward epoch block failed", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}
		signedTx, err := opts.Signer(opts.From, tx)
		if err != nil {
			return
		}
		err = o.l2Client.SendTransaction(o.ctx, signedTx)
		if err != nil {
			log.Error("send transaction failed,retry...", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}
		receipt, err := o.waitReceiptWithCtx(o.ctx, tx.Hash())
		if err != nil {
			log.Error("TransactionReceipt failed,retry...", "error", err)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Error("set stark block failed")
			continue
		}
		log.Info("set start block success")
	}
}

func (o *Oracle) waitReceiptWithCtx(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	t := time.NewTicker(time.Second)
	receipt := new(types.Receipt)
	var err error
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("timeout")
		case <-t.C:
			receipt, err = o.l2Client.TransactionReceipt(o.ctx, txHash)
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
