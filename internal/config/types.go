package config

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	ActiveEnvironment string `yaml:"active_environment" validate:"required"`

	Settings GlobalSettings `yaml:"settings" validate:"required"`

	Environments map[string]Environment `yaml:"environments" validate:"required,min=1"`

	Frontend FrontendConfig `yaml:"frontend" validate:"required"`
}

type GlobalSettings struct {
	DefaultTimeout int `yaml:"default_timeout" validate:"min=1000,max=300000"`

	// PageLoadTimeout int `yaml:"page_load_timeout" validate:"min=1000,max=300000"`

	// ScreenshotOnFailure bool `yaml:"screenshot_on_failure"`

	// VideoRecording bool `yaml:"video_recording"`

	Concurrency int `yaml:"concurrency" validate:"min=1,max=20"`

	Headless bool `yaml:"headless"`

	SlowMotion int `yaml:"slow_motion" validate:"omitempty" `

	ReportFormat string `yaml:"report_format" validate:"oneof=html json junit"`

	GherkinLocation string `yaml:"gherkin_location" validate:"required"`

	Tags string `yaml:"tags"`
}

type Environment struct {
	FrontendBaseURL string `yaml:"frontend_base_url" validate:"required,url"`
}

type FrontendConfig struct {
	Elements map[string]map[string][]string `yaml:"elements" validate:"required"`

	Pages map[string]string `yaml:"pages" validate:"required"`
}

type Endpoint struct {
	Method string `yaml:"method" validate:"required,oneof=GET POST PUT DELETE PATCH HEAD OPTIONS"`

	Path string `yaml:"path" validate:"required"`

	Description string `yaml:"description" validate:"required"`
}

func (c *Config) GetCurrentEnvironment() (Environment, error) {
	env, exists := c.Environments[c.ActiveEnvironment]
	if !exists {
		return Environment{}, fmt.Errorf("active environment '%s' not found in configuration", c.ActiveEnvironment)
	}
	return env, nil
}

func (c *Config) GetFrontendURL(page string) (string, error) {
	env, err := c.GetCurrentEnvironment()
	if err != nil {
		return "", err
	}

	if pagePath, ok := c.Frontend.Pages[GetLabelKey(page)]; ok {
		if strings.HasPrefix(pagePath, "http://") || strings.HasPrefix(pagePath, "https://") {
			return pagePath, nil
		}

		fullURL, errJoin := url.JoinPath(env.FrontendBaseURL, pagePath)
		if errJoin != nil {
			return "", err
		}

		return fullURL, nil
	}

	return env.FrontendBaseURL, nil
}

func (c *Config) GetElementSelectors(page, elementName string) []string {
	elementName = GetLabelKey(elementName)
	if pageElements, isPageElement := c.Frontend.Elements[page]; isPageElement {
		if selectors, ok := pageElements[elementName]; ok {
			return selectors
		}
	}

	if commonElements, isCommonElement := c.Frontend.Elements["common"]; isCommonElement {
		if selectors, ok := commonElements[elementName]; ok {
			return selectors
		}
	}

	return []string{}
}

func (c *Config) GetSlowMotion() time.Duration {
	if c.Settings.SlowMotion == 0 || c.Settings.Headless {
		return 0
	}

	duration, err := time.ParseDuration(fmt.Sprintf("%dms", c.Settings.SlowMotion))
	if err != nil {
		log.Printf("Invalid slow motion duration: %d, using 0", c.Settings.SlowMotion)
		return 0
	}
	return duration
}

func (c *Config) IsHeadlessModeEnabled() bool {
	return c.Settings.Headless
}

func (c *Config) GetConcurrency() int {
	if c.Settings.Concurrency <= 0 {
		return 1
	}
	return c.Settings.Concurrency
}

func (c *Config) GetTimeout() time.Duration {
	return time.Duration(c.Settings.DefaultTimeout) * time.Millisecond
}

// func (c *Config) GetPageLoadTimeout() time.Duration {
// 	return time.Duration(c.Settings.PageLoadTimeout) * time.Millisecond
// }

func (c *Config) ValidateConfiguration() error {
	if err := c.validateGlobalSettings(); err != nil {
		return err
	}

	if err := c.validateFrontend(); err != nil {
		return err
	}

	return nil
}

func (c *Config) validateFrontend() error {
	if len(c.Frontend.Elements) == 0 {
		return errors.New("frontend elements configuration is required")
	}

	if len(c.Frontend.Pages) == 0 {
		return errors.New("frontend pages configuration is required")
	}

	if c.Settings.DefaultTimeout < 1000 || c.Settings.DefaultTimeout > 300000 {
		return errors.New("default_timeout must be between 1000 and 300000 milliseconds")
	}

	// if c.Settings.PageLoadTimeout < 5000 || c.Settings.PageLoadTimeout > 300000 {
	// 	return errors.New("page_load_timeout must be between 5000 and 300000 milliseconds")
	// }
	return nil
}

func (c *Config) validateGlobalSettings() error {
	if c.ActiveEnvironment == "" {
		return errors.New("active_environment is required")
	}

	if _, exists := c.Environments[c.ActiveEnvironment]; !exists {
		return fmt.Errorf("active environment '%s' not found in environments", c.ActiveEnvironment)
	}

	if len(c.Environments) == 0 {
		return errors.New("at least one environment must be defined")
	}

	if c.Settings.Concurrency < 1 || c.Settings.Concurrency > 20 {
		return errors.New("concurrency must be between 1 and 20")
	}

	if c.Settings.GherkinLocation == "" {
		return errors.New("gherkin_location is required")
	}

	validReportFormats := []string{"html", "json", "junit"}
	reportFormatValid := false
	for _, validFormat := range validReportFormats {
		if c.Settings.ReportFormat == validFormat {
			reportFormatValid = true
			break
		}
	}
	if !reportFormatValid {
		return fmt.Errorf("report_format must be one of: %s", strings.Join(validReportFormats, ", "))
	}
	return nil
}
