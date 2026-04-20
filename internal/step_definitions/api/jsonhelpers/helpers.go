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

// ContainsJSON checks that all keys and values present in expected are also
// present in actual (deep subset check). Extra fields in actual are ignored.
func ContainsJSON(expected, actual []byte) error {
	var expectedData, actualData any

	if err := json.Unmarshal(expected, &expectedData); err != nil {
		return fmt.Errorf("failed to unmarshal expected JSON: %w", err)
	}

	if err := json.Unmarshal(actual, &actualData); err != nil {
		return fmt.Errorf("failed to unmarshal actual JSON: %w", err)
	}

	if err := containsValue("$", expectedData, actualData); err != nil {
		return fmt.Errorf("partial JSON mismatch: %w", err)
	}

	return nil
}

func containsValue(path string, expected, actual any) error {
	switch exp := expected.(type) {
	case map[string]any:
		return containsObject(path, exp, actual)
	case []any:
		return containsArray(path, exp, actual)
	default:
		return containsScalar(path, expected, actual)
	}
}

func containsObject(path string, exp map[string]any, actual any) error {
	act, ok := actual.(map[string]any)
	if !ok {
		return fmt.Errorf("at '%s': expected an object but got %T", path, actual)
	}
	for key, expVal := range exp {
		actVal, exists := act[key]
		if !exists {
			return fmt.Errorf("at '%s': key '%s' not found in actual response", path, key)
		}
		if err := containsValue(path+"."+key, expVal, actVal); err != nil {
			return err
		}
	}
	return nil
}

func containsArray(path string, exp []any, actual any) error {
	act, ok := actual.([]any)
	if !ok {
		return fmt.Errorf("at '%s': expected an array but got %T", path, actual)
	}
	if len(exp) != len(act) {
		return fmt.Errorf("at '%s': expected array of length %d but got %d", path, len(exp), len(act))
	}
	for i, expVal := range exp {
		if err := containsValue(fmt.Sprintf("%s[%d]", path, i), expVal, act[i]); err != nil {
			return err
		}
	}
	return nil
}

func containsScalar(path string, expected, actual any) error {
	expJSON, _ := json.Marshal(expected)
	actJSON, _ := json.Marshal(actual)
	if string(expJSON) != string(actJSON) {
		return fmt.Errorf("at '%s': expected '%s' but got '%s'", path, string(expJSON), string(actJSON))
	}
	return nil
}
