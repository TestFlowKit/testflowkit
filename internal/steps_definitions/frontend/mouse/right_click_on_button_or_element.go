package mouse

import (
	"context"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) rightClickOn() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user right clicks on {string}$`},
		func(ctx context.Context, label string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageName := scenarioCtx.GetCurrentPage()
			element, err := browser.GetElementByLabel(currentPage, pageName, label)
			if err != nil {
				return ctx, err
			}
			err = element.RightClick()
			return ctx, err

		},
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "performs a right click action on the element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of element to right click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user right clicks on \"Submit button\"",
			Category: stepbuilder.Mouse,
		},
	)
}
