package derivation

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/morph-l2/bindings/bindings"
	node "github.com/morph-l2/node/core"
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
	"github.com/scroll-tech/go-ethereum/eth/catalyst"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/ethclient/authclient"
	"github.com/scroll-tech/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

var (
	RollupEventTopic     = "CommitBatch(uint256,bytes32)"
	RollupEventTopicHash = crypto.Keccak256Hash([]byte(RollupEventTopic))
)

// BatchInfo is all rollup data of one l1 block,maybe contain many rollup batch
type BatchInfo struct {
	batchIndex       uint64
	blockNum         uint64
	txNum            uint64
	version          uint64
	dataHash         common.Hash
	batchHash        common.Hash
	chunks           []*Chunk
	l1BlockNumber    uint64
	txHash           common.Hash
	nonce            uint64
	lastBlockNumber  uint64
	firstBlockNumber uint64

	root                   common.Hash
	skippedL1MessageBitmap *big.Int
}

func (bi *BatchInfo) FirstBlockNumber() uint64 {
	return bi.firstBlockNumber
}

func (bi *BatchInfo) LastBlockNumber() uint64 {
	return bi.lastBlockNumber
}

func (bi *BatchInfo) BlockNum() uint64 {
	return bi.blockNum
}

func (bi *BatchInfo) TxNum() uint64 {
	return bi.txNum
}

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

	d.logger.Info("Stopping Derivation service")

	if d.cancel != nil {
		d.cancel()
	}
	<-d.stop
	d.logger.Info("Derivation service is stopped")
}

func (d *Derivation) derivationBlock(ctx context.Context) {
	latestDerivation := d.db.ReadLatestDerivationL1Height()
	latest := d.syncer.LatestSynced()
	start := *latestDerivation + 1
	end := latest
	if latest < start {
		d.logger.Info("latest less than or equal to start", "latest", latest, "start", start)
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
	// parse latest batch
	d.logger.Info(fmt.Sprintf("rollup latest batch index:%v", latestBatchIndex))
	d.logger.Info("fetched rollup tx", "txNum", len(logs))

	for _, lg := range logs {
		batchInfo, err := d.fetchRollupDataByTxHash(lg.TxHash, lg.BlockNumber)
		if err != nil {
			rollupCommitBatch, parseErr := d.rollup.ParseCommitBatch(lg)
			if parseErr != nil {
				d.logger.Error("get l2 BlockNumber", "err", err)
				return
			}
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
		d.logger.Info("batch derivation complete", "currentBatchEndBlock", lastHeader.Number.Uint64())
		d.metrics.SetL2DeriveHeight(lastHeader.Number.Uint64())
		if !bytes.Equal(lastHeader.Root.Bytes(), batchInfo.root.Bytes()) && d.validator != nil && d.validator.ChallengeEnable() {
			d.logger.Info("root hash is not equal", "originStateRootHash", batchInfo.root, "deriveStateRootHash", lastHeader.Root.Hex())
			return
		}
		//}
		d.db.WriteLatestDerivationL1Height(lg.BlockNumber)
		d.metrics.SetL1SyncHeight(lg.BlockNumber)
		d.logger.Info("WriteLatestDerivationL1Height success", "l1BlockNumber", lg.BlockNumber)
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
	}
	rollupData, err := d.parseBatch(batch)
	if err != nil {
		d.logger.Error("ParseBatch failed", "txNonce", tx.Nonce(), "txHash", txHash,
			"l1BlockNumber", blockNumber)
		return rollupData, fmt.Errorf("ParseBatch error:%v\n", err)
	}
	rollupData.l1BlockNumber = blockNumber
	rollupData.txHash = txHash
	rollupData.nonce = tx.Nonce()
	return rollupData, nil
}

type Chunk struct {
	blockContext []*BlockContext
	txsPayload   [][]*eth.Transaction
	txHashes     [][]common.Hash
	blockNum     int
}

type BlockContext struct {
	Number    uint64 `json:"number"`
	Timestamp uint64 `json:"timestamp"`
	BaseFee   *big.Int
	GasLimit  uint64
	txsNum    uint16
	l1MsgNum  uint16

	SafeL2Data *catalyst.SafeL2Data
}

func (b *BlockContext) Decode(bc []byte) error {
	reader := bytes.NewReader(bc)
	bsBaseFee := make([]byte, 32)
	if err := binary.Read(reader, binary.BigEndian, &b.Number); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &b.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &bsBaseFee); err != nil {
		return err
	}
	b.BaseFee = new(big.Int).SetBytes(bsBaseFee)
	if err := binary.Read(reader, binary.BigEndian, &b.GasLimit); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &b.txsNum); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &b.l1MsgNum); err != nil {
		return err
	}
	return nil
}

