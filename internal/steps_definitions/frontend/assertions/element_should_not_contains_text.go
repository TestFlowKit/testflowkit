package assertions

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) elementShouldNotContainsText() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the {string} should not contain the text {string}$`},
		func(ctx context.Context, name, unexpectedText string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetHTMLElementByLabel(name)
			if err != nil {
				return ctx, err
			}

			if !element.IsVisible() {
				return ctx, fmt.Errorf("%s is not visible", name)
			}

			if strings.Contains(element.TextContent(), unexpectedText) {
				return ctx, fmt.Errorf("%s unexpectedly contains text '%s'", name, unexpectedText)
			}

			return ctx, nil
		},
		func(name, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the element's visible text does not include the specified substring.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "unexpectedText", Description: "The text that should not be contained.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the "Welcome Message" element should not contain the text "Hello John"`,
			Category: stepbuilder.Visual,
		},
	)
}
