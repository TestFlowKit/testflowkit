package protocol

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/utils/fileutils"
	"testflowkit/pkg/graphql"
	"testflowkit/pkg/logger"
	"time"
)

type GraphQLAdapter struct{}

func NewGraphQLAdapter() *GraphQLAdapter {
	return &GraphQLAdapter{}
}

func (a *GraphQLAdapter) PrepareRequest(ctx context.Context, apiName string, reqName string) (context.Context, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	cfg := scenarioCtx.GetConfig()

	apiDef, errGetAPI := cfg.GetAPI(apiName)
	if errGetAPI != nil {
		return ctx, fmt.Errorf("API '%s' not found: %w", apiName, errGetAPI)
	}

	if apiDef.Type != config.APITypeGraphQL {
		err := fmt.Errorf("API '%s' is not a GraphQL API, got type '%s'", apiName, apiDef.Type)
		logger.Fatal("graphql adapter", err)
	}

	operation, exists := apiDef.Operations[reqName]
	if !exists {
		return ctx, fmt.Errorf("operation '%s' not found in API '%s'", reqName, apiName)
	}

	query, errGetQuery := a.getQuery(operation.Operation)
	if errGetQuery != nil {
		return ctx, errGetQuery
	}

	req := &graphql.Request{
		Query:     query,
		Variables: scenarioCtx.GetGraphQLVariables(),
	}

	scenarioCtx.SetGraphQLRequest(req)

	scenarioCtx.SetGraphQLEndpoint(apiDef.Endpoint)

	if len(apiDef.DefaultHeaders) > 0 {
		for key, value := range apiDef.DefaultHeaders {
			scenarioCtx.SetGraphQLHeader(key, value)
		}
	}

	scenarioCtx.GetBackendContext().SetProtocol(a)

	return ctx, nil
}

func (*GraphQLAdapter) getQuery(operation string) (string, error) {
	isGqlFilePath := strings.HasSuffix(operation, ".graphql") || strings.HasSuffix(operation, ".gql")
	if !isGqlFilePath {
		return operation, nil
	}

	errPath := fileutils.ValidatePath(operation)
	if errPath != nil {
		return "", errPath
	}
	content, err := os.ReadFile(operation)
	if err != nil {
		return "", fmt.Errorf("failed to read GraphQL query file '%s': %w", operation, err)
	}

	logger.InfoFf("GraphQL query loaded from file: %s", operation)
	return string(content), nil
}

func (a *GraphQLAdapter) SendRequest(ctx context.Context) (context.Context, error) {
	scenarioCtx := scenario.MustFromContext(ctx)

	request := scenarioCtx.GetGraphQLRequest()
	if request == nil {
		return ctx, errors.New("no GraphQL request is prepared to send")
	}

	endpoint := scenarioCtx.GetGraphQLEndpoint()
	if endpoint == "" {
		return ctx, errors.New("no GraphQL endpoint configured")
	}

	headers := scenarioCtx.GetGraphQLHeaders()

	var options []graphql.ClientOption
	if len(headers) > 0 {
		options = append(options, graphql.WithHeaders(headers))
	}
	client := graphql.NewClient(endpoint, options...)

	// Ensure variables are set in the request
	request.Variables = scenarioCtx.GetGraphQLVariables()
	if len(request.Variables) > 0 {
		logger.InfoFf("GraphQL Variables: %v", request.Variables)
	}

	startTime := time.Now()
	response, err := client.Execute(ctx, *request)
	duration := time.Since(startTime)
	if err != nil {
		return ctx, fmt.Errorf("failed to send GraphQL request: %w", err)
	}

	scenarioCtx.SetGraphQLResponse(response)

	logger.InfoFf("GraphQL request completed - Status: %d, Duration: %v, Errors: %d",
		response.StatusCode, duration, len(response.Errors))

	return ctx, nil
}

func (a *GraphQLAdapter) GetResponseBody(ctx context.Context) ([]byte, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	if !backend.HasResponse() {
		return nil, errors.New("no GraphQL response available")
	}

	return backend.GetResponseBody(), nil
}

func (a *GraphQLAdapter) GetStatusCode(ctx context.Context) (int, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	if !backend.HasResponse() {
		return 0, errors.New("no GraphQL response available")
	}

	return backend.GetStatusCode(), nil
}

func (a *GraphQLAdapter) HasErrors(ctx context.Context) bool {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	return backend.HasGraphQLErrors()
}

func (a *GraphQLAdapter) GetProtocolName() string {
	return string(ProtocolGraphQL)
}
