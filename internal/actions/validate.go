package actions

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testflowkit/internal/config"
	stepdefinitions "testflowkit/internal/step_definitions"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/gherkinparser"
	"testflowkit/pkg/logger"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func validate(appConfig *config.Config, cfgErr error) {
	if cfgErr != nil {
		logger.Fatal("VALIDATE", cfgErr)
	}

	displayConfigSummary(appConfig)

	logger.Info("Validate gherkin files ...")

	parsedFeatures := gherkinparser.Parse(appConfig.Settings.GherkinLocation)
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
		Tags:                appConfig.Settings.Tags,
		FeatureContents:     features,
	}

	ctx := stepbuilder.ValidatorContext{}
	testSuite := godog.TestSuite{
		Name:                 "validate",
		Options:              &opts,
		ScenarioInitializer:  validateScenarioInitializer(&ctx),
		TestSuiteInitializer: validateTestSuiteInitializer(&ctx),
	}

	testSuite.Run()
}

func validateScenarioInitializer(ctx *stepbuilder.ValidatorContext) func(*godog.ScenarioContext) {
	logger.Info("Initializing scenarios for validation ...")

	return func(sc *godog.ScenarioContext) {
		registerValidationStepDefinitions(sc, ctx)
		sc.StepContext().After(validateAfterStepHookInitializer(ctx))
	}
}

func registerValidationStepDefinitions(ctx *godog.ScenarioContext, vCtx *stepbuilder.ValidatorContext) {
	for _, step := range stepdefinitions.GetAll() {
		handler := step.Validate(vCtx)
		for _, sentence := range step.GetSentences() {
			ctx.Step(formatStep(sentence), handler)
		}
	}
}

func validateAfterStepHookInitializer(vCtx *stepbuilder.ValidatorContext) godog.AfterStepHook {
	return func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		if status == godog.StepUndefined {
			vCtx.AddUndefinedStep(st.Text)
		}
		return ctx, err
	}
}

func validateTestSuiteInitializer(validatorCtx *stepbuilder.ValidatorContext) func(*godog.TestSuiteContext) {
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
				undefinedStepsListFormatted := strings.Join(validatorCtx.GetUndefinedSteps(), "\n")
				msg := fmt.Sprintf("List of undefined steps: \n%s", undefinedStepsListFormatted)
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
