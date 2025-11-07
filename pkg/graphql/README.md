# GraphQL Client Library

A comprehensive GraphQL client library for Go applications with advanced error handling, response parsing, and request building capabilities.

## Features

- **HTTP Client**: Full-featured GraphQL HTTP client with configurable options
- **Request Builder**: Fluent API for building GraphQL requests
- **Response Parser**: Advanced JSON path-based data extraction
- **Error Handling**: Comprehensive error classification and severity analysis
- **Type Safety**: Strong typing for GraphQL requests and responses
- **Context Support**: Full context.Context support for cancellation and timeouts

## Installation

```bash
go get your-module/pkg/graphql
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "your-module/pkg/graphql"
)

func main() {
    // Create a new GraphQL client
    client := graphql.NewClient(
        "https://api.example.com/graphql",
        graphql.WithTimeout(10*time.Second),
        graphql.WithHeaders(map[string]string{
            "Authorization": "Bearer your-token",
        }),
    )
    
    // Execute a simple query
    query := `
        query GetUser($id: ID!) {
            user(id: $id) {
                id
                name
                email
            }
        }
    `
    
    variables := map[string]interface{}{
        "id": "123",
    }
    
    ctx := context.Background()
    response, err := client.Query(ctx, query, variables)
    if err != nil {
        log.Fatal(err)
    }
    
    // Parse the response
    parser := response.GetParser()
    
    if response.HasErrors() {
        fmt.Printf("GraphQL errors: %s\n", response.GetErrorsAsString())
        return
    }
    
    // Extract data
    userName, err := parser.GetDataAsString("user.name")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("User name: %s\n", userName)
}
```

## Advanced Usage

### Using Request Builder

```go
builder := graphql.NewRequestBuilder().
    WithQuery(`
        mutation CreateUser($input: CreateUserInput!) {
            createUser(input: $input) {
                id
                name
                email
            }
        }
    `).
    WithVariable("input", map[string]interface{}{
        "name":  "John Doe",
        "email": "john@example.com",
    })

response, err := client.ExecuteWithBuilder(ctx, builder)
```

### Error Handling

```go
response, err := client.Query(ctx, query, variables)
if err != nil {
    // Handle client errors (network, configuration, etc.)
    if clientErr, ok := err.(*graphql.ClientError); ok {
        fmt.Printf("Client error [%s]: %s\n", clientErr.Type, clientErr.Message)
        fmt.Printf("Details: %+v\n", clientErr.Details)
    }
    return
}

if response.HasErrors() {
    // Handle GraphQL errors
    summary := response.GetErrorSummary()
    fmt.Printf("Total errors: %d\n", summary.TotalErrors)
    
    // Get errors by classification
    authErrors := response.GetErrorsByClassification("AUTH_ERROR")
    if len(authErrors) > 0 {
        fmt.Println("Authentication errors found")
    }
    
    // Check for critical errors
    if response.HasCriticalErrors() {
        fmt.Println("Critical errors detected!")
    }
}
```

### Advanced Response Parsing

```go
parser := response.GetParser()

// Extract different data types
userID, err := parser.GetDataAsInt("user.id")
isActive, err := parser.GetDataAsBool("user.active")
tags, err := parser.GetDataAsArray("user.tags")

// Validate data paths and types
err = parser.ValidateDataPath("user.email", "string")
if err != nil {
    fmt.Printf("Validation error: %s\n", err)
}

// Extract complex data for variables
userData, err := parser.ExtractVariableData("user")
if err != nil {
    log.Fatal(err)
}

// Get response metadata
metadata := parser.GetResponseMetadata()
fmt.Printf("Response metadata: %+v\n", metadata)
```

### Client Configuration

```go
// Custom HTTP client
httpClient := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns: 100,
    },
}

client := graphql.NewClient(
    "https://api.example.com/graphql",
    graphql.WithHTTPClient(httpClient),
    graphql.WithHeaders(map[string]string{
        "User-Agent": "MyApp/1.0",
        "X-API-Key":  "your-api-key",
    }),
)

// Add headers dynamically
client.SetHeaders(map[string]string{
    "Authorization": "Bearer new-token",
})
```

## Error Types

The library provides comprehensive error handling with the following error types:

- **Configuration**: Client configuration errors
- **Network**: HTTP and network-related errors  
- **GraphQL**: GraphQL execution errors
- **Syntax**: GraphQL syntax errors
- **Schema**: Schema validation errors

## Error Classifications

GraphQL errors are automatically classified:

- **AUTH_ERROR**: Authentication/authorization errors
- **VALIDATION_ERROR**: Input validation errors
- **FIELD_ERROR**: Field resolution errors
- **SYNTAX_ERROR**: Query syntax errors
- **INTERNAL_ERROR**: Server internal errors
- **GRAPHQL_ERROR**: Generic GraphQL errors

## Error Severity Levels

- **CRITICAL**: Errors that prevent execution
- **HIGH**: Authentication/authorization errors
- **MEDIUM**: Validation and field errors
- **LOW**: Other errors

## API Reference

### Client

- `NewClient(endpoint, ...options)` - Create a new GraphQL client
- `Execute(ctx, request)` - Execute a GraphQL request
- `Query(ctx, query, variables)` - Execute a query
- `Mutate(ctx, mutation, variables)` - Execute a mutation
- `SetHeaders(headers map[string]string)` - Set request headers
- `GetEndpoint()` - Get the GraphQL endpoint URL

### RequestBuilder

- `NewRequestBuilder()` - Create a new request builder
- `WithQuery(query)` - Set the GraphQL query
- `WithVariable(name, value)` - Set a single variable
- `WithVariables(variables)` - Set multiple variables
- `Build()` - Build the final request

### ResponseParser

- `GetDataAtPath(path)` - Extract data at JSON path
- `GetDataAsString(path)` - Extract string data
- `GetDataAsInt(path)` - Extract integer data
- `GetDataAsFloat(path)` - Extract float data
- `GetDataAsBool(path)` - Extract boolean data
- `GetDataAsArray(path)` - Extract array data
- `ValidateDataPath(path, type)` - Validate path and type
- `ExtractVariableData(path)` - Extract data for variables

### Response

- `HasErrors()` - Check if response has errors
- `HasData()` - Check if response has data
- `IsSuccessful()` - Check if response is successful
- `GetErrorMessages()` - Get all error messages
- `GetErrorSummary()` - Get error summary
- `GetErrorsByClassification(class)` - Filter errors by classification
- `HasCriticalErrors()` - Check for critical errors

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This library is part of the TestFlowKit project.