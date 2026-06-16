---
title: Selectors
description: Learn how TestFlowKit finds and interacts with elements on web pages
navigation:
  title: Selectors
---

# Selectors

Steps reference elements by name — TestFlowKit looks up the selector in `testflowkit.yml` and tries each fallback until one matches.

```gherkin
When the user clicks the "submit_button" button
```

```yaml
frontend:
  elements:
    login:
      submit_button:
        - "[data-testid='submit']"
        - "#submit-btn"
        - "button[type='submit']"
```

## Defining elements

Group by page name (matching `pages`) or use `common` for shared elements:

```yaml
frontend:
  pages:
    login: "/login"
  elements:
    common:
      header: "#main-header"
    login:
      email_field: "#email"
      submit_button:
        - "[data-testid='submit']"
        - "#submit-btn"
```

## Selector types

| Type | Example |
|------|---------|
| CSS ID | `#email` |
| CSS class | `.btn-primary` |
| Data attribute | `[data-testid='login']` |
| Attribute | `[name='email']` |
| XPath | `xpath://button[contains(text(), 'Submit')]` |

Prefer CSS over XPath when possible. Prefix XPath selectors with `xpath:`.

## Best practices

| Do | Don't |
|----|-------|
| Use `data-testid` attributes | Rely on deep DOM paths like `div > div:nth-child(3) > button` |
| List fallbacks most-reliable first | Use a single brittle selector |
| Use semantic names (`submit_button`) | Use generic names (`btn1`) |

Work with developers to add `data-testid` to key interactive elements.

## Debugging

Test selectors in browser DevTools:

```javascript
document.querySelector('#my-element')       // CSS
$x("//button[text()='Submit']")           // XPath
```

Run with debug output to see which selector matched:

```bash
tkit run --debug
```

Increase `frontend.default_timeout` or `settings.think_time` if elements load slowly.

## Next Steps

- [Frontend Testing](/docs/guides/frontend-testing) — UI testing guide
- [testflowkit.yml](/docs/config/overview) — Full frontend config
