package config

import (
	"time"
)

type argsConfig struct {
	Run      *runCmd      `arg:"subcommand:run" help:"run tests"`
	Init     *initCmd     `arg:"subcommand:init" help:"init cmd config"`
	Validate *validateCmd `arg:"subcommand:validate" help:"validate gherkin files"`
}

type runCmd struct {
	GherkinLocation    string        `arg:"-l,--location" help:"path to gherkin files"`
	ClIConfigPath      string        `arg:"-c,--config" help:"app config path" default:"cmd.yml"`
	FrontendConfigPath string        `arg:"-f,--front-config" help:"front tests config path" default:"frontend.yml"`
	Tags               string        `arg:"-t,--tags" help:"tags"`
	Parallel           int           `arg:"-p,--parallel" help:"number of tests launch in parallel"`
	Timeout            time.Duration `arg:"--timeout" help:"test suite timeout"`
	Headless           bool          `arg:"--headless" help:"display browser" default:"true"`
	AppVersion         string        `arg:"-v,--version" help:"app version" default:"1.0"`
}

type initCmd struct {
	AppName        string `arg:"required" help:"app name"`
	AppDescription string `arg:"-d,--app-description" help:"app description" default:"Cool app"`
	AppVersion     string `arg:"-v,--version" help:"app version" default:"1.0"`
}

type validateCmd struct {
	GherkinLocation    string `arg:"-l,--location" help:"path to gherkin files"`
	ClIConfigPath      string `arg:"-c,--config" help:"app config path" default:"cmd.yml"`
	FrontendConfigPath string `arg:"-f,--front-config" help:"front tests config path" default:"frontend.yml"`
	Tags               string `arg:"-t,--tags" help:"tags"`
}
