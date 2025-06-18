package assertions

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) theFieldShouldContains() stepbuilder.TestStep {
	formatFieldID := func(fieldId string) string {
		return fmt.Sprintf("%s_field", fieldId)
	}

	return stepbuilder.NewStepWithTwoVariables(
		[]string{`^the value of the {string} field should be {string}`},
		func(ctx *stepbuilder.TestSuiteContext) func(string, string) error {
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
		func(fieldId, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsElementDefined(formatFieldID(fieldId)) {
				vc.AddMissingElement(formatFieldID(fieldId))
			}

			return vc
		},
		stepbuilder.StepDefDocParams{
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
