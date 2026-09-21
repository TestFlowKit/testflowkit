---
title: Add a GraphQL Operation
description: Declare a new GraphQL query or mutation and call it from a scenario
navigation:
  title: Add a GraphQL Operation
---

## Goal

Test a new GraphQL query or mutation without writing any code.

## Prerequisites

- A `graphql` API definition already exists in `testflowkit.yml`, or you are creating one.

## Configuration

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
        create_post:
          type: mutation
          operation: "graphql/queries/create_post.graphql"
```

`operation` can point to a `.graphql` file or contain the query/mutation inline.

## Scenario

```gherkin
Scenario: Fetch a user
  Given I prepare a request to "my_graphql.get_user"
  And I set the following GraphQL variables:
    | id | 1 |
  When I send the request
  Then the GraphQL response should not have errors
  And the response should have field "user.username"
```

## Verify

Run `tkit run` and check the scenario reports no GraphQL errors and the expected fields are present.

## Common issues

- **GraphQL errors in response** — assert with `the GraphQL response should not have errors` before checking fields.
- **Timeout** — precedence is operation → API → `apis.default_timeout` → 30s fallback.

## See also

- [API Testing](/docs/guides/api-testing) — full REST/GraphQL guide
- [testflowkit.yml](/docs/config/overview) — complete `apis` schema
- [Add a REST Endpoint](/docs/how-to/add-a-rest-endpoint)
