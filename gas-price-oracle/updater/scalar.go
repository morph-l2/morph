package updater

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/morph/bindings/bindings"
	"github.com/morph-l2/morph/gas-price-oracle/calc"
	"github.com/morph-l2/morph/gas-price-oracle/client"
	"github.com/morph-l2/morph/gas-price-oracle/metrics"
	"github.com/sirupsen/logrus"
)

// ScalarUpdater Scalar updater
type ScalarUpdater struct {
	l1Client       *client.L1Client
	l2Client       *client.L2Client
	beaconClient   *client.BeaconClient
	oracleContract *bindings.GasPriceOracle
	rollupContract *bindings.Rollup
	txManager      *TxManager
	calculator     *calc.ScalarCalculator
	blobProcessor  calc.BlobDataProcessor
	gasThreshold   uint64
	updateCounter  uint64
	updateInterval uint64
	log            *logrus.Entry
}

// NewScalarUpdater creates Scalar updater
func NewScalarUpdater(
	l1Client *client.L1Client,
	l2Client *client.L2Client,
	beaconClient *client.BeaconClient,
	oracleContract *bindings.GasPriceOracle,
	rollupContract *bindings.Rollup,
	txManager *TxManager,
	gasThreshold uint64,
	updateInterval uint64,
	txnPerBatch uint64,
) *ScalarUpdater {
	return &ScalarUpdater{
		l1Client:       l1Client,
		l2Client:       l2Client,
		beaconClient:   beaconClient,
		oracleContract: oracleContract,
		rollupContract: rollupContract,
		txManager:      txManager,
		calculator:     calc.NewScalarCalculator(txnPerBatch),
		blobProcessor:  &calc.DefaultBlobProcessor{}, // Using default processor
		gasThreshold:   gasThreshold,
		updateCounter:  0,
		updateInterval: updateInterval,
		log:            logrus.WithField("component", "scalar_updater"),
	}
}

// ShouldUpdate checks whether an update should be performed
func (u *ScalarUpdater) ShouldUpdate() bool {
	u.updateCounter++
	if u.updateCounter >= u.updateInterval {
		u.updateCounter = 0
		return true
	}
	return false
}

// Update performs one scalar update
func (u *ScalarUpdater) Update(ctx context.Context) error {
	// Step 1: Find recent commit batch events
	commitScalar, blobScalar, err := u.calculateLatestScalar(ctx)
	if err != nil {
		return err
	}

	// Step 2: Get current scalar values on L2
	callOpts := &bind.CallOpts{Context: ctx}

	currentCommitScalar, err := u.oracleContract.CommitScalar(callOpts)
	if err != nil {
		return err
	}

	currentBlobScalar, err := u.oracleContract.BlobScalar(callOpts)
	if err != nil {
		return err
	}

	u.log.WithFields(logrus.Fields{
		"latest_commit_scalar":  commitScalar,
		"current_commit_scalar": currentCommitScalar.Uint64(),
		"latest_blob_scalar":    blobScalar,
		"current_blob_scalar":   currentBlobScalar.Uint64(),
	}).Info("Calculated scalars")

	// Limit to maximum value
	if commitScalar > calc.MaxCommitScalar {
		commitScalar = calc.MaxCommitScalar
	}
	if blobScalar > calc.MaxBlobScalar {
		blobScalar = calc.MaxBlobScalar
	}

	// Update metrics
	metrics.CommitScalar.Set(float64(commitScalar / calc.Precision))
	metrics.BlobScalar.Set(float64(blobScalar) / calc.Precision)

	// Step 3: Check if commit scalar needs to be updated
	if calc.ShouldUpdate(commitScalar, currentCommitScalar.Uint64(), u.gasThreshold) {
		if err := u.updateCommitScalar(ctx, commitScalar); err != nil {
			return fmt.Errorf("failed to update commit scalar: %w", err)
		}
	}

	// Step 4: Check if blob scalar needs to be updated
	if calc.ShouldUpdate(blobScalar, currentBlobScalar.Uint64(), u.gasThreshold) {
		if err := u.updateBlobScalar(ctx, blobScalar); err != nil {
			return fmt.Errorf("failed to update blob scalar: %w", err)
		}
	}

	return nil
}

// calculateLatestScalar calculates the latest scalar values
func (u *ScalarUpdater) calculateLatestScalar(ctx context.Context) (commitScalar, blobScalar uint64, err error) {
	// Get CommitBatch events from last 100 blocks
	currentBlock, err := u.l1Client.GetBlockNumber(ctx)
	if err != nil {
		return 0, 0, err
	}

	startBlock := currentBlock - 100
	if startBlock > currentBlock {
		startBlock = 1
	}

	// Query CommitBatch events
	filterOpts := &bind.FilterOpts{
		Start:   startBlock,
		End:     &currentBlock,
		Context: ctx,
	}

	iterator, err := u.rollupContract.FilterCommitBatch(filterOpts, nil, nil)
	if err != nil {
		return 0, 0, err
	}
	defer iterator.Close()

	// Find the latest event
	var latestEvent *bindings.RollupCommitBatch
	for iterator.Next() {
		latestEvent = iterator.Event
	}

	if latestEvent == nil {
		u.log.Warn("No commit batch events found in recent blocks, skipping update")
		return 0, 0, fmt.Errorf("no commit batch events found")
	}

	u.log.WithFields(logrus.Fields{
		"tx_hash":     latestEvent.Raw.TxHash.Hex(),
		"block_num":   latestEvent.Raw.BlockNumber,
		"batch_index": latestEvent.BatchIndex,
	}).Info("Found latest commit batch event")

	// Calculate scalar from event
	return u.calculateScalarFromTx(ctx, latestEvent.Raw.TxHash, latestEvent.Raw.BlockNumber)
}

