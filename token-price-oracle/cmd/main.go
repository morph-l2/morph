package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"
	"morph-l2/token-price-oracle/client"
	"morph-l2/token-price-oracle/config"
	"morph-l2/token-price-oracle/flags"
	"morph-l2/token-price-oracle/metrics"
	"morph-l2/token-price-oracle/updater"
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
	app.Name = "token-price-oracle"
	app.Usage = "Gas Price Oracle Service"
	app.Description = "Service for monitoring L1 gas prices and updating L2 GasPriceOracle contract"
	app.Action = Main

	if err := app.Run(os.Args); err != nil {
		log.Crit("Application failed", "err", err)
	}
}

func Main(cliCtx *cli.Context) error {
	// Load configuration
	cfg, err := config.LoadConfig(cliCtx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Setup logging
	var logHandler log.Handler

	output := io.Writer(os.Stderr)
	if cfg.LogFilename != "" {
		dir := filepath.Dir(cfg.LogFilename) // handles "dir/filename" correctly
		if dir != "" && dir != "." {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("create log directory %q failed: %v", dir, err)
			}
		}
		f, err := os.OpenFile(cfg.LogFilename, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
		if err != nil {
			return fmt.Errorf("wrong log.filename set: %d", err)
		}
		f.Close()

		if cfg.LogFileMaxSize < 1 {
			return fmt.Errorf("wrong log.maxsize set: %d", cfg.LogFileMaxSize)
		}

		if cfg.LogFileMaxAge < 1 {
			return fmt.Errorf("wrong log.maxage set: %d", cfg.LogFileMaxAge)
		}
		logFile := &lumberjack.Logger{
			Filename: cfg.LogFilename,
			MaxSize:  cfg.LogFileMaxSize, // megabytes
			MaxAge:   cfg.LogFileMaxAge,  // days
			Compress: cfg.LogCompress,
		}
		output = io.MultiWriter(output, logFile)
	}
	if cfg.LogTerminal {
		logHandler = log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	} else {
		logHandler = log.StreamHandler(output, log.JSONFormat())
	}
	logLevel, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return err
	}
	log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize metrics if enabled
	if cfg.MetricsServerEnable {
		go func() {
			if err := metrics.StartMetricsServer(cfg.MetricAddress()); err != nil {
				log.Error("Metrics server failed", "err", err)
			}
		}()
		log.Info("Metrics server started", "address", cfg.MetricAddress())
	}

	// Create L2 client
	l2Client, err := client.NewL2Client(cfg.L2RPC, cfg.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to create L2 client: %w", err)
	}
	defer l2Client.Close()

	// Create transaction manager
	txManager := updater.CreateTxManager(l2Client)
	log.Info("Transaction manager initialized")

	priceUpdater, err := updater.CreatePriceUpdater(cfg, l2Client, txManager)
	if err != nil {
		return fmt.Errorf("failed to create price updater: %w", err)
	}

	if priceUpdater == nil {
		log.Warn("Price updater not created (no token IDs configured)")
	} else {
		log.Info("Price updater created", "updater", "price")
		if err := priceUpdater.Start(ctx); err != nil {
			return fmt.Errorf("failed to start updater: %w", err)
		}
	}

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigCh:
		log.Info("Received interrupt signal, shutting down...")
	case <-ctx.Done():
		log.Info("Context cancelled, shutting down...")
	}

	// Graceful shutdown
	cancel()

	if priceUpdater != nil {
		if err := priceUpdater.Stop(); err != nil {
			log.Warn("Failed to stop updater", "error", err)
		}
	}

	time.Sleep(2 * time.Second)

	log.Info("Token price Oracle stopped")
	return nil
}
