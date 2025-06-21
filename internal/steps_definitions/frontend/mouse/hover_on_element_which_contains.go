package mouse

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) hoverOnElementWhichContains() stepbuilder.TestStep {
	return stepbuilder.NewStepWithTwoVariables(
		[]string{`^the user hovers on {string} which contains {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(_ string, text string) error {
				element, err := ctx.GetCurrentPage().GetOneByTextContent(text)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.Hover()
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "Hover on an element which contains a specific text.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to hover on.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user hovers on \"Submit button\" which contains \"Submit\"",
			Category: shared.Mouse,
		},
	)
}
