package main

import (
	"etoolse/internal/actions"
	"etoolse/internal/config"
	"log"
)

func main() {
	cliConfig := config.Init()
	if cliConfig == nil {
		log.Fatal("no mode specified")
	}

	modes := map[config.Mode]func(*config.App){
		config.RunMode:        actions.Run,
		config.InitMode:       actions.Init,
		config.ValidationMode: actions.Validate,
	}

	if mode, ok := modes[cliConfig.Mode]; ok {
		mode(cliConfig)
	} else {
		log.Fatalf("unknown mode: %s", cliConfig.Mode)
	}
}
