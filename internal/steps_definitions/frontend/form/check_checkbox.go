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

func (s steps) userChecksCheckbox() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user checks the {string} checkbox$`},
		func(ctx *scenario.Context) func(string) error {
			return func(checkBoxName string) error {
				checkBox, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(checkBoxName))
				if err != nil {
					return err
				}

				if checkBox.IsChecked() {
					logger.Warn(fmt.Sprintf("%s checkbox is already checked", checkBoxName), []string{})
					return nil
				}

				return checkBox.Click()
			}
		},
		func(checkboxName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(checkboxName)
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "Selects or ticks a checkbox element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "checkbox logical name", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user checks the "Agree to terms" checkbox`,
			Category: stepbuilder.Form,
		},
	)
}
