package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (s steps) userSelectsRadioButton() stepbuilder.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "radio_button")
	}

	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user selects the {string} radio button$`},
		func(ctx *scenario.Context) func(string) error {
			return func(radioBtnName string) error {
				radioButton, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(radioBtnName))
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
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Selects a radio button by its logical name.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the radio button.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user selects the "Gender Male" radio button`,
			Category: shared.Form,
		},
	)
}
