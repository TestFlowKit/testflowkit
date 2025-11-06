/*
Package graphql provides a comprehensive GraphQL client library for Go applications.

This package offers a full-featured GraphQL HTTP client with advanced error handling,
response parsing, and request building capabilities. It's designed to be both easy to
use for simple cases and powerful enough for complex GraphQL operations.

# Key Features

- HTTP Client: Full-featured GraphQL HTTP client with configurable options
- Request Builder: Fluent API for building GraphQL requests
- Response Parser: Advanced JSON path-based data extraction using gjson
- Error Handling: Comprehensive error classification and severity analysis
- Type Safety: Strong typing for GraphQL requests and responses
- Context Support: Full context.Context support for cancellation and timeouts

# Basic Usage

	client := graphql.NewClient("https://api.example.com/graphql")

	query := `query GetUser($id: ID!) { user(id: $id) { id name email } }`
	variables := map[string]interface{}{"id": "123"}

	response, err := client.Query(context.Background(), query, variables)
	if err != nil {
		log.Fatal(err)
	}

	parser := response.GetParser()
	name, _ := parser.GetDataAsString("user.name")

# Advanced Configuration

	client := graphql.NewClient(
		"https://api.example.com/graphql",
		graphql.WithTimeout(10*time.Second),
		graphql.WithHeaders(map[string]string{
			"Authorization": "Bearer token",
		}),
	)

# Error Handling

The library provides comprehensive error handling with automatic classification:

	response, err := client.Query(ctx, query, variables)
	if err != nil {
		// Handle client errors (network, configuration, etc.)
		if clientErr, ok := err.(*graphql.ClientError); ok {
			fmt.Printf("Error [%s]: %s\n", clientErr.Type, clientErr.Message)
		}
		return
	}

	if response.HasErrors() {
		// Handle GraphQL errors
		authErrors := response.GetErrorsByClassification("AUTH_ERROR")
		if len(authErrors) > 0 {
			fmt.Println("Authentication required")
		}
	}

# Response Parsing

The response parser provides powerful data extraction capabilities:

	parser := response.GetParser()

	// Extract different data types
	name, _ := parser.GetDataAsString("user.name")
	age, _ := parser.GetDataAsInt("user.age")
	active, _ := parser.GetDataAsBool("user.active")
	tags, _ := parser.GetDataAsArray("user.tags")

	// Validate data paths and types
	err := parser.ValidateDataPath("user.email", "string")

# Request Building

Use the fluent request builder for complex requests:

	builder := graphql.NewRequestBuilder().
		WithQuery(`mutation CreateUser($input: CreateUserInput!) {
			createUser(input: $input) { id name }
		}`).
		WithVariable("input", map[string]interface{}{
			"name": "John Doe",
			"email": "john@example.com",
		})

	response, err := client.ExecuteWithBuilder(ctx, builder)

This package is part of the TestFlowKit project but can be used independently
in any Go application that needs GraphQL client functionality.
*/
package graphql
