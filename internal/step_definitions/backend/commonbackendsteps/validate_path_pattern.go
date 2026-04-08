package commonbackendsteps

import (
	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) validateJSONPathPattern() stepbuilder.Step {
	return s.newJSONPathPatternStep(
		`the response field {string} should match pattern {string}`,
		true,
		stepbuilder.DocParams{
			Description: "Validates that a specific response path field matches a regular expression pattern.",
			Variables: []stepbuilder.DocVariable{
				responsePathDocVariable("Response path to the field to validate (GJSON for JSON, XPath for XML)"),
				{Name: "pattern", Description: "Regular expression pattern to match against", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.email" should match pattern "^[a-z]+@example\\.com$"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) validateJSONPathNotPattern() stepbuilder.Step {
	return s.newJSONPathPatternStep(
		`the response field {string} should not match pattern {string}`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that a specific response path field does not match a regular expression pattern.",
			Variables: []stepbuilder.DocVariable{
				responsePathDocVariable("Response path to the field to validate (GJSON for JSON, XPath for XML)"),
				{Name: "pattern", Description: "Regular expression pattern that should not match", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.email" should not match pattern ".*@internal\\.com$"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) newJSONPathPatternStep(
	sentence string,
	shouldMatch bool,
	doc stepbuilder.DocParams,
) stepbuilder.Step {
	return s.newJSONPathStringStep(sentence, shouldMatch, doc, validation.ValidatePathPattern)
}
