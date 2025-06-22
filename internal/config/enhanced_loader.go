package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/goccy/go-yaml"
	"testflowkit/pkg/logger"
)

var (
	// enhancedConfig holds the loaded configuration (singleton pattern)
	enhancedConfig     *Config
	enhancedConfigOnce sync.Once
	enhancedConfigErr  error
)

// LoadEnhancedConfig loads and parses the enhanced YAML configuration file.
// This function implements the singleton pattern and provides comprehensive
// configuration loading with validation and environment variable substitution.
//
// The loading process:
// 1. Reads the YAML configuration file
// 2. Parses it into the Config struct
// 3. Applies environment variable overrides (TEST_ENV)
// 4. Substitutes environment variables (${VAR_NAME} format)
// 5. Validates the complete configuration
// 6. Caches the result for subsequent calls
//
// Parameters:
//   - configPath: Path to the YAML configuration file
//
// Returns:
//   - *Config: The loaded and validated configuration
//   - error: Any error encountered during loading or validation
func LoadEnhancedConfig(configPath string) (*Config, error) {
	enhancedConfigOnce.Do(func() {
		logger.Info("Loading enhanced configuration from: " + configPath)
		
		// Read configuration file
		data, err := os.ReadFile(configPath)
		if err != nil {
			enhancedConfigErr = fmt.Errorf("failed to read config file '%s': %w", configPath, err)
			return
		}

		// Parse YAML into Config struct
		var config Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			enhancedConfigErr = fmt.Errorf("failed to parse YAML configuration: %w", err)
			return
		}

		logger.Info("Configuration parsed successfully")

		// Apply environment variable overrides
		if err := applyEnvironmentOverrides(&config); err != nil {
			enhancedConfigErr = fmt.Errorf("failed to apply environment overrides: %w", err)
			return
		}

		// Substitute environment variables in configuration values
		if err := substituteEnvironmentVariables(&config); err != nil {
			enhancedConfigErr = fmt.Errorf("failed to substitute environment variables: %w", err)
			return
		}

		// Validate the complete configuration
		if err := config.ValidateConfiguration(); err != nil {
			enhancedConfigErr = fmt.Errorf("invalid configuration: %w", err)
			return
		}

		logger.InfoFf("Enhanced configuration loaded successfully for environment: %s", config.ActiveEnvironment)
		enhancedConfig = &config
	})

	return enhancedConfig, enhancedConfigErr
}

// GetEnhancedConfig returns the loaded enhanced configuration.
// This implements the singleton pattern and requires LoadEnhancedConfig to be called first.
//
// Returns:
//   - *Config: The cached configuration instance
//   - error: Error if configuration hasn't been loaded yet
func GetEnhancedConfig() (*Config, error) {
	if enhancedConfig == nil {
		return nil, fmt.Errorf("enhanced configuration not loaded - call LoadEnhancedConfig first")
	}
	return enhancedConfig, nil
}

// MustGetEnhancedConfig returns the loaded enhanced configuration or panics.
// Use this only when you're certain the configuration has been loaded.
func MustGetEnhancedConfig() *Config {
	config, err := GetEnhancedConfig()
	if err != nil {
		panic(fmt.Sprintf("Enhanced configuration not available: %v", err))
	}
	return config
}

// ResetEnhancedConfig clears the cached configuration.
// This is primarily useful for testing scenarios where you need to reload configuration.
func ResetEnhancedConfig() {
	enhancedConfig = nil
	enhancedConfigOnce = sync.Once{}
	enhancedConfigErr = nil
}

// applyEnvironmentOverrides applies environment variable overrides to the configuration.
// This allows runtime modification of configuration without changing the config file.
//
// Supported environment variables:
//   - TEST_ENV: Overrides the active_environment setting
//   - TEST_HEADLESS: Overrides the headless browser setting
//   - TEST_CONCURRENCY: Overrides the concurrency setting
//   - TEST_TAGS: Overrides the test tags setting
func applyEnvironmentOverrides(config *Config) error {
	// Override active environment if TEST_ENV is set
	if testEnv := os.Getenv("TEST_ENV"); testEnv != "" {
		logger.InfoFf("Overriding active environment from '%s' to '%s' via TEST_ENV", 
			config.ActiveEnvironment, testEnv)
		config.ActiveEnvironment = testEnv
		
		// Validate that the overridden environment exists
		if _, exists := config.Environments[testEnv]; !exists {
			return fmt.Errorf("environment '%s' specified in TEST_ENV not found in configuration", testEnv)
		}
	}

	// Override headless mode if TEST_HEADLESS is set
	if testHeadless := os.Getenv("TEST_HEADLESS"); testHeadless != "" {
		switch strings.ToLower(testHeadless) {
		case "true", "1", "yes":
			config.Settings.Headless = true
			logger.Info("Overriding headless mode to true via TEST_HEADLESS")
		case "false", "0", "no":
			config.Settings.Headless = false
			logger.Info("Overriding headless mode to false via TEST_HEADLESS")
		default:
			return fmt.Errorf("invalid TEST_HEADLESS value '%s', must be true/false, 1/0, or yes/no", testHeadless)
		}
	}

	// Override concurrency if TEST_CONCURRENCY is set
	if testConcurrency := os.Getenv("TEST_CONCURRENCY"); testConcurrency != "" {
		var concurrency int
		if _, err := fmt.Sscanf(testConcurrency, "%d", &concurrency); err != nil {
			return fmt.Errorf("invalid TEST_CONCURRENCY value '%s', must be a number: %w", testConcurrency, err)
		}
		if concurrency < 1 || concurrency > 20 {
			return fmt.Errorf("TEST_CONCURRENCY value %d out of range, must be between 1 and 20", concurrency)
		}
		config.Settings.Concurrency = concurrency
		logger.InfoFf("Overriding concurrency to %d via TEST_CONCURRENCY", concurrency)
	}

	// Override tags if TEST_TAGS is set
	if testTags := os.Getenv("TEST_TAGS"); testTags != "" {
		config.Settings.Tags = testTags
		logger.InfoFf("Overriding test tags to '%s' via TEST_TAGS", testTags)
	}

	return nil
}

