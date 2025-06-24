package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
)

func (s steps) selectRadioButton() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "radio_button")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user selects the {string} radio button$`},
		func(ctx *scenario.Context) func(string) error {
			return func(radioBtnName string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				radioButton, err := browser.GetElementByLabel(currentPage, pageName, formatLabel(radioBtnName))
				if err != nil {
					return err
				}

				if radioButton.IsChecked() {
					logger.Warn("Radio button already selected", []string{})
					return nil
				}

				return radioButton.Click()
			}
		},
		func(radioBtnName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(radioBtnName)
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
