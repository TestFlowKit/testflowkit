package main

import (
	"context"
	"cucumber/config"
	"cucumber/frontend"
	"cucumber/report"
	"log"
	"time"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func main() {
	log.Println("Starting tests execution ...")
	cliConfig := config.Init()

	if cliConfig.Mode == config.RunMode {
		var opts = godog.Options{
			Output:              &buffer.Writer{},
			Concurrency:         cliConfig.GetConcurrency(),
			Format:              "pretty",
			ShowStepDefinitions: false,
			Tags:                cliConfig.Tags,
			Paths:               []string{cliConfig.GherkinLocation},
		}

		testReport := report.New(cliConfig.AppName, cliConfig.AppDescription, cliConfig.ReportFormat)
		testSuite := godog.TestSuite{
			Name:                 cliConfig.AppName,
			Options:              &opts,
			TestSuiteInitializer: testSuiteInitializer(&testReport),
			ScenarioInitializer:  scenarioInitializer(cliConfig, &testReport),
		}

		log.Println("Running tests ...")
		status := testSuite.Run()
		log.Println("Tests execution finished")
		if status != 0 {
			log.Fatalf("zero status code expected, %d received", status)
		}
	} else {
		log.Fatalf("unknown mode: %s", cliConfig.Mode)
	}
}

func testSuiteInitializer(testReport *report.Report) func(*godog.TestSuiteContext) {
	return func(suiteContext *godog.TestSuiteContext) {
		suiteContext.BeforeSuite(func() {
			testReport.Start()
		})

		suiteContext.AfterSuite(func() {
			testReport.Write()
		})
	}
}
func scenarioInitializer(config *config.AppConfig, testReport *report.Report) func(*godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		frontend.InitializeScenario(sc, config)
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

func afterScenarioHookInitializer(testReport *report.Report, myCtx *myScenarioCtx) godog.AfterScenarioHook {
	return func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		myCtx.scenarioReport.SetTitle(sc.Name)

		myCtx.scenarioReport.End()
		testReport.AddScenario(myCtx.scenarioReport)
		return ctx, err
	}
}

func newScenarioCtx() myScenarioCtx {
	return myScenarioCtx{
		scenarioReport: report.NewScenario(),
	}
}

type myScenarioCtx struct {
	currentStepStartTime time.Time
	scenarioReport       report.Scenario
}

func (c *myScenarioCtx) addStep(title string, status godog.StepResultStatus, err error) {
	c.scenarioReport.AddStep(title, status, time.Since(c.currentStepStartTime), err)
}
