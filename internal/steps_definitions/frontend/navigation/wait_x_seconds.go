package navigation

import (
	"context"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"time"
)

type waitXSecondsHandler = func(context.Context, int) (context.Context, error)

func (steps) userWait() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user waits for {number} seconds$`},
		func(ctx context.Context, seconds int) (context.Context, error) {
			time.Sleep(time.Duration(seconds) * time.Second)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "waits for a specified number of seconds.",
			Variables: []stepbuilder.DocVariable{
				{Name: "seconds", Description: "The number of seconds to wait.", Type: stepbuilder.VarTypeInt},
			},
			Example:  "When the user waits for 3 seconds",
			Category: stepbuilder.Navigation,
		},
	)
}
