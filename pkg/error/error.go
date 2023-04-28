package error

import "github.com/giantswarm/microerror"

var InvalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == InvalidConfigError
}

var InvalidFlagError = &microerror.Error{
	Kind: "invalidFlagError",
}

func IsInvalidFlag(err error) bool {
	return microerror.Cause(err) == InvalidFlagError
}

var InvalidSchemaFile = &microerror.Error{
	Kind: "invalidSchemaFile",
}

func IsInvalidSchemaFile(err error) bool {
	return microerror.Cause(err) == InvalidSchemaFile
}

var DocsGenerationError = &microerror.Error{
	Kind: "docsGenerationError",
}

func IsDocsGenerationError(err error) bool {
	return microerror.Cause(err) == DocsGenerationError
}

var InvalidOutputFilePath = &microerror.Error{
	Kind: "invalidOutputFilePath",
}

func IsInvalidOutputFilePath(err error) bool {
	return microerror.Cause(err) == InvalidOutputFilePath
}

var InvalidFileError = &microerror.Error{
	Kind: "invalidFileError",
}

func IsInvalidFileError(err error) bool {
	return microerror.Cause(err) == InvalidFileError
}

var InvalidDocsPlaceholderError = &microerror.Error{
	Kind: "invalidDocsPlaceholderError",
}

func IsInvalidDocsPlaceholder(err error) bool {
	return microerror.Cause(err) == InvalidDocsPlaceholderError
}

var InvalidDocsError = &microerror.Error{
	Kind: "invalidDocumentationError",
}

func IsInvalidDocs(err error) bool {
	return microerror.Cause(err) == InvalidDocsError
}
