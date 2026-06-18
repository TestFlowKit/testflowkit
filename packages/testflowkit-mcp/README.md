# @testflowkit/mcp

TestFlowKit MCP server for Cursor and VS Code Copilot.

Provides the `testflowkit` MCP server with tools for generating valid Gherkin tests using the TestFlowKit step catalog.

## Prerequisites

- Node.js >= 22
- `tkit` CLI installed (the server exports the step catalog from the local CLI)
- A `testflowkit.yml` in your project root (run `tkit init` to generate one)

## Getting started

### 1. Add the `agent:` section to `testflowkit.yml`

Run `tkit init` to generate `testflowkit.yml` with the agent block included, or add it manually:

```yaml
agent:
  default_tags_for_draft: "@wip @ai-generated"
  run_command: "tkit run --tags @wip"

  step_catalog:
    # Omit file/url to use the local tkit CLI export (default).
    # file: "./build/step-definitions.json"
    # url: "https://..."
```

The `agent:` block is ignored by the `tkit` CLI — it is only read by this MCP server.

### 2. Configure Cursor

Add to your project's `.cursor/mcp.json` (or the global `~/.cursor/mcp.json`):

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

### 3. Configure VS Code (Copilot with MCP)

Add to `.vscode/mcp.json`:

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

### Development (local path)

```bash
cd packages/testflowkit-mcp
npm install
npm run build
```

Then in `.cursor/mcp.json`:

```json
{
  "mcpServers": {
    "testflowkit": {
      "command": "node",
      "args": ["/absolute/path/to/testflowkit/packages/testflowkit-mcp/dist/index.js"]
    }
  }
}
```

## Typical workflow

1. Call `get_step_categories` to list available step categories.
2. Call `get_step_catalog` (optionally with a `category`) to fetch matching sentences.
3. Call `get_config_schema` to load the authoritative `testflowkit.yml` schema before generating or editing config.
4. Call `read_test_config` to discover APIs, operations, pages, and element groups.
5. Call `write_feature` to create or update a `.feature` file under `settings.gherkin_location`.

## Available tools

| Tool | Description |
|------|-------------|
| `get_step_categories` | List step definition categories from the catalog |
| `get_step_catalog` | Full step catalog; pass optional `category` (from `get_step_categories`) to filter |
| `get_config_schema` | Full JSON schema exported from `tkit export-config-schema` for AI-safe config authoring |
| `read_test_config` | Summary of `testflowkit.yml` (APIs, pages, elements — secrets redacted) |
| `list_features` | List all `.feature` files under `settings.gherkin_location` |
| `read_feature` | Read a specific feature file (`path` relative to project root) |
| `write_feature` | Create or overwrite a feature file (path-guarded to `settings.gherkin_location`) |

## Catalog resolution

The server exports the step catalog directly from the installed `tkit` CLI (`tkit export-step-definitions` under the hood). The `agent.step_catalog.file` and `agent.step_catalog.url` fields in `testflowkit.yml` are reserved for a future release.

To generate a local catalog snapshot (useful on feature branches):

```bash
tkit export-step-definitions --format json > build/step-definitions.json
```

Set `TESTFLOWKIT_CLI_PATH` to point at a non-default `tkit` binary when needed.

## Configuration

The server loads `testflowkit.yml` from the workspace root. `config.yml` is accepted as a legacy fallback.

Feature file tools resolve paths from `settings.gherkin_location` (default `./features`).

## Troubleshooting

**MCP server not starting**

- Verify `testflowkit.yml` exists in the workspace root.
- Run the server manually to see errors: `npx @testflowkit/mcp`

**Catalog not loading**

- Ensure `tkit` is installed and `tkit version` works.
- Run `tkit export-step-definitions --format json` manually to confirm the CLI export works.

## `.gitignore` (recommended)

```gitignore
.testflowkit/
**/step-definitions.json
```

## License

MIT
