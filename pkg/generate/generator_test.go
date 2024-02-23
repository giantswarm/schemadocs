package generate

import (
	"os"
	"testing"

	"github.com/giantswarm/microerror"
	"github.com/google/go-cmp/cmp"

	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
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
			schemaPath: "testdata/schema.json",
			outputPath: "testdata/output.txt",
		},
		{
			name:        "case 1: Fail to generate markdown from an existing invalid JSON schema",
			layout:      "default",
			schemaPath:  "testdata/schema_invalid.json",
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
			schemaPath: "testdata/schema.json",
			outputPath: "testdata/output-linear.txt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			docs, err := Generate(tc.schemaPath, tc.layout)

			if err != nil {
				if err != tc.expectedErr && microerror.Cause(err) != tc.expectedErr {
					t.Fatalf("received unexpected error: %s\n", err)
				}
				return
			}

			expectedResult, err := os.ReadFile(tc.outputPath)
			if err != nil {
				t.Fatal(err)
			}

			diff := cmp.Diff(docs, string(expectedResult))
			if diff != "" {
				t.Fatalf("value of generated docs not expected, got: \n %s", diff)
			}
		})
	}

}
