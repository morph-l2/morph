package updater

import (
	"context"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/params"
	"github.com/morph-l2/morph/bindings/bindings"
	"github.com/morph-l2/morph/gas-price-oracle/calc"
	"github.com/morph-l2/morph/gas-price-oracle/client"
	"github.com/morph-l2/morph/gas-price-oracle/metrics"
	"github.com/sirupsen/logrus"
)

const (
	// MaxBaseFee maximum base fee (1000 Gwei)
	MaxBaseFee = 1000 * params.GWei
)

// BaseFeeUpdater Base Fee updater
type BaseFeeUpdater struct {
	l1Client       *client.L1Client
	l2Client       *client.L2Client
	oracleContract *bindings.GasPriceOracle
	txManager      *TxManager
	gasThreshold   uint64
	interval       time.Duration
	log            *logrus.Entry
}

// NewBaseFeeUpdater creates Base Fee updater
func NewBaseFeeUpdater(
	l1Client *client.L1Client,
	l2Client *client.L2Client,
	oracleContract *bindings.GasPriceOracle,
	txManager *TxManager,
	gasThreshold uint64,
	interval time.Duration,
) *BaseFeeUpdater {
	return &BaseFeeUpdater{
		l1Client:       l1Client,
		l2Client:       l2Client,
		oracleContract: oracleContract,
		txManager:      txManager,
		gasThreshold:   gasThreshold,
		interval:       interval,
		log:            logrus.WithField("component", "basefee_updater"),
	}
}

// Start starts the updater
func (u *BaseFeeUpdater) Start(ctx context.Context) {
	ticker := time.NewTicker(u.interval)
	defer ticker.Stop()

	u.log.Info("Base fee updater started")

	for {
		select {
		case <-ctx.Done():
			u.log.Info("Base fee updater stopped")
			return
		case <-ticker.C:
			if err := u.update(ctx); err != nil {
				u.log.WithError(err).Error("Failed to update base fee")
				metrics.UpdateErrors.WithLabelValues("basefee").Inc()
			}
		}
	}
}

// update performs one update
func (u *BaseFeeUpdater) update(ctx context.Context) error {
	// Step 1: Fetch L1 data
	l1BaseFee, l1BlobBaseFee, err := u.l1Client.GetBaseFee(ctx)
	if err != nil {
		return err
	}

	l1GasPrice, err := u.l1Client.GetGasPrice(ctx)
	if err != nil {
		metrics.L1RPCStatus.Set(2) // error status
		return err
	}
	metrics.L1RPCStatus.Set(0) // normal status

	// Validate data validity
	if l1BaseFee.Sign() == 0 || l1BlobBaseFee.Sign() == 0 || l1GasPrice.Sign() == 0 {
		return nil
	}

	// Record metrics
	l1BaseFeeGwei := new(big.Float).Quo(
		new(big.Float).SetInt(l1BaseFee),
		big.NewFloat(params.GWei),
	)
	l1BaseFeeFloat, _ := l1BaseFeeGwei.Float64()
	metrics.L1BaseFee.Set(l1BaseFeeFloat)

	u.log.WithFields(logrus.Fields{
		"l1_base_fee":      l1BaseFee,
		"l1_blob_base_fee": l1BlobBaseFee,
		"l1_gas_price":     l1GasPrice,
	}).Debug("Fetched L1 gas data")

	// Step 2: Get current values on L2
	callOpts := &bind.CallOpts{Context: ctx}

	l1BaseFeeOnL2, err := u.oracleContract.L1BaseFee(callOpts)
	if err != nil {
		return err
	}

	l1BlobBaseFeeOnL2, err := u.oracleContract.L1BlobBaseFee(callOpts)
	if err != nil {
		return err
	}

	// Record L2 metrics
	l1BaseFeeOnL2Gwei := new(big.Float).Quo(
		new(big.Float).SetInt(l1BaseFeeOnL2),
		big.NewFloat(params.GWei),
	)
	l1BaseFeeOnL2Float, _ := l1BaseFeeOnL2Gwei.Float64()
	metrics.L1BaseFeeOnL2.Set(l1BaseFeeOnL2Float)

	l1BlobBaseFeeOnL2Gwei := new(big.Float).Quo(
		new(big.Float).SetInt(l1BlobBaseFeeOnL2),
		big.NewFloat(params.GWei),
	)
	l1BlobBaseFeeOnL2Float, _ := l1BlobBaseFeeOnL2Gwei.Float64()
	metrics.L1BlobBaseFeeOnL2.Set(l1BlobBaseFeeOnL2Float)

	u.log.WithFields(logrus.Fields{
		"l1_base_fee_on_l2":      l1BaseFeeOnL2,
		"l1_blob_base_fee_on_l2": l1BlobBaseFeeOnL2,
	}).Debug("Fetched L2 oracle data")

	// Step 3: Update base fee
	if err := u.updateBaseFee(ctx, l1BaseFee, l1BlobBaseFee, l1BaseFeeOnL2, l1BlobBaseFeeOnL2); err != nil {
		return err
	}

	// Step 4: Update wallet balance metric
	balance, err := u.l2Client.GetBalance(ctx, u.l2Client.WalletAddress())
	if err != nil {
		u.log.WithError(err).Warn("Failed to get wallet balance")
	} else {
		balanceEth := new(big.Float).Quo(
			new(big.Float).SetInt(balance),
			big.NewFloat(params.Ether),
		)
		balanceFloat, _ := balanceEth.Float64()
		metrics.GasOracleOwnerBalance.Set(balanceFloat)
	}

	return nil
}

