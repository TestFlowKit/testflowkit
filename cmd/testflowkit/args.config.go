package main

import (
	"testflowkit/internal/config"
	"testflowkit/pkg/apperrors"
	"time"

	"github.com/alexflint/go-arg"
)

type argsConfig struct {
	Version    bool         `arg:"--version,-v" help:"show version information"`
	Run        *runCmd      `arg:"subcommand:run" help:"run tests"`
	Init       *initCmd     `arg:"subcommand:init" help:"init cmd config"`
	Install    *installCmd  `arg:"subcommand:install" help:"install browser drivers"`
	Validate   *validateCmd `arg:"subcommand:validate" help:"validate gherkin files"`
	VersionCmd *versionCmd  `arg:"subcommand:version" help:"show version information"`

	ExportStepDefinitions *exportStepDefinitionsCmd `arg:"subcommand:export-step-definitions" help:"export step definitions"`

	ExportConfigSchema *exportConfigSchemaCmd `arg:"subcommand:export-config-schema" help:"export config schema"`
}

func (a *argsConfig) getConfigPath() (string, error) {
	var path string

	switch {
	case a.Run != nil:
		path = a.Run.ConfigPath
	case a.Validate != nil:
		path = a.Validate.ConfigPath
	case a.Install != nil:
		path = a.Install.ConfigPath
	default:
		return "", apperrors.ErrNoConfigPath
	}

	return config.ResolveConfigPath(path), nil
}

func (a *argsConfig) getAppConfigOverrides() config.Overrides {
	if a.Run != nil {
		verbosity := a.Run.Verbosity
		if verbosity == 0 && a.Run.Debug {
			verbosity = 2
		}
		return config.Overrides{
			Settings: config.GlobalSettings{
				GherkinLocation: a.Run.GherkinLocation,
				Tags:            a.Run.Tags,
				EnvFile:         a.Run.EnvFile,
				Debug: config.DebugConfig{
					Verbosity: verbosity,
					Scopes:    a.Run.DebugScope,
					Scenario:  a.Run.DebugScenario,
					LogFile:   a.Run.LogFile,
					LogFormat: a.Run.LogFormat,
				},
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

	if a.Install != nil {
		return config.Overrides{
			Settings: config.GlobalSettings{
				EnvFile: a.Install.EnvFile,
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

	if a.Install != nil {
		return config.InstallMode, nil
	}

	if a.ExportStepDefinitions != nil {
		return config.ExportStepDefinitionsMode, nil
	}

	if a.ExportConfigSchema != nil {
		return config.ExportConfigSchemaMode, nil
	}

	return "", apperrors.ErrNoModeProvided
}

func (a *argsConfig) isSilent() bool {
	if a.ExportStepDefinitions != nil && a.ExportStepDefinitions.Silent {
		return true
	}
	return a.ExportConfigSchema != nil && a.ExportConfigSchema.Silent
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
	Headless      bool   `arg:"--headless" help:"headless mode" default:"true"`
	Timeout       string `arg:"--timeout" help:"timeout duration (e.g. 10s, 1m, 2h)"`
	Debug         bool   `arg:"--debug" help:"enable debug output (verbosity 2: headers and variable substitutions)"`
	Verbosity     int    `arg:"--verbosity" help:"debug verbosity: 1=summary, 2=detailed, 3=trace (overrides --debug)" default:"0"`
	DebugScope    string `arg:"--debug-scope" help:"restrict debug to comma-separated scopes: http,browser,variables,config"`
	DebugScenario string `arg:"--debug-scenario" help:"restrict debug output to scenarios whose name or tag contains this value"`
	LogFile       string `arg:"--log-file" help:"write debug output to this file path in addition to stdout"`
	LogFormat     string `arg:"--log-format" help:"log output format: text or json" default:"text"`
}

type initCmd struct {
}

type installCmd struct {
	ConfigPath string `arg:"-c,--config" help:"app config path" default:"testflowkit.yml"`
	EnvFile    string `arg:"--env-file" help:"path to env file"`
}

type validateCmd struct {
	commonCmd
}

type versionCmd struct {
	// No additional fields needed - simple command
}

type exportStepDefinitionsCmd struct {
	Format string `arg:"--format" help:"output format (json)" default:"json"`
	Silent bool   `arg:"--silent" help:"suppress all stderr output"`
}

type exportConfigSchemaCmd struct {
	Format string `arg:"--format" help:"output format (json)" default:"json"`
	Silent bool   `arg:"--silent" help:"suppress all stderr output"`
}

type commonCmd struct {
	GherkinLocation string `arg:"-l,--location" help:"path to gherkin files"`
	ConfigPath      string `arg:"-c,--config" help:"app config path" default:"testflowkit.yml"`
	Tags            string `arg:"-t,--tags" help:"tags"`
	EnvFile         string `arg:"--env-file" help:"path to env file"`
}

func getAppArgs() argsConfig {
	c := argsConfig{}
	arg.MustParse(&c)
	return c
}
