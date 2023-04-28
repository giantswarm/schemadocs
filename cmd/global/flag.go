package global

import "github.com/spf13/cobra"

const (
	flagNoColor = "no-color"
)

type Flag struct {
	NoColor bool
}

func (f *Flag) Init(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(&f.NoColor, flagNoColor, false, "Ignore colors in the output")
}

func (f *Flag) Validate() error {
	return nil
}
