package mouse

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type clickCommon struct {
	labelFormatter func(string) string
}

func (c clickCommon) handler() func(context.Context, string) (context.Context, error) {
	return func(ctx context.Context, label string) (context.Context, error) {
		scCtx := scenario.MustFromContext(ctx)
		element, err := scCtx.GetHTMLElementByLabel(c.labelFormatter(label))
		if err != nil {
			return ctx, err
		}
		err = element.Click()
		return ctx, err
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
