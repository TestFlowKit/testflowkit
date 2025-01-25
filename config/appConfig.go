package config

type AppConfig struct {
	Mode Mode
	appDetailsConfig
	reportingConfig
	testingConfig
}

func (c *AppConfig) GetConcurrency() int {
	if !c.Headless {
		return 0
	}
	return c.Parallel
}

func initAppConfig(args appArgsConfig, fileConfig appFileConfig) *AppConfig {
	c := AppConfig{
		appDetailsConfig: appDetailsConfig{
			AppName:        fileConfig.AppName,
			AppDescription: fileConfig.AppDescription,
		},
		testingConfig: testingConfig{
			GherkinLocation: fileConfig.GherkinLocation,
			Timeout:         fileConfig.Timeout,
		},
	}

	c.ReportFormat = fileConfig.ReportFormat
	c.SlowMotion = fileConfig.SlowMotion

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

func fillConfigForRunCmd(c *AppConfig, runArgs RunCmd) {
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
