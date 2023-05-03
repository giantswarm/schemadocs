package generate

import (
	"io"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/schemadocs/cmd/global"
	"github.com/giantswarm/schemadocs/pkg/cli"
	cmderror "github.com/giantswarm/schemadocs/pkg/error"
	"github.com/giantswarm/schemadocs/pkg/generate"
	"github.com/giantswarm/schemadocs/pkg/readme"
)

type runner struct {
	flag       *flag
	globalFlag *global.Flag
	stdout     io.Writer
	stderr     io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	cli.WriteOutput(r.stdout, "Generating documentation")

	err := r.run(cmd, args)

	cli.WriteNewLine(r.stdout)

	if err != nil {
		cli.WriteError(r.stderr, r.globalFlag.NoColor, "Failed to generate documentation", err)
	} else {
		cli.WriteSuccess(r.stdout, r.globalFlag.NoColor, "Documentation generated successfully!\n")
	}

	return err
}

func (r *runner) run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return microerror.Maskf(cmderror.InvalidConfigError, "requires at least 1 arg(s), only received 0")
	}

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	cli.WriteOutputF(r.stdout, "Schema: %s\n", args[0])

	docs, err := generate.Generate(args[0])
	if err != nil {
		return microerror.Mask(err)
	}

	readmeItem, err := readme.New(r.flag.outputPath, r.flag.docPlaceholderStart, r.flag.docPlaceholderEnd)
	if err != nil {
		return microerror.Mask(err)
	}

	cli.WriteOutputF(r.stdout, "Destination file: %s\n", readmeItem.Path())

	if r.flag.docPlaceholderStart != "" && r.flag.docPlaceholderEnd != "" {
		cli.WriteOutputF(r.stdout, "Using placeholders '%s' and '%s' from flags\n", readmeItem.StartPlaceholder(), readmeItem.EndPlaceholder())
	} else {
		cli.WriteOutputF(r.stdout, "Using default placeholders '%s' and '%s'\n", readmeItem.StartPlaceholder(), readmeItem.EndPlaceholder())
	}

	err = readmeItem.WriteDocs(docs)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
