package function

import "strings"

func (assigner *DefaultAssigner) Lpz(arg string, numberOfPad int) (hashed string, err error) {
	if len(arg) >= numberOfPad {
		return arg, nil
	}
	return strings.Repeat("0", numberOfPad-len(arg)) + arg, nil
}