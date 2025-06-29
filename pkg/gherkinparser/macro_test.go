package gherkinparser

import (
	"testing"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/stretchr/testify/assert"
)

func TestGetMacroTitles(t *testing.T) {
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

func TestGetMacros(t *testing.T) {
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
						{Name: "first macro"},
						{Name: "second macro"},
					},
				},
				{
					scenarios: []*scenario{
						{Name: "third macro"},
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

func TestApplyMacro(t *testing.T) {
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
				"When macro step 1\nAnd macro step 2",
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
				"When second macro step 1\nAnd second macro step 2",
				"Then a result",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			applyMacro(tt.scenario.Steps, tt.macroTitles, tt.macros, tt.featureContent)
			assert.Equal(t, tt.expected, tt.featureContent)
		})
	}
}
