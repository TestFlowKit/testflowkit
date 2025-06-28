package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) scrollToElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user scrolls to the {string} element$`},
		func(scenarioCtx *scenario.Context) func(context.Context, string) (context.Context, error) {
			return func(ctx context.Context, name string) (context.Context, error) {
				page, pageName := scenarioCtx.GetCurrentPage()
				element, err := browser.GetElementByLabel(page, pageName, fmt.Sprintf("%s_element", name))
				if err != nil {
					return ctx, err
				}

				scrollErr := element.ScrollIntoView()
				if scrollErr != nil {
					return ctx, scrollErr
				}

				return ctx, nil
			}
		},
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			variable := fmt.Sprintf("%s_element", name)
			if !config.IsElementDefined(variable) {
				vc.AddMissingElement(variable)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "Scrolls the page until the specified element is visible in the viewport.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user scrolls to the "Submit Button at the bottom" element`,
			Category: stepbuilder.Visual,
		},
	)
}
