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

// JsonUnpack will unpack json string based on provided keys
// if keys is empty, it will return the whole decoded json
// if keys is provided, it will traverse the decoded json based on the keys
// and return the value found at the end of the keys
// if any key is not found, it will return nil or empty string regarding on the type
// for example:
//
//	jsonStr := `{"a": {"b": {"c": 123}}}`
//	value, err := JsonUnpack(jsonStr, "a", "b", "c") // value will be 123
//	value, err := JsonUnpack(jsonStr, "a", "b") // value will be map[string]interface{}{"c": 123}
//	value, err := JsonUnpack(jsonStr, "a") // value will be map[string]interface{}{"b": map[string]interface{}{"c": 123}}
//	value, err := JsonUnpack(jsonStr) // value will be map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": 123}}}
func (assigner *DefaultAssigner) JsonUnpack(arg string, keys ...string) (value any, err error) {
	// Decode the JSON string
	var decoded map[string]interface{}
	err = json.Unmarshal([]byte(arg), &decoded)
	if err != nil {
		assigner.Logger.Error("JsonUnpack", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return nil, err
	}

	// If no keys provided, return the whole decoded json
	if len(keys) == 0 {
		return decoded, nil
	}

	// Traverse the decoded json based on the keys
	var current interface{} = decoded
	for _, key := range keys {
		// Check if current value is a map
		if currentMap, ok := current.(map[string]interface{}); ok {
			// Get the value for the key
			if val, exists := currentMap[key]; exists {
				current = val
			} else {
				// Key not found, return nil
				return nil, nil
			}
		} else {
			// Current value is not a map, cannot traverse further
			return nil, nil
		}
	}

	return current, nil
}
