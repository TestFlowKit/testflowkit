package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type clickCommon struct {
	labelFormatter func(string) string
}

func (c clickCommon) handler() func(*scenario.Context) func(string) error {
	return func(ctx *scenario.Context) func(string) error {
		return func(label string) error {
			element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), c.labelFormatter(label))
			if err != nil {
				return err
			}
			return element.Click()
		}
	}
}

func (c clickCommon) validation() func(string) stepbuilder.ValidationErrors {
	return func(label string) stepbuilder.ValidationErrors {
		vc := stepbuilder.ValidationErrors{}
		formattedLabel := c.labelFormatter(label)
		if !testsconfig.IsElementDefined(formattedLabel) {
			vc.AddMissingElement(formattedLabel)
		}
		return vc
	}
}

func clickCommonHandler(labelFormatter func(string) string) *clickCommon {
	return &clickCommon{
		labelFormatter: labelFormatter,
	}
}
