package mouse

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userClicksOnElement() core.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "element")
	}

	handler := clickCommonHandler(formatLabel)

	return core.NewStepWithOneVariable(
		[]string{`^the user clicks the {string} element$`},
		handler.handler(),
		handler.validation(),
		core.StepDefDocParams{
			Description: "performs a click action on the web element identified by its logical name",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of element to click on.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user clicks the \"Main Logo\" element",
			Category: shared.Mouse,
		},
	)
}
