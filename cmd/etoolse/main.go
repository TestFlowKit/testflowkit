package main

import (
	"etoolse/internal/actions"
	"etoolse/internal/config"
	"log"
)

func main() {
	cliConfig := config.Init()
	if cliConfig == nil {
		log.Fatal("no mode specified")
	}

	modes := map[config.Mode]actions.Type{
		config.RunMode:        runMode,
		config.InitMode:       actions.Init,
		config.ValidationMode: actions.Validate,
	}

	if mode, ok := modes[cliConfig.Mode]; ok {
		mode(cliConfig)
	} else {
		log.Fatalf("unknown mode: %s", cliConfig.Mode)
	}
}

func runMode(appConfig *config.App) {
	log.Println("--- configuration resume ---")

	log.Println("app name: ", appConfig.AppName)
	log.Println("app description: ", appConfig.AppDescription)
	log.Println("app version: ", appConfig.AppVersion)
	log.Println("app tags: ", appConfig.Tags)
	log.Println("app gherkin location: ", appConfig.GherkinLocation)
	log.Println("app reporters format: ", appConfig.ReportFormat)
	log.Println("app concurrency: ", appConfig.GetConcurrency())
	log.Println("app slow motion: ", appConfig.GetSlowMotion())
	log.Println("app test suite timeout: ", appConfig.Timeout)
	log.Println("app headless mode: ", appConfig.IsHeadlessModeEnabled())

	log.Print("--- configuration resume end ---\n\n")

	actions.Run(appConfig)
}
