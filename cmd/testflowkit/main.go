package main

import (
	"fmt"
	"testflowkit/internal/actions"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

// Version is injected at build time via ldflags.
var Version string

func main() {
	args := getAppArgs()

	mode, err := args.getMode()
	if err != nil {
		logger.Fatal("Failed to get mode", err)
	}

	cfg, err := getConfig(args, err)

	actions.Execute(cfg, err, mode)
}

func getConfig(args argsConfig, err error) (*config.Config, error) {
	cfgPath, getcfigPathErr := args.getConfigPath()
	defaultConf := &config.Config{}
	defaultConf.SetVersion(Version)
	if getcfigPathErr != nil {
		return defaultConf, fmt.Errorf("failed to get config path: %w", getcfigPathErr)
	}

	configLoadErr := config.Load(cfgPath, args.getAppConfigOverrides())
	if configLoadErr != nil {
		return defaultConf, fmt.Errorf("failed to load config: %w", configLoadErr)
	}

	cfg, configGetErr := config.Get()
	if configGetErr != nil {
		return defaultConf, fmt.Errorf("failed to get config: %w", err)
	}

	cfg.SetVersion(Version)

	return cfg, nil
}
