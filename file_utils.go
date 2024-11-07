package gfu

import (
	"os"
	"strings"
)

// ReadAllLines function that reads file content with auto-detection of line ending (\n, \r or \r\n) and return lines
// without ending symbols, empty lines could be omitted
func ReadAllLines(file string, omitEmpty bool) ([]string, error) {
	var err error
	var content []byte
	// 1. Read file in memory
	content, err = os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// 2. Define line separator, most popular \n (Linux) or \r\n (Windows), anyway \r\n already contains \r
	endSeparator := "\n"
	if strings.Contains(string(content), endSeparator) {
		// using Mac-like separator
		endSeparator = "\r"
	}
	rawLines := strings.Split(string(content), endSeparator)
	lines := make([]string, 0)

	finalLinesNumber := 0
	if !omitEmpty {
		finalLinesNumber = len(rawLines)
	}

	for _, l := range rawLines {
		l = strings.Trim(l, "\r\n")
		if omitEmpty {
			spaceTrimmedLine := strings.Trim(l, " \t")
			if len(spaceTrimmedLine) > 0 {
				lines = append(lines, l)
				finalLinesNumber++
			}
		} else {
			lines = append(lines, l)
		}
	}

	return lines[0:finalLinesNumber], err
}

// ReadAllText just wraps os.ReadFile and return reading result as Text (string)
func ReadAllText(file string) (string, error) {
	var err error
	var content []byte
	// 1. Read file in memory
	content, err = os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), err
}

func WriteAllLines(file string, lines []string, separator string) error {
	return nil
}

func WriteAllText(file string, text string) error {
	return nil
}
