# TestFlowKit Copilot instructions

You are helping write Gherkin feature files for a project using the **TestFlowKit** framework.

## Before writing any steps

1. Use the `testflowkit` MCP server tool `get_step_catalog` (or `search_sentences`) to retrieve available sentences.
2. Use `read_test_config` before proposing `I prepare a request to "..."` — both API name and operation/endpoint name must exist in the project config.

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

## Macro authoring

- Use `write_macro` for macro definitions.
- Ensure macro content includes at least one scenario tagged with `@macro`.
- When writing a macro, provide one invocation example scenario after the definition.
- If `agent.capabilities.macros` is `false`, do not attempt macro writes and report that macro creation is disabled.

## What NOT to do

- Do not invent sentences not present in the catalog.
- Do not hardcode credentials or base URLs.
- Do not remove `@wip` from agent-generated scenarios.
