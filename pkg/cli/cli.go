package cli

import (
	"fmt"
	"io"
)

func WriteOutput(w io.Writer, output string) error {
	_, writeErr := fmt.Println(output)
	return writeErr
}

func WriteError(w io.Writer, message string, err error) error {
	_, writeErr := fmt.Fprintf(w, "%s: %s\n", message, err)
	return writeErr
}
