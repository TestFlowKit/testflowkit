package config

type AppConfig struct {
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

func InitAppConfig(args appArgsConfig, fileConfig appFileConfig) *AppConfig {
	c := AppConfig{}

	c.AppName = fileConfig.AppName
	c.AppDescription = fileConfig.AppDescription
	c.ReportFormat = fileConfig.ReportFormat
	c.SlowMotion = fileConfig.SlowMotion

	c.Tags = args.Tags
	c.Parallel = args.Parallel
	c.AppVersion = args.AppVersion
	c.Headless = args.Headless

	if args.GherkinLocation == "" {
		c.GherkinLocation = fileConfig.GherkinLocation
	} else {
		c.GherkinLocation = args.GherkinLocation
	}

	if args.Timeout > 0 {
		c.Timeout = args.Timeout.String()
	} else {
		c.Timeout = fileConfig.Timeout
	}

	return &c
}
