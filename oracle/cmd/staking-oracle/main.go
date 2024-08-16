package main

import (
	"fmt"
	"os"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/morph-l2/go-ethereum/params"
	"github.com/urfave/cli"

	"morph-l2/oracle/flags"
	"morph-l2/oracle/oracle"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s", GitVersion, params.VersionWithCommit(GitCommit, GitDate))
	app.Name = "staking-oracle"
	app.Usage = "Staking oracle Service"

	app.Action = oracle.Main()
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
