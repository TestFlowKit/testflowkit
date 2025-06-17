package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (s steps) userSelectsRadioButton() core.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "radio_button")
	}

	doc := core.StepDefDocParams{
		Description: "Selects a radio button option identified by its logical name from a radio button group.",
		Variables: []shared.StepVariable{
			{Name: "name", Description: "the logical name of the radio button", Type: shared.DocVarTypeString},
		},
		Example:  `When the user selects the "Standard Delivery" radio button`,
		Category: shared.Form,
	}

	return core.NewStepWithOneVariable(
		[]string{`^the user selects the {string} radio button$`},
		func(ctx *core.TestSuiteContext) func(string) error {
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
		func(radioBtnName string) core.ValidationErrors {
			vErr := core.ValidationErrors{}
			label := formatLabel(radioBtnName)
			if !testsconfig.IsElementDefined(label) {
				vErr.AddMissingElement(label)
			}

			return vErr
		},
		doc,
	)
}
