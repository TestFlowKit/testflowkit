package httpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

// Enhanced API step definitions using the new configuration system.
// These steps provide comprehensive API testing capabilities with environment-aware
// endpoint resolution and detailed request/response handling.

// prepareRequestEnhanced prepares an HTTP request using enhanced configuration.
// This step resolves endpoint names to full URLs using the current environment
// and automatically applies default headers from configuration.
//
// Example usage in Gherkin:
//   Given I prepare a request for the "get_user" endpoint
//   Given I prepare a request for the "create_user" endpoint
func (st steps) prepareRequestEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I prepare a request for the {string} endpoint$`},
		func(ctx *scenario.Context) func(string) error {
			return func(endpointName string) error {
				// Get enhanced configuration
				cfg, err := config.GetEnhancedConfig()
				if err != nil {
					return fmt.Errorf("enhanced configuration required for API requests: %w", err)
				}

				// Resolve endpoint name to URL and get endpoint configuration
				url, endpoint, err := cfg.GetAPIEndpoint(endpointName)
				if err != nil {
					return fmt.Errorf("failed to resolve endpoint '%s': %w", endpointName, err)
				}

				logger.InfoFf("Preparing %s request for endpoint '%s': %s", endpoint.Method, endpointName, url)

				// Initialize HTTP context
				ctx.HttpContext.Method = endpoint.Method
				ctx.HttpContext.URL = url
				ctx.HttpContext.EndpointName = endpointName
				ctx.HttpContext.EndpointDescription = endpoint.Description

				// Initialize headers with default headers from configuration
				ctx.HttpContext.Headers = make(map[string]string)
				for key, value := range cfg.Backend.DefaultHeaders {
					ctx.HttpContext.Headers[key] = value
				}

				// Initialize other request components
				ctx.HttpContext.QueryParams = make(map[string]string)
				ctx.HttpContext.PathParams = make(map[string]string)
				ctx.HttpContext.RequestBody = nil

				logger.InfoFf("Request prepared for endpoint '%s' (%s)", endpointName, endpoint.Description)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Prepares an HTTP request for a configured endpoint with automatic method and URL resolution.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "endpointName",
					Description: "The logical endpoint name as defined in the configuration (e.g., 'get_user', 'create_product')",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Given I prepare a request for the "get_user" endpoint`,
			Category: stepbuilder.APITesting,
		},
	)
}

// setRequestHeaderEnhanced sets a custom header for the prepared request.
// This step allows overriding default headers or adding new ones.
func (st steps) setRequestHeaderEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^I set the request header {string} to {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(headerName, headerValue string) error {
				if ctx.HttpContext.Headers == nil {
					ctx.HttpContext.Headers = make(map[string]string)
				}

				ctx.HttpContext.Headers[headerName] = headerValue
				logger.InfoFf("Set request header '%s': %s", headerName, headerValue)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a custom header for the HTTP request.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "headerName",
					Description: "The header name (e.g., 'Authorization', 'Content-Type')",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "headerValue",
					Description: "The header value",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I set the request header "Authorization" to "Bearer token123"`,
			Category: stepbuilder.APITesting,
		},
	)
}

// setPathParameterEnhanced sets a path parameter for the endpoint URL.
// This step handles parameterized endpoints like "/users/{id}".
func (st steps) setPathParameterEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^I set the path parameter {string} to {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(paramName, paramValue string) error {
				if ctx.HttpContext.PathParams == nil {
					ctx.HttpContext.PathParams = make(map[string]string)
				}

				ctx.HttpContext.PathParams[paramName] = paramValue
				
				// Update the URL with the path parameter
				placeholder := fmt.Sprintf("{%s}", paramName)
				ctx.HttpContext.URL = strings.ReplaceAll(ctx.HttpContext.URL, placeholder, paramValue)

				logger.InfoFf("Set path parameter '%s' to '%s'", paramName, paramValue)
				logger.InfoFf("Updated URL: %s", ctx.HttpContext.URL)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a path parameter for parameterized endpoints (e.g., /users/{id}).",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "paramName",
					Description: "The parameter name as defined in the endpoint path",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "paramValue",
					Description: "The parameter value to substitute",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I set the path parameter "id" to "123"`,
			Category: stepbuilder.APITesting,
		},
	)
}

