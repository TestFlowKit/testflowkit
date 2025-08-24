package actions

import (
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

func version(conf *config.Config, _ error) {
	logger.InfoFf("testflowkit version %s\n", conf.GetVersion())
}
