package config

import (
	"log"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
)

type appDetailsConfig struct {
	AppName        string `yaml:"app_name"`
	AppDescription string `yaml:"app_description,omitempty"`
	AppVersion     string `yaml:"app_version"`
}

type reportingConfig struct {
	ReportFormat string `yaml:"report_format"`
}

type testingConfig struct {
	Timeout         string `yaml:"timeout"`
	Tags            string `yaml:"tags"`
	SlowMotion      string `yaml:"slowMotion"`
	Parallel        int
	Headless        bool
	GherkinLocation string `yaml:"gherkin_location"`
}

func (c *testingConfig) IsHeadlessModeEnabled() bool {
	return c.Headless
}

func (c *testingConfig) GetSlowMotion() time.Duration {
	if c.Headless {
		return 0
	}

	duration, err := time.ParseDuration(c.SlowMotion)
	if err != nil {
		log.Panicf("%s is not correct duration", c.SlowMotion)
	}

	return duration
}

type configType interface {
	testingConfig | reportingConfig | appDetailsConfig
}

// TODO: rename in order to specify the type of the config
func getConfig[T configType](file, path string, config *T) error {
	configPath, err := yaml.PathString(path)
	if err != nil {
		return err
	}

	if err = configPath.Read(strings.NewReader(file), config); err != nil {
		return err
	}

	return nil
}
