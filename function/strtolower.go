package function

import "strings"

// concat concatenates the provided strings with spaces in between.
func (assigner *DefaultAssigner) Strtolower(origin string) string {
	return strings.ToLower(origin)
}
