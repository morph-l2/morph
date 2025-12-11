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
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	geth "github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	nodecommon "morph-l2/node/common"
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
	l1Client              *ethclient.Client
	RollupContractAddress common.Address
	confirmations         rpc.BlockNumber
	l2Client              *types.RetryableClient
	validator             *validator.Validator
	logger                tmlog.Logger
	rollup                *bindings.Rollup
	metrics               *Metrics
	l1BeaconClient        *L1BeaconClient
	L2ToL1MessagePasser   *bindings.L2ToL1MessagePasser

	rollupABI             *abi.ABI
	legacyRollupABI       *abi.ABI // before remove skipMap
	beforeMoveBlockCtxABI *abi.ABI

	db Database

	cancel context.CancelFunc

	startHeight         uint64
	baseHeight          uint64
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
	// TODO
	laClient, err := authclient.DialContext(context.Background(), cfg.L2.EngineAddr, cfg.L2.JwtSecret)
	if err != nil {
		return nil, err
	}
	leClient, err := ethclient.Dial(cfg.L2.EthAddr)
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
	rollupAbi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	legacyRollupAbi, err := types.LegacyRollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	beforeMoveBlockCtxAbi, err := types.BeforeMoveBlockCtxABI.GetAbi()
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
		rollupABI:             rollupAbi,
		legacyRollupABI:       legacyRollupAbi,
		beforeMoveBlockCtxABI: beforeMoveBlockCtxAbi,
		logger:                logger,
		RollupContractAddress: cfg.RollupContractAddress,
		confirmations:         cfg.L1.Confirmations,
		l2Client:              types.NewRetryableClient(laClient, leClient, aClient, eClient, tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout))),
		cancel:                cancel,
		stop:                  make(chan struct{}),
		startHeight:           cfg.StartHeight,
		baseHeight:            cfg.BaseHeight,
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
	latest, err := d.getLatestConfirmedBlockNumber(d.ctx)
	if err != nil {
		d.logger.Error("get latest block number failed", "err", err)
		return
	}
	var start uint64
	if latestDerivation == nil {
		start = d.startHeight
	} else {
		start = *latestDerivation + 1
	}
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
	latestBatchIndex, err := d.rollup.LastCommittedBatchIndex(nil)
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
		if lastHeader.Number.Uint64() <= d.baseHeight {
			continue
		}
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
	batch, err := d.UnPackData(tx.Data())
	if err != nil {
		return nil, err
	}

	// Get block header to retrieve timestamp
	header, err := d.l1Client.HeaderByNumber(d.ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	// Get transaction blob hashes
	blobHashes := tx.BlobHashes()
	if len(blobHashes) > 0 {
		d.logger.Info("Transaction contains blobs", "txHash", txHash, "blobCount", len(blobHashes))

		// Initialize indexedBlobHashes as nil
		var indexedBlobHashes []IndexedBlobHash

		// Only try to build IndexedBlobHash array if not forcing get all blobs
		// Try to get the block to build IndexedBlobHash array
		block, err := d.l1Client.BlockByNumber(d.ctx, big.NewInt(int64(blockNumber)))
		if err == nil {
			// Successfully got the block, now build IndexedBlobHash array
			d.logger.Info("Building IndexedBlobHash array from block", "blockNumber", blockNumber)
			indexedBlobHashes = dataAndHashesFromTxs(block.Transactions(), tx)
			d.logger.Info("Built IndexedBlobHash array", "count", len(indexedBlobHashes))
		} else {
			d.logger.Info("Failed to get block, will try fetching all blobs", "blockNumber", blockNumber, "error", err)
		}

		// Get all blobs corresponding to this timestamp
		blobSidecars, err := d.l1BeaconClient.GetBlobSidecarsEnhanced(d.ctx, L1BlockRef{
			Time: header.Time,
		}, indexedBlobHashes)
		if err != nil {
			return nil, fmt.Errorf("failed to get blobs, continuing processing:%v", err)
		}
		if len(blobSidecars) > 0 {
			// Create blob sidecar
			var blobTxSidecar eth.BlobTxSidecar
			matchedCount := 0

			// Match blobs
			for _, sidecar := range blobSidecars {
				var commitment kzg4844.Commitment
				copy(commitment[:], sidecar.KZGCommitment[:])
				versionedHash := KZGToVersionedHash(commitment)

				for _, expectedHash := range blobHashes {
					if bytes.Equal(versionedHash[:], expectedHash[:]) {
						matchedCount++
						d.logger.Info("Found matching blob", "index", sidecar.Index, "hash", versionedHash.Hex())

						// Decode and process blob data
						var blob Blob
						b, err := hexutil.Decode(sidecar.Blob)
						if err != nil {
							d.logger.Error("Failed to decode blob data", "error", err)
							continue
						}
						copy(blob[:], b)

						// Verify blob
						//if err := VerifyBlobProof(&blob, commitment, kzg4844.Proof(sidecar.KZGProof)); err != nil {
						//	d.logger.Error("Blob verification failed", "error", err)
						//	continue
						//}

						// Add to sidecar
						blobTxSidecar.Blobs = append(blobTxSidecar.Blobs, *blob.KZGBlob())
						blobTxSidecar.Commitments = append(blobTxSidecar.Commitments, commitment)
						blobTxSidecar.Proofs = append(blobTxSidecar.Proofs, kzg4844.Proof(sidecar.KZGProof))
						break
					}
				}
			}

			d.logger.Info("Blob matching results", "matched", matchedCount, "expected", len(blobHashes))
			if matchedCount == 0 {
				return nil, fmt.Errorf("no matching versionedHash was found")
			}
			batch.Sidecar = blobTxSidecar
		} else {
			return nil, fmt.Errorf("not matched blob,txHash:%v,blockNumber:%v", txHash, blockNumber)
		}
	}

	// Get L2 height
	l2Height, err := d.l2Client.BlockNumber(d.ctx)
	if err != nil {
		return nil, fmt.Errorf("query l2 block number error:%v", err)
	}
	rollupData, err := d.parseBatch(batch, l2Height)
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

func (d *Derivation) UnPackData(data []byte) (geth.RPCRollupBatch, error) {
	var batch geth.RPCRollupBatch
	if bytes.Equal(d.beforeMoveBlockCtxABI.Methods["commitBatch"].ID, data[:4]) {
		args, err := d.beforeMoveBlockCtxABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("submitBatches Unpack error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version           uint8     "json:\"version\""
			ParentBatchHeader []uint8   "json:\"parentBatchHeader\""
			BlockContexts     []uint8   "json:\"blockContexts\""
			PrevStateRoot     [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot     [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot    [32]uint8 "json:\"withdrawalRoot\""
		})
		batch = geth.RPCRollupBatch{
			Version:           uint(rollupBatchData.Version),
			ParentBatchHeader: rollupBatchData.ParentBatchHeader,
			BlockContexts:     rollupBatchData.BlockContexts,
			PrevStateRoot:     common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:     common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:      common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
	} else if bytes.Equal(d.legacyRollupABI.Methods["commitBatch"].ID, data[:4]) {
		args, err := d.legacyRollupABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("submitBatches Unpack error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version                uint8     "json:\"version\""
			ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
			BlockContexts          []uint8   "json:\"blockContexts\""
			SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
			PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot          [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		})
		batch = geth.RPCRollupBatch{
			Version:           uint(rollupBatchData.Version),
			ParentBatchHeader: rollupBatchData.ParentBatchHeader,
			BlockContexts:     rollupBatchData.BlockContexts,
			PrevStateRoot:     common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:     common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:      common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
	} else if bytes.Equal(d.rollupABI.Methods["commitBatch"].ID, data[:4]) {
		args, err := d.rollupABI.Methods["commitBatch"].Inputs.Unpack(data[4:])
		if err != nil {
			return batch, fmt.Errorf("submitBatches Unpack error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version           uint8     "json:\"version\""
			ParentBatchHeader []uint8   "json:\"parentBatchHeader\""
			LastBlockNumber   uint64    "json:\"lastBlockNumber\""
			NumL1Messages     uint16    "json:\"numL1Messages\""
			PrevStateRoot     [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot     [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot    [32]uint8 "json:\"withdrawalRoot\""
		})
		batch = geth.RPCRollupBatch{
			Version:           uint(rollupBatchData.Version),
			ParentBatchHeader: rollupBatchData.ParentBatchHeader,
			LastBlockNumber:   rollupBatchData.LastBlockNumber,
			NumL1Messages:     rollupBatchData.NumL1Messages,
			PrevStateRoot:     common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:     common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:      common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
	} else {
		return batch, types.ErrNotCommitBatchTx
	}
	return batch, nil
}

func (d *Derivation) parseBatch(batch geth.RPCRollupBatch, l2Height uint64) (*BatchInfo, error) {
	batchInfo := new(BatchInfo)
	if err := batchInfo.ParseBatch(batch); err != nil {
		return nil, fmt.Errorf("parse batch error:%v", err)
	}
	if err := d.handleL1Message(batchInfo, batchInfo.parentTotalL1MessagePopped, l2Height); err != nil {
		return nil, fmt.Errorf("handle l1 message error:%v", err)
	}
	return batchInfo, nil
}

func (d *Derivation) handleL1Message(rollupData *BatchInfo, parentTotalL1MessagePopped, l2Height uint64) error {
	totalL1MessagePopped := parentTotalL1MessagePopped
	for bIndex, block := range rollupData.blockContexts {
		// This may happen to nodes started from snapshot, in which case we will no longer handle L1Msg
		if block.Number <= l2Height {
			continue
		}
		var l1Transactions []*eth.Transaction
		l1Messages, err := d.getL1Message(totalL1MessagePopped, uint64(block.l1MsgNum))
		if err != nil {
			return fmt.Errorf("get l1 message error:%v", err)
		}
		if len(l1Messages) != int(block.l1MsgNum) {
			return fmt.Errorf("invalid l1 msg num,expect %v,have %v", block.l1MsgNum, l1Messages)
		}
		totalL1MessagePopped += uint64(block.l1MsgNum)
		if len(l1Messages) > 0 {
			for _, l1Message := range l1Messages {
				transaction := eth.NewTx(&l1Message.L1MessageTx)
				l1Transactions = append(l1Transactions, transaction)
			}
		}
		rollupData.blockContexts[bIndex].SafeL2Data.Transactions = append(encodeTransactions(l1Transactions), rollupData.blockContexts[bIndex].SafeL2Data.Transactions...)
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
	for _, blockData := range rollupData.blockContexts {
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

	return lastHeader, nil
}

func (d *Derivation) getLatestConfirmedBlockNumber(ctx context.Context) (uint64, error) {
	return nodecommon.GetLatestConfirmedBlockNumber(ctx, d.l1Client, d.confirmations)
}
