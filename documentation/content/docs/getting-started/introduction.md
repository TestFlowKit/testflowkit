---
title: Introduction
description: Welcome to TestFlowKit - a powerful, Gherkin-based testing framework for web applications
navigation:
  title: Introduction
---

# Introduction to TestFlowKit

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

## Who it's for

- **QA** — Write and maintain tests without coding
- **Developers** — Integrate into CI/CD with a single binary
- **Product** — Read scenarios as living documentation

## Next Steps

[Installation](/docs/getting-started/installation) → [Quick Start](/docs/getting-started/quick-start)

Or browse the [documentation hub](/docs) to pick a path for your role.
