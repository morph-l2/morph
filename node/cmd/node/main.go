package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmnode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/privval"
	"github.com/urfave/cli"

	"morph-l2/bindings/bindings"
	"morph-l2/node/blocktag"
	"morph-l2/node/cmd/keyconverter"
	node "morph-l2/node/core"
	"morph-l2/node/db"
	"morph-l2/node/derivation"
	"morph-l2/node/flags"
	"morph-l2/node/l1sequencer"
	"morph-l2/node/sequencer"
	"morph-l2/node/sequencer/mock"
	"morph-l2/node/sync"
	"morph-l2/node/types"
	"morph-l2/node/validator"
)

var keyConverterCmd = cli.Command{
	Name:    "key-converter",
	Aliases: []string{"kc"},
	Usage:   "tools to convert base64-encoded keys(tendermint key/bls key) to the format used by contracts",
	Action:  keyconverter.ConvertKey,
	Flags:   keyconverter.Flags,
}

func main() {
	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Name = "morphnode"
	app.Action = L2NodeMain
	app.Commands = []cli.Command{
		keyConverterCmd,
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
		err         error
		executor    *node.Executor
		syncer      *sync.Syncer
		ms          *mock.Sequencer
		tmNode      *tmnode.Node
		dvNode      *derivation.Derivation
		blockTagSvc *blocktag.BlockTagService
		tracker  *l1sequencer.L1Tracker
		verifier *l1sequencer.SequencerVerifier
		signer   l1sequencer.Signer

		nodeConfig = node.DefaultConfig()
	)
	isMockSequencer := ctx.GlobalBool(flags.MockEnabled.Name)
	isValidator := ctx.GlobalBool(flags.ValidatorEnable.Name)

	if err = nodeConfig.SetCliContext(ctx); err != nil {
		return err
	}
	home, err := homeDir(ctx)
	if err != nil {
		return err
	}

	if isValidator {
		// configure store
		dbConfig := db.DefaultConfig()
		dbConfig.SetCliContext(ctx)
		store, err := db.NewStore(dbConfig, home)
		if err != nil {
			return err
		}
		derivationCfg := derivation.DefaultConfig()
		if err := derivationCfg.SetCliContext(ctx); err != nil {
			return fmt.Errorf("derivation set cli context error: %v", err)
		}
		syncConfig := sync.DefaultConfig()
		if err = syncConfig.SetCliContext(ctx); err != nil {
			return err
		}
		syncer, err = sync.NewSyncer(context.Background(), store, syncConfig, nodeConfig.Logger)
		if err != nil {
			return fmt.Errorf("failed to create syncer, error: %v", err)
		}
		validatorCfg := validator.NewConfig()
		if err := validatorCfg.SetCliContext(ctx); err != nil {
			return fmt.Errorf("validator set cli context error: %v", err)
		}
		l1Client, err := ethclient.Dial(derivationCfg.L1.Addr)
		if err != nil {
			return fmt.Errorf("dial l1 node error:%v", err)
		}
		rollup, err := bindings.NewRollup(derivationCfg.RollupContractAddress, l1Client)
		if err != nil {
			return fmt.Errorf("NewRollup error:%v", err)
		}
		vt, err := validator.NewValidator(validatorCfg, rollup, nodeConfig.Logger)
		if err != nil {
			return fmt.Errorf("new validator client error: %v", err)
		}

		dvNode, err = derivation.NewDerivationClient(context.Background(), derivationCfg, syncer, store, vt, rollup, nodeConfig.Logger)
		if err != nil {
			return fmt.Errorf("new derivation client error: %v", err)
		}
		dvNode.Start()
		nodeConfig.Logger.Info("derivation node starting")
	} else {
		// ========== Create Syncer and L1 Sequencer Components ==========
		syncer, err = node.NewSyncer(ctx, home, nodeConfig)
		if err != nil {
			return fmt.Errorf("failed to create syncer: %w", err)
		}

		tracker, verifier, signer, err = initL1SequencerComponents(ctx, syncer.L1Client(), nodeConfig.Logger)
		if err != nil {
			return fmt.Errorf("failed to init L1 sequencer components: %w", err)
		}

		// ========== Launch Tendermint Node ==========
		tmCfg, err := sequencer.LoadTmConfig(ctx, home)
		if err != nil {
			return err
		}
		tmVal := privval.LoadOrGenFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile())
		pubKey, _ := tmVal.GetPubKey()

		// Create executor with syncer
		newSyncerFunc := func() (*sync.Syncer, error) { return syncer, nil } // Reuse existing syncer
		executor, err = node.NewExecutor(newSyncerFunc, nodeConfig, pubKey)
		if err != nil {
			return err
		}
		if isMockSequencer {
			ms, err = mock.NewSequencer(executor)
			if err != nil {
				return err
			}
			go ms.Start()
		} else {
			tmNode, err = sequencer.SetupNode(tmCfg, tmVal, executor, nodeConfig.Logger, verifier, signer)
			if err != nil {
				return fmt.Errorf("failed to setup consensus node: %v", err)
			}
			if err = tmNode.Start(); err != nil {
				return fmt.Errorf("failed to start consensus node, error: %v", err)
			}
		}

		// ========== Initialize BlockTagService ==========
		blockTagSvc, err = initBlockTagService(ctx, syncer.L1Client(), executor, nodeConfig.Logger)
		if err != nil {
			return fmt.Errorf("failed to init BlockTagService: %w", err)
		}
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
	if blockTagSvc != nil {
		blockTagSvc.Stop()
	}
	if tracker != nil {
		tracker.Stop()
	}

	return nil
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
	contractAddr := common.HexToAddress(ctx.GlobalString(flags.L1SequencerContractAddr.Name))
	seqPrivKeyHex := ctx.GlobalString(flags.SequencerPrivateKey.Name)

	// Initialize L1 Tracker
	tracker := l1sequencer.NewL1Tracker(context.Background(), l1Client, lagThreshold, logger)
	if err := tracker.Start(); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to start L1 tracker: %w", err)
	}
	logger.Info("L1 Tracker started", "lagThreshold", lagThreshold)

	// Initialize Sequencer Verifier (optional)
	var verifier *l1sequencer.SequencerVerifier
	if contractAddr != (common.Address{}) {
		caller, err := bindings.NewL1SequencerCaller(contractAddr, l1Client)
		if err != nil {
			tracker.Stop()
			return nil, nil, nil, fmt.Errorf("failed to create L1Sequencer caller: %w", err)
		}
		verifier = l1sequencer.NewSequencerVerifier(caller, logger)
		logger.Info("Sequencer verifier initialized", "contract", contractAddr.Hex())
	} else {
		logger.Info("L1 Sequencer contract not configured, verifier disabled")
	}

	// Initialize Signer (optional)
	var signer l1sequencer.Signer
	if seqPrivKeyHex != "" {
		seqPrivKeyHex = strings.TrimPrefix(seqPrivKeyHex, "0x")
		privKey, err := crypto.HexToECDSA(seqPrivKeyHex)
		if err != nil {
			tracker.Stop()
			return nil, nil, nil, fmt.Errorf("invalid sequencer private key: %w", err)
		}
		signer, err = l1sequencer.NewLocalSigner(privKey, verifier, logger)
		if err != nil {
			tracker.Stop()
			return nil, nil, nil, err
		}
		logger.Info("Sequencer signer initialized", "address", signer.Address().Hex())
	} else {
		logger.Info("Sequencer private key not configured, signer disabled")
	}

	return tracker, verifier, signer, nil
}

// initBlockTagService initializes the block tag service
func initBlockTagService(
	ctx *cli.Context,
	l1Client *ethclient.Client,
	executor *node.Executor,
	logger tmlog.Logger,
) (*blocktag.BlockTagService, error) {
	config := blocktag.DefaultConfig()
	if err := config.SetCliContext(ctx); err != nil {
		return nil, err
	}

	svc, err := blocktag.NewBlockTagService(context.Background(), l1Client, executor.L2Client(), config, logger)
	if err != nil {
		return nil, err
	}

	if err := svc.Start(); err != nil {
		return nil, err
	}

	logger.Info("BlockTagService started")
	return svc, nil
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
