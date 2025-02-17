package actions

import (
	"context"
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/frontend"
	"etoolse/pkg/gherkinparser"
	"etoolse/pkg/logger"
	"etoolse/pkg/reporters"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func Run(appConfig *config.App) {
	logger.Info("Starting tests execution ...")

	parsedFeatures := gherkinparser.Parse(appConfig.GherkinLocation)
	features := make([]godog.Feature, len(parsedFeatures))
	for i, f := range parsedFeatures {
		features[i] = godog.Feature{
			Name:     f.Name,
			Contents: f.Contents,
		}
	}

	testReport := reporters.New(appConfig.AppName, appConfig.AppDescription, appConfig.ReportFormat)
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         appConfig.GetConcurrency(),
		Format:              "pretty",
		ShowStepDefinitions: false,
		Tags:                appConfig.Tags,
		FeatureContents:     features,
	}

	testSuite := godog.TestSuite{
		Name:                 appConfig.AppName,
		Options:              &opts,
		TestSuiteInitializer: testSuiteInitializer(&testReport),
		ScenarioInitializer:  scenarioInitializer(appConfig, &testReport),
	}

	logger.Info("Running tests ...")
	testSuite.Run()
	if testReport.AreAllTestsPassed {
		logger.Success("All tests passed")
		os.Exit(0)
	}
	logger.Error("Some tests failed", []string{
		"some selectors may be missing in the configuration file",
		"Some selectors may be malformed",
		"Some selectors may no longer be available",
		"Some selectors may be incorrect",
		"Tests steps may be malformed",
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
			testReport.Start()
		})

		suiteContext.AfterSuite(func() {
			if testReport.HasScenarios() {
				testReport.Write()
			} else {
				logger.Info("No scenarios executed")
			}
		})
	}
}
func scenarioInitializer(config *config.App, testReport *reporters.Report) func(*godog.ScenarioContext) {
	logger.Info("Initializing scenario for test running ...")
	return func(sc *godog.ScenarioContext) {
		frontend.InitTestRunnerScenarios(sc, config)
		myCtx := newScenarioCtx()
		sc.StepContext().Before(beforeStepHookInitializer(&myCtx))
		sc.StepContext().After(afterStepHookInitializer(&myCtx))
		sc.After(afterScenarioHookInitializer(testReport, &myCtx))
	}
}
func afterStepHookInitializer(myCtx *myScenarioCtx) godog.AfterStepHook {
	return func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		myCtx.addStep(st.Text, status, err)
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

func (c *myScenarioCtx) addStep(title string, status godog.StepResultStatus, err error) {
	c.scenarioReport.AddStep(title, status, time.Since(c.currentStepStartTime), err)
}
