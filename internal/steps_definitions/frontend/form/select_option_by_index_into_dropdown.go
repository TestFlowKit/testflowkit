package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) userSelectOptionByIndexIntoDropdown() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "dropdown")
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user selects the option at index {number} from the {string} dropdown$`},
		func(ctx context.Context, index int, dropdownId string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			input, err := scenarioCtx.GetHTMLElementByLabel(formatLabel(dropdownId))
			if err != nil {
				return ctx, err
			}

			return ctx, input.SelectByIndex(index)
		},
		func(_ int, dropdownName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(dropdownName)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "selects an option from a dropdown by its index.",
			Variables: []stepbuilder.DocVariable{
				{Name: "index", Description: "The index of the option to select.", Type: stepbuilder.VarTypeInt},
				{Name: "name", Description: "The logical name of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user selects the option at index 2 from the "Country" dropdown`,
			Category: stepbuilder.Form,
		},
	)
}
