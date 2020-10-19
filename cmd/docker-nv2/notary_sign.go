package main

import (
	"github.com/urfave/cli/v2"
)

var notarySignCommand = &cli.Command{
	Name:      "sign",
	Usage:     "Sign a docker image",
	ArgsUsage: "[<reference>]",
	Action:    notarySign,
}

func notarySign(ctx *cli.Context) error {
	panic("not implemented")
}
