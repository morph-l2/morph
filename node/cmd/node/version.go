package main

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli"
)

// Version information, set via -ldflags
var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildTime = "unknown"
)

var versionCmd = cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "show version information",
	Action: func(ctx *cli.Context) error {
		fmt.Printf("morphnode %s\n", Version)
		fmt.Printf("Git Commit: %s\n", GitCommit)
		fmt.Printf("Build Time: %s\n", BuildTime)
		fmt.Printf("Go Version: %s\n", runtime.Version())
		fmt.Printf("OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
		return nil
	},
}
