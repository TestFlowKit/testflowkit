package actions

import (
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/steps_definitions/frontend"
	"etoolse/pkg/logger"
	"os"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func Validate(appConfig *config.App) {
	logger.Info("Validate gherkin files ...")

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

	testSuite.Run()
}

func validateScenarioInitializer(ctx *core.ValidatorContext) func(*godog.ScenarioContext) {
	logger.Info("Initializing scenarios for validation ...")

	return func(sc *godog.ScenarioContext) {
		frontend.InitValidationScenarios(sc, ctx)
	}
}

func validateTestSuiteInitializer(validatorCtx *core.ValidatorContext) func(*godog.TestSuiteContext) {
	return func(suiteContext *godog.TestSuiteContext) {
		suiteContext.AfterSuite(func() {
			if !validatorCtx.HasErrors() {
				logger.Success("All is good !")
				os.Exit(0)
			}

			if validatorCtx.HasMissingElements() {
				logger.Error("Elements validation failed", []string{
					"Elements variables malformed in gherkin files",
					"Elements variables not defined in the config file",
				}, []string{
					"Verify the elements variables in the gherkin files",
					validatorCtx.GetElementsErrorsFormatted(),
				})
			}

			if validatorCtx.HasMissingPages() {
				logger.Error("Pages validation failed", []string{
					"Pages variables malformed in gherkin files",
					"Pages variables not defined in the config file",
				}, []string{
					"Verify the pages variables in the gherkin files",
					validatorCtx.GetPagesErrorsFormatted(),
				})
			}

			os.Exit(1)
		})
	}
}
