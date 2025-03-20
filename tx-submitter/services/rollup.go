package services

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/consensus/misc/eip4844"
	"github.com/morph-l2/go-ethereum/core"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/bls12381"
	"github.com/morph-l2/go-ethereum/crypto/kzg4844"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"
	"github.com/morph-l2/go-ethereum/params"
	"github.com/morph-l2/go-ethereum/rpc"
	"github.com/tendermint/tendermint/blssignatures"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/constants"
	"morph-l2/tx-submitter/db"
	"morph-l2/tx-submitter/event"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/l1checker"
	"morph-l2/tx-submitter/localpool"
	"morph-l2/tx-submitter/metrics"
	"morph-l2/tx-submitter/types"
	"morph-l2/tx-submitter/utils"
)

const (
	txSlotSize           = 32 * 1024
	txMaxSize            = 4 * txSlotSize // 128KB
	rotatorWait          = 3 * time.Second
	rollupSumKey         = "rollup_sum"
	finalizeSumKey       = "finalize_sum"
	collectedL1FeeSumKey = "collected_l1_fee_sum"
)

type Rollup struct {
	ctx         context.Context
	metrics     *metrics.Metrics
	l1RpcClient *rpc.Client
	L1Client    iface.Client
	L2Clients   []iface.L2Client
	Rollup      iface.IRollup
	Staking     iface.IL1Staking
	chainId     *big.Int
	privKey     *ecdsa.PrivateKey
	rollupAddr  common.Address
	abi         *abi.ABI
	// rotator
	rotator          *Rotator
	pendingTxs       *PendingTxs
	rollupFinalizeMu sync.Mutex
	externalRsaPriv  *rsa.PrivateKey
	// cfg
	cfg utils.Config
	// signer
	signer ethtypes.Signer
	// leveldb
	ldb *db.Db
	// rollupFeeSum
	rollupFeeSum float64
	// finalizeFeeSum
	finalizeFeeSum float64
	// collectedL1FeeSum
	collectedL1FeeSum float64
	// batchcache
	batchCache       *types.BatchCache
	bm               *l1checker.BlockMonitor
	eventInfoStorage *event.EventInfoStorage
	reorgDetector    iface.IReorgDetector
	// Cache for current submitter info
	submitterCache struct {
		submitter      common.Address
		submitterIdx   uint64
		lastUpdateTime time.Time
		mu             sync.RWMutex
	}
	// Cache refresh interval
	submitterCacheTimeout time.Duration
}

func NewRollup(
	ctx context.Context,
	metrics *metrics.Metrics,
	l1RpcClient *rpc.Client,
	l1 iface.Client,
	l2Clients []iface.L2Client,
	rollup iface.IRollup,
	staking iface.IL1Staking,
	chainId *big.Int,
	priKey *ecdsa.PrivateKey,
	rollupAddr common.Address,
	abi *abi.ABI,
	cfg utils.Config,
	rsaPriv *rsa.PrivateKey,
	rotator *Rotator,
	ldb *db.Db,
	bm *l1checker.BlockMonitor,
	eventInfoStorage *event.EventInfoStorage,
) *Rollup {
	batchFetcher := NewBatchFetcher(l2Clients)
	return &Rollup{
		ctx:              ctx,
		metrics:          metrics,
		l1RpcClient:      l1RpcClient,
		L1Client:         l1,
		Rollup:           rollup,
		Staking:          staking,
		L2Clients:        l2Clients,
		privKey:          priKey,
		chainId:          chainId,
		rollupAddr:       rollupAddr,
		abi:              abi,
		rotator:          rotator,
		cfg:              cfg,
		signer:           ethtypes.LatestSignerForChainID(chainId),
		externalRsaPriv:  rsaPriv,
		batchCache:       types.NewBatchCache(batchFetcher),
		ldb:              ldb,
		bm:               bm,
		eventInfoStorage: eventInfoStorage,
	}
}

func (r *Rollup) Start() error {

	// init rollup service
	if err := r.PreCheck(); err != nil {
		return err
	}

	// journal
	jn := localpool.New(r.cfg.JournalFilePath)
	err := jn.Init()
	if err != nil {
		log.Crit("journal file init failed", "err", err)
	}
	// pendingtxs
	r.pendingTxs = NewPendingTxs(r.abi.Methods[constants.MethodCommitBatch].ID, r.abi.Methods[constants.MethodFinalizeBatch].ID, jn)
	txs, err := jn.ParseAllTxs()
	if err != nil {
		log.Crit("parse l1 mempool error", "error", err)
	}
	if err := r.pendingTxs.Recover(txs, r.abi); err != nil {
		log.Crit("failed to recover pending transactions", "error", err)
	}

	// init fee metrics sum
	err = r.InitFeeMetricsSum()
	if err != nil {
		return fmt.Errorf("init fee metrics sum failed: %w", err)
	}

	/// start services
	// start l1 monitor
	go r.bm.StartMonitoring()

	// metrics
	go utils.Loop(r.ctx, 10*time.Second, func() {

		// get balacnce of wallet
		balance, err := r.L1Client.BalanceAt(context.Background(), r.WalletAddr(), nil)
		if err != nil {
			log.Error("get wallet balance error", "error", err)
			if utils.IsRpcErr(err) {
				r.metrics.IncRpcErrors()
			}
			return
		}
		// balance to eth
		balanceEth := new(big.Rat).SetFrac(balance, big.NewInt(params.Ether))

		// parse float64 from string
		balanceEthFloat, err := strconv.ParseFloat(balanceEth.FloatString(18), 64)
		if err != nil {
			log.Warn("parse balance to float error", "error", err)
			return
		}

		r.metrics.SetWalletBalance(balanceEthFloat)

	})

	go utils.Loop(r.ctx, r.cfg.RollupInterval, func() {
		r.rollupFinalizeMu.Lock()
		defer r.rollupFinalizeMu.Unlock()
		if err := r.rollup(); err != nil {
			if utils.IsRpcErr(err) {
				r.metrics.IncRpcErrors()
			}
			log.Error("rollup failed,wait for the next try", "error", err)
		}
	})

	if r.cfg.Finalize {

		go utils.Loop(r.ctx, r.cfg.FinalizeInterval, func() {
			r.rollupFinalizeMu.Lock()
			defer r.rollupFinalizeMu.Unlock()

			if err := r.finalize(); err != nil {
				log.Error("finalize failed", "error", err)
				if utils.IsRpcErr(err) {
					r.metrics.IncRpcErrors()
				}

			}
		})
	}

	var processtxMu sync.Mutex
	go utils.Loop(r.ctx, r.cfg.TxProcessInterval, func() {
		processtxMu.Lock()
		defer processtxMu.Unlock()
		if err := r.ProcessTx(); err != nil {
			log.Error("process tx err", "error", err)
			if utils.IsRpcErr(err) {
				r.metrics.IncRpcErrors()
			}
		}
	})
	return nil
}

func (r *Rollup) ProcessTx() error {
	// Check for reorgs first with exponential backoff retry
	_, _, err := r.detectReorgWithRetry()
	if err != nil {
		log.Warn("Failed to check for reorgs", "error", err)
	}

	// Get all local transactions
	txRecords := r.pendingTxs.GetAll()
	if len(txRecords) == 0 {
		return nil
	}

	// Check if this submitter should process transactions
	if err := r.checkSubmitterTurn(); err != nil {
		if err == errNotMyTurn {
			// Get current submitter index for logging
			activeSubmitter, activeIndex, _ := r.rotator.CurrentSubmitter(r.L2Clients, r.Staking)

			// Calculate rotation timing information
			past := (time.Now().Unix() - r.rotator.startTime.Int64()) % r.rotator.epoch.Int64()
			start := time.Now().Unix() - past
			end := start + r.rotator.epoch.Int64()
			timeLeft := end - time.Now().Unix()

			// Format timestamps for human readability
			endTimeFormatted := utils.FormatTime(big.NewInt(end))
			timeLeftFormatted := fmt.Sprintf("%dm%ds", timeLeft/60, timeLeft%60)

			log.Info("Awaiting turn for transaction processing",
				"current_submitter", activeSubmitter.Hex(),
				"submitter_index", activeIndex,
				"total_submitters", len(r.rotator.GetSubmitterSet()),
				"next_rotation", endTimeFormatted,
				"time_remaining", timeLeftFormatted)
			return nil
		}
		return err
	}

	// Process each transaction
	for _, txRecord := range txRecords {
		if err := r.processSingleTx(txRecord); err != nil {
			log.Error("Transaction processing failed",
				"tx_hash", txRecord.Tx.Hash().String(),
				"error", err)
			return fmt.Errorf("transaction processing failed: %w", err)
		}
	}

	return nil
}

