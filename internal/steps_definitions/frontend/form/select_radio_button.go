package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
)

func (steps) selectRadioButton() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "radio_button")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`the user selects the {string} radio button`},
		func(ctx context.Context, radioName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			radio, err := scenarioCtx.GetHTMLElementByLabel(formatLabel(radioName))
			if err != nil {
				return ctx, err
			}

			if radio.IsChecked() {
				logger.Warn("Radio button already selected", []string{})
				return ctx, nil
			}

			err = radio.Click()
			return ctx, err
		},
		func(radioName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(radioName)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "Selects a radio button by its logical name.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the radio button.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user selects the "Gender Male" radio button`,
			Category: stepbuilder.Form,
		},
	)
}
