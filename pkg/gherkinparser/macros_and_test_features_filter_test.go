package gherkinparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeparateMacrosAndTestFeatures_WithMacroAndTestFeatures(t *testing.T) {
	features := []*Feature{
		{uri: "mac1.macro.feature"},
		{uri: "mac2.macro.feature"},
		{uri: "feature2.feature"},
	}

	macroFeatures, testFeatures := separateMacroAndTestsFeatures(features)

	assert.Len(t, macroFeatures, 2)
	assert.Len(t, testFeatures, 1)
}

func TestSeparateMacrosAndTestFeatures_EmptyFeatures(t *testing.T) {
	var features []*Feature

	macroFeatures, testFeatures := separateMacroAndTestsFeatures(features)
	assert.Empty(t, macroFeatures)
	assert.Empty(t, testFeatures)
}

func TestSeparateFeaturesContainingMacrosOrNot_WithMacroTitles(t *testing.T) {
	testFeatures := []*Feature{
		{Contents: []byte("Scenario: Test scenario 1\nGiven a step\nWhen a macro step\nThen a result")},
		{Contents: []byte("Scenario: Test scenario 2\nGiven a step\nWhen another step\nThen a result")},
		{Contents: []byte("Scenario: Test scenario 3\nGiven a step\nWhen another step\nThen a result")},
	}
	macros := []*scenario{{Name: "a macro step"}}

	featuresContainingMacros, featuresWithoutMacros := separateFeaturesContainingMacrosOrNot(macros, testFeatures)

	assert.Len(t, featuresContainingMacros, 1)
	assert.Len(t, featuresWithoutMacros, 2)
}

func TestSeparateFeaturesContainingMacrosOrNot_Empty(t *testing.T) {
	testFeatures := []*Feature{
		{Contents: []byte("Scenario: Test scenario 1\nGiven a step\nWhen a step\nThen a result")},
		{Contents: []byte("Scenario: Test scenario 2\nGiven a step\nWhen another step\nThen a result")},
	}

	var macros []*scenario

	featuresContainingMacros, featuresWithoutMacros := separateFeaturesContainingMacrosOrNot(macros, testFeatures)

	assert.Empty(t, featuresContainingMacros)
	assert.Len(t, featuresWithoutMacros, 2)
}
