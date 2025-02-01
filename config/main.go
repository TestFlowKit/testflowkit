package config

import (
	"log"
	"os"

	"github.com/alexflint/go-arg"
)

func Init() *AppConfig {
	argsConfig := getAppArgs()
	if argsConfig.Run != nil {
		fileConfig := getFileConfig(argsConfig.Run.ClIConfigPath)
		FrontConfig{}.init(argsConfig.Run.FrontendConfigPath)
		return initRunConfig(argsConfig, fileConfig)
	}

	if argsConfig.Validate != nil {
		fileConfig := getFileConfig(argsConfig.Validate.ClIConfigPath)
		FrontConfig{}.init(argsConfig.Validate.FrontendConfigPath)
		curr := initAppConfig(argsConfig, fileConfig)
		curr.Tags = argsConfig.Validate.Tags
		curr.Mode = ValidationMode
		return curr
	}

	if argsConfig.Init != nil {
		return &AppConfig{
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

func getFileConfig(filePath string) appFileConfig {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("config file not found")
	}

	fileConfig := appFileConfig{}
	fileConfig.InitByFileContent(string(file))
	return fileConfig
}

func getAppArgs() appArgsConfig {
	args := appArgsConfig{}
	arg.MustParse(&args)
	return args
}