// substituteEnvironmentVariables replaces ${VAR_NAME} placeholders with environment variable values.
// This provides secure and flexible configuration management, especially for sensitive data like API tokens.
//
// The substitution pattern supports:
//   - ${VAR_NAME} - Required environment variable (fails if not set)
//   - ${VAR_NAME:-default} - Optional with default value (future enhancement)
//
// Currently implemented substitution locations:
//   - Backend default headers (commonly used for Authorization tokens)
//   - Environment URLs (for dynamic environment configuration)
func substituteEnvironmentVariables(config *Config) error {
	// Substitute in backend default headers
	for key, value := range config.Backend.DefaultHeaders {
		substituted, err := performEnvironmentSubstitution(value)
		if err != nil {
			return fmt.Errorf("failed to substitute environment variables in backend header '%s': %w", key, err)
		}
		config.Backend.DefaultHeaders[key] = substituted
	}

	// Substitute in environment URLs
	for envName, env := range config.Environments {
		// Substitute frontend base URL
		frontendURL, err := performEnvironmentSubstitution(env.FrontendBaseURL)
		if err != nil {
			return fmt.Errorf("failed to substitute environment variables in frontend URL for environment '%s': %w", envName, err)
		}
		env.FrontendBaseURL = frontendURL

		// Substitute backend base URL
		backendURL, err := performEnvironmentSubstitution(env.BackendBaseURL)
		if err != nil {
			return fmt.Errorf("failed to substitute environment variables in backend URL for environment '%s': %w", envName, err)
		}
		env.BackendBaseURL = backendURL

		// Update the environment in the map
		config.Environments[envName] = env
	}

	return nil
}

// performEnvironmentSubstitution performs the actual environment variable substitution.
// This function handles the ${VAR_NAME} syntax and provides detailed error messages
// for missing environment variables.
//
// Parameters:
//   - input: String that may contain environment variable placeholders
//
// Returns:
//   - string: Input string with variables substituted
//   - error: Error if required environment variables are missing
func performEnvironmentSubstitution(input string) (string, error) {
	// Regular expression to match ${VAR_NAME} patterns
	envVarRegex := regexp.MustCompile(`\$\{([^}]+)\}`)
	
	// Track any missing environment variables
	var missingVars []string
	
	// Perform substitution
	result := envVarRegex.ReplaceAllStringFunc(input, func(match string) string {
		// Extract variable name (remove ${ and })
		varName := match[2 : len(match)-1]
		
		// Get environment variable value
		if value := os.Getenv(varName); value != "" {
			return value
		}
		
		// Track missing variables for detailed error reporting
		missingVars = append(missingVars, varName)
		return match // Return original if not found
	})
	
	// Report missing environment variables
	if len(missingVars) > 0 {
		return "", fmt.Errorf("required environment variables not set: %s", strings.Join(missingVars, ", "))
	}
	
	return result, nil
}

// GetConfigurationSummary returns a human-readable summary of the current configuration.
// This is useful for logging and debugging configuration issues.
//
// Returns a formatted string with key configuration details.
func GetConfigurationSummary() string {
	config, err := GetEnhancedConfig()
	if err != nil {
		return fmt.Sprintf("Configuration not loaded: %v", err)
	}

	env, _ := config.GetCurrentEnvironment()
	
	return fmt.Sprintf(`Enhanced Configuration Summary:
  Active Environment: %s
  Frontend Base URL: %s
  Backend Base URL: %s
  Headless Mode: %t
  Concurrency: %d
  Report Format: %s
  Gherkin Location: %s
  Test Tags: %s
  Default Timeout: %dms
  Page Load Timeout: %dms
  Elements Configured: %d page groups
  API Endpoints: %d endpoints`,
		config.ActiveEnvironment,
		env.FrontendBaseURL,
		env.BackendBaseURL,
		config.Settings.Headless,
		config.Settings.Concurrency,
		config.Settings.ReportFormat,
		config.Settings.GherkinLocation,
		config.Settings.Tags,
		config.Settings.DefaultTimeout,
		config.Settings.PageLoadTimeout,
		len(config.Frontend.Elements),
		len(config.Backend.Endpoints))
} 