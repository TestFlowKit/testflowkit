package mouse

import (
	"context"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) doubleClickOn() stepbuilder.Step {
	const docDescription = "The logical name of element to double click on."

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user double clicks on {string}$`},
		func(scenarioCtx *scenario.Context) func(context.Context, string) (context.Context, error) {
			return func(ctx context.Context, label string) (context.Context, error) {
				currentPage, pageName := scenarioCtx.GetCurrentPage()
				element, err := browser.GetElementByLabel(currentPage, pageName, label)
				if err != nil {
					return ctx, err
				}
				err = element.DoubleClick()
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
			Description: "performs a double click action on the element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: docDescription, Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user double clicks on \"File item\"",
			Category: stepbuilder.Mouse,
		},
	)
}
