package commonbackendsteps

import (
	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathValue validates that a JSON path has a specific value.
func (s steps) validateJSONPathValue() stepbuilder.Step {
	return s.newJSONPathValueStep(
		`the response field {string} should be {string}`,
		true,
		newResponseFieldDocParams(
			"Validates that a specific response path has the expected value.",
			"Response path to validate (GJSON for JSON, XPath for XML)",
			"value",
			"Expected value at the path",
			`Then the response field "user.name" should be "John Doe"`,
		),
	)
}

// validateJSONPathValueNot validates that a JSON path does not have a specific value.
func (s steps) validateJSONPathValueNot() stepbuilder.Step {
	return s.newJSONPathValueStep(
		`the response field {string} should not be {string}`,
		false,
		newResponseFieldDocParams(
			"Validates that a specific response path does not have the specified value.",
			"Response path to validate (GJSON for JSON, XPath for XML)",
			"value",
			"Value that should not be at the path",
			`Then the response field "user.role" should not be "admin"`,
		),
	)
}

func (s steps) newJSONPathValueStep(
	sentence string,
	shouldEqual bool,
	doc stepbuilder.DocParams,
) stepbuilder.Step {
	return s.newJSONPathStringStep(sentence, shouldEqual, doc, validation.ValidatePathValue)
}
