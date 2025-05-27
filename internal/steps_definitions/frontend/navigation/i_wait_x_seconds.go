package navigation

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
	"time"
)

func (n navigation) iWait() core.TestStep {
	testDefinition := func(_ *core.TestSuiteContext) func(int) error {
		return func(seconds int) error {
			time.Sleep(time.Duration(seconds) * time.Second)
			return nil
		}
	}

	return core.NewStepWithOneVariable[int](
		[]string{`^I wait for {number} seconds$`},
		testDefinition,
		nil,
		core.StepDefDocParams{
			Description: "Waits for a specified number of seconds.",
			Variables: []shared.StepVariable{
				{Name: "seconds", Description: "The number of seconds to wait for.", Type: shared.DocVarTypeInt},
			},
			Example:  "When I wait for 5 seconds",
			Category: shared.Navigation,
		},
	)
}
