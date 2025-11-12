package updater

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/morph/bindings/bindings"
	"github.com/morph-l2/morph/gas-price-oracle/calc"
	"github.com/morph-l2/morph/gas-price-oracle/client"
	"github.com/morph-l2/morph/gas-price-oracle/metrics"
	"github.com/sirupsen/logrus"
)

// PriceUpdater handles token price updates
type PriceUpdater struct {
	l2Client         *client.L2Client
	registryContract *bindings.L2TokenRegistry
	priceFeed        client.PriceFeed
	txManager        *TxManager
	tokenIDs         []uint16
	interval         time.Duration
	priceThreshold   uint64
	log              *logrus.Entry

	// Cache of last updated prices
	lastPrices map[uint16]*big.Int
	mu         sync.RWMutex
	stopChan   chan struct{}
}

// NewPriceUpdater creates a new price updater
func NewPriceUpdater(
	l2Client *client.L2Client,
	registryContract *bindings.L2TokenRegistry,
	priceFeed client.PriceFeed,
	txManager *TxManager,
	tokenIDs []uint16,
	interval time.Duration,
	priceThreshold uint64,
) *PriceUpdater {
	return &PriceUpdater{
		l2Client:         l2Client,
		registryContract: registryContract,
		priceFeed:        priceFeed,
		txManager:        txManager,
		tokenIDs:         tokenIDs,
		interval:         interval,
		priceThreshold:   priceThreshold,
		log:              logrus.WithField("component", "price_updater"),
		lastPrices:       make(map[uint16]*big.Int),
		stopChan:         make(chan struct{}),
	}
}

// Name returns the updater name
func (u *PriceUpdater) Name() string {
	return "price"
}

// Start starts the price updater
func (u *PriceUpdater) Start(ctx context.Context) error {
	go func() {
		ticker := time.NewTicker(u.interval)
		defer ticker.Stop()

		u.log.WithFields(logrus.Fields{
			"token_ids":       u.tokenIDs,
			"interval":        u.interval,
			"price_threshold": u.priceThreshold,
		}).Info("Price updater started")

		// Perform initial update (fetch current prices from contract)
		if err := u.initializePriceCache(ctx); err != nil {
			u.log.WithError(err).Warn("Failed to initialize price cache, will start fresh")
		}

		// Perform first actual update
		if err := u.update(ctx); err != nil {
			u.log.WithError(err).Error("Initial price update failed")
		}

		for {
			select {
			case <-ctx.Done():
				u.log.Info("Price updater stopped by context")
				return
			case <-u.stopChan:
				u.log.Info("Price updater stopped")
				return
			case <-ticker.C:
				if err := u.update(ctx); err != nil {
					u.log.WithError(err).Error("Failed to update prices")
					metrics.UpdateErrors.WithLabelValues("price").Inc()
				}
			}
		}
	}()
	return nil
}

// Stop gracefully stops the updater
func (u *PriceUpdater) Stop() error {
	close(u.stopChan)
	u.log.Info("Price updater stop requested")
	return nil
}

// initializePriceCache fetches current prices from contract and caches them
func (u *PriceUpdater) initializePriceCache(ctx context.Context) error {
	callOpts := &bind.CallOpts{Context: ctx}

	u.mu.Lock()
	defer u.mu.Unlock()

	for _, tokenID := range u.tokenIDs {
		price, err := u.registryContract.GetTokenPrice(callOpts, tokenID)
		if err != nil {
			u.log.WithFields(logrus.Fields{
				"token_id": tokenID,
				"error":    err,
			}).Debug("Failed to get current price for token")
			continue
		}

		if price.Sign() > 0 {
			u.lastPrices[tokenID] = price
			u.log.WithFields(logrus.Fields{
				"token_id": tokenID,
				"price":    price.String(),
			}).Debug("Cached current price")
		}
	}

	u.log.WithField("cached_count", len(u.lastPrices)).Info("Initialized price cache")
	return nil
}

