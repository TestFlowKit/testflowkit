package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (s steps) userUnchecksCheckbox() stepbuilder.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user unchecks the {string} checkbox$`},
		func(ctx *stepbuilder.TestSuiteContext) func(string) error {
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
		stepbuilder.StepDefDocParams{
			Description: "Deselects or unticks a checkbox element identified by its logical name",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "the logical name of the checkbox", Type: shared.DocVarTypeString},
			},
			Example:  `When the user unchecks the "Subscribe to newsletter" checkbox`,
			Category: shared.Form,
		},
	)
}
