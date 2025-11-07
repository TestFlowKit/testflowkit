package scenario

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// VariableParser handles parsing of GraphQL variables from string representations
type VariableParser struct{}

// NewVariableParser creates a new variable parser
func NewVariableParser() *VariableParser {
	return &VariableParser{}
}

// ParseValue parses a string value into the appropriate Go type for GraphQL variables
// Supports: JSON arrays, JSON objects, booleans, numbers, and strings
func (vp *VariableParser) ParseValue(value string) (interface{}, error) {
	if value == "" {
		return "", nil
	}

	// Handle JSON arrays
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		var arr []interface{}
		if err := json.Unmarshal([]byte(value), &arr); err != nil {
			return nil, fmt.Errorf("failed to parse array value '%s': %w", value, err)
		}
		return arr, nil
	}

	// Handle JSON objects
	if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
		var obj map[string]interface{}
		if err := json.Unmarshal([]byte(value), &obj); err != nil {
			return nil, fmt.Errorf("failed to parse object value '%s': %w", value, err)
		}
		return obj, nil
	}

	// Handle booleans
	if value == "true" {
		return true, nil
	}
	if value == "false" {
		return false, nil
	}

	// Handle null
	if value == "null" {
		return nil, nil
	}

	// Handle numbers (integers and floats)
	if num, err := strconv.ParseFloat(value, 64); err == nil {
		// Check if it's an integer
		if num == float64(int64(num)) {
			return int64(num), nil
		}
		return num, nil
	}

	// Default to string
	return value, nil
}

// ParseVariables parses a map of string values into properly typed GraphQL variables
func (vp *VariableParser) ParseVariables(variables map[string]string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for key, value := range variables {
		parsedValue, err := vp.ParseValue(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse variable '%s': %w", key, err)
		}
		result[key] = parsedValue
	}

	return result, nil
}

// ValidateArrayValue validates that a string represents a valid JSON array
func (vp *VariableParser) ValidateArrayValue(value string) error {
	if !strings.HasPrefix(value, "[") || !strings.HasSuffix(value, "]") {
		return fmt.Errorf("array value must start with '[' and end with ']'")
	}

	var arr []interface{}
	if err := json.Unmarshal([]byte(value), &arr); err != nil {
		return fmt.Errorf("invalid JSON array format: %w", err)
	}

	return nil
}

// ValidateObjectValue validates that a string represents a valid JSON object
func (vp *VariableParser) ValidateObjectValue(value string) error {
	if !strings.HasPrefix(value, "{") || !strings.HasSuffix(value, "}") {
		return fmt.Errorf("object value must start with '{' and end with '}'")
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(value), &obj); err != nil {
		return fmt.Errorf("invalid JSON object format: %w", err)
	}

	return nil
}

// SerializeValue converts a Go value back to its string representation
// Useful for debugging and logging
func (vp *VariableParser) SerializeValue(value interface{}) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case bool:
		return strconv.FormatBool(v), nil
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), nil
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%g", v), nil
	case nil:
		return "null", nil
	default:
		// For arrays, objects, and other complex types, use JSON marshaling
		bytes, err := json.Marshal(value)
		if err != nil {
			return "", fmt.Errorf("failed to serialize value: %w", err)
		}
		return string(bytes), nil
	}
}

// GetVariableType returns a string representation of the variable's type
func (vp *VariableParser) GetVariableType(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case bool:
		return "boolean"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return "integer"
	case float32, float64:
		return "float"
	case []interface{}:
		return "array"
	case map[string]interface{}:
		return "object"
	case nil:
		return "null"
	default:
		return "unknown"
	}
}
