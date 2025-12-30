package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
	"github.com/morph-l2/go-ethereum/rlp"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/l2node"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"morph-l2/bindings/bindings"
	"morph-l2/node/sync"
	"morph-l2/node/types"
)

type NewSyncerFunc func() (*sync.Syncer, error)

type Executor struct {
	l2Client            *types.RetryableClient
	bc                  BlockConverter
	nextL1MsgIndex      uint64
	maxL1MsgNumPerBlock uint64
	l1MsgReader         types.L1MessageReader

	newSyncerFunc NewSyncerFunc
	syncer        *sync.Syncer

	govCaller       *bindings.GovCaller
	sequencerCaller *bindings.SequencerCaller
	l2StakingCaller *bindings.L2StakingCaller

	currentSeqHash *[32]byte
	valsByTmKey    map[[tmKeySize]byte]validatorInfo

	nextValidators [][]byte
	batchParams    tmproto.BatchParams
	tmPubKey       []byte
	isSequencer    bool
	devSequencer   bool

	UpgradeBatchTime      uint64
	blsKeyCheckForkHeight uint64
	rollupABI             *abi.ABI
	batchingCache         *BatchingCache

	logger  tmlog.Logger
	metrics *Metrics
}

func getNextL1MsgIndex(client *types.RetryableClient) (uint64, error) {
	currentHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return currentHeader.NextL1MsgIndex, nil
}

func NewExecutor(newSyncFunc NewSyncerFunc, config *Config, tmPubKey crypto.PubKey) (*Executor, error) {
	logger := config.Logger
	logger = logger.With("module", "executor")
	aClient, err := authclient.DialContext(context.Background(), config.L2.EngineAddr, config.L2.JwtSecret)
	if err != nil {
		return nil, err
	}
	eClient, err := ethclient.Dial(config.L2.EthAddr)
	if err != nil {
		return nil, err
	}

	l2Client := types.NewRetryableClient(aClient, eClient, config.Logger)
	index, err := getNextL1MsgIndex(l2Client)
	if err != nil {
		return nil, err
	}
	logger.Info("obtained next L1Message index when initilize executor", "index", index)

	sequencer, err := bindings.NewSequencerCaller(config.SequencerAddress, l2Client)
	if err != nil {
		return nil, err
	}
	gov, err := bindings.NewGovCaller(config.GovAddress, l2Client)
	if err != nil {
		return nil, err
	}
	l2Staking, err := bindings.NewL2StakingCaller(config.L2StakingAddress, l2Client)
	if err != nil {
		return nil, err
	}

	rollupAbi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	var tmPubKeyBytes []byte
	if tmPubKey != nil {
		tmPubKeyBytes = tmPubKey.Bytes()
	}
	executor := &Executor{
		l2Client:              l2Client,
		bc:                    &Version1Converter{},
		govCaller:             gov,
		sequencerCaller:       sequencer,
		l2StakingCaller:       l2Staking,
		tmPubKey:              tmPubKeyBytes,
		nextL1MsgIndex:        index,
		maxL1MsgNumPerBlock:   config.MaxL1MessageNumPerBlock,
		newSyncerFunc:         newSyncFunc,
		devSequencer:          config.DevSequencer,
		rollupABI:             rollupAbi,
		batchingCache:         NewBatchingCache(),
		UpgradeBatchTime:      config.UpgradeBatchTime,
		blsKeyCheckForkHeight: config.BlsKeyCheckForkHeight,
		logger:                logger,
		metrics:               PrometheusMetrics("morphnode"),
	}

	if config.DevSequencer {
		executor.syncer, err = executor.newSyncerFunc()
		if err != nil {
			return nil, err
		}
		//executor.syncer.Start()
		executor.l1MsgReader = executor.syncer
		return executor, nil
	}

	// Get current height for initial sequencer set update
	currentHeight, err := l2Client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	if _, err = executor.updateSequencerSet(currentHeight); err != nil {
		return nil, err
	}

	return executor, nil
}

var _ l2node.L2Node = (*Executor)(nil)

