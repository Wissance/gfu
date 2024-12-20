package gfu_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/gfu"
	"github.com/wissance/gwuu/testingutils"
	"os"
	"path"
	"testing"
)

func TestReadAllLine(t *testing.T) {
	testCases := []struct {
		name          string
		fileName      string
		omitEmpty     bool
		expectedLines []string
	}{
		{
			name:      "cr_ending_with_empty_lines_and_omit_empty_lines",
			fileName:  path.Join(".", "testfiles", "cr_ending_with_empty_lines.txt"),
			omitEmpty: true,
			expectedLines: []string{
				"Line one",
				"Line two",
				"Line three",
				"Line four",
				"{",
				"    \"id\": 1,",
				"    \"name\": \"Michael Ushakov\"",
				"}",
			},
		},
		{
			name:      "lf_ending_with_empty_lines_and_omit_empty_lines",
			fileName:  path.Join(".", "testfiles", "lf_ending_with_empty_lines.txt"),
			omitEmpty: true,
			expectedLines: []string{
				"Line one",
				"Line two",
				"Line three",
				"Line four",
				"{",
				"    \"id\": 1,",
				"    \"name\": \"Michael Ushakov\"",
				"}",
			},
		},
		{
			name:      "crlf_ending_with_empty_lines_and_omit_empty_lines",
			fileName:  path.Join(".", "testfiles", "crlf_ending_with_empty_lines.txt"),
			omitEmpty: true,
			expectedLines: []string{
				"Line one",
				"Line two",
				"Line three",
				"Line four",
				"{",
				"    \"id\": 1,",
				"    \"name\": \"Michael Ushakov\"",
				"}",
			},
		},
		{
			name:      "crlf_ending_with_empty_lines_and_not_omit_empty_lines",
			fileName:  path.Join(".", "testfiles", "crlf_ending_with_empty_lines.txt"),
			omitEmpty: false,
			expectedLines: []string{
				"Line one",
				"\t",
				"Line two",
				"Line three",
				"\t",
				"Line four",
				"\t",
				"{",
				"    \"id\": 1,",
				"",
				"    \"name\": \"Michael Ushakov\"",
				"}",
				"",
				"",
			},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			t.Parallel()
			actualLines, err := gfu.ReadAllLines(tCase.fileName, tCase.omitEmpty)
			assert.NoError(t, err)
			testingutils.CheckStrings(t, tCase.expectedLines, actualLines, true, true)
		})
	}
}

func TestWriteAllLines(t *testing.T) {
	lines := []string{
		"Line one",
		"Line two",
		"Line three",
		"Line four",
		"{",
		"    \"id\": 1,",
		"    \"name\": \"Michael Ushakov\"",
		"}",
	}
	testFile := "write_all_lines_test.txt"
	err := gfu.WriteAllLines(testFile, lines, "\n")
	assert.NoError(t, err)
	readLines, err := gfu.ReadAllLines(testFile, true)
	assert.NoError(t, err)
	testingutils.CheckStrings(t, lines, readLines, true, true)
	_ = os.Remove(testFile)
}

func TestAppendAllLines(t *testing.T) {
	lines := []string{
		"Line one",
		"Line two",
		"Line three",
		"Line four",
		"{",
		"    \"id\": 1,",
		"    \"name\": \"Michael Ushakov\"",
		"}",
	}
	testFile := "append_all_lines_test.txt"
	err := gfu.WriteAllLines(testFile, lines, "\n")
	assert.NoError(t, err)

	appendingLines := []string{
		"Line five",
		"Line six",
		"<html><head><title>This is the simplest demo page</title></head><body><h1>Main page header</h1></body></html>",
		"Line seven",
	}

	err = gfu.AppendAllLines(testFile, appendingLines, "\n")
	assert.NoError(t, err)
	expectedLines := make([]string, 0)
	expectedLines = append(expectedLines, lines...)
	expectedLines = append(expectedLines, appendingLines...)
	actualLines, err := gfu.ReadAllLines(testFile, true)
	assert.NoError(t, err)
	testingutils.CheckStrings(t, expectedLines, actualLines, true, true)
	_ = os.Remove(testFile)
}
