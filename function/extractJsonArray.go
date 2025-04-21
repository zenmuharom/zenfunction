package function

import (
	"errors"
	"strings"
)

func extractJSONArray(input string) (string, string, error) {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "[") {
		return "", input, errors.New("not a JSON array")
	}

	bracketCount := 0
	inString := false
	escaped := false
	end := -1

	for i, r := range input {
		switch {
		case r == '\\' && !escaped:
			escaped = true
			continue
		case r == '"' && !escaped:
			inString = !inString
		case !inString:
			if r == '[' {
				bracketCount++
			} else if r == ']' {
				bracketCount--
				if bracketCount == 0 {
					end = i
					break
				}
			}
		}
		escaped = false
	}

	if end == -1 {
		return "", input, errors.New("unclosed JSON array")
	}

	arrayPart := strings.TrimSpace(input[:end+1])
	rest := strings.TrimSpace(input[end+1:])
	return arrayPart, rest, nil
}
