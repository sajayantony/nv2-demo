package main

import (
	"github.com/urfave/cli/v2"
)

var pullCommand = &cli.Command{
	Name:      "pull",
	Usage:     "Pull an image or a repository from a registry",
	ArgsUsage: "[<reference>]",
	Action:    pullImage,
}

func pullImage(ctx *cli.Context) error {
	if err := passThroughIfNotaryDisabled(ctx); err != nil {
		return err
	}

	panic("not implemented")
}
