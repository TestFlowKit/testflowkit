package config

import (
	"errors"
	"fmt"
	"log"
	"maps"
	"net/url"
	"path/filepath"
	"strings"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/variables"
	"time"
)

type Config struct {
	Env map[string]any `yaml:"env"`

	Settings GlobalSettings `yaml:"settings" validate:"required"`

	Frontend *FrontendConfig `yaml:"frontend"`

	Backend BackendConfig `yaml:"backend"`

	Files FileConfig `yaml:"files"`

	configPath string
	appVersion string
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

	fullURL := filepath.Join(c.GetRestAPIBaseURL(), parsedURL.Path)
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

func (c *Config) GetRestAPIBaseURL() string {
	val, _ := variables.GetEnvVariable("rest_api_base_url")
	return val
}

func (c *Config) GetGraphQLEndpoint() (string, error) {
	if c.Backend.GraphQL == nil {
		return "", errors.New("GraphQL configuration not found")
	}

	if val, exists := variables.GetEnvVariable("graphql_endpoint"); exists && val != "" {
		return val, nil
	}

	return "", errors.New("GraphQL endpoint not defined in environment configuration")
}

func (c *Config) GetGraphQLOperation(operationName string) (GraphQLOperation, error) {
	if c.Backend.GraphQL == nil {
		return GraphQLOperation{}, errors.New("GraphQL configuration not found")
	}

	operation, exists := c.Backend.GraphQL.Operations[operationName]
	if !exists {
		return GraphQLOperation{}, fmt.Errorf("GraphQL operation '%s' not found in configuration", operationName)
	}

	return operation, nil
}

func (c *Config) GetGraphQLHeaders() map[string]string {
	headers := make(map[string]string)

	maps.Copy(headers, c.Backend.DefaultHeaders)

	if c.Backend.GraphQL != nil {
		maps.Copy(headers, c.Backend.GraphQL.DefaultHeaders)
	}

	return headers
}

func (c *Config) IsGraphQLConfigured() bool {
	return c.Backend.GraphQL != nil
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

	if err := c.validateGraphQL(); err != nil {
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

func (c *Config) validateGraphQL() error {
	if c.Backend.GraphQL == nil {
		logger.Info("GraphQL config is not defined")
		return nil
	}

	graphqlEndpoint, exists := variables.GetEnvVariable("graphql_endpoint")
	if !exists || graphqlEndpoint == "" {
		return errors.New("GraphQL endpoint must be defined (env.graphql_endpoint)")
	}
	if len(c.Backend.GraphQL.Operations) == 0 {
		return errors.New("at least one GraphQL operation must be defined when GraphQL config is present")
	}

	// Validate each operation
	for operationName, operation := range c.Backend.GraphQL.Operations {
		if operation.Type == "" {
			return fmt.Errorf("GraphQL operation '%s' must have a type", operationName)
		}

		if operation.Type != "query" && operation.Type != "mutation" {
			const msg = "GraphQL operation '%s' type must be 'query' or 'mutation', got '%s'"
			return fmt.Errorf(msg, operationName, operation.Type)
		}

		if operation.Operation == "" {
			const msg = "GraphQL operation '%s' must have an operation definition"
			return fmt.Errorf(msg, operationName)
		}

		if operation.Description == "" {
			const msg = "GraphQL operation '%s' should have a description"
			logger.Warn(fmt.Sprintf(msg, operationName), nil)
		}
	}

	return nil
}
