package assertions

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) elementShouldContainsExactText() stepbuilder.Step {
	return s.newElementExactTextStep(
		`the text of the {string} element should be exactly {string}`,
		true,
		stepbuilder.DocParams{
			Description: "This assertion checks if the element's visible text is an exact match to the specified string.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "expectedText", Description: "The exact text that should be contained.", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the text of the "Welcome Message" element should be exactly "Hello John".`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions, stepbuilder.Frontend},
		},
	)
}

func (s steps) elementShouldNotContainsExactText() stepbuilder.Step {
	return s.newElementExactTextStep(
		`the text of the {string} element should not be exactly {string}`,
		false,
		func() stepbuilder.DocParams {
			desc := "This assertion checks if the element's visible text is not an exact match to the specified string."
			str := stepbuilder.VarTypeString
			vars := []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: str},
				{Name: "unexpectedText", Description: "The exact text that should not be contained.", Type: str},
			}
			return stepbuilder.DocParams{
				Description: desc,
				Variables:   vars,
				Example:     `Then the text of the "Welcome Message" element should not be exactly "Error".`,
				Categories:  []stepbuilder.StepCategory{stepbuilder.Assertions, stepbuilder.Frontend},
			}
		}(),
	)
}

func (s steps) newElementExactTextStep(sentence string, shouldEqual bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, name, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetHTMLElementByLabel(name)
			if err != nil {
				return ctx, err
			}

			actualText := element.TextContent()
			actualTrimmed := strings.TrimSpace(actualText)
			if shouldEqual && actualTrimmed != text {
				return ctx, fmt.Errorf("element %s contains '%s' but expected '%s'", name, actualText, text)
			}

			if !shouldEqual && actualTrimmed == text {
				return ctx, fmt.Errorf("element %s should not be exactly '%s'", name, text)
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
		doc,
	)
}
