package config

import (
	"time"

	"github.com/alexflint/go-arg"
)

type appArgs struct {
	GherkinLocation    string        `arg:"-l,--location" help:"path to gherkin files" default:"features"`
	ClIConfigPath      string        `arg:"-c,--config" help:"app config path" default:"cli.yml"`
	FrontendConfigPath string        `arg:"-f,--front-config" help:"front tests config path" default:"frontend.yml"`
	Tags               string        `arg:"-t,--tags" help:"tags"`
	Parallel           int           `arg:"-p,--parallel" help:"number of tests launch in parallel"`
	Timeout            time.Duration `arg:"--timeout" help:"test suite timeout"`
	DisplayBrowser     bool          `arg:"-d,--display-browser" help:"display browser"`
}

func getAppArgs() appArgs {
	args := appArgs{}
	arg.MustParse(&args)

	return args
}

func overrideClIConfig(args appArgs, cliConf ClI) ClI {
	if args.GherkinLocation != "" {
		cliConf.GherkinLocation = args.GherkinLocation
	}

	if args.Timeout.Milliseconds() > 0 {
		cliConf.Timeout = args.Timeout.String()
	}

	if args.Tags != "" {
		cliConf.Tags = args.Tags
	}

	if args.Parallel >= 0 {
		cliConf.Parallel = args.Parallel
	}

	cliConf.displayBrowser = args.DisplayBrowser

	return cliConf
}
