package function

import "fmt"

func (assigner *DefaultAssigner) Substr(arg string, from int, to int) (result string, err error) {
	length := len(arg)
	result = arg
	if from < 0 || from >= length {
		err = fmt.Errorf("from index %d out of range (0-%d)", from, length-1)
		return
	}

	if to < 0 {
		err = fmt.Errorf("length must be non-negative, got %d", to)
		return
	}

	end := from + to
	if end > length {
		end = length
	}

	result = arg[from:end]
	return
}
