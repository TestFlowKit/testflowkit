package config

import (
	"log"
	"os"
)

type appFileConfig struct {
	reportingConfig
	AppName         string `yaml:"app_name"`
	AppDescription  string `yaml:"app_description,omitempty"`
	Timeout         string `yaml:"timeout"`
	SlowMotion      string `yaml:"slowMotion"`
	GherkinLocation string `yaml:"gherkin_location"`
}

func (c *appFileConfig) InitByFilePath(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("config file not found")
	}
	c.InitByFileContent(string(file))
}

func (c *appFileConfig) InitByFileContent(content string) {
	appDetails := c.getAppDetailsConfig(content)
	testing := c.getTestingConfig(content)
	reporting := c.getReportingConfig(content)

	c.AppName = appDetails.AppName
	c.AppDescription = appDetails.AppDescription
	c.Timeout = testing.Timeout
	c.SlowMotion = testing.SlowMotion
	c.GherkinLocation = testing.GherkinLocation
	c.reportingConfig = reporting
}

func (c *appFileConfig) getAppDetailsConfig(file string) appDetailsConfig {
	config := appDetailsConfig{}
	err := getConfig(file, "$.application", &config)
	if err != nil {
		return config
	}

	return config
}

func (c *appFileConfig) getTestingConfig(file string) testingConfig {
	config := testingConfig{}
	err := getConfig(file, "$.configuration", &config)
	if err != nil {
		log.Println(err)
		log.Panicln("testsuite config getting failed")
	}

	return config
}

func (c *appFileConfig) getReportingConfig(file string) reportingConfig {
	config := reportingConfig{}
	err := getConfig(file, "$.reporting", &config)
	if err != nil {
		log.Println(err)
		log.Panicln("reporting config getting failed")
	}
	return config
}
