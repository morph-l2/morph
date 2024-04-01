package derivation

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/node/sync"
	"github.com/morph-l2/node/types"
	"github.com/morph-l2/node/validator"
	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	eth "github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	geth "github.com/scroll-tech/go-ethereum/eth"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/ethclient/authclient"
	"github.com/scroll-tech/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

var (
	RollupEventTopic     = "CommitBatch(uint256,bytes32)"
	RollupEventTopicHash = crypto.Keccak256Hash([]byte(RollupEventTopic))
)

type Derivation struct {
	ctx                   context.Context
	syncer                *sync.Syncer
	l1Client              DeployContractBackend
	RollupContractAddress common.Address
	confirmations         rpc.BlockNumber
	l2Client              *types.RetryableClient
	validator             *validator.Validator
	logger                tmlog.Logger
	rollup                *bindings.Rollup
	metrics               *Metrics
	l1BeaconClient        *L1BeaconClient

	latestDerivation uint64
	db               Database

	cancel context.CancelFunc

	fetchBlockRange     uint64
	preBatchLastBlock   uint64
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

func NewDerivationClient(ctx context.Context, cfg *Config, syncer *sync.Syncer, db Database, validator *validator.Validator, rollup *bindings.Rollup, logger tmlog.Logger) (*Derivation, error) {
	l1Client, err := ethclient.Dial(cfg.L1.Addr)
	if err != nil {
		return nil, err
	}
	aClient, err := authclient.DialContext(context.Background(), cfg.L2.EngineAddr, cfg.L2.JwtSecret)
	if err != nil {
		return nil, err
	}
	eClient, err := ethclient.Dial(cfg.L2.EthAddr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	logger = logger.With("module", "derivation")
	metrics := PrometheusMetrics("morphnode")
	if cfg.MetricsServerEnable {
		go func() {
			_, err := metrics.Serve(cfg.MetricsHostname, cfg.MetricsPort)
			if err != nil {
				panic(fmt.Errorf("metrics server start error:%v", err))
			}
		}()
		logger.Info("metrics server enabled", "host", cfg.MetricsHostname, "port", cfg.MetricsPort)
	}
	baseHttp := NewBasicHTTPClient(cfg.BeaconRpc, logger)
	l1BeaconClient := NewL1BeaconClient(baseHttp)
	return &Derivation{
		ctx:                   ctx,
		db:                    db,
		l1Client:              l1Client,
		syncer:                syncer,
		validator:             validator,
		rollup:                rollup,
		logger:                logger,
		RollupContractAddress: cfg.RollupContractAddress,
		confirmations:         cfg.L1.Confirmations,
		l2Client:              types.NewRetryableClient(aClient, eClient, tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout))),
		cancel:                cancel,
		stop:                  make(chan struct{}),
		fetchBlockRange:       cfg.FetchBlockRange,
		pollInterval:          cfg.PollInterval,
		logProgressInterval:   cfg.LogProgressInterval,
		metrics:               metrics,
		l1BeaconClient:        l1BeaconClient,
	}, nil
}

func (d *Derivation) Start() {
	// block node startup during initial sync and print some helpful logs
	go func() {
		d.syncer.Start()
		t := time.NewTicker(d.pollInterval)
		defer t.Stop()

		for {
			// don't wait for ticker during startup
			d.derivationBlock(d.ctx)

			select {
			case <-d.ctx.Done():
				d.logger.Error("derivation node Unexpected exit")
				close(d.stop)
				return
			case <-t.C:
				continue
			}
		}
	}()
}

func (d *Derivation) Stop() {
	if d == nil {
		return
	}

	d.logger.Info("stopping derivation service")

	if d.cancel != nil {
		d.cancel()
	}
	<-d.stop
	d.logger.Info("derivation service is stopped")
}

