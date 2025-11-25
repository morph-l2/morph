package updater

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
	"morph-l2/bindings/bindings"
	"morph-l2/token-price-oracle/client"
	"morph-l2/token-price-oracle/metrics"
)

// PriceUpdater handles token price updates
type PriceUpdater struct {
	l2Client         *client.L2Client
	registryContract *bindings.L2TokenRegistry
	priceFeed        client.PriceFeed
	txManager        *TxManager
	tokenMapping     map[uint16]string // tokenID -> trading pair (e.g. 1 -> "BTCUSDT")
	interval         time.Duration
	priceThreshold   uint64
	stopChan         chan struct{}
	stopOnce         sync.Once // ensures stopChan is closed only once
}

// NewPriceUpdater creates a new price updater
func NewPriceUpdater(
	l2Client *client.L2Client,
	registryContract *bindings.L2TokenRegistry,
	priceFeed client.PriceFeed,
	txManager *TxManager,
	tokenMapping map[uint16]string,
	interval time.Duration,
	priceThreshold uint64,
) *PriceUpdater {
	return &PriceUpdater{
		l2Client:         l2Client,
		registryContract: registryContract,
		priceFeed:        priceFeed,
		txManager:        txManager,
		tokenMapping:     tokenMapping,
		interval:         interval,
		priceThreshold:   priceThreshold,
		stopChan:         make(chan struct{}),
	}
}

// Start starts the price updater
func (u *PriceUpdater) Start(ctx context.Context) error {
	go func() {
		log.Info("Price updater starting", "interval", u.interval)
		ticker := time.NewTicker(u.interval)
		defer ticker.Stop()

		if err := u.update(ctx); err != nil {
			log.Error("Initial price update failed")
		}

		for {
			select {
			case <-ctx.Done():
				log.Info("Price updater stopped by context")
				return
			case <-u.stopChan:
				log.Info("Price updater stopped")
				return
			case <-ticker.C:
				if err := u.update(ctx); err != nil {
					log.Error("Failed to update prices")
					metrics.UpdateErrors.WithLabelValues("price").Inc()
				}
			}
		}
	}()
	return nil
}

// Stop gracefully stops the updater
// This method is safe to call multiple times
func (u *PriceUpdater) Stop() error {
	u.stopOnce.Do(func() {
		close(u.stopChan)
		log.Info("Price updater stop requested")
	})
	return nil
}