// Helper function to detect reorgs with retry
func (r *Rollup) detectReorgWithRetry() (bool, uint64, error) {
	var lastErr error
	for i := range 3 { // Try up to 3 times
		hasReorg, depth, err := r.reorgDetector.DetectReorg(r.ctx)
		if err == nil {
			return hasReorg, depth, nil
		}
		lastErr = err
		time.Sleep(time.Duration(i+1) * time.Second) // Exponential backoff
	}
	return false, 0, lastErr
}

var errNotMyTurn = errors.New("not my turn")

func (r *Rollup) checkSubmitterTurn() error {
	if r.cfg.PriorityRollup {
		return nil
	}
	activeSubmitter, submitterIndex, err := r.getCachedSubmitter()
	if err != nil {
		return fmt.Errorf("rollup: get current submitter err, %w", err)
	}

	myAddress := r.WalletAddr().Hex()
	activeAddress := activeSubmitter.Hex()
	isMyTurn := activeAddress == myAddress

	// Calculate rotation timing information
	past := (time.Now().Unix() - r.rotator.startTime.Int64()) % r.rotator.epoch.Int64()
	start := time.Now().Unix() - past
	end := start + r.rotator.epoch.Int64()
	timeLeft := end - time.Now().Unix()

	// Format timestamps for human readability
	startTimeFormatted := utils.FormatTime(big.NewInt(start))
	endTimeFormatted := utils.FormatTime(big.NewInt(end))
	timeLeftFormatted := fmt.Sprintf("%dm%ds", timeLeft/60, timeLeft%60)

	if !isMyTurn {
		log.Debug("Not active submitter",
			"active_submitter", activeAddress,
			"index", submitterIndex,
			"my_address", myAddress,
			"total_submitters", len(r.rotator.GetSubmitterSet()),
			"rotation_end", endTimeFormatted,
			"time_remaining", timeLeftFormatted)
		return errNotMyTurn
	}

	log.Info("Active submitter status",
		"index", submitterIndex,
		"total_submitters", len(r.rotator.GetSubmitterSet()),
		"rotation_start", startTimeFormatted,
		"rotation_end", endTimeFormatted,
		"time_remaining", timeLeftFormatted)
	return nil
}

// Handle chain reorganization
func (r *Rollup) handleReorg(depth uint64) error {
	// Update metrics
	r.metrics.SetReorgDepth(float64(depth))
	r.metrics.IncReorgs()
	return nil
}

// Process a single transaction
func (r *Rollup) processSingleTx(txRecord *types.TxRecord) error {
	rtx := txRecord.Tx
	method := utils.ParseMethod(rtx, r.abi)

	log.Info("Processing transaction",
		"hash", rtx.Hash().String(),
		"nonce", rtx.Nonce(),
		"method", method,
		"query_times", txRecord.QueryTimes)

	// Check transaction status
	status, err := r.getTxStatus(rtx)
	if err != nil {
		return fmt.Errorf("get tx status error: %w", err)
	}

	switch status.state {
	case txStatusPending:
		return r.handlePendingTx(txRecord, rtx, method)
	case txStatusConfirmed:
		// Get current block number
		currentBlock, err := r.L1Client.BlockNumber(context.Background())
		if err != nil {
			return fmt.Errorf("get current block number error: %v", err)
		}

		// Check confirmation depth
		if status.receipt != nil && currentBlock >= status.receipt.BlockNumber.Uint64()+6 {
			// Update fee metrics before removing the transaction
			if err := r.updateFeeMetrics(rtx, status.receipt, method); err != nil {
				log.Error("Failed to update fee metrics",
					"error", err,
					"tx_hash", rtx.Hash().String())
			}

			// Transaction has 6 confirmations, remove it from tracking
			log.Info("Removing confirmed tx from tracking after 6 blocks",
				"tx_hash", rtx.Hash().String(),
				"block_number", status.receipt.BlockNumber.Uint64(),
				"current_block", currentBlock,
				"gas_used", status.receipt.GasUsed,
				"effective_gas_price", status.receipt.EffectiveGasPrice,
				"method", method)
			if err := r.pendingTxs.Remove(rtx.Hash()); err != nil {
				log.Error("failed to remove transaction", "hash", rtx.Hash().String(), "error", err)
			}
			r.metrics.IncTxConfirmed(method)
			return nil
		}
		return r.handleConfirmedTx(txRecord, rtx, method)
	case txStatusMissing:
		return r.handleMissingTx(txRecord, rtx, method)
	default:
		return fmt.Errorf("unknown transaction status: %v", status.state)
	}
}

// updateFeeMetrics updates all fee-related metrics for a confirmed transaction
func (r *Rollup) updateFeeMetrics(tx *ethtypes.Transaction, receipt *ethtypes.Receipt, method string) error {
	txFeeEth := calcFee(tx, receipt)
	txFeeFloat, _ := txFeeEth.Float64()

	// Update metrics based on transaction type
	if method == constants.MethodCommitBatch {
		r.rollupFeeSum += txFeeFloat
		r.metrics.RollupCostSum.Add(txFeeFloat)
		r.metrics.RollupCost.Set(txFeeFloat)
		// Update leveldb
		err := r.ldb.PutFloat(rollupSumKey, r.rollupFeeSum)
		if err != nil {
			return fmt.Errorf("failed to update rollup fee sum in leveldb: %w", err)
		}

		// Calculate and update L1 fee metrics
		batchIndex := utils.ParseParentBatchIndex(tx.Data()) + 1
		batch, ok := r.batchCache.Get(batchIndex)
		if ok {
			collectedL1Fee := new(big.Float).Quo(new(big.Float).SetInt(batch.CollectedL1Fee.ToInt()), new(big.Float).SetInt(big.NewInt(params.Ether)))
			collectedL1FeeFloat, _ := collectedL1Fee.Float64()

			// Update metrics
			r.collectedL1FeeSum += collectedL1FeeFloat
			r.metrics.CollectedL1FeeSum.Add(collectedL1FeeFloat)

			// Update leveldb
			err := r.ldb.PutFloat(collectedL1FeeSumKey, r.collectedL1FeeSum)
			if err != nil {
				log.Error("failed to update collected L1 fee sum in leveldb", "error", err)
			}

			log.Info("Updated L1 fee metrics",
				"batch_index", batchIndex,
				"l1_fee_eth", collectedL1FeeFloat)
		} else {
			log.Warn("batch not found in cache", "batch_index", batchIndex)
		}
	} else if method == constants.MethodFinalizeBatch {
		r.finalizeFeeSum += txFeeFloat
		r.metrics.FinalizeCostSum.Add(txFeeFloat)
		r.metrics.FinalizeCost.Set(txFeeFloat)
		// Update leveldb
		err := r.ldb.PutFloat(finalizeSumKey, r.finalizeFeeSum)
		if err != nil {
			return fmt.Errorf("failed to update finalize fee sum in leveldb: %w", err)
		}
	}

	return nil
}

type txStatus struct {
	state   int
	receipt *ethtypes.Receipt
}

const (
	txStatusPending = iota
	txStatusConfirmed
	txStatusMissing
)

func (r *Rollup) getTxStatus(tx *ethtypes.Transaction) (*txStatus, error) {
	receipt, err := r.L1Client.TransactionReceipt(context.Background(), tx.Hash())
	if err == nil {
		return &txStatus{state: txStatusConfirmed, receipt: receipt}, nil
	}

	if !utils.ErrStringMatch(err, ethereum.NotFound) {
		return nil, fmt.Errorf("query tx receipt error: %w", err)
	}

	// Check mempool
	_, isPending, err := r.L1Client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		if !utils.ErrStringMatch(err, ethereum.NotFound) {
			return nil, fmt.Errorf("query tx error: %w", err)
		}
		return &txStatus{state: txStatusMissing}, nil
	}

	if isPending {
		return &txStatus{state: txStatusPending}, nil
	}

	return &txStatus{state: txStatusMissing}, nil
}

