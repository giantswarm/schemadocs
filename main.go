package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemadocs/cmd"
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
			return fmt.Errorf("failed to create root command: %w", err)
		}
	}

	err = rootCommand.Execute()
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}
	return nil
}
