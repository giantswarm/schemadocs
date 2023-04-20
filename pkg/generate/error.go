package generate

import "github.com/giantswarm/microerror"

var invalidSchemaFile = &microerror.Error{
	Kind: "invalidSchemaFile",
}

func isInvalidSchemaFile(err error) bool {
	return microerror.Cause(err) == invalidSchemaFile
}

var docsGenerationError = &microerror.Error{
	Kind: "docsGenerationError",
}

func isDocsGenerationError(err error) bool {
	return microerror.Cause(err) == docsGenerationError
}
