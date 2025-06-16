package navigation

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
	"time"
)

func (n navigation) userWait() core.TestStep {
	testDefinition := func(_ *core.TestSuiteContext) func(int) error {
		return func(seconds int) error {
			time.Sleep(time.Duration(seconds) * time.Second)
			return nil
		}
	}

	return core.NewStepWithOneVariable(
		[]string{`^the user waits for {number} seconds$`},
		testDefinition,
		nil,
		core.StepDefDocParams{
			Description: "Pauses the test execution for a specified number of seconds.",
			Variables: []shared.StepVariable{
				{Name: "seconds", Description: "The number of seconds to wait for.", Type: shared.DocVarTypeInt},
			},
			Example:  "When the user waits for 3 seconds",
			Category: shared.Navigation,
		},
	)
}
