package gherkinparser

import (
	"testing"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/stretchr/testify/assert"
)

func TestApplyMacros_ReplacesMacroStepsWithStepsFromMacro(t *testing.T) {
	featuresContainingMacros := []*Feature{
		{
			Contents: []byte("Scenario: Test scenario\nGiven a step\nWhen a macro step\nThen a result"),
			scenarios: []*scenario{
				{
					Steps: []*step{
						{Keyword: "Given", Text: "a step"},
						{Keyword: "When", Text: "a macro step", Location: &messages.Location{Line: 3}},
						{Keyword: "Then", Text: "a result"},
					},
				},
			},
		},
	}
	macros := []*scenario{
		{
			Name: "a macro step",
			Steps: []*step{
				{Keyword: "Given", Text: "macro step 1"},
				{Keyword: "And", Text: "macro step 2"},
			},
		},
	}

	applyMacros(macros, featuresContainingMacros)

	expectedContent := "Scenario: Test scenario\nGiven a step\nWhen macro step 1\nAnd macro step 2\nThen a result"
	assert.Equal(t, expectedContent, string(featuresContainingMacros[0].Contents))
}

func TestApplyMacros_NoReplacementWhenNoMacroSteps(t *testing.T) {
	featuresContainingMacros := []*Feature{
		{
			Contents: []byte("Scenario: Test scenario\nGiven a step\nWhen another step\nThen a result"),
			scenarios: []*scenario{
				{
					Steps: []*step{
						{Keyword: "Given", Text: "a step"},
						{Keyword: "When", Text: "another step"},
						{Keyword: "Then", Text: "a result"},
					},
				},
			},
		},
	}
	macros := []*scenario{
		{
			Name: "a macro step",
			Steps: []*step{
				{Keyword: "Given", Text: "macro step 1"},
				{Keyword: "And", Text: "macro step 2"},
			},
		},
	}

	applyMacros(macros, featuresContainingMacros)

	expectedContent := "Scenario: Test scenario\nGiven a step\nWhen another step\nThen a result"
	assert.Equal(t, expectedContent, string(featuresContainingMacros[0].Contents))
}

func TestApplyMacros_EmptyFeaturesContainingMacros(t *testing.T) {
	featuresContainingMacros := []*Feature{}
	macros := []*scenario{
		{
			Name: "a macro step",
			Steps: []*step{
				{Keyword: "Given", Text: "macro step 1"},
				{Keyword: "And", Text: "macro step 2"},
			},
		},
	}

	applyMacros(macros, featuresContainingMacros)

	assert.Empty(t, featuresContainingMacros)
}

func TestApplyMacros_EmptyMacroTitles(t *testing.T) {
	featuresContainingMacros := []*Feature{
		{
			Contents: []byte("Scenario: Test scenario\nGiven a step\nWhen a macro step\nThen a result"),
			scenarios: []*scenario{
				{
					Steps: []*step{
						{Keyword: "Given", Text: "a step"},
						{Keyword: "When", Text: "a macro step"},
						{Keyword: "Then", Text: "a result"},
					},
				},
			},
		},
	}
	var macros []*scenario

	applyMacros(macros, featuresContainingMacros)

	expectedContent := "Scenario: Test scenario\nGiven a step\nWhen a macro step\nThen a result"
	assert.Equal(t, expectedContent, string(featuresContainingMacros[0].Contents))
}

func TestApplyMacros_ReplacesMacroStepsInBackground(t *testing.T) {
	featuresContainingMacros := []*Feature{
		{
			Contents: []byte(`Feature: Test feature

Background:
Given a background step
When a macro step

Scenario: Test scenario
Given a step
Then a result`),
			scenarios: []*scenario{
				{
					Steps: []*step{
						{Keyword: "Given", Text: "a step"},
						{Keyword: "Then", Text: "a result"},
					},
				},
			},
			background: &messages.Background{
				Steps: []*step{
					{Keyword: "Given", Text: "a background step", Location: &messages.Location{Line: 4}},
					{Keyword: "When", Text: "a macro step", Location: &messages.Location{Line: 5}},
				},
			},
		},
	}
	macros := []*scenario{
		{
			Name: "a macro step",
			Steps: []*step{
				{Keyword: "Given", Text: "macro step 1"},
				{Keyword: "And", Text: "macro step 2"},
			},
		},
	}

	applyMacros(macros, featuresContainingMacros)

	expectedContent := `Feature: Test feature

Background:
Given a background step
When macro step 1
And macro step 2

Scenario: Test scenario
Given a step
Then a result`
	assert.Equal(t, expectedContent, string(featuresContainingMacros[0].Contents))
}

func TestApplyMacros_NoReplacementInBackgroundWhenNoMacroSteps(t *testing.T) {
	featuresContainingMacros := []*Feature{
		{
			Contents: []byte(`Feature: Test feature

Background:
Given a background step
When another step

Scenario: Test scenario
Given a step
Then a result`),
			scenarios: []*scenario{
				{
					Steps: []*step{
						{Keyword: "Given", Text: "a step"},
						{Keyword: "Then", Text: "a result"},
					},
				},
			},
			background: &messages.Background{
				Steps: []*step{
					{Keyword: "Given", Text: "a background step", Location: &messages.Location{Line: 4}},
					{Keyword: "When", Text: "another step", Location: &messages.Location{Line: 5}},
				},
			},
		},
	}
	macros := []*scenario{
		{
			Name: "a macro step",
			Steps: []*step{
				{Keyword: "Given", Text: "macro step 1"},
				{Keyword: "And", Text: "macro step 2"},
			},
		},
	}

	applyMacros(macros, featuresContainingMacros)

	expectedContent := `Feature: Test feature

Background:
Given a background step
When another step

Scenario: Test scenario
Given a step
Then a result`
	assert.Equal(t, expectedContent, string(featuresContainingMacros[0].Contents))
}

type step = messages.Step
