---
title: Glossary
description: The words TestFlowKit uses, in one place
navigation:
  title: Glossary
---

| Term | Meaning |
|---|---|
| **`tkit`** | The TestFlowKit command line. `tkit run` executes your tests. |
| **`testflowkit.yml`** | The project config: settings, environment values, pages, elements, APIs, files. Legacy projects may use `config.yml`. See [testflowkit.yml](/docs/config/overview). |
| **Feature file** | A `.feature` text file holding one Feature and its scenarios, in the folder set by `settings.gherkin_location`. |
| **Gherkin** | The plain-English format of feature files: Feature, Scenario, Given / When / Then. |
| **Scenario** | One test case: a list of steps. |
| **Step** | One line of a scenario, such as `the user clicks the "submit" button`. Steps are matched by their text. All available steps are in the [Step Catalog](/sentences). |
| **Tag** | A label such as `@smoke` above a Feature or Scenario. Used with `--tags` to choose what runs. `@skip`, `@macro`, `@BeforeAll` and `@AfterAll` have a special meaning. |
| **Page** | A logical name for a URL path, declared under `frontend.pages`. Used as `the user goes to the "login" page`. |
| **Element** | A logical name for something on a page, declared under `frontend.elements` with a selector. Steps use the name, never the selector. See [Selectors](/docs/config/selectors). |
| **Selector** | The CSS or XPath expression that finds an element in the page. |
| **Driver** | The browser engine: `rod` (default, bundled) or `playwright` (needs `tkit install`). |
| **API definition** | A REST or GraphQL service declared under `apis.definitions`. A request is addressed as `"api_name.endpoint_name"`. |
| **Endpoint / Operation** | A REST call (method and path) or a GraphQL query or mutation inside an API definition. |
| **Environment value** | A value in the `env:` block or an env file, read as `{{ env.name }}`. Switch files with `--env-file`. See [Test multiple environments](/docs/how-to/test-multiple-environments). |
| **Variable** | A value stored during a run and read as `{{name}}`. See [Variables](/docs/patterns/variables). |
| **Macro** | A scenario tagged `@macro` that other scenarios call as a step with parameters. See [Macros](/docs/patterns/macros). |
| **Global hook** | A scenario tagged `@BeforeAll` or `@AfterAll`, run once before or after the whole suite. See [Global Hooks](/docs/patterns/global-hooks). |
| **Report** | The result file of a run: HTML, Cucumber JSON or JUnit XML. See [Reporters](/docs/reference/reporters). |
| **MCP server** | The tool that lets an AI agent in your IDE read the step catalog and write tests. See [IDE Agent](/docs/guides/ide-agent). |
