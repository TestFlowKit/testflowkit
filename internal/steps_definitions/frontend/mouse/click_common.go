package mouse

import (
	"context"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type clickCommon struct {
	labelFormatter func(string) string
}

func (c clickCommon) handler() func(*scenario.Context) func(context.Context, string) (context.Context, error) {
	return func(scenarioCtx *scenario.Context) func(context.Context, string) (context.Context, error) {
		return func(ctx context.Context, label string) (context.Context, error) {
			page, pageName := scenarioCtx.GetCurrentPage()
			element, err := browser.GetElementByLabel(page, pageName, c.labelFormatter(label))
			if err != nil {
				return ctx, err
			}
			err = element.Click()
			return ctx, err
		}
	}
}

func (c clickCommon) validation() func(string) stepbuilder.ValidationErrors {
	return func(label string) stepbuilder.ValidationErrors {
		vc := stepbuilder.ValidationErrors{}
		formattedLabel := c.labelFormatter(label)
		if !config.IsElementDefined(formattedLabel) {
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
