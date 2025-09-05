package error

import "errors"

var InvalidConfigError = errors.New("invalidConfigError")

func IsInvalidConfig(err error) bool {
	return errors.Is(err, InvalidConfigError)
}

var InvalidFlagError = errors.New("invalidFlagError")

func IsInvalidFlag(err error) bool {
	return errors.Is(err, InvalidFlagError)
}

var InvalidSchemaFile = errors.New("invalidSchemaFile")

func IsInvalidSchemaFile(err error) bool {
	return errors.Is(err, InvalidSchemaFile)
}

var DocsGenerationError = errors.New("docsGenerationError")

func IsDocsGenerationError(err error) bool {
	return errors.Is(err, DocsGenerationError)
}

var InvalidOutputFilePath = errors.New("invalidOutputFilePath")

func IsInvalidOutputFilePath(err error) bool {
	return errors.Is(err, InvalidOutputFilePath)
}

var InvalidFileError = errors.New("invalidFileError")

func IsInvalidFileError(err error) bool {
	return errors.Is(err, InvalidFileError)
}

var InvalidDocsPlaceholderError = errors.New("invalidDocsPlaceholderError")

func IsInvalidDocsPlaceholder(err error) bool {
	return errors.Is(err, InvalidDocsPlaceholderError)
}

var InvalidDocsError = errors.New("invalidDocumentationError")

func IsInvalidDocs(err error) bool {
	return errors.Is(err, InvalidDocsError)
}
