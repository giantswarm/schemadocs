package readme

import (
	"fmt"
	"os"
	"strings"

	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
)

func (r *Readme) Content() (string, error) {
	content, err := os.ReadFile(r.path)
	if err != nil {
		return "", fmt.Errorf("%s: %w", err.Error(), pkgerror.InvalidFileError)
	}

	return string(content), nil
}

func (r *Readme) Docs() (string, error) {
	content, err := r.Content()
	if err != nil {
		return "", fmt.Errorf("failed to read content: %w", err)
	}

	docs, err := docsFromContent(content, r.startPlaceholder, r.endPlaceholder)
	if err != nil {
		return "", fmt.Errorf("failed to extract docs from content: %w", err)
	}

	return docs, nil
}

func docsFromContent(content, startPlaceholder, endPlaceholder string) (string, error) {
	startIndex := strings.Index(content, startPlaceholder)
	if startIndex < 0 {
		return "", fmt.Errorf("start placeholder %s was not found: %w", startPlaceholder, pkgerror.InvalidDocsPlaceholderError)
	}

	lastStartIndex := strings.LastIndex(content, startPlaceholder)
	if startIndex != lastStartIndex {
		return "", fmt.Errorf("multiple start placeholders %s found: %w", startPlaceholder, pkgerror.InvalidDocsPlaceholderError)
	}

	endIndex := strings.Index(content, endPlaceholder)
	if endIndex < 0 {
		return "", fmt.Errorf("end placeholder %s was not found: %w", endPlaceholder, pkgerror.InvalidDocsPlaceholderError)
	}

	lastEndIndex := strings.LastIndex(content, endPlaceholder)
	if endIndex != lastEndIndex {
		return "", fmt.Errorf("multiple end placeholders %s found: %w", endPlaceholder, pkgerror.InvalidDocsPlaceholderError)
	}

	if endIndex < startIndex {
		return "", fmt.Errorf("end placeholders appears before start placeholder %s %s: %w", endPlaceholder, startPlaceholder, pkgerror.InvalidDocsPlaceholderError)
	}

	return content[startIndex : endIndex+len(endPlaceholder)], nil
}
