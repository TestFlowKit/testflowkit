package scenario

import (
	"encoding/json"
	"fmt"
	"testflowkit/pkg/logger"
)

// SubstituteVariables replaces all variables in the backend context using the provided context.
func (bc *BackendContext) SubstituteVariables(ctx *Context) error {
	// 1. Headers
	newHeaders := make(map[string]string)
	for k, v := range bc.Headers {
		newKey := ReplaceVariablesInString(ctx, k)
		newHeaders[newKey] = ReplaceVariablesInString(ctx, v)
	}
	bc.Headers = newHeaders

	// 2. Endpoint (REST)
	restEndpoint := bc.Rest.Endpoint
	if restEndpoint != nil {
		newQueryParams := make(map[string]string)
		for k, v := range restEndpoint.QueryParams {
			newKey := ReplaceVariablesInString(ctx, k)
			newQueryParams[newKey] = ReplaceVariablesInString(ctx, v)
		}
		restEndpoint.QueryParams = newQueryParams

		newPathParams := make(map[string]string)
		for k, v := range restEndpoint.PathParams {
			newKey := ReplaceVariablesInString(ctx, k)
			newPathParams[newKey] = ReplaceVariablesInString(ctx, v)
		}
		restEndpoint.PathParams = newPathParams
	}

	// 3. RequestBody (REST)
	if len(bc.Rest.RequestBody) > 0 {
		bodyStr := string(bc.Rest.RequestBody)
		newBody := ReplaceVariablesInString(ctx, bodyStr)
		bc.Rest.RequestBody = []byte(newBody)
	}

	if len(bc.GraphQL.Variables) > 0 {
		// Use JSON roundtrip to handle nested variables and type conversion
		err := bc.marshalAndSubstituteVariables(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bc *BackendContext) marshalAndSubstituteVariables(ctx *Context) error {
	jsonData, err := json.Marshal(bc.GraphQL.Variables)
	if err != nil {
		return fmt.Errorf("failed to marshal variables for substitution: %w", err)
	}

	substitutedJSON := ReplaceVariablesInString(ctx, string(jsonData))

	var newVariables map[string]any
	if errJSONDecode := json.Unmarshal([]byte(substitutedJSON), &newVariables); errJSONDecode != nil {
		return fmt.Errorf("failed to unmarshal variables after substitution: %w", errJSONDecode)
	}
	bc.GraphQL.Variables = newVariables

	// Post-processing: Try to parse strings that look like JSON or booleans/numbers
	// This ensures that variables that were substituted into complex types (arrays/objects)
	// are correctly parsed into their Go types.
	for k, v := range bc.GraphQL.Variables {
		if strVal, ok := v.(string); ok {
			if parsed, errParse := bc.parser.ParseValue(strVal); errParse == nil {
				bc.GraphQL.Variables[k] = parsed
			} else {
				logger.Warn(fmt.Sprintf("failed to parse variable '%s': %v", k, errParse), nil)
			}
		}
	}
	return nil
}
