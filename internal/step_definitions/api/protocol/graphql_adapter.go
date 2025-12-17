package protocol

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"os"
	"strings"
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

func (a *GraphQLAdapter) PrepareRequest(ctx context.Context, operationName string) (context.Context, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	cfg := scenarioCtx.GetConfig()

	operation, err := cfg.GetGraphQLOperation(operationName)
	if err != nil {
		return ctx, err
	}

	query, err := a.getQuery(operation.Operation)
	if err != nil {
		return ctx, err
	}

	req := &graphql.Request{
		Query:     query,
		Variables: scenarioCtx.GetGraphQLVariables(),
	}

	scenarioCtx.SetGraphQLRequest(req)

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
	cfg := scenarioCtx.GetConfig()

	request := scenarioCtx.GetGraphQLRequest()
	if request == nil {
		return ctx, errors.New("no GraphQL request is prepared to send")
	}

	endpoint, err := cfg.GetGraphQLEndpoint()
	if err != nil {
		return ctx, fmt.Errorf("failed to get GraphQL endpoint: %w", err)
	}

	headers := cfg.GetGraphQLHeaders()
	maps.Copy(headers, scenarioCtx.GetGraphQLHeaders())

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
