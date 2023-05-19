package function

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Ltrim(arg0, arg1 string) (trimmed string, err error) {
	argType := reflect.ValueOf(arg0)
	assigner.logger.Debug("Ltrim", zenlogger.ZenField{Key: "arg0", Value: fmt.Sprintf("%v", arg0)}, zenlogger.ZenField{Key: "arg1", Value: fmt.Sprintf("%v", arg1)}, zenlogger.ZenField{Key: "kind", Value: argType.Kind().String()})

	// trimmed = strings.TrimSpace(fmt.Sprintf("%v", argReal))
	trimmed = strings.TrimLeft(fmt.Sprintf("%v", arg0), fmt.Sprintf("%v", arg1))

	return
}