// update performs one price update
func (u *PriceUpdater) update(ctx context.Context) error {
	defer func() {
		if err := u.updateBalanceMetrics(ctx); err != nil {
			log.Warn("Failed to update balance metrics", "error", err)
		}
	}()

	// Fetch token IDs from contract if not configured
	tokenIDs, err := u.fetchTokenIDsFromContract(ctx)
	if err != nil {
		log.Error("Failed to fetch tokenIDs from contract, price updater will not start")
		return err
	}

	if len(tokenIDs) == 0 {
		log.Warn("No tokens to update, skipping price update cycle")
		return nil
	}

	// Step 1: Fetch new prices from feed (USD prices)
	tokenPrices, err := u.priceFeed.GetBatchTokenPrices(ctx, tokenIDs)
	if err != nil {
		return fmt.Errorf("failed to fetch token prices: %w", err)
	}

	// Step 2: Calculate price ratios using tokenInfo from contract
	newPriceRatios := make(map[uint16]*big.Int)
	for tokenID, tokenPrice := range tokenPrices {
		priceRatio, err := u.calculatePriceRatio(ctx, tokenID, tokenPrice)
		if err != nil {
			log.Warn("Failed to calculate price ratio, skipping",
				"token_id", tokenID,
				"error", err)
			continue
		}
		newPriceRatios[tokenID] = priceRatio
	}

	// Step 3: Fetch current prices from contract and filter prices that need updating
	var tokenIDsToUpdate []uint16
	var pricesToUpdate []*big.Int

	callOpts := &bind.CallOpts{Context: ctx}
	for tokenID, newPrice := range newPriceRatios {
		if newPrice == nil || newPrice.Sign() == 0 {
			log.Warn("Skipping zero price",
				"token_id", tokenID)
			continue
		}

		// Fetch current price from contract (not from cache)
		lastPrice, err := u.registryContract.GetTokenPrice(callOpts, tokenID)
		if err != nil {
			log.Warn("Failed to get current price from contract, will update anyway",
				"token_id", tokenID,
				"error", err)
			tokenIDsToUpdate = append(tokenIDsToUpdate, tokenID)
			pricesToUpdate = append(pricesToUpdate, newPrice)
			continue
		}

		// Check if price changed significantly
		if lastPrice.Sign() > 0 {
			// Calculate if price change exceeds threshold
			if !u.shouldUpdatePrice(lastPrice, newPrice) {
				log.Debug("Price change below threshold, skipping update",
					"token_id", tokenID,
					"last_price", lastPrice.String(),
					"new_price", newPrice.String(),
					"threshold", u.priceThreshold)
				continue
			}

			log.Info("Price change exceeds threshold, will update",
				"token_id", tokenID,
				"last_price", lastPrice.String(),
				"new_price", newPrice.String())
		} else {
			log.Info("First time update for token (no price in contract)",
				"token_id", tokenID,
				"new_price", newPrice.String())
		}

		tokenIDsToUpdate = append(tokenIDsToUpdate, tokenID)
		pricesToUpdate = append(pricesToUpdate, newPrice)
	}

	if len(tokenIDsToUpdate) == 0 {
		log.Debug("No prices need updating (all changes below threshold)")
		return nil
	}

	log.Info("Updating token prices",
		"token_count", len(tokenIDsToUpdate),
		"token_ids", tokenIDsToUpdate,
		"total_tokens", len(tokenIDs))

	// Step 3: Update prices on L2
	receipt, err := u.txManager.SendTransaction(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return u.registryContract.BatchUpdatePrices(auth, tokenIDsToUpdate, pricesToUpdate)
	})

	if err != nil {
		log.Error("Failed to send transaction", "error", err)
		return fmt.Errorf("failed to send batch update prices transaction: %w", err)
	}

	if receipt == nil {
		log.Error("Received nil receipt")
		return fmt.Errorf("received nil receipt")
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("Transaction failed on-chain",
			"tx_hash", receipt.TxHash.Hex(),
			"status", receipt.Status,
			"gas_used", receipt.GasUsed)
		return fmt.Errorf("transaction failed on-chain: %s", receipt.TxHash.Hex())
	}

	log.Info("Successfully updated token prices",
		"tx_hash", receipt.TxHash.Hex(),
		"gas_used", receipt.GasUsed,
		"token_count", len(tokenIDsToUpdate))

	// Step 5: Update metrics
	for i, tokenID := range tokenIDsToUpdate {
		log.Debug("Price updated",
			"token_id", tokenID,
			"price_ratio", pricesToUpdate[i].String())
	}

	return nil
}

// calculatePriceRatio calculates the price ratio for a token
// Formula: priceRatio = tokenScale * tokenPriceUSD * 10^(18 - tokenDecimals) / ethPriceUSD
// We do multiplications first, then division at the end to avoid precision loss
func (u *PriceUpdater) calculatePriceRatio(ctx context.Context, tokenID uint16, tokenPrice *client.TokenPrice) (*big.Int, error) {
	// Validate input price data to prevent nil pointer panics
	if tokenPrice == nil || tokenPrice.TokenPriceUSD == nil || tokenPrice.EthPriceUSD == nil {
		return nil, fmt.Errorf("token price data missing for token %d", tokenID)
	}

	// Fetch token info from contract
	tokenInfo, err := u.registryContract.GetTokenInfo(&bind.CallOpts{
		Context: ctx,
	}, tokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to get token info from contract: %w", err)
	}

	// Check if token is active
	if !tokenInfo.IsActive {
		return nil, fmt.Errorf("token %d is not active", tokenID)
	}

	tokenScale := tokenInfo.Scale
	tokenDecimals := tokenInfo.Decimals

	log.Debug("Token info from contract",
		"token_id", tokenID,
		"address", tokenInfo.TokenAddress.Hex(),
		"decimals", tokenDecimals,
		"token_scale", tokenScale.String(),
		"active", tokenInfo.IsActive)

	// Validate token decimals (must be <= 18 for our formula to work)
	if tokenDecimals > 18 {
		return nil, fmt.Errorf("unsupported token decimals %d (>18) for token %d", tokenDecimals, tokenID)
	}

	// Check ETH price is not zero
	if tokenPrice.EthPriceUSD.Cmp(big.NewFloat(0)) == 0 {
		return nil, fmt.Errorf("ETH price is zero")
	}

	// Check token price is not zero or negative
	if tokenPrice.TokenPriceUSD.Cmp(big.NewFloat(0)) <= 0 {
		return nil, fmt.Errorf("invalid token price %s for token %d", tokenPrice.TokenPriceUSD.String(), tokenID)
	}

	// Step 1: Start with tokenPriceUSD
	priceRatio := new(big.Float).Set(tokenPrice.TokenPriceUSD)

	// Step 2: Multiply by tokenScale
	tokenScaleFloat := new(big.Float).SetInt(tokenScale)
	priceRatio.Mul(priceRatio, tokenScaleFloat)

	// Step 3: Multiply by 10^(18 - tokenDecimals)
	// ETH has 18 decimals, so we need to adjust for token decimals
	decimalAdjustment := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18-tokenDecimals)), nil)
	decimalAdjustmentFloat := new(big.Float).SetInt(decimalAdjustment)
	priceRatio.Mul(priceRatio, decimalAdjustmentFloat)

	// Step 4: Finally divide by ethPriceUSD
	priceRatio.Quo(priceRatio, tokenPrice.EthPriceUSD)

	// Convert to big.Int with precision check
	priceRatioInt, accuracy := priceRatio.Int(nil)
	if accuracy != big.Exact {
		log.Warn("Price ratio conversion lost precision",
			"token_id", tokenID,
			"symbol", tokenPrice.Symbol,
			"accuracy", accuracy.String(),
			"float_value", priceRatio.String(),
			"int_value", priceRatioInt.String())
	}

	log.Info("Calculated price ratio",
		"token_id", tokenID,
		"symbol", tokenPrice.Symbol,
		"token_price_usd", tokenPrice.TokenPriceUSD.String(),
		"eth_price_usd", tokenPrice.EthPriceUSD.String(),
		"decimals", tokenDecimals,
		"token_scale", tokenScale.String(),
		"price_ratio", priceRatioInt.String())

	return priceRatioInt, nil
}

