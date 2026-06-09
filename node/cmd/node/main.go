package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmnode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/privval"
	tmsequencer "github.com/tendermint/tendermint/sequencer"
	"github.com/urfave/cli"

	"morph-l2/bindings/bindings"
	node "morph-l2/node/core"
	"morph-l2/node/db"
	"morph-l2/node/derivation"
	"morph-l2/node/flags"
	"morph-l2/node/hakeeper"
	"morph-l2/node/l1sequencer"
	"morph-l2/node/sequencer"
	"morph-l2/node/sequencer/mock"
	"morph-l2/node/sync"
	"morph-l2/node/types"
)

func main() {
	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Name = "morphnode"
	app.Action = L2NodeMain
	app.Commands = []cli.Command{
		versionCmd,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Application failed, message: ", err)
		os.Exit(1)
	}
}

func L2NodeMain(ctx *cli.Context) error {
	var (
		err       error
		executor  *node.Executor
		syncer    *sync.Syncer
		ms        *mock.Sequencer
		tmNode    *tmnode.Node
		dvNode    *derivation.Derivation
		tracker   *l1sequencer.L1Tracker
		verifier  *l1sequencer.SequencerVerifier
		signer    l1sequencer.Signer
		haService *hakeeper.HAService

		nodeConfig = node.DefaultConfig()
	)
	isMockSequencer := ctx.GlobalBool(flags.MockEnabled.Name)

	if ctx.GlobalBool(flags.MainnetFlag.Name) && ctx.GlobalBool(flags.HoodiFlag.Name) {
		return fmt.Errorf("--%s and --%s are mutually exclusive", flags.MainnetFlag.Name, flags.HoodiFlag.Name)
	}

	if err = nodeConfig.SetCliContext(ctx); err != nil {
		return err
	}
	home, err := homeDir(ctx)
	if err != nil {
		return err
	}

	// ========== Shared L1 client ==========
	// One ethclient.Dial per process — all L1-touching components (syncer,
	// derivation, l1sequencer Tracker/Verifier/Signer, rollup binding) share
	// the same connection pool, retry policy, and metrics surface. Adding a
	// new consumer means injecting this client, not opening a new one.
	l1RPC := ctx.GlobalString(flags.L1NodeAddr.Name)
	if l1RPC == "" {
		return fmt.Errorf("%s is required", flags.L1NodeAddr.Name)
	}
	l1Client, err := ethclient.Dial(l1RPC)
	if err != nil {
		return fmt.Errorf("dial l1 node error: %v", err)
	}

	// ========== Shared store + syncer (used by both executor and derivation) ==========
	dbConfig := db.DefaultConfig()
	dbConfig.SetCliContext(ctx)
	store, err := db.NewStore(dbConfig, home)
	if err != nil {
		return err
	}
	syncConfig := sync.DefaultConfig()
	if err = syncConfig.SetCliContext(ctx); err != nil {
		return err
	}
	syncer, err = sync.NewSyncer(context.Background(), store, syncConfig, nodeConfig.Logger, l1Client)
	if err != nil {
		return fmt.Errorf("failed to create syncer, error: %v", err)
	}

	tracker, verifier, signer, err = initL1SequencerComponents(ctx, l1Client, nodeConfig.Logger)
	if err != nil {
		return fmt.Errorf("failed to init L1 sequencer components: %w", err)
	}

	// ========== Executor + sequencer / mock ==========
	tmCfg, err := sequencer.LoadTmConfig(ctx, home)
	if err != nil {
		return err
	}
	tmVal := privval.LoadOrGenFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile())
	pubKey, _ := tmVal.GetPubKey()

	// Reuse the shared syncer instance -- DevSequencer mode is the only path
	// that pulls a syncer out of NewExecutor, so we hand back the same one
	// rather than letting NewExecutor open a second store + syncer.
	newSyncerFunc := func() (*sync.Syncer, error) { return syncer, nil }
	executor, err = node.NewExecutor(newSyncerFunc, nodeConfig, pubKey)
	if err != nil {
		return err
	}

	// Eagerly start the L1 message syncer for post-upgrade sequencer nodes that
	// are NOT in the PBFT validator set (separated-deployment / HA cluster).
	// In the combined-deployment case, updateSequencerSet already started the
	// syncer inside NewExecutor, so SetSyncer is a no-op there.
	if signer != nil && executor.Syncer() == nil {
		executor.SetSyncer(syncer)
		syncer.Start()
		nodeConfig.Logger.Info("L1 syncer start", "reason", "post-upgrade always start")
	}

	haService, err = initHAService(ctx, home, nodeConfig.Logger)
	if err != nil {
		return err
	}

	// ========== Derivation config (loaded early to drive the layer1 branch below) ==========
	derivationCfg := derivation.DefaultConfig()
	if err := derivationCfg.SetCliContext(ctx); err != nil {
		return fmt.Errorf("derivation set cli context error: %v", err)
	}

	switch {
	case isMockSequencer:
		ms, err = mock.NewSequencer(executor)
		if err != nil {
			return err
		}
		go ms.Start()
	case derivationCfg.VerifyMode == derivation.VerifyModeLayer1:
		nodeConfig.Logger.Info("layer1 verify mode: tendermint not started")
		// Other modes inherit /metrics from tendermint; layer1 has to bring
		// its own listener up on the same address.
		startMetricsServer(tmCfg.Instrumentation.PrometheusListenAddr, nodeConfig.Logger)
	default:
		// Convert typed nil (*HAService)(nil) to untyped nil interface to avoid
		// Go's nil interface gotcha: a typed nil satisfies (ha != nil) checks.
		var ha tmsequencer.SequencerHA
		if haService != nil {
			ha = haService
		}
		tmNode, err = sequencer.SetupNode(tmCfg, tmVal, executor, nodeConfig.Logger, verifier, signer, ha)
		if err != nil {
			return fmt.Errorf("failed to setup consensus node: %v", err)
		}
		if err = tmNode.Start(); err != nil {
			return fmt.Errorf("failed to start consensus node, error: %v", err)
		}
	}

	// ========== Derivation ==========
	// Only fullnodes (no signer) derive blocks from L1. Sequencer nodes
	// (HA leader/follower or single-sequencer) produce blocks locally and
	// replicate via Raft / broadcast — running derivation on a sequencer
	// is both redundant (it would re-fetch L1 batches it produced) and
	// unsafe (deriveForce would risk a self-reorg on transient divergence).
	if signer == nil {
		rollup, err := bindings.NewRollup(derivationCfg.RollupContractAddress, l1Client)
		if err != nil {
			return fmt.Errorf("NewRollup error: %v", err)
		}
		dvNode, err = derivation.NewDerivationClient(context.Background(), derivationCfg, syncer, store, rollup, l1Client, tmNode, verifier, nodeConfig.Logger)
		if err != nil {
			return fmt.Errorf("new derivation client error: %v", err)
		}
		dvNode.Start()
		nodeConfig.Logger.Info("derivation started")
	} else {
		nodeConfig.Logger.Info("derivation skipped: sequencer mode")
	}

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}...)
	<-interruptChannel

	if ms != nil {
		ms.Stop()
	}
	if tmNode != nil {
		if stopErr := tmNode.Stop(); stopErr != nil {
			nodeConfig.Logger.Error("failed to stop tendermint node", "err", stopErr)
			return stopErr
		}
	}
	if syncer != nil {
		syncer.Stop()
	}
	if dvNode != nil {
		dvNode.Stop()
	}
	if tracker != nil {
		tracker.Stop()
	}
	if verifier != nil {
		verifier.Stop()
	}

	return nil
}

