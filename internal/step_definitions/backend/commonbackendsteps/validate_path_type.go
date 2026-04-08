package commonbackendsteps

import (
	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathType validates that a JSON path field has a specific type.
func (s steps) validateJSONPathType() stepbuilder.Step {
	return s.newJSONPathTypeStep(
		`the response field {string} should have type {string}`,
		true,
		func() stepbuilder.DocParams {
			tdsc := "Expected type (string/text, number, integer, boolean, object, array/list, null)"
			vars := []stepbuilder.DocVariable{
				responsePathDocVariable("Response path to the field to validate (GJSON for JSON, XPath for XML)"),
				{Name: "type", Description: tdsc, Type: stepbuilder.VarTypeString},
			}
			return stepbuilder.DocParams{
				Description: "Validates that a specific response path field has the expected type.",
				Variables:   vars,
				Example:     `Then the response field "user.age" should have type "integer"`,
				Categories:  stepbuilder.Backend,
			}
		}(),
	)
}

// validateJSONPathNotType validates that a JSON path field does not have a specific type.
func (s steps) validateJSONPathNotType() stepbuilder.Step {
	return s.newJSONPathTypeStep(
		`the response field {string} should not have type {string}`,
		false,
		func() stepbuilder.DocParams {
			tdsc := "Type that must not match (string/text, number, integer, boolean, object, array/list, null)"
			vars := []stepbuilder.DocVariable{
				responsePathDocVariable("Response path to the field to validate (GJSON for JSON, XPath for XML)"),
				{Name: "type", Description: tdsc, Type: stepbuilder.VarTypeString},
			}
			return stepbuilder.DocParams{
				Description: "Validates that a specific response path field does not have the specified type.",
				Variables:   vars,
				Example:     `Then the response field "user.id" should not have type "string"`,
				Categories:  stepbuilder.Backend,
			}
		}(),
	)
}

func (s steps) newJSONPathTypeStep(
	sentence string,
	shouldMatchType bool,
	doc stepbuilder.DocParams,
) stepbuilder.Step {
	return s.newJSONPathStringStep(sentence, shouldMatchType, doc, validation.ValidatePathType)
}
