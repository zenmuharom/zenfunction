package function

import (
	"strings"
)

func (assigner *DefaultAssigner) Substr(arg string, from int, to int) (string, error) {
	var result strings.Builder
	visualCount := 0

	for i := 0; i < len(arg); i++ {
		ch := arg[i]

		// Untuk menghitung visual length
		visualLen := 1
		if ch == '\\' || ch == '\r' || ch == '\n' || ch == '\t' {
			visualLen = 2
		}

		if visualCount+visualLen <= from {
			visualCount += visualLen
			continue
		}
		if visualCount >= from+to {
			break
		}

		// kalau batas terpotong, dan karakter adalah '\', tambahkan \\ biar valid
		if visualCount+visualLen > from+to && ch == '\\' {
			result.WriteString(`\\`)
			break
		}

		result.WriteByte(ch)
		visualCount += visualLen
	}

	return result.String(), nil
}