// updateBalanceMetrics queries and updates balance metrics
func (u *PriceUpdater) updateBalanceMetrics(ctx context.Context) error {
	// Get account address
	account := u.l2Client.WalletAddress()

	// Query ETH balance
	ethBalance, err := u.l2Client.GetClient().BalanceAt(ctx, account, nil)
	if err != nil {
		return fmt.Errorf("failed to get ETH balance: %w", err)
	}

	// Convert to ETH (wei to ETH)
	ethBalanceFloat := new(big.Float).SetInt(ethBalance)
	ethBalanceFloat.Quo(ethBalanceFloat, big.NewFloat(1e18))
	ethBalanceEth, _ := ethBalanceFloat.Float64()

	// Update ETH balance metric
	metrics.AccountBalance.Set(ethBalanceEth)

	log.Info("Updated balance metrics",
		"account", account.Hex(),
		"eth_balance", ethBalanceEth)

	return nil
}

// shouldUpdatePrice checks if the price change exceeds the threshold
// Formula: |newPrice - lastPrice| / lastPrice * 10000 >= threshold
// Threshold is in basis points (bps): 1 bps = 0.01%, 100 bps = 1%, 10000 bps = 100%
// Example: if threshold is 100 (bps), price must change by at least 1% to trigger update
func (u *PriceUpdater) shouldUpdatePrice(lastPrice, newPrice *big.Int) bool {
	// Validate inputs
	if lastPrice == nil || newPrice == nil {
		log.Warn("shouldUpdatePrice called with nil price")
		return false
	}

	if lastPrice.Sign() == 0 {
		return true // Always update if no previous price
	}

	// Validate threshold is reasonable (should be <= 10000 for basis points)
	// If threshold is unreasonably large, log warning and cap at 100% (10000 bps)
	threshold := u.priceThreshold
	if threshold > 10000 {
		log.Warn("Price threshold is unusually large, capping at 100% (10000 bps)",
			"configured_threshold", threshold,
			"capped_threshold", 10000)
		threshold = 10000
	}

	// Calculate absolute difference: |newPrice - lastPrice|
	diff := new(big.Int).Sub(newPrice, lastPrice)
	diff.Abs(diff)

	// Calculate change in basis points: diff * 10000 / lastPrice
	// This gives us the change in bps (e.g., 100 for 1%, 10 for 0.1%, 1 for 0.01%)
	bps := new(big.Int).Mul(diff, big.NewInt(10000))
	bps.Div(bps, lastPrice)

	// Compare with threshold (both are in basis points)
	thresholdBig := big.NewInt(int64(threshold))
	shouldUpdate := bps.Cmp(thresholdBig) >= 0
	
	log.Debug("Price change check",
		"last_price", lastPrice.String(),
		"new_price", newPrice.String(),
		"change_bps", bps.String(),
		"threshold_bps", threshold,
		"should_update", shouldUpdate)
	
	return shouldUpdate
}

// fetchTokenIDsFromContract fetches supported token IDs from L2TokenRegistry contract
func (u *PriceUpdater) fetchTokenIDsFromContract(ctx context.Context) ([]uint16, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	// Call getSupportedIDList() on the contract
	return u.registryContract.GetSupportedIDList(callOpts)
}
