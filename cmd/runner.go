package cmd

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	"io"
)

type runner struct {
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	err := cmd.Help()
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
