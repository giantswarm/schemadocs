package validate

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

const (
	flagSchema              = "schema"
	flagDocPlaceholderStart = "doc-placeholder-start"
	flagDocPlaceholderEnd   = "doc-placeholder-end"
)

type flag struct {
	schema              string
	docPlaceholderStart string
	docPlaceholderEnd   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.schema, flagSchema, "", "Path to the JSON schema file")
	cmd.Flags().StringVar(&f.docPlaceholderStart, flagDocPlaceholderStart, "", "Placeholder string marking the start of the docs section in the output file")
	cmd.Flags().StringVar(&f.docPlaceholderEnd, flagDocPlaceholderEnd, "", "Placeholder string marking the end of the docs section in the output file")
}

func (f *flag) Validate() error {
	if f.schema == "" {
		return microerror.Maskf(invalidFlagError, "--%s must be set to a non-empty value", flagSchema)
	}
	if (f.docPlaceholderStart == "") != (f.docPlaceholderEnd == "") {
		return microerror.Maskf(invalidFlagError, "both --%s and --%s flags must be set to non-empty values", flagDocPlaceholderStart, flagDocPlaceholderEnd)
	}
	return nil
}
