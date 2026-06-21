package actions

import (
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

func version(conf *config.Config, _ error) {
	logger.Infof("testflowkit version %s\n", conf.GetVersion())
}
