---
title: Writing Tests
description: Gherkin syntax and patterns for writing TestFlowKit scenarios
navigation:
  title: Writing Tests
---

# Writing Tests

TestFlowKit scenarios live in `.feature` files written in Gherkin — plain English with a fixed structure. No coding required.

## Basic structure

```gherkin
Feature: Login

  Scenario: Successful login
    Given the user goes to the "login" page
    When the user enters "john@example.com" into the "email" field
    And the user enters "mypassword" into the "password" field
    And the user clicks the "login" button
    Then the page title should be "Dashboard"
```

| Keyword | Role |
|---------|------|
| **Feature** | What you're testing |
| **Background** | Steps run before every scenario in the file |
| **Scenario** | One test case |
| **Given / When / Then / And / But** | Steps (interchangeable — use for readability) |

TestFlowKit matches steps by their text, not the keyword.

## Writing a good scenario

1. Name the scenario clearly — what behavior and outcome?
2. Write steps as you'd explain them to a colleague
3. Keep one behavior per scenario
4. Use tags to group tests: `@smoke`, `@regression`

```bash
tkit run --tags @smoke
tkit run --tags "@login and not @slow"
```

## Data in steps

**Tables** — key/value pairs:

```gherkin
And I set the following path parameters:
  | id | 123 |
```

**Doc strings** — multi-line text:

```gherkin
And I set the request body to:
  """
  { "name": "John", "email": "john@example.com" }
  """
```

## Scenario Outline

Run the same steps with different data:

```gherkin
Scenario Outline: Login attempts
  When the user enters "<email>" into the "email" field
  And the user enters "<password>" into the "password" field
  Then the "<result>" should be visible

  Examples:
    | email             | password | result        |
    | valid@example.com | correct  | dashboard     |
    | bad@example.com   | wrong    | error_message |
```

## Common patterns

**Navigation:**
```gherkin
Given the user goes to the "home" page
When the user clicks the "about" link
Then the current URL should contain "/about"
```

**Forms:**
```gherkin
When the user enters "Jane Doe" into the "name" field
And the user clicks the "submit" button
Then the "success_message" should be visible
```

**Visibility:**
```gherkin
Then the "sidebar" should be visible
And the "login_button" should not be visible
```

## What developers set up for you

- **testflowkit.yml** — Page URLs, element selectors, API definitions
- **Macros** — Reusable login/setup flows ([Macros](/docs/patterns/macros))

When a test fails, share the scenario name, error message, and what you expected.

Use [macros](/docs/patterns/macros) to hide repetitive setup steps.

## Step catalog

Browse the **[Step Definitions catalog](/sentences)** for every available sentence — searchable by keyword and category.

## Next Steps

- [Frontend Testing](/docs/guides/frontend-testing) — UI testing guide
- [Variables](/docs/patterns/variables) — Dynamic test data
- [Common Issues](/docs/troubleshooting/common-issues) — Troubleshooting
