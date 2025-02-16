package config

import (
	"etoolse/internal/config/testsconfig"
	"etoolse/pkg/logger"
	"os"

	"github.com/alexflint/go-arg"
)

func Init() *App {
	appArgsConfig := getAppArgs()

	if appArgsConfig.Run != nil {
		cliFileConfig := getCLIFileConfig(appArgsConfig.Run.ClIConfigPath)
		initFrontTestsConfig(appArgsConfig.Run.FrontendConfigPath)

		return initAppConfig(appArgsConfig, cliFileConfig, RunMode)
	}

	if appArgsConfig.Validate != nil {
		fileConfig := getCLIFileConfig(appArgsConfig.Validate.ClIConfigPath)
		initFrontTestsConfig(appArgsConfig.Validate.FrontendConfigPath)

		curr := initAppConfig(appArgsConfig, fileConfig, ValidationMode)
		curr.Tags = appArgsConfig.Validate.Tags

		return curr
	}

	if appArgsConfig.Init != nil {
		return &App{
			Mode: InitMode,
			appDetailsConfig: appDetailsConfig{
				AppName:        appArgsConfig.Init.AppName,
				AppDescription: appArgsConfig.Init.AppDescription,
				AppVersion:     appArgsConfig.Init.AppVersion,
			},
		}
	}

	return nil
}

func getCLIFileConfig(filePath string) cliConfig {
	configFileContent, err := os.ReadFile(filePath)
	if err != nil {
		logger.Fatal("CLI config file not found", err)
	}
	cliFileConfig := cliConfig{}

	cliFileConfig.init(string(configFileContent))

	return cliFileConfig
}

func getAppArgs() argsConfig {
	config := argsConfig{}
	arg.MustParse(&config)
	return config
}

func initFrontTestsConfig(filePath string) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		logger.Fatal("config file not found", err)
	}

	testsconfig.Init(string(configFile))
}
