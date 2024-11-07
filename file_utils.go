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

// WriteAllLines write lines in file truncating it, if file does not exists
// it will be created
func WriteAllLines(file string, lines []string, separator string) error {
	bytes := prepareBytes(lines, separator)
	return os.WriteFile(file, bytes, 0666)
}

func AppendAllLines(file string, lines []string, separator string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	bytes := prepareBytes(lines, separator)
	_, err = f.Write(bytes)
	if err != nil {
		_ = f.Close()
		return err
	}
	err = f.Close()
	return err
}

func WriteAllText(file string, text string) []byte {
	return nil
}

func prepareBytes(lines []string, separator string) []byte {
	textBuilder := &strings.Builder{}
	textBuilder.Grow(8192)

	for _, l := range lines {
		textBuilder.WriteString(l)
		if !strings.HasSuffix(l, separator) {
			textBuilder.WriteString(separator)
		}
	}
	return []byte(textBuilder.String())
}
