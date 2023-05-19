package function

import (
	"fmt"
	"reflect"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Substr(arg string, from int, to int) (substred string, err error) {
	argType := reflect.ValueOf(arg)
	assigner.logger.Debug("Substr", zenlogger.ZenField{Key: "arg", Value: fmt.Sprintf("%v", arg)}, zenlogger.ZenField{Key: "from", Value: from}, zenlogger.ZenField{Key: "to", Value: to}, zenlogger.ZenField{Key: "kind", Value: argType.Kind().String()})

	until := len(fmt.Sprintf("%v", arg))

	// validasi start of range substring
	if from > until {
		from = until
	}

	// validasi end of range substring
	if from+to <= until {
		until = from + to
	}

	substred = fmt.Sprintf("%v", arg)[from:until]
	return
}
