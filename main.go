package main

import (
	"cucumber/frontend"
	"log"
	"os"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func main() {
	const concurrency = 2
	var opts = godog.Options{
		Output:      colors.Colored(os.Stdout),
		Concurrency: concurrency,
		Format:      "pretty",
		Paths:       []string{"features"},
	}

	testSuite := godog.TestSuite{
		Name:    "App",
		Options: &opts,
		ScenarioInitializer: func(context *godog.ScenarioContext) {
			frontend.InitializeScenario(context)
		},
	}

	status := testSuite.Run()

	if status != 0 {
		log.Fatalf("zero status code expected, %d received", status)
	}
}
