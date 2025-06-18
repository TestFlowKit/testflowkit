package mouse

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userClicksOnElement() stepbuilder.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "element")
	}

	handler := clickCommonHandler(formatLabel)

	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user clicks the {string} element$`},
		handler.handler(),
		handler.validation(),
		stepbuilder.StepDefDocParams{
			Description: "performs a click action on the web element identified by its logical name",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of element to click on.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user clicks the \"Main Logo\" element",
			Category: shared.Mouse,
		},
	)
}
