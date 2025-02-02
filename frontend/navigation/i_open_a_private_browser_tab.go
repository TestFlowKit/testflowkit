package navigation

import (
	"etoolse/frontend/common"
)

func (n navigation) iOpenNewPrivateBrowserTab() common.TestStep {
	return common.NewStepWithoutVariables(
		[]string{"I open a new private browser tab"},
		func(ctx *common.TestSuiteContext) func() error {
			return func() error {
				ctx.InitBrowser(true)
				return nil
			}
		},
		nil,
	)
}
