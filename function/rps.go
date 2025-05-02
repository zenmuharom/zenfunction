package function

import (
	"strconv"
	"strings"
)

func (assigner *DefaultAssigner) Rps(arg string, numberOfPad int) (result string, err error) {
	unquotedArg, errUnquote := strconv.Unquote(arg)
	if errUnquote == nil {
		arg = unquotedArg
	}

	if len(arg) >= numberOfPad {
		return arg, nil
	}
	result = arg + strings.Repeat(" ", numberOfPad-len(arg))
	return
}