// setQueryParameterEnhanced sets a query parameter for the request.
func (st steps) setQueryParameterEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^I set the query parameter {string} to {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(paramName, paramValue string) error {
				if ctx.HttpContext.QueryParams == nil {
					ctx.HttpContext.QueryParams = make(map[string]string)
				}

				ctx.HttpContext.QueryParams[paramName] = paramValue
				logger.InfoFf("Set query parameter '%s' to '%s'", paramName, paramValue)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a query parameter for the HTTP request.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "paramName",
					Description: "The query parameter name",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "paramValue",
					Description: "The query parameter value",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I set the query parameter "page" to "1"`,
			Category: stepbuilder.APITesting,
		},
	)
}

// setRequestBodyEnhanced sets the request body with enhanced JSON validation.
func (st steps) setRequestBodyEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^I set the request body to:$`},
		func(ctx *scenario.Context) func(string) error {
			return func(body string) error {
				// Validate JSON if Content-Type is application/json
				if contentType, ok := ctx.HttpContext.Headers["Content-Type"]; ok {
					if strings.Contains(contentType, "application/json") {
						var jsonTest interface{}
						if err := json.Unmarshal([]byte(body), &jsonTest); err != nil {
							return fmt.Errorf("invalid JSON in request body: %w", err)
						}
						logger.Info("Request body validated as valid JSON")
					}
				}

				ctx.HttpContext.RequestBody = []byte(body)
				logger.InfoFf("Set request body (%d bytes)", len(body))
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the request body content with automatic JSON validation.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "body",
					Description: "The request body content (JSON, XML, or plain text)",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `And I set the request body to:
  """
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  """`,
			Category: stepbuilder.APITesting,
		},
	)
}

// sendRequestEnhanced sends the prepared HTTP request with enhanced error handling.
func (st steps) sendRequestEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`^I send the request$`},
		func(ctx *scenario.Context) func() error {
			return func() error {
				logger.InfoFf("Sending %s request to: %s", ctx.HttpContext.Method, ctx.HttpContext.URL)

				// Create HTTP client with timeout from configuration
				cfg, err := config.GetEnhancedConfig()
				if err != nil {
					return fmt.Errorf("enhanced configuration required for request: %w", err)
				}

				client := &http.Client{
					Timeout: time.Duration(cfg.Settings.DefaultTimeout) * time.Millisecond,
				}

				// Build complete URL with query parameters
				finalURL := ctx.HttpContext.URL
				if len(ctx.HttpContext.QueryParams) > 0 {
					params := make([]string, 0, len(ctx.HttpContext.QueryParams))
					for key, value := range ctx.HttpContext.QueryParams {
						params = append(params, fmt.Sprintf("%s=%s", key, value))
					}
					finalURL += "?" + strings.Join(params, "&")
				}

				// Create request
				var bodyReader io.Reader
				if ctx.HttpContext.RequestBody != nil {
					bodyReader = strings.NewReader(string(ctx.HttpContext.RequestBody))
				}

				req, err := http.NewRequest(ctx.HttpContext.Method, finalURL, bodyReader)
				if err != nil {
					return fmt.Errorf("failed to create HTTP request: %w", err)
				}

				// Add headers
				for key, value := range ctx.HttpContext.Headers {
					req.Header.Set(key, value)
				}

				// Send request
				startTime := time.Now()
				resp, err := client.Do(req)
				duration := time.Since(startTime)

				if err != nil {
					return fmt.Errorf("failed to send request: %w", err)
				}
				defer resp.Body.Close()

				// Read response body
				responseBody, err := io.ReadAll(resp.Body)
				if err != nil {
					return fmt.Errorf("failed to read response body: %w", err)
				}

				// Store response in context
				ctx.HttpContext.Response = &http.Response{
					StatusCode: resp.StatusCode,
					Status:     resp.Status,
					Header:     resp.Header,
				}
				ctx.HttpContext.ResponseBody = responseBody
				ctx.HttpContext.RequestDuration = duration

				logger.InfoFf("Request completed - Status: %d, Duration: %v, Response size: %d bytes", 
					resp.StatusCode, duration, len(responseBody))

				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sends the prepared HTTP request and stores the response.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I send the request`,
			Category:    stepbuilder.APITesting,
		},
	)
}

// verifyResponseStatusEnhanced verifies the response status code.
func (st steps) verifyResponseStatusEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the response status code should be {int}$`},
		func(ctx *scenario.Context) func(int) error {
			return func(expectedStatus int) error {
				if ctx.HttpContext.Response == nil {
					return fmt.Errorf("no response available - send a request first")
				}

				actualStatus := ctx.HttpContext.Response.StatusCode
				if actualStatus != expectedStatus {
					return fmt.Errorf("expected status code %d, but got %d", expectedStatus, actualStatus)
				}

				logger.InfoFf("Response status code verified: %d", actualStatus)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the response has the expected HTTP status code.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "statusCode",
					Description: "The expected HTTP status code (e.g., 200, 201, 404)",
					Type:        stepbuilder.VarTypeInt,
				},
			},
			Example:  `Then the response status code should be 200`,
			Category: stepbuilder.APITesting,
		},
	)
}

