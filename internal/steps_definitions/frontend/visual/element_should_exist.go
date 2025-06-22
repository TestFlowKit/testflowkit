package visual

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) elementShouldExist() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
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
		stepbuilder.DocParams{
			Description: "This assertion checks if the element is present in the DOM, regardless of its visibility.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the submit button should exist",
			Category: stepbuilder.Visual,
		},
	)
}
