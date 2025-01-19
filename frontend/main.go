package frontend

import (
	"etoolse/config"
	"slices"

	"etoolse/frontend/common"
	"etoolse/frontend/form"
	"etoolse/frontend/keyboard"
	"etoolse/frontend/navigation"
	"etoolse/frontend/visual"
	"etoolse/utils"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext, config *config.AppConfig) {
	frontendCtx := common.NewFrontendContext(config.Timeout, config.IsHeadlessModeEnabled(), config.GetSlowMotion())
	allSteps := slices.Concat(form.GetSteps(), keyboard.GetSteps(), navigation.GetSteps(), visual.GetSteps())
	for _, step := range allSteps {
		handler := step.Definition(frontendCtx)
		for _, sentence := range step.Sentences {
			ctx.Step(utils.ConvertWildcards(sentence), handler)
		}
	}
}
