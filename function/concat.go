package function

import "strings"

// concat concatenates the provided strings with spaces in between.
func (assigner *DefaultAssigner) Concat(strs ...string) string {
	return strings.Join(strs, "")
}
