package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
)

// Config represents the complete enhanced configuration structure.
// It provides multi-environment support, organized frontend selectors,
// comprehensive backend API definitions, and centralized settings management.
//
// This configuration follows the principle of making dependencies explicit
// and supports environment-specific overrides through TEST_ENV environment variable.
type Config struct {
	// ActiveEnvironment specifies which environment configuration to use.
	// Can be overridden by TEST_ENV environment variable.
	ActiveEnvironment string `yaml:"active_environment" validate:"required"`

	// Settings contains global framework configuration.
	Settings GlobalSettings `yaml:"settings" validate:"required"`

	// Environments maps environment names to their specific configurations.
	Environments map[string]Environment `yaml:"environments" validate:"required,min=1"`

	// Frontend contains UI element selectors and page definitions.
	Frontend FrontendConfig `yaml:"frontend" validate:"required"`

	// Backend contains API endpoint definitions and default headers.
	Backend BackendConfig `yaml:"backend" validate:"required"`
}

// GlobalSettings defines framework-wide configuration parameters.
// These settings apply regardless of the active environment and control
// TestFlowKit's behavior at the framework level.
type GlobalSettings struct {
	// DefaultTimeout specifies the default timeout in milliseconds for element waits.
	DefaultTimeout int `yaml:"default_timeout" validate:"min=1000,max=300000"`

	// PageLoadTimeout specifies the maximum time to wait for page loads in milliseconds.
	PageLoadTimeout int `yaml:"page_load_timeout" validate:"min=5000,max=300000"`

	// ScreenshotOnFailure enables automatic screenshots when tests fail.
	ScreenshotOnFailure bool `yaml:"screenshot_on_failure"`

	// VideoRecording enables video recording of test execution.
	VideoRecording bool `yaml:"video_recording"`

	// Concurrency sets the number of parallel test executions.
	Concurrency int `yaml:"concurrency" validate:"min=1,max=20"`

	// Headless controls whether to run browser in headless mode.
	Headless bool `yaml:"headless"`

	// SlowMotion adds delay between actions for debugging (e.g., "100ms").
	SlowMotion string `yaml:"slow_motion" validate:"omitempty"`

	// ReportFormat specifies the output format for test reports.
	ReportFormat string `yaml:"report_format" validate:"oneof=html json junit"`

	// GherkinLocation specifies the directory containing feature files.
	GherkinLocation string `yaml:"gherkin_location" validate:"required"`

	// Tags specifies which tagged scenarios to run.
	Tags string `yaml:"tags"`
}

// Environment represents environment-specific configuration.
// Each environment (local, staging, production) can have different URLs
// and service endpoints while maintaining the same test structure.
type Environment struct {
	// FrontendBaseURL is the base URL for the web application frontend.
	FrontendBaseURL string `yaml:"frontend_base_url" validate:"required,url"`

	// BackendBaseURL is the base URL for API endpoints.
	BackendBaseURL string `yaml:"backend_base_url" validate:"required,url"`
}

// FrontendConfig organizes UI element selectors and page routes.
// This structure supports fallback selector strategies and page-aware element resolution.
type FrontendConfig struct {
	// Elements maps page/component names to element selectors with fallback strategies.
	// Structure: page_name -> element_name -> []selector
	// Example: login_page -> email_field -> ["[data-testid='email']", "#email", "input[type='email']"]
	Elements map[string]map[string][]string `yaml:"elements" validate:"required"`

	// Pages maps logical page names to their routes or full URLs.
	// Example: "login" -> "/login" or "google" -> "https://www.google.com"
	Pages map[string]string `yaml:"pages" validate:"required"`
}

// BackendConfig defines API testing configuration including endpoints and default headers.
// This supports comprehensive API testing capabilities with detailed endpoint specifications.
type BackendConfig struct {
	// DefaultHeaders are applied to all API requests unless overridden.
	// Supports environment variable substitution (${VAR_NAME} format).
	DefaultHeaders map[string]string `yaml:"default_headers"`

	// Endpoints maps endpoint names to their detailed configurations.
	Endpoints map[string]Endpoint `yaml:"endpoints" validate:"required"`
}