func (d *Derivation) derivationBlock(ctx context.Context) {
	latestDerivation := d.db.ReadLatestDerivationL1Height()
	latest := d.syncer.LatestSynced()
	start := *latestDerivation + 1
	end := latest
	if latest < start {
		d.logger.Info("latest less than start", "latest", latest, "start", start)
		return
	} else if latest-start >= d.fetchBlockRange {
		end = start + d.fetchBlockRange
	} else {
		end = latest
	}
	d.logger.Info("derivation start pull rollupData form l1", "startBlock", start, "end", end)
	logs, err := d.fetchRollupLog(ctx, start, end)
	if err != nil {
		d.logger.Error("eth_getLogs failed", "err", err)
		return
	}
	latestBatchIndex, err := d.rollup.LastCommittedBatchIndex(nil)
	if err != nil {
		d.logger.Error("query rollup latestCommitted batch Index failed", "err", err)
		return
	}
	d.logger.Info("fetched rollup tx", "txNum", len(logs), "latestBatchIndex", latestBatchIndex)

	for _, lg := range logs {
		batchInfo, err := d.fetchRollupDataByTxHash(lg.TxHash, lg.BlockNumber)
		if err != nil {
			rollupCommitBatch, parseErr := d.rollup.ParseCommitBatch(lg)
			if parseErr != nil {
				d.logger.Error("parse commit batch failed", "err", err)
				return
			}
			// ignore genesis batch
			if rollupCommitBatch.BatchIndex.Uint64() == 0 {
				continue
			}
			d.logger.Error("fetch batch info failed", "txHash", lg.TxHash, "blockNumber", lg.BlockNumber, "error", err)
			return
		}
		d.logger.Info("fetch rollup transaction success", "txNonce", batchInfo.nonce, "txHash", batchInfo.txHash,
			"l1BlockNumber", batchInfo.l1BlockNumber, "firstL2BlockNumber", batchInfo.firstBlockNumber, "lastL2BlockNumber", batchInfo.lastBlockNumber)

		// derivation
		lastHeader, err := d.derive(batchInfo)
		if err != nil {
			d.logger.Error("derive blocks interrupt", "error", err)
			return
		}
		// only last block of batch
		d.logger.Info("batch derivation complete", "batch_index", batchInfo.batchIndex, "currentBatchEndBlock", lastHeader.Number.Uint64())
		d.metrics.SetL2DeriveHeight(lastHeader.Number.Uint64())
		if !bytes.Equal(lastHeader.Root.Bytes(), batchInfo.root.Bytes()) {
			d.metrics.SetBatchStatus(stateException)
			// TODO The challenge switch is currently on and will be turned on in the future
			if d.validator != nil && d.validator.ChallengeEnable() {
				if err := d.validator.ChallengeState(batchInfo.batchIndex); err != nil {
					d.logger.Error("challenge state failed")
					return
				}
			}
			d.logger.Info("root hash is not equal", "originStateRootHash", batchInfo.root, "deriveStateRootHash", lastHeader.Root.Hex())
			return
		} else {
			d.metrics.SetBatchStatus(stateNormal)
		}
		d.db.WriteLatestDerivationL1Height(lg.BlockNumber)
		d.metrics.SetL1SyncHeight(lg.BlockNumber)
		d.logger.Info("write latest derivation l1 height success", "l1BlockNumber", lg.BlockNumber)
	}

	d.db.WriteLatestDerivationL1Height(end)
	d.metrics.SetL1SyncHeight(end)
}

func (d *Derivation) fetchRollupLog(ctx context.Context, from, to uint64) ([]eth.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(from),
		ToBlock:   big.NewInt(0).SetUint64(to),
		Addresses: []common.Address{
			d.RollupContractAddress,
		},
		Topics: [][]common.Hash{
			{RollupEventTopicHash},
		},
	}
	return d.l1Client.FilterLogs(ctx, query)
}

