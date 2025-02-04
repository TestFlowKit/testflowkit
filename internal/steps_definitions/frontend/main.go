package frontend

import (
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/steps_definitions/frontend/form"
	"etoolse/internal/steps_definitions/frontend/keyboard"
	"etoolse/internal/steps_definitions/frontend/navigation"
	"etoolse/internal/steps_definitions/frontend/visual"
	"slices"

	"github.com/cucumber/godog"
)

func InitTestRunnerScenarios(ctx *godog.ScenarioContext, config *config.App) {
	frontendCtx := core.NewFrontendContext(config.Timeout, config.IsHeadlessModeEnabled(), config.GetSlowMotion())
	allSteps := slices.Concat(form.GetSteps(), keyboard.GetSteps(), navigation.GetSteps(), visual.GetSteps())

	for _, step := range allSteps {
		handler := step.GetDefinition(frontendCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(core.ConvertWildcards(sentence), handler)
		}
	}
}

func InitValidationScenarios(ctx *godog.ScenarioContext, vCtx *core.ValidatorContext) {
	allSteps := slices.Concat(form.GetSteps(), keyboard.GetSteps(), navigation.GetSteps(), visual.GetSteps())

	for _, step := range allSteps {
		handler := step.Validate(vCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(core.ConvertWildcards(sentence), handler)
		}
	}
}
