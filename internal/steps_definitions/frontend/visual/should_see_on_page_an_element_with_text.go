package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) shouldSeeElementWhichContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user should see a (link|button|element) which contains "{string}"$`},
		func(ctx context.Context, elementLabel, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			cases := map[string]string{
				"link":    "a",
				"button":  "button",
				"element": "*",
			}

			xPath := fmt.Sprintf("//%s[contains(text(),\"%s\")]", cases[elementLabel], text)
			element, err := scenarioCtx.GetCurrentPageOnly().GetOneByXPath(xPath)
			cErr := fmt.Errorf("no %s is visible with text \"%s\"", elementLabel, text)
			if err != nil {
				return ctx, cErr
			}

			visible := element.IsVisible()
			if !visible {
				return ctx, cErr
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a link, button or element is visible and contains a specific text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the user should see a button which contains \"Submit\"",
			Category: stepbuilder.Visual,
		},
	)
}
