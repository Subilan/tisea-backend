package cli

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	&cli.Command{
		Name: "help",
		Usage: "Get help messages",
		Action: func(ctx *cli.Context) error {
			return cli.ShowAppHelp(ctx)
		},
	},
	&cli.Command{
		Name: "run",
		Usage: "Run backend built on gin",
		Action: func(ctx *cli.Context) error {
			return RunBackend()
		},
	},
}