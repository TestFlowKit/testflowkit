package graphql

import "testing"

var testCases = []struct {
	name           string
	message        string
	subStringPacks [][]string
	expected       bool
}{
	{
		name:    "single pack with all substrings found",
		message: "This is a validation error message",
		subStringPacks: [][]string{
			{"validation", "error"},
		},
		expected: true,
	},
	{
		name:    "single pack with some substrings missing",
		message: "This is a validation message",
		subStringPacks: [][]string{
			{"validation", "error"},
		},
		expected: false,
	},
	{
		name:    "single pack with all substrings found in different order",
		message: "error validation occurred",
		subStringPacks: [][]string{
			{"validation", "error"},
		},
		expected: true,
	},
	{
		name:    "multiple packs, first pack matches",
		message: "field not found in the query",
		subStringPacks: [][]string{
			{"field", "not found"},
			{"validation", "error"},
		},
		expected: true,
	},
	{
		name:    "multiple packs, second pack matches",
		message: "validation error occurred",
		subStringPacks: [][]string{
			{"field", "not found"},
			{"validation", "error"},
		},
		expected: true,
	},
	{
		name:    "multiple packs, none match",
		message: "some other error message",
		subStringPacks: [][]string{
			{"field", "not found"},
			{"validation", "error"},
		},
		expected: false,
	},
	{
		name:    "multiple packs, first pack partially matches",
		message: "field exists but validation failed",
		subStringPacks: [][]string{
			{"field", "not found"},
			{"validation", "error"},
		},
		expected: false,
	},
	{
		name:    "empty message",
		message: "",
		subStringPacks: [][]string{
			{"validation", "error"},
		},
		expected: false,
	},
	{
		name:           "empty subStringPacks",
		message:        "some message",
		subStringPacks: [][]string{},
		expected:       false,
	},
	{
		name:    "empty pack in subStringPacks",
		message: "some message",
		subStringPacks: [][]string{
			{},
			{"validation", "error"},
		},
		expected: true, // Empty pack matches (all 0 substrings found)
	},
	{
		name:    "single substring in pack",
		message: "unauthorized access",
		subStringPacks: [][]string{
			{"unauthorized"},
		},
		expected: true,
	},
	{
		name:    "case sensitive matching",
		message: "Validation Error",
		subStringPacks: [][]string{
			{"validation", "error"},
		},
		expected: false, // strings.Contains is case-sensitive
	},
	{
		name:    "case sensitive matching - lowercase",
		message: "validation error",
		subStringPacks: [][]string{
			{"validation", "error"},
		},
		expected: true,
	},
	{
		name:    "substrings as part of other words",
		message: "invalid token provided",
		subStringPacks: [][]string{
			{"invalid", "token"},
		},
		expected: true,
	},
	{
		name:    "overlapping substrings",
		message: "invalid password",
		subStringPacks: [][]string{
			{"invalid", "password"},
		},
		expected: true,
	},
	{
		name:    "three substrings in pack, all found",
		message: "invalid authentication token",
		subStringPacks: [][]string{
			{"invalid", "auth", "token"},
		},
		expected: true,
	},
	{
		name:    "three substrings in pack, one missing",
		message: "invalid token",
		subStringPacks: [][]string{
			{"invalid", "auth", "token"},
		},
		expected: false,
	},
	{
		name:    "multiple packs with different lengths",
		message: "field unknown",
		subStringPacks: [][]string{
			{"field", "not found"},
			{"field", "unknown"},
			{"validation"},
		},
		expected: true,
	},
	{
		name:    "substring appears multiple times",
		message: "error: validation error occurred",
		subStringPacks: [][]string{
			{"validation", "error"},
		},
		expected: true,
	},
	{
		name:    "empty string in pack",
		message: "any message",
		subStringPacks: [][]string{
			{"", "message"},
		},
		expected: true, // Empty string is always found
	},
	{
		name:    "all empty strings in pack",
		message: "any message",
		subStringPacks: [][]string{
			{"", ""},
		},
		expected: true, // All empty strings found
	},
}

func TestIsMessageContainsSubstrings(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := isMessageContainsSubstrings(tt.message, tt.subStringPacks)
			if result != tt.expected {
				t.Errorf("isMessageContainsSubstrings(%q, %v) = %v, want %v",
					tt.message, tt.subStringPacks, result, tt.expected)
			}
		})
	}
}
