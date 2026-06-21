package actionutils

import (
	"strings"
	"testflowkit/internal/config"
	stepdefinitions "testflowkit/internal/step_definitions"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/stepexpr"
)

func FormatStep(sentence string) string {
	cleanedSentence := strings.TrimPrefix(sentence, "^")
	cleanedSentence = strings.TrimSuffix(cleanedSentence, "$")

	pattern := "^" + stepexpr.ConvertWildcards(cleanedSentence) + "$"
	return pattern
}

func DisplayConfigSummary(cfg *config.Config) {
	if cfg == nil {
		return
	}
	logger.Infof("Testflowkit version %s", cfg.GetVersion())

	logger.Info("--- Configuration Summary ---")

	logger.Infof("Available Steps: %d", len(stepdefinitions.GetAll()))

	logger.Infof("Concurrency: %d", cfg.Settings.Concurrency)
	logger.Infof("Report Format: %s", cfg.Settings.ReportFormat)
	logger.Infof("Gherkin Location: %s", cfg.Settings.GherkinLocation)
	logger.Infof("Test Tags: %s", cfg.Settings.Tags)
	logger.Infof("Environment Variables: %d defined", len(cfg.Env))

	if cfg.IsFrontendDefined() {
		displayFrontSummary(cfg)
	}

	if cfg.APIs != nil && cfg.APIs.Definitions != nil {
		logger.Infof("APIs configured: %d", len(cfg.APIs.Definitions))
		for apiName, apiDef := range cfg.APIs.Definitions {
			logger.Infof("  - %s (%s)", apiName, apiDef.Type)
		}
	}

	logger.Info("--- Configuration Summary End ---\n")
}

func displayFrontSummary(conf *config.Config) {
	frontConf := conf.Frontend
	logger.Infof("Headless Mode: %t", frontConf.Headless)
	logger.Infof("Default Timeout: %dms", frontConf.DefaultTimeout)
	logger.Infof("Frontend Base URL: %s", conf.GetFrontendBaseURL())
	logger.Infof("Pages Configured: %d", len(frontConf.Pages))
	logger.Infof("Think Time Between Actions: %dms", frontConf.ThinkTime)
	logger.Infof("Screenshot on Failure: %t", frontConf.ScreenshotOnFailure)
	logger.Infof("Elements Configured: %d page groups", len(frontConf.Elements))
}
