---
title: Fill in a Form
description: Type in fields, pick options, tick boxes and check the result
navigation:
  title: Fill in a Form
---

## Goal

Fill a web form from a scenario and check that the values were taken.

## 1. Declare the page and its fields

Every field you touch needs a logical name under `frontend.elements`, grouped by page. See [Selectors](/docs/config/selectors).

```yaml
frontend:
  pages:
    signup: "/signup"
  elements:
    signup:
      email: "#email"
      country: "#country"
      newsletter: "#newsletter"
      plan_pro: "#plan-pro"
      submit: "button[type=submit]"
```

## 2. Write the scenario

```gherkin
Scenario: Sign up with the Pro plan
  Given the user goes to the "signup" page
  When the user enters "jane@example.com" into the "email" field
  And the user selects the option with text "France" from the "country" dropdown
  And the user checks the "newsletter" checkbox
  And the user selects the "plan_pro" radio button
  And the user clicks the "submit" button
  Then the current URL should contain "/welcome"
```

## Other useful steps

| Need | Step |
|---|---|
| Empty a field | `the user clears the "email" field` |
| Untick a box | `the user unchecks the "newsletter" checkbox` |
| Pick by value or index | `the user selects the option with value "fr" from the "country" dropdown` |
| Press a key | `the user presses the enter key` |
| Random test data | `the user enters "{{ rand:email }}" into the "email" field` |

Random values are covered in [Random Data](/docs/patterns/random-data).

## Verify

Assert on the form state before submitting if it matters:

```gherkin
Then the value of the "email" field should be "jane@example.com"
And the "newsletter" checkbox should be checked
And the "country" dropdown should have "France" selected
```

Run `tkit validate` first: it flags a step that does not exist. If an element is not found, see [Common Issues](/docs/troubleshooting/common-issues).

## See also

- [Step Catalog](/sentences): every form step with its exact wording
- [Upload a file in the browser](/docs/how-to/upload-a-file-in-the-browser)
- [Selectors](/docs/config/selectors)
