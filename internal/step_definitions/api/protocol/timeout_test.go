package protocol

import (
	"context"
	"testing"
	"time"

	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRESTAPIAdapterPrepareRequestSetsResolvedTimeout(t *testing.T) {
	intPtr := func(v int) *int { return &v }

	cfg := &config.Config{
		APIs: &config.APIsConfig{
			DefaultTimeout: 30000,
			Definitions: map[string]config.APIDefinition{
				"users": {
					Type:    config.APITypeREST,
					BaseURL: "https://api.example.com",
					Timeout: intPtr(12000),
					Endpoints: map[string]config.Endpoint{
						"get_user": {
							Method:      "GET",
							Path:        "/users/{id}",
							Description: "Retrieve a user",
							Timeout:     intPtr(5000),
						},
					},
				},
			},
		},
	}

	scenarioCtx := scenario.NewContext(cfg, nil, nil)
	ctx := scenario.WithContext(context.Background(), scenarioCtx)

	adapter := NewRESTAPIAdapter()
	_, err := adapter.PrepareRequest(ctx, "users", "get_user")
	require.NoError(t, err)

	assert.Equal(t, 5*time.Second, scenarioCtx.GetBackendContext().Timeout)
}

func TestGraphQLAdapterPrepareRequestSetsResolvedTimeout(t *testing.T) {
	intPtr := func(v int) *int { return &v }

	cfg := &config.Config{
		APIs: &config.APIsConfig{
			DefaultTimeout: 30000,
			Definitions: map[string]config.APIDefinition{
				"content": {
					Type:     config.APITypeGraphQL,
					Endpoint: "https://api.example.com/graphql",
					Timeout:  intPtr(15000),
					Operations: map[string]config.GraphQLOperation{
						"get_user": {
							Type:        "query",
							Operation:   "query { user { id } }",
							Description: "Fetch a user",
							Timeout:     intPtr(7000),
						},
					},
				},
			},
		},
	}

	scenarioCtx := scenario.NewContext(cfg, nil, nil)
	ctx := scenario.WithContext(context.Background(), scenarioCtx)

	adapter := NewGraphQLAdapter()
	_, err := adapter.PrepareRequest(ctx, "content", "get_user")
	require.NoError(t, err)

	assert.Equal(t, 7*time.Second, scenarioCtx.GetBackendContext().Timeout)
}
