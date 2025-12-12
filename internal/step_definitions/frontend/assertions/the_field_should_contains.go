package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) theFieldShouldContain() stepbuilder.Step {
	formatFieldID := func(fieldId string) string {
		return fieldId + "_field"
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`the value of the {string} field should be {string}`},
		func(ctx context.Context, fieldId, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			input, err := scenarioCtx.GetHTMLElementByLabel(formatFieldID(fieldId))
			if err != nil {
				return ctx, err
			}

			if input.TextContent() == text {
				return ctx, nil
			}

			return ctx, fmt.Errorf(`field should be contains "%s" but contains "%s"`, text, input.TextContent())
		},
		func(fieldId, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(formatFieldID(fieldId)) {
				vc.AddMissingElement(formatFieldID(fieldId))
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the current value of an input field matches the specified value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "fieldId", Description: "The id of the field.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the value of the "Username" field should be "myUsername".`,
			Category: stepbuilder.Form,
		},
	)
}
