package assertions

import (
	"fmt"
	"strings"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) elementShouldContainText() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the {string} should contain the text {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(name, text string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
				}

				if !element.IsVisible() {
					return fmt.Errorf("%s is not visible", name)
				}

				if !strings.Contains(element.TextContent(), text) {
					return fmt.Errorf("%s does not contain text '%s'", name, text)
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
			Description: "This assertion checks if the element’s visible text includes the specified substring.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the welcome card should contain the text "Hello John"`,
			Category: stepbuilder.Visual,
		},
	)
}
