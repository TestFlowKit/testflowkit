package navigation

import (
	"cucumber/frontend/common"
)

func (n navigation) iOpenNewBrowserTab() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{"I open a new browser tab"},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func() error {
				ctx.InitBrowser(false)
				return nil
			}
		},
	}
}
