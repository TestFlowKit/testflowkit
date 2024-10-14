package frontend

import (
	"slices"

	"github.com/cucumber/godog"

	"cucumber/config"
	"cucumber/frontend/common"
	"cucumber/frontend/form"
	"cucumber/frontend/keyboard"
	"cucumber/frontend/navigation"
	"cucumber/frontend/visual"
	"cucumber/utils"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	config.InitializeFrontConfig()
	frontendCtx := common.NewFrontendContext()

	allSteps := slices.Concat(form.Steps, keyboard.Steps, navigation.Steps, visual.Steps)

	for _, step := range allSteps {
		handler := step.Definition(frontendCtx)
		for _, sentence := range step.Sentences {
			ctx.Step(utils.ConvertWildcards(sentence), handler)
		}
	}
}
