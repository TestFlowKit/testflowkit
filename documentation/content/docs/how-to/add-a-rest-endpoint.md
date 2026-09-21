---
title: Add a REST Endpoint
description: Declare a new REST endpoint in testflowkit.yml and call it from a scenario
navigation:
  title: Add a REST Endpoint
---

# Add a REST Endpoint

## Goal

Test a new REST endpoint of your backend without writing any code.

## Prerequisites

- An `apis` definition already exists in `testflowkit.yml`, or you are creating one.

## Configuration

Add the endpoint under an existing (or new) `rest` API definition:

```yaml
apis:
  definitions:
    my_api:
      type: rest
      base_url: "{{ env.api_base_url }}"
      default_headers:
        Content-Type: "application/json"
      endpoints:
        get_users:
          method: GET
          path: "/users"
        create_user:
          method: POST
          path: "/users"
        get_user_by_id:
          method: GET
          path: "/users/{id}"
```

Path parameters use the `{name}` syntax and are filled with `I set the path parameter "name" to "value"`.

## Scenario

```gherkin
Scenario: Create a user
  Given I prepare a request to "my_api.create_user"
  And I set the request body to:
    """
    { "name": "Jane Doe", "email": "jane@example.com" }
    """
  When I send the request
  Then the response status code should be 201
  And the response should have field "id"
```

## Verify

Run `tkit run` and check the scenario passes with the expected status code and response fields.

## Common issues

- **404 / wrong path** — check `base_url` + `path` concatenation and path parameter names.
- **Timeout** — precedence is endpoint → API → `apis.default_timeout` → 30s fallback.

## See also

- [API Testing](/docs/guides/api-testing) — full REST/GraphQL guide
- [testflowkit.yml](/docs/config/overview) — complete `apis` schema
- [Add a GraphQL Operation](/docs/how-to/add-a-graphql-operation)
