package jsonhelpers

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

// GetPathValue extracts a value from JSON data at the specified path.
func GetPathValue(jsonBody []byte, path string) (any, error) {
	var data any
	if err := json.Unmarshal(jsonBody, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON body: %w", err)
	}

	value := gjson.Get(string(jsonBody), path)
	if !value.Exists() {
		return nil, fmt.Errorf("JSON path '%s' does not exist in response body", path)
	}

	return value.Value(), nil
}

// PathExists checks if a JSON path exists in the data.
func PathExists(jsonBody []byte, path string) bool {
	value := gjson.Get(string(jsonBody), path)
	return value.Exists()
}

// GetPathValueAsString extracts a value as string from JSON data at the specified path.
func GetPathValueAsString(jsonBody []byte, path string) (string, error) {
	value := gjson.Get(string(jsonBody), path)
	if !value.Exists() {
		return "", fmt.Errorf("JSON path '%s' does not exist in response body", path)
	}
	return value.String(), nil
}

// CompareJSON compares two JSON byte arrays for equality.
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
