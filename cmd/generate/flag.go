package generate

import (
	"fmt"

	"github.com/spf13/cobra"

	cmderror "github.com/giantswarm/schemadocs/pkg/error"
)

const (
	flagOutputPath          = "output-path"
	flagDocPlaceholderStart = "doc-placeholder-start"
	flagDocPlaceholderEnd   = "doc-placeholder-end"
	flagLayout              = "layout"
)

type flag struct {
	outputPath          string
	docPlaceholderStart string
	docPlaceholderEnd   string
	layout              string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.outputPath, flagOutputPath, "o", "", "Path to file to output the generated documentation")
	cmd.Flags().StringVar(&f.docPlaceholderStart, flagDocPlaceholderStart, "", "Placeholder string marking the start of the docs section in the output file")
	cmd.Flags().StringVar(&f.docPlaceholderEnd, flagDocPlaceholderEnd, "", "Placeholder string marking the end of the docs section in the output file")
	cmd.Flags().StringVarP(&f.layout, "layout", "l", "default", "Layout of the generated documentation")
}

func (f *flag) Validate() error {
	if (f.docPlaceholderStart == "") != (f.docPlaceholderEnd == "") {
		return fmt.Errorf("both --%s and --%s flags must be set to non-empty values: %w", flagDocPlaceholderStart, flagDocPlaceholderEnd, cmderror.InvalidFlagError)
	}
	return nil
}
