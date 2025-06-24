package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
)

func (s steps) uncheckCheckbox() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user unchecks the {string} checkbox$`},
		func(ctx *scenario.Context) func(string) error {
			return func(checkBoxName string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				checkBox, err := browser.GetElementByLabel(currentPage, pageName, formatLabel(checkBoxName))
				if err != nil {
					return err
				}

				if checkBox.IsChecked() {
					return checkBox.Click()
				}

				logger.Warn(fmt.Sprintf("%s checkbox is not unchecked because it is already unchecked", checkBoxName), []string{})
				return nil
			}
		},
		func(checkBoxName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(checkBoxName)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "unchecks a checkbox if it is currently checked.",
			Variables: []stepbuilder.DocVariable{
				{Name: "checkBoxName", Description: "The name of the checkbox to uncheck.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user unchecks the \"Newsletter\" checkbox",
			Category: stepbuilder.Form,
		},
	)
}
