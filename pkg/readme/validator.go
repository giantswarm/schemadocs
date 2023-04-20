package readme

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/schemadocs/pkg/generate"
	"github.com/google/go-cmp/cmp"
	"strings"
)

func (r *Readme) Validate(schemaPath string) error {
	docsFromReadme, err := r.Docs()
	if err != nil {
		return microerror.Mask(err)
	}

	docsFromSchema, err := generate.Generate(schemaPath)
	if err != nil {
		return microerror.Mask(err)
	}

	diff := cmp.Diff(strings.TrimSpace(docsFromSchema), trimDocs(docsFromReadme, r.startPlaceholder, r.endPlaceholder))
	if diff != "" {
		return microerror.Maskf(invalidDocsError, "docs from readme %s do not match docs generated from %s:\n%s\n", r.path, schemaPath, diff)
	}

	return nil
}

func trimDocs(docs, startPlaceholder, endPlaceholder string) string {
	if strings.HasPrefix(docs, startPlaceholder) {
		docs = docs[len(startPlaceholder):]
	}
	if strings.HasSuffix(docs, endPlaceholder) {
		docs = docs[:len(docs)-len(endPlaceholder)]
	}
	return strings.TrimSpace(docs)
}
