package main

import (
	"flag"
	"fmt"
	"os"

	"testflowkit/internal/actions"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

// Version information - set during build
var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

func main() {
	// Define command line flags following Go best practices
	var (
		useEnhanced       = flag.Bool("enhanced", false, "Use enhanced configuration system")
		configPath        = flag.String("config", "", "Path to configuration file")
		showVersion       = flag.Bool("version", false, "Show version information")
		showConfigSummary = flag.Bool("show-config", false, "Display configuration summary")
		verbose           = flag.Bool("verbose", false, "Enable verbose logging")
	)

	// Check if enhanced configuration is requested via command line
	// We need to parse flags early to check for enhanced mode
	for _, arg := range os.Args[1:] {
		if arg == "-enhanced" || arg == "--enhanced" {
			*useEnhanced = true
			break
		}
	}

	// Show version information if requested
	if *showVersion {
		fmt.Printf("TestFlowKit %s\n", Version)
		fmt.Printf("Git Commit: %s\n", GitCommit)
		fmt.Printf("Build Date: %s\n", BuildDate)
		os.Exit(0)
	}

	// Configure logging based on verbose flag
	if *verbose {
		logger.Info("Verbose logging enabled")
	}

	// Branch execution based on configuration system
	if *useEnhanced {
		handleEnhancedMode(*configPath, *showConfigSummary)
	} else {
		handleLegacyMode()
	}
}

// handleEnhancedMode manages execution with the enhanced configuration system
func handleEnhancedMode(configPath string, showConfigSummary bool) {
	logger.Info("Using enhanced configuration system")

	// Determine configuration file path
	enhancedConfigPath := configPath
	if enhancedConfigPath == "" {
		enhancedConfigPath = "config-enhanced.yml"
		logger.InfoFf("No config path specified, using default: %s", enhancedConfigPath)
	}

	// Load and validate enhanced configuration
	_, err := config.LoadEnhancedConfig(enhancedConfigPath)
	if err != nil {
		logger.FatalFf("Failed to load enhanced configuration from '%s': %v", enhancedConfigPath, err)
	}

	// Show configuration summary if requested
	if showConfigSummary {
		fmt.Println(config.GetConfigurationSummary())
		os.Exit(0)
	}

	// Get the loaded enhanced configuration
	cfg, err := config.GetEnhancedConfig()
	if err != nil {
		logger.FatalFf("Enhanced configuration not available: %v", err)
	}

	// Display enhanced configuration summary
	displayEnhancedConfigSummary(cfg)

	// Convert enhanced config to legacy format for backward compatibility
	legacyConfig := convertToLegacyConfig(cfg)

	// Execute the run action (enhanced mode currently focuses on test execution)
	logger.Info("Running tests with enhanced configuration")
	actions.Run(legacyConfig)
}

// handleLegacyMode manages execution with the original configuration system
func handleLegacyMode() {
	logger.Info("Using legacy configuration system")
	
	// Use the existing configuration initialization
	cliConfig := config.Init()
	if cliConfig == nil {
		logger.Fatal("no mode specified", nil)
	}

	// Execute based on configured mode
	modes := map[config.Mode]actions.Type{
		config.RunMode:        runMode,
		config.InitMode:       actions.Init,
		config.ValidationMode: actions.Validate,
	}

	if mode, ok := modes[cliConfig.Mode]; ok {
		mode(cliConfig)
	} else {
		logger.FatalFf("unknown mode: %s", cliConfig.Mode)
	}
}

// runMode handles the legacy run mode execution
func runMode(appConfig *config.App) {
	displayLegacyConfigSummary(appConfig)
	actions.Run(appConfig)
}

// displayEnhancedConfigSummary shows the enhanced configuration summary
func displayEnhancedConfigSummary(cfg *config.Config) {
	logger.Info("--- Enhanced Configuration Summary ---")

	env, _ := cfg.GetCurrentEnvironment()
	
	logger.InfoFf("Active Environment: %s", cfg.ActiveEnvironment)
	logger.InfoFf("Frontend Base URL: %s", env.FrontendBaseURL)
	logger.InfoFf("Backend Base URL: %s", env.BackendBaseURL)
	logger.InfoFf("Headless Mode: %t", cfg.Settings.Headless)
	logger.InfoFf("Concurrency: %d", cfg.Settings.Concurrency)
	logger.InfoFf("Report Format: %s", cfg.Settings.ReportFormat)
	logger.InfoFf("Gherkin Location: %s", cfg.Settings.GherkinLocation)
	logger.InfoFf("Test Tags: %s", cfg.Settings.Tags)
	logger.InfoFf("Default Timeout: %dms", cfg.Settings.DefaultTimeout)
	logger.InfoFf("Page Load Timeout: %dms", cfg.Settings.PageLoadTimeout)
	logger.InfoFf("Screenshot on Failure: %t", cfg.Settings.ScreenshotOnFailure)
	logger.InfoFf("Video Recording: %t", cfg.Settings.VideoRecording)
	logger.InfoFf("Slow Motion: %s", cfg.Settings.SlowMotion)
	logger.InfoFf("Elements Configured: %d page groups", len(cfg.Frontend.Elements))
	logger.InfoFf("API Endpoints: %d endpoints", len(cfg.Backend.Endpoints))

	logger.Info("--- Enhanced Configuration Summary End ---\n")
}

// displayLegacyConfigSummary shows the legacy configuration summary
func displayLegacyConfigSummary(appConfig *config.App) {
	logger.Info("--- Legacy Configuration Summary ---")

	logger.InfoFf("App Name: %s", appConfig.AppName)
	logger.InfoFf("App Description: %s", appConfig.AppDescription)
	logger.InfoFf("App Version: %s", appConfig.AppVersion)
	logger.InfoFf("App Tags: %s", appConfig.Tags)
	logger.InfoFf("Gherkin Location: %s", appConfig.GherkinLocation)
	logger.InfoFf("Report Format: %s", appConfig.ReportFormat)
	logger.InfoFf("Concurrency: %d", appConfig.GetConcurrency())
	logger.InfoFf("Slow Motion: %s", appConfig.GetSlowMotion())
	logger.InfoFf("Test Suite Timeout: %s", appConfig.Timeout)
	logger.InfoFf("Headless Mode: %t", appConfig.IsHeadlessModeEnabled())

	logger.Info("--- Legacy Configuration Summary End ---\n")
}

// convertToLegacyConfig converts enhanced configuration to legacy format.
// This bridge function enables gradual migration while maintaining backward compatibility.
//
// The conversion process:
// 1. Maps enhanced settings to legacy config fields
// 2. Resolves current environment URLs
// 3. Preserves essential configuration for existing actions
func convertToLegacyConfig(enhanced *config.Config) *config.App {
	// Get current environment configuration
	env, err := enhanced.GetCurrentEnvironment()
	if err != nil {
		logger.WarnFf("Failed to get current environment, using defaults: %v", err)
		env = config.Environment{
			FrontendBaseURL: "http://localhost:3000",
			BackendBaseURL:  "http://localhost:8080/api",
		}
	}

	// Create legacy configuration with enhanced settings
	legacy := &config.App{
		AppName:             "TestFlowKit",
		AppDescription:      "Web Test Automation Framework",
		AppVersion:          Version,
		Tags:                enhanced.Settings.Tags,
		GherkinLocation:     enhanced.Settings.GherkinLocation,
		ReportFormat:        enhanced.Settings.ReportFormat,
		Timeout:             fmt.Sprintf("%dms", enhanced.Settings.DefaultTimeout),
		SlowMotion:          enhanced.Settings.SlowMotion,
		Headless:            enhanced.Settings.Headless,
		ScreenshotOnFailure: enhanced.Settings.ScreenshotOnFailure,
		VideoRecording:      enhanced.Settings.VideoRecording,
		Concurrency:         enhanced.Settings.Concurrency,
		Browser:             "chrome",           // Default browser for legacy compatibility
		BaseURL:             env.FrontendBaseURL, // Use current environment frontend URL
	}

	// Log the conversion for debugging
	logger.InfoFf("Converted enhanced config to legacy format - Environment: %s, BaseURL: %s", 
		enhanced.ActiveEnvironment, legacy.BaseURL)

	return legacy
}
