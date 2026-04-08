package commonbackendsteps

import (
	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathContains validates that a JSON path field contains specific text.
func (s steps) validateJSONPathContains() stepbuilder.Step {
	return s.newJSONPathContainsStep(
		`the response field "{string}" should contain "{string}"`,
		true,
		newResponseFieldDocParams(
			"Validates that a specific response path field contains the expected text (substring match).",
			"Response path to the field to validate (GJSON for JSON, XPath for XML)",
			"text",
			"Text that should be contained in the field value",
			`Then the response field "user.email" should contain "@example.com"`,
		),
	)
}

// validateJSONPathNotContains validates that a JSON path field does not contain specific text.
func (s steps) validateJSONPathNotContains() stepbuilder.Step {
	return s.newJSONPathContainsStep(
		`the response field "{string}" should not contain "{string}"`,
		false,
		newResponseFieldDocParams(
			"Validates that a specific response path field does not contain the specified text.",
			"Response path to the field to validate (GJSON for JSON, XPath for XML)",
			"text",
			"Text that should not be present in the field value",
			`Then the response field "user.email" should not contain "@internal"`,
		),
	)
}

func (s steps) newJSONPathContainsStep(
	sentence string,
	shouldContain bool,
	doc stepbuilder.DocParams,
) stepbuilder.Step {
	return s.newJSONPathStringStep(sentence, shouldContain, doc, validation.ValidatePathContains)
}
