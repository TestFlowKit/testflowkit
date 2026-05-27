---
title: IDE agent
description: Use the TestFlowKit MCP server to generate tests with Cursor or VS Code Copilot
navigation:
  title: IDE Agent
---

# IDE agent (Cursor / VS Code Copilot)

The TestFlowKit MCP server (`@testflowkit/mcp`) connects your IDE to the step catalog and project configuration, enabling AI-assisted Gherkin test generation that stays within the registered sentence library.

## What the agent can do

- Fetch the step catalog matching your installed `tkit` version automatically.
- Search sentences by keyword or category.
- Read your `testflowkit.yml` (APIs, operations, pages, elements) to propose valid references.
- Create and update `.feature` files within your project.
- Report missing sentences when the catalog cannot express your intent.

## Prerequisites

- `tkit` CLI installed (for automatic catalog version detection).
- Node.js >= 18.
- Cursor or VS Code with MCP support.

## Quick start

### 1. Run `tkit init`

`tkit init` creates all required files automatically:

```bash
tkit init
```

Generated files:

| File | Purpose |
|------|---------|
| `testflowkit.agent.yml` | IDE agent configuration |
| `.cursor/rules/testflowkit-agent.mdc` | Cursor guardrails |
| `copilot-instructions.md` | VS Code Copilot guardrails |

### 2. Configure the MCP server

#### Cursor

Add `.cursor/mcp.json` at your project root (or globally at `~/.cursor/mcp.json`):

```json
{
  "mcpServers": {
    "testflowkit": {
      "command": "npx",
      "args": ["-y", "@testflowkit/mcp"]
    }
  }
}
```

#### VS Code (MCP extension)

Add `.vscode/mcp.json`:

```json
{
  "servers": {
    "testflowkit": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "@testflowkit/mcp"]
    }
  }
}
```

### 3. Start writing tests

Open any `.feature` file (or ask the agent to create one) and prompt:

> "Write a test for user registration using the available API operations."

The agent will call `get_step_catalog`, `read_test_config`, and then draft a scenario using only registered sentences.

## `testflowkit.agent.yml` reference

```yaml
version: 1

project:
  test_config: "./testflowkit.yml"        # path to your runtime config
  features_glob: "features/**/*.feature"   # where feature files live

step_catalog:
  # Omit file/url to auto-fetch from GitHub Release for your tkit version.
  # file: "./build/step-definitions.json"  # local override
  # url: "https://..."                     # explicit URL
  release:
    repository: "TestFlowKit/testflowkit"
    asset: "step-definitions.json"
  cache:
    path: ".testflowkit/cache/step-definitions.json"

agent:
  capabilities:
    macros: true
  default_tags_for_draft: "@wip @ai-generated"
  run_command: "tkit run -c testflowkit.yml --tags @wip"
```

See the [full field reference](/docs/reference/agent-config) for details.

## Available MCP tools

| Tool | Description |
|------|-------------|
| `get_step_catalog` | Full list of registered Gherkin sentences |
| `search_sentences` | Filter by keyword and category |
| `read_test_config` | API names, operations/endpoints, pages, elements (secrets redacted) |
| `list_features` | All feature files under `features_glob` |
| `read_feature` | Read a specific feature file |
| `get_guidelines` | Return bundled concept guidelines (e.g. `concept: "macro"`; unknown/empty concept returns all) |
| `write_feature` | Create or update a feature file (path-guarded) |
| `write_macro` | Create or update macro feature files (requires `@macro` in content) |

## Available MCP resources

| Resource URI | Description |
|---|---|
| `testflowkit://guidelines/macros` | Macro authoring and usage documentation |
| `testflowkit://guidelines/ide-agent` | IDE agent setup and usage documentation |
| `testflowkit://guidelines/copilot-instructions` | Copilot instruction template rules |

If a documentation resource does not exist in the current workspace, use `get_guidelines` as fallback.

## Macro permissions

Macro creation is controlled by `agent.capabilities.macros` in `testflowkit.agent.yml`.

- `true` (default): the agent can write macro scenarios.
- `false`: `write_macro` is denied and `write_feature` rejects content containing `@macro`.

## Catalog resolution

The server resolves the step catalog in this order:

1. `step_catalog.file` (local path)
2. `step_catalog.url`
3. Installed `tkit version` → `https://github.com/TestFlowKit/testflowkit/releases/download/{version}/step-definitions.json`
4. Cache at `step_catalog.cache.path`

For canary builds (e.g. `3.6.1-canary.abc1234`), the catalog for the semver base (`3.6.1`) is fetched.

To use a local catalog (e.g. on a branch with new steps):

```bash
make export-step-definitions
```

Then add to `testflowkit.agent.yml`:

```yaml
step_catalog:
  file: "./build/step-definitions.json"
```

## Missing sentences

When no registered sentence matches your intent, the agent outputs a `missing_sentence` block instead of inventing a step:

```yaml
missing_sentence:
  intent: "assert websocket message received"
  closest_matches:
    - "the response should contain {string}"
  proposed_sentence: "the websocket should receive a message containing {string}"
  proposed_category: "backend"
```

Collect these and open a GitHub issue (or implement the Go step) to extend the catalog.

## `.gitignore` recommendations

```gitignore
# Agent cache
.testflowkit/

# Generated catalog (rebuild with make export-step-definitions)
**/step-definitions.json
```

## Environment variable

| Variable | Purpose |
|----------|---------|
| `TESTFLOWKIT_AGENT_CONFIG` | Override path to `testflowkit.agent.yml` |

## Troubleshooting

### MCP server not starting

- Check that Node.js >= 18 is installed.
- Verify `testflowkit.agent.yml` exists in the workspace root (or set `TESTFLOWKIT_AGENT_CONFIG`).
- Run the server manually to see errors:

  ```bash
  npx @testflowkit/mcp
  ```

### Catalog not loading

- Ensure `tkit` CLI is installed and `tkit version` works.
- For canary builds: add `step_catalog.file` pointing to a local export.
- Check write permission for the cache directory (`.testflowkit/cache/`).

### Agent uses wrong sentences

- The Cursor rule (`.cursor/rules/testflowkit-agent.mdc`) instructs the model to call `get_step_catalog` first. If the rule is missing, run `tkit init` again in your project.
