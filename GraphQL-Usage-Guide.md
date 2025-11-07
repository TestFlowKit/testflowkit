# GraphQL Testing with TestFlowKit

This comprehensive guide covers GraphQL testing capabilities in TestFlowKit, including configuration, step definitions, variable handling, and advanced testing patterns.

## Table of Contents

1. [Overview](#overview)
2. [Configuration](#configuration)
3. [Step Definitions](#step-definitions)
4. [Variable Management](#variable-management)
5. [Testing Patterns](#testing-patterns)
6. [Schema Validation](#schema-validation)
7. [Error Handling](#error-handling)
8. [Best Practices](#best-practices)
9. [Examples](#examples)

## Overview

TestFlowKit provides comprehensive GraphQL testing capabilities that allow you to:

- Execute GraphQL queries, mutations, and subscriptions
- Handle complex variable types including arrays and objects
- Validate GraphQL responses and errors
- Store and reuse data from GraphQL responses
- Perform schema introspection and validation
- Test GraphQL APIs with the same ease as REST APIs

## Configuration

### Basic GraphQL Configuration

Configure GraphQL endpoints and operations in your `config.yml`:

```yaml
backend:
  graphql:
    endpoint: "/graphql"                    # GraphQL endpoint path
    introspection_enabled: true             # Enable schema introspection
    default_headers:                        # Default headers for all requests
      Content-Type: "application/json"
      Accept: "application/json"
    operations:                             # Define your GraphQL operations
      get_user_profile:
        type: "query"
        operation: |
          query GetUserProfile($userId: ID!) {
            user(id: $userId) {
              id
              name
              email
              profile {
                avatar
                bio
              }
            }
          }
        description: "Fetch user profile with nested data"
```

### Environment-Specific Configuration

Use different GraphQL endpoints for different environments:

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
    endpoint: "/graphql"  # Combined with api_base_url
```

### Authentication Configuration

Include authentication headers:

```yaml
backend:
  default_headers:
    Authorization: "Bearer ${API_TOKEN}"
  
  graphql:
    default_headers:
      Content-Type: "application/json"
      X-API-Key: "${GRAPHQL_API_KEY}"
```

## Step Definitions

### Request Preparation

#### Prepare GraphQL Request

```gherkin
Given I prepare a GraphQL request for the "get_user_profile" operation
```

Prepares a GraphQL request using a predefined operation from your configuration.

#### Set GraphQL Variables (Table Format)

```gherkin
And I set the following GraphQL variables:
  | userId    | 123                           |
  | tags      | ["frontend", "testing"]       |
  | filters   | {"status": "active"}          |
  | isActive  | true                          |
```

Sets multiple GraphQL variables using a table format. Supports all variable types.

#### Set Individual GraphQL Variable

```gherkin
And I set the GraphQL variable "userId" to "123"
And I set the GraphQL variable "tags" to array ["frontend", "backend", "testing"]
```

Sets individual GraphQL variables, with special support for array variables.

#### Set GraphQL Headers

```gherkin
And I set the following GraphQL headers:
  | Authorization | Bearer token123 |
  | X-Client-ID   | test-client     |
```

Sets custom headers for the GraphQL request.

### Request Execution

#### Send GraphQL Request

```gherkin
When I send the GraphQL request
```

Executes the prepared GraphQL request and stores the response.

### Response Validation

#### Validate No Errors

```gherkin
Then the GraphQL response should not contain errors
```

Validates that the GraphQL response doesn't contain any errors.

#### Validate Response Data

```gherkin
Then the GraphQL response should contain data at path "user.name"
And the GraphQL response should contain data at path "user.profile.avatar"
```

Validates that specific data paths exist in the GraphQL response.

#### Validate Errors Present

```gherkin
Then the GraphQL response should contain errors
```

Validates that the GraphQL response contains errors (useful for negative testing).

### Data Storage

#### Store GraphQL Data

```gherkin
And I store the GraphQL data at path "user.id" into "userId" variable
And I store the GraphQL data at path "user.name" into "userName" variable
```

Extracts data from GraphQL responses and stores it in variables for later use.

#### Store GraphQL Array Data

```gherkin
And I store the GraphQL array at path "user.tags" into "userTags" variable
And I store the GraphQL array at path "posts" into "userPosts" variable
```

Extracts array data from GraphQL responses and stores it in variables.

## Variable Management

### Variable Types

TestFlowKit supports comprehensive variable types for GraphQL operations:

#### Primitive Types

```gherkin
And I set the following GraphQL variables:
  | name      | "John Doe"    |  # String
  | age       | 30            |  # Number
  | isActive  | true          |  # Boolean
  | userId    | "user-123"    |  # ID (string)
```

#### Array Types

```gherkin
# String arrays
And I set the GraphQL variable "tags" to array ["frontend", "backend", "testing"]

# Number arrays
And I set the GraphQL variable "scores" to array [95, 87, 92]

# Boolean arrays
And I set the GraphQL variable "flags" to array [true, false, true]

# Mixed arrays (in table format)
And I set the following GraphQL variables:
  | categories | ["tech", "tutorial", "beginner"] |
  | ratings    | [4.5, 3.8, 5.0]                  |
```

#### Object Types

```gherkin
And I set the following GraphQL variables:
  | input | {"name": "John Doe", "email": "john@example.com"} |
  | filters | {"status": "active", "verified": true} |
```

#### Complex Nested Types

```gherkin
And I set the following GraphQL variables:
  | createPostInput | {"title": "New Post", "content": "Content here", "tags": ["tech", "tutorial"], "author": {"id": "123", "name": "John"}} |
```

### Variable Parsing

TestFlowKit automatically parses variable values based on their format:

- **JSON Arrays**: `["item1", "item2"]` → Array
- **JSON Objects**: `{"key": "value"}` → Object
- **Booleans**: `true`/`false` → Boolean
- **Numbers**: `123` or `45.67` → Number
- **Strings**: Everything else → String

## Testing Patterns

### Basic Query Testing

```gherkin
Scenario: Get user profile
  Given I prepare a GraphQL request for the "get_user_profile" operation
  And I set the following GraphQL variables:
    | userId | 123 |
  When I send the GraphQL request
  Then the GraphQL response should not contain errors
  And the GraphQL response should contain data at path "user.name"
  And the GraphQL response should contain data at path "user.email"
```

### Mutation Testing

```gherkin
Scenario: Create new user
  Given I prepare a GraphQL request for the "create_user" operation
  And I set the following GraphQL variables:
    | input | {"name": "Jane Doe", "email": "jane@example.com"} |
  When I send the GraphQL request
  Then the GraphQL response should not contain errors
  And the GraphQL response should contain data at path "createUser.user.id"
  And I store the GraphQL data at path "createUser.user.id" into "newUserId" variable
```

### Array Variable Testing

```gherkin
Scenario: Search users with multiple criteria
  Given I prepare a GraphQL request for the "search_users" operation
  And I set the following GraphQL variables:
    | tags     | ["frontend", "testing", "automation"] |
    | statuses | ["active", "verified"]               |
    | limit    | 10                                   |
  When I send the GraphQL request
  Then the GraphQL response should not contain errors
  And the GraphQL response should contain data at path "users"
  And I store the GraphQL array at path "users" into "foundUsers" variable
```

### Error Handling Testing

```gherkin
Scenario: Handle invalid user ID
  Given I prepare a GraphQL request for the "get_user_profile" operation
  And I set the following GraphQL variables:
    | userId | invalid_id |
  When I send the GraphQL request
  Then the GraphQL response should contain errors
```

### Data Flow Testing

```gherkin
Scenario: Create user and then fetch profile
  # Create user
  Given I prepare a GraphQL request for the "create_user" operation
  And I set the following GraphQL variables:
    | input | {"name": "Test User", "email": "test@example.com"} |
  When I send the GraphQL request
  Then the GraphQL response should not contain errors
  And I store the GraphQL data at path "createUser.user.id" into "createdUserId" variable
  
  # Fetch created user
  Given I prepare a GraphQL request for the "get_user_profile" operation
  And I set the following GraphQL variables:
    | userId | {{createdUserId}} |
  When I send the GraphQL request
  Then the GraphQL response should not contain errors
  And the GraphQL response should contain data at path "user.name"
```

## Schema Validation

### Enable Schema Introspection

```yaml
backend:
  graphql:
    introspection_enabled: true
```

When enabled, TestFlowKit will:

1. Automatically fetch the GraphQL schema from the endpoint
2. Validate operations against the schema before execution
3. Provide detailed error messages for invalid operations
4. Cache the schema for performance

### Schema Validation Benefits

- **Early Error Detection**: Catch invalid operations before execution
- **Type Safety**: Ensure variables match expected types
- **Field Validation**: Verify that requested fields exist
- **Better Error Messages**: Get specific feedback about schema mismatches

## Error Handling

### GraphQL Error Types

TestFlowKit handles different types of GraphQL errors:

1. **Network Errors**: HTTP-level connection issues
2. **GraphQL Errors**: Application-level GraphQL errors
3. **Validation Errors**: Schema validation failures
4. **Configuration Errors**: Missing or invalid operation definitions

### Error Testing Patterns

```gherkin
# Test for specific error conditions
Scenario: Test authentication error
  Given I prepare a GraphQL request for the "get_user_profile" operation
  And I set the following GraphQL headers:
    | Authorization | Bearer invalid_token |
  And I set the following GraphQL variables:
    | userId | 123 |
  When I send the GraphQL request
  Then the GraphQL response should contain errors

# Test for validation errors
Scenario: Test invalid input
  Given I prepare a GraphQL request for the "create_user" operation
  And I set the following GraphQL variables:
    | input | {"email": "invalid-email"} |
  When I send the GraphQL request
  Then the GraphQL response should contain errors
```

## Best Practices

### 1. Operation Organization

```yaml
# Group related operations logically
backend:
  graphql:
    operations:
      # User operations
      get_user_profile:
        type: "query"
        # ...
      
      create_user:
        type: "mutation"
        # ...
      
      # Post operations
      get_posts:
        type: "query"
        # ...
      
      create_post:
        type: "mutation"
        # ...
```

### 2. Use Descriptive Names

```yaml
operations:
  get_user_with_posts_and_comments:  # Descriptive name
    type: "query"
    description: "Fetch user profile with associated posts and comments"
    # ...
```

### 3. Include Error Fields

```graphql
mutation CreateUser($input: CreateUserInput!) {
  createUser(input: $input) {
    user {
      id
      name
      email
    }
    errors {
      field
      message
      code
    }
  }
}
```

### 4. Use Variables for Reusability

```gherkin
# Store commonly used values
Given I store the "active" into "userStatus" variable
And I store the "verified" into "userVerification" variable

# Use variables in GraphQL requests
And I set the following GraphQL variables:
  | statuses | ["{{userStatus}}", "{{userVerification}}"] |
```

### 5. Test Both Success and Error Cases

```gherkin
# Success case
Scenario: Valid user creation
  Given I prepare a GraphQL request for the "create_user" operation
  # ... valid input
  Then the GraphQL response should not contain errors

# Error case
Scenario: Invalid user creation
  Given I prepare a GraphQL request for the "create_user" operation
  # ... invalid input
  Then the GraphQL response should contain errors
```

## Examples

### Complete User Management Test

```gherkin
Feature: User Management GraphQL API
  As a developer
  I want to test user management operations
  So that I can ensure the GraphQL API works correctly

  Background:
    Given the API is available

  Scenario: Create and retrieve user with tags
    # Create user with tags
    Given I prepare a GraphQL request for the "create_user" operation
    And I set the following GraphQL variables:
      | input | {"name": "John Doe", "email": "john@example.com", "tags": ["developer", "frontend"]} |
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    And the GraphQL response should contain data at path "createUser.user.id"
    And I store the GraphQL data at path "createUser.user.id" into "userId" variable
    
    # Retrieve created user
    Given I prepare a GraphQL request for the "get_user_profile" operation
    And I set the following GraphQL variables:
      | userId | {{userId}} |
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    And the GraphQL response should contain data at path "user.name"
    And the GraphQL response should contain data at path "user.tags"
    And I store the GraphQL array at path "user.tags" into "userTags" variable
    
    # Update user tags
    Given I prepare a GraphQL request for the "update_user_tags" operation
    And I set the GraphQL variable "userId" to "{{userId}}"
    And I set the GraphQL variable "tags" to array ["developer", "frontend", "testing"]
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    And the GraphQL response should contain data at path "updateUser.tags"

  Scenario: Search users by multiple criteria
    Given I prepare a GraphQL request for the "search_users" operation
    And I set the following GraphQL variables:
      | tags     | ["developer", "frontend"] |
      | statuses | ["active", "verified"]    |
      | limit    | 5                         |
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    And the GraphQL response should contain data at path "users"
    And I store the GraphQL array at path "users" into "searchResults" variable

  Scenario: Handle authentication errors
    Given I prepare a GraphQL request for the "get_user_profile" operation
    And I set the following GraphQL headers:
      | Authorization | Bearer invalid_token |
    And I set the following GraphQL variables:
      | userId | 123 |
    When I send the GraphQL request
    Then the GraphQL response should contain errors
```

### Complex Data Flow Test

```gherkin
Feature: Blog Post Management
  As a content creator
  I want to manage blog posts through GraphQL
  So that I can create and organize content

  Scenario: Complete blog post workflow
    # Create author
    Given I prepare a GraphQL request for the "create_user" operation
    And I set the following GraphQL variables:
      | input | {"name": "Jane Author", "email": "jane@example.com", "role": "AUTHOR"} |
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    And I store the GraphQL data at path "createUser.user.id" into "authorId" variable
    
    # Create blog post
    Given I prepare a GraphQL request for the "create_post" operation
    And I set the following GraphQL variables:
      | input | {"title": "GraphQL Testing Guide", "content": "Comprehensive guide to GraphQL testing", "authorId": "{{authorId}}", "tags": ["graphql", "testing", "tutorial"]} |
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    And I store the GraphQL data at path "createPost.post.id" into "postId" variable
    
    # Add comments to post
    Given I prepare a GraphQL request for the "add_comment" operation
    And I set the following GraphQL variables:
      | input | {"postId": "{{postId}}", "content": "Great tutorial!", "authorName": "Reader One"} |
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    
    # Fetch post with comments
    Given I prepare a GraphQL request for the "get_post_with_comments" operation
    And I set the following GraphQL variables:
      | postId | {{postId}} |
    When I send the GraphQL request
    Then the GraphQL response should not contain errors
    And the GraphQL response should contain data at path "post.title"
    And the GraphQL response should contain data at path "post.comments"
    And I store the GraphQL array at path "post.comments" into "postComments" variable
```

This guide provides comprehensive coverage of GraphQL testing capabilities in TestFlowKit. For more specific examples and advanced patterns, refer to the example feature files in the `examples/` directory.