package gherkinparser

import (
	"strings"
	"testing"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/stretchr/testify/assert"
)

func Test_GetMacroTitles(t *testing.T) {
	tests := []struct {
		name     string
		macros   []*scenario
		expected []string
	}{
		{
			name: "returns titles from multiple macros",
			macros: []*scenario{
				{Name: "first macro"},
				{Name: "second macro"},
				{Name: "third macro"},
			},
			expected: []string{"first macro", "second macro", "third macro"},
		},
		{
			name:     "returns empty slice for empty macros",
			macros:   []*scenario{},
			expected: nil,
		},
		{
			name: "handles macros with empty names",
			macros: []*scenario{
				{Name: ""},
				{Name: "named macro"},
			},
			expected: []string{"", "named macro"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMacroTitles(tt.macros)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_GetMacros(t *testing.T) {
	tests := []struct {
		name          string
		macroFeatures []*Feature
		expected      int // expected number of scenarios
	}{
		{
			name: "extracts scenarios from multiple features",
			macroFeatures: []*Feature{
				{
					scenarios: []*scenario{
						{Name: "first macro", Tags: []*messages.Tag{{Name: MacroTag}}},
						{Name: "second macro", Tags: []*messages.Tag{{Name: MacroTag}}},
						{Name: "none macro"},
					},
				},
				{
					scenarios: []*scenario{
						{Name: "third macro", Tags: []*messages.Tag{{Name: MacroTag}}},
						{Name: "simple scenario", Tags: []*messages.Tag{{Name: "Test"}}},
					},
				},
			},
			expected: 3,
		},
		{
			name:          "returns empty slice for empty features",
			macroFeatures: []*Feature{},
			expected:      0,
		},
		{
			name: "handles features with empty scenarios",
			macroFeatures: []*Feature{
				{
					scenarios: []*scenario{},
				},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMacros(tt.macroFeatures)
			assert.Len(t, result, tt.expected)
		})
	}
}

func Test_ApplyMacro(t *testing.T) {
	tests := []struct {
		name           string
		scenario       *scenario
		macroTitles    []string
		macros         []*scenario
		featureContent []string
		expected       []string
	}{
		{
			name: "replaces single macro step",
			scenario: &scenario{
				Steps: []*step{
					{Keyword: "Given", Text: "a step"},
					{Keyword: "When", Text: "a macro step", Location: &messages.Location{Line: 3}},
					{Keyword: "Then", Text: "a result"},
				},
			},
			macroTitles: []string{"a macro step"},
			macros: []*scenario{
				{
					Name: "a macro step",
					Steps: []*step{
						{Keyword: "Given", Text: "macro step 1"},
						{Keyword: "And", Text: "macro step 2"},
					},
				},
			},
			featureContent: []string{
				"Scenario: Test scenario",
				"Given a step",
				"When a macro step",
				"Then a result",
			},
			expected: []string{
				"Scenario: Test scenario",
				"Given a step",
				"When macro step 1",
				"And macro step 2",
				"Then a result",
			},
		},
		{
			name: "handles scenario with no macro steps",
			scenario: &scenario{
				Steps: []*step{
					{Keyword: "Given", Text: "a step"},
					{Keyword: "When", Text: "another step", Location: &messages.Location{Line: 3}},
					{Keyword: "Then", Text: "a result"},
				},
			},
			macroTitles: []string{"a macro step"},
			macros: []*scenario{
				{
					Name: "a macro step",
					Steps: []*step{
						{Keyword: "Given", Text: "macro step 1"},
					},
				},
			},
			featureContent: []string{
				"Scenario: Test scenario",
				"Given a step",
				"When another step",
				"Then a result",
			},
			expected: []string{
				"Scenario: Test scenario",
				"Given a step",
				"When another step",
				"Then a result",
			},
		},
		{
			name: "handles multiple macro steps",
			scenario: &scenario{
				Steps: []*step{
					{Keyword: "Given", Text: "first macro", Location: &messages.Location{Line: 2}},
					{Keyword: "When", Text: "second macro", Location: &messages.Location{Line: 3}},
					{Keyword: "Then", Text: "a result"},
				},
			},
			macroTitles: []string{"first macro", "second macro"},
			macros: []*scenario{
				{
					Name: "first macro",
					Steps: []*step{
						{Keyword: "Given", Text: "first macro step 1"},
					},
				},
				{
					Name: "second macro",
					Steps: []*step{
						{Keyword: "When", Text: "second macro step 1"},
						{Keyword: "And", Text: "second macro step 2"},
					},
				},
			},
			featureContent: []string{
				"Scenario: Test scenario",
				"Given first macro",
				"When second macro",
				"Then a result",
			},
			expected: []string{
				"Scenario: Test scenario",
				"Given first macro step 1",
				"When second macro step 1",
				"And second macro step 2",
				"Then a result",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the feature content for testing
			testContent := make([]string, len(tt.featureContent))
			copy(testContent, tt.featureContent)

			applyMacro(tt.scenario.Steps, tt.macros, &testContent)
			assert.Equal(t, tt.expected, testContent)
		})
	}
}

func Test_ApplyMacro_WithDocstringAndDatatable(t *testing.T) {
	tests := []struct {
		name           string
		scenario       *messages.Scenario
		macros         []*messages.Scenario
		featureContent []string
		expected       []string
	}{
		{
			name: "replaces macro step with only docstring",
			scenario: &messages.Scenario{
				Steps: []*messages.Step{
					{Keyword: "Given", Text: "a step"},
					{Keyword: "When", Text: "a macro step with doc", Location: &messages.Location{Line: 3}},
					{Keyword: "Then", Text: "a result"},
				},
			},
			macros: []*messages.Scenario{
				{
					Name: "a macro step with doc",
					Steps: []*messages.Step{
						{Keyword: "Given", Text: "macro step 1"},
						{Keyword: "And", Text: "macro step 2", DocString: &messages.DocString{
							Content: "This is a docstring",
						}},
					},
				},
			},
			featureContent: []string{
				"Scenario: Test scenario",
				"Given a step",
				"When a macro step with doc",
				"Then a result",
			},
			expected: []string{
				"Scenario: Test scenario",
				"Given a step",
				"When macro step 1",
				"And macro step 2",
				"\"\"\"",
				"This is a docstring",
				"\"\"\"",
				"Then a result",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the feature content for testing
			testContent := make([]string, len(tt.featureContent))
			copy(testContent, tt.featureContent)

			applyMacro(tt.scenario.Steps, tt.macros, &testContent)

			// Trim any trailing empty lines for comparison
			for len(testContent) > 0 && strings.TrimSpace(testContent[len(testContent)-1]) == "" {
				testContent = testContent[:len(testContent)-1]
			}

			assert.Equal(t, tt.expected, testContent)
		})
	}
}
