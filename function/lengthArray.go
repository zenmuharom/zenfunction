package function

import (
	"encoding/json"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) LengthArray(arg string) (length int, err error) {

	var decoded []interface{}
	err = json.Unmarshal([]byte(arg), &decoded)
	if err != nil {
		assigner.Logger.Error("LengthArray", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return
	}

	length = len(decoded)

	return
}
