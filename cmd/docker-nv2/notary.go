package main

import (
	"os"

	"github.com/sajayantony/nv2-demo/pkg/config"
	"github.com/urfave/cli/v2"
)

var notaryCommand = &cli.Command{
	Name: "notary",
	Subcommands: []*cli.Command{
		notarySignCommand,
	},
	Flags: []cli.Flag{
		notaryEnabledFlag,
	},
	Action: setNotary,
}

var notaryEnabledFlag = &cli.BoolFlag{
	Name:  "enabled",
	Usage: "Enable Notary support",
}

func setNotary(ctx *cli.Context) error {
	if !ctx.IsSet(notaryEnabledFlag.Name) {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}

	conf, err := config.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		conf = config.New()
	}
	conf.Enabled = ctx.Bool(notaryEnabledFlag.Name)
	return conf.Save()
}
