package generate

import (
	"github.com/giantswarm/microerror"
	cmderror "github.com/giantswarm/schemadocs/pkg/error"
	"github.com/spf13/cobra"
)

const (
	flagOutputPath          = "output-path"
	flagDocPlaceholderStart = "doc-placeholder-start"
	flagDocPlaceholderEnd   = "doc-placeholder-end"
)

type flag struct {
	outputPath          string
	docPlaceholderStart string
	docPlaceholderEnd   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.outputPath, flagOutputPath, "o", "", "Path to file to output the generated documentation")
	cmd.Flags().StringVar(&f.docPlaceholderStart, flagDocPlaceholderStart, "", "Placeholder string marking the start of the docs section in the output file")
	cmd.Flags().StringVar(&f.docPlaceholderEnd, flagDocPlaceholderEnd, "", "Placeholder string marking the end of the docs section in the output file")
}

func (f *flag) Validate() error {
	if (f.docPlaceholderStart == "") != (f.docPlaceholderEnd == "") {
		return microerror.Maskf(cmderror.InvalidFlagError, "both --%s and --%s flags must be set to non-empty values", flagDocPlaceholderStart, flagDocPlaceholderEnd)
	}
	return nil
}
