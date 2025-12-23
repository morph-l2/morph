package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/morph-l2/go-ethereum/log"
	"morph-l2/token-price-oracle/metrics"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	// Setup basic logging
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
	log.Info("Starting token-price-oracle (metrics only mode)", "version", fmt.Sprintf("%s-%s-%s", GitVersion, GitCommit, GitDate))

	// Metrics server address
	metricsAddr := "0.0.0.0:6060"

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
}
