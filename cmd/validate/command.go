package validate

import (
	"io"
	"os"

	"github.com/giantswarm/schemadocs/cmd/global"
	"github.com/spf13/cobra"
)

const (
	name            = "validate"
	description     = "Validate generated JSON schema documentation in a text file"
	longDescription = `Validate documentation in the given text bile by comparing it
to the documentation generated from the provided JSON schema.

Use --schema to specify the JSON schema.

The input text file needs to contain placeholders indicating the start and end of the documentation,
Default placeholders are <!-- DOCS_START --> and <!-- DOCS_END -->.
Use --doc-placeholder-start and --doc-placeholder-end to specify different placeholders.
`
	example = `  schemadocs validate README.md --schema schema.json
  schemadocs validate README.md --schema schema.json --doc-placeholder-start [DOCS_START] --doc-placeholder-end [DOCS_END]
`
)

type Config struct {
	GlobalFlag *global.Flag
	Stderr     io.Writer
	Stdout     io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	f := &flag{}

	r := &runner{
		flag:       f,
		globalFlag: config.GlobalFlag,
		stderr:     config.Stderr,
		stdout:     config.Stdout,
	}

	c := &cobra.Command{
		Use:           name,
		Short:         description,
		Long:          longDescription,
		Example:       example,
		RunE:          r.Run,
		Args:          cobra.MinimumNArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	f.Init(c)

	return c, nil
}
