package frontend

import (
	"slices"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/steps_definitions/frontend/assertions"
	"testflowkit/internal/steps_definitions/frontend/form"
	"testflowkit/internal/steps_definitions/frontend/keyboard"
	"testflowkit/internal/steps_definitions/frontend/mouse"
	"testflowkit/internal/steps_definitions/frontend/navigation"
	"testflowkit/internal/steps_definitions/frontend/visual"

	"github.com/cucumber/godog"
)

func InitTestRunnerScenarios(ctx *godog.ScenarioContext, config *config.Config) {
	frontendCtx := scenario.NewContext(config)
	for _, step := range getAllSteps() {
		handler := step.GetDefinition(frontendCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(core.ConvertWildcards(sentence), handler)
		}
	}
}

func InitValidationScenarios(ctx *godog.ScenarioContext, vCtx *stepbuilder.ValidatorContext) {
	for _, step := range getAllSteps() {
		handler := step.Validate(vCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(core.ConvertWildcards(sentence), handler)
		}
	}
}

func getAllSteps() []stepbuilder.Step {
	return slices.Concat(
		form.GetSteps(),
		keyboard.GetSteps(),
		navigation.GetSteps(),
		visual.GetSteps(),
		mouse.GetSteps(),
		assertions.GetSteps(),
	)
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range getAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
