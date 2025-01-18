package config

import (
	"time"

	"github.com/alexflint/go-arg"
)

type RunCmd struct {
	GherkinLocation    string        `arg:"-l,--location" help:"path to gherkin files"`
	ClIConfigPath      string        `arg:"-c,--config" help:"app config path" default:"cli.yml"`
	FrontendConfigPath string        `arg:"-f,--front-config" help:"front tests config path" default:"frontend.yml"`
	Tags               string        `arg:"-t,--tags" help:"tags"`
	Parallel           int           `arg:"-p,--parallel" help:"number of tests launch in parallel"`
	Timeout            time.Duration `arg:"--timeout" help:"test suite timeout"`
	Headless           bool          `arg:"--headless" help:"display browser" default:"true"`
	AppVersion         string        `arg:"-v,--version" help:"app version" default:"1.0"`
}

type appArgsConfig struct {
	Run *RunCmd `arg:"subcommand:run" help:"run tests"`
}

func getAppArgs() appArgsConfig {
	args := appArgsConfig{}
	arg.MustParse(&args)

	validateSubcommand(args)

	return args
}

func validateSubcommand(args appArgsConfig) {
	if args.Run == nil {
		panic("subcommand is required")
	}
}
