package generate

import (
	"errors"
	"flag"
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func Test_Generator(t *testing.T) {

	testCases := []struct {
		name        string
		layout      string
		schemaPath  string
		outputPath  string
		expectedErr error
	}{
		{
			name:       "case 0: Generate markdown from a valid JSON schema",
			layout:     "default",
			schemaPath: "schema.json",
			outputPath: "output_tabular.golden",
		},
		{
			name:        "case 1: Fail to generate markdown from an existing invalid JSON schema",
			layout:      "default",
			schemaPath:  "schema_invalid.json",
			expectedErr: pkgerror.InvalidSchemaFile,
		},
		{
			name:        "case 2: Fail to generate markdown from a on-existent JSON schema",
			layout:      "default",
			expectedErr: pkgerror.InvalidSchemaFile,
		},
		{
			name:       "case 3: Generate markdown from a valid JSON schema in linear layout",
			layout:     "linear",
			schemaPath: "schema.json",
			outputPath: "output_linear.golden",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			docs, err := Generate("testdata/"+tc.schemaPath, tc.layout)

			if err != nil {
				if err != tc.expectedErr && !errors.Is(err, tc.expectedErr) {
					t.Fatalf("received unexpected error: %s\n", err)
				}
				return
			}

			expectedResult := goldenValue(t, tc.outputPath, docs, *update)
			diff := cmp.Diff(docs, expectedResult)
			if diff != "" {
				t.Fatalf("value of generated docs not expected, got: \n %s", diff)
			}
		})
	}

}

func goldenValue(t *testing.T, goldenFile string, actual string, update bool) string {
	t.Helper()
	goldenPath := "testdata/" + goldenFile

	f, err := os.OpenFile(goldenPath, os.O_RDWR, 0644) //nolint:gosec
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	defer func() { _ = f.Close() }()

	if update {
		_, err := f.WriteString(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	return string(content)
}
