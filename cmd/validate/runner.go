package validate

import (
	"fmt"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/schemadocs/pkg/cli"
	"github.com/giantswarm/schemadocs/pkg/readme"
	"github.com/spf13/cobra"
	"io"
)

type runner struct {
	flag   *flag
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	var writeErr error

	err := r.run(cmd, args)
	if err != nil && isInvalidFlag(err) || isInvalidConfig(err) {
		writeErr = cli.WriteError(r.stderr, "Failed to validate documentation", err)
	} else if err != nil {
		writeErr = cli.WriteOutput(r.stderr, err.Error())
	}

	if writeErr != nil {
		return writeErr
	}
	return err
}

func (r *runner) run(cmd *cobra.Command, args []string) error {
	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	readmeItem, err := readme.New(args[0], r.flag.docPlaceholderStart, r.flag.docPlaceholderEnd)
	if err != nil {
		return microerror.Mask(err)
	}

	err = readmeItem.Validate(r.flag.schema)
	if err != nil {
		_, writeErr := fmt.Fprintln(r.stderr, err.Error())
		if writeErr != nil {
			return writeErr
		}
		return err
	}

	return cli.WriteOutput(r.stdout, "Documentation is valid!")
}
