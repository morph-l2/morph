package tx_summitter

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/metrics"
	"morph-l2/tx-submitter/services"
	"morph-l2/tx-submitter/utils"
)

// Main is the entrypoint into the batch submitter service. This method returns
// a closure that executes the service and blocks until the service exits. The
// use of a closure allows the parameters bound to the top-level main package,
// e.g. GitVersion, to be captured and used once the function is executed.
func Main() func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := utils.NewConfig(cliCtx)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Set up our logging. If Sentry is enabled, we will use our custom log
		// handler that logs to stdout and forwards any error messages to Sentry
		// for collection. Otherwise, logs will only be posted to stdout.
		output := io.Writer(os.Stdout)
		if cfg.LogFilename != "" {
			f, err := os.OpenFile(cfg.LogFilename, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
			if err != nil {
				return fmt.Errorf("wrong log.filename set: %d", err)
			}
			_ = f.Close()

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

		logHandler := log.StreamHandler(output, log.TerminalFormat(false))

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
		l1Rollup, err := bindings.NewRollup(*rollupAddr, l1Client)
		if err != nil {
			return err
		}
		m := metrics.NewMetrics()
		abi, _ := bindings.RollupMetaData.GetAbi()

		// l1 staking
		l1Staking, err := bindings.NewL1Staking(common.HexToAddress(cfg.L1StakingAddress), l1Client)
		if err != nil {
			return fmt.Errorf("failed to connect to l1 staking contract")
		}

		sr := services.NewRollup(
			ctx,
			m,
			l1Client,
			l2Clients,
			l1Rollup,
			l1Staking,
			chainID,
			privKey,
			*rollupAddr,
			abi,
			cfg,
		)
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
			}
			log.Info("metrics server enabled", "host", cfg.MetricsHostname, "port", cfg.MetricsPort)
		}

		log.Info("starting tx submitter",
			"l1_rpc", cfg.L1EthRpc,
			"l2_rpcs", cfg.L2EthRpcs,
			"rollup_addr", rollupAddr.Hex(),
			"chainid", chainID.String(),
			"submitter_addr", cfg.SubmitterAddress,
			"sequencer_addr", cfg.SequencerAddress,
			"finalize_enable", cfg.Finalize,
			"priority_rollup_enable", cfg.PriorityRollup,
		)
		sr.Start()

		return nil
	}
}
