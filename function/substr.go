package function

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Substr(arg string, from int, to int) (substred string, err error) {
	argType := reflect.ValueOf(arg)

	escapedArg := strconv.Quote(arg)               // escape special chars explicitly
	escapedArg = escapedArg[1 : len(escapedArg)-1] // remove surrounding quotes ("...")
	assigner.Logger.Debug("Substr", zenlogger.ZenField{Key: "arg", Value: fmt.Sprintf("%q", escapedArg)}, zenlogger.ZenField{Key: "from", Value: from}, zenlogger.ZenField{Key: "to", Value: to}, zenlogger.ZenField{Key: "kind", Value: argType.Kind().String()})

	until := len(escapedArg)

	// validasi start of range substring
	if from > until {
		from = until
	}

	// validasi end of range substring
	if from+to <= until {
		until = from + to
	}

	substred = escapedArg[from:until]
	return
}
