package actions

import (
	"context"
	"os"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/frontend"
	"testflowkit/pkg/gherkinparser"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/reporters"
	"time"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

const screenshotDir = "report/screenshots"

func run(appConfig *config.Config) {
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
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         appConfig.GetConcurrency(),
		Format:              "pretty",
		ShowStepDefinitions: false,
		Tags:                appConfig.Settings.Tags,
		FeatureContents:     features,
	}

	testSuite := godog.TestSuite{
		Name:                 "Test Suite",
		Options:              &opts,
		TestSuiteInitializer: testSuiteInitializer(&testReport),
		ScenarioInitializer:  scenarioInitializer(appConfig, &testReport),
	}

	logger.Info("Running tests ...")
	testSuite.Run()

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

func testSuiteInitializer(testReport *reporters.Report) func(*godog.TestSuiteContext) {
	return func(suiteContext *godog.TestSuiteContext) {

		suiteContext.BeforeSuite(func() {
			err := os.RemoveAll(screenshotDir)
			if err != nil {
				logger.Error("Failed to remove screenshot directory", []string{
					"please check the permissions of the directory",
					"please check the directory path",
				}, []string{
					"please see logs for more details",
					"please see the test report for more details",
				})
				os.Exit(1)
			}
			mkdirErr := os.MkdirAll(screenshotDir, 0755)
			if mkdirErr != nil {
				logger.Error("Failed to create screenshot directory", []string{
					"please check the permissions of the directory",
					"please check the directory path",
				}, []string{
					"please see logs for more details",
					"please see the test report for more details",
				})
				os.Exit(1)
			}
			testReport.Start()
		})

		suiteContext.AfterSuite(func() {
			if testReport.HasScenarios() {
				testReport.Write()
			}
		})
	}
}
func scenarioInitializer(config *config.Config, testReport *reporters.Report) func(*godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		scenarioCtx := scenario.NewContext(config)
		frontend.InitTestRunnerScenarios(sc, config)
		myCtx := newScenarioCtx()
		sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
			logger.InfoFf("Running scenario: %s", sc.Name)
			ctx = scenario.WithContext(ctx, scenarioCtx)
			return ctx, nil
		})
		sc.StepContext().Before(beforeStepHookInitializer(&myCtx))
		sc.StepContext().After(afterStepHookInitializer(&myCtx))
		sc.After(afterScenarioHookInitializer(testReport, &myCtx))
	}
}
func afterStepHookInitializer(myCtx *myScenarioCtx) godog.AfterStepHook {
	return func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		myCtx.addStep(st.Text, status, stepError{
			error:          err,
			screenshotPath: screenshotDir + "/" + st.Text + ".png",
		})
		return ctx, err
	}
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
		return ctx, err
	}
}

func newScenarioCtx() myScenarioCtx {
	return myScenarioCtx{
		scenarioReport: reporters.NewScenario(),
	}
}

type myScenarioCtx struct {
	currentStepStartTime time.Time
	scenarioReport       reporters.Scenario
}

func (c *myScenarioCtx) addStep(title string, status godog.StepResultStatus, err stepError) {
	c.scenarioReport.AddStep(title, status, time.Since(c.currentStepStartTime), err)
}

type stepError struct {
	error
	screenshotPath string
}
