package function

import "strings"

func (assigner *DefaultAssigner) Rps(arg string, numberOfPad int) (hashed string, err error) {
	if len(arg) >= numberOfPad {
		return arg, nil
	}
	return arg + strings.Repeat(" ", numberOfPad-len(arg)), nil
}