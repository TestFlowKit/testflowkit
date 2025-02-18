package actions

import (
	"context"
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/steps_definitions/frontend"
	"etoolse/pkg/gherkinparser"
	"etoolse/pkg/logger"
	"fmt"
	"os"
	"strings"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func Validate(appConfig *config.App) {
	logger.Info("Validate gherkin files ...")

	parsedFeatures := gherkinparser.Parse(appConfig.GherkinLocation)
	features := make([]godog.Feature, len(parsedFeatures))
	for i, f := range parsedFeatures {
		features[i] = godog.Feature{
			Name:     f.Name,
			Contents: f.Contents,
		}
	}

	const concurrency = 5
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         concurrency,
		ShowStepDefinitions: false,
		Format:              "pretty",
		Tags:                appConfig.Tags,
		FeatureContents:     features,
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
		sc.StepContext().After(validateAfterStepHookInitializer(ctx))
	}
}

func validateAfterStepHookInitializer(vCtx *core.ValidatorContext) godog.AfterStepHook {
	return func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		if status == godog.StepUndefined {
			vCtx.AddUndefinedStep(st.Text)
		}
		return ctx, err
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

			if validatorCtx.HasUndefinedSteps() {
				indent := 3
				indents := logger.GetIndents(indent)
				steps := strings.Join(validatorCtx.GetUndefinedSteps(), indents)
				msg := fmt.Sprintf("this is the list of undefined steps: \n%s%s", indents, steps)
				logger.Error("Steps validation failed",
					[]string{
						"Steps are malformed in the gherkin files",
						msg,
					},
					[]string{
						"Verify the steps in the gherkin files",
						"Please refer to documentation for complete list of steps",
					})
			}

			os.Exit(1)
		})
	}
}
