package navigation

import (
	"cucumber/frontend/common"
)

func (n navigation) iOpenNewPrivateBrowserTab() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{"I open a new private browser tab"},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func() error {
				ctx.InitBrowser(true)
				return nil
			}
		},
	}
}
