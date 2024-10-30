package config

func Init() ClI {
	cli := ClI{}
	cli.InitByFilePath("cli.yml")

	FrontConfig{}.init("frontend.yml")

	return cli
}
