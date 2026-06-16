---
title: Skipping Tests
description: Skip individual scenarios or entire features using the @skip tag
navigation:
  title: Skipping Tests
---

# Skipping Tests

Tag a scenario or feature with `@skip` to exclude it from all test runs.

## Skip one scenario

```gherkin
Feature: User account

  Scenario: Successful login
    Given I prepare a request to "my_api.login"
    When I send the request
    Then the response status code should be 200

  @skip
  Scenario: Forgot password
    Given I prepare a request to "my_api.forgot_password"
    When I send the request
    Then the response status code should be 200
```

## Skip an entire feature

```gherkin
@skip
Feature: Experimental payments
  Scenario: Pay with crypto
    ...
```

All scenarios in a `@skip` feature are excluded.

Skipped scenarios log a warning so you can spot them in CI output:

```
⚠  Skipping scenario: "Forgot password" in feature: "User account"
```

Use for flaky tests, work-in-progress scenarios, or features waiting on unfinished APIs.

## Next Steps

- [CLI Reference](/docs/reference/cli) — Tag filtering with `--tags`
