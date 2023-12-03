package main

import (
	"log"
	"os"
	"tisea-backend/cli"

	cliv2 "github.com/urfave/cli/v2"
)

func main() {
	app := &cliv2.App{
		Name: "tisea",
		Usage: "Central management and entrance of the backend part of Tisea.",
		Action: func(ctx *cliv2.Context) error {
			if ctx.NArg() == 0 {
				return cliv2.ShowAppHelp(ctx)
			}
			return nil
		},
		Commands: cli.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}