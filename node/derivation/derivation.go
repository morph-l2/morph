package derivation

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	geth "github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	"morph-l2/node/sync"
	"morph-l2/node/types"
	"morph-l2/node/validator"
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
	L2ToL1MessagePasser   *bindings.L2ToL1MessagePasser

	db Database

	cancel context.CancelFunc

	fetchBlockRange     uint64
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
	msgPasser, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, eClient)
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
		L2ToL1MessagePasser:   msgPasser,
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
	}
	d.logger.Info("derivation start pull rollupData form l1", "startBlock", start, "end", end)
	logs, err := d.fetchRollupLog(ctx, start, end)
	if err != nil {
		d.logger.Error("eth_getLogs failed", "err", err)
		return
	}
	latestBatchIndex, err := d.rollup.LastCommittedBatchIndex(&bind.CallOpts{
		BlockNumber: big.NewInt(int64(latest)),
	})
	if err != nil {
		d.logger.Error("query rollup latestCommitted batch Index failed", "err", err)
		return
	}
	d.metrics.SetLatestBatchIndex(latestBatchIndex.Uint64())
	d.logger.Info("fetched rollup tx", "txNum", len(logs), "latestBatchIndex", latestBatchIndex)

	for _, lg := range logs {
		batchInfo, err := d.fetchRollupDataByTxHash(lg.TxHash, lg.BlockNumber)
		if err != nil {
			if errors.Is(err, types.ErrNotCommitBatchTx) {
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
		d.metrics.SetSyncedBatchIndex(batchInfo.batchIndex)
		withdrawalRoot, err := d.L2ToL1MessagePasser.MessageRoot(&bind.CallOpts{
			BlockNumber: lastHeader.Number,
		})
		if err != nil {
			d.logger.Error("get withdrawal root failed", "error", err)
			return
		}
		if !bytes.Equal(lastHeader.Root.Bytes(), batchInfo.root.Bytes()) || !bytes.Equal(withdrawalRoot[:], batchInfo.withdrawalRoot.Bytes()) {
			d.metrics.SetBatchStatus(stateException)
			// TODO The challenge switch is currently on and will be turned on in the future
			if d.validator != nil && d.validator.ChallengeEnable() {
				if err := d.validator.ChallengeState(batchInfo.batchIndex); err != nil {
					d.logger.Error("challenge state failed")
					return
				}
			}
			d.logger.Info("root hash or withdrawal hash is not equal",
				"originStateRootHash", batchInfo.root,
				"deriveStateRootHash", lastHeader.Root.Hex(),
				"batchWithdrawalRoot", batchInfo.withdrawalRoot.Hex(),
				"deriveWithdrawalRoot", common.BytesToHash(withdrawalRoot[:]).Hex(),
			)
			return
		} else {
			d.metrics.SetBatchStatus(stateNormal)
		}
		d.metrics.SetL1SyncHeight(lg.BlockNumber)
	}

	d.db.WriteLatestDerivationL1Height(end)
	d.metrics.SetL1SyncHeight(end)
	d.logger.Info("write latest derivation l1 height success", "l1BlockNumber", end)
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
	if !bytes.Equal(abi.Methods["commitBatch"].ID, tx.Data()[:4]) {
		return nil, types.ErrNotCommitBatchTx
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
	})

	// query blob
	block, err := d.l1Client.BlockByNumber(d.ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	indexedBlobHashes := dataAndHashesFromTxs(block.Transactions(), tx)
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
	parentBatchHeader, err := types.DecodeBatchHeader(batch.ParentBatchHeader)
	if err != nil {
		return nil, fmt.Errorf("decode batch header error:%v", err)
	}
	batchInfo := new(BatchInfo)
	if err := batchInfo.ParseBatch(batch); err != nil {
		return nil, fmt.Errorf("parse batch error:%v", err)
	}
	if err := d.handleL1Message(batchInfo, parentBatchHeader.TotalL1MessagePopped); err != nil {
		return nil, fmt.Errorf("handle l1 message error:%v", err)
	}
	batchInfo.batchIndex = parentBatchHeader.BatchIndex + 1
	return batchInfo, nil
}

func (d *Derivation) handleL1Message(rollupData *BatchInfo, parentTotalL1MessagePopped uint64) error {
	totalL1MessagePopped := parentTotalL1MessagePopped
	for index, chunk := range rollupData.chunks {
		for bIndex, block := range chunk.blockContextes {
			var l1Transactions []*eth.Transaction
			l1Messages, err := d.getL1Message(totalL1MessagePopped, uint64(block.l1MsgNum))
			if err != nil {
				return fmt.Errorf("get l1 message error:%v", err)
			}
			totalL1MessagePopped += uint64(block.l1MsgNum)
			if len(l1Messages) > 0 {
				for _, l1Message := range l1Messages {
					if rollupData.skippedL1MessageBitmap.Bit(int(l1Message.QueueIndex)-int(parentTotalL1MessagePopped)) == 1 {
						continue
					}
					transaction := eth.NewTx(&l1Message.L1MessageTx)
					l1Transactions = append(l1Transactions, transaction)
				}
			}
			rollupData.chunks[index].blockContextes[bIndex].SafeL2Data.Transactions = append(encodeTransactions(l1Transactions), chunk.blockContextes[bIndex].SafeL2Data.Transactions...)
		}
	}
	return nil
}

func (d *Derivation) getL1Message(l1MessagePopped, l1MsgNum uint64) ([]types.L1Message, error) {
	if l1MsgNum == 0 {
		return nil, nil
	}
	return d.syncer.ReadL1MessagesInRange(l1MessagePopped, l1MessagePopped+l1MsgNum-1), nil
}

func (d *Derivation) derive(rollupData *BatchInfo) (*eth.Header, error) {
	var lastHeader *eth.Header
	for _, chunk := range rollupData.chunks {
		for _, blockData := range chunk.blockContextes {
			latestBlockNumber, err := d.l2Client.BlockNumber(context.Background())
			if err != nil {
				return nil, fmt.Errorf("get derivation geth block number error:%v", err)
			}
			if blockData.SafeL2Data.Number <= latestBlockNumber {
				d.logger.Info("new L2 Data block number less than latestBlockNumber", "safeL2DataNumber", blockData.SafeL2Data.Number, "latestBlockNumber", latestBlockNumber)
				lastHeader, err = d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(blockData.SafeL2Data.Number)))
				if err != nil {
					return nil, fmt.Errorf("query header by number error:%v", err)
				}
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
