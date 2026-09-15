# @testflowkit/ts

A TypeScript sibling of the TestFlowKit Go framework: the same Gherkin step sentences, running on [`@cucumber/cucumber`](https://github.com/cucumber/cucumber-js) and [Playwright](https://playwright.dev/).

This is a v1 with a deliberately small step catalog — frontend interaction/assertion steps plus scenario variables. It is not a full port of the Go framework; see [Not yet implemented](#not-yet-implemented).

## Configuration

Configuration is split across two files:

- **`testflowkit.yml`** — element selectors and page names, in the same shape as the Go framework's `frontend.elements` / `frontend.pages`. Element lookup uses the *current page name* as the group, falling back to a `common` group; each element can list multiple fallback selectors (raced concurrently, first to attach wins); prefix a selector with `xpath:` to use XPath instead of CSS.

  ```yaml
  settings:
    gherkin_location: "features"

  frontend:
    elements:
      login_e2e:
        email_field: ["#email"]
        login_button: ["#login-button"]
      common:
        submit_button: ["#submit", "button[type=submit]"]
    pages:
      login_e2e: "login"
      google: "{{ env.google_url }}"
  ```

- **`playwright.config.ts`** — everything browser-level (base URL, headless, timeout, viewport, locale, timezone, user agent), using Playwright's own `defineConfig`. This package reads `use.*` and `timeout` from it at runtime; it does not run tests through the Playwright Test runner.

  ```ts
  import { defineConfig } from '@playwright/test';

  export default defineConfig({
    timeout: 30_000,
    use: {
      baseURL: process.env.FRONTEND_BASE_URL,
      headless: true,
    },
  });
  ```

Override the file paths with `TESTFLOWKIT_CONFIG` / `TESTFLOWKIT_PLAYWRIGHT_CONFIG` env vars if they aren't at the project root.

## Usage

```js
// cucumber.mjs
export default {
  default: {
    import: ['dist/index.js'], // or '@testflowkit/ts' once published
    paths: ['features/**/*.feature'],
  },
};
```

```gherkin
Feature: Login

  Scenario: successful login
    Given the user goes to the "login_e2e" page
    When the user enters "user@example.com" into the "email" field
    And the user enters "secret" into the "password" field
    And the user clicks the "login" button
    Then the "success_message" should be visible
```

```
npx cucumber-js
```

## Step catalog (v1)

**Navigation**: `the user goes to the "{page}" page`, `the user navigate to the URL "{url}"`, `the user should be navigated to the "{page}" page`, `the user navigates back`, `the user refreshes the page`, `the user waits for {n} seconds`.

**Form**: `the user enters "{text}" into the "{field}" field`, `the user clears the "{field}" field`, `the user checks/unchecks the "{name}" checkbox`, `the user selects the "{name}" radio button`, `the user selects the option with text/value "{value}" from the "{name}" dropdown`.

**Mouse**: `the user clicks the "{name}" button/element/link`, `the user hovers on "{name}"`.

**Keyboard**: `the user presses the "{key}" key` (Enter, Tab, Delete, Escape, Space, Arrow Up/Right/Down/Left).

**Assertions**: `the "{name}" should (not) be visible`, `the "{name}" should (not) exist`, `the "{name}" should (not) contain the text "{text}"`, `the current URL should (not) contain "{text}"`, `the page title should (not) be "{text}"`, `the "{name}" checkbox should be checked/unchecked`, `the "{name}" radio button should be selected/unselected`.

**Variables**: `I store the value "{value}" into "{name}" variable`, `I store the content of "{name}" into "{varName}" variable`, `I display the value of variable "{name}"`, `the variable "{name}" should (not) contain "{text}"`.

Any step's string arguments accept `{{ varName }}` (scenario variable) and `{{ env.VAR }}` (process env) placeholders.

## Not yet implemented

Dropdown multi-select, file upload, table/visual assertions, window/tab management, `think_time`, REST/GraphQL steps, macros, HTML/JSON reporters. These exist in the Go framework and are candidates for a later version.

## Development

```
npm install
npm run build
npm run lint   # tsc --noEmit
npm test       # unit tests (node --test)
npm run e2e    # runs features/ against the example config above
```
