package navigation

import (
	"etoolse/frontend/common"
)

func (n navigation) iOpenNewBrowserTab() common.TestStep {
	return common.NewStepWithoutVariables(
		[]string{"I open a new browser tab"},
		func(ctx *common.TestSuiteContext) func() error {
			return func() error {
				ctx.InitBrowser(false)
				return nil
			}
		},
		nil,
	)
}
