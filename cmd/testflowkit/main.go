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
	if cfg != nil {
		cfg.SetVersion(Version)
	}

	actions.Execute(cfg, err, mode)
}

func getConfig(args argsConfig, err error) (*config.Config, error) {
	cfgPath, getcfigPathErr := args.getConfigPath()
	if getcfigPathErr != nil {
		logger.Fatal("Failed to get config path", getcfigPathErr)
	}

	configLoadErr := config.Load(cfgPath, args.getAppConfigOverrides())
	if configLoadErr != nil {
		return nil, fmt.Errorf("failed to load config: %w", configLoadErr)
	}

	cfg, configGetErr := config.Get()
	if configGetErr != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	return cfg, nil
}
