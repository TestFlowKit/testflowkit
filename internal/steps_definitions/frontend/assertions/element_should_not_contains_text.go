package assertions

import (
	"fmt"
	"strings"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) elementShouldNotContainsText() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the {string} should not contain the text {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(name, unexpectedText string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				element, err := browser.GetElementByLabel(currentPage, pageName, name)
				if err != nil {
					return err
				}

				if !element.IsVisible() {
					return fmt.Errorf("%s is not visible", name)
				}

				if strings.Contains(element.TextContent(), unexpectedText) {
					return fmt.Errorf("%s unexpectedly contains text '%s'", name, unexpectedText)
				}

				return nil
			}
		},
		func(name, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the element’s visible text does not include the specified substring.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "unexpectedText", Description: "The text that should not be contained.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the "Welcome Message" element should not contain the text "Hello John"`,
			Category: stepbuilder.Visual,
		},
	)
}
