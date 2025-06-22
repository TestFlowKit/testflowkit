package config

import (
	"testflowkit/pkg/logger"
)

const defaultCliConfigPath = "cliConfig.yml"
const testingYamlPath = "$.configuration"
const reportingYamlPath = "$.reporting"

type cliConfig struct {
	reportingConfig
	Timeout         string `yaml:"timeout"`
	SlowMotion      string `yaml:"slowMotion"`
	GherkinLocation string `yaml:"gherkin_location"`
}

func (c *cliConfig) init(content string) {
	testing := c.getTestingConfig(content)
	reporting := c.getReportingConfig(content)

	c.Timeout = testing.Timeout
	c.SlowMotion = testing.SlowMotion
	c.GherkinLocation = testing.GherkinLocation
	c.reportingConfig = reporting
}

func (c *cliConfig) getTestingConfig(file string) testingConfig {
	config := testingConfig{}
	err := getConfig(file, testingYamlPath, &config)
	if err != nil {
		logger.Fatal("testing config getting failed", err)
	}

	return config
}

func (c *cliConfig) getReportingConfig(file string) reportingConfig {
	config := reportingConfig{}
	err := getConfig(file, reportingYamlPath, &config)
	if err != nil {
		logger.Fatal("reporting config getting failed", err)
	}
	return config
}
