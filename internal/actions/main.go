package actions

import (
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

func Execute(cfg *config.Config, mode config.Mode) {
	modes := map[config.Mode]func(*config.Config){
		config.RunMode:        run,
		config.InitMode:       initMode,
		config.ValidationMode: validate,
	}

	if action, ok := modes[mode]; ok {
		action(cfg)
	} else {
		logger.Fatal(fmt.Sprintf("unknown mode: %s", mode), nil)
	}
}
