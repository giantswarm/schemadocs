package generate

import (
	"github.com/giantswarm/schemadocs/cmd/global"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const (
	name            = "generate"
	description     = "Generate documentation from JSON schema"
	longDescription = `Generate documentation from the given JSON schema input and store it in a README file

By default the documentation will be stored to README.md file in the current working directory.
Use --output-path / -o to specify a target path.

The output file needs to contain placeholders indicating the start and end of the documentation,
Default placeholders are {::comment} # DOCS_START {/:comment} and {::comment} # DOCS_END {/:comment}.
Use --doc-placeholder-start and --doc-placeholder-end to specify different placeholders.
`
	example = `  schemadocs generate schema.json
  schemadocs generate schema.json -o README.md
  schemadocs generate schema.json --doc-placeholder-start [DOCS_START] --doc-placeholder-end [DOCS_END]
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
		ArgAliases:    []string{"PATH"},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	f.Init(c)

	return c, nil
}
