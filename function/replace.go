package function

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Replace(text string, toReplace, replaceTo string, count int) (result string, err error) {
	textType := reflect.ValueOf(text)
	assigner.Logger.Debug("Replace", zenlogger.ZenField{Key: "text", Value: fmt.Sprintf("%v", text)}, zenlogger.ZenField{Key: "toReplace", Value: toReplace}, zenlogger.ZenField{Key: "replaceTo", Value: replaceTo}, zenlogger.ZenField{Key: "count", Value: count}, zenlogger.ZenField{Key: "kind", Value: textType.Kind().String()})

	// Using strings.Replace with count parameter
	result = strings.Replace(text, toReplace, replaceTo, count)

	return
}
