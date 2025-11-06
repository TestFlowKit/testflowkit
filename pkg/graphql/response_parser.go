package graphql

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// parseResponse parses a raw HTTP response body into a structured GraphQL response.
func parseResponse(responseBody []byte, statusCode int) (*Response, error) {
	// Handle empty response body
	if len(responseBody) == 0 {
		return nil, NewNetworkError(
			"received empty response body",
			map[string]interface{}{
				"status_code": statusCode,
			},
		)
	}

	// Parse JSON response
	var graphqlResp Response
	if err := json.Unmarshal(responseBody, &graphqlResp); err != nil {
		return nil, NewNetworkError(
			"failed to parse GraphQL response JSON",
			map[string]interface{}{
				"parse_error":   err.Error(),
				"response_body": string(responseBody),
				"status_code":   statusCode,
			},
		)
	}

	// Set status code for reference
	graphqlResp.StatusCode = statusCode

	// Validate response structure
	if err := validateResponseStructure(&graphqlResp); err != nil {
		return nil, NewGraphQLError(
			"invalid GraphQL response structure",
			map[string]interface{}{
				"validation_error": err.Error(),
				"response_body":    string(responseBody),
			},
		)
	}

	// Process and enhance error information
	processGraphQLErrors(&graphqlResp)

	return &graphqlResp, nil
}

// validateResponseStructure ensures the response has a valid GraphQL structure.
func validateResponseStructure(resp *Response) error {
	// A valid GraphQL response must have either data or errors (or both)
	hasData := len(resp.Data) > 0
	hasErrors := len(resp.Errors) > 0

	if !hasData && !hasErrors {
		return errors.New("response must contain either data or errors")
	}

	// Validate error structure if errors exist
	if hasErrors {
		for i, err := range resp.Errors {
			if err.Message == "" {
				return fmt.Errorf("error at index %d is missing required 'message' field", i)
			}
		}
	}

	return nil
}

// processGraphQLErrors enhances error information for better handling.
func processGraphQLErrors(resp *Response) {
	if len(resp.Errors) == 0 {
		return
	}

	// Process each error to ensure consistent structure
	for i := range resp.Errors {
		err := &resp.Errors[i]

		// Ensure extensions map is initialized
		if err.Extensions == nil {
			err.Extensions = make(map[string]interface{})
		}

		// Add error classification if not present
		if _, exists := err.Extensions["classification"]; !exists {
			err.Extensions["classification"] = classifyGraphQLError(err)
		}

		// Add severity if not present
		if _, exists := err.Extensions["severity"]; !exists {
			err.Extensions["severity"] = determineSeverity(err)
		}
	}
}

// classifyGraphQLError attempts to classify the error type based on message and path.
func classifyGraphQLError(err *Error) string {
	message := strings.ToLower(err.Message)

	if isAuthError := checkIfIsAuthError(message); isAuthError {
		return "AUTH_ERROR"
	}

	// Validation errors
	validationErrorTerms := [][]string{
		{"validation"},
		{"invalid"},
	}

	if isMessageContainsSubstrings(message, validationErrorTerms) {
		return "VALIDATION_ERROR"
	}

	fieldErrorTerms := [][]string{
		{"field", "not found"},
		{"field", "unknown"},
	}

	if isMessageContainsSubstrings(message, fieldErrorTerms) {
		return "FIELD_ERROR"
	}

	// Syntax errors
	if strings.Contains(message, "syntax") || strings.Contains(message, "parse") {
		return "SYNTAX_ERROR"
	}

	// Internal server errors

	internalErrorTerms := [][]string{
		{"internal"},
		{"server error"},
	}

	if isMessageContainsSubstrings(message, internalErrorTerms) {
		return "INTERNAL_ERROR"
	}

	// Default classification
	return "GRAPHQL_ERROR"
}

