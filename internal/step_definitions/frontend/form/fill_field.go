package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) userEntersTextIntoField() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "field")
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`the user enters {string} into the {string} field`},
		func(ctx context.Context, text, inputLabel string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			input, err := scenarioCtx.GetHTMLElementByLabel(formatLabel(inputLabel))
			if err != nil {
				return ctx, err
			}
			err = input.Input(text)
			return ctx, err
		},
		func(_, inputLabel string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(inputLabel)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "Types the specified text into an input field identified by its logical name.",
			Variables: []stepbuilder.DocVariable{
				{Name: "text", Description: "The text to type.", Type: stepbuilder.VarTypeString},
				{Name: "name", Description: "The logical name of the input field.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user enters "myUsername" into the "Username" field`,
			Category: stepbuilder.Form,
		},
	)
}
