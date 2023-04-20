package generate

import (
	"github.com/spf13/cobra"
	"io"
	"os"
)

const (
	name        = "generate"
	description = "Generate documentation from JSON schema"
)

type Config struct {
	Stderr io.Writer
	Stdout io.Writer
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
		flag:   f,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
		Args:  cobra.MinimumNArgs(1),
	}

	f.Init(c)

	return c, nil
}
