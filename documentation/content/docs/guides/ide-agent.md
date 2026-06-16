---
title: IDE agent
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
| `search_sentences` | Filter by keyword/category |
| `read_test_config` | APIs, pages, elements (secrets redacted) |
| `list_features` / `read_feature` | Browse feature files |
| `write_feature` / `write_macro` | Create or update `.feature` files |
| `get_guidelines` | Macro and agent authoring docs |

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
