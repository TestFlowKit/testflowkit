# @testflowkit/mcp

TestFlowKit MCP server for Cursor and VS Code Copilot.

Provides the `testflowkit` MCP server with tools for generating valid Gherkin tests using the TestFlowKit step catalog.

## Relationship with `@testflowkit/cli`

This package is a companion to the TestFlowKit CLI (`tkit`, published as [`@testflowkit/cli`](https://www.npmjs.com/package/@testflowkit/cli)). It does not run tests itself — you still execute scenarios with `tkit run`.

| Concern | CLI (`tkit`) | MCP server (`@testflowkit/mcp`) |
|---------|--------------|----------------------------------|
| Run and validate tests | Yes | No |
| Export step catalog | `tkit export-step-definitions` | Calls the CLI command at runtime |
| Export config schema | `tkit export-config-schema` | Calls the CLI command at runtime |
| Project config | Reads `testflowkit.yml` / `config.yml` | Reads the same files from the workspace root |
| `agent:` block in config | Ignored | Used for draft tags, run command, and future catalog overrides |

**Install order:** install `@testflowkit/cli` globally (or ensure `tkit` is on your `PATH`), then configure this MCP server in your IDE. The server shells out to your local `tkit` binary; set `TESTFLOWKIT_CLI_PATH` if it is not on `PATH`.

**Typical loop:** use MCP tools in the IDE to browse the catalog and write `.feature` files, then run `tkit run` (or the `agent.run_command` from `testflowkit.yml`) to execute them.

CLI package README: [`npm/README.md`](https://github.com/TestFlowKit/testflowkit/tree/main/npm)

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

1. Pin framework documentation resources (e.g. `docs://framework/features/api_testing`) when you need macros, variables, hooks, random data, API testing, or frontend testing mechanics.
2. Call `get_step_categories` to list available step categories.
3. Call `get_step_catalog` (optionally with a `category`) to fetch matching sentences.
4. Call `get_config_schema` to load the authoritative `testflowkit.yml` schema before generating or editing config.
5. Call `read_test_config` to discover APIs, operations, pages, and element groups.
6. Call `write_gherkin_file` to create or update a `.feature` file under `settings.gherkin_location`.

## Available tools

| Tool | Description |
|------|-------------|
| `get_step_categories` | List step definition categories from the catalog |
| `get_step_catalog` | Full step catalog; pass optional `category` (from `get_step_categories`) to filter |
| `get_config_schema` | Full JSON schema exported from `tkit export-config-schema` for AI-safe config authoring |
| `read_test_config` | Summary of `testflowkit.yml` (APIs, pages, elements — secrets redacted) |
| `list_gherkin_files` | List all `.feature` files under `settings.gherkin_location` |
| `read_gherkin_file` | Read a specific Gherkin feature file (`path` relative to project root) |
| `write_gherkin_file` | Create or overwrite a Gherkin feature file (path-guarded to `settings.gherkin_location`) |

## Framework documentation resources

Framework mechanics are exposed as **MCP resources** (not tools). Pin them in Cursor or read them via `resources/read`:

| URI | Description |
|-----|-------------|
| `docs://framework/features/index` | Index of all framework documentation pages |
| `docs://framework/features/macros` | Macros — reusable parameterized scenarios |
| `docs://framework/features/random_data` | Random data generators (`{{ rand:... }}`) |
| `docs://framework/features/global_hooks` | `@BeforeAll` / `@AfterAll` setup and teardown |
| `docs://framework/features/variables` | Scenario, environment, and global variables |
| `docs://framework/features/api_testing` | REST and GraphQL API testing |
| `docs://framework/features/frontend_testing` | Browser automation and UI testing |

Docs are synced from `documentation/content/docs/patterns/` and `documentation/content/docs/guides/` at build time into `docs/features/` inside this package.

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
