# TestFlowKit Copilot instructions

You are helping write Gherkin feature files for a project using the **TestFlowKit** framework.

## Before writing any steps

1. Run `tkit version` in the **project workspace terminal** to capture the installed CLI version.
2. Call the `testflowkit` MCP tool `initialize_session({ cliVersion: "<version>" })` with that version string. Do this once per connection — all subsequent catalog tools will use it automatically.
3. Use the `get_step_catalog` tool (or `search_sentences`) to retrieve available sentences.
4. Use `read_test_config` before proposing `I prepare a request to "..."` — both API name and operation/endpoint name must exist in the project config.

> If `tkit` is not installed locally, try `npx --yes @testflowkit/cli version`.
> Use `get_session_info` at any time to inspect the resolved CLI version and catalog source.

## Sentence rules

- Only use sentences that exist verbatim in the catalog.
- If no sentence fits, output a `missing_sentence` block instead of inventing a step:

```yaml
missing_sentence:
  intent: "what you wanted to express"
  closest_matches:
    - "closest existing sentence"
  proposed_sentence: "proposed new sentence pattern"
  proposed_category: "restapi | frontend | graphql | variable | assertions"
```

## API and GraphQL steps

- `I prepare a request to "api_name.operation_name"` — both parts must exist in `read_test_config`.
- Use `I set the following GraphQL variables:` with a data table for GraphQL variables.
- Never hardcode secrets or tokens. Reference them via `{{ env.VAR_NAME }}`.

## Frontend steps

- Element names must match keys in `frontend.elements` from the config.
- Page names must match keys in `frontend.pages` from the config.

## Scenario tags

- Add `@wip @ai-generated` to every new scenario (configurable in `testflowkit.agent.yml`).
- Use `@macro` for reusable scenarios; use `${variable_name}` for placeholders.

## What NOT to do

- Do not invent sentences not present in the catalog.
- Do not hardcode credentials or base URLs.
- Do not remove `@wip` from agent-generated scenarios.
