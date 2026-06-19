---
title: Variables
description: Store and reuse dynamic data throughout your test scenarios
---

# Variables

Use `{{variable_name}}` in any step — strings, tables, request bodies. Values are substituted automatically at runtime.

```gherkin
When I store the value "John Doe" into "user_name" variable
When the user enters "{{user_name}}" into the "name" field
Then the value of the "name" field should be "{{user_name}}"
```

## Variable types

| Type | Store | Use |
|------|-------|-----|
| **Scenario** | `I store the value "..." into "name" variable` | `{{name}}` — cleared after the scenario |
| **Environment** | Defined in `testflowkit.yml` | `{{ env.api_key }}`, `{{ env.database.host }}` |
| **Response path** | `I store the response path "data.id" from the response into "user_id" variable` | JSON (GJSON) or XML (XPath) |
| **Page element** | `I store the content of "page_title" into "title" variable` | Text from a UI element |
| **Random data** | `I store the value "{{ rand:uuid }}" into "id" variable` | See [Random Data](./random_data.md) |

### Environment variables

Define in `testflowkit.yml`:

```yaml
env:
  api_key: "your-api-key"
  base_url: "https://api.example.com"
  database:
    host: "localhost"
    port: "5432"
```

Use in Gherkin with dot notation for nested values:

```gherkin
Given the user goes to "{{ env.base_url }}"
And I set the header "Authorization" to "Bearer {{ env.api_key }}"
```

Load environment-specific files at runtime:

```bash
tkit run --env-file .env.staging.yml
```

### Response paths

JSON responses use [GJSON](https://github.com/tidwall/gjson) syntax. XML responses use XPath.

| Expression | Example |
|------------|---------|
| `user.email` | Nested field |
| `items.0` | First array item |
| `items.#` | Array length |
| `//user/id` | XML element (XPath) |

```gherkin
When I send the request
And I store the response path "data.id" from the response into "user_id" variable
And I store the response path "data.name" from the response into "user_name" variable
```

## API → UI flow

```gherkin
Scenario: Create user via API and verify in UI
  Given I prepare a request to "my_api.create_user"
  And I set the request body to:
    """
    { "name": "Test User", "email": "test@example.com" }
    """
  When I send the request
  And I store the response path "data.id" from the response into "user_id" variable
  Then the response status code should be 201

  Given the user goes to the "users" page
  When the user enters "{{user_id}}" into the "search" field
  Then the "user_row" should contain the text "Test User"
```

## Scope

| Variable | Scope |
|----------|-------|
| Scenario (`{{name}}`) | Current scenario only |
| Environment (`{{ env.* }}`) | All scenarios, loaded at startup |
| Global (`{{name}}` set in hooks) | All scenarios — see [Global Hooks](./global_hooks.md) |

::alert{type="warning"}
Scenario variables do not persist across scenarios. Use global hooks when you need to share data between scenarios.
::

## Next Steps

- [Random Data](./random_data.md) — Generate UUIDs, emails, dates inline
- [Macros](./macros.md) — Reusable parameterized steps
- [Global Hooks](./global_hooks.md) — Cross-scenario setup and teardown
