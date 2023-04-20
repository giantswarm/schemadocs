package readme

import "github.com/giantswarm/microerror"

var invalidOutputFilePath = &microerror.Error{
	Kind: "invalidOutputFilePath",
}

func isInvalidOutputFilePath(err error) bool {
	return microerror.Cause(err) == invalidOutputFilePath
}

var invalidFileError = &microerror.Error{
	Kind: "invalidFileError",
}

func isInvalidFileError(err error) bool {
	return microerror.Cause(err) == invalidFileError
}

var invalidDocsPlaceholderError = &microerror.Error{
	Kind: "invalidDocsPlaceholderError",
}

func isInvalidDocsPlaceholder(err error) bool {
	return microerror.Cause(err) == invalidDocsPlaceholderError
}

var invalidDocsError = &microerror.Error{
	Kind: "invalidDocsError",
}

func isInvalidDocs(err error) bool {
	return microerror.Cause(err) == invalidDocsError
}
