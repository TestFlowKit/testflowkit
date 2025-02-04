package actions

import (
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/steps_definitions/frontend"
	"log"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func Validate(appConfig *config.App) {
	log.Println("Validate gherkin files ...")

	const concurrency = 5
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         concurrency,
		Format:              "pretty",
		ShowStepDefinitions: false,
		Tags:                appConfig.Tags,
		Paths:               []string{appConfig.GherkinLocation},
	}

	ctx := core.ValidatorContext{}
	testSuite := godog.TestSuite{
		Name:                 "validate",
		Options:              &opts,
		ScenarioInitializer:  validateScenarioInitializer(&ctx),
		TestSuiteInitializer: validateTestSuiteInitializer(&ctx),
	}

	status := testSuite.Run()
	if status != 0 {
		log.Fatalf("zero status code expected, %d received", status)
	}
}

func validateScenarioInitializer(ctx *core.ValidatorContext) func(*godog.ScenarioContext) {
	log.Println("Initializing scenarios for validation ...")

	return func(sc *godog.ScenarioContext) {
		frontend.InitValidationScenarios(sc, ctx)
	}
}

func validateTestSuiteInitializer(validatorCtx *core.ValidatorContext) func(*godog.TestSuiteContext) {
	return func(suiteContext *godog.TestSuiteContext) {
		suiteContext.AfterSuite(func() {
			log.Println("Errors:")
			log.Println(validatorCtx.GetErrors())
		})
	}
}
