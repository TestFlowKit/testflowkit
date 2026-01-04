package actionutils

import (
	"strings"
	"testflowkit/internal/config"
	stepdefinitions "testflowkit/internal/step_definitions"
	"testflowkit/internal/step_definitions/core"
	"testflowkit/pkg/logger"
)

func FormatStep(sentence string) string {
	cleanedSentence := strings.TrimPrefix(sentence, "^")
	cleanedSentence = strings.TrimSuffix(cleanedSentence, "$")

	pattern := "^" + core.ConvertWildcards(cleanedSentence) + "$"
	return pattern
}

func DisplayConfigSummary(cfg *config.Config) {
	if cfg == nil {
		return
	}
	logger.InfoFf("Testflowkit version %s", cfg.GetVersion())

	logger.Info("--- Configuration Summary ---")

	logger.InfoFf("Available Steps: %d", len(stepdefinitions.GetAll()))

	logger.InfoFf("Active Environment: %s", cfg.ActiveEnvironment)
	logger.InfoFf("Concurrency: %d", cfg.Settings.Concurrency)
	logger.InfoFf("Report Format: %s", cfg.Settings.ReportFormat)
	logger.InfoFf("Gherkin Location: %s", cfg.Settings.GherkinLocation)
	logger.InfoFf("Think Time: %v", cfg.Settings.ThinkTime)
	logger.InfoFf("Test Tags: %s", cfg.Settings.Tags)

	if cfg.IsFrontendDefined() {
		displayFrontSummary(cfg)
	}

	env, _ := cfg.GetCurrentEnvironment()

	logger.InfoFf("API Base URL: %s", env.RestAPIBaseURL)
	logger.InfoFf("API GraphQL Endpoint: %s", env.GraphQLEndpoint)
	logger.InfoFf("API Endpoints: %d endpoints", len(cfg.Backend.Endpoints))

	logger.Info("--- Configuration Summary End ---\n")
}

func displayFrontSummary(conf *config.Config) {
	env, _ := conf.GetCurrentEnvironment()

	frontConf := conf.Frontend
	logger.InfoFf("Headless Mode: %t", frontConf.Headless)
	logger.InfoFf("Default Timeout: %dms", frontConf.DefaultTimeout)
	logger.InfoFf("Frontend Base URL: %s", env.FrontendBaseURL)
	logger.InfoFf("Screenshot on Failure: %t", frontConf.ScreenshotOnFailure)
	logger.InfoFf("Elements Configured: %d page groups", len(frontConf.Elements))
}
