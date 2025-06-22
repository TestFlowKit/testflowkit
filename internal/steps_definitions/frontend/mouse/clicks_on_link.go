package mouse

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) userClicksOnLink() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "link")
	}

	handler := clickCommonHandler(formatLabel)

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user clicks the {string} link$`},
		handler.handler(),
		handler.validation(),
		stepbuilder.DocParams{
			Description: "Performs a click action on a hyperlink element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the link to click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clicks the \"Forgot Password\" link",
			Category: stepbuilder.Mouse,
		},
	)
}