// Endpoint represents a complete API endpoint definition.
// This provides self-documenting API configurations with method, path, and description.
type Endpoint struct {
	// Method specifies the HTTP method (GET, POST, PUT, DELETE, PATCH, etc.).
	Method string `yaml:"method" validate:"required,oneof=GET POST PUT DELETE PATCH HEAD OPTIONS"`

	// Path specifies the endpoint path, supporting path parameters with {param} syntax.
	// Example: "/users/{id}" for parameterized paths.
	Path string `yaml:"path" validate:"required"`

	// Description provides human-readable documentation for the endpoint.
	Description string `yaml:"description" validate:"required"`
}

// GetCurrentEnvironment returns the configuration for the currently active environment.
// This is a convenience method that handles environment resolution and validation.
func (c *Config) GetCurrentEnvironment() (Environment, error) {
	env, exists := c.Environments[c.ActiveEnvironment]
	if !exists {
		return Environment{}, fmt.Errorf("active environment '%s' not found in configuration", c.ActiveEnvironment)
	}
	return env, nil
}

// GetFrontendURL constructs the complete URL for a given page.
// It handles both relative paths (prefixed with environment base URL) and absolute URLs.
//
// Parameters:
//   - page: The logical page name as defined in frontend.pages
//
// Returns the complete URL or the environment's frontend base URL if page is not found.
func (c *Config) GetFrontendURL(page string) (string, error) {
	env, err := c.GetCurrentEnvironment()
	if err != nil {
		return "", err
	}

	if pagePath, ok := c.Frontend.Pages[page]; ok {
		// Handle absolute URLs (external sites)
		if strings.HasPrefix(pagePath, "http://") || strings.HasPrefix(pagePath, "https://") {
			return pagePath, nil
		}
		// Handle relative paths
		return env.FrontendBaseURL + pagePath, nil
	}

	// Default to base URL if page not found
	return env.FrontendBaseURL, nil
}

// GetAPIEndpoint constructs the complete API URL and returns endpoint configuration.
// This method handles environment-specific base URLs and path parameter placeholders.
//
// Parameters:
//   - endpointName: The logical endpoint name as defined in backend.endpoints
//
// Returns:
//   - Complete URL with environment base URL
//   - Endpoint configuration with method, path, and description
//   - Error if endpoint not found
func (c *Config) GetAPIEndpoint(endpointName string) (string, Endpoint, error) {
	env, err := c.GetCurrentEnvironment()
	if err != nil {
		return "", Endpoint{}, err
	}

	endpoint, ok := c.Backend.Endpoints[endpointName]
	if !ok {
		return "", Endpoint{}, fmt.Errorf("endpoint '%s' not found in configuration", endpointName)
	}

	fullURL := env.BackendBaseURL + endpoint.Path
	return fullURL, endpoint, nil
}

// GetElementSelectors returns all available selectors for a specific element.
// This implements the fallback selector strategy: page-specific selectors first,
// then common selectors as fallbacks.
//
// Parameters:
//   - page: The page/component name where the element resides
//   - element: The element name to find selectors for
//
// Returns slice of selectors in priority order (most specific first).
func (c *Config) GetElementSelectors(page, element string) []string {
	// Try page-specific selectors first
	if pageElements, ok := c.Frontend.Elements[page]; ok {
		if selectors, ok := pageElements[element]; ok {
			return selectors
		}
	}

	// Fallback to common selectors
	if commonElements, ok := c.Frontend.Elements["common"]; ok {
		if selectors, ok := commonElements[element]; ok {
			return selectors
		}
	}

	return []string{} // No selectors found
}

