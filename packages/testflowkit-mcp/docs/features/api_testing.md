---
title: API Testing
description: REST API and GraphQL testing capabilities
---

# API Testing

Define APIs in `testflowkit.yml`, then call them with `I prepare a request to "api_name.endpoint_name"`. Same syntax for REST and GraphQL.

## REST configuration

```yaml
apis:
  default_timeout: 10000
  definitions:
    jsonplaceholder:
      type: rest
      base_url: "https://jsonplaceholder.typicode.com"
      default_headers:
        Content-Type: "application/json"
      endpoints:
        get_posts:
          method: GET
          path: "/posts"
        create_post:
          method: POST
          path: "/posts"
        get_post_by_id:
          method: GET
          path: "/posts/{id}"
```

Timeout precedence: endpoint → API → `apis.default_timeout` → 30s fallback.

## Typical flow

Most API scenarios follow the same pattern:

1. **Prepare** — `I prepare a request to "api_name.endpoint_name"`. This resets the current request state (headers, body, path/query parameters, GraphQL variables, and any previous response), then loads the selected endpoint with its defaults from `testflowkit.yml`.
2. **Configure** — set body, path/query params, headers (merged with `default_headers`)
3. **Send** — `I send the request`
4. **Assert** — status code, response fields (GJSON for JSON, XPath for XML), headers
5. **Extract** — `I store the response path "data.id" from the response into "id" variable`

Use `{{variable}}` anywhere in step values.

## REST example

```gherkin
Scenario: Create and verify a post
  Given I prepare a request to "jsonplaceholder.create_post"
  And I set the request body to:
    """
    {
      "title": "Test Post",
      "body": "Test body",
      "userId": 1
    }
    """
  When I send the request
  Then the response status code should be 201
  And the response should have field "id"
  And the response field "title" should be "Test Post"
```

## GraphQL

```yaml
apis:
  definitions:
    my_graphql:
      type: graphql
      endpoint: "{{ env.graphql_endpoint }}"
      operations:
        get_user:
          type: query
          operation: "graphql/queries/get_user.graphql"
```

```gherkin
Scenario: Fetch a user
  Given I prepare a request to "my_graphql.get_user"
  And I set the following GraphQL variables:
    | id | 1 |
  When I send the request
  Then the GraphQL response should not have errors
  And the response should have field "user.username"
```

Operations can also be defined inline in config — see [testflowkit.yml](/docs/config/overview).

## Authentication

Configure reusable auth with `security_schemes` and `security_ref` (bearer, basic, apikey, oauth2), or set a header per request:

```gherkin
And I set the header "Authorization" to "Bearer {{auth_token}}"
```

## Step catalog

For the full list of API sentences, browse the **[Step Definitions catalog](/sentences)** — searchable by keyword and category.

## Next Steps

- [Variables](./variables.md) — Store and reuse response data
- [Random Data](./random_data.md) — Dynamic request payloads
- [Global Hooks](./global_hooks.md) — Auth and data setup before tests
