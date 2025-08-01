package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (assigner *DefaultAssigner) RemoveItemOnObject(args ...string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("removeItemOnObject requires at least 2 arguments: JSON object and keys to remove")
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(args[0]), &obj); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	for _, key := range args[1:] {
		delete(obj, key)
	}

	encoded, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("failed to marshal result: %w", err)
	}

	return string(encoded), nil
}
