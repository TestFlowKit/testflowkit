package main

import (
	"etoolse/cli"
	"etoolse/config"
	"log"
)

func main() {
	cliConfig := config.Init()

	modes := map[config.Mode]func(*config.AppConfig){
		config.RunMode:  cli.Run,
		config.InitMode: cli.Init,
	}

	if mode, ok := modes[cliConfig.Mode]; ok {
		mode(cliConfig)
	} else {
		log.Fatalf("unknown mode: %s", cliConfig.Mode)
	}
}
