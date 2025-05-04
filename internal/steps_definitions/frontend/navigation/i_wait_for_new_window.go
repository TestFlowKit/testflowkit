package navigation

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
	"time"
)

func (n navigation) iWaitForNewWindow() core.TestStep {
	const docDescription = "Maximum time to wait for a new window (e.g., \"5s\", \"500ms\")."

	return core.NewStepWithOneVariable(
		[]string{"^I wait for a new window to open within {string}$"},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(waitTime string) error {
				duration, err := time.ParseDuration(waitTime)
				if err != nil {
					logger.Error(fmt.Sprintf("Invalid duration format: %s", waitTime), []string{
						"Duration should be in the format of 1s, 500ms, etc.",
					}, nil)
					return err
				}

				initialPageCount := len(ctx.GetPages())
				logger.Info(fmt.Sprintf("Waiting for new window. Current window count: %d", initialPageCount))

				// Wait and check periodically for new windows
				// TODO: refactor in order to use the Wait function
				startTime := time.Now()
				for {
					// Check if we've exceeded the wait time
					if time.Since(startTime) > duration {
						return fmt.Errorf("no new window detected within %s", waitTime)
					}

					// Check if the number of pages has increased
					currentPageCount := len(ctx.GetPages())
					if currentPageCount > initialPageCount {
						logger.Info(fmt.Sprintf("New window detected! Page count increased from %d to %d",
							initialPageCount, currentPageCount))
						return nil
					}

					// Wait a short time before checking again
					const milliseconds = 100
					time.Sleep(milliseconds * time.Millisecond)
				}
			}
		},
		func(duration string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			_, err := time.ParseDuration(duration)
			if err != nil {
				vc.AddError(fmt.Sprintf("Invalid duration format: %s", duration))
			}
			return vc
		},
		core.StepDefDocParams{
			Description: "waits for a new browser window to open within the specified timeout.",
			Variables: []shared.StepVariable{
				{Name: "waitTime", Description: docDescription, Type: shared.DocVarTypeString},
			},
			Example:  "When I wait for a new window to open within \"5s\"",
			Category: shared.Navigation,
		},
	)
}
