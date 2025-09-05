package error

import "errors"

var ErrInvalidConfig = errors.New("invalidConfigError")

func IsInvalidConfig(err error) bool {
	return errors.Is(err, ErrInvalidConfig)
}

var ErrInvalidFlag = errors.New("invalidFlagError")

func IsInvalidFlag(err error) bool {
	return errors.Is(err, ErrInvalidFlag)
}

var ErrInvalidSchemaFile = errors.New("invalidSchemaFile")

func IsInvalidSchemaFile(err error) bool {
	return errors.Is(err, ErrInvalidSchemaFile)
}

var ErrDocsGeneration = errors.New("docsGenerationError")

func IsDocsGenerationError(err error) bool {
	return errors.Is(err, ErrDocsGeneration)
}

var ErrInvalidOutputFilePath = errors.New("invalidOutputFilePath")

func IsInvalidOutputFilePath(err error) bool {
	return errors.Is(err, ErrInvalidOutputFilePath)
}

var ErrInvalidFile = errors.New("invalidFileError")

func IsInvalidFileError(err error) bool {
	return errors.Is(err, ErrInvalidFile)
}

var ErrInvalidDocsPlaceholder = errors.New("invalidDocsPlaceholderError")

func IsInvalidDocsPlaceholder(err error) bool {
	return errors.Is(err, ErrInvalidDocsPlaceholder)
}

var ErrInvalidDocs = errors.New("invalidDocumentationError")

func IsInvalidDocs(err error) bool {
	return errors.Is(err, ErrInvalidDocs)
}
