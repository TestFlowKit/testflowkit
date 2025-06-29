package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) dropdownHasValuesSelected() stepbuilder.Step {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	doc := stepbuilder.DocParams{
		Description: "checks if the dropdown has the specified values selected.",
		Variables: []stepbuilder.DocVariable{
			{Name: "dropdownId", Description: "The id of the dropdown.", Type: stepbuilder.VarTypeString},
			{Name: "optionLabels", Description: "The labels of the options to check.", Type: stepbuilder.VarTypeString},
		},
		Example:  `Then the "country" dropdown should have "USA,Canada" selected`,
		Category: stepbuilder.Form,
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`the {string} dropdown should have "{string}" selected`},
		func(ctx context.Context, dropdownId, optionLabels string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage := scenarioCtx.GetCurrentPageOnly()
			selector, err := currentPage.GetAllBySelector(formatVar(dropdownId))
			if err != nil {
				return ctx, err
			}

			labels := stringutils.SplitAndTrim(optionLabels, ",")

			result := currentPage.ExecuteJS(`(selector, labels) => {
				const selectedOpts = Array.from(document.querySelector(selector).selectedOptions).map(opt => opt.label)
				return labels.every(label => selectedOpts.includes(label))
			}`, selector, labels)

			if result == "true" {
				return ctx, nil
			}
			return ctx, fmt.Errorf("%s value is not selected in %s dropdown", optionLabels, dropdownId)
		},
		func(dropdownId, _ string) stepbuilder.ValidationErrors {
			vErr := stepbuilder.ValidationErrors{}
			label := formatVar(dropdownId)
			if !config.IsElementDefined(label) {
				vErr.AddMissingElement(label)
			}

			return vErr
		},
		doc,
	)
}
