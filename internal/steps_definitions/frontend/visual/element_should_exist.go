package visual

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) elementShouldExist() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the {string} should exist$`},
		func(ctx *scenario.Context) func(string) error {
			return func(name string) error {
				_, err := browser.GetElementByLabel(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
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
			Description: "This assertion checks if the element is present in the DOM, regardless of its visibility.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then the submit button should exist",
			Category: shared.Visual,
		},
	)
}
