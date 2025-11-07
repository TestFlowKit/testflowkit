# GraphQL Configuration Guide

This guide explains how to configure GraphQL operations in TestFlowKit for automated testing.

## Configuration Structure

GraphQL configuration is defined under the `backend.graphql` section in your `config.yml` file:

```yaml
backend:
  graphql:
    endpoint: "/graphql"                    # GraphQL endpoint path
    introspection_enabled: true             # Enable schema introspection
    default_headers:                        # Default headers for all GraphQL requests
      Content-Type: "application/json"
      Accept: "application/json"
    operations:                             # Define your GraphQL operations
      operation_name:
        type: "query"                       # Operation type: query, mutation, or subscription
        operation: |                        # GraphQL operation string
          query OperationName($var: Type!) {
            field(arg: $var) {
              subfield
            }
          }
        description: "Description of the operation"
```

## Configuration Options

### GraphQL Section

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `endpoint` | string | Yes | The GraphQL endpoint path (e.g., "/graphql") |
| `introspection_enabled` | boolean | No | Enable GraphQL schema introspection for validation (default: false) |
| `default_headers` | map | No | Default HTTP headers to include with all GraphQL requests |
| `operations` | map | Yes | Map of named GraphQL operations |

### Operation Definition

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | string | Yes | Operation type: "query", "mutation", or "subscription" |
| `operation` | string | Yes | The GraphQL operation string (query/mutation/subscription) |
| `description` | string | Yes | Human-readable description of the operation |

## Variable Types Support

TestFlowKit supports various variable types in GraphQL operations:

### Primitive Types
- **String**: `"hello world"`
- **Number**: `123` or `45.67`
- **Boolean**: `true` or `false`
- **ID**: `"user-123"`

### Complex Types
- **Arrays**: `["item1", "item2", "item3"]` or `[1, 2, 3]`
- **Objects**: `{"key": "value", "nested": {"field": true}}`
- **Mixed Arrays**: `[{"id": 1, "name": "John"}, {"id": 2, "name": "Jane"}]`

## Example Configurations

### Basic Query Example

```yaml
backend:
  graphql:
    endpoint: "/graphql"
    operations:
      get_user:
        type: "query"
        operation: |
          query GetUser($userId: ID!) {
            user(id: $userId) {
              id
              name
              email
            }
          }
        description: "Fetch user by ID"
```

### Query with Array Variables

```yaml
backend:
  graphql:
    operations:
      search_users:
        type: "query"
        operation: |
          query SearchUsers($tags: [String!]!, $statuses: [UserStatus!]!) {
            users(tags: $tags, statuses: $statuses) {
              id
              name
              email
              tags
              status
            }
          }
        description: "Search users by tags and statuses"
```

### Mutation with Complex Input

```yaml
backend:
  graphql:
    operations:
      create_post:
        type: "mutation"
        operation: |
          mutation CreatePost($input: CreatePostInput!) {
            createPost(input: $input) {
              post {
                id
                title
                content
                author {
                  id
                  name
                }
                tags
                createdAt
              }
              errors {
                field
                message
              }
            }
          }
        description: "Create a new post with complex input"
```

### Using Fragments

```yaml
backend:
  graphql:
    operations:
      get_user_with_posts:
        type: "query"
        operation: |
          query GetUserWithPosts($userId: ID!) {
            user(id: $userId) {
              ...UserDetails
              posts {
                ...PostSummary
              }
            }
          }
          
          fragment UserDetails on User {
            id
            name
            email
            profile {
              avatar
              bio
            }
          }
          
          fragment PostSummary on Post {
            id
            title
            publishedAt
            tags
          }
        description: "Get user with posts using fragments"
```

## Environment-Specific Configuration

You can use different GraphQL endpoints for different environments:

```yaml
environments:
  development:
    api_base_url: "http://localhost:4000"
  
  staging:
    api_base_url: "https://api-staging.example.com"
  
  production:
    api_base_url: "https://api.example.com"

backend:
  graphql:
    endpoint: "/graphql"  # Will be combined with api_base_url
```

## Authentication

Include authentication headers in your GraphQL requests:

```yaml
backend:
  default_headers:
    Authorization: "Bearer ${API_TOKEN}"  # Use environment variables
  
  graphql:
    default_headers:
      Content-Type: "application/json"
      X-API-Key: "${GRAPHQL_API_KEY}"     # GraphQL-specific headers
```

## Schema Introspection

Enable schema introspection for operation validation:

```yaml
backend:
  graphql:
    introspection_enabled: true  # Enables schema validation
```

When enabled, TestFlowKit will:
- Fetch the GraphQL schema from the endpoint
- Validate operations against the schema
- Provide detailed error messages for invalid operations
- Cache the schema for performance

## Best Practices

1. **Use Descriptive Names**: Choose clear, descriptive names for your operations
2. **Include Descriptions**: Always provide meaningful descriptions for operations
3. **Group Related Operations**: Organize operations logically in your configuration
4. **Use Variables**: Parameterize your operations with variables for reusability
5. **Handle Errors**: Include error fields in your mutations for proper error handling
6. **Use Fragments**: Leverage GraphQL fragments to avoid duplication
7. **Environment Variables**: Use environment variables for sensitive data like API keys

## Common Patterns

### Pagination
```graphql
query GetPosts($first: Int, $after: String) {
  posts(first: $first, after: $after) {
    edges {
      node {
        id
        title
      }
      cursor
    }
    pageInfo {
      hasNextPage
      endCursor
    }
  }
}
```

### Error Handling
```graphql
mutation CreateUser($input: CreateUserInput!) {
  createUser(input: $input) {
    user {
      id
      name
    }
    errors {
      field
      message
      code
    }
  }
}
```

### Nested Relationships
```graphql
query GetUserWithPosts($userId: ID!) {
  user(id: $userId) {
    id
    name
    posts {
      id
      title
      comments {
        id
        content
        author {
          name
        }
      }
    }
  }
}
```