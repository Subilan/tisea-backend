package main

import (
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "tisea",
		Usage: "Central management and entrance of the backend part of Tisea.",
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() == 0 {
				return cli.ShowAppHelp(ctx)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}