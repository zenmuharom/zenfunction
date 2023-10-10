package function

import (
	"errors"
	"strings"
)

func escapedCommas(str string) (newStr string) {
	newStr = strings.ReplaceAll(str, ",", `\,`)
	return
}

func unEscapedCommas(str string) (newStr string) {
	newStr = strings.ReplaceAll(str, `\,`, ",")
	return
}

func splitWithEscapedCommas(str string) []string {
	var splitStrings []string
	var currentString string
	escaped := false

	for _, char := range str {
		if char == '\\' && !escaped {
			escaped = true
			continue
		}

		if char == ',' && !escaped {
			splitStrings = append(splitStrings, currentString)
			currentString = ""
			continue
		}

		currentString += string(char)
		escaped = false
	}

	splitStrings = append(splitStrings, currentString)
	return splitStrings
}

func splitWithEscapedCommas2(str string) ([]string, error) {
	var splitStrings []string
	var currentString string
	escaped := false
	inQuotes := false

	for _, char := range str {
		if char == '\\' && !escaped {
			escaped = true
			continue
		}

		if char == '"' && !escaped {
			inQuotes = !inQuotes
			if inQuotes {
				currentString += string(char)
			}
			continue
		}

		if char == ',' && !escaped && !inQuotes {
			splitStrings = append(splitStrings, strings.TrimSpace(currentString))
			currentString = ""
			continue
		}

		currentString += string(char)
		escaped = false
	}

	splitStrings = append(splitStrings, strings.TrimSpace(currentString))

	if inQuotes {
		return nil, errors.New("unmatched double quotes")
	}

	return splitStrings, nil
}
