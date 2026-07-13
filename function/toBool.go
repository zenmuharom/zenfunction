package function

import (
	"errors"
	"strconv"
)

func (assigner *DefaultAssigner) ToBool(args ...string) (string, error) {
	if len(args) != 1 {
		return "invalid parameter", errors.New("invalid parameter")
	}

	boolVal, err := convertToBool(args[0])
	if err != nil {
		return "invalid parameter", errors.New("invalid parameter")
	}

	return strconv.FormatBool(boolVal), nil
}