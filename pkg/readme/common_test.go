package readme

import (
	"fmt"
	"github.com/giantswarm/microerror"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func createTempFileWithContent(path, content string) (*os.File, error) {
	if content == "" {
		return nil, nil
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(content)
	if err != nil {
		deleteTempFile(file)
		return nil, err
	}

	return file, nil
}

func deleteTempFile(tempFile *os.File) {
	if tempFile == nil {
		return
	}
	deleteTempResourceAtPath(tempFile.Name())
}

func deleteTempResourceAtPath(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
}

func checkFileContent(t *testing.T, content string, expectedContent string, err error, expectedErr error) {
	if expectedErr != nil && err != expectedErr && microerror.Cause(err) != expectedErr {
		t.Fatalf("expected error %s, received %s", expectedErr, err)
	}

	if expectedErr == nil {
		if err != nil {
			t.Fatalf("received unexpected error %s", err)
		}

		diff := cmp.Diff(content, expectedContent)
		if diff != "" {
			t.Fatalf("received unexpected file content:\n%s\n", diff)
		}
	}
}

func checkErr(t *testing.T, err error, expectedErr error) {
	if err != expectedErr && microerror.Cause(err) != expectedErr {
		t.Fatalf("received unexpected error: %s\n", err)
	}
}
