package assertions

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) elementShouldNotExist() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the {string} should not exist$`},
		func(ctx *scenario.Context) func(string) error {
			return func(name string) error {
				page, pageName := ctx.GetCurrentPage()
				_, err := browser.GetElementByLabel(page, pageName, name)
				if err == nil {
					return fmt.Errorf("%s exists but should not", name)
				}

				return nil
			}
		},
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the element is not present in the DOM.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the submit button should not exist",
			Category: stepbuilder.Visual,
		},
	)
}
