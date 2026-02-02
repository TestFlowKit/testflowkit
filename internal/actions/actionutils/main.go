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

	logger.InfoFf("Concurrency: %d", cfg.Settings.Concurrency)
	logger.InfoFf("Report Format: %s", cfg.Settings.ReportFormat)
	logger.InfoFf("Gherkin Location: %s", cfg.Settings.GherkinLocation)
	logger.InfoFf("Test Tags: %s", cfg.Settings.Tags)
	logger.InfoFf("Environment Variables: %d defined", len(cfg.Env))

	if cfg.IsFrontendDefined() {
		displayFrontSummary(cfg)
	}

	if cfg.APIs != nil && cfg.APIs.Definitions != nil {
		logger.InfoFf("APIs configured: %d", len(cfg.APIs.Definitions))
		for apiName, apiDef := range cfg.APIs.Definitions {
			logger.InfoFf("  - %s (%s)", apiName, apiDef.Type)
		}
	}

	logger.Info("--- Configuration Summary End ---\n")
}

func displayFrontSummary(conf *config.Config) {
	frontConf := conf.Frontend
	logger.InfoFf("Headless Mode: %t", frontConf.Headless)
	logger.InfoFf("Default Timeout: %dms", frontConf.DefaultTimeout)
	logger.InfoFf("Frontend Base URL: %s", conf.GetFrontendBaseURL())
	logger.InfoFf("Pages Configured: %d", len(frontConf.Pages))
	logger.InfoFf("Think Time Between Actions: %dms", frontConf.ThinkTime)
	logger.InfoFf("Screenshot on Failure: %t", frontConf.ScreenshotOnFailure)
	logger.InfoFf("Elements Configured: %d page groups", len(frontConf.Elements))
}
