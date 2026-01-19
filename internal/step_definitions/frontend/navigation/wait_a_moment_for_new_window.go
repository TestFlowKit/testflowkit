package navigation

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"time"
)

func (steps) waitAMomentForNewWindow() stepbuilder.Step {
	const docDescription = "Maximum time to wait for a new window (e.g., \"5s\", \"500ms\")."

	return stepbuilder.NewWithOneVariable(
		[]string{"the user waits for a new window to open within {string}"},
		func(ctx context.Context, waitTime string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			duration, err := time.ParseDuration(waitTime)
			if err != nil {
				logger.Error("Invalid duration format: "+waitTime, []string{
					"Duration should be in the format of 1s, 500ms, etc.",
				}, nil)
				return ctx, err
			}

			initialPageCount := len(scenarioCtx.GetPages())
			logger.Info(fmt.Sprintf("Waiting for new window. Current window count: %d", initialPageCount))

			// Wait and check periodically for new windows
			// TODO: refactor in order to use the Wait function
			startTime := time.Now()
			for {
				if time.Since(startTime) > duration {
					return ctx, fmt.Errorf("no new window detected within %s", waitTime)
				}

				currentPageCount := len(scenarioCtx.GetPages())
				if currentPageCount > initialPageCount {
					logger.Info(fmt.Sprintf("New window detected! Page count increased from %d to %d",
						initialPageCount, currentPageCount))
					return ctx, nil
				}

				const milliseconds = 100
				time.Sleep(milliseconds * time.Millisecond)
			}
		},
		func(duration string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			_, err := time.ParseDuration(duration)
			if err != nil {
				vc.AddError("Invalid duration format: " + duration)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "waits for a new browser window to open within the specified timeout.",
			Variables: []stepbuilder.DocVariable{
				{Name: "waitTime", Description: docDescription, Type: stepbuilder.VarTypeString},
			},
			Example:    "When the user waits for a new window to open within \"5s\"",
			Categories: []stepbuilder.StepCategory{stepbuilder.Navigation},
		},
	)
}
