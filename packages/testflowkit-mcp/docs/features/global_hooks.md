---
title: Global Hooks
description: Setup and teardown operations across your test suite
---

# Global Hooks

Run setup and teardown once per test run, before or after all scenarios. Use them to authenticate, seed data, or clean up — and to share values across scenarios via global variables.

## Tags

| Tag | When it runs |
|-----|--------------|
| `@BeforeAll` | Once, before any test scenario |
| `@AfterAll` | Once, after all test scenarios (even if tests fail) |

Place the tag on a feature file in your `features/` directory:

```gherkin
@BeforeAll
Feature: Test Setup

  Scenario: Authenticate
    Given I prepare a request to "my_api.login"
    And I set the request body to:
      """
      { "email": "test@example.com", "password": "testpass123" }
      """
    When I send the request
    And I save the response path "token" as global variable "auth_token"
    Then the response status code should be 200
```

```gherkin
@AfterAll
Feature: Test Cleanup

  Scenario: Delete test user
    Given I prepare a request to "my_api.delete_user"
    And I set the following path parameters:
      | id | {{test_user_id}} |
    And I set the header "Authorization" to "Bearer {{auth_token}}"
    When I send the request
    Then the response status code should be 204
```

## Global variables

Store values that all scenarios can read:

```gherkin
# From an API response
And I save the response path "token" as global variable "auth_token"

# Direct value
And I store the value "https://api.example.com" into global variable "api_base_url"
```

Use them with the same `{{variable_name}}` syntax:

```gherkin
And I set the header "Authorization" to "Bearer {{auth_token}}"
```

## Scope

| Variable | Scope | Set with |
|----------|-------|----------|
| Scenario | Current scenario | `I store the value ... into "name" variable` |
| Global | Entire test run | `... into global variable "name"` or `I save the response path ... as global variable "name"` |

## Execution order

1. `@BeforeAll` scenarios
2. All regular test scenarios
3. `@AfterAll` scenarios

::alert{type="warning"}
If a `@BeforeAll` hook fails, no test scenarios run. Keep hooks focused and resilient.
::

## Next Steps

- [Variables](./variables.md) — Scenario and environment variables
- [API Testing](/docs/guides/api-testing) — API steps for hooks
