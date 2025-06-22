package navigation

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"time"
)

func (n navigation) userWait() stepbuilder.Step {
	testDefinition := func(_ *scenario.Context) func(int) error {
		return func(seconds int) error {
			time.Sleep(time.Duration(seconds) * time.Second)
			return nil
		}
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user waits for {number} seconds$`},
		testDefinition,
		nil,
		stepbuilder.DocParams{
			Description: "Pauses the test execution for a specified number of seconds.",
			Variables: []stepbuilder.DocVariable{
				{Name: "seconds", Description: "The number of seconds to wait for.", Type: stepbuilder.VarTypeInt},
			},
			Example:  "When the user waits for 3 seconds",
			Category: stepbuilder.Navigation,
		},
	)
}
