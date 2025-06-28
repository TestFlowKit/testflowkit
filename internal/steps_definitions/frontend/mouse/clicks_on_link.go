package mouse

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) userClicksOnLink() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "link")
	}

	common := clickCommonHandler(formatLabel)
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user clicks the {string} link$`},
		common.handler(),
		common.validation(),
		stepbuilder.DocParams{
			Description: "performs a click action on the link identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of link to click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clicks the \"Forgot Password\" link",
			Category: stepbuilder.Mouse,
		},
	)
}
