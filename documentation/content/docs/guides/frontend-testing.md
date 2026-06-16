---
title: Frontend Testing
description: Browser automation and UI testing capabilities
navigation:
  title: Frontend Testing
---

# Frontend Testing

Define pages and elements in `testflowkit.yml`, then reference them by name in your scenarios.

```yaml
frontend:
  driver: "rod"              # or "playwright"
  base_url: "{{ env.base_url }}"
  default_timeout: 10000
  headless: false
  screenshot_on_failure: true
  pages:
    login: "/login"
  elements:
    login:
      email_field: "#email"
      submit_button: "#submit"
```

```gherkin
Given the user goes to the "login" page
When the user enters "john@example.com" into the "email_field" field
And the user clicks the "submit_button" button
Then the "dashboard" should be visible
```

::alert{type="warning"}
The `driver` field is required when the `frontend` block is defined. Use `tkit install` to set up Playwright.
::

## Typical flow

Most UI scenarios follow the same pattern:

1. **Navigate** — `the user goes to the "page_name" page`
2. **Interact** — enter text, click buttons, select dropdowns, upload files
3. **Assert** — check visibility, text content, URL, or field values

TestFlowKit waits for elements automatically before interacting. Adjust `default_timeout` in config if needed.

Use `{{variable_name}}` and `{{ rand:email }}` in any step value — see [Variables](/docs/patterns/variables).

## Example

```gherkin
@FRONTEND
Feature: User registration

  Scenario: Register with valid data
    Given the user goes to the "registration" page
    When the user enters "jane@example.com" into the "email" field
    And the user enters "SecurePass123" into the "password" field
    And the user checks the "terms" checkbox
    And the user clicks the "submit" button
    Then the "confirmation" should be visible
    And the current URL should contain "/welcome"
```

## Step catalog

For the full list of frontend sentences, browse the **[Step Definitions catalog](/sentences)** — searchable by keyword and category.

## Next Steps

- [Selectors](/docs/config/selectors) — How elements are resolved
- [testflowkit.yml](/docs/config/overview) — Full frontend options
