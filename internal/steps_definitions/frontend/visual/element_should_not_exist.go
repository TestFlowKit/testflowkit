package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) elementShouldNotExist() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{`^the {string} should not exist$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(name string) error {
				_, err := browser.GetElementByLabel(ctx.GetCurrentPage(), name)
				if err == nil {
					return fmt.Errorf("%s exists but should not", name)
				}

				return nil
			}
		},
		func(name string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		core.StepDefDocParams{
			Description: "This assertion checks if the element is not present in the DOM.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then the submit button should not exist",
			Category: shared.Visual,
		},
	)
}
