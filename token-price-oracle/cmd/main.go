package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"
	"morph-l2/token-price-oracle/flags"
	"morph-l2/token-price-oracle/metrics"
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
	// Setup basic logging
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
	log.Info("Starting token-price-oracle (metrics only mode)")

	// Get metrics config
	metricsHost := cliCtx.String(flags.MetricsHostnameFlag.Name)
	metricsPort := cliCtx.Uint64(flags.MetricsPortFlag.Name)
	metricsAddr := fmt.Sprintf("%s:%d", metricsHost, metricsPort)

	// Start metrics server
	go func() {
		if err := metrics.StartMetricsServer(metricsAddr); err != nil {
			log.Error("Metrics server failed", "err", err)
		}
	}()
	log.Info("Metrics server started", "address", metricsAddr)

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	log.Info("Received interrupt signal, shutting down...")

	time.Sleep(1 * time.Second)
	log.Info("Token price Oracle stopped")
	return nil
}
