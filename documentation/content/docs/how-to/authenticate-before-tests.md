---
title: Authenticate Before Tests
description: Log in once, extract the token, and reuse it in every scenario
navigation:
  title: Authenticate Before Tests
---

# Authenticate Before Tests

## Goal

Log in through your API's login endpoint once before the test run, then reuse the returned token in every scenario. Use this when credentials are not static (see [Configure Authentication](/docs/how-to/configure-authentication) for fixed tokens/OAuth2 client credentials).

## Prerequisites

- A login endpoint declared under `apis` (see [Add a REST Endpoint](/docs/how-to/add-a-rest-endpoint)).

## Setup

Tag a feature with `@BeforeAll` so it runs once, before any other scenario:

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

## Use the token

Reference the global variable in any scenario:

```gherkin
Scenario: Read protected data
  Given I prepare a request to "my_api.get_data"
  And I set the header "Authorization" to "Bearer {{auth_token}}"
  When I send the request
  Then the response status code should be 200
```

## Verify

Run `tkit run` and confirm the `@BeforeAll` scenario runs first and subsequent scenarios succeed with the injected header.

## Common issues

- If the `@BeforeAll` hook fails, no test scenarios run — keep it focused and resilient.
- Global variables (set in hooks) are available to all scenarios; scenario variables (`I store the value ... into "name" variable`) are cleared after each scenario.

## See also

- [Global Hooks](/docs/patterns/global-hooks) — full setup/teardown reference
- [Variables](/docs/patterns/variables) — variable scope rules
- [Configure Authentication](/docs/how-to/configure-authentication) — static credentials
