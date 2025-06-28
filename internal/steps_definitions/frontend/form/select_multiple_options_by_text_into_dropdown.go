package form

import (
	"context"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) selectMultipleOptionsByTextIntoDropdown() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "dropdown")
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user selects the options with text {string} from the {string} dropdown$`},
		func(scenarioCtx *scenario.Context) func(context.Context, string, string) (context.Context, error) {
			return func(ctx context.Context, optionLabels, dropdownId string) (context.Context, error) {
				currentPage, pageName := scenarioCtx.GetCurrentPage()
				input, err := browser.GetElementByLabel(currentPage, pageName, formatLabel(dropdownId))
				if err != nil {
					return ctx, err
				}
				return ctx, input.SelectByText(stringutils.SplitAndTrim(optionLabels, ","))
			}
		},
		func(_, dropdownName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(dropdownName)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "selects multiple options from a dropdown by their text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "options", Description: "Comma-separated list of option texts to select.", Type: stepbuilder.VarTypeString},
				{Name: "name", Description: "The logical name of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user selects the options with text "Konoha,Hidden Leaf Village" from the "Country" dropdown`,
			Category: stepbuilder.Form,
		},
	)
}
