package mouse

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) userClicksOnButton() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "button")
	}

	handler := clickCommonHandler(formatLabel)
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user clicks the {string} button$`},
		handler.handler(),
		handler.validation(),
		stepbuilder.DocParams{
			Description: "performs a click action specifically on a button element identified by its logical name.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the button to click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clicks the \"Submit Order\" button",
			Category: stepbuilder.Mouse,
		},
	)
}
