package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) iTypeXXXIntoInput() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I type "{string}" into the {string}`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(text, inputLabel string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), inputLabel)
				if err != nil {
					return err
				}
				return input.Input(text)
			}
		},
		func(_, inputLabel string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(inputLabel) {
				vc.AddMissingElement(inputLabel)
			}

			return vc
		},
		core.StepDefDocParams{
			Description: "types the specified text into the input.",
			Variables: []shared.StepVariable{
				{Name: "text", Description: "The text to type.", Type: shared.DocVarTypeString},
				{Name: "inputLabel", Description: "The label of the input.", Type: shared.DocVarTypeString},
			},
			Example:  `When I type "John" into the "username field"`,
			Category: shared.Form,
		},
	)
}
