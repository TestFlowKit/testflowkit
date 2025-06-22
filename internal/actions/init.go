package actions

import (
	_ "embed"
	"os"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

//go:embed boilerplate/cli.boilerplate.yml
var cliConfigTemplate string

//go:embed boilerplate/frontend.boilerplate.yml
var frontTestsConfigTemplate string

func Init(_ *config.App) {
	logger.Info("init cmd config ...")

	if _, err := os.Stat("cmd.yml"); err == nil {
		logger.Fatal("cmd already initialized", err)
	}

	if _, err := os.Stat("frontend.yml"); err == nil {
		logger.Fatal("cmd already initialized", err)
	}

	err := os.WriteFile("cmd.yml", []byte(cliConfigTemplate), 0600)
	if err != nil {
		logger.Fatal("failed to write cmd config: ", err)
	}

	err = os.WriteFile("frontend.yml", []byte(frontTestsConfigTemplate), 0600)
	if err != nil {
		os.Remove("cmd.yml")
		logger.Fatal("failed to write frontend config: ", err)
	}
}
