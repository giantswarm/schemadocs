package cmd

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/schemadocs/cmd/generate"
	"github.com/giantswarm/schemadocs/cmd/validate"
	"github.com/giantswarm/schemadocs/pkg/project"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const (
	name        = "schemadoc"
	description = "Generator and validator of Markdown docs from JSON schema files in text files"
)

type Config struct {
	Stdout io.Writer
	Stderr io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	var err error

	f := &flag{}

	r := &runner{
		flag:   f,
		stdout: config.Stdout,
		stderr: config.Stderr,
	}

	c := &cobra.Command{
		Use:          name,
		Short:        description,
		RunE:         r.Run,
		SilenceUsage: true,
		Version:      project.Version(),
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	f.Init(c)

	var generateCmd *cobra.Command
	{
		c := generate.Config{
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		generateCmd, err = generate.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var validateCmd *cobra.Command
	{
		c := validate.Config{
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		validateCmd, err = validate.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c.AddCommand(generateCmd)
	c.AddCommand(validateCmd)

	return c, nil
}
