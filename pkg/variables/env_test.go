package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGetEnvVariable(t *testing.T) {
	// Reset before test
	ResetEnvVariables()

	// Test setting and getting
	SetEnvVariables(map[string]string{
		"api.url": "https://api.example.com",
		"api.key": "secret",
		"timeout": "30s",
	})

	val, ok := GetEnvVariable("api.url")
	assert.True(t, ok)
	assert.Equal(t, "https://api.example.com", val)

	val, ok = GetEnvVariable("api.key")
	assert.True(t, ok)
	assert.Equal(t, "secret", val)

	val, ok = GetEnvVariable("nonexistent")
	assert.False(t, ok)
	assert.Equal(t, "", val)
}

func TestSetEnvVariablesWithNil(t *testing.T) {
	// Should not panic with nil input
	SetEnvVariables(nil)

	all := GetAllEnvVariables()
	assert.Empty(t, all)
}

func TestGetAllEnvVariables(t *testing.T) {
	ResetEnvVariables()

	vars := map[string]string{
		"var1": "value1",
		"var2": "value2",
	}
	SetEnvVariables(vars)

	all := GetAllEnvVariables()
	assert.Equal(t, 2, len(all))
	assert.Equal(t, "value1", all["var1"])
	assert.Equal(t, "value2", all["var2"])

	// Modifying returned map should not affect store
	all["var3"] = "value3"

	val, ok := GetEnvVariable("var3")
	assert.False(t, ok)
	assert.Equal(t, "", val)
}

func TestFindUndefinedEnvReferences(t *testing.T) {
	ResetEnvVariables()
	SetEnvVariables(map[string]string{
		"defined.var": "value",
		"api.url":     "https://api.example.com",
	})

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "no references",
			content:  "just plain text",
			expected: []string{},
		},
		{
			name:     "empty content",
			content:  "",
			expected: []string{},
		},
		{
			name:     "all defined",
			content:  "{{ env.defined.var }} and {{ env.api.url }}",
			expected: []string{},
		},
		{
			name:     "one undefined",
			content:  "{{ env.defined.var }} and {{ env.undefined }}",
			expected: []string{"undefined"},
		},
		{
			name:     "multiple undefined",
			content:  "{{ env.missing1 }} and {{ env.missing2 }}",
			expected: []string{"missing1", "missing2"},
		},
		{
			name:     "duplicate references",
			content:  "{{ env.missing }} and {{ env.missing }} again",
			expected: []string{"missing"},
		},
		{
			name:     "with spaces",
			content:  "{{  env.missing  }}",
			expected: []string{"missing"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindUndefinedEnvReferences(tt.content)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReplaceEnvVariables(t *testing.T) {
	ResetEnvVariables()
	SetEnvVariables(map[string]string{
		"api.url":       "https://api.example.com",
		"api.key":       "secret123",
		"database.host": "localhost",
		"database.port": "5432",
	})

	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "no variables",
			content:  "plain text",
			expected: "plain text",
		},
		{
			name:     "empty content",
			content:  "",
			expected: "",
		},
		{
			name:     "single variable",
			content:  "URL: {{ env.api.url }}",
			expected: "URL: https://api.example.com",
		},
		{
			name:     "multiple variables",
			content:  "{{ env.api.url }}/data?key={{ env.api.key }}",
			expected: "https://api.example.com/data?key=secret123",
		},
		{
			name:     "nested path variables",
			content:  "Connect to {{ env.database.host }}:{{ env.database.port }}",
			expected: "Connect to localhost:5432",
		},
		{
			name:     "undefined variable stays unchanged",
			content:  "{{ env.undefined }}",
			expected: "{{ env.undefined }}",
		},
		{
			name:     "mixed defined and undefined",
			content:  "{{ env.api.url }} and {{ env.missing }}",
			expected: "https://api.example.com and {{ env.missing }}",
		},
		{
			name:     "with spaces",
			content:  "{{  env.api.url  }}",
			expected: "https://api.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceEnvVariables(tt.content)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestResetEnvVariables(t *testing.T) {
	SetEnvVariables(map[string]string{
		"var1": "value1",
		"var2": "value2",
	})

	all := GetAllEnvVariables()
	assert.Equal(t, 2, len(all))

	ResetEnvVariables()

	all = GetAllEnvVariables()
	assert.Empty(t, all)

	val, ok := GetEnvVariable("var1")
	assert.False(t, ok)
	assert.Equal(t, "", val)
}
