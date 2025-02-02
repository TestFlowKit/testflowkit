package frontend

import (
	"etoolse/internal/config"
	"log"
	"slices"

	"etoolse/frontend/common"
	"etoolse/frontend/form"
	"etoolse/frontend/keyboard"
	"etoolse/frontend/navigation"
	"etoolse/frontend/visual"
	"etoolse/utils"

	"github.com/cucumber/godog"
)

func InitTestRunnerScenarios(ctx *godog.ScenarioContext, config *config.App) {
	frontendCtx := common.NewFrontendContext(config.Timeout, config.IsHeadlessModeEnabled(), config.GetSlowMotion())
	allSteps := slices.Concat(form.GetSteps(), keyboard.GetSteps(), navigation.GetSteps(), visual.GetSteps())

	log.Println("Initializing scenario for test running ...")
	for _, step := range allSteps {
		handler := step.GetDefinition(frontendCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(utils.ConvertWildcards(sentence), handler)
		}
	}
}

func InitValidationScenarios(ctx *godog.ScenarioContext, vCtx *common.ValidatorContext) {
	allSteps := slices.Concat(form.GetSteps(), keyboard.GetSteps(), navigation.GetSteps(), visual.GetSteps())

	log.Println("Initializing scenarios for validation ...")
	for _, step := range allSteps {
		handler := step.Validate(vCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(utils.ConvertWildcards(sentence), handler)
		}
	}
}
