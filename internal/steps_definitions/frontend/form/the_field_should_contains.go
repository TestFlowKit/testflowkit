package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) theFieldShouldContains() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^the {string} should be contain "{string}"`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(fieldId, text string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), fieldId)
				if err != nil {
					return err
				}

				if input.TextContent() == text {
					return nil
				}

				return fmt.Errorf("field should be contains %s but contains %s", text, input.TextContent())
			}
		},
		func(fieldId, _ string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(fieldId) {
				vc.AddMissingElement(fieldId)
			}

			return vc
		},
		core.StepDefDocParams{
			Description: "checks if the field contains the specified text.",
			Variables: []shared.StepVariable{
				{Name: "fieldId", Description: "The id of the field.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text to check.", Type: shared.DocVarTypeString},
			},
			Example:  `Then the "username" should be contain "John"`,
			Category: shared.Form,
		},
	)
}
