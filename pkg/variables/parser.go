package variables

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const nullString = "null"

type Parser struct {
	store Store
}

func NewParser(store Store) *Parser {
	return &Parser{store: store}
}

// Supports: JSON arrays, JSON objects, booleans, numbers, and strings.
func (p *Parser) ParseValue(value string) (any, error) {
	if value == "" {
		return "", nil
	}

	// Handle JSON arrays
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		var arr []any
		if err := json.Unmarshal([]byte(value), &arr); err != nil {
			return nil, fmt.Errorf("failed to parse array value '%s': %w", value, err)
		}
		return arr, nil
	}

	// Handle JSON objects
	if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
		var obj map[string]any
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
	if value == nullString {
		return (*string)(nil), nil
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

func (p *Parser) ParseVariables(variables map[string]string) (map[string]any, error) {
	result := make(map[string]any)

	for key, value := range variables {
		parsedValue, err := p.ParseValue(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse variable '%s': %w", key, err)
		}
		result[key] = parsedValue
	}

	return result, nil
}

// ReplaceInString replaces all variable placeholders in the format {{variableName}} with their actual values
// This is used for REST API variable substitution in URLs, headers, and body content
// NOTE: ReplaceInString/Bytes/Map functions moved to runtime.go to avoid import cycles

// ValidateArrayValue validates that a string represents a valid JSON array.
func (p *Parser) ValidateArrayValue(value string) error {
	if !strings.HasPrefix(value, "[") || !strings.HasSuffix(value, "]") {
		return errors.New("array value must start with '[' and end with ']'")
	}

	var arr []any
	if err := json.Unmarshal([]byte(value), &arr); err != nil {
		return fmt.Errorf("invalid JSON array format: %w", err)
	}

	return nil
}

// ValidateObjectValue validates that a string represents a valid JSON object.
func (p *Parser) ValidateObjectValue(value string) error {
	if !strings.HasPrefix(value, "{") || !strings.HasSuffix(value, "}") {
		return errors.New("object value must start with '{' and end with '}'")
	}

	var obj map[string]any
	if err := json.Unmarshal([]byte(value), &obj); err != nil {
		return fmt.Errorf("invalid JSON object format: %w", err)
	}

	return nil
}

// SerializeValue converts a Go value back to its string representation
// Useful for debugging, logging, and variable substitution.
func (p *Parser) SerializeValue(value any) (string, error) {
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
		return nullString, nil
	default:
		// For arrays, objects, and other complex types, use JSON marshaling
		bytes, err := json.Marshal(value)
		if err != nil {
			return "", fmt.Errorf("failed to serialize value: %w", err)
		}
		return string(bytes), nil
	}
}

// GetVariableType returns a string representation of the variable's type.
func (p *Parser) GetVariableType(value any) string {
	switch value.(type) {
	case string:
		return "string"
	case bool:
		return "boolean"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return "integer"
	case float32, float64:
		return "float"
	case []any:
		return "array"
	case map[string]any:
		return "object"
	case nil:
		return nullString
	default:
		return "unknown"
	}
}

// HasVariablePlaceholders checks if a string contains any variable placeholders.
func (p *Parser) HasVariablePlaceholders(input string) bool {
	return variablePlaceholderRe.MatchString(input)
}

// ExtractVariableNames returns all variable names found in a string.
func (p *Parser) ExtractVariableNames(input string) []string {
	matches := extractVariablesNameRe.FindAllStringSubmatch(input, -1)

	names := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) > 1 {
			names = append(names, strings.TrimSpace(match[1]))
		}
	}

	return names
}

var extractVariablesNameRe = regexp.MustCompile(`\{\{([^}]+)\}\}`)
var variablePlaceholderRe = regexp.MustCompile(`\{\{[^}]+\}\}`)
