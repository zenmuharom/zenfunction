package function

import "errors"

func (assigner *DefaultAssigner) ToString(args ...string) (string, error) {
	if len(args) != 1 {
		return "invalid parameter", errors.New("invalid parameter")
	}

	return convertToString(args[0]), nil
}