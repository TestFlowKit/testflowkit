package config

import (
	"log"
)

func Init() *AppConfig {
	argsConfig := getAppArgs()

	return initRunConfig(argsConfig)
}

func initRunConfig(argsConfig appArgsConfig) *AppConfig {
	fileConfig := appFileConfig{}
	appConfig := InitAppConfig(argsConfig, fileConfig)
	runConfig := argsConfig.Run

	fileConfig.InitByFilePath(runConfig.ClIConfigPath)
	FrontConfig{}.init(runConfig.FrontendConfigPath)

	log.Println("--- configuration resume ---")
	log.Println("cli config path: ", runConfig.ClIConfigPath)
	log.Println("frontend config path: ", runConfig.FrontendConfigPath)

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

	log.Print("--- configuration resume end ---\n\n")

	return appConfig
}
