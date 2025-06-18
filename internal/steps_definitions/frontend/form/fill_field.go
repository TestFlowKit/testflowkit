package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userEntersTextIntoField() stepbuilder.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "field")
	}

	return stepbuilder.NewStepWithTwoVariables(
		[]string{`^the user enters {string} into the {string} field`},
		func(ctx *stepbuilder.TestSuiteContext) func(string, string) error {
			return func(text, inputLabel string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(inputLabel))
				if err != nil {
					return err
				}
				return input.Input(text)
			}
		},
		func(_, inputLabel string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(inputLabel)
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Types the specified text into an input field identified by its logical name.",
			Variables: []shared.StepVariable{
				{Name: "text", Description: "The text to type.", Type: shared.DocVarTypeString},
				{Name: "name", Description: "The logical name of the input field.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user enters "myUsername" into the "Username" field`,
			Category: shared.Form,
		},
	)
}
