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
		schemaPath  string
		outputPath  string
		expectedErr error
	}{
		{
			name:       "case 0: Generate markdown from a valid JSON schema",
			schemaPath: "testdata/schema.json",
			outputPath: "testdata/output.txt",
		},
		{
			name:        "case 1: Fail to generate markdown from an existing invalid JSON schema",
			schemaPath:  "testdata/schema_invalid.json",
			expectedErr: pkgerror.InvalidSchemaFile,
		},
		{
			name:        "case 2: Fail to generate markdown from a on-existent JSON schema",
			expectedErr: pkgerror.InvalidSchemaFile,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			docs, err := Generate(tc.schemaPath)

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
