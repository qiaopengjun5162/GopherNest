package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/log"
)

var (
	GitCommit = ""
	gitDate   = ""
)

func main() {
	app := NewCli(GitCommit, gitDate)
	ctx := WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("Application failed")
		os.Exit(1)
	}
}
