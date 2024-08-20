package function

import "strings"

func (assigner *DefaultAssigner) Lps(arg string, numberOfPad int) (hashed string, err error) {
	if len(arg) >= numberOfPad {
		return arg, nil
	}
	return strings.Repeat(" ", numberOfPad-len(arg)) + arg, nil
}