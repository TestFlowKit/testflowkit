package main

import (
	"log"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/httpapi"
)

func main() {

	steps := httpapi.GetAllSteps()

	handler := steps[0].GetDefinition(&scenario.Context{}).(func(string, string) error)

	err := handler("GET", "test")
	if err != nil {
		log.Fatal(err)
	}
}
