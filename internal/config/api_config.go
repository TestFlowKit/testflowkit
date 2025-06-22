package config

import (
	"testflowkit/pkg/logger"
)

var apiConfigContent string

const apiConfigYamlPath = "$.default"

func InitAPIConfig(fileContent string) {
	apiConfigContent = fileContent
}

func GetAPIConfig() (APIConfig, error) {
	var cfg APIConfig
	err := getConfig(apiConfigContent, apiConfigYamlPath, &cfg)
	if err != nil {
		logger.Fatal("API config getting failed", err)
	}
	return cfg, err
}
