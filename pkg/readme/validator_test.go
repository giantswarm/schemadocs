package readme

import "testing"

func Test_Validate(t *testing.T) {
	testCases := []struct {
		name             string
		readmeFilePath   string
		schemaFilePath   string
		startPlaceholder string
		endPlaceholder   string
		expectedErr      error
	}{
		{
			name:           "case 0: Succeed when the provided and generated docs match",
			readmeFilePath: "testdata/readme1.md",
			schemaFilePath: "testdata/schema.json",
		},
		{
			name:           "case 1: Fail when the provided and generated docs do not match",
			readmeFilePath: "testdata/readme2.md",
			schemaFilePath: "testdata/schema.json",
			expectedErr:    invalidDocsError,
		},
		{
			name:           "case 2: Fail when the provided docs is not valid",
			schemaFilePath: "testdata/schema.json",
			expectedErr:    invalidFileError,
		},
		{
			name:             "case 2: Fail when the docs placeholders are not valid",
			readmeFilePath:   "testdata/readme1.md",
			schemaFilePath:   "testdata/schema.json",
			startPlaceholder: "[DOCS_START]",
			expectedErr:      invalidDocsPlaceholderError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			readme, err := New(tc.readmeFilePath, tc.startPlaceholder, tc.endPlaceholder)
			if err == nil {
				err = readme.Validate(tc.schemaFilePath)
			}
			checkErr(t, err, tc.expectedErr)
		})
	}
}