// initHAService builds the HA config and creates the HAService.
// Loading order: defaults → config file → flag overrides → auto-resolve → validate.
// Returns nil (no error) if HA is not enabled.
func initHAService(ctx *cli.Context, home string, logger tmlog.Logger) (*hakeeper.HAService, error) {
	cfg := hakeeper.DefaultConfig()

	if cfgPath := ctx.GlobalString(flags.SequencerHAConfig.Name); cfgPath != "" {
		if err := cfg.LoadFile(cfgPath); err != nil {
			return nil, fmt.Errorf("HA config: %w", err)
		}
	}

	if ctx.GlobalBool(flags.SequencerHAEnabled.Name) {
		cfg.Enabled = true
	}
	if ctx.GlobalBool(flags.SequencerHABootstrap.Name) {
		cfg.Bootstrap = true
	}
	if addrs := ctx.GlobalStringSlice(flags.SequencerHAJoin.Name); len(addrs) > 0 {
		cfg.JoinAddrs = addrs
	}
	if id := ctx.GlobalString(flags.SequencerHAServerID.Name); id != "" {
		cfg.ServerID = id
	}
	if addr := ctx.GlobalString(flags.SequencerHAAdvertisedAddr.Name); addr != "" {
		cfg.Consensus.AdvertisedAddr = addr
	}
	if token := ctx.GlobalString(flags.SequencerHARPCToken.Name); token != "" {
		cfg.RPC.Token = token
	}

	if !cfg.Enabled {
		return nil, nil
	}

	// Propagate node log level to Raft internal logger
	if logLevel := ctx.GlobalString(flags.LogLevel.Name); logLevel == "debug" {
		cfg.Debug = true
	}

	if err := cfg.Resolve(home); err != nil {
		return nil, fmt.Errorf("HA config resolve: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("HA config: %w", err)
	}

	cfg.LogEffectiveConfig(logger)
	return hakeeper.New(cfg, logger.With("module", "hakeeper"))
}

// initL1SequencerComponents initializes all L1 sequencer related components:
// - L1Tracker: monitors L1 sync status
// - SequencerCache: caches L1 sequencer address (nil if contract not configured)
// - Signer: signs blocks (nil if private key not configured)
func initL1SequencerComponents(
	ctx *cli.Context,
	l1Client *ethclient.Client,
	logger tmlog.Logger,
) (*l1sequencer.L1Tracker, *l1sequencer.SequencerVerifier, l1sequencer.Signer, error) {
	if l1Client == nil {
		return nil, nil, nil, fmt.Errorf("L1 client is required, check l1.rpc configuration")
	}

	// Get config from flags
	lagThreshold := ctx.GlobalDuration(flags.L1SyncLagThreshold.Name)
	if lagThreshold == 0 {
		lagThreshold = 5 * time.Minute // default
	}
	var contractAddr common.Address
	if ctx.GlobalIsSet(flags.L1SequencerContractAddr.Name) {
		contractAddr = common.HexToAddress(ctx.GlobalString(flags.L1SequencerContractAddr.Name))
	} else if ctx.GlobalBool(flags.MainnetFlag.Name) {
		contractAddr = types.MainnetL1SequencerContractAddress
	} else if ctx.GlobalBool(flags.HoodiFlag.Name) {
		contractAddr = types.HoodiL1SequencerContractAddress
	}
	seqPrivKeyHex := ctx.GlobalString(flags.SequencerPrivateKey.Name)
	enclaveSignerAddr := ctx.GlobalString(flags.SequencerEnclaveSignerAddr.Name)

	// Initialize L1 Tracker
	tracker := l1sequencer.NewL1Tracker(context.Background(), l1Client, lagThreshold, logger)
	if err := tracker.Start(); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to start L1 tracker: %w", err)
	}
	logger.Info("L1 Tracker started", "lagThreshold", lagThreshold)

	// Initialize Sequencer Verifier
	var verifier *l1sequencer.SequencerVerifier
	if contractAddr != (common.Address{}) {
		caller, err := bindings.NewL1SequencerCaller(contractAddr, l1Client)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to create L1Sequencer caller: %w", err)
		}
		verifier = l1sequencer.NewSequencerVerifier(caller, logger)
		logger.Info("Sequencer verifier initialized", "contract", contractAddr.Hex())
	} else {
		return nil, nil, nil, fmt.Errorf("L1 Sequencer contract address is required, check l1.sequencerContract configuration")
	}

	// Initialize Signer (optional). Three mutually exclusive modes:
	//   1) sequencer.privateKey set        → LocalSigner (plaintext key in node memory)
	//   2) sequencer.enclaveSignerAddr     → EnclaveSigner (vsock to Nitro Enclave; key never in node)
	//   3) neither                         → no signer, sequencer can't produce blocks
	// Both at once is rejected so config mistakes surface at boot, not at first block.
	var signer l1sequencer.Signer
	switch {
	case seqPrivKeyHex != "" && enclaveSignerAddr != "":
		return nil, nil, nil, fmt.Errorf(
			"sequencer.privateKey and sequencer.enclaveSignerAddr are mutually exclusive; pick one")
	case seqPrivKeyHex != "":
		seqPrivKeyHex = strings.TrimPrefix(seqPrivKeyHex, "0x")
		privKey, err := crypto.HexToECDSA(seqPrivKeyHex)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("invalid sequencer private key: %w", err)
		}
		signer, err = l1sequencer.NewLocalSigner(privKey, logger)
		if err != nil {
			return nil, nil, nil, err
		}
		logger.Info("Sequencer signer initialized (local)", "address", signer.Address().Hex())
	case enclaveSignerAddr != "":
		var err error
		signer, err = l1sequencer.NewEnclaveSigner(enclaveSignerAddr, logger)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("init enclave signer: %w", err)
		}
		logger.Info("Sequencer signer initialized (enclave)",
			"addr", enclaveSignerAddr, "address", signer.Address().Hex())
	default:
		logger.Info("Sequencer signer not configured (no privateKey, no enclaveSignerAddr)")
	}

	return tracker, verifier, signer, nil
}

func startMetricsServer(addr string, logger tmlog.Logger) {
	if addr == "" {
		logger.Info("metrics server disabled (instrumentation.prometheus_listen_addr is empty)")
		return
	}
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 10 * time.Second}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("metrics server", "addr", addr, "err", err)
		}
	}()
	logger.Info("metrics server started", "addr", addr)
}

func homeDir(ctx *cli.Context) (string, error) {
	home := ctx.GlobalString(flags.Home.Name)
	if home == "" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		home = filepath.Join(userHome, types.DefaultHomeDir)
	}
	return home, nil
}