func (d *Derivation) fetchRollupDataByTxHash(txHash common.Hash, blockNumber uint64) (*BatchInfo, error) {
	tx, pending, err := d.l1Client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, err
	}
	if pending {
		return nil, errors.New("pending transaction")
	}
	abi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	args, err := abi.Methods["commitBatch"].Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		return nil, fmt.Errorf("submitBatches Unpack error:%v", err)
	}

	rollupBatchData := args[0].(struct {
		Version                uint8     "json:\"version\""
		ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
		Chunks                 [][]uint8 "json:\"chunks\""
		SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
		PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
		PostStateRoot          [32]uint8 "json:\"postStateRoot\""
		WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		Signature              struct {
			Version   *big.Int   "json:\"version\""
			Signers   []*big.Int "json:\"signers\""
			Signature []uint8    "json:\"signature\""
		} "json:\"signature\""
	})

	// query blob
	block, err := d.l1Client.BlockByNumber(d.ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	indexedBlobHashes := dataAndHashesFromTxs(block.Transactions(), *tx)
	header, err := d.l1Client.HeaderByNumber(d.ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	var bts eth.BlobTxSidecar
	if len(indexedBlobHashes) != 0 {
		bts, err = d.l1BeaconClient.GetBlobSidecar(context.Background(), L1BlockRef{
			Time: header.Time,
		}, indexedBlobHashes)
		if err != nil {
			return nil, fmt.Errorf("getBlockSidecar error:%v", err)
		}
	}

	var chunks []hexutil.Bytes
	for _, chunk := range rollupBatchData.Chunks {
		chunks = append(chunks, chunk)
	}
	batch := geth.RPCRollupBatch{
		Version:                uint(rollupBatchData.Version),
		ParentBatchHeader:      rollupBatchData.ParentBatchHeader,
		Chunks:                 chunks,
		SkippedL1MessageBitmap: rollupBatchData.SkippedL1MessageBitmap,
		PrevStateRoot:          common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
		PostStateRoot:          common.BytesToHash(rollupBatchData.PostStateRoot[:]),
		WithdrawRoot:           common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		Sidecar:                bts,
	}

	rollupData, err := d.parseBatch(batch)
	if err != nil {
		d.logger.Error("parse batch failed", "txNonce", tx.Nonce(), "txHash", txHash,
			"l1BlockNumber", blockNumber)
		return rollupData, fmt.Errorf("parse batch error:%v", err)
	}
	rollupData.l1BlockNumber = blockNumber
	rollupData.txHash = txHash
	rollupData.nonce = tx.Nonce()
	return rollupData, nil
}

func (d *Derivation) parseBatch(batch geth.RPCRollupBatch) (*BatchInfo, error) {
	blobHashes := batch.Sidecar.BlobHashes()
	parentBatchHeader, err := types.DecodeBatchHeader(batch.ParentBatchHeader)
	if err != nil {
		return nil, fmt.Errorf("decode batch header error:%v", err)
	}
	batchInfo := new(BatchInfo)
	if err := batchInfo.ParseBatch(batch); err != nil {
		return nil, fmt.Errorf("parse batch error:%v", err)
	}
	if err := d.handleL1Message(batchInfo, &parentBatchHeader, blobHashes); err != nil {
		return nil, fmt.Errorf("handle l1 message error:%v", err)
	}
	batchInfo.batchIndex = parentBatchHeader.BatchIndex + 1
	return batchInfo, nil
}

func (d *Derivation) handleL1Message(rollupData *BatchInfo, parentBatchHeader *types.BatchHeader, blobHashes []common.Hash) error {
	batchHeader := types.BatchHeader{
		Version:                uint8(rollupData.version),
		BatchIndex:             parentBatchHeader.BatchIndex + 1,
		DataHash:               rollupData.dataHash,
		ParentBatchHash:        parentBatchHeader.ParentBatchHash,
		SkippedL1MessageBitmap: rollupData.skippedL1MessageBitmap.Bytes(),
	}
	var l1MessagePopped, totalL1MessagePopped uint64
	totalL1MessagePopped = parentBatchHeader.TotalL1MessagePopped
	var chunkHashes []byte
	for index, chunk := range rollupData.chunks {
		var chunkTxHashes []byte
		for bIndex, block := range chunk.blockContextes {
			var l1Transactions []*eth.Transaction
			l1Messages, err := d.getL1Message(totalL1MessagePopped, uint64(block.l1MsgNum))
			if err != nil {
				return fmt.Errorf("get l1 message error:%v", err)
			}
			l1MessagePopped += uint64(block.l1MsgNum)
			totalL1MessagePopped += uint64(block.l1MsgNum)
			if len(l1Messages) > 0 {
				for _, l1Message := range l1Messages {
					if rollupData.skippedL1MessageBitmap.Bit(int(l1Message.QueueIndex)-int(parentBatchHeader.TotalL1MessagePopped)) == 1 {
						continue
					}
					transaction := eth.NewTx(&l1Message.L1MessageTx)
					l1Transactions = append(l1Transactions, transaction)
				}
			}
			rollupData.chunks[index].blockContextes[bIndex].SafeL2Data.Transactions = append(encodeTransactions(l1Transactions), chunk.blockContextes[bIndex].SafeL2Data.Transactions...)
		}

		for _, bc := range rollupData.chunks[index].blockContextes {
			txs := decodeTransactions(bc.SafeL2Data.Transactions)
			for _, tx := range txs {
				chunkTxHashes = append(chunkTxHashes, tx.Hash().Bytes()...)
			}
		}
		chunk.Raw.SetTxHashBytes(chunkTxHashes)
		chunkHashes = append(chunkHashes, chunk.Raw.Hash().Bytes()...)
	}
	rollupData.dataHash = crypto.Keccak256Hash(chunkHashes)
	batchHeader.TotalL1MessagePopped = totalL1MessagePopped
	batchHeader.L1MessagePopped = l1MessagePopped
	rollupData.batchHash = types.NewBatchHeaderWithBlobHashes(batchHeader, blobHashes).BatchHash()
	return nil
}

func (d *Derivation) getL1Message(l1MessagePopped, l1MsgNum uint64) ([]types.L1Message, error) {
	if l1MsgNum == 0 {
		return nil, nil
	}
	start := l1MessagePopped
	end := l1MessagePopped + l1MsgNum - 1
	return d.syncer.ReadL1MessagesInRange(start, end), nil
}

func (d *Derivation) derive(rollupData *BatchInfo) (*eth.Header, error) {
	var lastHeader *eth.Header
	for _, chunk := range rollupData.chunks {
		for _, blockData := range chunk.blockContextes {
			batchHash := rollupData.batchHash
			blockData.SafeL2Data.BatchHash = &batchHash
			latestBlockNumber, err := d.l2Client.BlockNumber(context.Background())
			if err != nil {
				return nil, fmt.Errorf("get derivation geth block number error:%v", err)
			}
			if blockData.SafeL2Data.Number <= latestBlockNumber {
				d.logger.Info("new L2 Data block number less than latestBlockNumber", "safeL2DataNumber", blockData.SafeL2Data.Number, "latestBlockNumber", latestBlockNumber)
				lastHeader, err = d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(latestBlockNumber)))
				continue
			}
			err = func() error {
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
				defer cancel()
				lastHeader, err = d.l2Client.NewSafeL2Block(ctx, blockData.SafeL2Data)
				if err != nil {
					d.logger.Error("new l2 block failed", "latestBlockNumber", latestBlockNumber, "error", err)
					return err
				}
				return nil
			}()
			if err != nil {
				return nil, fmt.Errorf("derivation error:%v", err)
			}
			d.logger.Info("new l2 block success", "blockNumber", blockData.Number)
		}
	}

	return lastHeader, nil
}