func (e *Executor) RequestBlockData(height int64) (txs [][]byte, blockMeta []byte, collectedL1Msgs bool, err error) {
	if e.l1MsgReader == nil {
		err = fmt.Errorf("RequestBlockData is not alllowed to be called")
		return
	}
	e.logger.Info("RequestBlockData request", "height", height)
	// read the l1 messages
	fromIndex := e.nextL1MsgIndex
	l1Messages := e.l1MsgReader.ReadL1MessagesInRange(fromIndex, fromIndex+e.maxL1MsgNumPerBlock-1)
	transactions := make(eth.Transactions, len(l1Messages))

	var collectedL1TxHashes []common.Hash
	if len(l1Messages) > 0 {
		queueIndex := fromIndex
		for i, l1Message := range l1Messages {
			transaction := eth.NewTx(&l1Message.L1MessageTx)
			transactions[i] = transaction
			if queueIndex != l1Message.QueueIndex {
				e.logger.Error("unexpected l1message queue index", "expected", queueIndex, "actual", l1Message.QueueIndex)
				err = types.ErrInvalidL1MessageOrder
				return
			}
			collectedL1TxHashes = append(collectedL1TxHashes, l1Message.L1TxHash)
			queueIndex++
		}
		collectedL1Msgs = true
	}
	for _, tx := range transactions {
		l1MsgTx := tx.AsL1MessageTx()
		e.logger.Info("[debug]queueIndex in transactions: ", "queueIndex", l1MsgTx.QueueIndex, "gas", l1MsgTx.Gas, "hash", tx.Hash().String())
	}
	l2Block, err := e.l2Client.AssembleL2Block(context.Background(), big.NewInt(height), transactions)
	if err != nil {
		e.logger.Error("failed to assemble block", "height", height, "error", err)
		return
	}
	e.logger.Info("AssembleL2Block returns l2Block", "tx length", len(l2Block.Transactions))

	wb := types.WrappedBlock{
		ParentHash:          l2Block.ParentHash,
		Miner:               l2Block.Miner,
		Number:              l2Block.Number,
		GasLimit:            l2Block.GasLimit,
		BaseFee:             l2Block.BaseFee,
		Timestamp:           l2Block.Timestamp,
		StateRoot:           l2Block.StateRoot,
		GasUsed:             l2Block.GasUsed,
		ReceiptRoot:         l2Block.ReceiptRoot,
		LogsBloom:           l2Block.LogsBloom,
		WithdrawTrieRoot:    l2Block.WithdrawTrieRoot,
		NextL1MessageIndex:  l2Block.NextL1MessageIndex,
		Hash:                l2Block.Hash,
		CollectedL1TxHashes: collectedL1TxHashes,
	}
	blockMeta, err = wb.MarshalBinary()
	txs = l2Block.Transactions
	e.logger.Info("RequestBlockData response",
		"txs.length", len(txs),
		"collectedL1Msgs", collectedL1Msgs)
	return
}

func (e *Executor) CheckBlockData(txs [][]byte, metaData []byte) (valid bool, err error) {
	if e.l1MsgReader == nil {
		return false, fmt.Errorf("RequestBlockData is not alllowed to be called")
	}
	if len(metaData) == 0 {
		e.logger.Error("metaData cannot be nil")
		return false, nil
	}

	wrappedBlock := new(types.WrappedBlock)
	if err = wrappedBlock.UnmarshalBinary(metaData); err != nil {
		e.logger.Error("failed to UnmarshalBinary meta data bytes", "err", err)
		return false, nil
	}

	e.logger.Info("CheckBlockData requests",
		"txs.length", len(txs),
		"metaData length", len(metaData),
		"eth block number", wrappedBlock.Number)

	l2Block := &catalyst.ExecutableL2Data{
		ParentHash:         wrappedBlock.ParentHash,
		Miner:              wrappedBlock.Miner,
		Number:             wrappedBlock.Number,
		GasLimit:           wrappedBlock.GasLimit,
		BaseFee:            wrappedBlock.BaseFee,
		Timestamp:          wrappedBlock.Timestamp,
		StateRoot:          wrappedBlock.StateRoot,
		GasUsed:            wrappedBlock.GasUsed,
		ReceiptRoot:        wrappedBlock.ReceiptRoot,
		LogsBloom:          wrappedBlock.LogsBloom,
		WithdrawTrieRoot:   wrappedBlock.WithdrawTrieRoot,
		NextL1MessageIndex: wrappedBlock.NextL1MessageIndex,
		Hash:               wrappedBlock.Hash,

		Transactions: txs,
	}
	if err = e.validateL1Messages(l2Block, wrappedBlock.CollectedL1TxHashes); err != nil {
		if !errors.Is(err, types.ErrQueryL1Message) { // hide error if it is not ErrQueryL1Message
			err = nil
		}
		return false, err
	}
	l2Block.WithdrawTrieRoot = wrappedBlock.WithdrawTrieRoot

	validated, err := e.l2Client.ValidateL2Block(context.Background(), l2Block)
	e.logger.Info("CheckBlockData response", "validated", validated, "error", err)
	return validated, err
}

