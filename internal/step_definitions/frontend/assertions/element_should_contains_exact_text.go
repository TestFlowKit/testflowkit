package assertions

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) elementShouldContainsExactText() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the text of the {string} element should be exactly {string}`},
		func(ctx context.Context, name, expectedText string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetHTMLElementByLabel(name)
			if err != nil {
				return ctx, err
			}

			actualText := element.TextContent()
			if strings.TrimSpace(actualText) != expectedText {
				return ctx, fmt.Errorf("element %s contains '%s' but expected '%s'", name, actualText, expectedText)
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
			Description: "This assertion checks if the element's visible text is an exact match to the specified string.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "expectedText", Description: "The exact text that should be contained.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the text of the "Welcome Message" element should be exactly "Hello John".`,
			Category: stepbuilder.Visual,
		},
	)
}
