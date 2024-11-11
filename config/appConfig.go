package config

import (
	"errors"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

type ClI struct {
	appConfig
	reportingConfig
	testingConfig
}

func (c *ClI) InitByFilePath(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("config file not found")
	}
	c.InitByFileContent(string(file))
}

func (c *ClI) InitByFileContent(content string) {
	c.appConfig = c.getAppDetailsConfig(content)
	c.reportingConfig = c.getReportingConfig(content)
	c.testingConfig = c.getTestingConfig(content)
}

func (c *ClI) getAppDetailsConfig(file string) appConfig {
	config := appConfig{}
	err := getConfig(file, "$.application", &config)
	if err != nil {
		return config
	}

	return config
}

func (c *ClI) getReportingConfig(file string) reportingConfig {
	config := reportingConfig{}
	err := getConfig(file, "$.reporting", &config)
	if err != nil && errors.Is(err, yaml.ErrNotFoundNode) {
		config.ReportEnabled = false
	}
	return config
}

func (c *ClI) getTestingConfig(file string) testingConfig {
	config := testingConfig{}
	err := getConfig(file, "$.configuration", &config)
	if err != nil {
		log.Println(err)
		log.Panicln("testsuite config getting failed")
	}

	if config.Timeout == "" {
		config.Timeout = "1s"
	}

	return config
}

func (c *ClI) GetConcurrency() int {
	if c.displayBrowser {
		return 0
	}
	return c.Parallel
}