func parseChunk(chunkBytes []byte) (*types.Chunk, error) {
	reader := bytes.NewReader(chunkBytes)
	var blockNum uint8
	if err := binary.Read(reader, binary.BigEndian, &blockNum); err != nil {
		return nil, err
	}

	blockCtx := make([]byte, 0)
	for i := 0; i < int(blockNum); i++ {
		bc := make([]byte, 60)
		if err := binary.Read(reader, binary.BigEndian, &bc); err != nil {
			return nil, err
		}
		blockCtx = append(blockCtx, bc...)
	}
	txsPayload := make([]byte, len(chunkBytes)-int(blockNum)*60-1)
	if err := binary.Read(reader, binary.BigEndian, &txsPayload); err != nil {
		return nil, err
	}
	chunk := types.NewChunk(blockCtx, txsPayload, nil, nil)
	chunk.ResetBlockNum(int(blockNum))
	return chunk, nil
}

func (d *Derivation) parseBatch(batch geth.RPCRollupBatch) (*BatchInfo, error) {
	parentBatchHeader, err := types.DecodeBatchHeader(batch.ParentBatchHeader)
	if err != nil {
		return nil, fmt.Errorf("DecodeBatchHeader error:%v", err)
	}
	rollupData, err := ParseBatch(batch)
	if err != nil {
		return nil, fmt.Errorf("parse batch error:%v", err)
	}
	if err := d.handleL1Message(rollupData, &parentBatchHeader); err != nil {
		return nil, fmt.Errorf("handleL1Message error:%v", err)
	}
	rollupData.batchIndex = parentBatchHeader.BatchIndex + 1
	return rollupData, nil
}

func ParseBatch(batch geth.RPCRollupBatch) (*BatchInfo, error) {
	var rollupData BatchInfo
	rollupData.root = batch.PostStateRoot
	rollupData.skippedL1MessageBitmap = new(big.Int).SetBytes(batch.SkippedL1MessageBitmap[:])
	rollupData.version = uint64(batch.Version)
	chunks := types.NewChunks()
	for cbIndex, chunkByte := range batch.Chunks {
		chunk, err := parseChunk(chunkByte)
		if err != nil {
			return nil, fmt.Errorf("parse chunk error:%v", err)
		}
		rollupData.blockNum += uint64(chunk.BlockNum())
		chunks.Append(chunk.BlockContext(), chunk.TxsPayload(), nil, nil)
		ck := Chunk{}
		var txsNum uint64
		var l1MsgNum uint64
		reader := bytes.NewReader(chunk.TxsPayload())
		for i := 0; i < chunk.BlockNum(); i++ {
			var block BlockContext
			err = block.Decode(chunk.BlockContext()[i*60 : i*60+60])
			if err != nil {
				return nil, fmt.Errorf("decode chunk block context error:%v", err)
			}
			if cbIndex == 0 && i == 0 {
				rollupData.firstBlockNumber = block.Number
			}
			if cbIndex == len(batch.Chunks)-1 && i == chunk.BlockNum()-1 {
				rollupData.lastBlockNumber = block.Number
			}
			var safeL2Data catalyst.SafeL2Data
			safeL2Data.Number = block.Number
			safeL2Data.GasLimit = block.GasLimit
			safeL2Data.BaseFee = block.BaseFee
			safeL2Data.Timestamp = block.Timestamp
			if block.BaseFee != nil && block.BaseFee.Cmp(big.NewInt(0)) == 0 {
				safeL2Data.BaseFee = nil
			}
			if block.txsNum < block.l1MsgNum {
				return nil, fmt.Errorf("txsNum must be or equal to or greater than l1MsgNum,txsNum:%v,l1MsgNum:%v", block.txsNum, block.l1MsgNum)
			}

			txs, err := node.DecodeTxsPayload(reader, int(block.txsNum)-int(block.l1MsgNum))
			if err != nil {
				return nil, fmt.Errorf("DecodeTxsPayload error:%v", err)
			}
			txsNum += uint64(block.txsNum)
			l1MsgNum += uint64(block.l1MsgNum)
			safeL2Data.Transactions = encodeTransactions(txs)
			if block.txsNum > 0 {
				safeL2Data.Transactions = encodeTransactions(txs)
			} else {
				safeL2Data.Transactions = [][]byte{}
			}
			block.SafeL2Data = &safeL2Data
			ck.blockContext = append(ck.blockContext, &block)
		}
		rollupData.txNum += txsNum
		rollupData.chunks = append(rollupData.chunks, &ck)
	}
	rollupData.dataHash = chunks.DataHash()
	return &rollupData, nil
}

