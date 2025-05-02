package function

import (
	"fmt"
	"strconv"
	"strings"
)

func (assigner *DefaultAssigner) Rps(arg string, numberOfPad int) (result string, err error) {
	unquotedArg, errUnquote := strconv.Unquote(arg)
	if errUnquote == nil {
		arg = unquotedArg
	}

	if len(arg) >= numberOfPad {
		fmt.Println(arg)
		return arg, nil
	}
	result = arg + strings.Repeat(" ", numberOfPad-len(arg))
	fmt.Println(result)
	return
}
