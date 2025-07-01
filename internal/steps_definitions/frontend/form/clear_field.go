package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) clearField() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "field")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`the user clears the {string} field`},
		func(ctx context.Context, inputLabel string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			input, err := scenarioCtx.GetHTMLElementByLabel(formatLabel(inputLabel))
			if err != nil {
				return ctx, err
			}

			err = input.Clear()
			return ctx, err
		},
		func(inputLabel string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(inputLabel)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "clears the content of an input field.",
			Variables: []stepbuilder.DocVariable{
				{Name: "inputLabel", Description: "The label of the input field to clear.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clears the \"Username\" field",
			Category: stepbuilder.Form,
		},
	)
}
