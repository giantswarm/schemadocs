package generate

import (
	"fmt"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/schemadocs/pkg/cli"
	"github.com/giantswarm/schemadocs/pkg/generate"
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
	err := r.run(cmd, args)
	if err != nil {
		writeErr := cli.WriteError(r.stderr, "Failed to generate documentation", err)
		if writeErr != nil {
			return writeErr
		}
	}
	return err
}

func (r *runner) run(cmd *cobra.Command, args []string) error {
	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	docs, err := generate.Generate(args[0])
	if err != nil {
		return microerror.Mask(err)
	}

	readmeItem, err := readme.New(r.flag.readme, r.flag.docPlaceholderStart, r.flag.docPlaceholderEnd)
	if err != nil {
		return microerror.Mask(err)
	}

	err = readmeItem.WriteDocs(docs)
	if err != nil {
		return microerror.Mask(err)
	}

	return cli.WriteOutput(r.stdout, fmt.Sprintf("Successfully generated documentation from %s and stored it in %s", args[0], r.flag.readme))
}