func (r *Rollup) handlePendingTx(txRecord *types.TxRecord, tx *ethtypes.Transaction, method string) error {
	// Check for timeout
	if txRecord.SendTime+uint64(r.cfg.TxTimeout.Seconds()) < uint64(time.Now().Unix()) {
		log.Info("Transaction timed out",
			"tx", tx.Hash().Hex(),
			"nonce", tx.Nonce(),
			"method", method)

		// Try to replace the transaction with higher gas price
		return r.replaceTimedOutTx(tx)
	}

	// Check if transaction might fail
	if method == constants.MethodCommitBatch {
		batchIndex := utils.ParseParentBatchIndex(tx.Data()) + 1
		lastCommitted, err := r.Rollup.LastCommittedBatchIndex(nil)
		if err != nil {
			return fmt.Errorf("get last committed batch index error: %w", err)
		}

		if batchIndex <= lastCommitted.Uint64() {
			// This batch is already committed by another submitter
			log.Info("Batch already committed by another submitter, trying to cancel transaction",
				"batch_index", batchIndex,
				"last_committed", lastCommitted.Uint64(),
				"tx_hash", tx.Hash().String())

			// Try to cancel the transaction since it will fail
			cancelTx, err := r.CancelTx(tx)
			if err != nil {
				log.Error("Failed to cancel commit batch transaction",
					"error", err,
					"tx", tx.Hash().Hex(),
					"nonce", tx.Nonce(),
					"gas", tx.Gas(),
					"gas_tip_cap", tx.GasTipCap().String(),
					"gas_fee_cap", tx.GasFeeCap().String(),
					"blob_fee_cap", tx.BlobGasFeeCap().String(),
					"batch_index", batchIndex,
					"last_committed", lastCommitted.Uint64())
				return fmt.Errorf("cancel commit batch transaction failed: %w", err)
			}

			log.Info("Successfully sent cancel transaction for commit batch",
				"old_tx", tx.Hash().Hex(),
				"new_tx", cancelTx.Hash().String(),
				"nonce", tx.Nonce())
			if err := r.pendingTxs.Remove(tx.Hash()); err != nil {
				log.Error("failed to remove transaction", "hash", tx.Hash().String(), "error", err)
			}
			if err := r.pendingTxs.Add(cancelTx); err != nil {
				log.Error("failed to add cancel transaction", "hash", cancelTx.Hash().String(), "error", err)
			}
			return nil
		}
	} else if method == constants.MethodFinalizeBatch {
		batchIndex := utils.ParseFBatchIndex(tx.Data())
		lastFinalized, err := r.Rollup.LastFinalizedBatchIndex(nil)
		if err != nil {
			return fmt.Errorf("get last finalized batch index error: %w", err)
		}

		if batchIndex <= lastFinalized.Uint64() {
			// This batch is already finalized by another submitter
			log.Info("Batch already finalized by another submitter, trying to cancel transaction",
				"batch_index", batchIndex,
				"last_finalized", lastFinalized.Uint64(),
				"tx_hash", tx.Hash().String())

			// Try to cancel the transaction since it will fail
			cancelTx, err := r.CancelTx(tx)
			if err != nil {
				log.Error("Failed to cancel finalize batch transaction",
					"error", err,
					"tx", tx.Hash().Hex(),
					"nonce", tx.Nonce(),
					"gas", tx.Gas(),
					"gas_tip_cap", tx.GasTipCap().String(),
					"gas_fee_cap", tx.GasFeeCap().String(),
					"blob_fee_cap", tx.BlobGasFeeCap().String(),
					"batch_index", batchIndex,
					"last_finalized", lastFinalized.Uint64())
				return fmt.Errorf("cancel finalize batch transaction failed: %w", err)
			}

			log.Info("Successfully sent cancel transaction for finalize batch",
				"old_tx", tx.Hash().Hex(),
				"new_tx", cancelTx.Hash().String(),
				"nonce", tx.Nonce())
			if err := r.pendingTxs.Remove(tx.Hash()); err != nil {
				log.Error("failed to remove transaction", "hash", tx.Hash().String(), "error", err)
			}
			if err := r.pendingTxs.Add(cancelTx); err != nil {
				log.Error("failed to add cancel transaction", "hash", cancelTx.Hash().String(), "error", err)
			}
			return nil
		}
	}

	return nil
}

func (r *Rollup) replaceTimedOutTx(tx *ethtypes.Transaction) error {
	newTx, err := r.ReSubmitTx(false, tx)
	if err != nil {
		log.Error("Failed to resubmit transaction",
			"error", err,
			"tx", tx.Hash().Hex(),
			"nonce", tx.Nonce())
		return fmt.Errorf("resubmit tx error: %w", err)
	}

	log.Info("Successfully replaced transaction",
		"old_tx", tx.Hash().Hex(),
		"new_tx", newTx.Hash().String(),
		"nonce", tx.Nonce())

	if err := r.pendingTxs.Remove(tx.Hash()); err != nil {
		log.Error("failed to remove transaction", "hash", tx.Hash().String(), "error", err)
	}
	if err := r.pendingTxs.Add(newTx); err != nil {
		log.Error("failed to add new transaction", "hash", newTx.Hash().String(), "error", err)
	}
	return nil
}

func (r *Rollup) handleMissingTx(txRecord *types.TxRecord, tx *ethtypes.Transaction, method string) error {
	r.pendingTxs.IncQueryTimes(tx.Hash())

	// Mark transaction as unconfirmed since it's missing
	txRecord.Confirmed = false

	// Only resubmit after several retries
	if txRecord.QueryTimes >= 5 {
		return r.handleDiscardedTx(txRecord, tx, method)
	}

	log.Info("Transaction not found in mempool or chain",
		"hash", tx.Hash().String(),
		"nonce", tx.Nonce(),
		"query_times", txRecord.QueryTimes)

	return nil
}

func (r *Rollup) handleDiscardedTx(txRecord *types.TxRecord, tx *ethtypes.Transaction, method string) error {
	log.Warn("Transaction discarded",
		"hash", tx.Hash().String(),
		"nonce", tx.Nonce(),
		"query_times", txRecord.QueryTimes)

	// Try to resubmit with original parameters
	replacedTx, err := r.ReSubmitTx(true, tx)
	if err != nil {
		if utils.ErrStringMatch(err, core.ErrNonceTooLow) {
			// Transaction was probably confirmed in a reorg
			log.Info("Discarded transaction removed (nonce too low)",
				"hash", tx.Hash().String(),
				"nonce", tx.Nonce(),
				"method", method)
			if err := r.pendingTxs.Remove(tx.Hash()); err != nil {
				log.Error("failed to remove transaction", "hash", tx.Hash().String(), "error", err)
			}
			return nil
		}
		return fmt.Errorf("resend discarded tx: %w", err)
	}

	if err := r.pendingTxs.Remove(tx.Hash()); err != nil {
		log.Error("failed to remove transaction", "hash", tx.Hash().String(), "error", err)
	}
	if err := r.pendingTxs.Add(replacedTx); err != nil {
		log.Error("failed to add replaced transaction", "hash", replacedTx.Hash().String(), "error", err)
	}
	log.Info("Successfully resubmitted discarded transaction",
		"old_tx", tx.Hash().String(),
		"new_tx", replacedTx.Hash().String(),
		"nonce", replacedTx.Nonce())

	return nil
}

