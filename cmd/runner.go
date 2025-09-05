package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type runner struct {
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	err := cmd.Help()
	if err != nil {
		return fmt.Errorf("failed to show help: %w", err)
	}

	return nil
}
