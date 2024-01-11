package tx_summitter

import (
	"context"
	"fmt"
	"os"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/tx-submitter/iface"
	"github.com/morph-l2/tx-submitter/metrics"
	"github.com/morph-l2/tx-submitter/services"
	"github.com/morph-l2/tx-submitter/utils"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/urfave/cli"
)

// Main is the entrypoint into the batch submitter service. This method returns
// a closure that executes the service and blocks until the service exits. The
// use of a closure allows the parameters bound to the top-level main package,
// e.g. GitVersion, to be captured and used once the function is executed.
func Main(gitCommit string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Set up our logging. If Sentry is enabled, we will use our custom log
		// handler that logs to stdout and forwards any error messages to Sentry
		// for collection. Otherwise, logs will only be posted to stdout.
		logHandler := log.StreamHandler(os.Stdout, log.TerminalFormat(true))

		logLevel, err := log.LvlFromString(cfg.LogLevel)
		if err != nil {
			return err
		}

		log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))

		// Parse sequencer private key and rollup contract address.

		privKey, rollupAddr, err := utils.ParsePkAndWallet(cfg.PrivateKey, cfg.RollupAddress)
		if err != nil {
			return err
		}

		// Connect to L1 and L2 providers. Perform these last since they are the
		// most expensive.
		l1Client, err := ethclient.DialContext(ctx, cfg.L1EthRpc)
		if err != nil {
			return err
		}
		// l2 rpcs
		var l2Clients []iface.L2Client
		for _, rpc := range cfg.L2EthRpcs {

			l2Client, err := ethclient.DialContext(ctx, rpc)
			if err != nil {
				log.Warn("failed to connect to L2 provider", "url", rpc)
				continue
			}
			l2Clients = append(l2Clients, l2Client)
		}
		if len(l2Clients) == 0 {
			return fmt.Errorf("cannot connect to any l2 rpc")
		}

		chainID, err := l1Client.ChainID(ctx)
		if err != nil {
			return err
		}
		rollup, err := bindings.NewRollup(*rollupAddr, l1Client)
		if err != nil {
			return err
		}
		m := metrics.NewMetrics()
		abi, _ := bindings.RollupMetaData.GetAbi()

		//
		submitterAddr := common.HexToAddress(cfg.SubmitterAddress)
		l2Submitter, err := bindings.NewSubmitter(submitterAddr, l2Clients[0])
		if err != nil {
			return err
		}

		sr := services.NewSR(ctx, l1Client, l2Clients, rollup, cfg.PollInterval, chainID, privKey, *rollupAddr, m, abi, cfg.TxTimeout, cfg.MaxBlock, cfg.MinBlock, l2Submitter, cfg.Finalize, cfg.MaxFinalizeNum, cfg.PriorityRollup)
		if err := sr.Init(); err != nil {
			return err
		}
		// metrics
		{
			if cfg.MetricsServerEnable {
				go func() {
					_, err := m.Serve(cfg.MetricsHostname, cfg.MetricsPort)
					if err != nil {
						log.Error("metrics server failed to start", "err", err)
					}
				}()
				log.Info("metrics server enabled", "host", cfg.MetricsHostname, "port", cfg.MetricsPort)
			}
		}

		sr.Start()

		return nil
	}
}
