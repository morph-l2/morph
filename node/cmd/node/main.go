package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/morph-l2/go-ethereum/ethclient"
	tmnode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/privval"
	"github.com/urfave/cli"

	"morph-l2/bindings/bindings"
	"morph-l2/node/cmd/keyconverter"
	node "morph-l2/node/core"
	"morph-l2/node/db"
	"morph-l2/node/derivation"
	"morph-l2/node/flags"
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
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Application failed, message: ", err)
		os.Exit(1)
	}
}

func L2NodeMain(ctx *cli.Context) error {
	var (
		err      error
		executor *node.Executor
		syncer   *sync.Syncer
		ms       *mock.Sequencer
		tmNode   *tmnode.Node
		dvNode   *derivation.Derivation

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
		// launch tendermint node
		tmCfg, err := sequencer.LoadTmConfig(ctx, home)
		if err != nil {
			return err
		}
		tmVal := privval.LoadOrGenFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile())
		pubKey, _ := tmVal.GetPubKey()
		newSyncerFunc := func() (*sync.Syncer, error) { return node.NewSyncer(ctx, home, nodeConfig) }
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
			if tmNode, err = sequencer.SetupNode(tmCfg, tmVal, executor, nodeConfig.Logger); err != nil {
				return fmt.Errorf("failed to setup consensus node, error: %v", err)
			}
			if err = tmNode.Start(); err != nil {
				return fmt.Errorf("failed to start consensus node, error: %v", err)
			}
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

	return nil
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
