---
title: IDE Agent
description: Use the TestFlowKit MCP server to generate tests with Cursor or VS Code Copilot
navigation:
  title: IDE Agent
---

# IDE Agent (Cursor / VS Code)

The `@testflowkit/mcp` server connects your IDE to the step catalog and project config, so AI-generated tests use only registered sentences.

## Setup

**1. Initialize the project:**

```bash
tkit init
```

**2. Add MCP config:**

Cursor (`.cursor/mcp.json`):

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

VS Code (`.vscode/mcp.json`):

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

**3. Prompt the agent:**

> "Write a test for user registration using the available API operations."

The agent reads the step catalog and `testflowkit.yml`, then drafts scenarios with valid steps only.

## Requirements

- `tkit` CLI installed (for catalog version detection)
- Node.js >= 22
- `testflowkit.yml` in the project root

## MCP tools

| Tool | Purpose |
|------|---------|
| `get_step_catalog` | Full sentence list |
| `read_test_config` | APIs, pages, elements (secrets redacted) |
| `list_gherkin_files` / `read_gherkin_file` | Browse Gherkin feature files |
| `write_gherkin_file` | Create or update `.feature` files |

## Framework documentation resources

Framework mechanics are exposed as **MCP resources**. Pin them in your IDE context:

| URI | Topic |
|-----|-------|
| `docs://framework/features/index` | Documentation index |
| `docs://framework/features/macros` | Reusable parameterized scenarios |
| `docs://framework/features/random_data` | `{{ rand:... }}` generators |
| `docs://framework/features/global_hooks` | `@BeforeAll` / `@AfterAll` hooks |
| `docs://framework/features/variables` | Scenario and environment variables |
| `docs://framework/features/api_testing` | REST and GraphQL API testing |
| `docs://framework/features/frontend_testing` | Browser automation and UI testing |

## Agent config in testflowkit.yml

```yaml
agent:
  default_tags_for_draft: "@wip @ai-generated"
  run_command: "tkit run --tags @wip"
```

The `agent:` block is ignored by `tkit run` — only the MCP server reads it.

## Missing sentences

If no registered step matches your intent, the agent reports a `missing_sentence` block instead of inventing one. Use these to request new steps via GitHub issues.

## Troubleshooting

| Problem | Fix |
|---------|-----|
| MCP won't start | Run `npx @testflowkit/mcp` manually; check Node >= 22 |
| Catalog empty | Verify `tkit version` works |
| Wrong steps used | Ask the agent to call `get_step_catalog` first |

## Next Steps

- [Step Catalog](/sentences) — Browse all available sentences
- [testflowkit.yml](/docs/config/overview) — Project configuration
