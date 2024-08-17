package main

import (
	"os"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"

	"morph-l2/morph-deployer/cmd/genesis"
	"morph-l2/morph-deployer/flags"
	oplog "morph-l2/morph-deployer/log"
	"morph-l2/morph-deployer/version"
)

var (
	GitCommit = ""
	GitDate   = ""
)

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	v := version.Version
	if GitCommit != "" {
		v += "-" + GitCommit[:8]
	}
	if GitDate != "" {
		v += "-" + GitDate
	}
	if version.Meta != "" {
		v += "-" + version.Meta
	}
	return v
}()

func main() {
	// Set up logger with a default INFO level in case we fail to parse flags,
	// otherwise the final critical log won't show what the parsing error was.
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Version = VersionWithMeta
	app.Flags = flags.Flags
	app.Name = "op-node"
	app.Usage = "Optimism Rollup Node"
	app.Description = "The Optimism Rollup Node derives L2 block inputs from L1 data and drives an external L2 Execution Engine to build a L2 chain."
	//app.Action = RollupNodeMain
	app.Commands = []cli.Command{
		{
			Name:        "genesis",
			Subcommands: genesis.Subcommands,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
