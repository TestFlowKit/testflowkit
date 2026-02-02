package storage

import (
	"context"
	"errors"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) clearSessionStorage() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I clear sessionStorage`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			err := clearStorage(scenarioCtx, "sessionStorage")
			if err != nil {
				return ctx, errors.New("failed to clear sessionStorage: " + err.Error())
			}

			logger.InfoFf("Cleared all sessionStorage items")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Removes all items from sessionStorage.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I clear sessionStorage`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.Browser},
		},
	)
}
