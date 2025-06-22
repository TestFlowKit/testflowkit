package assertions

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) elementShouldContainExactText() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the text of the {string} element should be exactly {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(name, text string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
				}

				if !element.IsVisible() {
					return fmt.Errorf("%s is not visible", name)
				}

				if element.TextContent() != text {
					return fmt.Errorf("%s does not contain exact text '%s'", name, text)
				}

				return nil
			}
		},
		func(name, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the element’s visible text is an exact match to the specified string.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the text of the "Welcome Message" element should be exactly "Hello John".`,
			Category: stepbuilder.Visual,
		},
	)
}
