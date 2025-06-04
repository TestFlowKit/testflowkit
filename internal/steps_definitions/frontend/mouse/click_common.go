package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
)

type clickCommon struct {
	labelFormatter func(string) string
}

func (c clickCommon) handler() func(*core.TestSuiteContext) func(string) error {
	return func(ctx *core.TestSuiteContext) func(string) error {
		return func(label string) error {
			element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), c.labelFormatter(label))
			if err != nil {
				return err
			}
			return element.Click()
		}
	}
}

func (c clickCommon) validation() func(string) core.ValidationErrors {
	return func(label string) core.ValidationErrors {
		vc := core.ValidationErrors{}
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
