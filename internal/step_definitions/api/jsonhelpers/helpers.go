package jsonhelpers

import (
	"encoding/json"
	"fmt"
)

func CompareJSON(expected, actual []byte) error {
	var expectedData, actualData any

	if err := json.Unmarshal(expected, &expectedData); err != nil {
		return fmt.Errorf("failed to unmarshal expected JSON: %w", err)
	}

	if err := json.Unmarshal(actual, &actualData); err != nil {
		return fmt.Errorf("failed to unmarshal actual JSON: %w", err)
	}

	expectedJSON, _ := json.Marshal(expectedData)
	actualJSON, _ := json.Marshal(actualData)

	if string(expectedJSON) != string(actualJSON) {
		return fmt.Errorf("JSON mismatch:\nExpected: %s\nActual: %s", string(expectedJSON), string(actualJSON))
	}

	return nil
}

// PrettyPrint formats JSON for readable output.
func PrettyPrint(data []byte) (string, error) {
	var obj any
	if err := json.Unmarshal(data, &obj); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}

	return string(pretty), nil
}

func IsValid(data []byte) bool {
	var js any
	return json.Unmarshal(data, &js) == nil
}
