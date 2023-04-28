package readme

import (
	"github.com/giantswarm/microerror"
	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
	"os"
)

const (
	defaultOutputPath1      = "./README.md"
	defaultOutputPath2      = "./Readme.md"
	defaultOutputPath3      = "./readme.md"
	defaultStartPlaceholder = "{::comment} # DOCS_START {:/comment}"
	defaultEndPlaceholder   = "{::comment} # DOCS_END {:/comment}"
)

type Readme struct {
	startPlaceholder string
	endPlaceholder   string
	path             string
}

func New(path, startPlaceholder, endPlaceholder string) (Readme, error) {
	var readme Readme
	var err error

	if path == "" {
		path, err = resolveReadmeFilePathStep([]string{defaultOutputPath1, defaultOutputPath2, defaultOutputPath3})
	} else {
		path, err = resolveReadmeFilePathStep([]string{path})
	}

	if err != nil {
		return readme, microerror.Maskf(pkgerror.InvalidFileError, err.Error())
	}

	if startPlaceholder == "" {
		startPlaceholder = defaultStartPlaceholder
	}
	if endPlaceholder == "" {
		endPlaceholder = defaultEndPlaceholder
	}

	readme = Readme{
		startPlaceholder: startPlaceholder,
		endPlaceholder:   endPlaceholder,
		path:             path,
	}

	return readme, nil
}

func (r *Readme) StartPlaceholder() string {
	return r.startPlaceholder
}

func (r *Readme) EndPlaceholder() string {
	return r.endPlaceholder
}

func (r *Readme) Path() string {
	return r.path
}

func resolveReadmeFilePathStep(paths []string) (string, error) {
	if len(paths) == 0 {
		return "", microerror.Maskf(pkgerror.InvalidOutputFilePath, "valid output file path is not specified")
	}
	if paths[0] != "" {
		_, err := os.Stat(paths[0])
		if err == nil {
			return paths[0], nil
		} else if len(paths) == 1 {
			return "", microerror.Mask(err)
		}
	}
	return resolveReadmeFilePathStep(paths[1:])
}
