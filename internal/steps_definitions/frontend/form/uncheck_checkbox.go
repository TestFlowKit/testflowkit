package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
)

func (s steps) userUnchecksCheckbox() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user unchecks the {string} checkbox$`},
		func(ctx *scenario.Context) func(string) error {
			return func(checkBoxName string) error {
				checkBox, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(checkBoxName))
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
		func(checkboxId string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(checkboxId)
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "Deselects or unticks a checkbox element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "the logical name of the checkbox", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user unchecks the "Subscribe to newsletter" checkbox`,
			Category: stepbuilder.Form,
		},
	)
}
