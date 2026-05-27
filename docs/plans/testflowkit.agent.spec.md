# `testflowkit.agent.yml` specification

Field reference for the TestFlowKit IDE agent configuration file.

The file lives at the **project root** alongside `testflowkit.yml` (or the legacy `config.yml`). It is loaded by the `testflowkit-mcp` server and is intentionally separate from the runtime test configuration — it is never read by the `tkit` CLI.

---

## File discovery (MCP server)

The MCP server resolves `testflowkit.agent.yml` in order:

1. Environment variable `TESTFLOWKIT_AGENT_CONFIG` (absolute or workspace-relative path).
2. Walk upward from `process.cwd()` looking for `testflowkit.agent.yml` (max 5 levels).
3. Fatal error with a clear message if not found.

---

## Full example

```yaml
# testflowkit.agent.yml — TestFlowKit IDE agent configuration
version: 1

project:
  test_config: "./testflowkit.yml"
  features_glob: "features/**/*.feature"

step_catalog:
  # Uncomment ONE of the options below, or omit all to auto-fetch from GitHub Release.

  # Option A — local file (for contributors, offline, or canary builds)
  # file: "./build/step-definitions.json"

  # Option B — explicit URL (forks, mirrors, GitHub Enterprise)
  # url: "https://example.com/artifacts/step-definitions.json"

agent:
  default_tags_for_draft: "@wip @ai-generated"
  run_command: "tkit run -c testflowkit.yml --tags @wip"
```

---

## Fields reference

### `version`

| | |
|-|-|
| Type | integer |
| Required | yes |
| Allowed | `1` |

Schema version. Currently only `1` is valid. The MCP server rejects unsupported versions with a clear error.

---

### `project`

#### `project.test_config`

| | |
|-|-|
| Type | string (path) |
| Default | `./testflowkit.yml` |

Path to the TestFlowKit runtime configuration. Resolved relative to the directory containing `testflowkit.agent.yml`.

Legacy `config.yml` is accepted automatically when `testflowkit.yml` does not exist in the same directory (mirrors `config.ResolveConfigPath` behavior in the Go CLI).

#### `project.features_glob`

| | |
|-|-|
| Type | string (glob) |
| Default | `features/**/*.feature` |

Glob pattern for feature files, relative to the agent file directory. Used by `list_features`, `read_feature`, and to validate paths in `write_feature`.

---

### `step_catalog`

Controls how the MCP server resolves the step definitions catalog (`step-definitions.json`).

**Resolution order** (first match wins):

1. `step_catalog.file` — read directly from disk.
2. `step_catalog.url` — HTTP GET.
3. Auto-detect: run `tkit version`, strip `-canary.<sha>` suffix, fetch from `https://github.com/TestFlowKit/testflowkit/releases/download/{cliVersion}/step-definitions.json`.
4. Read/write cache at `.testflowkit/cache/step-definitions-<version>.json`.

#### `step_catalog.file`

| | |
|-|-|
| Type | string (path) |
| Default | unset |

Local path to a pre-generated `step-definitions.json`. Useful for:

- Contributors working on a branch with new steps (run `make export-step-definitions` first).
- Air-gapped environments.
- Canary CLI installs where no release catalog exists yet.

Resolved relative to the agent file directory.

#### `step_catalog.url`

| | |
|-|-|
| Type | string (HTTPS URL) |
| Default | unset |

Direct URL to download the catalog JSON. Supports `{{ env.VAR }}` interpolation for tokens or private URLs. Takes precedence over the automatic release fetch.

#### `step_catalog.release`

This section is not user-configurable. The MCP server always fetches from the official TestFlowKit release source when `step_catalog.file` and `step_catalog.url` are both unset.

#### `step_catalog.cache`

This section is not user-configurable. The MCP server manages cache location and keying implicitly.

Fetched catalogs are written under `.testflowkit/cache/` using versioned filenames:

- `.testflowkit/cache/step-definitions-<resolved-base-version>.json`

This prevents collisions across CLI versions while keeping cache files out of git (add `.testflowkit/` to `.gitignore`; `tkit init` already does this).

---

### `agent`

#### `agent.default_tags_for_draft`

| | |
|-|-|
| Type | string |
| Default | `@wip @ai-generated` |

Space-separated Gherkin tags added to every scenario written by the agent. Keeps AI-generated scenarios isolated so they can be reviewed before running in CI.

#### `agent.capabilities`

##### `agent.capabilities.macros`

| | |
|-|-|
| Type | boolean |
| Default | `true` |

Controls macro authoring from MCP tools.

- `true`: `write_macro` can create/update macro files and `write_feature` accepts macro content.
- `false`: `write_macro` is denied and `write_feature` rejects content containing `@macro`.

#### `agent.run_command`

| | |
|-|-|
| Type | string |
| Default | `tkit run -c testflowkit.yml --tags @wip` |

The CLI command used to run draft tests. Documented for Phase 3 (run loop). Not executed by the MCP server in Phase 1.

---

## Version alignment rules

| Condition | Behavior |
|-----------|----------|
| CLI `1.2.4`, catalog `1.2.4` | OK |
| CLI `1.2.4`, catalog `1.2.3` (patch) | Warning in tool metadata |
| CLI `1.2.4`, catalog `1.1.x` (minor) | Error in tool result; use `step_catalog.file` or pin version |
| CLI `1.2.4`, catalog `2.x.x` (major) | Error in tool result |
| Canary CLI `3.6.1-canary.abc1234` | Strip suffix → fetch `3.6.1` stable catalog |

---

## Environment variable interpolation

The following fields support `{{ env.VAR }}` placeholders (replaced from `process.env` by the MCP server at load time):

- `step_catalog.url`

All other fields use literal values only. For secrets such as authentication tokens needed to fetch a private catalog URL, store them in env and reference via `url`.

---

## `.gitignore` recommendations

```gitignore
# TestFlowKit agent cache
.testflowkit/

# Generated step catalog (rebuild with make export-step-definitions)
**/step-definitions.json
```

---

## Related

| Document | Path |
|----------|------|
| Agent roadmap | [testflowkit-agent-roadmap.md](./testflowkit-agent-roadmap.md) |
| Publish plan | [publish-step-definitions-on-release.md](./publish-step-definitions-on-release.md) |
| MCP server README | `packages/testflowkit-mcp/README.md` |
| Export script | `scripts/step_definitions_export/main.go` |
