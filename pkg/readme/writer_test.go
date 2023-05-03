package readme

import (
	"os"
	"path"
	"testing"

	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
)

func Test_Write(t *testing.T) {
	testCases := []struct {
		name             string
		fileName         string
		content          string
		docs             string
		startPlaceholder string
		endPlaceholder   string
		expectedContent  string
		expectedErr      error
	}{
		{
			name:             "case 0: Write content to valid file",
			fileName:         "test.txt",
			content:          "Text [START_PLACEHOLDER] Docs [END_PLACEHOLDER] Text",
			docs:             "New Docs",
			startPlaceholder: "[START_PLACEHOLDER]",
			endPlaceholder:   "[END_PLACEHOLDER]",
			expectedContent:  "Text [START_PLACEHOLDER]\n\nNew Docs\n\n[END_PLACEHOLDER] Text",
		},
		{
			name:             "case 1: Write empty content to valid file",
			fileName:         "test.txt",
			content:          "Text [START_PLACEHOLDER] Docs [END_PLACEHOLDER] Text",
			docs:             "",
			startPlaceholder: "[START_PLACEHOLDER]",
			endPlaceholder:   "[END_PLACEHOLDER]",
			expectedContent:  "Text [START_PLACEHOLDER]\n\n\n\n[END_PLACEHOLDER] Text",
		},
		{
			name:        "case 2: Write content to valid file with invalid placeholders",
			fileName:    "test.txt",
			content:     "Text [START_PLACEHOLDER] Docs [END_PLACEHOLDER] Text",
			docs:        "New Docs",
			expectedErr: pkgerror.InvalidDocsPlaceholderError,
		},
		{
			name:             "case 3: Write content to invalid file",
			fileName:         "test.txt",
			startPlaceholder: "[START_PLACEHOLDER]",
			endPlaceholder:   "[END_PLACEHOLDER]",
			docs:             "New Docs",
			expectedErr:      pkgerror.InvalidFileError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "*")
			if err != nil {
				t.Fatal(err)
			}
			defer DeleteTempResourceAtPath(tempDir)

			tempFilePath := path.Join(tempDir, tc.fileName)
			tempFile, err := CreateTempFileWithContent(tempFilePath, tc.content)
			if err != nil {
				t.Fatal(err)
			}
			defer DeleteTempFile(tempFile)

			readme, err := New(tempFilePath, tc.startPlaceholder, tc.endPlaceholder)
			if err != nil {
				checkErr(t, err, tc.expectedErr)
				return
			}

			content, err := readme.Content()
			checkFileContent(t, content, tc.content, err, nil)

			err = readme.WriteDocs(tc.docs)
			if err != nil {
				checkErr(t, err, tc.expectedErr)
				return
			}

			content, err = readme.Content()
			checkFileContent(t, content, tc.expectedContent, err, tc.expectedErr)
		})
	}
}
