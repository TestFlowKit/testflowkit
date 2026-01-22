package main

import (
	"errors"
	"testflowkit/internal/config"
	"time"

	"github.com/alexflint/go-arg"
)

type argsConfig struct {
	Version    bool         `arg:"--version,-v" help:"show version information"`
	Run        *runCmd      `arg:"subcommand:run" help:"run tests"`
	Init       *initCmd     `arg:"subcommand:init" help:"init cmd config"`
	Validate   *validateCmd `arg:"subcommand:validate" help:"validate gherkin files"`
	VersionCmd *versionCmd  `arg:"subcommand:version" help:"show version information"`
}

func (a *argsConfig) getConfigPath() (string, error) {
	if a.Run != nil {
		return a.Run.ConfigPath, nil
	}

	if a.Validate != nil {
		return a.Validate.ConfigPath, nil
	}

	return "", errors.New("no config path provided")
}

func (a *argsConfig) getAppConfigOverrides() config.Overrides {
	if a.Run != nil {
		return config.Overrides{
			Settings: config.GlobalSettings{
				GherkinLocation: a.Run.GherkinLocation,
				Tags:            a.Run.Tags,
				EnvFile:         a.Run.EnvFile,
			},
			Frontend: config.FrontendConfig{
				DefaultTimeout: a.GetTimeout(),
				Headless:       a.Run.Headless,
			},
		}
	}

	if a.Validate != nil {
		return config.Overrides{
			Settings: config.GlobalSettings{
				GherkinLocation: a.Validate.GherkinLocation,
				Tags:            a.Validate.Tags,
				EnvFile:         a.Validate.EnvFile,
			},
		}
	}

	return config.Overrides{}
}

func (a *argsConfig) getMode() (config.Mode, error) {
	if a.Version || a.VersionCmd != nil {
		return config.VersionMode, nil
	}

	if a.Run != nil {
		return config.RunMode, nil
	}

	if a.Init != nil {
		return config.InitMode, nil
	}

	if a.Validate != nil {
		return config.ValidationMode, nil
	}

	return "", errors.New("no mode provided")
}

func (a *argsConfig) GetTimeout() int {
	if a.Run != nil {
		timeout, err := time.ParseDuration(a.Run.Timeout)
		if err != nil {
			return 0
		}
		return int(timeout.Milliseconds())
	}

	return 0
}

type runCmd struct {
	commonCmd
	Headless bool   `arg:"--headless" help:"headless mode" default:"true"`
	Timeout  string `arg:"--timeout" help:"timeout duration (e.g. 10s, 1m, 2h)"`
}

type initCmd struct {
}

type validateCmd struct {
	commonCmd
}

type versionCmd struct {
	// No additional fields needed - simple command
}

type commonCmd struct {
	GherkinLocation string `arg:"-l,--location" help:"path to gherkin files"`
	ConfigPath      string `arg:"-c,--config" help:"app config path" default:"config.yml"`
	Tags            string `arg:"-t,--tags" help:"tags"`
	EnvFile         string `arg:"--env-file" help:"path to env file"`
}

func getAppArgs() argsConfig {
	c := argsConfig{}
	arg.MustParse(&c)
	return c
}
