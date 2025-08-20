package config

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"testflowkit/pkg/logger"

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

		data, err := os.ReadFile(configFilePath)
		if err != nil {
			errCfg = fmt.Errorf("failed to read config file '%s': %w", configFilePath, err)
			return
		}

		var config Config
		if yamlErr := yaml.Unmarshal(data, &config); yamlErr != nil {
			errCfg = fmt.Errorf("failed to parse YAML configuration: %w", err)
			return
		}

		logger.Info("Configuration parsed successfully")

		applyOverrides(&config, overrides)

		if validateErr := config.ValidateConfiguration(); validateErr != nil {
			errCfg = fmt.Errorf("invalid configuration: %w", validateErr)
			return
		}

		logger.InfoFf("Configuration loaded successfully for environment: %s", config.ActiveEnvironment)
		cfg = &config
	})

	return errCfg
}

func applyOverrides(config *Config, overrides Overrides) {
	if overrides.ActiveEnvironment != "" {
		config.ActiveEnvironment = overrides.ActiveEnvironment
	}

	if overrides.Settings.GherkinLocation != "" {
		config.Settings.GherkinLocation = overrides.Settings.GherkinLocation
	}

	if overrides.Settings.Tags != "" {
		config.Settings.Tags = overrides.Settings.Tags
	}

	if overrides.Frontend.DefaultTimeout > 0 {
		config.Frontend.DefaultTimeout = overrides.Frontend.DefaultTimeout
	}

	config.Frontend.Headless = overrides.Frontend.Headless
}

func Get() (*Config, error) {
	if cfg == nil {
		return nil, errors.New("configuration not loaded - call Load first")
	}
	return cfg, nil
}

type Overrides struct {
	ActiveEnvironment string
	Settings          GlobalSettings
	Frontend          FrontendConfig
}
