package function

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) ReplaceAll(text string, toReplace, replaceTo string) (result string, err error) {
	textType := reflect.ValueOf(text)
	assigner.Logger.Debug("ReplaceAll", zenlogger.ZenField{Key: "text", Value: fmt.Sprintf("%v", text)}, zenlogger.ZenField{Key: "toReplace", Value: toReplace}, zenlogger.ZenField{Key: "replaceTo", Value: replaceTo}, zenlogger.ZenField{Key: "kind", Value: textType.Kind().String()})

	// Using the ReplaceAll Function
	result = strings.ReplaceAll(text, toReplace, replaceTo)

	return
}
