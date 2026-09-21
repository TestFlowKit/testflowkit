---
title: Introduction
description: Welcome to TestFlowKit - a powerful, Gherkin-based testing framework for web applications
navigation:
  title: Introduction
---

TestFlowKit lets you write automated tests in plain text (Gherkin). No Go or JavaScript required — define pages and APIs in YAML, write scenarios in `.feature` files, run with `tkit run`.

```gherkin
Feature: User Login

  Scenario: Successful login
    Given the user goes to the "login" page
    When the user enters "john@example.com" into the "email" field
    And the user enters "password123" into the "password" field
    And the user clicks the "submit" button
    Then the page title should be "Dashboard"
```

## What you get

| Area | Capabilities |
|------|--------------|
| **Frontend** | Browser automation (Rod or Playwright) |
| **Backend** | REST and GraphQL testing |
| **Data** | Variables, random data, macros, global hooks |
| **Tooling** | HTML reports, MCP server for AI-assisted test writing |

## How it fits together

Three things, and nothing else to install:

1. **`testflowkit.yml`** declares your pages, elements and APIs under logical names.
2. **`.feature` files** describe scenarios using those names.
3. **`tkit`** runs the scenarios and writes a report.

Unfamiliar word? See the [Glossary](/docs/reference/glossary).

## Who it's for

- **QA** — Write and maintain tests without coding
- **Developers** — Integrate into CI/CD with a single binary
- **Product** — Read scenarios as living documentation

## Next Steps

[Installation](/docs/getting-started/installation) → [Quick Start](/docs/getting-started/quick-start)

Or browse the [documentation hub](/docs) to pick a path for your role. After your first run, pick a recipe under **Recipes** in the sidebar.
