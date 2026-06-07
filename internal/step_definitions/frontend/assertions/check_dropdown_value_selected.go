package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) dropdownHasValuesSelected() stepbuilder.Step {
	return s.newDropdownSelectedStep(
		`the {string} dropdown should have {string} selected`,
		true,
		stepbuilder.DocParams{
			Description: "checks if the dropdown has the specified values selected.",
			Variables: []stepbuilder.DocVariable{
				{Name: "dropdownId", Description: "The id of the dropdown.", Type: stepbuilder.VarTypeString},
				{Name: "optionLabels", Description: "The labels of the options to check.", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the "country" dropdown should have "USA,Canada" selected`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Form, stepbuilder.Frontend},
		},
	)
}

func (s steps) dropdownShouldNotHaveValuesSelected() stepbuilder.Step {
	str := stepbuilder.VarTypeString
	vars := []stepbuilder.DocVariable{
		{Name: "dropdownId", Description: "The id of the dropdown.", Type: str},
		{Name: "optionLabels", Description: "The labels of the options that should not all be selected.", Type: str},
	}

	return s.newDropdownSelectedStep(
		`the {string} dropdown should not have {string} selected`,
		false,
		stepbuilder.DocParams{
			Description: "checks if the dropdown does not have the specified values selected.",
			Variables:   vars,
			Example:     `Then the "country" dropdown should not have "USA,Canada" selected`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.Form, stepbuilder.Frontend},
		},
	)
}

func (s steps) newDropdownSelectedStep(
	sentence string, shouldBeSelected bool, doc stepbuilder.DocParams,
) stepbuilder.Step {
	formatVar := func(label string) string {
		return label + "_dropdown"
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, dropdownId, optionLabels string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			selector, err := currentPage.GetAllBySelector(formatVar(dropdownId))
			if err != nil {
				return ctx, err
			}

			labels := stringutils.SplitAndTrim(optionLabels, ",")

			result := currentPage.ExecuteJS(`(selector, labels) => {
				const selectedOpts = Array.from(document.querySelector(selector).selectedOptions).map(opt => opt.label)
				return labels.every(label => selectedOpts.includes(label))
			}`, selector, labels)

			isSelected := result == "true"
			if shouldBeSelected && !isSelected {
				return ctx, fmt.Errorf("%s value is not selected in %s dropdown", optionLabels, dropdownId)
			}

			if !shouldBeSelected && isSelected {
				return ctx, fmt.Errorf("%s value should not be selected in %s dropdown", optionLabels, dropdownId)
			}

			return ctx, nil
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
