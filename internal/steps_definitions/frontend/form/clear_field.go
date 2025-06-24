package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) clearField() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "field")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user clears the {string} field$`},
		func(ctx *scenario.Context) func(string) error {
			return func(inputLabel string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				input, err := browser.GetElementByLabel(currentPage, pageName, formatLabel(inputLabel))
				if err != nil {
					return err
				}

				return input.Clear()
			}
		},
		func(inputLabel string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(inputLabel)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "clears the content of an input field.",
			Variables: []stepbuilder.DocVariable{
				{Name: "inputLabel", Description: "The label of the input field to clear.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clears the \"Username\" field",
			Category: stepbuilder.Form,
		},
	)
}