func checkIfIsAuthError(message string) bool {
	authErrorsSubStrings := [][]string{
		{"unauthorized"},
		{"forbidden"},
		{"authentication"},
		{"permission"},
		{"unauthenticated"},
		{"invalid", "token"},
		{"invalid", "auth"},
		{"invalid", "password"},
		{"invalid", "credentials"},
	}

	return isMessageContainsSubstrings(message, authErrorsSubStrings)
}

// isMessageContainsSubstrings checks if all substrings in any pack are present in the message.
//
// Parameters:
//   - message: The string to search within.
//   - subStringPacks: A slice of packs, where each pack is a slice of substrings.
//     The function returns true if, for any pack, all substrings in that pack are found in the message.
//
// Returns:
//   - bool: true if at least one pack is fully matched in the message; false otherwise.
func isMessageContainsSubstrings(message string, subStringPacks [][]string) bool {
	for _, subStrings := range subStringPacks {
		totalFound := 0
		for _, sub := range subStrings {
			if strings.Contains(message, sub) {
				totalFound++
			}
		}

		if totalFound == len(subStrings) {
			return true
		}
	}
	return false
}

// determineSeverity determines error severity based on error characteristics.
func determineSeverity(err *Error) string {
	message := strings.ToLower(err.Message)

	// Critical errors that prevent execution
	if strings.Contains(message, "syntax") || strings.Contains(message, "parse") ||
		strings.Contains(message, "internal") || strings.Contains(message, "server error") {
		return "CRITICAL"
	}

	// High severity errors
	if strings.Contains(message, "unauthorized") || strings.Contains(message, "forbidden") {
		return "HIGH"
	}

	// Medium severity errors
	if strings.Contains(message, "validation") || strings.Contains(message, "field") {
		return "MEDIUM"
	}

	// Default to low severity
	return "LOW"
}

// GetErrorSummary returns a structured summary of all errors in the response.
func (r *Response) GetErrorSummary() ErrorSummary {
	summary := ErrorSummary{
		TotalErrors:      len(r.Errors),
		ErrorsByType:     make(map[string]int),
		ErrorsBySeverity: make(map[string]int),
		Messages:         make([]string, len(r.Errors)),
	}

	for i, err := range r.Errors {
		summary.Messages[i] = err.Message

		// Count by classification
		if classification, exists := err.Extensions["classification"]; exists {
			if classStr, ok := classification.(string); ok {
				summary.ErrorsByType[classStr]++
			}
		}

		// Count by severity
		if severity, exists := err.Extensions["severity"]; exists {
			if sevStr, ok := severity.(string); ok {
				summary.ErrorsBySeverity[sevStr]++
			}
		}
	}

	return summary
}

// HasCriticalErrors checks if the response contains any critical errors.
func (r *Response) HasCriticalErrors() bool {
	for _, err := range r.Errors {
		if severity, exists := err.Extensions["severity"]; exists {
			if sevStr, ok := severity.(string); ok && sevStr == "CRITICAL" {
				return true
			}
		}
	}
	return false
}

// GetErrorsByClassification returns errors filtered by classification.
func (r *Response) GetErrorsByClassification(classification string) []Error {
	filteredErrors := make([]Error, 0, len(r.Errors))
	for _, err := range r.Errors {
		if class, exists := err.Extensions["classification"]; exists {
			if classStr, ok := class.(string); ok && classStr == classification {
				filteredErrors = append(filteredErrors, err)
			}
		}
	}
	return filteredErrors
}

// GetErrorsBySeverity returns errors filtered by severity level.
func (r *Response) GetErrorsBySeverity(severity string) []Error {
	filteredErrors := make([]Error, 0, len(r.Errors))
	for _, err := range r.Errors {
		if sev, exists := err.Extensions["severity"]; exists {
			if sevStr, ok := sev.(string); ok && sevStr == severity {
				filteredErrors = append(filteredErrors, err)
			}
		}
	}
	return filteredErrors
}
