package assertions

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) theFieldShouldContains() core.TestStep {
	formatFieldID := func(fieldId string) string {
		return fmt.Sprintf("%s_field", fieldId)
	}

	return core.NewStepWithTwoVariables(
		[]string{`^the value of the {string} field should be {string}`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(fieldId, text string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatFieldID(fieldId))
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
			if !testsconfig.IsElementDefined(formatFieldID(fieldId)) {
				vc.AddMissingElement(formatFieldID(fieldId))
			}

			return vc
		},
		core.StepDefDocParams{
			Description: "This assertion checks if the current value of an input field matches the specified value.",
			Variables: []shared.StepVariable{
				{Name: "fieldId", Description: "The id of the field.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text to check.", Type: shared.DocVarTypeString},
			},
			Example:  `Then the value of the "Username" field should be "myUsername".`,
			Category: shared.Form,
		},
	)
}
