package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"testflowkit/pkg/logger"
)

type Config struct {
	Env map[string]any `yaml:"env"`

	Settings GlobalSettings `yaml:"settings" validate:"required"`

	Frontend *FrontendConfig `yaml:"frontend"`

	APIs *APIsConfig `yaml:"apis"`

	Files FileConfig `yaml:"files"`

	configPath string
	appVersion string
}

func (c *Config) GetAPI(apiName string) (*APIDefinition, error) {
	if c.APIs == nil || c.APIs.Definitions == nil {
		return nil, errors.New("no APIs configured")
	}

	apiDef, exists := c.APIs.Definitions[apiName]
	if !exists {
		return nil, fmt.Errorf("API '%s' not found in configuration", apiName)
	}

	return &apiDef, nil
}

func (c *Config) GetAPITimeout(apiName string) int {
	const defaultTimeoutInMS = 30000
	if c.APIs == nil {
		return defaultTimeoutInMS
	}

	apiDef, err := c.GetAPI(apiName)
	if err != nil {
		if c.APIs.DefaultTimeout > 0 {
			return c.APIs.DefaultTimeout
		}
		return defaultTimeoutInMS
	}

	if apiDef.Timeout != nil {
		return *apiDef.Timeout
	}

	if c.APIs.DefaultTimeout > 0 {
		return c.APIs.DefaultTimeout
	}

	return defaultTimeoutInMS
}

func (c *Config) IsAPIsConfigured() bool {
	return c.APIs != nil && len(c.APIs.Definitions) > 0
}

func (c *Config) GetFileDefinitions() map[string]string {
	return c.Files.Definitions
}

func (c *Config) GetConcurrency() int {
	if c.Settings.Concurrency <= 0 {
		return 1
	}
	return c.Settings.Concurrency
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

	if err := c.validateAPIs(); err != nil {
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

func (c *Config) SetVersion(ver string) {
	c.appVersion = ver
}

func (c *Config) SetConfigPath(path string) {
	c.configPath = path
}

func (c *Config) GetConfigPath() string {
	return c.configPath
}

func (c *Config) GetVersion() string {
	if c.appVersion == "" {
		return "unknown"
	}
	return c.appVersion
}

func (c *Config) validateAPIs() error {
	if c.APIs == nil || len(c.APIs.Definitions) == 0 {
		logger.Info("APIs config is not defined")
		return nil
	}

	for apiName, apiDef := range c.APIs.Definitions {
		if apiDef.Type != APITypeREST && apiDef.Type != APITypeGraphQL {
			return fmt.Errorf("API '%s' has invalid type '%s', must be 'rest' or 'graphql'", apiName, apiDef.Type)
		}

		if apiDef.Type == APITypeREST {
			err := c.validateRestAPI(apiDef, apiName)
			if err != nil {
				return err
			}
		}

		if apiDef.Type == APITypeGraphQL {
			err := c.validateGraphQLApi(apiDef, apiName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (*Config) validateGraphQLApi(apiDef APIDefinition, apiName string) error {
	if apiDef.Endpoint == "" {
		return fmt.Errorf("GraphQL API '%s' must have an endpoint", apiName)
	}
	if len(apiDef.Operations) == 0 {
		return fmt.Errorf("GraphQL API '%s' must have at least one operation", apiName)
	}
	for operationName, operation := range apiDef.Operations {
		if operation.Type != "query" && operation.Type != "mutation" {
			return fmt.Errorf("GraphQL operation '%s.%s' type must be 'query' or 'mutation'", apiName, operationName)
		}
		if operation.Operation == "" {
			return fmt.Errorf("GraphQL operation '%s.%s' must have an operation definition", apiName, operationName)
		}
		if operation.Description == "" {
			logger.Warn(fmt.Sprintf("GraphQL operation '%s.%s' should have a description", apiName, operationName), nil)
		}
	}
	return nil
}

func (*Config) validateRestAPI(apiDef APIDefinition, apiName string) error {
	if apiDef.BaseURL == "" {
		return fmt.Errorf("REST API '%s' must have a base_url", apiName)
	}
	if len(apiDef.Endpoints) == 0 {
		return fmt.Errorf("REST API '%s' must have at least one endpoint", apiName)
	}
	for endpointName, endpoint := range apiDef.Endpoints {
		if endpoint.Method == "" {
			return fmt.Errorf("endpoint '%s.%s' must have a method", apiName, endpointName)
		}
		if endpoint.Path == "" {
			return fmt.Errorf("endpoint '%s.%s' must have a path", apiName, endpointName)
		}
	}
	return nil
}
