package actions

import (
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

func Execute(cfg *config.Config, cfgErr error, mode config.Mode) {
	modes := map[config.Mode]func(*config.Config, error){
		config.RunMode:        run,
		config.InitMode:       initMode,
		config.ValidationMode: validate,
		config.VersionMode:    version,
	}

	if action, ok := modes[mode]; ok {
		action(cfg, cfgErr)
	} else {
		logger.Fatal(fmt.Sprintf("unknown mode: %s", mode), nil)
	}
}
