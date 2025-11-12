package updater

import (
	"fmt"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/morph/bindings/bindings"
	"github.com/morph-l2/morph/gas-price-oracle/client"
	"github.com/morph-l2/morph/gas-price-oracle/config"
	"github.com/sirupsen/logrus"
)

// CreateUpdaters creates all enabled updaters based on config
func CreateUpdaters(
	cfg *config.Config,
	l1Client *client.L1Client,
	l2Client *client.L2Client,
	beaconClient *client.BeaconClient,
	oracleContract *bindings.GasPriceOracle,
	rollupContract *bindings.Rollup,
	txManager *TxManager,
) ([]Updater, error) {
	var updaters []Updater

	// Base fee updater (optional)
	if cfg.BaseFeeUpdateEnabled {
		baseFeeUpdater := NewBaseFeeUpdater(
			l1Client,
			l2Client,
			oracleContract,
			txManager,
			cfg.GasThreshold,
			cfg.Interval,
		)
		updaters = append(updaters, baseFeeUpdater)
		logrus.WithField("updater", "basefee").Info("Base fee updater created")
	} else {
		logrus.Warn("Base fee updater disabled")
	}

	// Scalar updater (optional)
	if cfg.ScalarUpdateEnabled {
		scalarUpdater := NewScalarUpdater(
			l1Client,
			l2Client,
			beaconClient,
			oracleContract,
			rollupContract,
			txManager,
			cfg.GasThreshold,
			cfg.OverheadInterval,
			cfg.TxnPerBatch,
		)
		updaters = append(updaters, scalarUpdater)
		logrus.WithField("updater", "scalar").Info("Scalar updater created")
	} else {
		logrus.Warn("Scalar updater disabled")
	}

	// Price updater (optional)
	if cfg.PriceUpdateEnabled {
		priceUpdater, err := createPriceUpdater(cfg, l2Client, txManager)
		if err != nil {
			return nil, fmt.Errorf("failed to create price updater: %w", err)
		}
		if priceUpdater != nil {
			updaters = append(updaters, priceUpdater)
			logrus.WithField("updater", "price").Info("Price updater created")
		}
	} else {
		logrus.Info("Price updater disabled")
	}

	if len(updaters) == 0 {
		logrus.Warn("No updaters enabled!")
	}

	return updaters, nil
}

// createPriceUpdater creates price updater if conditions are met
func createPriceUpdater(
	cfg *config.Config,
	l2Client *client.L2Client,
	txManager *TxManager,
) (*PriceUpdater, error) {
	if cfg.L2TokenRegistryAddr == (common.Address{}) {
		return nil, fmt.Errorf("price update enabled but token registry address not set")
	}

	if len(cfg.TokenIDs) == 0 {
		logrus.Warn("Price update enabled but no token IDs specified, skipping")
		return nil, nil
	}

	// Create registry contract
	registryContract, err := bindings.NewL2TokenRegistry(cfg.L2TokenRegistryAddr, l2Client.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to create TokenRegistry contract: %w", err)
	}
	logrus.WithField("address", cfg.L2TokenRegistryAddr.Hex()).Info("TokenRegistry contract bound")

	// Create price feed
	priceFeed, err := createPriceFeed(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create price feed: %w", err)
	}

	// Create price updater
	priceUpdater := NewPriceUpdater(
		l2Client,
		registryContract,
		priceFeed,
		txManager,
		cfg.TokenIDs,
		cfg.PriceUpdateInterval,
		cfg.PriceThreshold,
	)

	logrus.WithFields(logrus.Fields{
		"token_ids": cfg.TokenIDs,
		"interval":  cfg.PriceUpdateInterval,
		"threshold": cfg.PriceThreshold,
	}).Info("Price updater configured")

	return priceUpdater, nil
}

// createPriceFeed creates appropriate price feed based on config
func createPriceFeed(cfg *config.Config) (client.PriceFeed, error) {
	switch cfg.PriceFeedType {
	case "mock":
		feed := client.NewMockPriceFeed(cfg.BasePrice, cfg.PriceVariation)
		logrus.WithFields(logrus.Fields{
			"type":       "mock",
			"base_price": cfg.BasePrice.String(),
			"variation":  cfg.PriceVariation,
		}).Info("Mock price feed created")
		return feed, nil

	case "bitget":
		if len(cfg.TokenMapping) == 0 {
			return nil, fmt.Errorf("bitget price feed requires token mapping")
		}
		feed := client.NewBitgetSDKPriceFeed(cfg.TokenMapping)
		logrus.WithFields(logrus.Fields{
			"type":    "bitget-sdk",
			"mapping": cfg.TokenMapping,
		}).Info("Bitget SDK price feed created")
		return feed, nil

	default:
		return nil, fmt.Errorf("unsupported price feed type: %s", cfg.PriceFeedType)
	}
}

// CreateTxManager creates transaction manager
func CreateTxManager(l2Client *client.L2Client) *TxManager {
	return NewTxManager(l2Client)
}

// BindContracts binds all required contracts
func BindContracts(
	cfg *config.Config,
	l1Client *client.L1Client,
	l2Client *client.L2Client,
) (*bindings.GasPriceOracle, *bindings.Rollup, error) {
	// Bind GasPriceOracle contract
	oracleContract, err := bindings.NewGasPriceOracle(cfg.L2GasPriceOracleAddr, l2Client.GetClient())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to bind GasPriceOracle contract: %w", err)
	}
	logrus.WithField("address", cfg.L2GasPriceOracleAddr.Hex()).Info("GasPriceOracle contract bound")

	// Bind Rollup contract
	rollupContract, err := bindings.NewRollup(cfg.L1RollupAddress, l1Client.GetClient())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to bind Rollup contract: %w", err)
	}
	logrus.WithField("address", cfg.L1RollupAddress.Hex()).Info("Rollup contract bound")

	return oracleContract, rollupContract, nil
}
