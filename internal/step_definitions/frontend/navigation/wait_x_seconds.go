package navigation

import (
	"context"
	"fmt"
	"strconv"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"time"
)

func (steps) userWait() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user waits for {number} seconds`},
		func(ctx context.Context, seconds string) (context.Context, error) {
			secondsInt, err := strconv.Atoi(seconds)
			if err != nil {
				return ctx, fmt.Errorf("invalid seconds: %s", seconds)
			}

			time.Sleep(time.Duration(secondsInt) * time.Second)
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
