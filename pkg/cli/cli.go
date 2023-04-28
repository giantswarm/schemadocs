package cli

import (
	"fmt"
	"io"
	"log"

	"github.com/fatih/color"
)

const (
	errorPrefix   = "[ERROR]"
	successPrefix = "[SUCCESS]"
	failurePrefix = "[FAILURE]"
)

var (
	SuccessColor = color.New(color.FgGreen)
	ErrorColor   = color.New(color.FgRed).Add(color.Bold)
)

func WriteError(w io.Writer, noColor bool, message string, err error) {
	WriteOutputF(w, "%s %s\nReason: %s\n", prefix(errorPrefix, noColor, ErrorColor.Sprint), message, err)
}

func WriteSuccess(w io.Writer, noColor bool, output string) {
	WriteOutputF(w, "%s %s", prefix(successPrefix, noColor, SuccessColor.Sprint), output)
}

func WriteFailure(w io.Writer, noColor bool, message string, err error) {
	WriteOutputF(w, "%s %s\nReason: %s\n", prefix(failurePrefix, noColor, ErrorColor.Sprint), message, err)
}

func WriteNewLine(w io.Writer) {
	WriteOutput(w, "")
}

func WriteOutput(w io.Writer, output string) {
	_, err := fmt.Fprintln(w, output)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteOutputF(w io.Writer, format string, args ...interface{}) {
	_, err := fmt.Fprintf(w, format, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func prefix(prefixStr string, noColor bool, colorF func(...interface{}) string) string {
	if noColor {
		return prefixStr
	}
	return colorF(prefixStr)
}
