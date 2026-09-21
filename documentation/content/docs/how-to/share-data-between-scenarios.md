---
title: Share Data Between Scenarios
description: Persist a value across scenarios using global variables
navigation:
  title: Share Data Between Scenarios
---

## Goal

Compute or fetch a value once (a token, an ID, a generated email) and reuse it in several scenarios. Regular scenario variables are cleared after each scenario, so they can't be used for this.

## Setup

Set a global variable from a `@BeforeAll` hook or from any scenario:

```gherkin
# From an API response
And I save the response path "token" as global variable "auth_token"

# Direct value
And I store the value "https://api.example.com" into global variable "api_base_url"
```

## Use it anywhere

```gherkin
Scenario: Use the shared value
  Given I prepare a request to "my_api.get_data"
  And I set the header "Authorization" to "Bearer {{auth_token}}"
  When I send the request
  Then the response status code should be 200
```

## Scope reference

| Variable | Scope | Set with |
|----------|-------|----------|
| Scenario | Current scenario only | `I store the value ... into "name" variable` |
| Global | Entire test run | `... into global variable "name"` or `I save the response path ... as global variable "name"` |
| Environment | All scenarios, from config | Defined in `testflowkit.yml` under `env` |

## Verify

Set a global variable in one scenario (or a `@BeforeAll` hook) and confirm a later scenario can read it via `{{name}}`.

## Common issues

- Using `{{name}}` from a scenario variable in another scenario silently resolves to an empty value — use a global variable instead.

## See also

- [Variables](/docs/patterns/variables) — full variable reference
- [Global Hooks](/docs/patterns/global-hooks) — `@BeforeAll` / `@AfterAll`
- [Authenticate Before Tests](/docs/how-to/authenticate-before-tests)
