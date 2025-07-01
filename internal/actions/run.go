package actions

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/steps_definitions/core/scenario"
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
				logger.Error("Failed to remove screenshot directory: "+err.Error(), []string{
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
		if err == nil {
			myCtx.addStep(st.Text, status, nil)
			return ctx, nil
		}

		screenshotPath := ""
		if config.Settings.ScreenshotOnFailure {
			currentPage := scenario.MustFromContext(ctx).GetCurrentPageOnly()
			if currentPage != nil {
				screenshotPath = takeScreenshot(st, currentPage)
			}
		}

		myCtx.addStep(st.Text, status, stepError{
			error:          err,
			screenshotPath: screenshotPath,
		})
		return ctx, err
	}
}

func takeScreenshot(st *godog.Step, currentPage common.Page) string {
	safeStepName := sanitizeFilename(st.Text)
	timestamp := time.Now().Format("20060102_150405_000")
	screenshotPath := filepath.Join(screenshotDir, safeStepName+"_"+timestamp+".png")

	screenshotData, screenshotErr := currentPage.Screenshot()
	if screenshotErr != nil {
		logger.Warn("Failed to take screenshot on step failure", []string{
			"step: " + st.Text,
			"error: " + screenshotErr.Error(),
		})
		screenshotPath = ""
	} else {
		if writeErr := os.WriteFile(screenshotPath, screenshotData, 0600); writeErr != nil {
			logger.Warn("Failed to save screenshot file", []string{
				"step: " + st.Text,
				"path: " + screenshotPath,
				"error: " + writeErr.Error(),
			})
			screenshotPath = ""
		}
	}
	return screenshotPath
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
	for _, step := range GetAllSteps() {
		handler := step.GetDefinition()
		for _, sentence := range step.GetSentences() {
			ctx.Step(core.ConvertWildcards(sentence), handler)
		}
	}
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
	screenshotPath string
}

func (se stepError) ScreenshotPath() string {
	return se.screenshotPath
}

func sanitizeFilename(input string) string {
	re := regexp.MustCompile(`[^\w\s-]`)
	cleaned := re.ReplaceAllString(input, "_")

	cleaned = strings.ReplaceAll(cleaned, " ", "_")

	re = regexp.MustCompile(`_+`)
	cleaned = re.ReplaceAllString(cleaned, "_")

	re = regexp.MustCompile(`^_+|_+$`)
	cleaned = re.ReplaceAllString(cleaned, "")
	const maxLength = 100
	if len(cleaned) > maxLength {
		cleaned = cleaned[:maxLength]
	}

	return cleaned
}
