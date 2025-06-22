package httpapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (st *steps) sendRequest() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`^I send the request$`},
		func(ctx *scenario.Context) func() error {
			return func() error {
				if ctx.HttpContext.Method == "" {
					return fmt.Errorf("request has not been prepared. Please use 'I prepare a ... request' first")
				}
				url := ctx.HttpContext.EndpointConfigs[strings.ToLower(ctx.HttpContext.EndpointName)]
				req, err := http.NewRequest(ctx.HttpContext.Method, url, bytes.NewBuffer(ctx.HttpContext.RequestBody))
				if err != nil {
					return fmt.Errorf("failed to build http request: %w", err)
				}
				for key, value := range ctx.HttpContext.Headers {
					req.Header.Set(key, value)
				}
				params := req.URL.Query()
				for key, value := range ctx.HttpContext.QueryParams {
					params.Add(key, value)
				}
				req.URL.RawQuery = params.Encode()
				resp, err := ctx.HttpContext.Client.Do(req)
				if err != nil {
					return fmt.Errorf("failed to send http request: %w", err)
				}
				defer resp.Body.Close()
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					return fmt.Errorf("failed to read response body: %w", err)
				}
				ctx.HttpContext.Response = scenario.HttpResponse{StatusCode: resp.StatusCode, Headers: resp.Header, Body: bodyBytes}
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sends the prepared HTTP request.",
			Variables:   nil,
			Example:     "When I send the request",
			Category:    stepbuilder.Navigation,
		},
	)
}
