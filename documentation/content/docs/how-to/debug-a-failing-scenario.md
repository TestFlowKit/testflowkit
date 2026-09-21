---
title: Debug a Failing Scenario
description: Find out why a step fails, from the report to full request traces
navigation:
  title: Debug a Failing Scenario
---

## Goal

Go from "the scenario is red" to the cause, using the least noisy tool that answers your question.

## 1. Read the failure

Open the report (`report.html`, see [Reporters](/docs/reference/reporters)) and find the first failed step. The error message and, for UI steps with [screenshot on failure](/docs/how-to/capture-a-failure-screenshot), the picture usually tell you which case you are in:

| Symptom | Likely cause |
|---|---|
| Element not found / timeout | Wrong selector, wrong page name, or the page is slow: [Wait for an element](/docs/how-to/wait-for-an-element) |
| `{{ ... }}` or `${...}` shows up literally | Variable never set or misspelled: [Variables](/docs/patterns/variables) |
| Unexpected status code or body | Wrong endpoint, auth or payload: rerun with the HTTP scope below |
| Fails only in headless / CI | Window size, timing, or a missing env value |

## 2. Rerun only that scenario, with logs

```bash
tkit run --debug --debug-scenario "Login"
```

`--debug-scenario` keeps the output to scenarios whose name or tag contains the text. Then narrow by layer:

```bash
tkit run --debug --debug-scenario "Login" --debug-scope http       # requests and responses
tkit run --debug --debug-scenario "Login" --debug-scope variables  # how values were resolved
tkit run --debug --debug-scenario "Login" --debug-scope browser    # browser actions
```

Need the full request and response bodies? Use `--verbosity 3`. Keep the output for later with `--log-file debug.log`, or `--log-format json` to pipe it into `jq`. All levels are listed in the [CLI Reference](/docs/reference/cli).

## 3. Watch the browser

`tkit run` is headless by default. To see the window:

```bash
tkit run --headless=false --tags "@login"
```

Set `settings.think_time` (ms) in `testflowkit.yml` to slow the actions down so you can follow them.

## 4. Inspect inside the scenario

Temporary steps that print what TestFlowKit sees:

```gherkin
And I display the value of variable "token"
And I display the request curl
And I debug the api response
```

Remove them once fixed. Use `I debug the api request` to print the request that is about to be sent.

## Verify

The failing step now shows the value or request that explains it. Fix it, then rerun the scenario without `--debug` and check that it passes.

## See also

- [Common Issues](/docs/troubleshooting/common-issues)
- [CLI Reference](/docs/reference/cli): debug flags
- [Step Catalog](/sentences): debug steps
