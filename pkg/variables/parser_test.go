package variables

import (
	"testing"
)

func TestParseValue(t *testing.T) {
	parser := NewParser(&mockStore{})

	tests := []struct {
		name     string
		input    string
		expected any
		wantErr  bool
	}{
		{"empty string", "", "", false},
		{"simple string", "hello", "hello", false},
		{"boolean true", "true", true, false},
		{"boolean false", "false", false, false},
		{"null", "null", (*string)(nil), false},
		{"integer", "42", int64(42), false},
		{"float", "3.14", 3.14, false},
		{"json array", `["a", "b"]`, []any{"a", "b"}, false},
		{"json object", `{"key": "value"}`, map[string]any{"key": "value"}, false},
		{"invalid json array", `[invalid]`, nil, true},
		{"invalid json object", `{invalid}`, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseValue(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareValues(result, tt.expected) {
				t.Errorf("ParseValue() = %v (%T), want %v (%T)", result, result, tt.expected, tt.expected)
			}
		})
	}
}

func TestReplaceInString(t *testing.T) {
	// Create a test context with variables
	mockStore := &mockStore{
		variables: map[string]any{
			"userId": "12345",
			"token":  "abc-token",
			"count":  42,
		},
	}

	parser := NewParser(mockStore)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"no variables", "plain text", "plain text"},
		{"single variable", "User ID: {{userId}}", "User ID: 12345"},
		{"multiple variables", "Token: {{token}}, User: {{userId}}", "Token: abc-token, User: 12345"},
		{"number variable", "Count: {{count}}", "Count: 42"},
		{"unknown variable", "Unknown: {{unknown}}", "Unknown: {{unknown}}"},
		{"variable with spaces", "{{ userId }}", "12345"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.ReplaceInString(tt.input)
			if result != tt.expected {
				t.Errorf("ReplaceInString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSerializeValue(t *testing.T) {
	parser := NewParser(&mockStore{})

	tests := []struct {
		name     string
		input    any
		expected string
		wantErr  bool
	}{
		{"string", "hello", "hello", false},
		{"boolean true", true, "true", false},
		{"boolean false", false, "false", false},
		{"integer", 42, "42", false},
		{"float", 3.14, "3.14", false},
		{"null", nil, "null", false},
		{"array", []any{"a", "b"}, `["a","b"]`, false},
		{"object", map[string]any{"key": "value"}, `{"key":"value"}`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.SerializeValue(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerializeValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("SerializeValue() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractVariableNames(t *testing.T) {
	parser := NewParser(&mockStore{})

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"no variables", "plain text", []string{}},
		{"single variable", "{{userId}}", []string{"userId"}},
		{"multiple variables", "{{userId}} and {{token}}", []string{"userId", "token"}},
		{"variable with spaces", "{{ userId }}", []string{"userId"}},
		{"duplicate variables", "{{userId}} and {{userId}}", []string{"userId", "userId"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.ExtractVariableNames(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ExtractVariableNames() = %v, want %v", result, tt.expected)
				return
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("ExtractVariableNames()[%d] = %v, want %v", i, v, tt.expected[i])
				}
			}
		})
	}
}

func compareValues(a, b any) bool {
	// Simple comparison for basic types
	// For production, use a more robust comparison like reflect.DeepEqual
	switch va := a.(type) {
	case []any:
		vb, ok := b.([]any)
		if !ok || len(va) != len(vb) {
			return false
		}
		for i := range va {
			if !compareValues(va[i], vb[i]) {
				return false
			}
		}
		return true
	case map[string]any:
		vb, ok := b.(map[string]any)
		if !ok || len(va) != len(vb) {
			return false
		}
		for k, v := range va {
			if !compareValues(v, vb[k]) {
				return false
			}
		}
		return true
	default:
		return a == b
	}
}

type mockStore struct {
	variables map[string]any
}

func (ms *mockStore) GetVariable(name string) (any, bool) {
	value, exists := ms.variables[name]
	return value, exists
}
