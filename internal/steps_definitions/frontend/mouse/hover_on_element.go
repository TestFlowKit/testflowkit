package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) hoverOnElement() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user hovers on {string}$`},
		func(ctx *stepbuilder.TestSuiteContext) func(label string) error {
			return func(label string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.Hover()
			}
		},
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Hover on a element.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to hover on.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user hovers on \"Submit button\"",
			Category: shared.Mouse,
		},
	)
}
