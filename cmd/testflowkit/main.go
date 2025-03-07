package main

import (
	"fmt"
	"testflowkit/internal/actions"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

func main() {
	cliConfig := config.Init()
	if cliConfig == nil {
		logger.Fatal("no mode specified", nil)
	}

	modes := map[config.Mode]actions.Type{
		config.RunMode:        runMode,
		config.InitMode:       actions.Init,
		config.ValidationMode: actions.Validate,
	}

	if mode, ok := modes[cliConfig.Mode]; ok {
		mode(cliConfig)
	} else {
		logger.Fatal(fmt.Sprintf("unknown mode: %s", cliConfig.Mode), nil)
	}
}

func runMode(appConfig *config.App) {
	logger.Info("--- configuration resume ---")

	logger.InfoFf("app name: %s", appConfig.AppName)
	logger.InfoFf("app description: %s", appConfig.AppDescription)
	logger.InfoFf("app version: %s", appConfig.AppVersion)
	logger.InfoFf("app tags: %s", appConfig.Tags)
	logger.InfoFf("app gherkin location: %s", appConfig.GherkinLocation)
	logger.InfoFf("app reporters format: %s", appConfig.ReportFormat)
	logger.InfoFf("app concurrency: %d", appConfig.GetConcurrency())
	logger.InfoFf("app slow motion: %s", appConfig.GetSlowMotion())
	logger.InfoFf("app test suite timeout: %s", appConfig.Timeout)
	logger.InfoFf("app headless mode: %t", appConfig.IsHeadlessModeEnabled())

	logger.Info("--- configuration resume end ---\n\n")

	actions.Run(appConfig)
}
