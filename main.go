package main

import (
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/schemadocs/cmd"
	"github.com/spf13/cobra"
)

func main() {
	if err := mainE(); err != nil {
		os.Exit(1)
	}
}

func mainE() error {
	var err error

	var rootCommand *cobra.Command
	{
		c := cmd.Config{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		rootCommand, err = cmd.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = rootCommand.Execute()
	if err != nil {
		return microerror.Mask(err)
	}
	return nil
}
