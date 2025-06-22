package httpapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

func getValueFromDotNotation(jsonBody []byte, path string) (interface{}, error) {
	var data interface{}
	if err := json.Unmarshal(jsonBody, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body: %w", err)
	}
	parts := strings.Split(path, ".")
	current := data
	for _, part := range parts {
		m, ok := current.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("path part '%s' requires a map, but got %T", part, current)
		}
		current, ok = m[part]
		if !ok {
			return nil, fmt.Errorf("key '%s' not found in map", part)
		}
	}
	return current, nil
}