// handleConfirmedTx handles a confirmed transaction
func (r *Rollup) handleConfirmedTx(txRecord *types.TxRecord, tx *ethtypes.Transaction, txType string) error {
	status, err := r.getTxStatus(tx)
	if err != nil {
		return fmt.Errorf("get tx status error: %w", err)
	}

	// Get current block number for confirmation count
	currentBlock, err := r.L1Client.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("get current block number error: %w", err)
	}

	confirmations := currentBlock - status.receipt.BlockNumber.Uint64()
	log.Info("Transaction confirmation status",
		"hash", tx.Hash().String(),
		"block_number", status.receipt.BlockNumber.Uint64(),
		"current_block", currentBlock,
		"confirmations", confirmations)

	method := utils.ParseMethod(tx, r.abi)
	if status.receipt.Status == ethtypes.ReceiptStatusFailed {
		if method == constants.MethodCommitBatch {
			batchIndex := utils.ParseParentBatchIndex(tx.Data()) + 1
			lastCommitted, err := r.Rollup.LastCommittedBatchIndex(nil)
			if err != nil {
				return fmt.Errorf("get last committed batch index error: %w", err)
			}

			if batchIndex <= lastCommitted.Uint64() {
				// Another submitter has already committed this batch
				log.Warn("Batch commit transaction failed but batch is already committed by another submitter", "batch_index", batchIndex, "tx_hash", tx.Hash().String())
				// Clean up batch from cache since it's already committed
				r.batchCache.Delete(batchIndex)
			} else {
				// Contract bug detected - batch is not committed by anyone else but our transaction failed
				log.Warn("Critical error: batch commit transaction failed and batch is not committed by anyone", "batch_index", batchIndex, "tx_hash", tx.Hash().String())
			}
		} else if method == constants.MethodFinalizeBatch {
			batchIndex := utils.ParseFBatchIndex(tx.Data())
			lastFinalized, err := r.Rollup.LastFinalizedBatchIndex(nil)
			if err != nil {
				return fmt.Errorf("get last finalized batch index error: %w", err)
			}

			if batchIndex <= lastFinalized.Uint64() {
				// Another submitter has already finalized this batch
				log.Warn("Batch finalize transaction failed but batch is already finalized by another submitter", "batch_index", batchIndex, "tx_hash", tx.Hash().String())
			} else {
				// Contract bug detected - batch is not finalized by anyone else but our transaction failed
				log.Warn("Critical error: batch finalize transaction failed and batch is not finalized by anyone", "batch_index", batchIndex, "tx_hash", tx.Hash().String())
			}
		}
	} else { // Transaction succeeded
		// Get current block number for confirmation count only for successful transactions
		currentBlock, err := r.L1Client.BlockNumber(context.Background())
		if err != nil {
			return fmt.Errorf("get current block number error: %w", err)
		}
		confirmations := currentBlock - status.receipt.BlockNumber.Uint64()

		if method == constants.MethodCommitBatch {
			batchIndex := utils.ParseParentBatchIndex(tx.Data()) + 1
			log.Info("Successfully committed batch", "batch_index", batchIndex, "tx_hash", tx.Hash().String(), "block_number", status.receipt.BlockNumber.Uint64(), "gas_used", status.receipt.GasUsed, "confirm", confirmations)

			// Clean up batch from cache after successful commit
			r.batchCache.Delete(batchIndex)
		} else if method == constants.MethodFinalizeBatch {
			batchIndex := utils.ParseFBatchIndex(tx.Data())
			log.Info("Successfully finalized batch", "batch_index", batchIndex, "tx_hash", tx.Hash().String(), "block_number", status.receipt.BlockNumber.Uint64(), "gas_used", status.receipt.GasUsed, "confirm", confirmations)
		}
	}

	r.pendingTxs.MarkConfirmed(tx.Hash())
	return nil
}

func (r *Rollup) finalize() error {
	// get last finalized
	lastFinalized, err := r.Rollup.LastFinalizedBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last finalized error:%v", err)
	}
	// get last committed
	lastCommitted, err := r.Rollup.LastCommittedBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last committed error:%v", err)
	}

	target := big.NewInt(int64(r.pendingTxs.pfinalize + 1))
	if target.Cmp(lastFinalized) <= 0 {
		target = new(big.Int).Add(lastFinalized, big.NewInt(1))
	}

	if target.Cmp(lastCommitted) > 0 {
		log.Info("no need to finalize", "last_finalized", lastFinalized.Uint64(), "last_committed", lastCommitted.Uint64())
		return nil
	}

	log.Info("finalize info",
		"last_fianlzied", lastFinalized,
		"last_committed", lastCommitted,
		"finalize_index", target,
	)

	// batch exist
	existed, err := r.Rollup.BatchExist(nil, target)
	if err != nil {
		log.Error("query batch exist", "err", err)
		return err
	}
	if !existed {
		log.Warn("finalized batch not existed")
		return nil
	}

	// in challenge window
	inWindow, err := r.Rollup.BatchInsideChallengeWindow(nil, target)
	if err != nil {
		return fmt.Errorf("get batch inside challenge window error:%v", err)
	}
	if inWindow {
		log.Info("batch inside challenge window, wait")
		return nil
	}
	// finalize

	// get next batch
	nextBatchIndex := target.Uint64() + 1
	batch, err := GetRollupBatchByIndex(nextBatchIndex, r.L2Clients)
	if err != nil {
		log.Error("get next batch by index error",
			"batch_index", nextBatchIndex,
		)
		return fmt.Errorf("get next batch by index err:%v", err)
	}
	if batch == nil {
		log.Info("next batch is nil,wait next batch header to finalize", "next_batch_index", nextBatchIndex)
		return nil
	}

	// calldata
	calldata, err := r.abi.Pack("finalizeBatch", []byte(batch.ParentBatchHeader))
	if err != nil {
		return fmt.Errorf("pack finalizeBatch error:%v", err)
	}
	tip, feecap, _, err := r.GetGasTipAndCap()
	if err != nil {
		log.Error("get gas tip and cap error", "business", "finalize")
		return fmt.Errorf("get gas tip and cap error:%v", err)
	}

	gas, err := r.EstimateGas(r.WalletAddr(), r.rollupAddr, calldata, feecap, tip)
	if err != nil {
		log.Warn("estimate finalize tx gas error",
			"error", err,
		)

		if r.cfg.RoughEstimateGas {
			gas = r.RoughFinalizeGasEstimate()
			log.Info("rough estimate finalize gas", "gas", gas)
		} else {
			return fmt.Errorf("estimate finalize gas error:%v", err)
		}
	}

	// gas bump
	gas = r.BumpGas(gas)

	var nonce uint64
	if r.pendingTxs.pnonce != 0 {
		nonce = r.pendingTxs.pnonce + 1
	} else {
		nonce, err = r.L1Client.PendingNonceAt(context.Background(), r.WalletAddr())
		if err != nil {
			return fmt.Errorf("query layer1 nonce error:%v", err.Error())
		}
	}

	tx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   r.chainId,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: feecap,
		Gas:       gas,
		To:        &r.rollupAddr,
		Data:      calldata,
	})

	if uint64(tx.Size()) > txMaxSize {
		return core.ErrOversizedData
	}

	signedTx, err := r.Sign(tx)
	if err != nil {
		return fmt.Errorf("sign tx error:%v", err)
	}

	log.Info("finalize tx info",
		"batch_index", target,
		"last_committed", lastCommitted,
		"last_finalized", lastFinalized,
		"hash", signedTx.Hash().String(),
		"type", signedTx.Type(),
		"nonce", signedTx.Nonce(),
		"gas", signedTx.Gas(),
		"tip", signedTx.GasTipCap().String(),
		"fee_cap", signedTx.GasFeeCap().String(),
		"size", signedTx.Size(),
	)

	err = r.SendTx(signedTx)
	if err != nil {
		log.Error("send finalize tx to mempool", "error", err.Error())
		if utils.ErrStringMatch(err, core.ErrNonceTooLow) {
			// adjust nonce
			n1, _, err := utils.ParseNonce(err.Error())
			if err != nil {
				return fmt.Errorf("parse nonce err: %w", err)
			}
			r.pendingTxs.SetNonce(n1 - 1)
			log.Info("update pnonce", "nonce", n1-1)
		}
		return fmt.Errorf("send tx error:%v", err.Error())
	} else {
		log.Info("finalzie tx sent")

		r.pendingTxs.SetNonce(signedTx.Nonce())
		r.pendingTxs.SetPFinalize(target.Uint64())
		if err := r.pendingTxs.Add(signedTx); err != nil {
			log.Error("failed to add signed transaction", "hash", signedTx.Hash().String(), "error", err)
		}
	}

	return nil

}

