package gfu_test

import (
	"gfu"
	"github.com/stretchr/testify/assert"
	"github.com/wissance/gwuu/testingutils"
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
