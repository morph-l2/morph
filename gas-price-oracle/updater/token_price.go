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
	"morph-l2/gas-price-oracle/client"
	"morph-l2/gas-price-oracle/metrics"
)

// PriceUpdater handles token price updates
type PriceUpdater struct {
	l2Client         *client.L2Client
	registryContract *bindings.L2TokenRegistry
	priceFeed        client.PriceFeed
	txManager        *TxManager
	tokenIDs         []uint16
	tokenMapping     map[uint16]string // tokenID -> trading pair (e.g. 1 -> "BTCUSDT")
	interval         time.Duration
	priceThreshold   uint64

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
	tokenMapping map[uint16]string,
	interval time.Duration,
	priceThreshold uint64,
) *PriceUpdater {
	return &PriceUpdater{
		l2Client:         l2Client,
		registryContract: registryContract,
		priceFeed:        priceFeed,
		txManager:        txManager,
		tokenIDs:         tokenIDs,
		tokenMapping:     tokenMapping,
		interval:         interval,
		priceThreshold:   priceThreshold,
		lastPrices:       make(map[uint16]*big.Int),
		stopChan:         make(chan struct{}),
	}
}

// Start starts the price updater
func (u *PriceUpdater) Start(ctx context.Context) error {
	go func() {
		fmt.Println("u.interval", u.interval)
		ticker := time.NewTicker(u.interval)
		defer ticker.Stop()

		// Fetch token IDs from contract if not configured
		// TODO: Uncomment when contract has getSupportedIDList method
		// if len(u.tokenIDs) == 0 {
		// 	log.Info("No tokenIDs configured, fetching from contract...")
		// 	if err := u.fetchTokenIDsFromContract(ctx); err != nil {
		// 		log.Error("Failed to fetch tokenIDs from contract, price updater will not start")
		// 		return
		// 	}
		// }

		// Filter tokenIDs to only those in tokenMapping
		u.filterTokenIDsByMapping()

		log.Info("Price updater started",
			"token_ids", u.tokenIDs,
			"token_mapping", u.tokenMapping,
			"interval", u.interval,
			"price_threshold", u.priceThreshold)

		// Perform first update immediately (will fetch current prices from contract)
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
func (u *PriceUpdater) Stop() error {
	close(u.stopChan)
	log.Info("Price updater stop requested")
	return nil
}

// update performs one price update
func (u *PriceUpdater) update(ctx context.Context) error {
	if len(u.tokenIDs) == 0 {
		log.Crit("No tokens to update")
		return nil
	}

	// Step 1: Fetch new prices from feed (USD prices)
	tokenPrices, err := u.priceFeed.GetBatchTokenPrices(ctx, u.tokenIDs)
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
		"total_tokens", len(u.tokenIDs))

	// Step 3: Update prices on L2
	receipt, err := u.txManager.SendTransaction(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return u.registryContract.BatchUpdatePrices(auth, tokenIDsToUpdate, pricesToUpdate)
	})
	if err != nil {
		return fmt.Errorf("failed to send batch update prices transaction: %w", err)
	}

	if receipt.Status == 0 {
		log.Error("Transaction failed", "tx_hash", receipt.TxHash.Hex())
		return fmt.Errorf("transaction failed on-chain: %s", receipt.TxHash.Hex())
	}

	// Step 4: Update cache with new prices
	u.mu.Lock()
	for i, tokenID := range tokenIDsToUpdate {
		u.lastPrices[tokenID] = pricesToUpdate[i]
	}
	u.mu.Unlock()

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

	metrics.ScalarUpdateCount.Inc()

	return nil
}

// calculatePriceRatio calculates the price ratio for a token
// Formula: priceRatio = tokenScale * tokenPriceUSD * 10^(18 - tokenDecimals) / ethPriceUSD
// We do multiplications first, then division at the end to avoid precision loss
func (u *PriceUpdater) calculatePriceRatio(ctx context.Context, tokenID uint16, tokenPrice *client.TokenPrice) (*big.Int, error) {
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

	// Check ETH price is not zero
	if tokenPrice.EthPriceUSD.Cmp(big.NewFloat(0)) == 0 {
		return nil, fmt.Errorf("ETH price is zero")
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

	// Convert to big.Int
	priceRatioInt, _ := priceRatio.Int(nil)

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

// shouldUpdatePrice checks if the price change exceeds the threshold
// Formula: |newPrice - lastPrice| / lastPrice * 100 >= threshold
func (u *PriceUpdater) shouldUpdatePrice(lastPrice, newPrice *big.Int) bool {
	if lastPrice.Sign() == 0 {
		return true // Always update if no previous price
	}

	// Calculate absolute difference: |newPrice - lastPrice|
	diff := new(big.Int).Sub(newPrice, lastPrice)
	diff.Abs(diff)

	// Calculate percentage change: diff * 100 / lastPrice
	percentage := new(big.Int).Mul(diff, big.NewInt(100))
	percentage.Div(percentage, lastPrice)

	// Compare with threshold
	threshold := big.NewInt(int64(u.priceThreshold))
	return percentage.Cmp(threshold) >= 0
}

// UpdateTokenList updates the list of tokens to monitor
func (u *PriceUpdater) UpdateTokenList(tokenIDs []uint16) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.tokenIDs = tokenIDs
	log.Info("Updated token list", "token_ids", tokenIDs)
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

// fetchTokenIDsFromContract fetches supported token IDs from L2TokenRegistry contract
func (u *PriceUpdater) fetchTokenIDsFromContract(ctx context.Context) error {
	callOpts := &bind.CallOpts{Context: ctx}

	// Call getSupportedIDList() on the contract
	tokenIDs, err := u.registryContract.GetSupportedIDList(callOpts)
	if err != nil {
		return fmt.Errorf("failed to call getSupportedIDList: %w", err)
	}

	if len(tokenIDs) == 0 {
		log.Warn("Contract returned empty token ID list")
		return nil
	}

	u.mu.Lock()
	u.tokenIDs = tokenIDs
	u.mu.Unlock()

	log.Info("Fetched token IDs from contract",
		"token_ids", tokenIDs,
		"count", len(tokenIDs))

	return nil
}

// filterTokenIDsByMapping filters tokenIDs to only include those that have a mapping configured
func (u *PriceUpdater) filterTokenIDsByMapping() {
	if len(u.tokenMapping) == 0 {
		log.Error("No token mapping configured for current price feed type, price updater will not work. Please configure the appropriate token-mapping flag (e.g., --token-mapping-bitget, --token-mapping-binance, --token-mapping-mock)")
		u.tokenIDs = []uint16{}
		return
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	var filtered []uint16
	var unmapped []uint16
	for _, tokenID := range u.tokenIDs {
		if _, exists := u.tokenMapping[tokenID]; exists {
			filtered = append(filtered, tokenID)
		} else {
			unmapped = append(unmapped, tokenID)
			log.Warn("Token ID not in mapping, skipping price update for this token",
				"token_id", tokenID)
		}
	}

	u.tokenIDs = filtered

	if len(unmapped) > 0 {
		log.Warn("Some token IDs from contract are not mapped to trading pairs. Please update token mapping configuration if needed.",
			"unmapped_token_ids", unmapped,
			"mapped_token_ids", filtered)
	}

	log.Info("Filtered token IDs by mapping",
		"filtered_count", len(filtered),
		"token_ids", filtered)
}
