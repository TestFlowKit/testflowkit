package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) theFieldValueShouldEqualsTo() stepbuilder.Step {
	return s.newFieldValueStep(
		`the value of the {string} field should be {string}`,
		true,
		stepbuilder.DocParams{
			Description: "This assertion checks if the current value of an input field matches the specified value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "fieldId", Description: "The id of the field.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text to check.", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the value of the "Username" field should be "myUsername".`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Form},
		},
	)
}

func (s steps) theFieldValueShouldNotEqualsTo() stepbuilder.Step {
	return s.newFieldValueStep(
		`the value of the {string} field should not be {string}`,
		false,
		stepbuilder.DocParams{
			Description: "This assertion checks if the current value of an input field is different from the specified value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "fieldId", Description: "The id of the field.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text that should not match.", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the value of the "Username" field should not be "admin".`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Form},
		},
	)
}

func (s steps) newFieldValueStep(sentence string, shouldEqual bool, doc stepbuilder.DocParams) stepbuilder.Step {
	formatFieldID := func(fieldId string) string {
		return fieldId + "_field"
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, fieldId, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			input, err := scenarioCtx.GetHTMLElementByLabel(formatFieldID(fieldId))
			if err != nil {
				return ctx, err
			}

			actualValue := input.InputValue()
			if shouldEqual && actualValue != text {
				return ctx, fmt.Errorf(`field should be "%s" but was "%s"`, text, actualValue)
			}

			if !shouldEqual && actualValue == text {
				return ctx, fmt.Errorf(`field value should not be "%s"`, text)
			}

			return ctx, nil
		},
		func(fieldId, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(formatFieldID(fieldId)) {
				vc.AddMissingElement(formatFieldID(fieldId))
			}

			return vc
		},
		doc,
	)
}
