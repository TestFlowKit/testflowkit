package httpapi

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func userSendsAPostRequest() stepbuilder.TestStep {
	return stepbuilder.NewStepWithTwoVariables(
		[]string{"Post request on api {string} at url {string}"},
		func(ctx *scenario.Context) func(string, string) error {
			return func(apiName, url string) error {
				return nil
			}
		},
		func(url, body string) stepbuilder.ValidationErrors {
			return stepbuilder.ValidationErrors{}
		},
		stepbuilder.StepDefDocParams{
			Description: "Sends a POST request to the specified URL with the given body",
		},
	)
}
