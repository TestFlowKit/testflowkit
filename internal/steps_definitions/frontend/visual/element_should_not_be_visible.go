package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) elementShouldNotBeVisible() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{`^the {string} should not be visible$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(name string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
				}

				if element.IsVisible() {
					return fmt.Errorf("%s is visible", name)
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
			Description: "This assertion checks if the element is present in the DOM but not displayed",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then \"Submit button\" should not be visible",
			Category: shared.Visual,
		},
	)
}
