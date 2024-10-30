package frontend

import (
	"cucumber/config"
	"slices"

	"cucumber/frontend/common"
	"cucumber/frontend/form"
	"cucumber/frontend/keyboard"
	"cucumber/frontend/navigation"
	"cucumber/frontend/visual"
	"cucumber/utils"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext, config config.ClI) {
	frontendCtx := common.NewFrontendContext(config.Timeout, !config.DisplayBrowser)

	allSteps := slices.Concat(form.Steps, keyboard.Steps, navigation.Steps, visual.Steps)

	for _, step := range allSteps {
		handler := step.Definition(frontendCtx)
		for _, sentence := range step.Sentences {
			ctx.Step(utils.ConvertWildcards(sentence), handler)
		}
	}
}
