package main

import (
	"os"
	"path/filepath"

	cmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	"github.com/tendermint/tendermint/cmd/tendermint/commands/debug"
	"github.com/tendermint/tendermint/libs/cli"
	nm "github.com/tendermint/tendermint/node"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.GenValidatorCmd,
		cmd.InitFilesCmd,
		cmd.ProbeUpnpCmd,
		cmd.LightCmd,
		cmd.ReplayCmd,
		cmd.ReplayConsoleCmd,
		cmd.ResetAllCmd,
		cmd.ResetPrivValidatorCmd,
		cmd.ResetStateCmd,
		cmd.ShowValidatorCmd,
		cmd.TestnetFilesCmd,
		cmd.ShowNodeIDCmd,
		cmd.GenNodeKeyCmd,
		cmd.VersionCmd,
		cmd.RollbackStateCmd,
		cmd.RewindCmd,
		cmd.CompactGoLevelDBCmd,
		debug.DebugCmd,
		cli.NewCompletionCmd(rootCmd, true),
	)
	rootCmd.AddCommand(cmd.NewRunNodeCmd(nm.DefaultNewNode))

	command := cli.PrepareBaseCmd(rootCmd, "TM", os.ExpandEnv(filepath.Join("$HOME", ".morphnode")))
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
