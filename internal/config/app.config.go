package config

type App struct {
	Mode Mode
	appDetailsConfig
	reportingConfig
	testingConfig
}

func (c *App) GetConcurrency() int {
	if !c.Headless {
		return 0
	}
	return c.Parallel
}

func initAppConfig(args argsConfig, cliConfig cliConfig) *App {
	c := App{
		appDetailsConfig: appDetailsConfig{
			AppName:        cliConfig.AppName,
			AppDescription: cliConfig.AppDescription,
		},
		testingConfig: testingConfig{
			GherkinLocation: cliConfig.GherkinLocation,
			Timeout:         cliConfig.Timeout,
		},
	}

	c.ReportFormat = cliConfig.ReportFormat
	c.SlowMotion = cliConfig.SlowMotion

	if args.Run != nil {
		fillConfigForRunCmd(&c, *args.Run)
	}

	if c.GherkinLocation == "" {
		c.GherkinLocation = defaultCliConfigPath
	}

	if c.ReportFormat == "" {
		c.ReportFormat = defaultReportFormat
	}

	if c.Timeout == "" {
		c.Timeout = defaultTimeout
	}

	return &c
}

func fillConfigForRunCmd(c *App, runArgs runCmd) {
	c.Mode = "run"
	c.Tags = runArgs.Tags
	c.Parallel = runArgs.Parallel
	c.AppVersion = runArgs.AppVersion
	c.Headless = runArgs.Headless

	if runArgs.GherkinLocation != "" {
		c.GherkinLocation = runArgs.GherkinLocation
	}

	if runArgs.Timeout > 0 {
		c.Timeout = runArgs.Timeout.String()
	}
}

const (
	defaultReportFormat = "html"
	defaultTimeout      = "3m"
)
