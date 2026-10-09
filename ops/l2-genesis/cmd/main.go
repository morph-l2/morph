package main

import (
	"os"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"

	"morph-l2/morph-deployer/cmd/genesis"
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
		commit := GitCommit
		if len(commit) > 8 {
			commit = commit[:8]
		}
		v += "-" + commit
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
	app.Flags = oplog.CLIFlags("MORPH_GENESIS")
	app.Name = "morph-genesis"
	app.Usage = "Generate Morph L2 genesis artifacts"
	app.Description = "Create L2 genesis and rollup configuration from confirmed L1 deployment records."
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
