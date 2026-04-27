package protocol

import (
	"context"
	"testing"

	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRESTAPIAdapterGetCURLCommand(t *testing.T) {
	cfg := &config.Config{
		APIs: &config.APIsConfig{
			Definitions: map[string]config.APIDefinition{
				"users": {
					Type:    config.APITypeREST,
					BaseURL: "https://api.example.com",
					DefaultHeaders: map[string]string{
						"Accept": "application/json",
					},
					SecurityRef: &config.SecurityRef{Name: "bearer_auth"},
					Endpoints: map[string]config.Endpoint{
						"create_user": {
							Method:      "POST",
							Path:        "/users/{id}",
							Description: "Create a user",
						},
					},
				},
			},
		},
		SecuritySchemes: map[string]config.SecurityScheme{
			"bearer_auth": {
				Type:  config.SecurityTypeBearer,
				Token: "secret-token",
			},
		},
	}

	scenarioCtx := scenario.NewContext(cfg, nil, nil)
	ctx := scenario.WithContext(context.Background(), scenarioCtx)

	adapter := NewRESTAPIAdapter()
	_, err := adapter.PrepareRequest(ctx, "users", "create_user")
	require.NoError(t, err)

	scenarioCtx.AddPathParam("id", "42")
	scenarioCtx.AddQueryParam("expand", "roles")
	scenarioCtx.AddHeader("X-Trace-Id", "trace-123")
	require.NoError(t, scenarioCtx.SetRequestBody([]byte(`{"name":"Alice"}`)))

	curlCommand, err := adapter.GetCURLCommand(ctx)
	require.NoError(t, err)

	expectedCurl :=
		`curl -X POST ` +
			`-H "Accept: application/json" ` +
			`-H "Authorization: Bearer secret-token" ` +
			`-H "Content-Type: application/json" ` +
			`-H "X-Trace-Id: trace-123" ` +
			`--data-raw "{\"name\":\"Alice\"}" ` +
			`"https://api.example.com/users/42?expand=roles"`
	assert.Equal(t, expectedCurl, curlCommand)
}

func TestGraphQLAdapterGetCURLCommand(t *testing.T) {
	cfg := &config.Config{
		APIs: &config.APIsConfig{
			Definitions: map[string]config.APIDefinition{
				"catalog": {
					Type:        config.APITypeGraphQL,
					Endpoint:    "https://api.example.com/graphql",
					SecurityRef: &config.SecurityRef{Name: "api_key_query"},
					Operations: map[string]config.GraphQLOperation{
						"find_user": {
							Type:      "query",
							Operation: "query FindUser($id: ID!) { user(id: $id) { id name } }",
						},
					},
				},
			},
		},
		SecuritySchemes: map[string]config.SecurityScheme{
			"api_key_query": {
				Type:       config.SecurityTypeAPIKey,
				Key:        "key-123",
				Placement:  config.APIKeyPlacementQuery,
				QueryParam: "api_key",
			},
		},
	}

	scenarioCtx := scenario.NewContext(cfg, nil, nil)
	ctx := scenario.WithContext(context.Background(), scenarioCtx)

	adapter := NewGraphQLAdapter()
	_, err := adapter.PrepareRequest(ctx, "catalog", "find_user")
	require.NoError(t, err)

	scenarioCtx.GetBackendContext().GraphQL.Variables["id"] = "42"
	scenarioCtx.AddHeader("X-Trace-Id", "trace-graphql")

	curlCommand, err := adapter.GetCURLCommand(ctx)
	require.NoError(t, err)

	expectedCurl :=
		`curl -X POST ` +
			`-H "Content-Type: application/json" ` +
			`-H "X-Trace-Id: trace-graphql" ` +
			`--data-raw "{\"query\":\"query FindUser($id: ID!) { user(id: $id) { id name } }\",` +
			`\"variables\":{\"id\":\"42\"}}" ` +
			`"https://api.example.com/graphql?api_key=key-123"`
	assert.Equal(t, expectedCurl, curlCommand)
}
