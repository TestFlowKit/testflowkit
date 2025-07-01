package main

import (
	"testflowkit/internal/actions"
	"testflowkit/internal/config"
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
		logger.Fatal("Failed to load config", err)
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

	env, _ := cfg.GetCurrentEnvironment()
	logger.InfoFf("Available Steps: %d", len(actions.GetAllSteps()))

	logger.InfoFf("Active Environment: %s", cfg.ActiveEnvironment)
	logger.InfoFf("Frontend Base URL: %s", env.FrontendBaseURL)
	logger.InfoFf("API Base URL: %s", env.APIBaseURL)
	logger.InfoFf("Headless Mode: %t", cfg.Settings.Headless)
	logger.InfoFf("Concurrency: %d", cfg.Settings.Concurrency)
	logger.InfoFf("Report Format: %s", cfg.Settings.ReportFormat)
	logger.InfoFf("Gherkin Location: %s", cfg.Settings.GherkinLocation)
	logger.InfoFf("Test Tags: %s", cfg.Settings.Tags)
	logger.InfoFf("Default Timeout: %dms", cfg.Settings.DefaultTimeout)
	// logger.InfoFf("Page Load Timeout: %dms", cfg.Settings.PageLoadTimeout)
	logger.InfoFf("Screenshot on Failure: %t", cfg.Settings.ScreenshotOnFailure)
	// logger.InfoFf("Video Recording: %t", cfg.Settings.VideoRecording)
	logger.InfoFf("Think Time: %v", cfg.Settings.ThinkTime)
	logger.InfoFf("Elements Configured: %d page groups", len(cfg.Frontend.Elements))
	logger.InfoFf("API Endpoints: %d endpoints", len(cfg.Backend.Endpoints))

	logger.Info("--- Configuration Summary End ---\n")
}
