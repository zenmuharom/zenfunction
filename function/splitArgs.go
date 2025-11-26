package function

import (
	"strconv"
	"strings"
)

func splitArgs(input string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false
	escaped := false
	depth := 0
	braceDepth := 0   // Track curly braces for JSON objects
	bracketDepth := 0 // Track square brackets for JSON arrays

	for i, r := range input {
		switch {
		case escaped:
			current.WriteRune('\\') // keep escape
			current.WriteRune(r)
			escaped = false
		case r == '\\':
			escaped = true
		case r == '"':
			inQuotes = !inQuotes
			current.WriteRune(r)
		case r == '(' && !inQuotes:
			depth++
			current.WriteRune(r)
		case r == ')' && !inQuotes:
			depth--
			current.WriteRune(r)
		case r == '{' && !inQuotes:
			braceDepth++
			current.WriteRune(r)
		case r == '}' && !inQuotes:
			braceDepth--
			current.WriteRune(r)
		case r == '[' && !inQuotes:
			bracketDepth++
			current.WriteRune(r)
		case r == ']' && !inQuotes:
			bracketDepth--
			current.WriteRune(r)
		case r == ',' && !inQuotes && depth == 0 && braceDepth == 0 && bracketDepth == 0:
			part := strings.TrimSpace(current.String())
			if unquoted, err := strconv.Unquote(part); err == nil {
				args = append(args, unquoted)
			} else {
				args = append(args, part)
			}
			current.Reset()
		default:
			current.WriteRune(r)
		}

		if i == len(input)-1 {
			part := strings.TrimSpace(current.String())
			if unquoted, err := strconv.Unquote(part); err == nil {
				args = append(args, unquoted)
			} else {
				args = append(args, part)
			}
		}
	}

	return args
}
