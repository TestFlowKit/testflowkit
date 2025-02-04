package navigation

import (
	"etoolse/internal/steps_definitions/core"
)

func (n navigation) iOpenNewPrivateBrowserTab() core.TestStep {
	return core.NewStepWithoutVariables(
		[]string{"I open a new private browser tab"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				ctx.InitBrowser(true)
				return nil
			}
		},
		nil,
	)
}
