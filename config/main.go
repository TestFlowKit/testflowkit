package config

import (
	"log"
	"os"

	"github.com/alexflint/go-arg"
)

func Init() *AppConfig {
	argsConfig := getAppArgs()
	if argsConfig.Run != nil {
		return initRunConfig(argsConfig)
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

	fileConfig := getFileConfig(argsConfig.Run.ClIConfigPath)

	FrontConfig{}.init(argsConfig.Run.FrontendConfigPath)

	return initRunConfig(argsConfig, fileConfig)
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

	validateSubcommand(args)

	return args
}