// GetSlowMotion parses and returns the slow motion duration.
// Returns 0 duration if headless mode is enabled or if parsing fails.
func (c *Config) GetSlowMotion() time.Duration {
	if c.Settings.Headless {
		return 0
	}

	if c.Settings.SlowMotion == "" {
		return 0
	}

	duration, err := time.ParseDuration(c.Settings.SlowMotion)
	if err != nil {
		log.Printf("Invalid slow motion duration: %s, using 0", c.Settings.SlowMotion)
		return 0
	}
	return duration
}

// IsHeadlessModeEnabled returns whether the browser should run in headless mode.
func (c *Config) IsHeadlessModeEnabled() bool {
	return c.Settings.Headless
}

// GetConcurrency returns the number of parallel test executions allowed.
// Ensures a minimum of 1 concurrent execution.
func (c *Config) GetConcurrency() int {
	if c.Settings.Concurrency <= 0 {
		return 1
	}
	return c.Settings.Concurrency
}

// GetTimeout returns the default timeout duration for element operations.
func (c *Config) GetTimeout() time.Duration {
	return time.Duration(c.Settings.DefaultTimeout) * time.Millisecond
}

// GetPageLoadTimeout returns the timeout duration for page load operations.
func (c *Config) GetPageLoadTimeout() time.Duration {
	return time.Duration(c.Settings.PageLoadTimeout) * time.Millisecond
}

// ValidateConfiguration performs comprehensive validation of the configuration.
// This ensures all required fields are present and values are within acceptable ranges.
func (c *Config) ValidateConfiguration() error {
	// Validate active environment exists
	if c.ActiveEnvironment == "" {
		return fmt.Errorf("active_environment is required")
	}

	if _, exists := c.Environments[c.ActiveEnvironment]; !exists {
		return fmt.Errorf("active environment '%s' not found in environments", c.ActiveEnvironment)
	}

	// Validate at least one environment is defined
	if len(c.Environments) == 0 {
		return fmt.Errorf("at least one environment must be defined")
	}

	// Validate frontend configuration
	if len(c.Frontend.Elements) == 0 {
		return fmt.Errorf("frontend elements configuration is required")
	}

	if len(c.Frontend.Pages) == 0 {
		return fmt.Errorf("frontend pages configuration is required")
	}

	// Validate backend configuration
	if len(c.Backend.Endpoints) == 0 {
		return fmt.Errorf("backend endpoints configuration is required")
	}

	// Validate individual endpoints
	for name, endpoint := range c.Backend.Endpoints {
		if endpoint.Method == "" {
			return fmt.Errorf("endpoint '%s': method is required", name)
		}
		if endpoint.Path == "" {
			return fmt.Errorf("endpoint '%s': path is required", name)
		}
		if endpoint.Description == "" {
			return fmt.Errorf("endpoint '%s': description is required", name)
		}
		
		// Validate HTTP method
		validMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
		methodValid := false
		for _, validMethod := range validMethods {
			if strings.ToUpper(endpoint.Method) == validMethod {
				methodValid = true
				break
			}
		}
		if !methodValid {
			return fmt.Errorf("endpoint '%s': invalid HTTP method '%s'", name, endpoint.Method)
		}
	}

	// Validate settings
	if c.Settings.DefaultTimeout < 1000 || c.Settings.DefaultTimeout > 300000 {
		return fmt.Errorf("default_timeout must be between 1000 and 300000 milliseconds")
	}

	if c.Settings.PageLoadTimeout < 5000 || c.Settings.PageLoadTimeout > 300000 {
		return fmt.Errorf("page_load_timeout must be between 5000 and 300000 milliseconds")
	}

	if c.Settings.Concurrency < 1 || c.Settings.Concurrency > 20 {
		return fmt.Errorf("concurrency must be between 1 and 20")
	}

	if c.Settings.GherkinLocation == "" {
		return fmt.Errorf("gherkin_location is required")
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