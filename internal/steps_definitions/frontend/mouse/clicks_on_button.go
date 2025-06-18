package mouse

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userClicksOnButton() stepbuilder.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "button")
	}

	handler := clickCommonHandler(formatLabel)
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user clicks the {string} button$`},
		handler.handler(),
		handler.validation(),
		stepbuilder.StepDefDocParams{
			Description: "performs a click action specifically on a button element identified by its logical name.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the button to click on.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user clicks the \"Submit Order\" button",
			Category: shared.Mouse,
		},
	)
}
