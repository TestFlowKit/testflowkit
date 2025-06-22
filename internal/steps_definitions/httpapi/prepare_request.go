package httpapi

import (
	"fmt"
	"strings"

	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st steps) prepareRequest() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^I prepare a (GET|POST|PUT|DELETE) request for the {string} endpoint$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(method, endpointName string) error {
				apiCfg, err := config.GetAPIConfig()
				if err != nil {
					return err
				}
				endpointKey := strings.ToLower(endpointName)
				endpointPath, ok := apiCfg.Endpoints[endpointKey]
				if !ok {
					return fmt.Errorf("endpoint name '%s' not found in configuration", endpointName)
				}
				fullURL := apiCfg.BaseURL + endpointPath
				ctx.HttpContext.EndpointConfigs = map[string]string{endpointKey: fullURL}
				ctx.HttpContext.Method = method
				ctx.HttpContext.EndpointName = endpointName
				ctx.HttpContext.Headers = make(map[string]string)
				for k, v := range apiCfg.DefaultHeaders {
					ctx.HttpContext.Headers[k] = v
				}
				ctx.HttpContext.QueryParams = make(map[string]string)
				ctx.HttpContext.RequestBody = nil
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Prepares an HTTP request with the specified method and endpoint.",
			Variables: []stepbuilder.DocVariable{
				{Name: "method", Description: "The HTTP method (GET, POST, PUT, DELETE).", Type: stepbuilder.VarTypeEnum("GET", "POST", "PUT", "DELETE")},
				{Name: "endpointName", Description: "The name of the endpoint to request.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When I prepare a GET request for the \"users\" endpoint",
			Category: stepbuilder.Navigation,
		},
	)
}
