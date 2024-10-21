package main

import (
	"context"
	"cucumber/frontend"
	"cucumber/report"
	"log"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func main() {
	const concurrency = 2
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         concurrency,
		Format:              "pretty",
		ShowStepDefinitions: false,
		Paths:               []string{"features"},
	}

	testReport := report.New()

	testSuite := godog.TestSuite{
		Name:                 "App",
		Options:              &opts,
		TestSuiteInitializer: testSuiteInitializer(&testReport),
		ScenarioInitializer:  scenarioInitializer(&testReport),
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
func scenarioInitializer(testReport *report.Report) func(*godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		frontend.InitializeScenario(sc)
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