// verifyResponseContainsEnhanced verifies that the response body contains specific text.
func (st steps) verifyResponseContainsEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the response body should contain {string}$`},
		func(ctx *scenario.Context) func(string) error {
			return func(expectedText string) error {
				if ctx.HttpContext.ResponseBody == nil {
					return fmt.Errorf("no response body available - send a request first")
				}

				responseBodyStr := string(ctx.HttpContext.ResponseBody)
				if !strings.Contains(responseBodyStr, expectedText) {
					return fmt.Errorf("response body does not contain expected text '%s'", expectedText)
				}

				logger.InfoFf("Response body contains expected text: '%s'", expectedText)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the response body contains the specified text.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "text",
					Description: "The text that should be present in the response body",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And the response body should contain "success"`,
			Category: stepbuilder.APITesting,
		},
	)
}

// verifyResponseHeaderEnhanced verifies a specific response header value.
func (st steps) verifyResponseHeaderEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the response header {string} should be {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(headerName, expectedValue string) error {
				if ctx.HttpContext.Response == nil {
					return fmt.Errorf("no response available - send a request first")
				}

				actualValue := ctx.HttpContext.Response.Header.Get(headerName)
				if actualValue != expectedValue {
					return fmt.Errorf("expected header '%s' to be '%s', but got '%s'", 
						headerName, expectedValue, actualValue)
				}

				logger.InfoFf("Response header '%s' verified: %s", headerName, actualValue)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that a response header has the expected value.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "headerName",
					Description: "The header name to verify",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "expectedValue",
					Description: "The expected header value",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And the response header "Content-Type" should be "application/json"`,
			Category: stepbuilder.APITesting,
		},
	)
}

// storeResponseValueEnhanced stores a value from the response for later use.
func (st steps) storeResponseValueEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^I store the response body path {string} as {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(jsonPath, variableName string) error {
				if ctx.HttpContext.ResponseBody == nil {
					return fmt.Errorf("no response body available - send a request first")
				}

				// Parse JSON response
				var jsonData interface{}
				if err := json.Unmarshal(ctx.HttpContext.ResponseBody, &jsonData); err != nil {
					return fmt.Errorf("failed to parse response as JSON: %w", err)
				}

				// Extract value using simple path notation (e.g., "user.id", "data.items.0.name")
				value, err := extractJSONPath(jsonData, jsonPath)
				if err != nil {
					return fmt.Errorf("failed to extract path '%s' from response: %w", jsonPath, err)
				}

				// Store value in context for later use
				if ctx.StoredValues == nil {
					ctx.StoredValues = make(map[string]interface{})
				}
				ctx.StoredValues[variableName] = value

				logger.InfoFf("Stored response value '%s' as '%s': %v", jsonPath, variableName, value)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a value from the JSON response body for later use in the test.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "jsonPath",
					Description: "The JSON path to extract (e.g., 'user.id', 'data.items.0.name')",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "variableName",
					Description: "The variable name to store the value under",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I store the response body path "user.id" as "user_id"`,
			Category: stepbuilder.APITesting,
		},
	)
}

// Helper function to extract values from JSON using simple path notation
func extractJSONPath(data interface{}, path string) (interface{}, error) {
	parts := strings.Split(path, ".")
	current := data

	for _, part := range parts {
		switch v := current.(type) {
		case map[string]interface{}:
			if value, ok := v[part]; ok {
				current = value
			} else {
				return nil, fmt.Errorf("key '%s' not found in object", part)
			}
		case []interface{}:
			// Handle array index (simple implementation)
			return nil, fmt.Errorf("array indexing not yet implemented in path '%s'", path)
		default:
			return nil, fmt.Errorf("cannot navigate path '%s' - invalid structure at '%s'", path, part)
		}
	}

	return current, nil
} 