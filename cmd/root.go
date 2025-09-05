package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemadocs/cmd/generate"
	"github.com/giantswarm/schemadocs/cmd/global"
	"github.com/giantswarm/schemadocs/cmd/validate"
	"github.com/giantswarm/schemadocs/pkg/project"
)

const (
	name        = "schemadocs"
	description = "Genera and validate Markdown docs from JSON schema files in text files"
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

	gf := &global.Flag{}

	r := &runner{
		stdout: config.Stdout,
		stderr: config.Stderr,
	}

	c := &cobra.Command{
		Use:          name,
		Short:        description,
		RunE:         r.Run,
		SilenceUsage: true,
		//SilenceErrors: true,
		Version: project.Version(),
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	gf.Init(c)

	var generateCmd *cobra.Command
	{
		cfg := generate.Config{
			GlobalFlag: gf,
			Stderr:     config.Stderr,
			Stdout:     config.Stdout,
		}

		generateCmd, err = generate.New(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create generate command: %w", err)
		}
	}

	var validateCmd *cobra.Command
	{
		cfg := validate.Config{
			GlobalFlag: gf,
			Stderr:     config.Stderr,
			Stdout:     config.Stdout,
		}

		validateCmd, err = validate.New(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create validate command: %w", err)
		}
	}

	c.AddCommand(generateCmd)
	c.AddCommand(validateCmd)

	return c, nil
}
