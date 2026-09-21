---
title: Skip a Test
description: Temporarily exclude a scenario or feature from test runs
navigation:
  title: Skip a Test
---

# Skip a Test

## Goal

Exclude a flaky, work-in-progress, or blocked scenario from every run without deleting it.

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

## Verify

Run `tkit run` and check the skipped scenario logs a warning instead of executing:

```
⚠  Skipping scenario: "Forgot password" in feature: "User account"
```

## Alternative: filter by tag at run time

To exclude/include scenarios without editing feature files, use `--tags` instead:

```bash
tkit run --tags "not @wip"
```

## See also

- [Skipping Tests](/docs/patterns/skip-tag) — full reference
- [CLI Reference](/docs/reference/cli) — `--tags` filtering
