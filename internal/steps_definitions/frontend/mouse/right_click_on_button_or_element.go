package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) rightClickOn() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user right clicks on {string}$`},
		func(ctx *scenario.Context) func(string) error {
			return func(label string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				element, err := browser.GetElementByLabel(currentPage, pageName, label)
				if err != nil {
					return err
				}
				return element.RightClick()
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
			Description: "Right clicks on a button or element.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to right click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user right clicks on \"Submit button\"",
			Category: stepbuilder.Mouse,
		},
	)
}
