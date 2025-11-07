package graphql

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewIntrospectionClient(t *testing.T) {
	client := NewClient("http://example.com/graphql")
	introspectionClient := NewIntrospectionClient(client)

	if introspectionClient == nil {
		t.Fatal("Expected introspection client to be created")
	}

	if introspectionClient.client != client {
		t.Error("Expected introspection client to use provided GraphQL client")
	}

	if !introspectionClient.cache.enabled {
		t.Error("Expected cache to be enabled by default")
	}

	if introspectionClient.cache.ttl != 5*time.Minute {
		t.Error("Expected default cache TTL to be 5 minutes")
	}
}

func TestNewIntrospectionClientWithOptions(t *testing.T) {
	client := NewClient("http://example.com/graphql")
	options := IntrospectionOptions{
		CacheEnabled: false,
		CacheTTL:     10 * time.Minute,
	}
	introspectionClient := NewIntrospectionClient(client, options)

	if introspectionClient.cache.enabled {
		t.Error("Expected cache to be disabled")
	}

	if introspectionClient.cache.ttl != 10*time.Minute {
		t.Error("Expected cache TTL to be 10 minutes")
	}
}

func TestBuildIntrospectionQuery(t *testing.T) {
	client := NewClient("http://example.com/graphql")
	introspectionClient := NewIntrospectionClient(client)

	query := introspectionClient.buildIntrospectionQuery()

	if query == "" {
		t.Fatal("Expected introspection query to be built")
	}

	// Check that the query contains essential introspection elements
	expectedElements := []string{
		"IntrospectionQuery",
		"__schema",
		"queryType",
		"mutationType",
		"types",
		"FullType",
		"TypeRef",
		"InputValue",
	}

	for _, element := range expectedElements {
		if !contains(query, element) {
			t.Errorf("Expected introspection query to contain '%s'", element)
		}
	}
}

func TestGetSchema(t *testing.T) {
	// Create a mock GraphQL server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock introspection response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"__schema": map[string]interface{}{
					"queryType": map[string]interface{}{
						"name": "Query",
					},
					"mutationType": map[string]interface{}{
						"name": "Mutation",
					},
					"types": []map[string]interface{}{
						{
							"kind": "OBJECT",
							"name": "User",
							"fields": []map[string]interface{}{
								{
									"name": "id",
									"type": map[string]interface{}{
										"kind": "SCALAR",
										"name": "ID",
									},
									"args": []interface{}{},
								},
								{
									"name": "name",
									"type": map[string]interface{}{
										"kind": "SCALAR",
										"name": "String",
									},
									"args": []interface{}{},
								},
							},
						},
					},
					"directives": []interface{}{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	introspectionClient := NewIntrospectionClient(client)

	schema, err := introspectionClient.GetSchema(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if schema == nil {
		t.Fatal("Expected schema to be returned")
	}

	if schema.QueryType == nil || schema.QueryType.Name != "Query" {
		t.Error("Expected query type to be 'Query'")
	}

	if schema.MutationType == nil || schema.MutationType.Name != "Mutation" {
		t.Error("Expected mutation type to be 'Mutation'")
	}

	if len(schema.Types) != 1 {
		t.Errorf("Expected 1 type, got %d", len(schema.Types))
	}

	userType := schema.Types[0]
	if userType.Name != "User" {
		t.Errorf("Expected type name to be 'User', got '%s'", userType.Name)
	}

	if len(userType.Fields) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(userType.Fields))
	}
}

