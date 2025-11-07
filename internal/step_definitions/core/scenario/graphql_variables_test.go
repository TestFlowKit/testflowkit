package scenario

import (
	"testing"
)

func TestVariableParser_ParseValue(t *testing.T) {
	parser := NewVariableParser()

	tests := []struct {
		name     string
		input    string
		expected interface{}
		hasError bool
	}{
		// String values
		{
			name:     "simple string",
			input:    "hello world",
			expected: "hello world",
			hasError: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
			hasError: false,
		},

		// Boolean values
		{
			name:     "boolean true",
			input:    "true",
			expected: true,
			hasError: false,
		},
		{
			name:     "boolean false",
			input:    "false",
			expected: false,
			hasError: false,
		},

		// Null value
		{
			name:     "null value",
			input:    "null",
			expected: nil,
			hasError: false,
		},

		// Integer values
		{
			name:     "positive integer",
			input:    "123",
			expected: int64(123),
			hasError: false,
		},
		{
			name:     "negative integer",
			input:    "-456",
			expected: int64(-456),
			hasError: false,
		},
		{
			name:     "zero",
			input:    "0",
			expected: int64(0),
			hasError: false,
		},

		// Float values
		{
			name:     "positive float",
			input:    "123.45",
			expected: 123.45,
			hasError: false,
		},
		{
			name:     "negative float",
			input:    "-67.89",
			expected: -67.89,
			hasError: false,
		},

		// Array values
		{
			name:     "string array",
			input:    `["hello", "world"]`,
			expected: []interface{}{"hello", "world"},
			hasError: false,
		},
		{
			name:     "number array",
			input:    `[1, 2, 3]`,
			expected: []interface{}{float64(1), float64(2), float64(3)},
			hasError: false,
		},
		{
			name:     "mixed array",
			input:    `["hello", 123, true, null]`,
			expected: []interface{}{"hello", float64(123), true, nil},
			hasError: false,
		},
		{
			name:     "empty array",
			input:    `[]`,
			expected: []interface{}{},
			hasError: false,
		},
		{
			name:     "invalid array",
			input:    `[invalid, json]`,
			expected: nil,
			hasError: true,
		},

		// Object values
		{
			name:     "simple object",
			input:    `{"name": "John", "age": 30}`,
			expected: map[string]interface{}{"name": "John", "age": float64(30)},
			hasError: false,
		},
		{
			name:     "nested object",
			input:    `{"user": {"name": "John", "active": true}}`,
			expected: map[string]interface{}{"user": map[string]interface{}{"name": "John", "active": true}},
			hasError: false,
		},
		{
			name:     "empty object",
			input:    `{}`,
			expected: map[string]interface{}{},
			hasError: false,
		},
		{
			name:     "invalid object",
			input:    `{invalid: json}`,
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseValue(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// For complex types, we need to compare differently
			switch expected := tt.expected.(type) {
			case []interface{}:
				resultArray, ok := result.([]interface{})
				if !ok {
					t.Errorf("expected array but got %T", result)
					return
				}
				if len(resultArray) != len(expected) {
					t.Errorf("expected array length %d but got %d", len(expected), len(resultArray))
					return
				}
				for i, expectedItem := range expected {
					if resultArray[i] != expectedItem {
						t.Errorf("array item %d: expected %v but got %v", i, expectedItem, resultArray[i])
					}
				}
			case map[string]interface{}:
				resultMap, ok := result.(map[string]interface{})
				if !ok {
					t.Errorf("expected map but got %T", result)
					return
				}
				if len(resultMap) != len(expected) {
					t.Errorf("expected map length %d but got %d", len(expected), len(resultMap))
					return
				}
				for key, expectedValue := range expected {
					if resultValue, exists := resultMap[key]; !exists {
						t.Errorf("expected key %s not found in result", key)
					} else {
						// Handle nested objects
						if expectedNested, ok := expectedValue.(map[string]interface{}); ok {
							if resultNested, ok := resultValue.(map[string]interface{}); ok {
								for nestedKey, nestedExpected := range expectedNested {
									if nestedResult := resultNested[nestedKey]; nestedResult != nestedExpected {
										t.Errorf("nested key %s.%s: expected %v but got %v", key, nestedKey, nestedExpected, nestedResult)
									}
								}
							} else {
								t.Errorf("expected nested object for key %s but got %T", key, resultValue)
							}
						} else if resultValue != expectedValue {
							t.Errorf("key %s: expected %v but got %v", key, expectedValue, resultValue)
						}
					}
				}
			default:
				if result != tt.expected {
					t.Errorf("expected %v (%T) but got %v (%T)", tt.expected, tt.expected, result, result)
				}
			}
		})
	}
}

