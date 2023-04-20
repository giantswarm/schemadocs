package generate

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

const (
	flagReadme              = "readme"
	flagDocPlaceholderStart = "doc-placeholder-start"
	flagDocPlaceholderEnd   = "doc-placeholder-end"
)

type flag struct {
	readme              string
	docPlaceholderStart string
	docPlaceholderEnd   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.readme, flagReadme, "", "Path to file to output the generated documentation")
	cmd.Flags().StringVar(&f.docPlaceholderStart, flagDocPlaceholderStart, "", "Placeholder string marking the start of the docs section in the output file")
	cmd.Flags().StringVar(&f.docPlaceholderEnd, flagDocPlaceholderEnd, "", "Placeholder string marking the end of the docs section in the output file")
}

func (f *flag) Validate() error {
	if (f.docPlaceholderStart == "") != (f.docPlaceholderEnd == "") {
		return microerror.Maskf(invalidFlagError, "both --% and --S flags must be set to non-empty values", flagDocPlaceholderStart, flagDocPlaceholderEnd)
	}
	return nil
}
