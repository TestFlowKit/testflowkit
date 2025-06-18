package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userClearsFormField() stepbuilder.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "field")
	}

	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user clears the {string} field.`},
		func(ctx *stepbuilder.TestSuiteContext) func(string) error {
			return func(inputLabel string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(inputLabel))
				if err != nil {
					return err
				}
				return input.Clear()
			}
		},
		func(inputLabel string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(inputLabel)
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Removes any existing text or value from an input field identified by its logical name.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the input field.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user clears the "Search" field.`,
			Category: shared.Form,
		},
	)
}
