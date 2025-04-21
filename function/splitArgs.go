package function

import (
	"strings"
	"unicode"
)

func splitArgs(input string) []string {
	var args []string
	runes := []rune(input)
	n := len(runes)
	i := 0

	for i < n {
		for i < n && unicode.IsSpace(runes[i]) {
			i++
		}
		if i >= n {
			break
		}

		if runes[i] == '"' {
			j := i + 1
			var sb strings.Builder
			for j < n {
				if runes[j] == '\\' && j+1 < n {
					sb.WriteRune(runes[j])
					sb.WriteRune(runes[j+1])
					j += 2
					continue
				}
				if runes[j] == '"' {
					break
				}
				sb.WriteRune(runes[j])
				j++
			}
			if sb.Len() > 0 {
				args = append(args, sb.String())
			}
			i = j + 1
			continue
		}

		if unicode.IsDigit(runes[i]) {
			j := i
			for j < n && unicode.IsDigit(runes[j]) {
				j++
			}
			args = append(args, string(runes[i:j]))
			i = j
			continue
		}

		j := i
		depth := 0
		for j < n {
			if runes[j] == '(' {
				depth++
			} else if runes[j] == ')' {
				if depth > 0 {
					depth--
				}
			} else if runes[j] == ',' && depth == 0 {
				break
			}
			j++
		}
		arg := strings.TrimSpace(string(runes[i:j]))
		if arg != "" {
			args = append(args, arg)
		}
		i = j + 1
	}
	return args
}
