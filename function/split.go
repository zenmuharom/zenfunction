package function

import "strings"

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
			splitStrings = append(splitStrings, strings.Replace(currentString, " ", "", 1))
			currentString = ""
			continue
		}

		currentString += string(char)
		escaped = false
	}

	// splitStrings = append(splitStrings, strings.TrimSpace(currentString))
	splitStrings = append(splitStrings, currentString)
	return splitStrings
}
