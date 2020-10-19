package main

import (
	"github.com/sajayantony/nv2-demo/pkg/config"
	"github.com/urfave/cli/v2"
)

var notarySignCommand = &cli.Command{
	Name:      "sign",
	Usage:     "Sign a docker image",
	ArgsUsage: "[<reference>]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:      "key",
			Aliases:   []string{"k"},
			Usage:     "signing key file",
			TakesFile: true,
			Required:  true,
		},
		&cli.StringFlag{
			Name:      "cert",
			Aliases:   []string{"c"},
			Usage:     "signing cert",
			TakesFile: true,
		},
	},
	Action: notarySign,
}

func notarySign(ctx *cli.Context) error {
	if err := config.CheckNotaryEnabled(); err != nil {
		return err
	}

	panic("not implemented")
}
