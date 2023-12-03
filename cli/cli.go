package cli

import (
	// "os"
	// "os/signal"
	// "syscall"

	cliv2 "github.com/urfave/cli/v2"
)

var Commands = []*cliv2.Command{
	{
		Name: "help",
		Usage: "Get help messages",
		Action: func(ctx *cliv2.Context) error {
			return cliv2.ShowAppHelp(ctx)
		},
	},
	{
		Name: "run",
		Usage: "Run the backend",
		Action: func(ctx *cliv2.Context) error {
			return RunBackend()
		},
	},
}