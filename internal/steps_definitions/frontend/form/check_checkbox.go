package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) checkCheckbox() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user checks the {string} checkbox$`},
		func(ctx *scenario.Context) func(string) error {
			return func(checkBoxName string) error {
				page, pageName := ctx.GetCurrentPage()
				checkBox, err := browser.GetElementByLabel(page, pageName, formatLabel(checkBoxName))
				if err != nil {
					return err
				}

				if !checkBox.IsChecked() {
					return checkBox.Click()
				}

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
			Description: "checks a checkbox if it is not already checked.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "checkbox logical name", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user checks the \"Terms\" checkbox",
			Category: stepbuilder.Form,
		},
	)
}