// updateBaseFee updates base fee to L2
func (u *BaseFeeUpdater) updateBaseFee(ctx context.Context, l1BaseFee, l1BlobBaseFee, l1BaseFeeOnL2, l1BlobBaseFeeOnL2 *big.Int) error {
	// Limit to maximum value
	if l1BaseFee.Cmp(big.NewInt(MaxBaseFee)) > 0 {
		l1BaseFee = big.NewInt(MaxBaseFee)
	}
	if l1BlobBaseFee.Cmp(big.NewInt(MaxBaseFee)) > 0 {
		l1BlobBaseFee = big.NewInt(MaxBaseFee)
	}

	// Check if blob base fee needs to be updated
	needUpdateBlob := calc.ShouldUpdateBigInt(l1BlobBaseFee, l1BlobBaseFeeOnL2, u.gasThreshold)

	u.log.WithFields(logrus.Fields{
		"need_update_blob":       needUpdateBlob,
		"l1_blob_base_fee":       l1BlobBaseFee,
		"l1_blob_base_fee_on_l2": l1BlobBaseFeeOnL2,
	}).Info("Check blob base fee update")

	if needUpdateBlob && l1BlobBaseFee.Sign() > 0 {
		// Update both base fee and blob base fee
		receipt, err := u.txManager.SendTransaction(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
			return u.oracleContract.SetL1BaseFeeAndBlobBaseFee(auth, l1BaseFee, l1BlobBaseFee)
		})
		if err != nil {
			u.log.WithError(err).Error("Failed to send set L1 base fee and blob base fee transaction")
			return err
		}

		if receipt.Status == 0 {
			u.log.WithField("tx_hash", receipt.TxHash.Hex()).Error("Transaction failed")
			return nil
		}

		u.log.WithFields(logrus.Fields{
			"tx_hash":  receipt.TxHash.Hex(),
			"gas_used": receipt.GasUsed,
		}).Info("Successfully updated L1 base fee and blob base fee")

		metrics.BaseFeeUpdateCount.Inc()
		return nil
	}

	// Only check if base fee needs to be updated
	needUpdateBase := calc.ShouldUpdateBigInt(l1BaseFee, l1BaseFeeOnL2, u.gasThreshold)

	u.log.WithFields(logrus.Fields{
		"need_update_base":  needUpdateBase,
		"l1_base_fee":       l1BaseFee,
		"l1_base_fee_on_l2": l1BaseFeeOnL2,
	}).Info("Check base fee update")

	if needUpdateBase && l1BaseFee.Sign() > 0 {
		receipt, err := u.txManager.SendTransaction(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
			return u.oracleContract.SetL1BaseFee(auth, l1BaseFee)
		})
		if err != nil {
			u.log.WithError(err).Error("Failed to send set L1 base fee transaction")
			return err
		}

		if receipt.Status == 0 {
			u.log.WithField("tx_hash", receipt.TxHash.Hex()).Error("Transaction failed")
			return nil
		}

		u.log.WithFields(logrus.Fields{
			"tx_hash":  receipt.TxHash.Hex(),
			"gas_used": receipt.GasUsed,
		}).Info("Successfully updated L1 base fee")

		metrics.BaseFeeUpdateCount.Inc()
	}

	return nil
}
