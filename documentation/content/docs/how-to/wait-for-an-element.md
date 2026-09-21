---
title: Wait for an Element or a Page
description: Deal with slow pages, late elements and new windows
navigation:
  title: Wait for an Element
---

## Goal

Stop a scenario from failing because the page or an element is not ready yet.

## Interactions already wait

Steps that act on an element (click, type, select) wait for it to appear before acting. The wait lasts `frontend.default_timeout`, in milliseconds. Raise it if your app is slow:

```yaml
frontend:
  default_timeout: 20000
```

For a single run, override it: `tkit run --timeout 30s`.

## Wait for something to show up

Assert on it. The scenario moves on as soon as the element is visible:

```gherkin
When the user clicks the "search" button
Then the "results_list" should be visible
```

## Wait for a new window

When a click opens a popup or a new tab:

```gherkin
When the user clicks the "open_receipt" link
And the user waits for a new window to open within "5s"
And the user switches to the newly opened window
Then the page title should be "Receipt"
When the user switches back to the original window
```

## Last resort: a fixed pause

```gherkin
And the user waits for 2 seconds
```

A fixed pause slows every run and still breaks on a slower machine. Prefer an assertion, and use a pause only when there is nothing on the page to check.

## Verify

Run the scenario with `tkit run --debug --debug-scope browser` to see the browser actions and how long each took. If the step still times out, the selector is probably wrong: see [Debug a failing scenario](/docs/how-to/debug-a-failing-scenario).

## See also

- [Frontend Testing](/docs/guides/frontend-testing)
- [Step Catalog](/sentences): navigation and visibility steps
