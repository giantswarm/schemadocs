package readme

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func Test_Content(t *testing.T) {
	testCases := []struct {
		name        string
		fileName    string
		content     string
		expectedErr error
	}{
		{
			name:     "case 0: Read content from valid file",
			fileName: "text.txt",
			content:  "Docs",
		},
		{
			name:        "case 1: Read content from invalid file",
			fileName:    "text.txt",
			expectedErr: invalidFileError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "*")
			if err != nil {
				t.Fatal(err)
			}
			defer deleteTempResourceAtPath(tempDir)

			tempFilePath := path.Join(tempDir, tc.fileName)
			tempFile, err := createTempFileWithContent(tempFilePath, tc.content)
			defer deleteTempFile(tempFile)
			if err != nil {
				t.Fatal(err)
			}

			var readme Readme
			var content string

			readme, err = New(tempFilePath, "", "")
			if err == nil {
				content, err = readme.Content()
			}

			checkFileContent(t, content, tc.content, err, tc.expectedErr)
		})
	}
}

func Test_Docs(t *testing.T) {
	testCases := []struct {
		name             string
		fileName         string
		content          string
		startPlaceholder string
		endPlaceholder   string
		expectedDocs     string
		expectedErr      error
	}{
		{
			name:             "case 0: Read docs from a valid file with valid placeholders",
			fileName:         "test.txt",
			content:          "Text [START_PLACEHOLDER] Docs [END_PLACEHOLDER] Text",
			startPlaceholder: "[START_PLACEHOLDER]",
			endPlaceholder:   "[END_PLACEHOLDER]",
			expectedDocs:     "[START_PLACEHOLDER] Docs [END_PLACEHOLDER]",
		},
		{
			name:         "case 1: Read docs from a valid file with empty placeholders",
			fileName:     "test.txt",
			content:      fmt.Sprintf("Text %s Docs %s Text", defaultStartPlaceholder, defaultEndPlaceholder),
			expectedDocs: fmt.Sprintf("%s Docs %s", defaultStartPlaceholder, defaultEndPlaceholder),
		},
		{
			name:        "case 2: Read docs from a valid file with invalid placeholders",
			fileName:    "test.txt",
			content:     "Text [START_PLACEHOLDER] Docs [END_PLACEHOLDER] Text",
			expectedErr: invalidDocsPlaceholderError,
		},
		{
			name:        "case 3: Read docs from file with invalid content",
			fileName:    "test.txt",
			content:     fmt.Sprintf("Text %s Text", defaultEndPlaceholder),
			expectedErr: invalidDocsPlaceholderError,
		},
		{
			name:        "case 4: Read docs from invalid file",
			fileName:    "test.txt",
			content:     "",
			expectedErr: invalidFileError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "*")
			if err != nil {
				t.Fatal(err)
			}
			defer deleteTempResourceAtPath(tempDir)

			tempFilePath := path.Join(tempDir, tc.fileName)
			tempFile, err := createTempFileWithContent(tempFilePath, tc.content)
			if err != nil {
				t.Fatal(err)
			}
			defer deleteTempFile(tempFile)

			var readme Readme
			var docs string

			readme, err = New(tempFilePath, tc.startPlaceholder, tc.endPlaceholder)
			if err == nil {
				docs, err = readme.Docs()
			}

			checkFileContent(t, docs, tc.expectedDocs, err, tc.expectedErr)
		})
	}
}
