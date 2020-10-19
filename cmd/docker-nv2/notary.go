package main

import (
	"github.com/urfave/cli/v2"
)

var notaryCommand = &cli.Command{
	Name: "notary",
	Subcommands: []*cli.Command{
		notarySignCommand,
	},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "enabled",
			Usage: "Enable Notary support",
		},
	},
	Action: setNotary,
}

func setNotary(ctx *cli.Context) error {
	panic("not implemented")
}
