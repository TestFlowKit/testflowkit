package graphql

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	endpoint := "https://api.example.com/graphql"

	// Test basic client creation
	client := NewClient(endpoint)
	if client.GetEndpoint() != endpoint {
		t.Errorf("Expected endpoint %s, got %s", endpoint, client.GetEndpoint())
	}

	// Test with options
	headers := map[string]string{
		"Authorization": "Bearer token",
		"User-Agent":    "TestClient/1.0",
	}

	client = NewClient(endpoint,
		WithTimeout(10*time.Second),
		WithHeaders(headers),
	)

	clientHeaders := client.GetHeaders()
	for key, expectedValue := range headers {
		if actualValue, exists := clientHeaders[key]; !exists || actualValue != expectedValue {
			t.Errorf("Expected header %s: %s, got %s", key, expectedValue, actualValue)
		}
	}
}

func TestClient_SetHeaders(t *testing.T) {
	client := NewClient("https://api.example.com/graphql")

	// Test setting single header
	client.SetHeaders(map[string]string{
		"Authorization": "Bearer token",
	})
	headers := client.GetHeaders()
	if headers["Authorization"] != "Bearer token" {
		t.Error("Failed to set single header")
	}

	// Test setting multiple headers
	newHeaders := map[string]string{
		"X-API-Key":  "api-key",
		"User-Agent": "TestClient/1.0",
	}
	client.SetHeaders(newHeaders)

	headers = client.GetHeaders()
	for key, expectedValue := range newHeaders {
		if actualValue := headers[key]; actualValue != expectedValue {
			t.Errorf("Expected header %s: %s, got %s", key, expectedValue, actualValue)
		}
	}
}

func TestClient_Execute(t *testing.T) {
	// Create a test server
	server := createTestServer(t)
	defer server.Close()

	// Create client
	client := NewClient(server.URL)

	// Execute request
	req := Request{
		Query: `query { user { id name } }`,
		Variables: map[string]interface{}{
			"id": "123",
		},
	}

	ctx := context.Background()
	response, err := client.Execute(ctx, req)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !response.HasData() {
		t.Error("Expected response to have data")
	}

	if response.HasErrors() {
		t.Error("Expected response to have no errors")
	}

	// Test data extraction
	parser := response.GetParser()
	userNameResult, err := parser.GetDataAtPath("user.name")
	if err != nil {
		t.Fatalf("Failed to extract user name: %v", err)
	}

	if userNameResult.String() != "John Doe" {
		t.Errorf("Expected user name 'John Doe', got %s", userNameResult.String())
	}
}

func createTestServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and content type
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
		}

		// Parse request body
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Failed to decode request: %v", err)
		}

		// Verify request content
		if req.Query == "" {
			t.Error("Expected non-empty query")
		}

		// Send response
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"user": map[string]interface{}{
					"id":   "123",
					"name": "John Doe",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("Failed to encode response: %v", err)
		}
	}))
	return server
}

func TestClient_Query(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"users": []map[string]interface{}{
					{"id": "1", "name": "Alice"},
					{"id": "2", "name": "Bob"},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("Failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient(server.URL)

	query := `query GetUsers($limit: Int) { users(limit: $limit) { id name } }`
	variables := map[string]interface{}{
		"limit": 10,
	}

	ctx := context.Background()
	response, err := client.Query(ctx, query, variables)
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}

	if !response.IsSuccessful() {
		t.Error("Expected successful response")
	}

	parser := response.GetParser()
	users, err := parser.GetDataAtPath("users")
	if err != nil {
		t.Fatalf("Failed to extract users: %v", err)
	}

	if len(users.Array()) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users.Array()))
	}
}

func TestClient_Mutate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"createUser": map[string]interface{}{
					"id":   "123",
					"name": "New User",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("Failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient(server.URL)

	mutation := `mutation CreateUser($input: CreateUserInput!) { createUser(input: $input) { id name } }`
	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"name": "New User",
		},
	}

	ctx := context.Background()
	response, err := client.Mutate(ctx, mutation, variables)
	if err != nil {
		t.Fatalf("Mutate failed: %v", err)
	}

	if !response.IsSuccessful() {
		t.Error("Expected successful response")
	}

	parser := response.GetParser()
	userIDResult, err := parser.GetDataAtPath("createUser.id")
	if err != nil {
		t.Fatalf("Failed to extract user ID: %v", err)
	}

	if userIDResult.String() != "123" {
		t.Errorf("Expected user ID '123', got %s", userIDResult.String())
	}
}

func TestClient_ExecuteWithBuilder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"user": map[string]interface{}{
					"id":    "456",
					"email": "test@example.com",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("Failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient(server.URL)

	builder := NewRequestBuilder().
		WithQuery(`query GetUser($id: ID!) { user(id: $id) { id email } }`).
		WithVariable("id", "456")

	ctx := context.Background()
	response, err := client.ExecuteWithBuilder(ctx, builder)
	if err != nil {
		t.Fatalf("ExecuteWithBuilder failed: %v", err)
	}

	if !response.IsSuccessful() {
		t.Error("Expected successful response")
	}

	parser := response.GetParser()
	email, err := parser.GetDataAtPath("user.email")
	if err != nil {
		t.Fatalf("Failed to extract email: %v", err)
	}

	if email.String() != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got %s", email.String())
	}
}

func TestClient_ErrorHandling(t *testing.T) {
	// Test configuration error
	client := NewClient("https://api.example.com/graphql")

	req := Request{
		Query: "", // Empty query should cause configuration error
	}

	ctx := context.Background()
	_, err := client.Execute(ctx, req)
	if err == nil {
		t.Error("Expected configuration error for empty query")
	}

	clientErr := &ClientError{}
	if errors.As(err, &clientErr) {
		if clientErr.Type != ErrorTypeConfiguration {
			t.Errorf("Expected configuration error, got %s", clientErr.Type)
		}
	} else {
		t.Error("Expected ClientError type")
	}
}

func TestClient_NetworkError(t *testing.T) {
	// Test network error with invalid endpoint
	client := NewClient("http://invalid-endpoint-that-does-not-exist.com/graphql")

	req := Request{
		Query: `query { test }`,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := client.Execute(ctx, req)
	if err == nil {
		t.Error("Expected network error for invalid endpoint")
	}

	clientErr := &ClientError{}
	if errors.As(err, &clientErr) {
		if clientErr.Type != ErrorTypeNetwork {
			t.Errorf("Expected network error, got %s", clientErr.Type)
		}
	} else {
		t.Error("Expected ClientError type")
	}
}