func (e *Executor) DeliverBlock(txs [][]byte, metaData []byte, consensusData l2node.ConsensusData) (nextBatchParams *tmproto.BatchParams, nextValidatorSet [][]byte, err error) {
	e.logger.Info("DeliverBlock request", "txs length", len(txs),
		"blockMeta length", len(metaData),
		"batchHash", hexutil.Encode(consensusData.BatchHash))
	height, err := e.l2Client.BlockNumber(context.Background())
	if err != nil {
		return nil, nil, err
	}
	if metaData == nil {
		e.logger.Error("blockMeta cannot be nil")
		return nil, nil, errors.New("blockMeta cannot be nil")
	}

	wrappedBlock := new(types.WrappedBlock)
	if err = wrappedBlock.UnmarshalBinary(metaData); err != nil {
		e.logger.Error("failed to UnmarshalBinary meta data bytes", "err", err)
		return nil, nil, err
	}

	if wrappedBlock.Number <= height {
		e.logger.Info("ignore it, the block was delivered", "block number", wrappedBlock.Number)
		if e.devSequencer {
			return nil, consensusData.ValidatorSet, nil
		}
		return e.getParamsAndValsAtHeight(int64(wrappedBlock.Number))
	}

	// We only accept the continuous blocks for now.
	// It acts like full sync. Snap sync is not enabled until the Geth enables snapshot with zkTrie
	if wrappedBlock.Number > height+1 {
		return nil, nil, types.ErrWrongBlockNumber
	}

	if len(consensusData.BatchHash) > 0 {
		e.metrics.BatchPointHeight.Set(float64(wrappedBlock.Number))
	}

	l2Block := &catalyst.ExecutableL2Data{
		ParentHash:         wrappedBlock.ParentHash,
		Miner:              wrappedBlock.Miner,
		Number:             wrappedBlock.Number,
		GasLimit:           wrappedBlock.GasLimit,
		BaseFee:            wrappedBlock.BaseFee,
		Timestamp:          wrappedBlock.Timestamp,
		StateRoot:          wrappedBlock.StateRoot,
		GasUsed:            wrappedBlock.GasUsed,
		ReceiptRoot:        wrappedBlock.ReceiptRoot,
		LogsBloom:          wrappedBlock.LogsBloom,
		WithdrawTrieRoot:   wrappedBlock.WithdrawTrieRoot,
		NextL1MessageIndex: wrappedBlock.NextL1MessageIndex,
		Hash:               wrappedBlock.Hash,

		Transactions: txs,
	}
	var batchHash *common.Hash
	if consensusData.BatchHash != nil {
		batchHash = new(common.Hash)
		copy(batchHash[:], consensusData.BatchHash)
	}
	err = e.l2Client.NewL2Block(context.Background(), l2Block, batchHash)
	if err != nil {
		e.logger.Error("failed to NewL2Block", "error", err)
		return nil, nil, err
	}

	// end block
	e.updateNextL1MessageIndex(l2Block)

	var newValidatorSet = consensusData.ValidatorSet
	var newBatchParams *tmproto.BatchParams
	if !e.devSequencer {
		if newValidatorSet, err = e.updateSequencerSet(l2Block.Number); err != nil {
			return nil, nil, err
		}
		if newBatchParams, err = e.batchParamsUpdates(l2Block.Number); err != nil {
			return nil, nil, err
		}
	}

	e.metrics.Height.Set(float64(l2Block.Number))

	return newBatchParams, newValidatorSet,
		nil
}

// EncodeTxs
// decode each transaction bytes into Transaction, and wrap them into an array, then rlpEncode the whole array
func (e *Executor) EncodeTxs(batchTxs [][]byte) ([]byte, error) {
	if len(batchTxs) == 0 {
		return []byte{}, nil
	}
	transactions := make([]*eth.Transaction, len(batchTxs))
	for i, txBz := range batchTxs {
		if len(txBz) == 0 {
			return nil, fmt.Errorf("transaction %d is empty", i)
		}
		var tx eth.Transaction
		if err := tx.UnmarshalBinary(txBz); err != nil {
			return nil, fmt.Errorf("transaction %d is not valid: %v", i, err)
		}
		transactions[i] = &tx
	}
	return rlp.EncodeToBytes(transactions)
}

func (e *Executor) RequestHeight(tmHeight int64) (int64, error) {
	curHeight, err := e.l2Client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return int64(curHeight), nil
}

