package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) elementShouldBeVisible() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the {string} should be visible$`},
		func(ctx *stepbuilder.TestSuiteContext) func(string) error {
			return func(name string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
				}

				if !element.IsVisible() {
					return fmt.Errorf("%s is not visible", name)
				}

				return nil
			}
		},
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "This assertion checks if the element is present in the DOM and displayed on the page.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then the submit button should be visible",
			Category: shared.Visual,
		},
	)
}
