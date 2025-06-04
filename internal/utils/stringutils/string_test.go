package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_SplitAndTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expected []string
	}{
		{
			name:     "simple string with comma separator",
			input:    "a,b,c",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "string with spaces to be trimmed",
			input:    " a , b , c ",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "string with different separator",
			input:    "a|b|c",
			sep:      "|",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty string",
			input:    "",
			sep:      ",",
			expected: []string{""},
		},
		{
			name:     "string with only separator",
			input:    ",",
			sep:      ",",
			expected: []string{"", ""},
		},
		{
			name:     "string with mixed whitespace",
			input:    "  first item  ,second item,   third item  ",
			sep:      ",",
			expected: []string{"first item", "second item", "third item"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitAndTrim(tt.input, tt.sep)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestString_AddSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffix   string
		expected string
	}{
		{
			name:     "simple string with suffix",
			input:    "button",
			suffix:   "primary",
			expected: "button_primary",
		},
		{
			name:     "string with spaces to be trimmed",
			input:    " button ",
			suffix:   "primary",
			expected: "button_primary",
		},
		{
			name:     "string with different suffix",
			input:    "button",
			suffix:   "secondary",
			expected: "button_secondary",
		},
		{
			name:     "empty string",
			input:    "",
			suffix:   "primary",
			expected: "_primary",
		},
		{
			name:     "string with only suffix",
			input:    "",
			suffix:   "primary",
			expected: "_primary",
		},
		{
			name:     "string with mixed whitespace",
			input:    "  button  ",
			suffix:   "primary",
			expected: "button_primary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SuffixWithUnderscore(tt.input, tt.suffix)
			assert.Equal(t, tt.expected, result)
		})
	}
}
