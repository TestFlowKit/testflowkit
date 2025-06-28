package mouse

import (
	"context"
	"testflowkit/internal/browser"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/steps_definitions/core/scenario"
)

type simpleElementInteractionFunc = func(ctx context.Context, label string) (context.Context, error)

func commonSimpleElementInteraction(action func(common.Element) error) simpleElementInteractionFunc {
	return func(ctx context.Context, label string) (context.Context, error) {
		scenarioCtx := scenario.MustFromContext(ctx)
		currentPage, pageName := scenarioCtx.GetCurrentPage()
		element, err := browser.GetElementByLabel(currentPage, pageName, label)
		if err != nil {
			return ctx, err
		}
		err = action(element)
		return ctx, err
	}
}
