package config

func Init() ClI {
	args := getAppArgs()

	cli := ClI{}
	cli.InitByFilePath(args.ClIConfigPath)

	FrontConfig{}.init(args.FrontendConfigPath)

	return overrideClIConfig(args, cli)
}
