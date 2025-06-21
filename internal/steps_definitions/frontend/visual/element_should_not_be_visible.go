package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) elementShouldNotBeVisible() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the {string} should not be visible$`},
		func(ctx *scenario.Context) func(string) error {
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
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "This assertion checks if the element is present in the DOM but not displayed",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then \"Submit button\" should not be visible",
			Category: shared.Visual,
		},
	)
}
