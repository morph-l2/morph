package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	tmnode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/privval"
	"github.com/urfave/cli"

	"morph-l2/node/blocktag"
	"morph-l2/node/cmd/keyconverter"
	node "morph-l2/node/core"
	"morph-l2/node/derivation"
	"morph-l2/node/flags"
	"morph-l2/node/sequencer"
	"morph-l2/node/sequencer/mock"
	"morph-l2/node/sync"
	"morph-l2/node/types"
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
		ms          *mock.Sequencer
		tmNode      *tmnode.Node
		blockTagSvc *blocktag.BlockTagService

		nodeConfig = node.DefaultConfig()
	)
	isMockSequencer := ctx.GlobalBool(flags.MockEnabled.Name)

	if err = nodeConfig.SetCliContext(ctx); err != nil {
		return err
	}
	home, err := homeDir(ctx)
	if err != nil {
		return err
	}

	// launch tendermint node
	tmCfg, err := sequencer.LoadTmConfig(ctx, home)
	if err != nil {
		return err
	}
	tmVal := privval.LoadOrGenFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile())
	pubKey, err := tmVal.GetPubKey()
	if err != nil {
		return fmt.Errorf("failed to get validator public key: %w", err)
	}
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

	// Start BlockTagService
	blockTagConfig := blocktag.DefaultConfig()
	if err := blockTagConfig.SetCliContext(ctx); err != nil {
		return fmt.Errorf("blocktag config set cli context error: %w", err)
	}

	// Build BatchVerifier for full batch validation.
	// It reuses the same L1 addr / rollup address / L2 eth addr already parsed above.
	bvCfg := &derivation.Config{
		L1:                    &types.L1Config{Addr: blockTagConfig.L1Addr},
		L2:                    &types.L2Config{EthAddr: nodeConfig.L2.EthAddr},
		RollupContractAddress: blockTagConfig.RollupAddress,
		BeaconRpc:             ctx.GlobalString(flags.L1BeaconAddr.Name),
		BaseHeight:            ctx.GlobalUint64(flags.DerivationBaseHeight.Name),
	}
	bv, bvErr := derivation.NewBatchVerifier(context.Background(), bvCfg, nil, nodeConfig.Logger)
	if bvErr != nil {
		// BatchVerifier is non-critical; fall back to lightweight state-root-only check
		nodeConfig.Logger.Error("failed to create BatchVerifier, falling back to state-root-only validation", "error", bvErr)
		bv = nil
	}

	blockTagSvc, err = blocktag.NewBlockTagService(context.Background(), executor.L2Client(), blockTagConfig, bv, nodeConfig.Logger)
	if err != nil {
		return fmt.Errorf("failed to create BlockTagService: %w", err)
	}
	if err := blockTagSvc.Start(); err != nil {
		return fmt.Errorf("failed to start BlockTagService: %w", err)
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
	if blockTagSvc != nil {
		blockTagSvc.Stop()
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
