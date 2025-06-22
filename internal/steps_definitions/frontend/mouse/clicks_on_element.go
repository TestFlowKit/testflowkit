package mouse

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) userClicksOnElement() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "element")
	}

	handler := clickCommonHandler(formatLabel)

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user clicks the {string} element$`},
		handler.handler(),
		handler.validation(),
		stepbuilder.DocParams{
			Description: "performs a click action on the web element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of element to click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clicks the \"Main Logo\" element",
			Category: stepbuilder.Mouse,
		},
	)
}
