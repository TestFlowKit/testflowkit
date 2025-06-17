package mouse

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userClicksOnLink() core.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "link")
	}

	handler := clickCommonHandler(formatLabel)

	return core.NewStepWithOneVariable(
		[]string{`^the user clicks the {string} link$`},
		handler.handler(),
		handler.validation(),
		core.StepDefDocParams{
			Description: "Performs a click action on a hyperlink element identified by its logical name",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the link to click on.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user clicks the \"Forgot Password\" link",
			Category: shared.Mouse,
		},
	)
}
