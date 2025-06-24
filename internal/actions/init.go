package actions

import (
	_ "embed"
	"os"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

//go:embed boilerplate/config.boilerplate.yml
var configTemplate string

func initMode(_ *config.Config) {
	logger.Info("init cmd config ...")

	if _, err := os.Stat("config.yml"); err == nil {
		logger.Fatal("config already initialized", err)
	}

	err := os.WriteFile("config.yml", []byte(configTemplate), 0600)
	if err != nil {
		logger.Fatal("failed to write cmd config: ", err)
	}
}
