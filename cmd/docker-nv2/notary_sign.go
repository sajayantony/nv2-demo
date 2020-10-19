package main

import (
	"fmt"

	"github.com/sajayantony/nv2-demo/pkg/config"
	"github.com/sajayantony/nv2-demo/pkg/crypto"
	"github.com/sajayantony/nv2-demo/pkg/docker"
	"github.com/sajayantony/nv2-demo/pkg/util"
	"github.com/urfave/cli/v2"
)

var notarySignCommand = &cli.Command{
	Name:      "sign",
	Usage:     "Sign a docker image",
	ArgsUsage: "[<reference>]",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:      "key",
			Aliases:   []string{"k"},
			Usage:     "signing key file",
			TakesFile: true,
			Required:  true,
		},
		&cli.PathFlag{
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

	service, err := crypto.GetSigningService(
		ctx.Path("key"),
		ctx.Path("cert"),
	)
	if err != nil {
		return err
	}

	reference := ctx.Args().First()
	fmt.Println("Generating Docker mainfest:", reference)
	desc, err := docker.GenerateManifestOCIDescriptor(reference)
	if err != nil {
		return err
	}

	fmt.Println("Signing", desc.Digest)
	sig, err := service.Sign(ctx.Context, desc, reference)
	if err != nil {
		return err
	}
	sigPath := config.SignaturePath(desc.Digest)
	if err := util.WriteFile(sigPath, sig); err != nil {
		return err
	}
	fmt.Println("Signature saved to", sigPath)

	return nil
}
