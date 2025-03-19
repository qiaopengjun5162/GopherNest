package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/version"
	"github.com/urfave/cli/v2" // https://cli.urfave.org/v2/getting-started/
)

// Semantic holds the textual version string for major.minor.patch.
var Semantic = fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)

// WithMeta holds the textual version string including the metadata.
var WithMeta = func() string {
	v := Semantic
	if version.Meta != "" {
		v += "-" + version.Meta
	}
	return v
}()

func withCommit(gitCommit, gitDate string) string {
	vsn := WithMeta
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	if (version.Meta != "stable") && (gitDate != "") {
		vsn += "-" + gitDate
	}
	return vsn
}

func NewCli(GitCommit string, gitDate string) *cli.App {
	return &cli.App{
		Name:                 "GopherNest",
		Usage:                "GopherNest",
		Description:          "GopherNest",
		Version:              withCommit(GitCommit, gitDate),
		EnableBashCompletion: true, // Boolean to enable bash completion commands
		Commands: []*cli.Command{
			{
				Name:        "version",
				Usage:       "Show project version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}

type interruptContextKeyType struct{}

var blockerContextKey = interruptContextKeyType{}

type interruptCatcher struct {
	incoming chan os.Signal
}

func (c *interruptCatcher) Block(ctx context.Context) {
	select {
	case <-c.incoming:
	case <-ctx.Done():
	}
}

var DefaultInterruptSignals = []os.Signal{
	os.Interrupt,
	os.Kill,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

type BlockFn func(ctx context.Context)

func WithInterruptBlocker(ctx context.Context) context.Context {
	if ctx.Value(blockerContextKey) != nil { // already has an interrupt handler
		return ctx
	}
	catcher := &interruptCatcher{
		incoming: make(chan os.Signal, 10),
	}
	signal.Notify(catcher.incoming, DefaultInterruptSignals...)

	return context.WithValue(ctx, blockerContextKey, BlockFn(catcher.Block))
}
