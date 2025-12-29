package actionrun

import (
	"context"
	"fmt"
	"os"
	"testflowkit/internal/actions/actionutils"
	"testflowkit/internal/config"
	stepdefinitions "testflowkit/internal/step_definitions"
	"testflowkit/internal/step_definitions/core/scenario"

	"testflowkit/pkg/browser"
	"testflowkit/pkg/gherkinparser"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/reporters"
	"testflowkit/pkg/variables"
	"time"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func Execute(appConfig *config.Config, errCfg error) {
	if errCfg != nil {
		logger.Fatal("RUN", errCfg)
	}

	actionutils.DisplayConfigSummary(appConfig)
	logger.Info("Starting tests execution ...")

	parsedFeatures := gherkinparser.Parse(appConfig.Settings.GherkinLocation)
	features := make([]godog.Feature, len(parsedFeatures))
	for i, f := range parsedFeatures {
		features[i] = godog.Feature{
			Name:     f.Name,
			Contents: f.Contents,
		}
	}

	testReport := reporters.New(appConfig.Settings.ReportFormat)

	runBeforeAllScenarios(appConfig, testReport, features)

	runMainScenarios(appConfig, testReport, features)

	runAfterAllScenarios(appConfig, testReport, features)

	if testReport.HasScenarios() {
		testReport.Write()
	}

	if !testReport.HasScenarios() {
		logger.Warn("No scenarios executed!", []string{
			"Make sure your tags are correct",
			"Make sure your gherkin files directory is configured",
		})
		os.Exit(0)
	}

	if testReport.AreAllTestsPassed {
		logger.Success("All tests passed")
		os.Exit(0)
	}

	logger.Error("Some tests failed", []string{
		"some selectors may be missing in the configuration file",
		"Some selectors may be malformed",
		"Some selectors may no longer be available",
		"Some selectors may be incorrect",
		"Teststeps may be malformed",
	}, []string{
		"please check the configuration file",
		"please check the test steps",
		"please verify the availability of the selectors",
		"please verify the correctness of the selectors",
		"please verify the correctness of the test steps",
		"please see logs for more details",
		"please see the test report for more details",
	})

	os.Exit(1)
}

func runAfterAllScenarios(appConfig *config.Config, testReport *reporters.Report, features []godog.Feature) {
	afterAllSuite := createTestSuite(createTestSuiteParams{
		appConfig:   appConfig,
		testReport:  testReport,
		features:    features,
		tags:        "@AfterAll",
		concurrency: 1,
	})
	afterAllSuite.Run()
}

func runMainScenarios(appConfig *config.Config, testReport *reporters.Report, features []godog.Feature) {
	mainTags := getTagsExludingMacros(appConfig.Settings.Tags)
	// Exclude BeforeAll and AfterAll from main run
	mainTags = fmt.Sprintf("%s && ~@BeforeAll && ~@AfterAll", mainTags)

	mainSuite := createTestSuite(createTestSuiteParams{
		appConfig:   appConfig,
		testReport:  testReport,
		features:    features,
		tags:        mainTags,
		concurrency: appConfig.GetConcurrency(),
	})
	mainSuite.Run()
}

func runBeforeAllScenarios(appConfig *config.Config, testReport *reporters.Report, features []godog.Feature) {
	beforeAllSuite := createTestSuite(createTestSuiteParams{
		appConfig:   appConfig,
		testReport:  testReport,
		features:    features,
		tags:        "@BeforeAll",
		concurrency: 1,
	})
	if status := beforeAllSuite.Run(); status != 0 {
		logger.Error("BeforeAll hooks failed", []string{
			"Setup scenarios failed",
		}, []string{
			"Check the report for details",
		})
		testReport.Write()
		os.Exit(1)
	}
}

func createTestSuite(params createTestSuiteParams) godog.TestSuite {
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         params.concurrency,
		Format:              "pretty",
		ShowStepDefinitions: false,
		Tags:                params.tags,
		FeatureContents:     params.features,
	}

	return godog.TestSuite{
		Name:                 "Test Suite",
		Options:              &opts,
		TestSuiteInitializer: testSuiteInitializer(params.testReport),
		ScenarioInitializer:  scenarioInitializer(params.appConfig, params.testReport),
	}
}

type createTestSuiteParams struct {
	appConfig   *config.Config
	testReport  *reporters.Report
	features    []godog.Feature
	tags        string
	concurrency int
}