func (e *Executor) getParamsAndValsAtHeight(height int64) (*tmproto.BatchParams, [][]byte, error) {
	callOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(height),
	}
	batchBlockInterval, err := e.govCaller.BatchBlockInterval(callOpts)
	if err != nil {
		return nil, nil, err
	}
	batchTimeout, err := e.govCaller.BatchTimeout(callOpts)
	if err != nil {
		return nil, nil, err
	}
	// fetch current sequencerSet info at certain height
	addrs, err := e.sequencerCaller.GetSequencerSet2(callOpts)
	if err != nil {
		return nil, nil, err
	}
	stakesInfo, err := e.l2StakingCaller.GetStakesInfo(callOpts, addrs)
	if err != nil {
		return nil, nil, err
	}
	newValidators := make([][]byte, 0, len(addrs))
	for i := range stakesInfo {
		// validate blsKey to keep consistent with sequencerSetUpdates
		if _, err := decodeBlsPubKey(stakesInfo[i].BlsKey); err != nil {
			e.logger.Error("getParamsAndValsAtHeight: failed to decode bls key", "key bytes", hexutil.Encode(stakesInfo[i].BlsKey), "error", err)
			if e.isBlsKeyCheckFork(uint64(height)) {
				continue
			}
		}
		newValidators = append(newValidators, stakesInfo[i].TmKey[:])
	}

	return &tmproto.BatchParams{
		BlocksInterval: batchBlockInterval.Int64(),
		Timeout:        time.Duration(batchTimeout.Int64() * int64(time.Second)),
	}, newValidators, nil
}

func (e *Executor) L2Client() *types.RetryableClient {
	return e.l2Client
}

// ============================================================================
// L2NodeV2 interface implementation for sequencer mode
// ============================================================================

// RequestBlockDataV2 requests block data based on parent hash.
// This differs from RequestBlockData which uses height.
// Using parent hash allows for proper fork chain handling in the future.
func (e *Executor) RequestBlockDataV2(parentHashBytes []byte) (*l2node.BlockV2, bool, error) {
	if e.l1MsgReader == nil {
		return nil, false, fmt.Errorf("RequestBlockDataV2 is not allowed to be called")
	}
	parentHash := common.BytesToHash(parentHashBytes)

	// Read L1 messages
	fromIndex := e.nextL1MsgIndex
	l1Messages := e.l1MsgReader.ReadL1MessagesInRange(fromIndex, fromIndex+e.maxL1MsgNumPerBlock-1)
	transactions := make(eth.Transactions, len(l1Messages))

	collectedL1Msgs := false
	if len(l1Messages) > 0 {
		queueIndex := fromIndex
		for i, l1Message := range l1Messages {
			transaction := eth.NewTx(&l1Message.L1MessageTx)
			transactions[i] = transaction
			if queueIndex != l1Message.QueueIndex {
				e.logger.Error("unexpected l1message queue index", "expected", queueIndex, "actual", l1Message.QueueIndex)
				return nil, false, types.ErrInvalidL1MessageOrder
			}
			queueIndex++
		}
		collectedL1Msgs = true
	}

	// Call geth to assemble block based on parent hash
	l2Block, err := e.l2Client.AssembleL2BlockV2(context.Background(), parentHash, transactions)
	if err != nil {
		e.logger.Error("failed to assemble block v2", "parentHash", parentHash.Hex(), "error", err)
		return nil, false, err
	}

	e.logger.Info("AssembleL2BlockV2 success ",
		"number", l2Block.Number,
		"hash", l2Block.Hash.Hex(),
		"parentHash", l2Block.ParentHash.Hex(),
		"tx length", len(l2Block.Transactions),
		"collectedL1Msgs", collectedL1Msgs,
	)

	return executableL2DataToBlockV2(l2Block), collectedL1Msgs, nil
}

// ApplyBlockV2 applies a block to the L2 execution layer.
// This is used in sequencer mode after block validation.
func (e *Executor) ApplyBlockV2(block *l2node.BlockV2) error {
	// Convert BlockV2 to ExecutableL2Data for geth
	execBlock := blockV2ToExecutableL2Data(block)

	// Check if block is already applied
	height, err := e.l2Client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	if execBlock.Number <= height {
		e.logger.Info("ignore it, the block was already applied", "block number", execBlock.Number)
		return nil
	}

	// We only accept continuous blocks
	if execBlock.Number > height+1 {
		return types.ErrWrongBlockNumber
	}

	// Apply the block (no batch hash in sequencer mode for now)
	err = e.l2Client.NewL2Block(context.Background(), execBlock, nil)
	if err != nil {
		e.logger.Error("failed to apply block v2", "error", err)
		return err
	}

	// Update L1 message index
	e.updateNextL1MessageIndex(execBlock)

	e.metrics.Height.Set(float64(execBlock.Number))
	e.logger.Info("ApplyBlockV2 success", "number", execBlock.Number, "hash", execBlock.Hash.Hex())

	return nil
}

