---
title: Capture a Screenshot on Failure
description: See what the browser showed when a step failed
navigation:
  title: Screenshot on Failure
---

## Goal

Get a picture of the page at the moment a UI step fails, without adding any step to your scenarios.

## Turn it on

```yaml
frontend:
  screenshot_on_failure: true
```

## Where to find it

The screenshot is stored inside the report, next to the failed step. There is no separate image file.

| `report_format` | Where |
|---|---|
| `html` | Expand **View Screenshot** under the failed step in `report.html` |
| `json` | Embedded as a base64 PNG in `report/report.json`, on the failed step |
| `junit` | Not included |

## Verify

Break a step on purpose (wrong text in an assertion), run `tkit run`, and open the report. The failed step should have the screenshot.

Screenshots are taken only for UI steps that have a page open. An API-only failure has none.

## See also

- [Reporters](/docs/reference/reporters)
- [Debug a failing scenario](/docs/how-to/debug-a-failing-scenario)
