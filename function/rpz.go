package function

import "strings"

func (assigner *DefaultAssigner) Rpz(arg string, numberOfPad int) (hashed string, err error) {
	if len(arg) >= numberOfPad {
		return arg, nil
	}
	return arg + strings.Repeat("0", numberOfPad-len(arg)), nil
}