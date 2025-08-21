package config

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"testflowkit/pkg/logger"
	"time"
)

type Config struct {
	ActiveEnvironment string `yaml:"active_environment" validate:"required"`

	Settings GlobalSettings `yaml:"settings" validate:"required"`

	Environments map[string]Environment `yaml:"environments" validate:"required,min=1"`

	Frontend *FrontendConfig `yaml:"frontend"`

	Backend BackendConfig `yaml:"backend"`

	Files FileConfig `yaml:"files"`
}

func (c *Config) GetCurrentEnvironment() (Environment, error) {
	env, exists := c.Environments[c.ActiveEnvironment]
	if !exists {
		return Environment{}, fmt.Errorf("active environment '%s' not found in configuration", c.ActiveEnvironment)
	}
	return env, nil
}

func (c *Config) GetAPIEndpoint(endpointName string) (string, Endpoint, error) {
	endpoint, exists := c.Backend.Endpoints[endpointName]
	if !exists {
		return "", Endpoint{}, fmt.Errorf("endpoint '%s' not found in configuration", endpointName)
	}

	parsedURL, err := url.Parse(endpoint.Path)
	if err != nil {
		return "", Endpoint{}, fmt.Errorf("failed to parse endpoint path: %w", err)
	}

	if parsedURL.Scheme != "" {
		return parsedURL.String(), endpoint, nil
	}

	fullURL := filepath.Join(c.GetBackendBaseURL(), parsedURL.Path)
	return fullURL, endpoint, nil
}

func (c *Config) GetFileDefinitions() map[string]string {
	return c.Files.Definitions
}

func (c *Config) GetThinkTime() time.Duration {
	if c.Settings.ThinkTime == 0 {
		return 0
	}

	duration, err := time.ParseDuration(fmt.Sprintf("%dms", c.Settings.ThinkTime))
	if err != nil {
		log.Printf("Invalid think time duration: %d, using 0", c.Settings.ThinkTime)
		return 0
	}
	return duration
}

func (c *Config) GetConcurrency() int {
	if c.Settings.Concurrency <= 0 {
		return 1
	}
	return c.Settings.Concurrency
}

func (c *Config) GetBackendBaseURL() string {
	env, err := c.GetCurrentEnvironment()
	if err != nil {
		return ""
	}
	return env.APIBaseURL
}

func (c *Config) GetFileBaseDirectory() string {
	return c.Files.BaseDirectory
}

func (c *Config) GetFilesPaths(fileNames []string) ([]string, error) {
	defs := c.GetFileDefinitions()
	if defs == nil {
		return []string{}, errors.New("no file definitions configured")
	}

	filePaths := []string{}
	notFoundFiles := []string{}
	for _, fileName := range fileNames {
		filePath, exists := defs[fileName]
		if !exists {
			notFoundFiles = append(notFoundFiles, fileName)
			continue
		}

		fullPath := filepath.Join(c.GetFileBaseDirectory(), filePath)

		filePaths = append(filePaths, fullPath)
	}

	if len(notFoundFiles) > 0 {
		return nil, fmt.Errorf("files do not exist: %v", notFoundFiles)
	}

	return filePaths, nil
}

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
	if !c.IsFrontendDefined() {
		logger.Info("frontend config is not defined")
		return nil
	}
	if len(c.Frontend.Elements) == 0 {
		return errors.New("frontend elements configuration is required")
	}

	if len(c.Frontend.Pages) == 0 {
		return errors.New("frontend pages configuration is required")
	}

	if c.Frontend.DefaultTimeout < 1000 || c.Frontend.DefaultTimeout > 300000 {
		return errors.New("default_timeout (element search timeout) must be between 1000 and 300000 milliseconds")
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