func (d *Derivation) handleL1Message(rollupData *BatchInfo, parentBatchHeader *types.BatchHeader) error {
	batchHeader := types.BatchHeader{
		Version:                uint8(rollupData.version),
		BatchIndex:             parentBatchHeader.BatchIndex + 1,
		DataHash:               rollupData.dataHash,
		ParentBatchHash:        parentBatchHeader.ParentBatchHash,
		SkippedL1MessageBitmap: rollupData.skippedL1MessageBitmap.Bytes(),
	}
	var l1MessagePopped, totalL1MessagePopped uint64
	totalL1MessagePopped = parentBatchHeader.TotalL1MessagePopped
	for index, chunk := range rollupData.chunks {
		for bIndex, block := range chunk.blockContext {
			var l1Transactions []*eth.Transaction
			l1Messages, err := d.getL1Message(totalL1MessagePopped, uint64(block.l1MsgNum))
			if err != nil {
				return fmt.Errorf("getL1Message error:%v", err)
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
			rollupData.chunks[index].blockContext[bIndex].SafeL2Data.Transactions = append(encodeTransactions(l1Transactions), chunk.blockContext[bIndex].SafeL2Data.Transactions...)
		}

	}
	batchHeader.TotalL1MessagePopped = totalL1MessagePopped
	batchHeader.L1MessagePopped = l1MessagePopped
	batchHeader.Encode()
	rollupData.batchHash = batchHeader.Hash()
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
		for _, blockData := range chunk.blockContext {
			blockData.SafeL2Data.BatchHash = &rollupData.batchHash
			latestBlockNumber, err := d.l2Client.BlockNumber(context.Background())
			if err != nil {
				return nil, fmt.Errorf("get derivation geth block number error:%v", err)
			}
			if blockData.SafeL2Data.Number <= latestBlockNumber {
				d.logger.Info("SafeL2Data block number less than latestBlockNumber", "safeL2DataNumber", blockData.SafeL2Data.Number, "latestBlockNumber", latestBlockNumber)
				lastHeader, err = d.l2Client.HeaderByNumber(d.ctx, big.NewInt(int64(latestBlockNumber)))
				continue
			}
			err = func() error {
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
				defer cancel()
				lastHeader, err = d.l2Client.NewSafeL2Block(ctx, blockData.SafeL2Data)
				if err != nil {
					d.logger.Error("NewL2Block failed", "latestBlockNumber", latestBlockNumber, "error", err)
					return err
				}
				return nil
			}()
			if err != nil {
				return nil, fmt.Errorf("derivation error:%v", err)
			}
			d.logger.Info("NewSafeL2Block success", "blockNumber", blockData.Number)
		}
	}

	return lastHeader, nil
}

func (d *Derivation) findBatchIndex(txHash common.Hash, blockNumber uint64) (uint64, error) {
	receipt, err := d.l1Client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return 0, err
	}
	if receipt.Status == eth.ReceiptStatusFailed {
		return 0, err
	}

	return 0, fmt.Errorf("event not found")
}

func encodeTransactions(txs []*eth.Transaction) [][]byte {
	var enc = make([][]byte, len(txs))
	for i, tx := range txs {
		enc[i], _ = tx.MarshalBinary()
	}
	return enc
}
