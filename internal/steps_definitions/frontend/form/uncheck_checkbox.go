package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (s steps) userUnchecksCheckbox() core.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	doc := core.StepDefDocParams{
		Description: "Deselects or unticks a checkbox element identified by its logical name",
		Variables: []shared.StepVariable{
			{Name: "checkbox logical name", Description: "checkbox name", Type: shared.DocVarTypeString},
		},
		Example:  `When the user unchecks the "Subscribe to newsletter" checkbox`,
		Category: shared.Form,
	}

	return core.NewStepWithOneVariable(
		[]string{`^the user unchecks the {string} checkbox$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(checkboxName string) error {
				checkBox, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(checkboxName))
				if err != nil {
					return err
				}

				if checkBox.IsChecked() {
					return checkBox.Click()
				}
				logger.Warn(fmt.Sprintf("%s checkbox is not unchecked because it is already unchecked", checkboxName), []string{})
				return nil
			}
		},
		func(checkboxName string) core.ValidationErrors {
			vErr := core.ValidationErrors{}
			label := formatLabel(checkboxName)
			if !testsconfig.IsElementDefined(label) {
				vErr.AddMissingElement(label)
			}

			return vErr
		},
		doc,
	)
}