func TestGetSchemaWithErrors(t *testing.T) {
	// Create a mock GraphQL server that returns errors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"errors": []map[string]interface{}{
				{
					"message": "Introspection is disabled",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	introspectionClient := NewIntrospectionClient(client)

	_, err := introspectionClient.GetSchema(context.Background())
	if err == nil {
		t.Fatal("Expected error when introspection returns errors")
	}

	if !contains(err.Error(), "introspection query returned errors") {
		t.Errorf("Expected error message to mention introspection errors, got: %v", err)
	}
}

func TestSchemaCache(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"__schema": map[string]interface{}{
					"queryType": map[string]interface{}{
						"name": "Query",
					},
					"types":      []interface{}{},
					"directives": []interface{}{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	introspectionClient := NewIntrospectionClient(client)

	// First call should hit the server
	_, err := introspectionClient.GetSchema(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected 1 server call, got %d", callCount)
	}

	// Second call should use cache
	_, err = introspectionClient.GetSchema(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected 1 server call (cached), got %d", callCount)
	}
}

func TestSchemaCacheDisabled(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"__schema": map[string]interface{}{
					"queryType": map[string]interface{}{
						"name": "Query",
					},
					"types":      []interface{}{},
					"directives": []interface{}{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	options := IntrospectionOptions{
		CacheEnabled: false,
		CacheTTL:     5 * time.Minute,
	}
	introspectionClient := NewIntrospectionClient(client, options)

	// First call
	_, err := introspectionClient.GetSchema(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected 1 server call, got %d", callCount)
	}

	// Second call should also hit the server (cache disabled)
	_, err = introspectionClient.GetSchema(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if callCount != 2 {
		t.Errorf("Expected 2 server calls (cache disabled), got %d", callCount)
	}
}

func TestSchemaGetTypeByName(t *testing.T) {
	schema := &Schema{
		Types: []Type{
			{Name: "User", Kind: "OBJECT"},
			{Name: "String", Kind: "SCALAR"},
		},
	}

	userType := schema.GetTypeByName("User")
	if userType == nil {
		t.Fatal("Expected to find User type")
	}

	if userType.Name != "User" {
		t.Errorf("Expected type name to be 'User', got '%s'", userType.Name)
	}

	nonExistentType := schema.GetTypeByName("NonExistent")
	if nonExistentType != nil {
		t.Error("Expected nil for non-existent type")
	}
}

func TestTypeGetFieldByName(t *testing.T) {
	userType := &Type{
		Name: "User",
		Fields: []Field{
			{Name: "id", Type: Type{Kind: "SCALAR", Name: "ID"}},
			{Name: "name", Type: Type{Kind: "SCALAR", Name: "String"}},
		},
	}

	idField := userType.GetFieldByName("id")
	if idField == nil {
		t.Fatal("Expected to find id field")
	}

	if idField.Name != "id" {
		t.Errorf("Expected field name to be 'id', got '%s'", idField.Name)
	}

	nonExistentField := userType.GetFieldByName("nonexistent")
	if nonExistentField != nil {
		t.Error("Expected nil for non-existent field")
	}
}

func TestTypeKindCheckers(t *testing.T) {
	tests := []struct {
		name     string
		typeKind string
		checker  func(*Type) bool
		expected bool
	}{
		{"IsScalarType with SCALAR", "SCALAR", (*Type).IsScalarType, true},
		{"IsScalarType with OBJECT", "OBJECT", (*Type).IsScalarType, false},
		{"IsObjectType with OBJECT", "OBJECT", (*Type).IsObjectType, true},
		{"IsObjectType with SCALAR", "SCALAR", (*Type).IsObjectType, false},
		{"IsListType with LIST", "LIST", (*Type).IsListType, true},
		{"IsListType with OBJECT", "OBJECT", (*Type).IsListType, false},
		{"IsNonNullType with NON_NULL", "NON_NULL", (*Type).IsNonNullType, true},
		{"IsNonNullType with OBJECT", "OBJECT", (*Type).IsNonNullType, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			typ := &Type{Kind: test.typeKind}
			result := test.checker(typ)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestValidateOperation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"__schema": map[string]interface{}{
					"queryType": map[string]interface{}{
						"name": "Query",
					},
					"types": []map[string]interface{}{
						{
							"kind": "OBJECT",
							"name": "Query",
							"fields": []map[string]interface{}{
								{
									"name": "user",
									"type": map[string]interface{}{
										"kind": "OBJECT",
										"name": "User",
									},
									"args": []interface{}{},
								},
							},
						},
						{
							"kind": "OBJECT",
							"name": "User",
							"fields": []map[string]interface{}{
								{
									"name": "id",
									"type": map[string]interface{}{
										"kind": "SCALAR",
										"name": "ID",
									},
									"args": []interface{}{},
								},
								{
									"name": "name",
									"type": map[string]interface{}{
										"kind": "SCALAR",
										"name": "String",
									},
									"args": []interface{}{},
								},
							},
						},
					},
					"directives": []interface{}{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	introspectionClient := NewIntrospectionClient(client)

	// Test valid operation
	operation := `query { user { id name } }`
	err := introspectionClient.ValidateOperation(context.Background(), operation)
	if err != nil {
		t.Errorf("Expected no error for valid operation, got: %v", err)
	}

	// Test invalid operation
	invalidOperation := `query { nonexistentField }`
	err = introspectionClient.ValidateOperation(context.Background(), invalidOperation)
	if err == nil {
		t.Error("Expected error for invalid operation")
	}
}

func TestValidateOperationDetailed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"__schema": map[string]interface{}{
					"queryType": map[string]interface{}{
						"name": "Query",
					},
					"types": []map[string]interface{}{
						{
							"kind": "OBJECT",
							"name": "Query",
							"fields": []map[string]interface{}{
								{
									"name": "user",
									"type": map[string]interface{}{
										"kind": "OBJECT",
										"name": "User",
									},
									"args": []interface{}{},
								},
							},
						},
					},
					"directives": []interface{}{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	introspectionClient := NewIntrospectionClient(client)

	// Test valid operation
	operation := `query { user }`
	result, err := introspectionClient.ValidateOperationDetailed(context.Background(), operation)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if !result.Valid {
		t.Error("Expected operation to be valid")
	}

	// Test invalid operation
	invalidOperation := `query { nonexistentField }`
	result, err = introspectionClient.ValidateOperationDetailed(context.Background(), invalidOperation)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if result.Valid {
		t.Error("Expected operation to be invalid")
	}

	if len(result.Errors) == 0 {
		t.Error("Expected validation errors")
	}
}

func TestValidateFieldPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"__schema": map[string]interface{}{
					"queryType": map[string]interface{}{
						"name": "Query",
					},
					"types": []map[string]interface{}{
						{
							"kind": "OBJECT",
							"name": "Query",
							"fields": []map[string]interface{}{
								{
									"name": "user",
									"type": map[string]interface{}{
										"kind": "OBJECT",
										"name": "User",
									},
									"args": []interface{}{},
								},
							},
						},
						{
							"kind": "OBJECT",
							"name": "User",
							"fields": []map[string]interface{}{
								{
									"name": "id",
									"type": map[string]interface{}{
										"kind": "SCALAR",
										"name": "ID",
									},
									"args": []interface{}{},
								},
							},
						},
					},
					"directives": []interface{}{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	introspectionClient := NewIntrospectionClient(client)

	// Test valid field path
	err := introspectionClient.ValidateFieldPath(context.Background(), "user.id", "query")
	if err != nil {
		t.Errorf("Expected no error for valid field path, got: %v", err)
	}

	// Test invalid field path
	err = introspectionClient.ValidateFieldPath(context.Background(), "user.nonexistent", "query")
	if err == nil {
		t.Error("Expected error for invalid field path")
	}
}

func TestValidateOperation_IntrospectionFailure(t *testing.T) {
	// Create a server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client := NewClient(server.URL, nil)
	introspectionClient := NewIntrospectionClient(client)

	// Test that introspection failure is handled gracefully
	operation := `query { user }`
	err := introspectionClient.ValidateOperation(context.Background(), operation)
	if err == nil {
		t.Error("Expected error when introspection fails")
	}

	if !contains(err.Error(), "schema validation unavailable") {
		t.Errorf("Expected error message to mention schema validation unavailable, got: %v", err)
	}

	// Test detailed validation with introspection failure
	result, err := introspectionClient.ValidateOperationDetailed(context.Background(), operation)
	if err != nil {
		t.Errorf("Expected no error from detailed validation, got: %v", err)
	}

	if result.Valid {
		t.Error("Expected validation result to be invalid when introspection fails")
	}

	if len(result.Errors) == 0 {
		t.Error("Expected validation errors when introspection fails")
	}

	if result.Errors[0].Type != ErrorTypeSchemaUnavailable {
		t.Errorf("Expected error type %v, got %v", ErrorTypeSchemaUnavailable, result.Errors[0].Type)
	}
}

func TestClearCache(t *testing.T) {
	client := NewClient("http://example.com/graphql")
	introspectionClient := NewIntrospectionClient(client)

	// Add something to cache
	introspectionClient.cache.set("test", &Schema{})

	// Verify cache has content
	if len(introspectionClient.cache.schemas) != 1 {
		t.Error("Expected cache to have 1 entry")
	}

	// Clear cache
	introspectionClient.ClearCache()

	// Verify cache is empty
	if len(introspectionClient.cache.schemas) != 0 {
		t.Error("Expected cache to be empty after clearing")
	}
}

func TestSetCacheEnabled(t *testing.T) {
	client := NewClient("http://example.com/graphql")
	introspectionClient := NewIntrospectionClient(client)

	// Initially enabled
	if !introspectionClient.cache.enabled {
		t.Error("Expected cache to be enabled initially")
	}

	// Disable cache
	introspectionClient.SetCacheEnabled(false)
	if introspectionClient.cache.enabled {
		t.Error("Expected cache to be disabled")
	}

	// Re-enable cache
	introspectionClient.SetCacheEnabled(true)
	if !introspectionClient.cache.enabled {
		t.Error("Expected cache to be enabled")
	}
}

func TestSetCacheTTL(t *testing.T) {
	client := NewClient("http://example.com/graphql")
	introspectionClient := NewIntrospectionClient(client)

	newTTL := 10 * time.Minute
	introspectionClient.SetCacheTTL(newTTL)

	if introspectionClient.cache.ttl != newTTL {
		t.Errorf("Expected cache TTL to be %v, got %v", newTTL, introspectionClient.cache.ttl)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
