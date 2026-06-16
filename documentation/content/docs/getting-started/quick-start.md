---
title: Quick Start
description: Write and run your first automated test with TestFlowKit
navigation:
  title: Quick Start
---

# Quick Start

## 1. Initialize a project

```bash
tkit init
```

Creates `testflowkit.yml`, a sample feature file, and project structure. (Legacy projects may use `config.yml` — both work.)

## 2. Or create manually

```
my-tests/
├── testflowkit.yml
└── features/
    └── login.feature
```

**testflowkit.yml:**

```yaml
settings:
  gherkin_location: "features"
  report_format: "html"

env:
  base_url: "http://localhost:3000"

frontend:
  driver: "rod"
  base_url: "{{ env.base_url }}"
  default_timeout: 10000
  headless: false
  pages:
    login: "/login"
  elements:
    login:
      email_field: "#email"
      password_field: "#password"
      submit_button: "#submit"
```

**features/login.feature:**

```gherkin
Feature: Login

  Scenario: Valid credentials
    Given the user goes to the "login" page
    When the user enters "user@example.com" into the "email_field" field
    And the user enters "password123" into the "password_field" field
    And the user clicks the "submit_button" button
    Then the page title should be "Dashboard"
```

## 3. Run

```bash
tkit run
```

TestFlowKit launches the browser, executes steps, and generates an HTML report.

## Common commands

| Command | Description |
|---------|-------------|
| `tkit run` | Run all tests |
| `tkit run --tags @smoke` | Run tagged scenarios |
| `tkit run --env-file .env.staging.yml` | Switch environment |
| `tkit validate` | Check config and feature files |
| `tkit init` | Scaffold a new project |

## Next Steps

- [Writing Tests](/docs/guides/writing-tests) — Gherkin syntax and patterns
- [testflowkit.yml](/docs/config/overview) — Full config options
- [Frontend Testing](/docs/guides/frontend-testing) — UI testing guide
