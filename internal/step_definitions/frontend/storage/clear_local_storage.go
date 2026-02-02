package storage

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) clearLocalStorage() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I clear localStorage`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := clearStorage(scenarioCtx, "localStorage")
			if err != nil {
				return ctx, err
			}

			logger.InfoFf("Cleared all localStorage items")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Removes all items from localStorage.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I clear localStorage`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
