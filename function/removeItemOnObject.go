package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (a *DefaultAssigner) RemoveItemOnObject(args ...string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("removeItemOnObject requires at least 2 arguments: object and at least one key to remove")
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(args[0]), &data); err != nil {
		return "", fmt.Errorf("invalid JSON object: %w", err)
	}

	for _, key := range args[1:] {
		delete(data, key)
	}

	result, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal result: %w", err)
	}

	return string(result), nil
}
