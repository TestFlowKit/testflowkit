# @testflowkit/mcp

TestFlowKit MCP server for Cursor and VS Code Copilot.

Provides the `testflowkit` MCP server with tools for generating valid Gherkin tests using the TestFlowKit step catalog.

## Prerequisites

- Node.js >= 18
- `tkit` CLI installed (for automatic catalog version resolution)
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

Also add a `copilot-instructions.md` at your project root to guide the Copilot model.

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

## Available tools

| Tool | Description |
|------|-------------|
| `get_step_catalog` | Full step definitions catalog from the resolved source |
| `search_sentences` | Search catalog by keyword and/or category |
| `read_test_config` | Summary of `testflowkit.yml` (APIs, pages, elements — secrets redacted) |
| `list_features` | List all `.feature` files under `features_glob` |
| `read_feature` | Read a specific feature file |
| `get_guidelines` | Return bundled concept guidelines (e.g. `concept: "macro"`; unknown/empty concept returns all) |
| `write_feature` | Create or update a feature file (path-guarded to `features_glob`) |
| `write_macro` | Create or update macro feature files containing `@macro` scenarios |

## Available resources

| Resource URI | Description |
|---|---|
| `testflowkit://guidelines/macros` | Macro authoring and usage documentation |
| `testflowkit://guidelines/ide-agent` | IDE agent setup and usage documentation |
| `testflowkit://guidelines/copilot-instructions` | Copilot instruction template rules |

If a documentation resource is missing in the workspace, the server returns a fallback note and you can still call `get_guidelines` for bundled guidance.

## Catalog resolution

The server exports the step catalog directly from the installed `tkit` CLI. The `agent.step_catalog.file` and `agent.step_catalog.url` fields in `testflowkit.yml` are reserved for a future release.

## `.gitignore` (recommended)

```gitignore
.testflowkit/
**/step-definitions.json
```

## License

MIT
