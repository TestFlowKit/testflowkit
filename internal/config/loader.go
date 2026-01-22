package config

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/variables"

	"github.com/goccy/go-yaml"
)

var (
	cfg     *Config
	cfgOnce sync.Once
	errCfg  error
)

func Load(configFilePath string, overrides Overrides) error {
	cfgOnce.Do(func() {
		logger.InfoFf("Loading configuration from: %s", configFilePath)

		// 1. Read file raw content
		data, err := os.ReadFile(configFilePath)
		if err != nil {
			errCfg = fmt.Errorf("failed to read config file '%s': %w", configFilePath, err)
			return
		}

		// 2. Pre-parse to get Env settings
		var preCfg struct {
			Settings GlobalSettings `yaml:"settings"`
			Env      map[string]any `yaml:"env"`
		}
		if yamlErr := yaml.Unmarshal(data, &preCfg); yamlErr != nil {
			errCfg = fmt.Errorf("failed to parse YAML configuration (pre-load): %w", yamlErr)
			return
		}

		// 3. Determine EnvFile (CLI overrides Config)
		envFile := preCfg.Settings.EnvFile
		if overrides.Settings.EnvFile != "" {
			envFile = overrides.Settings.EnvFile
		}

		// 4. Load & Merge Environment Variables
		envVars := FlattenMap(preCfg.Env, "")

		if envFile != "" {
			logger.InfoFf("Loading environment variables from: %s", envFile)
			fileVars, err := LoadEnvFile(envFile)
			if err != nil {
				errCfg = fmt.Errorf("failed to load env file '%s': %w", envFile, err)
				return
			}
			// File vars override inline vars
			for k, v := range fileVars {
				envVars[k] = v
			}
			logger.InfoFf("Loaded %d environment variables from file", len(fileVars))
		}

		variables.SetEnvVariables(envVars)
		logger.InfoFf("Environment variables initialized: %d total", len(envVars))

		// 5. Substitute variables in config content
		dataStr := variables.ReplaceEnvVariables(string(data))

		// 6. Unmarshal final config
		var config Config
		if yamlErr := yaml.Unmarshal([]byte(dataStr), &config); yamlErr != nil {
			errCfg = fmt.Errorf("failed to parse YAML configuration: %w", yamlErr)
			return
		}

		applyOverrides(&config, overrides)

		if validateErr := config.ValidateConfiguration(); validateErr != nil {
			errCfg = fmt.Errorf("invalid configuration: %w", validateErr)
			return
		}

		logger.Info("Configuration loaded successfully")
		cfg = &config
	})

	return errCfg
}

func Get() (*Config, error) {
	if cfg == nil {
		return nil, errors.New("configuration not loaded - call Load first")
	}
	return cfg, nil
}

func applyOverrides(config *Config, overrides Overrides) {
	if overrides.Settings.GherkinLocation != "" {
		config.Settings.GherkinLocation = overrides.Settings.GherkinLocation
	}

	if overrides.Settings.Tags != "" {
		config.Settings.Tags = overrides.Settings.Tags
	}

	if overrides.Settings.EnvFile != "" {
		config.Settings.EnvFile = overrides.Settings.EnvFile
	}

	if config.IsFrontendDefined() {
		if overrides.Frontend.DefaultTimeout > 0 {
			config.Frontend.DefaultTimeout = overrides.Frontend.DefaultTimeout
		}
		config.Frontend.Headless = overrides.Frontend.Headless
	}
}

type Overrides struct {
	Settings GlobalSettings
	Frontend FrontendConfig
}
