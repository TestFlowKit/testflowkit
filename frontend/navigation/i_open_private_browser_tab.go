package navigation

import (
	"cucumber/frontend/common"
)

var iOpenAPrivateBrowserTab = common.FrontStep{
	Sentences: []string{"I open a new private browser tab"},
	Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
		return func() error {
			ctx.InitBrowser()
			return nil
		}
	},
}