func (r *Rollup) rollup() error {
	// Get current block height
	if !r.cfg.PriorityRollup {
		activeSubmitter, activeIndex, err := r.getCachedSubmitter()
		if err != nil {
			return fmt.Errorf("rollup: get current submitter err, %w", err)
		}

		err = r.eventInfoStorage.Load()
		if err != nil {
			return fmt.Errorf("failed to load storage in rollup: %w", err)
		}
		log.Debug("Indexer status",
			"blocks_processed", r.eventInfoStorage.BlockProcessed(),
			"last_event_time", r.eventInfoStorage.BlockTime())

		// get current blocknumber
		blockNumber, err := r.L1Client.BlockNumber(context.Background())
		if err != nil {
			return fmt.Errorf("failed to get block number in rollup: %w", err)
		}
		// set metrics
		r.metrics.SetIndexerBlockProcessed(r.eventInfoStorage.BlockProcessed())
		// check if indexed block number is too old
		if blockNumber > r.eventInfoStorage.BlockProcessed()+100 {
			log.Info("Indexer sync required",
				"module", r.GetModuleName(),
				"current_block", blockNumber,
				"processed_block", r.eventInfoStorage.BlockProcessed(),
				"blocks_behind", blockNumber-r.eventInfoStorage.BlockProcessed())
			return nil
		}

		past := (time.Now().Unix() - r.rotator.startTime.Int64()) % r.rotator.epoch.Int64()
		start := time.Now().Unix() - past
		end := start + r.rotator.epoch.Int64()

		// Calculate time remaining in current turn
		timeLeft := end - time.Now().Unix()
		myAddress := r.WalletAddr().Hex()
		activeAddress := activeSubmitter.Hex()
		isMyTurn := activeAddress == myAddress
		totalSubmitters := len(r.rotator.GetSubmitterSet())

		// Format timestamps for human readability
		startTimeFormatted := utils.FormatTime(big.NewInt(start))
		endTimeFormatted := utils.FormatTime(big.NewInt(end))
		timeLeftFormatted := fmt.Sprintf("%dm%ds", timeLeft/60, timeLeft%60)

		log.Info("Rotation status",
			"submitter_index", activeIndex,
			"active_submitter", activeAddress,
			"my_address", myAddress,
			"total_submitters", totalSubmitters,
			"is_my_turn", isMyTurn,
			"rotation_start", startTimeFormatted,
			"rotation_end", endTimeFormatted,
			"time_remaining", timeLeftFormatted)

		if isMyTurn {
			if timeLeft < r.cfg.RotatorBuffer {
				bufferFormatted := fmt.Sprintf("%dm%ds", r.cfg.RotatorBuffer/60, r.cfg.RotatorBuffer%60)
				log.Info("Insufficient time for rollup",
					"time_remaining", timeLeftFormatted,
					"buffer_required", bufferFormatted)
				return nil
			}

			log.Info("Starting rollup",
				"submitter_index", activeIndex,
				"total_submitters", totalSubmitters)
		} else {
			log.Debug("Skipping rollup - not active submitter",
				"active_index", activeIndex,
				"active_submitter", activeAddress)
			return nil
		}
	}

	if len(r.pendingTxs.txinfos) > int(r.cfg.MaxTxsInPendingPool) {
		log.Info("Pending pool full",
			"current_size", len(r.pendingTxs.txinfos),
			"max_size", r.cfg.MaxTxsInPendingPool)
		return nil
	}

	var batchIndex uint64

	cindexBig, err := r.Rollup.LastCommittedBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last committed batch index error:%v", err)
	}
	cindex := cindexBig.Uint64()

	switch {
	case r.pendingTxs.pindex != 0:
		batchIndex = max(cindex, r.pendingTxs.pindex) + 1
	default:
		batchIndex = cindex + 1
	}

	log.Debug("Batch status",
		"last_committed", cindex,
		"next_batch", batchIndex,
		"current_processing", r.pendingTxs.pindex)

	if r.pendingTxs.ExistedIndex(batchIndex) {
		log.Debug("Batch already committed",
			"batch_index", batchIndex)
		return nil
	}

	batch, ok := r.batchCache.Get(batchIndex)
	if !ok {
		log.Info("Batch not found in cache",
			"batch_index", batchIndex)
		return nil
	}

	signature, err := r.buildSignatureInput(batch)
	if err != nil {
		return err
	}
	rollupBatch := bindings.IRollupBatchDataInput{
		Version:           uint8(batch.Version),
		ParentBatchHeader: batch.ParentBatchHeader,
		BlockContexts:     batch.BlockContexts,
		PrevStateRoot:     batch.PrevStateRoot,
		PostStateRoot:     batch.PostStateRoot,
		WithdrawalRoot:    batch.WithdrawRoot,
	}

	// tip and cap
	tip, gasFeeCap, blobFee, err := r.GetGasTipAndCap()
	if err != nil {
		return fmt.Errorf("get gas tip and cap error:%v", err)
	}

	// calldata encode
	calldata, err := r.abi.Pack("commitBatch", rollupBatch, *signature)
	if err != nil {
		return fmt.Errorf("pack calldata error:%v", err)
	}
	// Estimate gas for transaction
	gas, err := r.EstimateGas(r.WalletAddr(), r.rollupAddr, calldata, gasFeeCap, tip)
	if err != nil {
		log.Warn("Estimate gas failed", "batch_index", batchIndex, "error", err)
		// Use rough estimation based on L1 message count
		if r.cfg.RoughEstimateGas {
			msgcnt := utils.ParseL1MessageCnt(batch.BlockContexts)
			gas = r.RoughRollupGasEstimate(msgcnt)
			log.Info("Using rough gas estimation",
				"batch_index", batchIndex,
				"gas_limit", gas,
				"l1_messages", msgcnt)
		} else {
			return nil
		}
	}

	// Apply gas buffer
	gas = r.BumpGas(gas)

	// Get next nonce
	nonce := r.getNextNonce()
	if nonce == 0 {
		return fmt.Errorf("failed to get next nonce")
	}

	// Create and sign transaction
	tx, err := r.createRollupTx(batch, nonce, gas, tip, gasFeeCap, blobFee, calldata)
	if err != nil {
		return fmt.Errorf("failed to create rollup tx: %w", err)
	}

	signedTx, err := r.Sign(tx)
	if err != nil {
		return fmt.Errorf("failed to sign tx: %w", err)
	}

	// Log transaction details before sending
	r.logTxInfo(signedTx, batchIndex)

	// Send transaction
	if err := r.SendTx(signedTx); err != nil {
		return fmt.Errorf("failed to send tx: %w", err)
	}

	// Update pending state
	r.pendingTxs.SetPindex(batchIndex)
	r.pendingTxs.SetNonce(tx.Nonce())
	if err := r.pendingTxs.Add(signedTx); err != nil {
		log.Error("Failed to track transaction", "error", err)
	}

	return nil
}

func (r *Rollup) getNextNonce() uint64 {
	if r.pendingTxs.pnonce != 0 {
		return r.pendingTxs.pnonce + 1
	}

	nonce, err := r.L1Client.PendingNonceAt(context.Background(), r.WalletAddr())
	if err != nil {
		log.Error("Failed to get nonce", "error", err)
		return 0
	}
	return nonce
}

func (r *Rollup) createRollupTx(batch *eth.RPCRollupBatch, nonce, gas uint64, tip, gasFeeCap, blobFee *big.Int, calldata []byte) (*ethtypes.Transaction, error) {
	if len(batch.Sidecar.Blobs) > 0 {
		return r.createBlobTx(batch, nonce, gas, tip, gasFeeCap, blobFee, calldata)
	}
	return r.createDynamicFeeTx(nonce, gas, tip, gasFeeCap, calldata)
}

func (r *Rollup) createBlobTx(batch *eth.RPCRollupBatch, nonce, gas uint64, tip, gasFeeCap, blobFee *big.Int, calldata []byte) (*ethtypes.Transaction, error) {
	versionedHashes := make([]common.Hash, 0, len(batch.Sidecar.Commitments))
	for _, commit := range batch.Sidecar.Commitments {
		versionedHashes = append(versionedHashes, kZGToVersionedHash(commit))
	}

	return ethtypes.NewTx(&ethtypes.BlobTx{
		ChainID:    uint256.MustFromBig(r.chainId),
		Nonce:      nonce,
		GasTipCap:  uint256.MustFromBig(big.NewInt(tip.Int64())),
		GasFeeCap:  uint256.MustFromBig(big.NewInt(gasFeeCap.Int64())),
		Gas:        gas,
		To:         r.rollupAddr,
		Data:       calldata,
		BlobFeeCap: uint256.MustFromBig(blobFee),
		BlobHashes: versionedHashes,
		Sidecar: &ethtypes.BlobTxSidecar{
			Blobs:       batch.Sidecar.Blobs,
			Commitments: batch.Sidecar.Commitments,
			Proofs:      batch.Sidecar.Proofs,
		},
	}), nil
}

func (r *Rollup) createDynamicFeeTx(nonce, gas uint64, tip, gasFeeCap *big.Int, calldata []byte) (*ethtypes.Transaction, error) {
	return ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   r.chainId,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &r.rollupAddr,
		Data:      calldata,
	}), nil
}

func (r *Rollup) logTxInfo(tx *ethtypes.Transaction, batchIndex uint64) {
	log.Info("Rollup transaction created",
		"batch_index", batchIndex,
		"hash", tx.Hash().String(),
		"type", tx.Type(),
		"nonce", tx.Nonce(),
		"gas", tx.Gas(),
		"tip", tx.GasTipCap().String(),
		"fee_cap", tx.GasFeeCap().String(),
		"blob_fee_cap", tx.BlobGasFeeCap(),
		"blob_gas", tx.BlobGas(),
		"size", tx.Size(),
		"blob_count", len(tx.BlobHashes()),
	)
}