func TestVariableParser_ParseVariables(t *testing.T) {
	parser := NewVariableParser()

	input := map[string]string{
		"userId":   "123",
		"tags":     `["frontend", "testing"]`,
		"filters":  `{"status": "active"}`,
		"isActive": "true",
		"score":    "95.5",
	}

	result, err := parser.ParseVariables(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check userId
	if userId, ok := result["userId"]; !ok {
		t.Error("userId not found in result")
	} else if userId != int64(123) {
		t.Errorf("userId: expected %v but got %v", int64(123), userId)
	}

	// Check tags
	if tags, ok := result["tags"]; !ok {
		t.Error("tags not found in result")
	} else if tagsArray, ok := tags.([]interface{}); !ok {
		t.Error("tags is not an array")
	} else if len(tagsArray) != 2 || tagsArray[0] != "frontend" || tagsArray[1] != "testing" {
		t.Errorf("tags: expected [\"frontend\", \"testing\"] but got %v", tagsArray)
	}

	// Check filters
	if filters, ok := result["filters"]; !ok {
		t.Error("filters not found in result")
	} else if filtersMap, ok := filters.(map[string]interface{}); !ok {
		t.Error("filters is not a map")
	} else if status := filtersMap["status"]; status != "active" {
		t.Errorf("filters.status: expected \"active\" but got %v", status)
	}

	// Check isActive
	if isActive, ok := result["isActive"]; !ok {
		t.Error("isActive not found in result")
	} else if isActive != true {
		t.Errorf("isActive: expected true but got %v", isActive)
	}

	// Check score
	if score, ok := result["score"]; !ok {
		t.Error("score not found in result")
	} else if score != 95.5 {
		t.Errorf("score: expected 95.5 but got %v", score)
	}
}

func TestGraphQLContext_VariableManagement(t *testing.T) {
	ctx := NewGraphQLContext()

	// Test SetVariableFromString
	err := ctx.SetVariableFromString("userId", "123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test GetVariable
	if value, exists := ctx.GetVariable("userId"); !exists {
		t.Error("userId not found")
	} else if value != int64(123) {
		t.Errorf("userId: expected %v but got %v", int64(123), value)
	}

	// Test SetArrayVariable
	err = ctx.SetArrayVariable("tags", `["frontend", "testing"]`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value, exists := ctx.GetVariable("tags"); !exists {
		t.Error("tags not found")
	} else if tagsArray, ok := value.([]interface{}); !ok {
		t.Error("tags is not an array")
	} else if len(tagsArray) != 2 {
		t.Errorf("tags: expected length 2 but got %d", len(tagsArray))
	}

	// Test SetObjectVariable
	err = ctx.SetObjectVariable("filters", `{"status": "active"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value, exists := ctx.GetVariable("filters"); !exists {
		t.Error("filters not found")
	} else if filtersMap, ok := value.(map[string]interface{}); !ok {
		t.Error("filters is not a map")
	} else if status := filtersMap["status"]; status != "active" {
		t.Errorf("filters.status: expected \"active\" but got %v", status)
	}

	// Test GetVariableType
	if varType := ctx.GetVariableType("userId"); varType != "integer" {
		t.Errorf("userId type: expected \"integer\" but got %s", varType)
	}

	if varType := ctx.GetVariableType("tags"); varType != "array" {
		t.Errorf("tags type: expected \"array\" but got %s", varType)
	}

	if varType := ctx.GetVariableType("filters"); varType != "object" {
		t.Errorf("filters type: expected \"object\" but got %s", varType)
	}

	// Test HasVariable
	if !ctx.HasVariable("userId") {
		t.Error("HasVariable should return true for userId")
	}

	if ctx.HasVariable("nonexistent") {
		t.Error("HasVariable should return false for nonexistent variable")
	}

	// Test GetVariableCount
	if count := ctx.GetVariableCount(); count != 3 {
		t.Errorf("expected 3 variables but got %d", count)
	}

	// Test RemoveVariable
	ctx.RemoveVariable("userId")
	if ctx.HasVariable("userId") {
		t.Error("userId should be removed")
	}

	if count := ctx.GetVariableCount(); count != 2 {
		t.Errorf("expected 2 variables after removal but got %d", count)
	}
}
