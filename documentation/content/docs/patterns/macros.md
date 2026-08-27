---
title: Macros
description: Create reusable, parameterized test patterns with macros
navigation:
  title: Macros
---

# Macros

Macros are reusable scenarios you define once and call with different parameters. Use `${variable}` placeholders in the macro, then pass values via a data table at call time.

## Define a macro

Tag the scenario with `@macro` and use `${variable_name}` for placeholders:

```gherkin
@macro
Scenario: Login as user
  Given the user goes to the "login" page
  When the user enters "${email}" into the "email" field
  And the user enters "${password}" into the "password" field
  And the user clicks the "login" button
  Then the "dashboard" should be visible
```

`@macro` scenarios are **not executed by `tkit run`**. They are excluded from normal test runs and only run when invoked as a step from a regular (non-`@macro`) scenario.

## Call a macro

Use the scenario name as a step, followed by a two-column table (`variable | value`):

```gherkin
Scenario: Admin can access dashboard
  Given Login as user
    | email    | admin@example.com |
    | password | admin123          |
  When the user clicks the "admin_panel" link
  Then the "admin_dashboard" should be visible
```

Each row maps one `${variable}` to a value. Missing variables cause the macro to fail.

## Reuse the same macro from multiple scenarios

The whole point of a macro is calling it more than once with different data:

```gherkin
Scenario: Admin can access dashboard
  Given Login as user
    | email    | admin@example.com |
    | password | admin123          |
  When the user clicks the "admin_panel" link
  Then the "admin_dashboard" should be visible

Scenario: Regular user cannot access dashboard
  Given Login as user
    | email    | user@example.com |
    | password | user123          |
  Then the "admin_panel" link should not be visible
```

## Discover existing macros with the AI agent

If you use the TestFlowKit MCP server ([IDE Agent](/docs/guides/ide-agent)), call
the `list_macros` tool before drafting a new scenario. It scans your project's
feature files for `@macro` scenarios and returns each macro's name, its
required `${variable}` placeholders, and a ready-to-paste call example —
so the agent reuses existing macros instead of duplicating steps or
guessing variable names.

## API example

`I prepare a request to ...` resets the current request state — headers, body, path/query parameters, GraphQL variables, and any previous response — then loads the endpoint defaults from `testflowkit.yml`.

```gherkin
@macro
Scenario: I try to create a new post with the following details:
  Given I prepare a request to "jsonplaceholder.create_post"
  When I set the request body to:
    """
    {
      "title": "${title}",
      "body": "${body}",
      "userId": ${userId}
    }
    """
  And I send the request

Scenario: Create a new post
  When I try to create a new post with the following details:
    | title  | Test Post Title |
    | body   | Test body       |
    | userId | 1               |
  Then the response status code should be 201
```

## Tips

- Keep macros focused — one clear action per macro
- Put macro definitions in a dedicated `features/macros/` folder
- `@macro` scenarios never run on their own with `tkit run` — call them from a regular scenario
- Variable names in the table must match `${placeholders}` exactly
- A plain description can go under `Scenario:` before the first step — it's ignored by both `tkit` and the `list_macros` MCP tool

## Troubleshooting

| Problem | Fix |
|---------|-----|
| Macro not found | Check `@macro` tag and exact scenario name |
| `${var}` appears literally | Variable name in the table doesn't match the placeholder |
| Wrong values | Each table row must have exactly 2 columns: `name \| value` |

## Next Steps

- [Global Hooks](/docs/patterns/global-hooks) — Setup and teardown across the suite
- [Variables](/docs/patterns/variables) — Scenario and environment variables
