---
title: Quick Start
description: Write and run your first automated test with TestFlowKit in about 10 minutes
navigation:
  title: Quick Start
---

Goal: run one UI scenario and open its report. You need [TestFlowKit installed](/docs/getting-started/installation) and a web app to test.

## 1. Check the install

```bash
tkit --version
```

If the command is not found, see [Installation](/docs/getting-started/installation) or [Common Issues](/docs/troubleshooting/common-issues).

## 2. Create the project

```bash
mkdir my-tests && cd my-tests
tkit init
```

`tkit init` scaffolds `testflowkit.yml` and a sample feature file. The default `rod` driver needs nothing more; if you set `driver: "playwright"`, run `tkit install` once to download its browser.

Prefer to write the files yourself? Use this layout:

```
my-tests/
├── testflowkit.yml
└── features/
    └── login.feature
```

**testflowkit.yml**

```yaml
settings:
  gherkin_location: "features"
  report_format: "html"

env:
  base_url: "http://localhost:3000"   # change to your app

frontend:
  base_url: "{{ env.base_url }}"
  default_timeout: 10000
  pages:
    login: "/login"
  elements:
    login:
      email_field: "#email"
      password_field: "#password"
      submit_button: "#submit"
```

**features/login.feature**

```gherkin
Feature: Login

  Scenario: Valid credentials
    Given the user goes to the "login" page
    When the user enters "user@example.com" into the "email_field" field
    And the user enters "password123" into the "password_field" field
    And the user clicks the "submit_button" button
    Then the page title should be "Dashboard"
```

The names `login`, `email_field` and `submit_button` in the feature are the keys you declared under `pages` and `elements`. See [Selectors](/docs/config/selectors).

## 3. Validate, then run

```bash
tkit validate
tkit run
```

`tkit validate` catches typos in the config and unknown steps before a browser starts. `tkit run` plays the steps in a headless browser and writes the report. To watch the browser, run `tkit run --headless=false`.

## 4. Read the result

The exit code is `0` when every scenario passes and non-zero otherwise, so CI can gate on it.

| `report_format` | Output |
|---|---|
| `html` | `report.html` in the current directory |
| `json` | `report/report.json` (Cucumber JSON) |
| `junit` | `report/report.xml` (JUnit XML) |

If a step fails, re-run with `tkit run --debug` to see requests and variable substitutions. See [Common Issues](/docs/troubleshooting/common-issues).

## Useful commands

| Command | Description |
|---------|-------------|
| `tkit run --tags "@smoke"` | Run tagged scenarios |
| `tkit run --env-file .env.staging.yml` | Switch environment |
| `tkit run --debug` | Verbose logs for a failing scenario |
| `tkit validate` | Check config and feature files |

Every flag is listed in the [CLI Reference](/docs/reference/cli).

## What next?

Pick the task you have in front of you:

- Test an API: [Add a REST endpoint](/docs/how-to/add-a-rest-endpoint), [Add a GraphQL operation](/docs/how-to/add-a-graphql-operation)
- Log in once for all tests: [Authenticate before tests](/docs/how-to/authenticate-before-tests)
- Run against staging: [Test multiple environments](/docs/how-to/test-multiple-environments)
- Reuse values between scenarios: [Share data between scenarios](/docs/how-to/share-data-between-scenarios)
- Look for a step to use: [Step Catalog](/sentences)
