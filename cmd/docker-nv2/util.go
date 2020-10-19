package main

import (
	"os"
	"os/exec"

	"github.com/sajayantony/nv2-demo/pkg/config"
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
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			os.Exit(err.ExitCode())
		}
		return err
	}
	os.Exit(0)
	panic("process should be terminated")
}
