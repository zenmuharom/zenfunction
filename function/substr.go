package function

import (
	"strconv"
)

func (assigner *DefaultAssigner) Substr(arg string, from int, to int) (substred string, err error) {
	// escape explicitly
	escapedArg := strconv.Quote(arg)
	escapedArg = escapedArg[1 : len(escapedArg)-1] // remove surrounding quotes

	// boundary check
	until := len(escapedArg)
	if from > until {
		from = until
	}
	if from+to <= until {
		until = from + to
	}

	// get the substring from escaped string
	subEscaped := escapedArg[from:until]

	// unescape substring explicitly
	unquoted, err := strconv.Unquote(`"` + subEscaped + `"`)
	if err != nil {
		return "", err
	}

	return unquoted, nil
}
