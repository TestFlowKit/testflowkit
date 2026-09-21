---
title: Reporters
description: Report formats, where they are written and how to read them
navigation:
  title: Reporters
---

Choose the format with `settings.report_format` in `testflowkit.yml`. It accepts `html`, `json` or `junit`, and produces one report per run.

```yaml
settings:
  report_format: "junit"
```

## Formats

| `report_format` | File | Best for |
|---|---|---|
| `html` | `report.html` in the directory you run from | Reading results yourself, with failure screenshots |
| `json` | `report/report.json` | Cucumber tooling and dashboards |
| `junit` | `report/report.xml` | CI test tabs (GitHub Actions, GitLab, Jenkins) |

The `report/` folder is created if needed and the file is overwritten on each run.

## HTML

Scenarios with their steps and status. A failed step has an expandable **View Screenshot** when [`frontend.screenshot_on_failure`](/docs/how-to/capture-a-failure-screenshot) is on.

## Cucumber JSON

Follows the Cucumber JSON schema: features, then scenarios, then steps, with tags and durations in nanoseconds. A failure screenshot is an embedding with MIME type `image/png` and base64 data on the failed step. Any tool that reads Cucumber JSON can consume it.

## JUnit XML

One test suite per feature file, one test case per scenario. A failed case carries the error message and the list of step statuses. Scenarios with every step skipped are reported as skipped.

Hooks (`@BeforeAll`, `@AfterAll`) appear only when they fail, in a suite named `hooks`, so a broken setup is visible in CI.

## Exit code

Independent of the format: `0` when all tests pass, non-zero otherwise. See [CLI Reference](/docs/reference/cli).

## See also

- [Run in CI](/docs/how-to/run-in-ci)
- [Debug a failing scenario](/docs/how-to/debug-a-failing-scenario)
- [Migration Guide](/docs/troubleshooting/migration-guide): moving from the previous JSON report
