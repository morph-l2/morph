package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/morph-l2/morph/gas-price-oracle/client"
	"github.com/morph-l2/morph/gas-price-oracle/config"
	"github.com/morph-l2/morph/gas-price-oracle/flags"
	"github.com/morph-l2/morph/gas-price-oracle/metrics"
	"github.com/morph-l2/morph/gas-price-oracle/updater"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", GitVersion, GitCommit, GitDate)
	app.Name = "gas-price-oracle"
	app.Usage = "Gas Price Oracle Service"
	app.Description = "Service for monitoring L1 gas prices and updating L2 GasPriceOracle contract"
	app.Action = Main

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("Application failed")
	}
}

func Main(cliCtx *cli.Context) error {
	// Load configuration
	cfg, err := config.LoadConfig(cliCtx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Setup logging
	if err := setupLogging(cfg); err != nil {
		return fmt.Errorf("failed to setup logging: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"version":           GitVersion,
		"commit":            GitCommit,
		"date":              GitDate,
		"l1_rpc":            cfg.L1RPC,
		"l2_rpc":            cfg.L2RPC,
		"l1_beacon_rpc":     cfg.L1BeaconRPC,
		"l1_rollup":         cfg.L1RollupAddress.Hex(),
		"l2_oracle":         cfg.L2GasPriceOracleAddr.Hex(),
		"l2_token_registry": cfg.L2TokenRegistryAddr.Hex(),
		"gas_threshold":     cfg.GasThreshold,
		"interval":          cfg.Interval,
		"overhead_interval": cfg.OverheadInterval,
		"txn_per_batch":     cfg.TxnPerBatch,
		"basefee_enabled":   cfg.BaseFeeUpdateEnabled,
		"scalar_enabled":    cfg.ScalarUpdateEnabled,
		"price_enabled":     cfg.PriceUpdateEnabled,
	}).Info("Starting Gas Price Oracle")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize metrics if enabled
	if cfg.MetricsServerEnable {
		go func() {
			if err := metrics.StartMetricsServer(cfg.MetricAddress()); err != nil {
				logrus.WithError(err).Error("Metrics server failed")
			}
		}()
		logrus.WithField("address", cfg.MetricAddress()).Info("Metrics server started")
	}

	// Initialize clients
	l1Client, l2Client, beaconClient, err := initializeClients(cfg)
	if err != nil {
		return err
	}
	defer l1Client.Close()
	defer l2Client.Close()

	// Bind contracts
	oracleContract, rollupContract, err := updater.BindContracts(cfg, l1Client, l2Client)
	if err != nil {
		return err
	}

	// Create transaction manager
	txManager := updater.CreateTxManager(l2Client)
	logrus.Info("Transaction manager initialized")

	// Create all updaters (basefee, scalar, price)
	updaters, err := updater.CreateUpdaters(
		cfg,
		l1Client,
		l2Client,
		beaconClient,
		oracleContract,
		rollupContract,
		txManager,
	)
	if err != nil {
		return fmt.Errorf("failed to create updaters: %w", err)
	}

	// Start all updaters
	for _, upd := range updaters {
		if err := upd.Start(ctx); err != nil {
			return fmt.Errorf("failed to start %s updater: %w", upd.Name(), err)
		}
		logrus.WithField("updater", upd.Name()).Info("Updater started")
	}

	logrus.WithField("count", len(updaters)).Info("All updaters started successfully")

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigCh:
		logrus.Info("Received interrupt signal, shutting down...")
	case <-ctx.Done():
		logrus.Info("Context cancelled, shutting down...")
	}

	// Graceful shutdown
	cancel()

	// Stop all updaters
	for _, upd := range updaters {
		if err := upd.Stop(); err != nil {
			logrus.WithFields(logrus.Fields{
				"updater": upd.Name(),
				"error":   err,
			}).Warn("Failed to stop updater")
		} else {
			logrus.WithField("updater", upd.Name()).Info("Updater stopped")
		}
	}

	time.Sleep(2 * time.Second)

	logrus.Info("Gas Price Oracle stopped")
	return nil
}

// initializeClients creates and connects all RPC clients
func initializeClients(cfg *config.Config) (*client.L1Client, *client.L2Client, *client.BeaconClient, error) {
	// Create L1 client
	l1Client, err := client.NewL1Client(cfg.L1RPC)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to connect to L1: %w", err)
	}
	logrus.Info("L1 client connected")

	// Create L2 client
	l2Client, err := client.NewL2Client(cfg.L2RPC, cfg.PrivateKey)
	if err != nil {
		l1Client.Close()
		return nil, nil, nil, fmt.Errorf("failed to create L2 client: %w", err)
	}
	logrus.WithField("address", l2Client.GetAuth().From.Hex()).Info("L2 client initialized")

	// Create beacon client
	beaconClient := client.NewBeaconClient(cfg.L1BeaconRPC)
	logrus.Info("Beacon client initialized")

	return l1Client, l2Client, beaconClient, nil
}

func setupLogging(cfg *config.Config) error {
	// Parse log level
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}
	logrus.SetLevel(level)

	// Set formatter
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Setup file logging if configured
	if cfg.LogFilename != "" {
		logFile := &lumberjack.Logger{
			Filename:   cfg.LogFilename,
			MaxSize:    cfg.LogFileMaxSize,
			MaxAge:     cfg.LogFileMaxAge,
			MaxBackups: 10,
			Compress:   cfg.LogCompress,
		}

		// Use multi-writer to write to both stdout and file
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		logrus.SetOutput(multiWriter)

		logrus.WithField("filename", cfg.LogFilename).Info("File logging enabled")
	}

	return nil
}
