package main

import (
	"github.com/urfave/cli/v2"
)

var pushCommand = &cli.Command{
	Name:      "push",
	Usage:     "Push an image or a repository to a registry",
	ArgsUsage: "[<reference>]",
	Action:    pushImage,
}

func pushImage(ctx *cli.Context) error {
	panic("not implemented")
}
