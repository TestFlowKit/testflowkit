package mouse

import (
	"context"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type hoverOnElementHandler = func(context.Context, string) (context.Context, error)

func (steps) hoverOnElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user hovers on {string}$`},
		func(scenarioCtx *scenario.Context) hoverOnElementHandler {
			return func(ctx context.Context, label string) (context.Context, error) {
				currentPage, pageName := scenarioCtx.GetCurrentPage()
				element, err := browser.GetElementByLabel(currentPage, pageName, label)
				if err != nil {
					return ctx, err
				}
				err = element.Hover()
				return ctx, err
			}
		},
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "performs a hover action on the element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of element to hover on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user hovers on \"Submit button\"",
			Category: stepbuilder.Mouse,
		},
	)
}
