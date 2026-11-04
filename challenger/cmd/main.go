package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"

	"morph-l2/challenger/client"
	"morph-l2/challenger/config"
	"morph-l2/challenger/flags"
	"morph-l2/challenger/httpapi"
	"morph-l2/challenger/metrics"
	"morph-l2/challenger/service"
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
	app.Name = "challenger"
	app.Usage = "Submit Rollup.challengeState when an external caller triggers it"
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Crit("Application failed", "err", err)
	}
}

func run(cliCtx *cli.Context) error {
	cfg, err := config.LoadConfig(cliCtx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	if err := setupLogger(cfg); err != nil {
		return err
	}

	if cfg.HTTPAuthToken == "" {
		log.Warn("HTTP auth token is empty; POST /challenge is open")
	}

	if cfg.MetricsServerEnable {
		go func() {
			if err := metrics.StartServer(cfg.MetricAddress()); err != nil {
				log.Error("Metrics server failed", "err", err)
			}
		}()
	}

	l1Client, err := client.NewL1Client(cfg)
	if err != nil {
		return fmt.Errorf("failed to create L1 client: %w", err)
	}
	defer l1Client.Close()

	challenger, err := service.New(l1Client, cfg.RollupAddress, cfg.ReceiptTimeout)
	if err != nil {
		return fmt.Errorf("failed to create challenger: %w", err)
	}

	server := httpapi.NewServer(cfg.HTTPAddress(), cfg.HTTPAuthToken, challenger, cfg.ReceiptTimeout)
	errCh := make(chan error, 1)
	go func() {
		log.Info("Challenge HTTP server started", "address", cfg.HTTPAddress())
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errCh:
		return err
	case <-sigCh:
		log.Info("Received interrupt signal, shutting down")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}

func setupLogger(cfg *config.Config) error {
	output := io.Writer(os.Stderr)
	if cfg.LogFilename != "" {
		dir := filepath.Dir(cfg.LogFilename)
		if dir != "" && dir != "." {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("create log directory %q failed: %w", dir, err)
			}
		}
		file, err := os.OpenFile(cfg.LogFilename, os.O_CREATE|os.O_RDWR, 0o600)
		if err != nil {
			return fmt.Errorf("open log file: %w", err)
		}
		file.Close()
		if cfg.LogFileMaxSize < 1 {
			return fmt.Errorf("log max size must be positive")
		}
		if cfg.LogFileMaxAge < 1 {
			return fmt.Errorf("log max age must be positive")
		}
		output = io.MultiWriter(output, &lumberjack.Logger{
			Filename: cfg.LogFilename,
			MaxSize:  cfg.LogFileMaxSize,
			MaxAge:   cfg.LogFileMaxAge,
			Compress: cfg.LogCompress,
		})
	}

	var handler log.Handler
	if cfg.LogTerminal {
		handler = log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	} else {
		handler = log.StreamHandler(output, log.JSONFormat())
	}
	level, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return err
	}
	log.Root().SetHandler(log.LvlFilterHandler(level, handler))
	return nil
}