// update performs one price update
func (u *PriceUpdater) update(ctx context.Context) error {
	if len(u.tokenIDs) == 0 {
		u.log.Debug("No tokens to update")
		return nil
	}

	// Step 1: Fetch new prices from feed
	newPrices, err := u.priceFeed.GetBatchPriceRatios(ctx, u.tokenIDs)
	if err != nil {
		return fmt.Errorf("failed to fetch price ratios: %w", err)
	}

	// Step 2: Filter prices that need updating based on threshold
	var tokenIDsToUpdate []uint16
	var pricesToUpdate []*big.Int

	u.mu.RLock()
	for tokenID, newPrice := range newPrices {
		if newPrice == nil || newPrice.Sign() == 0 {
			u.log.WithField("token_id", tokenID).Warn("Skipping zero price")
			continue
		}

		// Check if price changed significantly
		lastPrice, exists := u.lastPrices[tokenID]
		if exists && lastPrice.Sign() > 0 {
			// Calculate if price change exceeds threshold
			if !u.shouldUpdatePrice(lastPrice, newPrice) {
				u.log.WithFields(logrus.Fields{
					"token_id":   tokenID,
					"last_price": lastPrice.String(),
					"new_price":  newPrice.String(),
					"threshold":  u.priceThreshold,
				}).Debug("Price change below threshold, skipping update")
				continue
			}

			u.log.WithFields(logrus.Fields{
				"token_id":   tokenID,
				"last_price": lastPrice.String(),
				"new_price":  newPrice.String(),
			}).Info("Price change exceeds threshold, will update")
		} else {
			u.log.WithFields(logrus.Fields{
				"token_id":  tokenID,
				"new_price": newPrice.String(),
			}).Info("First time update for token")
		}

		tokenIDsToUpdate = append(tokenIDsToUpdate, tokenID)
		pricesToUpdate = append(pricesToUpdate, newPrice)
	}
	u.mu.RUnlock()

	if len(tokenIDsToUpdate) == 0 {
		u.log.Debug("No prices need updating (all changes below threshold)")
		return nil
	}

	u.log.WithFields(logrus.Fields{
		"token_count":  len(tokenIDsToUpdate),
		"token_ids":    tokenIDsToUpdate,
		"total_tokens": len(u.tokenIDs),
	}).Info("Updating token prices")

	// Step 3: Update prices on L2
	receipt, err := u.txManager.SendTransaction(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return u.registryContract.BatchUpdatePrices(auth, tokenIDsToUpdate, pricesToUpdate)
	})
	if err != nil {
		return fmt.Errorf("failed to send batch update prices transaction: %w", err)
	}

	if receipt.Status == 0 {
		u.log.WithField("tx_hash", receipt.TxHash.Hex()).Error("Transaction failed")
		return fmt.Errorf("transaction failed on-chain: %s", receipt.TxHash.Hex())
	}

	// Step 4: Update cache with new prices
	u.mu.Lock()
	for i, tokenID := range tokenIDsToUpdate {
		u.lastPrices[tokenID] = pricesToUpdate[i]
	}
	u.mu.Unlock()

	u.log.WithFields(logrus.Fields{
		"tx_hash":     receipt.TxHash.Hex(),
		"gas_used":    receipt.GasUsed,
		"token_count": len(tokenIDsToUpdate),
	}).Info("Successfully updated token prices")

	// Step 5: Update metrics
	for i, tokenID := range tokenIDsToUpdate {
		u.log.WithFields(logrus.Fields{
			"token_id":    tokenID,
			"price_ratio": pricesToUpdate[i].String(),
		}).Debug("Price updated")
	}

	metrics.ScalarUpdateCount.Inc()

	return nil
}

// shouldUpdatePrice checks if the price change exceeds the threshold
// Uses the same logic as calc.ShouldUpdateBigInt
func (u *PriceUpdater) shouldUpdatePrice(lastPrice, newPrice *big.Int) bool {
	return calc.ShouldUpdateBigInt(newPrice, lastPrice, u.priceThreshold)
}

// UpdateTokenList updates the list of tokens to monitor
func (u *PriceUpdater) UpdateTokenList(tokenIDs []uint16) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.tokenIDs = tokenIDs
	u.log.WithField("token_ids", tokenIDs).Info("Updated token list")
}

// GetTokenList returns current token list
func (u *PriceUpdater) GetTokenList() []uint16 {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.tokenIDs
}

// GetLastPrice returns the last updated price for a token
func (u *PriceUpdater) GetLastPrice(tokenID uint16) *big.Int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	if price, exists := u.lastPrices[tokenID]; exists {
		return new(big.Int).Set(price)
	}
	return nil
}

// GetAllLastPrices returns a copy of all cached prices
func (u *PriceUpdater) GetAllLastPrices() map[uint16]*big.Int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	result := make(map[uint16]*big.Int)
	for tokenID, price := range u.lastPrices {
		result[tokenID] = new(big.Int).Set(price)
	}
	return result
}
