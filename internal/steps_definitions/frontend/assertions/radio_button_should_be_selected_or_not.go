package assertions

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) radioButtonShouldBeSelectedOrNot() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "radio_button")
	}

	definition := func(ctx *scenario.Context) func(string, string) error {
		return func(radioId, expectedState string) error {
			currentPage, pageName := ctx.GetCurrentPage()
			input, err := browser.GetElementByLabel(currentPage, pageName, formatLabel(radioId))
			if err != nil {
				return err
			}

			isSelected := input.IsChecked()
			expectedSelected := expectedState == "selected"

			if isSelected != expectedSelected {
				return fmt.Errorf("radio button %s is %s but should be %s", radioId, getRadioState(isSelected), expectedState)
			}

			return nil
		}
	}

	validator := func(radioBtnName, _ string) stepbuilder.ValidationErrors {
		vc := stepbuilder.ValidationErrors{}
		radioLabel := formatLabel(radioBtnName)

		if !config.IsElementDefined(radioLabel) {
			vc.AddMissingElement(radioLabel)
		}

		return vc
	}

	statusType := stepbuilder.VarTypeEnum("selected", "unselected")

	return stepbuilder.NewWithTwoVariables(
		[]string{`the {string} radio button should be (selected|unselected)`},
		definition,
		validator,
		stepbuilder.DocParams{
			Description: "checks if a radio button is in the expected state (selected or not selected).",
			Variables: []stepbuilder.DocVariable{
				{Name: "radio button", Description: "The name of the radio button.", Type: stepbuilder.VarTypeString},
				{Name: "status", Description: "The status of the radio button.", Type: statusType},
			},
			Example:  `Then the "terms" radio button should be selected`,
			Category: stepbuilder.Form,
		},
	)
}

func getRadioState(isSelected bool) string {
	if isSelected {
		return "selected"
	}
	return "not selected"
}