func (r *Rollup) buildSignatureInput(batch *eth.RPCRollupBatch) (*bindings.IRollupBatchSignatureInput, error) {
	blsSignatures := batch.Signatures
	if len(blsSignatures) == 0 {
		return nil, fmt.Errorf("invalid batch signature")
	}
	signers := make([]common.Address, len(blsSignatures))
	sigs := make([]blssignatures.Signature, 0)
	for i, bz := range blsSignatures {
		if len(bz.Signature) > 0 {
			sig, err := blssignatures.SignatureFromBytes(bz.Signature)
			if err != nil {
				return nil, err
			}
			sigs = append(sigs, sig)
			signers[i] = bz.Signer
		}
	}
	aggregatedSig := blssignatures.AggregateSignatures(sigs)
	blsSignature := bls12381.NewG1().EncodePoint(aggregatedSig)

	// query bitmap of signers
	bm, err := r.Staking.GetStakersBitmap(nil, signers)
	if err != nil {
		return nil, fmt.Errorf("query stakers bitmap error:%v", err)
	}
	if bm == nil {
		return nil, errors.New("stakers bitmap is nil")
	}

	sigData := bindings.IRollupBatchSignatureInput{
		SignedSequencersBitmap: bm,
		SequencerSets:          batch.CurrentSequencerSetBytes,
		Signature:              blsSignature,
	}
	return &sigData, nil
}

func (r *Rollup) GetGasTipAndCap() (*big.Int, *big.Int, *big.Int, error) {

	head, err := r.L1Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, nil, err
	}
	if head.BaseFee != nil {
		log.Info("market fee info", "feecap", head.BaseFee)
		if r.cfg.MaxBaseFee > 0 && head.BaseFee.Cmp(big.NewInt(int64(r.cfg.MaxBaseFee))) > 0 {
			return nil, nil, nil, fmt.Errorf("base fee is too high, base fee %v exceeds max %v", head.BaseFee, r.cfg.MaxBaseFee)
		}
	}

	tip, err := r.L1Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	log.Info("market fee info", "tip", tip)

	if r.cfg.TipFeeBump > 0 {
		tip = new(big.Int).Mul(tip, big.NewInt(int64(r.cfg.TipFeeBump)))
		tip = new(big.Int).Div(tip, big.NewInt(100))
	}
	if r.cfg.MaxTip > 0 && tip.Cmp(big.NewInt(int64(r.cfg.MaxTip))) > 0 {
		return nil, nil, nil, fmt.Errorf("tip is too high, tip %v exceeds max %v", tip, r.cfg.MaxTip)
	}

	var gasFeeCap *big.Int
	if head.BaseFee != nil {
		gasFeeCap = new(big.Int).Add(
			tip,
			new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
		)
	} else {
		gasFeeCap = new(big.Int).Set(tip)
	}

	// calc blob fee cap
	var blobFee *big.Int
	if head.ExcessBlobGas != nil {
		blobFee = eip4844.CalcBlobFee(*head.ExcessBlobGas)
		// Set to 3x to handle blob market congestion
		blobFee = new(big.Int).Mul(blobFee, big.NewInt(3))
	}

	log.Info("fee info after bump",
		"tip", tip,
		"feecap", gasFeeCap,
		"blobfee", blobFee,
	)

	return tip, gasFeeCap, blobFee, nil
}

// PreCheck is run before the submitter to check whether the submitter can be started
func (r *Rollup) PreCheck() error {

	// debug stakers
	stakers, err := r.Staking.GetStakers(nil)
	if err != nil {
		log.Debug("get stakers error", "err", err)
	} else {
		log.Debug("stakers", "len", len(stakers))
		for _, s := range stakers {
			log.Debug("staker", "addr", s.Hex())
		}
	}
	// debug active stakers
	activeStakers, err := r.Staking.GetActiveStakers(nil)
	if err != nil {
		log.Debug("get active stakers error", "err", err)
	} else {
		log.Debug("active stakers", "len", len(activeStakers))
		for _, s := range activeStakers {
			log.Debug("active staker", "addr", s.Hex())
		}
	}

	isStaker, err := r.IsStaker()
	if err != nil {
		return fmt.Errorf("check if this account is sequencer error:%v", err)
	}

	if !isStaker {
		return fmt.Errorf("this account is not staker, can not rollup")
	}

	return nil
}

func (r *Rollup) WalletAddr() common.Address {

	if r.cfg.ExternalSign {
		return common.HexToAddress(r.cfg.ExternalSignAddress)
	} else {
		return crypto.PubkeyToAddress(r.privKey.PublicKey)
	}

}

func GetRollupBatchByIndex(index uint64, clients []iface.L2Client) (*eth.RPCRollupBatch, error) {
	if len(clients) < 1 {
		return nil, fmt.Errorf("no client to query")
	}
	for _, client := range clients {
		batch, err := client.GetRollupBatchByIndex(context.Background(), index)
		if err != nil {
			log.Warn("failed to get batch", "error", err)
			continue
		}
		if batch != nil && len(batch.Signatures) > 0 {
			return batch, nil
		}
	}

	return nil, nil
}

// query sequencer set from sequencer contract on l2
func QuerySequencerSet(addr common.Address, clients []iface.L2Client) ([]common.Address, error) {
	if len(clients) < 1 {
		return nil, fmt.Errorf("no client to query sequencer set")
	}
	for _, client := range clients {
		// l2 sequencer
		l2Seqencer, err := bindings.NewSequencer(addr, client)
		if err != nil {
			log.Warn("failed to connect to sequencer", "error", err)
			continue
		}
		// get sequencer set
		seqSet, err := l2Seqencer.GetSequencerSet2(nil)
		if err != nil {
			log.Warn("failed to get sequencer set", "error", err)
			continue
		}
		return seqSet, nil
	}
	return nil, fmt.Errorf("no sequencer set found after querying all clients")
}

// query epoch from gov contract on l2
func GetEpoch(addr common.Address, clients []iface.L2Client) (*big.Int, error) {
	if len(clients) < 1 {
		return nil, fmt.Errorf("no client to query epoch")
	}
	for _, client := range clients {
		// l2 gov
		l2Gov, err := bindings.NewGov(addr, client)
		if err != nil {
			log.Warn("failed to connect to gov", "error", err)
			continue
		}
		// get epoch
		epoch, err := l2Gov.RollupEpoch(nil)
		if err != nil {
			log.Warn("failed to get epoch", "error", err)
			continue
		}
		return epoch, nil
	}
	return nil, fmt.Errorf("no epoch found after querying all clients")
}

// query sequencer set update time from sequencer contract on l2
func GetSequencerSetUpdateTime(addr common.Address, clients []iface.L2Client) (*big.Int, error) {

	if len(clients) < 1 {
		return nil, fmt.Errorf("no client to query sequencer set update time")
	}
	for _, client := range clients {
		// l2 sequencer
		l2Seqencer, err := bindings.NewSequencer(addr, client)
		if err != nil {
			log.Warn("failed to connect to sequencer", "error", err)
			continue
		}
		// get sequencer set update time
		updateTime, err := l2Seqencer.UpdateTime(nil)
		if err != nil {
			log.Warn("failed to get sequencer set update time", "error", err)
			continue
		}
		return updateTime, nil
	}
	return nil, fmt.Errorf("no sequencer set update time found after querying all clients")
}

// query epoch update time from gov contract on l2
func GetEpochUpdateTime(addr common.Address, clients []iface.L2Client) (*big.Int, error) {
	if len(clients) < 1 {
		return nil, fmt.Errorf("no client to query epoch update time")
	}
	for _, client := range clients {
		// l2 gov
		l2Gov, err := bindings.NewGov(addr, client)
		if err != nil {
			log.Warn("failed to connect to gov", "error", err)
			continue
		}
		// get epoch update time
		updateTime, err := l2Gov.RollupEpochUpdateTime(nil)
		if err != nil {
			log.Warn("failed to get epoch update time", "error", err)
			continue
		}
		return updateTime, nil

	}
	return nil, fmt.Errorf("no epoch update time found after querying all clients")

}