func testSuiteInitializer(testReport *reporters.Report) func(*godog.TestSuiteContext) {
	return func(suiteContext *godog.TestSuiteContext) {
		suiteContext.BeforeSuite(func() {
			if !testReport.IsStarted() {
				testReport.Start()
			}
		})
	}
}

func scenarioInitializer(config *config.Config, testReport *reporters.Report) func(*godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		// Inject global variables here
		scenarioCtx := scenario.NewContext(config, variables.GetGlobalVariables())
		registerTestRunnerStepDefinitions(sc)
		myCtx := newScenarioCtx()
		sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
			logger.InfoFf("Running scenario: %s", sc.Name)
			ctx = scenario.WithContext(ctx, scenarioCtx)
			return ctx, nil
		})
		sc.StepContext().Before(beforeStepHookInitializer(&myCtx))
		sc.StepContext().After(afterStepHookInitializer(&myCtx, config))
		sc.After(afterScenarioHookInitializer(testReport, &myCtx))
	}
}

func afterStepHookInitializer(myCtx *myScenarioCtx, config *config.Config) godog.AfterStepHook {
	return func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		scenarioCtx := scenario.MustFromContext(ctx)

		stepText := scenario.ReplaceVariablesInString(scenarioCtx, st.Text)
		if err == nil {
			myCtx.addStep(stepText, status, nil)
			return ctx, nil
		}

		screenshotBase64 := ""
		if config.IsScreenshotOnFailureEnabled() {
			currentPage, errPage := scenarioCtx.GetCurrentPageOnly()
			if errPage == nil && currentPage != nil {
				screenshotBase64 = takeScreenshot(st, currentPage)
			}
		}

		myCtx.addStep(stepText, status, stepError{
			error:            err,
			screenshotBase64: screenshotBase64,
		})
		return ctx, err
	}
}

func takeScreenshot(st *godog.Step, currentPage browser.Page) string {
	screenshotData, screenshotErr := currentPage.Screenshot()
	if screenshotErr != nil {
		logger.Warn("Failed to take screenshot on step failure", []string{
			"step: " + st.Text,
			"error: " + screenshotErr.Error(),
		})
		return ""
	}

	base64Str, err := reporters.OptimizeAndEncodeScreenshot(screenshotData)
	if err != nil {
		logger.Warn("Failed to optimize and encode screenshot", []string{
			"step: " + st.Text,
			"error: " + err.Error(),
		})
		return ""
	}
	return base64Str
}

func beforeStepHookInitializer(myCtx *myScenarioCtx) godog.BeforeStepHook {
	return func(ctx context.Context, _ *godog.Step) (context.Context, error) {
		myCtx.currentStepStartTime = time.Now()
		return ctx, nil
	}
}

func afterScenarioHookInitializer(testReport *reporters.Report, myCtx *myScenarioCtx) godog.AfterScenarioHook {
	return func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		myCtx.scenarioReport.SetTitle(sc.Name)

		myCtx.scenarioReport.End()
		testReport.AddScenario(myCtx.scenarioReport)

		scenario.MustFromContext(ctx).Done()

		return ctx, err
	}
}

func newScenarioCtx() myScenarioCtx {
	return myScenarioCtx{
		scenarioReport: reporters.NewScenario(),
	}
}

func registerTestRunnerStepDefinitions(ctx *godog.ScenarioContext) {
	for _, step := range stepdefinitions.GetAll() {
		handler := step.GetDefinition()
		for _, sentence := range step.GetSentences() {
			ctx.Step(actionutils.FormatStep(sentence), handler)
		}
	}
}

func getTagsExludingMacros(tags string) string {
	excludeMacros := "~" + gherkinparser.MacroTag
	if tags == "" {
		return excludeMacros
	}

	return fmt.Sprintf("%s && %s", tags, excludeMacros)
}

type myScenarioCtx struct {
	currentStepStartTime time.Time
	scenarioReport       reporters.Scenario
}

func (c *myScenarioCtx) addStep(title string, status godog.StepResultStatus, err error) {
	c.scenarioReport.AddStep(title, status, time.Since(c.currentStepStartTime), err)
}

type stepError struct {
	error
	screenshotBase64 string
}

func (se stepError) ScreenshotBase64() string {
	return se.screenshotBase64
}
