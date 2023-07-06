package function

import (
	"encoding/json"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) JsonDecode(arg string) (decoded map[string]interface{}, err error) {

	err = json.Unmarshal([]byte(arg), &decoded)
	if err != nil {
		assigner.Logger.Error("JsonDecode", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return
	}
	return
}