// calculateScalarFromTx calculates scalar from transaction
func (u *ScalarUpdater) calculateScalarFromTx(ctx context.Context, txHash common.Hash, blockNum uint64) (commitScalar, blobScalar uint64, err error) {
	// Step 1: Get transaction
	tx, _, err := u.l1Client.GetTransaction(ctx, txHash)
	if err != nil {
		return 0, 0, err
	}

	// Step 2: Parse commitBatch call data
	l2TxCount, chunks, err := u.parseCommitBatchData(tx.Data())
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse commit batch data: %w", err)
	}

	u.log.WithFields(logrus.Fields{
		"l2_tx_count": l2TxCount,
		"num_chunks":  len(chunks),
	}).Debug("Parsed commit batch data")

	// Update metrics
	metrics.TxnPerBatch.Set(float64(l2TxCount))

	// Step 3: Get transaction receipt
	receipt, err := u.l1Client.GetTransactionReceipt(ctx, txHash)
	if err != nil {
		return 0, 0, err
	}

	rollupGasUsed := receipt.GasUsed
	if rollupGasUsed == 0 {
		return 0, 0, fmt.Errorf("rollup gas used is zero")
	}

	// Step 4: Calculate L2 data length (from blob)
	// TODO: Implement blob data extraction
	// Currently using placeholder value
	l2DataLen := uint64(0)

	// If transaction contains blob, try to extract (skip for now)
	u.log.Warn("Blob data extraction not implemented, using default blob scalar")

	// Step 5: Calculate scalar
	commitScalar, blobScalar = u.calculator.CalculateScalars(rollupGasUsed, l2TxCount, l2DataLen)

	u.log.WithFields(logrus.Fields{
		"rollup_gas_used": rollupGasUsed,
		"l2_tx_count":     l2TxCount,
		"l2_data_len":     l2DataLen,
		"commit_scalar":   commitScalar / calc.Precision,
		"blob_scalar":     float64(blobScalar) / calc.Precision,
	}).Info("Calculated scalars")

	return commitScalar, blobScalar, nil
}

// parseCommitBatchData parses commitBatch call data
func (u *ScalarUpdater) parseCommitBatchData(data []byte) (l2TxCount uint64, chunks [][]byte, err error) {
	// TODO: Parse according to actual Rollup contract ABI
	// This provides a simplified implementation framework

	// Parse ABI
	rollupABI, err := abi.JSON(strings.NewReader(bindings.RollupMetaData.ABI))
	if err != nil {
		return 0, nil, err
	}

	// Parse method call
	method, err := rollupABI.MethodById(data[:4])
	if err != nil {
		return 0, nil, err
	}

	if method.Name != "commitBatch" {
		return 0, nil, fmt.Errorf("unexpected method: %s", method.Name)
	}

	// Parse parameters
	params := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(params, data[4:])
	if err != nil {
		return 0, nil, err
	}

	// Extract chunks
	// TODO: Adjust according to actual contract structure
	// Parse based on your commitBatch function signature

	// Temporary implementation: Return default values
	u.log.Warn("Chunk parsing not fully implemented, using default values")

	return 50, [][]byte{}, nil // Return default values
}

// updateCommitScalar updates commit scalar
func (u *ScalarUpdater) updateCommitScalar(ctx context.Context, commitScalar uint64) error {
	receipt, err := u.txManager.SendTransaction(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return u.oracleContract.SetCommitScalar(auth, new(big.Int).SetUint64(commitScalar))
	})
	if err != nil {
		u.log.WithError(err).Error("Failed to send set commit scalar transaction")
		return err
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed: %s", receipt.TxHash.Hex())
	}

	u.log.WithFields(logrus.Fields{
		"tx_hash":       receipt.TxHash.Hex(),
		"gas_used":      receipt.GasUsed,
		"commit_scalar": commitScalar,
	}).Info("Successfully updated commit scalar")

	metrics.ScalarUpdateCount.Inc()
	return nil
}

// updateBlobScalar updates blob scalar
func (u *ScalarUpdater) updateBlobScalar(ctx context.Context, blobScalar uint64) error {
	receipt, err := u.txManager.SendTransaction(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return u.oracleContract.SetBlobScalar(auth, new(big.Int).SetUint64(blobScalar))
	})
	if err != nil {
		u.log.WithError(err).Error("Failed to send set blob scalar transaction")
		return err
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed: %s", receipt.TxHash.Hex())
	}

	u.log.WithFields(logrus.Fields{
		"tx_hash":     receipt.TxHash.Hex(),
		"gas_used":    receipt.GasUsed,
		"blob_scalar": blobScalar,
	}).Info("Successfully updated blob scalar")

	metrics.ScalarUpdateCount.Inc()
	return nil
}
