package config

import (
	"etoolse/internal/config/testsConfig"
	"log"
	"os"

	"github.com/alexflint/go-arg"
)

func Init() *App {
	argsConfig := getAppArgs()

	if argsConfig.Run != nil {
		cliFileConfig := getCLIFileConfig(argsConfig.Run.ClIConfigPath)
		initFrontTestsConfig(argsConfig.Run.FrontendConfigPath)

		return initRunConfig(argsConfig, cliFileConfig)
	}

	if argsConfig.Validate != nil {
		fileConfig := getCLIFileConfig(argsConfig.Validate.ClIConfigPath)
		initFrontTestsConfig(argsConfig.Validate.FrontendConfigPath)

		curr := initAppConfig(argsConfig, fileConfig)
		curr.Tags = argsConfig.Validate.Tags
		curr.Mode = ValidationMode

		return curr
	}

	if argsConfig.Init != nil {
		return &App{
			Mode: InitMode,
			appDetailsConfig: appDetailsConfig{
				AppName:        argsConfig.Init.AppName,
				AppDescription: argsConfig.Init.AppDescription,
				AppVersion:     argsConfig.Init.AppVersion,
			},
		}
	}

	return nil
}

func getCLIFileConfig(filePath string) cliConfig {
	configFileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("CLI config file not found")
	}
	CLIFileConfig := cliConfig{}

	CLIFileConfig.init(string(configFileContent))

	return CLIFileConfig
}

func getAppArgs() argsConfig {
	config := argsConfig{}
	arg.MustParse(&config)
	return config
}

func initFrontTestsConfig(filePath string) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("config file not found")
	}

	testsConfig.Init(string(configFile))
}
