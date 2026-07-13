package function

import (
	"errors"
	"strconv"
)

func (assigner *DefaultAssigner) ToFloat(args ...string) (string, error) {
	if len(args) != 1 {
		return "invalid parameter", errors.New("invalid parameter")
	}

	floatVal, err := convertToFloat64(args[0])
	if err != nil {
		return "invalid parameter", errors.New("invalid parameter")
	}

	return strconv.FormatFloat(floatVal, 'f', -1, 64), nil
}