package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"

	cmderror "github.com/giantswarm/schemadocs/pkg/error"
)

func Test_Root(t *testing.T) {
	testCases := []struct {
		name                 string
		command              string
		argFilePath          string
		flag                 string
		flagFilePath         string
		expectedOutput       string
		formatExpectedOutput bool
		expectedErr          error
	}{
		{
			name:                 "case 0: Generate command - show success message with valid inputs",
			command:              "generate",
			argFilePath:          "testdata/schema.json",
			flag:                 "--output-path",
			flagFilePath:         "testdata/readme.md",
			expectedOutput:       "testdata/output_format_generate_success.golden.txt",
			formatExpectedOutput: true,
		},
		{
			name:                 "case 1: Generate command - show error message when no flags are provided",
			command:              "generate",
			argFilePath:          "testdata/schema.json",
			expectedOutput:       "testdata/output_format_generate_failure_1.golden.txt",
			formatExpectedOutput: true,
			expectedErr:          cmderror.InvalidFileError,
		},
		{
			name:                 "case 2: Generate command - show error message when valid arg points to an invalid schema file",
			command:              "generate",
			argFilePath:          "testdata/schema2.json",
			flag:                 "--output-path",
			flagFilePath:         "testdata/readme.md",
			expectedOutput:       "testdata/output_format_generate_failure_2.golden.txt",
			formatExpectedOutput: true,
			expectedErr:          cmderror.InvalidSchemaFile,
		},
		{
			name:                 "case 3: Generate command - show error message when valid flag points to an invalid readme file",
			command:              "generate",
			argFilePath:          "testdata/schema.json",
			flag:                 "--output-path",
			flagFilePath:         "testdata/readme2.md",
			expectedOutput:       "testdata/output_format_generate_failure_3.golden.txt",
			formatExpectedOutput: true,
			expectedErr:          cmderror.InvalidFileError,
		},

		{
			name:                 "case 4: Validate command - show success message with valid inputs",
			command:              "validate",
			argFilePath:          "testdata/readme.md",
			flag:                 "--schema",
			flagFilePath:         "testdata/schema.json",
			expectedOutput:       "testdata/output_format_validate_success.golden.txt",
			formatExpectedOutput: true,
		},
		{
			name:                 "case 5: Validate command - show error message when no flags are provided",
			command:              "validate",
			argFilePath:          "testdata/readme.md",
			expectedOutput:       "testdata/output_format_validate_failure_1.golden.txt",
			formatExpectedOutput: false,
			expectedErr:          cmderror.InvalidFlagError,
		},
		{
			name:                 "case 6: Validate command - show error message when valid arg points to an invalid readme file",
			command:              "validate",
			argFilePath:          "testdata/readme2.md",
			flag:                 "--schema",
			flagFilePath:         "testdata/schema.json",
			expectedOutput:       "testdata/output_format_validate_failure_2.golden.txt",
			formatExpectedOutput: true,
			expectedErr:          cmderror.InvalidFileError,
		},
		{
			name:                 "case 7: Validate command - show error message when valid flag points to an invalid schema file",
			command:              "validate",
			argFilePath:          "testdata/readme.md",
			flag:                 "--schema",
			flagFilePath:         "testdata/schema2.json",
			expectedOutput:       "testdata/output_format_validate_failure_3.golden.txt",
			formatExpectedOutput: true,
			expectedErr:          cmderror.InvalidSchemaFile,
		},
		{
			name:                 "case 8: Validate command - show error message when the generated and provided documentation do not match",
			command:              "validate",
			argFilePath:          "testdata/readme_diff.md",
			flag:                 "--schema",
			flagFilePath:         "testdata/schema.json",
			expectedOutput:       "testdata/output_format_validate_failure_4.golden.txt",
			formatExpectedOutput: true,
			expectedErr:          cmderror.InvalidDocsError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rootDir, err := os.MkdirTemp("", "*")
			checkError(t, err)
			defer deleteResourcesAtPath(rootDir)

			argFilePath, err := cloneFileToTempDir(tc.argFilePath, rootDir)
			checkError(t, err)

			flagFilePath, err := cloneFileToTempDir(tc.flagFilePath, rootDir)
			checkError(t, err)

			output := new(bytes.Buffer)

			args := []string{"schemadocs"}
			if tc.command != "" {
				args = append(args, tc.command)
			}
			if argFilePath != "" {
				args = append(args, argFilePath)
			}
			if tc.flag != "" {
				args = append(args, tc.flag)
			}
			if flagFilePath != "" {
				args = append(args, flagFilePath)
			}
			args = append(args, "--no-color")

			os.Args = args

			cmd, err := New(Config{Stderr: output, Stdout: output})
			checkError(t, err)

			err = cmd.Execute()

			if tc.expectedErr != nil {
				if err != tc.expectedErr && !errors.Is(err, tc.expectedErr) {
					t.Fatalf("received unexpected error: expected '%s', received '%s'\n", tc.expectedErr, err)
				}
				checkOutput(t, output.String(), tc.expectedOutput, argFilePath, flagFilePath, tc.formatExpectedOutput)
			} else {
				checkError(t, err)
				checkOutput(t, output.String(), tc.expectedOutput, argFilePath, flagFilePath, tc.formatExpectedOutput)
			}
		})
	}
}

func checkOutput(t *testing.T, actual, expected, argFilePath, flagFilePath string, formatExpectedOutput bool) {
	expectedOutput, err := readGoldenFile(expected)
	checkError(t, err)
	if formatExpectedOutput {
		expectedOutput = fmt.Sprintf(expectedOutput, argFilePath, flagFilePath)
	}
	diff := cmp.Diff(expectedOutput, actual)
	if diff != "" {
		t.Fatalf("received unexpected output: %s\n", diff)
	}
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("received unexpected error: %s\n", err)
	}
}

func deleteResourcesAtPath(path string) {
	fmt.Printf("Deleting %s\n", path)
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func cloneFileToTempDir(srcFilePath string, dstRootDir string) (string, error) {
	var dstFilePath, dstFileDir string
	if srcFilePath != "" {
		dstFilePath = path.Join(dstRootDir, srcFilePath)
		dstFileDir = path.Dir(dstFilePath)
		err := os.MkdirAll(dstFileDir, os.ModePerm) //nolint:gosec
		if err != nil {
			return "", err
		}

		data, err := os.ReadFile(srcFilePath) //nolint:gosec
		if err == nil {
			err = os.WriteFile(dstFilePath, data, 0600)
			if err != nil {
				return "", err
			}
		}
	}
	return dstFilePath, nil
}

func readGoldenFile(path string) (string, error) {
	data, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		return "", err
	}
	return string(data), nil
}
