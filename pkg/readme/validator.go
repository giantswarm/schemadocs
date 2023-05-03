package readme

import (
	"strings"

	"github.com/giantswarm/microerror"
	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
	"github.com/giantswarm/schemadocs/pkg/generate"
	"github.com/google/go-cmp/cmp"
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
		return microerror.Maskf(pkgerror.InvalidDocsError, "documentation from readme %s do not match output generated from %s\n", r.path, schemaPath)
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
