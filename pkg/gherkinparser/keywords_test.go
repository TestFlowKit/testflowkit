package gherkinparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepKeywordsByID(t *testing.T) {
	content := []byte("Feature: f\n  Background:\n    Given a\n  Scenario: s\n    When b\n    And c\n    Then d\n")

	keywords := StepKeywordsByID(content)

	assert.ElementsMatch(t, []string{"Given", "When", "And", "Then"}, valuesOf(keywords))
}

func valuesOf(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}
