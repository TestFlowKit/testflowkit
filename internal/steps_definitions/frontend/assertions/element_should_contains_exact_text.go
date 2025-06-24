package assertions

import (
	"fmt"
	"strings"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) elementShouldContainsExactText() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the text of the {string} element should be exactly {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(name, expectedText string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				element, err := browser.GetElementByLabel(currentPage, pageName, name)
				if err != nil {
					return err
				}

				actualText := element.TextContent()
				if strings.TrimSpace(actualText) != expectedText {
					return fmt.Errorf("element %s contains '%s' but expected '%s'", name, actualText, expectedText)
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
			Description: "This assertion checks if the element’s visible text is an exact match to the specified string.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "expectedText", Description: "The exact text that should be contained.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the text of the "Welcome Message" element should be exactly "Hello John".`,
			Category: stepbuilder.Visual,
		},
	)
}
