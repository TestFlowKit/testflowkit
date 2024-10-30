package main

import (
	"context"
	"cucumber/config"
	"cucumber/frontend"
	"cucumber/report"
	"log"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func main() {
	cliConfig := config.Init()

	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         cliConfig.GetConcurrency(),
		Format:              "pretty",
		ShowStepDefinitions: false,
		Tags:                cliConfig.GetTagsExpression(),
		Paths:               []string{cliConfig.GherkinLocation},
	}

	testReport := report.New(cliConfig.AppName, cliConfig.AppDescription, cliConfig.ReportEnabled, cliConfig.ReportFormat)
	testSuite := godog.TestSuite{
		Name:                 cliConfig.AppName,
		Options:              &opts,
		TestSuiteInitializer: testSuiteInitializer(&testReport),
		ScenarioInitializer:  scenarioInitializer(cliConfig, &testReport),
	}

	status := testSuite.Run()

	if status != 0 {
		log.Fatalf("zero status code expected, %d received", status)
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
func scenarioInitializer(config config.ClI, testReport *report.Report) func(*godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		frontend.InitializeScenario(sc, config)
		scenarioReport := report.NewScenario()

		sc.StepContext().After(afterStepHookInitializer(&scenarioReport))
		sc.After(afterScenarioHookInitializer(testReport, &scenarioReport))
	}
}
func afterStepHookInitializer(scenarioReport *report.Scenario) godog.AfterStepHook {
	return func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		scenarioReport.AddStep(st.Text, status, err)
		return ctx, err
	}
}
func afterScenarioHookInitializer(testReport *report.Report, scenarioReport *report.Scenario) godog.AfterScenarioHook {
	return func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		scenarioReport.SetTitle(sc.Name)

		scenarioReport.End()
		testReport.AddScenario(*scenarioReport)
		return ctx, err
	}
}
