package main

import (
	"fmt"
	"os"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"

	"morph-l2/tx-submitter/flags"
)

func main() {
	// Set up logger with a default INFO level
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		),
	)

	app := cli.NewApp()
	app.Name = "cancel-tx"
	app.Usage = "Cancel a transaction"
	app.Flags = flags.CancleTxFlags

	app.Action = func(ctx *cli.Context) error {
		txHash := ctx.String("tx_hash")
		fmt.Printf("Cancelling transaction with hash: %s\n", txHash)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
