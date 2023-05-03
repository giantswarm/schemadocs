package validate

import (
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/schemadocs/cmd/global"
	"github.com/giantswarm/schemadocs/pkg/cli"
	cmderror "github.com/giantswarm/schemadocs/pkg/error"
	"github.com/giantswarm/schemadocs/pkg/readme"
	"github.com/spf13/cobra"
)

type runner struct {
	flag       *flag
	globalFlag *global.Flag
	stdout     io.Writer
	stderr     io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	cli.WriteOutput(r.stdout, "Validating documentation")

	err := r.run(cmd, args)

	cli.WriteNewLine(r.stdout)

	if err != nil && cmderror.IsInvalidDocs(err) {
		cli.WriteFailure(r.stderr, r.globalFlag.NoColor, "Documentation is not valid", err)
	} else if err != nil {
		cli.WriteError(r.stderr, r.globalFlag.NoColor, "Failed to validate documentation", err)
	} else {
		cli.WriteSuccess(r.stdout, r.globalFlag.NoColor, "Documentation is valid!\n")
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

	cli.WriteOutputF(r.stdout, "Source file: %s\n", args[0])
	if err != nil {
		return err
	}

	readmeItem, err := readme.New(args[0], r.flag.docPlaceholderStart, r.flag.docPlaceholderEnd)
	if err != nil {
		return microerror.Mask(err)
	}

	cli.WriteOutputF(r.stdout, "Schema file: %s\n", r.flag.schema)

	if r.flag.docPlaceholderStart != "" && r.flag.docPlaceholderEnd != "" {
		cli.WriteOutputF(r.stdout, "Using placeholders '%s' and '%s' from flags\n", readmeItem.StartPlaceholder(), readmeItem.EndPlaceholder())
	} else {
		cli.WriteOutputF(r.stdout, "Using default placeholders '%s' and '%s'\n", readmeItem.StartPlaceholder(), readmeItem.EndPlaceholder())
	}

	err = readmeItem.Validate(r.flag.schema)
	if err != nil {
		return err
	}

	return nil
}
