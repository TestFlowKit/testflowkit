package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func GetJSONPathValue(jsonBody []byte, path string) (any, error) {
	var data interface{}
	if err := json.Unmarshal(jsonBody, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON body: %w", err)
	}

	value := gjson.Get(string(jsonBody), path)
	if !value.Exists() {
		return nil, fmt.Errorf("JSON path '%s' does not exist in response body", path)
	}

	return value.Value(), nil
}