// GetBlockByNumber retrieves a block by its number from the L2 execution layer.
// Uses standard eth_getBlockByNumber JSON-RPC.
func (e *Executor) GetBlockByNumber(height uint64) (*l2node.BlockV2, error) {
	block, err := e.l2Client.BlockByNumber(context.Background(), big.NewInt(int64(height)))
	if err != nil {
		e.logger.Error("failed to get block by number", "height", height, "error", err)
		return nil, err
	}
	return ethBlockToBlockV2(block)
}

// GetLatestBlockV2 returns the latest block from the L2 execution layer.
// Uses standard eth_getBlockByNumber JSON-RPC with nil (latest).
func (e *Executor) GetLatestBlockV2() (*l2node.BlockV2, error) {
	block, err := e.l2Client.BlockByNumber(context.Background(), nil)
	if err != nil {
		e.logger.Error("failed to get latest block", "error", err)
		return nil, err
	}
	return ethBlockToBlockV2(block)
}

// ==================== Type Conversion Functions ====================

// ethBlockToBlockV2 converts eth.Block to BlockV2
func ethBlockToBlockV2(block *eth.Block) (*l2node.BlockV2, error) {
	if block == nil {
		return nil, fmt.Errorf("block is nil")
	}
	header := block.Header()

	// Encode transactions using MarshalBinary (handles EIP-2718 typed transactions correctly)
	// Initialize as empty slice (not nil) to ensure JSON serialization produces [] instead of null
	txs := make([][]byte, 0, len(block.Transactions()))
	for _, tx := range block.Transactions() {
		bz, err := tx.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal tx, error: %v", err)
		}
		txs = append(txs, bz)
	}

	return &l2node.BlockV2{
		ParentHash:         header.ParentHash,
		Miner:              header.Coinbase,
		Number:             header.Number.Uint64(),
		GasLimit:           header.GasLimit,
		BaseFee:            header.BaseFee,
		Timestamp:          header.Time,
		Transactions:       txs,
		StateRoot:          header.Root,
		GasUsed:            header.GasUsed,
		ReceiptRoot:        header.ReceiptHash,
		LogsBloom:          header.Bloom.Bytes(),
		NextL1MessageIndex: header.NextL1MsgIndex,
		Hash:               block.Hash(),
	}, nil
}

// blockV2ToExecutableL2Data converts BlockV2 to ExecutableL2Data
func blockV2ToExecutableL2Data(block *l2node.BlockV2) *catalyst.ExecutableL2Data {
	if block == nil {
		return nil
	}
	// Ensure Transactions is not nil (JSON requires [] not null)
	txs := block.Transactions
	if txs == nil {
		txs = make([][]byte, 0)
	}
	return &catalyst.ExecutableL2Data{
		ParentHash:         block.ParentHash,
		Miner:              block.Miner,
		Number:             block.Number,
		GasLimit:           block.GasLimit,
		BaseFee:            block.BaseFee,
		Timestamp:          block.Timestamp,
		Transactions:       txs,
		StateRoot:          block.StateRoot,
		GasUsed:            block.GasUsed,
		ReceiptRoot:        block.ReceiptRoot,
		LogsBloom:          block.LogsBloom,
		WithdrawTrieRoot:   block.WithdrawTrieRoot,
		NextL1MessageIndex: block.NextL1MessageIndex,
		Hash:               block.Hash,
	}
}

// executableL2DataToBlockV2 converts ExecutableL2Data to BlockV2
func executableL2DataToBlockV2(data *catalyst.ExecutableL2Data) *l2node.BlockV2 {
	if data == nil {
		return nil
	}
	return &l2node.BlockV2{
		ParentHash:         data.ParentHash,
		Miner:              data.Miner,
		Number:             data.Number,
		GasLimit:           data.GasLimit,
		BaseFee:            data.BaseFee,
		Timestamp:          data.Timestamp,
		Transactions:       data.Transactions,
		StateRoot:          data.StateRoot,
		GasUsed:            data.GasUsed,
		ReceiptRoot:        data.ReceiptRoot,
		LogsBloom:          data.LogsBloom,
		WithdrawTrieRoot:   data.WithdrawTrieRoot,
		NextL1MessageIndex: data.NextL1MessageIndex,
		Hash:               data.Hash,
	}
}
