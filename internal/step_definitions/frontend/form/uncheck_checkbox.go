package form

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
)

func (steps) uncheckCheckbox() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "checkbox")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`the user unchecks the {string} checkbox`},
		func(ctx context.Context, checkBoxName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			checkBox, err := scenarioCtx.GetHTMLElementByLabel(formatLabel(checkBoxName))
			if err != nil {
				return ctx, err
			}

			if checkBox.IsChecked() {
				err = checkBox.Click()
				return ctx, err
			}

			logger.Warn(fmt.Sprintf("%s checkbox is not unchecked because it is already unchecked", checkBoxName), []string{})
			return ctx, nil
		},
		func(checkBoxName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(checkBoxName)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "unchecks a checkbox if it is currently checked.",
			Variables: []stepbuilder.DocVariable{
				{Name: "checkBoxName", Description: "The name of the checkbox to uncheck.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user unchecks the \"Newsletter\" checkbox",
			Category: stepbuilder.Form,
		},
	)
}
