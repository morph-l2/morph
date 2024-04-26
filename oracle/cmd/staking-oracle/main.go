package main

import (
	"fmt"
	"github.com/morph-l2/morph/oracle/flags"
	"github.com/morph-l2/morph/oracle/oracle"

	"os"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/params"
	"github.com/urfave/cli"
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
