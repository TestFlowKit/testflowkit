package main

import (
	"errors"
	"testflowkit/internal/config"
	"time"

	"github.com/alexflint/go-arg"
)

type argsConfig struct {
	Run      *runCmd      `arg:"subcommand:run" help:"run tests"`
	Init     *initCmd     `arg:"subcommand:init" help:"init cmd config"`
	Validate *validateCmd `arg:"subcommand:validate" help:"validate gherkin files"`
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
			ActiveEnvironment: a.Run.Env,
			Settings: config.GlobalSettings{
				GherkinLocation: a.Run.GherkinLocation,
				Tags:            a.Run.Tags,
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
			},
		}
	}

	return config.Overrides{}
}

func (a *argsConfig) getMode() (config.Mode, error) {
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
	Env      string `arg:"--env" help:"environment"`
}

type initCmd struct {
}

type validateCmd struct {
	commonCmd
}

type commonCmd struct {
	GherkinLocation string `arg:"-l,--location" help:"path to gherkin files"`
	ConfigPath      string `arg:"-c,--config" help:"app config path" default:"config.yml"`
	Tags            string `arg:"-t,--tags" help:"tags"`
}

func getAppArgs() argsConfig {
	c := argsConfig{}
	arg.MustParse(&c)
	return c
}
