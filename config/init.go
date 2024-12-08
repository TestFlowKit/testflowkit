package config

import (
	"log"
)

func Init() *AppConfig {
	argsConfig, fileConfig := getAppArgs(), appFileConfig{}

	fileConfig.InitByFilePath(argsConfig.ClIConfigPath)
	FrontConfig{}.init(argsConfig.FrontendConfigPath)
	appConfig := InitAppConfig(argsConfig, fileConfig)

	log.Println("--- configuration resume ---")
	log.Println("cli config path: ", argsConfig.ClIConfigPath)
	log.Println("frontend config path: ", argsConfig.FrontendConfigPath)

	log.Println("app name: ", appConfig.AppName)
	log.Println("app description: ", appConfig.AppDescription)
	log.Println("app version: ", appConfig.AppVersion)
	log.Println("app tags: ", appConfig.Tags)
	log.Println("app gherkin location: ", appConfig.GherkinLocation)
	log.Println("app report format: ", appConfig.ReportFormat)
	log.Println("app concurrency: ", appConfig.GetConcurrency())
	log.Println("app slow motion: ", appConfig.GetSlowMotion())
	log.Println("app test suite timeout: ", appConfig.Timeout)
	log.Println("app headless mode: ", appConfig.IsHeadlessModeEnabled())

	log.Println("--- configuration resume end ---")
	log.Println()

	return appConfig
}
