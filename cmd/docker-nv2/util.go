package main

import (
	"os"
	"os/exec"

	"github.com/notaryproject/notary/v2"
	"github.com/sajayantony/nv2-demo/pkg/config"
	"github.com/sajayantony/nv2-demo/pkg/crypto"
	"github.com/urfave/cli/v2"
)

func passThroughIfNotaryDisabled(ctx *cli.Context) error {
	err := config.CheckNotaryEnabled()
	if err == nil {
		return nil
	}
	if err != config.ErrNotaryDisabled {
		return err
	}

	args := append([]string{ctx.Command.Name}, ctx.Args().Slice()...)
	if err := runCommand("docker", args...); err != nil {
		return err
	}
	os.Exit(0)
	panic("process should be terminated")
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			os.Exit(err.ExitCode())
		}
		return err
	}
	return nil
}

func getVerificationService() (notary.SigningService, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	return crypto.GetSigningService("", cfg.VerificationCerts...)
}
