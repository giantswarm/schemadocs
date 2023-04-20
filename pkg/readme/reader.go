package readme

import (
	"github.com/giantswarm/microerror"
	"os"
	"strings"
)

func (r *Readme) Content() (string, error) {
	content, err := os.ReadFile(r.path)
	if err != nil {
		return "", microerror.Maskf(invalidFileError, err.Error())
	}

	return string(content), nil
}

func (r *Readme) Docs() (string, error) {
	content, err := r.Content()
	if err != nil {
		return "", microerror.Mask(err)
	}

	docs, err := docsFromContent(content, r.startPlaceholder, r.endPlaceholder)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return docs, nil
}

func docsFromContent(content, startPlaceholder, endPlaceholder string) (string, error) {
	startIndex := strings.Index(content, startPlaceholder)
	if startIndex < 0 {
		return "", microerror.Maskf(invalidDocsPlaceholderError, "start placeholder %s was not found", startPlaceholder)
	}

	lastStartIndex := strings.LastIndex(content, startPlaceholder)
	if startIndex != lastStartIndex {
		return "", microerror.Maskf(invalidDocsPlaceholderError, "multiple start placeholders %s found", startPlaceholder)
	}

	endIndex := strings.Index(content, endPlaceholder)
	if endIndex < 0 {
		return "", microerror.Maskf(invalidDocsPlaceholderError, "end placeholder %s was not found", endPlaceholder)
	}

	lastEndIndex := strings.LastIndex(content, endPlaceholder)
	if endIndex != lastEndIndex {
		return "", microerror.Maskf(invalidDocsPlaceholderError, "multiple end placeholders %s found", endPlaceholder)
	}

	if endIndex < startIndex {
		return "", microerror.Maskf(invalidDocsPlaceholderError, "end placeholders appears before start placeholder %s %s", endPlaceholder, startPlaceholder)
	}

	return content[startIndex : endIndex+len(endPlaceholder)], nil
}
