package readme

import (
	"fmt"
	"os"
	"strings"

	"github.com/giantswarm/microerror"
)

func (r *Readme) WriteDocs(docs string) error {
	content, err := r.Content()
	if err != nil {
		return microerror.Mask(err)
	}

	docsToReplace, err := docsFromContent(content, r.startPlaceholder, r.endPlaceholder)
	if err != nil {
		return microerror.Mask(err)
	}

	newDocs := fmt.Sprintf("%s\n\n%s\n\n%s", r.startPlaceholder, docs, r.endPlaceholder)
	newContent := strings.ReplaceAll(content, docsToReplace, newDocs)

	err = os.WriteFile(r.path, []byte(newContent), 0600)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
