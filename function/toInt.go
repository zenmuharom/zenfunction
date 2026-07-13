package function

import (
	"errors"
	"strconv"
)

func (assigner *DefaultAssigner) ToInt(args ...string) (string, error) {
	if len(args) != 1 {
		return "invalid parameter", errors.New("invalid parameter")
	}

	intVal, err := convertToInt64(args[0])
	if err != nil {
		return "invalid parameter", errors.New("invalid parameter")
	}

	return strconv.FormatInt(intVal, 10), nil
}