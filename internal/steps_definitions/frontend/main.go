package frontend

import (
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/steps_definitions/frontend/form"
	"etoolse/internal/steps_definitions/frontend/keyboard"
	"etoolse/internal/steps_definitions/frontend/navigation"
	"etoolse/internal/steps_definitions/frontend/visual"
	"etoolse/shared"
	"slices"

	"github.com/cucumber/godog"
)

func InitTestRunnerScenarios(ctx *godog.ScenarioContext, config *config.App) {
	frontendCtx := core.NewFrontendContext(config.Timeout, config.IsHeadlessModeEnabled(), config.GetSlowMotion())
	for _, step := range getAllSteps() {
		handler := step.GetDefinition(frontendCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(core.ConvertWildcards(sentence), handler)
		}
	}
}

func InitValidationScenarios(ctx *godog.ScenarioContext, vCtx *core.ValidatorContext) {
	for _, step := range getAllSteps() {
		handler := step.Validate(vCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(core.ConvertWildcards(sentence), handler)
		}
	}
}

func getAllSteps() []core.TestStep {
	return slices.Concat(form.GetSteps(), keyboard.GetSteps(), navigation.GetSteps(), visual.GetSteps())
}

func GetDocs() []shared.StepDocumentation {
	var docs []shared.StepDocumentation
	for _, step := range getAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