func UpdateGasLimit(tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
	// add buffer to gas limit (*1.2)
	newGasLimit := tx.Gas() * 12 / 10

	var newTx *ethtypes.Transaction
	if tx.Type() == ethtypes.LegacyTxType {

		newTx = ethtypes.NewTx(&ethtypes.LegacyTx{
			Nonce:    tx.Nonce(),
			GasPrice: big.NewInt(tx.GasPrice().Int64()),
			Gas:      newGasLimit,
			To:       tx.To(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		})
	} else if tx.Type() == ethtypes.DynamicFeeTxType {
		newTx = ethtypes.NewTx(&ethtypes.DynamicFeeTx{
			Nonce:     tx.Nonce(),
			GasTipCap: big.NewInt(tx.GasTipCap().Int64()),
			GasFeeCap: big.NewInt(tx.GasFeeCap().Int64()),
			Gas:       newGasLimit,
			To:        tx.To(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
	} else if tx.Type() == ethtypes.BlobTxType {
		newTx = ethtypes.NewTx(&ethtypes.BlobTx{
			ChainID:    uint256.MustFromBig(tx.ChainId()),
			Nonce:      tx.Nonce(),
			GasTipCap:  uint256.MustFromBig(big.NewInt(tx.GasTipCap().Int64())),
			GasFeeCap:  uint256.MustFromBig(big.NewInt(tx.GasFeeCap().Int64())),
			Gas:        newGasLimit,
			To:         *tx.To(),
			Value:      uint256.MustFromBig(tx.Value()),
			Data:       tx.Data(),
			BlobFeeCap: uint256.MustFromBig(tx.BlobGasFeeCap()),
			BlobHashes: tx.BlobHashes(),
			Sidecar:    tx.BlobTxSidecar(),
		})

	} else {
		return nil, fmt.Errorf("unknown tx type:%v", tx.Type())
	}
	return newTx, nil
}

// send tx to l1 with business logic check
func (r *Rollup) SendTx(tx *ethtypes.Transaction) error {

	// judge tx info is valid
	if tx == nil {
		return errors.New("nil tx")
	}
	// l1 health check
	if r.bm != nil && !r.bm.IsGrowth() {
		return fmt.Errorf("block not growth in %d blocks time", r.cfg.BlockNotIncreasedThreshold)
	}

	err := sendTx(r.L1Client, r.cfg.TxFeeLimit, tx)
	if err != nil {
		return err
	}

	// after send tx
	// add to pending txs
	if r.pendingTxs != nil {
		if err := r.pendingTxs.Add(tx); err != nil {
			log.Error("failed to add transaction", "hash", tx.Hash().String(), "error", err)
		}
	}

	return nil

}

// send tx to l1 with business logic check
func sendTx(client iface.Client, txFeeLimit uint64, tx *ethtypes.Transaction) error {
	// fee limit
	if txFeeLimit > 0 {
		var fee uint64
		// calc tx gas fee
		if tx.Type() == ethtypes.BlobTxType {
			// blob fee
			fee = tx.BlobGasFeeCap().Uint64() * tx.BlobGas()
			// tx fee
			fee += tx.GasPrice().Uint64() * tx.Gas()
		} else {
			fee = tx.GasPrice().Uint64() * tx.Gas()
		}

		if fee > txFeeLimit {
			return fmt.Errorf("%v:limit=%v,but got=%v", utils.ErrExceedFeeLimit, txFeeLimit, fee)
		}
	}

	return client.SendTransaction(context.Background(), tx)
}

func (r *Rollup) ReSubmitTx(resend bool, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
	if tx == nil {
		return nil, errors.New("nil tx")
	}

	method := "replaced tx"
	if resend {
		method = "resubmitted tx"
	}

	// replaced tx info
	log.Info(method,
		"hash", tx.Hash().String(),
		"gas_fee_cap", tx.GasFeeCap().String(),
		"gas_tip", tx.GasTipCap().String(),
		"blob_fee_cap", tx.BlobGasFeeCap().String(),
		"gas", tx.Gas(),
		"nonce", tx.Nonce(),
	)

	tip, gasFeeCap, blobFeeCap, err := r.GetGasTipAndCap()
	if err != nil {
		return nil, fmt.Errorf("get gas tip and cap error:%w", err)
	}
	if !resend {
		// bump tip & feeCap
		bumpedFeeCap := calcThresholdValue(tx.GasFeeCap(), tx.Type() == ethtypes.BlobTxType)
		bumpedTip := calcThresholdValue(tx.GasTipCap(), tx.Type() == ethtypes.BlobTxType)

		// if bumpedTip > tip
		if bumpedTip.Cmp(tip) > 0 {
			tip = bumpedTip
		}

		if bumpedFeeCap.Cmp(gasFeeCap) > 0 {
			gasFeeCap = bumpedFeeCap
		}

		if tx.Type() == ethtypes.BlobTxType {
			bumpedBlobFeeCap := calcThresholdValue(tx.BlobGasFeeCap(), tx.Type() == ethtypes.BlobTxType)
			if bumpedBlobFeeCap.Cmp(blobFeeCap) > 0 {
				blobFeeCap = bumpedBlobFeeCap
			}
		}

		if r.cfg.MinTip > 0 && tip.Cmp(big.NewInt(int64(r.cfg.MinTip))) < 0 {
			log.Info("replace tip is too low, update tip to min tip ", "tip", tip, "min_tip", r.cfg.MinTip)
			tip = big.NewInt(int64(r.cfg.MinTip))
			// recalc feecap
			head, err := r.L1Client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				return nil, fmt.Errorf("get l1 head error:%w", err)
			}
			if head.BaseFee != nil {
				gasFeeCap = new(big.Int).Add(
					tip,
					new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
				)
			} else {
				gasFeeCap = new(big.Int).Set(tip)
			}
		}
	}

	var newTx *ethtypes.Transaction
	switch tx.Type() {
	case ethtypes.DynamicFeeTxType:
		newTx = ethtypes.NewTx(&ethtypes.DynamicFeeTx{
			ChainID:   tx.ChainId(),
			To:        tx.To(),
			Nonce:     tx.Nonce(),
			GasFeeCap: gasFeeCap,
			GasTipCap: tip,
			Gas:       tx.Gas(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
	case ethtypes.BlobTxType:

		newTx = ethtypes.NewTx(&ethtypes.BlobTx{
			ChainID:    uint256.MustFromBig(tx.ChainId()),
			Nonce:      tx.Nonce(),
			GasTipCap:  uint256.MustFromBig(tip),
			GasFeeCap:  uint256.MustFromBig(gasFeeCap),
			Gas:        tx.Gas(),
			To:         *tx.To(),
			Value:      uint256.MustFromBig(tx.Value()),
			Data:       tx.Data(),
			BlobFeeCap: uint256.MustFromBig(blobFeeCap),
			BlobHashes: tx.BlobHashes(),
			Sidecar:    tx.BlobTxSidecar(),
		})

	default:
		return nil, fmt.Errorf("replace unknown tx type:%v", tx.Type())

	}

	// weiToGwei converts wei value to gwei string representation
	weiToGwei := func(wei *big.Int) string {
		if wei == nil {
			return "0"
		}
		gwei := new(big.Float).Quo(
			new(big.Float).SetInt(wei),
			new(big.Float).SetInt64(1e9),
		)
		return gwei.Text('f', 6)
	}

	log.Info("new tx info",
		"tx_type", newTx.Type(),
		"gas_tip_gwei", weiToGwei(tip),
		"gas_fee_cap_gwei", weiToGwei(gasFeeCap),
		"blob_fee_cap_gwei", weiToGwei(blobFeeCap),
	)
	// sign tx
	newTx, err = r.Sign(newTx)
	if err != nil {
		return nil, fmt.Errorf("sign tx error:%w", err)
	}
	// send tx
	err = r.SendTx(newTx)
	if err != nil {
		return nil, fmt.Errorf("send tx error:%w", err)
	}

	return newTx, nil
}

func (r *Rollup) IsStaker() (bool, error) {

	isStaker, err := r.Staking.IsStaker(nil, r.WalletAddr())
	if err != nil {
		return false, fmt.Errorf("call IsStaker err :%v", err)
	}
	return isStaker, nil
}

func (r *Rollup) EstimateGas(from, to common.Address, data []byte, feecap *big.Int, tip *big.Int) (uint64, error) {

	gas, err := r.L1Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasFeeCap: feecap,
		GasTipCap: tip,
		Data:      data,
	})
	if err != nil {
		return 0, fmt.Errorf("call estimate gas error:%v", err)
	}
	return gas, nil
}

func (r *Rollup) BumpGas(origin uint64) uint64 {
	if r.cfg.GasLimitBuffer != 0 {
		return origin * (100 + r.cfg.GasLimitBuffer) / 100
	} else {
		return origin
	}
}

// for rollup
func (r *Rollup) RoughRollupGasEstimate(msgcnt uint64) uint64 {
	return r.cfg.RollupTxGasBase + msgcnt*r.cfg.RollupTxGasPerL1Msg
}

func (r *Rollup) RoughFinalizeGasEstimate() uint64 {
	return 500_000
}

func (r *Rollup) GetModuleName() string {
	return "rollup"
}

func (r *Rollup) InitFeeMetricsSum() error {
	// try to init rollupFeeSum & finalizeFeeSum
	// read rollupFeeSum
	rollupFeeSum, err := r.ldb.GetFloat(rollupSumKey)
	if err != nil {
		log.Warn("read rollupFeeSum from leveldb failed", "error", err)
		if utils.ErrStringMatch(err, db.ErrKeyNotFound) {
			err = r.ldb.PutFloat(rollupSumKey, 0)
			if err != nil {
				return fmt.Errorf("put rollupFeeSum to leveldb failed, key: %s, %w", rollupSumKey, err)
			}
		} else {
			return fmt.Errorf("get data from leveldb faild, key: %s, %w", rollupSumKey, err)
		}
	}
	log.Info(fmt.Sprintf("rollupFeeSum: %f", rollupFeeSum))
	finalizeFeeSum, err := r.ldb.GetFloat(finalizeSumKey)
	if err != nil {
		log.Warn("read finalizeFeeSum from leveldb failed", "error", err)
		if utils.ErrStringMatch(err, db.ErrKeyNotFound) {
			err = r.ldb.PutFloat(finalizeSumKey, 0)
			if err != nil {
				return fmt.Errorf("put finalizeFeeSum to leveldb failed, key: %s, %w", finalizeSumKey, err)
			}
		} else {
			return fmt.Errorf("get data from leveldb faild, key: %s, %w", finalizeSumKey, err)
		}
	}
	log.Info(fmt.Sprintf("finalizeFeeSum: %f", finalizeFeeSum))
	collectedL1FeeSum, err := r.ldb.GetFloat(collectedL1FeeSumKey)
	if err != nil {
		log.Warn("read collectedL1FeeSum from leveldb failed", "error", err)
		if utils.ErrStringMatch(err, db.ErrKeyNotFound) {
			err = r.ldb.PutFloat(collectedL1FeeSumKey, 0)
			if err != nil {
				return fmt.Errorf("put collectedL1FeeSum to leveldb failed, key: %s, %w", collectedL1FeeSumKey, err)
			}
		} else {
			return fmt.Errorf("get data from leveldb faild, key: %s, %w", collectedL1FeeSumKey, err)
		}
	}
	r.collectedL1FeeSum = collectedL1FeeSum
	log.Info(fmt.Sprintf("collectedL1FeeSum: %f", collectedL1FeeSum))

	r.rollupFeeSum = rollupFeeSum
	r.finalizeFeeSum = finalizeFeeSum
	r.collectedL1FeeSum = collectedL1FeeSum
	// set fee sum init val
	r.metrics.RollupCostSum.Add(r.rollupFeeSum)
	r.metrics.FinalizeCostSum.Add(r.finalizeFeeSum)
	r.metrics.CollectedL1FeeSum.Add(r.collectedL1FeeSum)
	return nil
}

// ClearPendingTxs clears all pending transactions
func (p *PendingTxs) ClearPendingTxs() {
	p.txinfos = make(map[common.Hash]*types.TxRecord)
}

// MarkUnconfirmed marks a transaction as unconfirmed in the pending pool
func (p *PendingTxs) MarkUnconfirmed(hash common.Hash) {
	if txRecord, ok := p.txinfos[hash]; ok {
		txRecord.Confirmed = false
	}
}

// CancelTx creates a new transaction with empty calldata to cancel the original transaction
func (r *Rollup) CancelTx(tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
	if tx == nil {
		return nil, errors.New("nil tx")
	}

	log.Info("canceling transaction",
		"hash", tx.Hash().String(),
		"gas_fee_cap", tx.GasFeeCap().String(),
		"gas_tip", tx.GasTipCap().String(),
		"blob_fee_cap", tx.BlobGasFeeCap().String(),
		"gas", tx.Gas(),
		"nonce", tx.Nonce(),
	)

	tip, gasFeeCap, blobFeeCap, err := r.GetGasTipAndCap()
	if err != nil {
		return nil, fmt.Errorf("get gas tip and cap error:%w", err)
	}

	// bump tip & feeCap
	bumpedFeeCap := calcThresholdValue(tx.GasFeeCap(), tx.Type() == ethtypes.BlobTxType)
	bumpedTip := calcThresholdValue(tx.GasTipCap(), tx.Type() == ethtypes.BlobTxType)

	// if bumpedTip > tip
	if bumpedTip.Cmp(tip) > 0 {
		tip = bumpedTip
	}

	if bumpedFeeCap.Cmp(gasFeeCap) > 0 {
		gasFeeCap = bumpedFeeCap
	}

	if tx.Type() == ethtypes.BlobTxType {
		bumpedBlobFeeCap := calcThresholdValue(tx.BlobGasFeeCap(), tx.Type() == ethtypes.BlobTxType)
		if bumpedBlobFeeCap.Cmp(blobFeeCap) > 0 {
			blobFeeCap = bumpedBlobFeeCap
		}
	}

	var newTx *ethtypes.Transaction
	switch tx.Type() {
	case ethtypes.DynamicFeeTxType:
		newTx = ethtypes.NewTx(&ethtypes.DynamicFeeTx{
			ChainID:   tx.ChainId(),
			To:        tx.To(),
			Nonce:     tx.Nonce(),
			GasFeeCap: gasFeeCap,
			GasTipCap: tip,
			Gas:       tx.Gas(),
			Value:     tx.Value(),
			Data:      []byte{}, // Empty calldata for cancellation
		})
	case ethtypes.BlobTxType:
		// For blob transactions, we need to keep one empty blob
		var emptyBlob kzg4844.Blob
		emptyCommitment, err := kzg4844.BlobToCommitment(&emptyBlob)
		if err != nil {
			return nil, fmt.Errorf("failed to create empty blob commitment: %w", err)
		}
		emptyProof, err := kzg4844.ComputeBlobProof(&emptyBlob, emptyCommitment)
		if err != nil {
			return nil, fmt.Errorf("failed to create empty blob proof: %w", err)
		}

		newTx = ethtypes.NewTx(&ethtypes.BlobTx{
			ChainID:    uint256.MustFromBig(tx.ChainId()),
			Nonce:      tx.Nonce(),
			GasTipCap:  uint256.MustFromBig(tip),
			GasFeeCap:  uint256.MustFromBig(gasFeeCap),
			Gas:        tx.Gas(),
			To:         *tx.To(),
			Value:      uint256.MustFromBig(tx.Value()),
			Data:       []byte{}, // Empty calldata for cancellation
			BlobFeeCap: uint256.MustFromBig(blobFeeCap),
			BlobHashes: []common.Hash{kZGToVersionedHash(emptyCommitment)},
			Sidecar: &ethtypes.BlobTxSidecar{
				Blobs:       []kzg4844.Blob{emptyBlob},
				Commitments: []kzg4844.Commitment{emptyCommitment},
				Proofs:      []kzg4844.Proof{emptyProof},
			},
		})
	default:
		return nil, fmt.Errorf("cancel unknown tx type:%v", tx.Type())
	}

	log.Info("new cancel tx info",
		"tx_type", newTx.Type(),
		"gas_tip_gwei", utils.WeiToGwei(tip),
		"gas_fee_cap_gwei", utils.WeiToGwei(gasFeeCap),
		"blob_fee_cap_gwei", utils.WeiToGwei(blobFeeCap),
	)

	// sign tx
	newTx, err = r.Sign(newTx)
	if err != nil {
		return nil, fmt.Errorf("sign tx error:%w", err)
	}

	// send tx
	err = r.SendTx(newTx)
	if err != nil {
		return nil, fmt.Errorf("send tx error:%w", err)
	}

	return newTx, nil
}

// Add this method to get cached submitter info
func (r *Rollup) getCachedSubmitter() (common.Address, uint64, error) {
	r.submitterCache.mu.RLock()
	if time.Since(r.submitterCache.lastUpdateTime) < r.submitterCacheTimeout {
		defer r.submitterCache.mu.RUnlock()
		return r.submitterCache.submitter, r.submitterCache.submitterIdx, nil
	}
	r.submitterCache.mu.RUnlock()

	// Need to update cache
	r.submitterCache.mu.Lock()
	defer r.submitterCache.mu.Unlock()

	// Double check after acquiring write lock
	if time.Since(r.submitterCache.lastUpdateTime) < r.submitterCacheTimeout {
		return r.submitterCache.submitter, r.submitterCache.submitterIdx, nil
	}

	// Get fresh data
	submitter, idx, err := r.rotator.CurrentSubmitter(r.L2Clients, r.Staking)
	if err != nil {
		return common.Address{}, 0, err
	}

	// Update cache with proper type conversion
	r.submitterCache.submitter = *submitter
	r.submitterCache.submitterIdx = uint64(idx)
	r.submitterCache.lastUpdateTime = time.Now()

	return *submitter, uint64(idx), nil
}
