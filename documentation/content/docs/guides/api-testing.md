---
title: API Testing
description: How an API scenario works and where to find each recipe
navigation:
  title: API Testing
---

Define APIs in `testflowkit.yml`, then call them with `I prepare a request to "api_name.endpoint_name"`. The syntax is the same for REST and GraphQL.

## Typical flow

1. **Prepare**: `I prepare a request to "api_name.endpoint_name"`. This resets the current request state (headers, body, path and query parameters, GraphQL variables, any previous response), then loads the endpoint defaults from `testflowkit.yml`.
2. **Configure**: set the body, path or query parameters, headers (merged with `default_headers`).
3. **Send**: `I send the request`.
4. **Assert**: status code, response fields (GJSON for JSON, XPath for XML), headers.
5. **Extract**: `I store the response path "data.id" from the response into "id" variable`.

Use `{{variable}}` anywhere in a step value.

```gherkin
Scenario: Create and verify a post
  Given I prepare a request to "jsonplaceholder.create_post"
  And I set the request body to:
    """
    { "title": "Test Post", "userId": 1 }
    """
  When I send the request
  Then the response status code should be 201
  And the response field "title" should be "Test Post"
```

## Find your recipe

| I want to... | Go to |
|---|---|
| Declare and call a REST endpoint | [Add a REST endpoint](/docs/how-to/add-a-rest-endpoint) |
| Declare and call a GraphQL query or mutation | [Add a GraphQL operation](/docs/how-to/add-a-graphql-operation) |
| Send a JSON body, a body from a file, or upload a file | [Send a JSON body or file](/docs/how-to/send-a-json-body-or-file) |
| Have every request authenticated | [Configure authentication](/docs/how-to/configure-authentication) |
| Log in first and reuse the token | [Authenticate before tests](/docs/how-to/authenticate-before-tests) |
| Reuse a value from one response in the next request | [Variables](/docs/patterns/variables) |
| Generate unique payload data | [Random Data](/docs/patterns/random-data) |
| Look up an exact step | [Step Catalog](/sentences) |

Timeouts, `default_headers` and the full `apis` block are documented in [testflowkit.yml](/docs/config/overview).
