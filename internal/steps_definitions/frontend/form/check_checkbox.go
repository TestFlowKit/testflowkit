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

func (s steps) userChecksCheckbox() core.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	doc := core.StepDefDocParams{
		Description: "Selects or ticks a checkbox element identified by its logical name",
		Variables: []shared.StepVariable{
			{Name: "name", Description: "checkbox logical name", Type: shared.DocVarTypeString},
		},
		Example:  `When the user checks the "Agree to terms" checkbox`,
		Category: shared.Form,
	}

	return core.NewStepWithOneVariable(
		[]string{`^the user checks the {string} checkbox$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(checkboxName string) error {
				checkBox, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(checkboxName))
				if err != nil {
					return err
				}

				if checkBox.IsChecked() {
					logger.Warn(fmt.Sprintf("%s checkbox is already checked", checkboxName), []string{})
					return nil
				}

				return checkBox.Click()
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
