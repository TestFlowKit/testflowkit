package main

import (
	"etoolse/cli"
	"etoolse/config"
	"log"
)

func main() {
	cliConfig := config.Init()
	if cliConfig == nil {
		log.Fatal("no mode specified")
	}

	modes := map[config.Mode]func(*config.AppConfig){
		config.RunMode:        cli.Run,
		config.InitMode:       cli.Init,
		config.ValidationMode: cli.Validate,
	}

	if mode, ok := modes[cliConfig.Mode]; ok {
		mode(cliConfig)
	} else {
		log.Fatalf("unknown mode: %s", cliConfig.Mode)
	}
}
