package main

import (
	"testflowkit/internal/actions"
	"testflowkit/internal/config"
	stepdefinitions "testflowkit/internal/step_definitions"
	"testflowkit/pkg/logger"
)

func main() {
	args := getAppArgs()

	mode, err := args.getMode()
	if err != nil {
		logger.Fatal("Failed to get mode", err)
	}

	if mode == config.InitMode {
		actions.Execute(nil, mode)
		return
	}

	cfgPath, getcfigPathErr := args.getConfigPath()
	if getcfigPathErr != nil {
		logger.Fatal("Failed to get config path", getcfigPathErr)
	}

	configLoadErr := config.Load(cfgPath, args.getAppConfigOverrides())
	if configLoadErr != nil {
		logger.Fatal("Failed to load config", configLoadErr)
	}

	cfg, configGetErr := config.Get()
	if configGetErr != nil {
		logger.Fatal("Failed to get config", err)
	}

	displayConfigSummary(cfg)

	actions.Execute(cfg, mode)
}

func displayConfigSummary(cfg *config.Config) {
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
	logger.InfoFf("API Base URL: %s", env.APIBaseURL)
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
