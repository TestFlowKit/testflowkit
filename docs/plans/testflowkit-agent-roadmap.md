# TestFlowKit agent roadmap

End-to-end plan for the **TestFlowKit IDE agent** (Cursor / VS Code Copilot) and the **step catalog** published on GitHub Releases.

| Document | Scope |
|----------|--------|
| [publish-step-definitions-on-release.md](./publish-step-definitions-on-release.md) | Release CI, catalog artifact contract, download URLs |
| **This file** | All phases: foundation → IDE agent → schema → run loop → hardening |

---

## Vision

Use AI to write **Gherkin E2E tests** from natural language, using only **registered TestFlowKit sentences** and project **`config.yml`** (APIs, GraphQL operations, pages, elements). The agent should:

- Generate valid `.feature` files and optional macros.
- Suggest missing sentences when the catalog cannot express intent.
- Eventually run tests and iterate on failures.
- Discover gaps in the sentence library for framework contributors.

---

## Roadmap at a glance

| Phase | Name | Status | Outcome |
|-------|------|--------|---------|
| **0** | Foundation — step catalog on release | **Done** | `step-definitions.json` on every stable GitHub Release |
| **1** | IDE agent (Cursor / Copilot) | **Next** | Rules + MCP + `testflowkit.agent.yml` + catalog fetch |
| **2** | Schema-aware generation | Planned | GraphQL / OpenAPI from file or URL |
| **3** | Run loop & golden tests | Planned | `tkit run` / validate, acceptance scenarios, gap reports |
| **4** | Hardening & distribution | Planned | Schema file, checksums, npm bundle, enterprise mirrors |

```text
Phase 0 ──► Phase 1 ──► Phase 2 ──► Phase 3 ──► Phase 4
 (release)   (copilot)   (schemas)   (run/QA)    (ops)
```

---

## Phase 0 — Foundation: step catalog on GitHub Release

**Status:** Done

### Goals

- Ship machine-readable **step definitions** with every **stable** release.
- Same export path locally, on PR CI, and on release.
- Enable IDE agents to fetch the catalog without committing it to test repos.

### Delivered

| Item | Location / command |
|------|-------------------|
| Export target | `make export-step-definitions` → `build/step-definitions.json` |
| Stable release upload | `.github/workflows/release.yml` |
| PR validation | `.github/workflows/pull-request.yml` (`jq` checks) |
| Post-upload verify | `release.yml` downloads asset and validates JSON |
| Documentation | `README.md`, [publish plan](./publish-step-definitions-on-release.md) |

### Published artifact

| Channel | Catalog on GitHub? | How agents resolve |
|---------|-------------------|-------------------|
| **Stable** (`1.2.4`, no `v` prefix) | Yes | `https://github.com/TestFlowKit/testflowkit/releases/download/<tag>/step-definitions.json` |
| **Canary** (`3.6.1-canary.abc1234` on npm) | No | Local `make export-step-definitions`, or stable URL for semver base `3.6.1` |

### Catalog entry shape

```json
{
  "sentence": "I prepare a request to {string}",
  "description": "...",
  "categories": ["restapi"],
  "example": "Given I prepare a request to \"users_api.getUser\"",
  "variables": [
    { "name": "request name", "description": "...", "type": "string" }
  ]
}
```

### Agent rules (catalog)

1. Fetch catalog for **installed `tkit version`**, not GitHub `latest`.
2. For canary CLI, prefer `step_catalog.file` or stable release matching the **semver base** (segment before `-canary.`).
3. Cache under `.testflowkit/cache/step-definitions-<version>.json`.

### Optional follow-ups (still Phase 0 / release hygiene)

- [ ] Confirm asset on next semantic-release tag (manual smoke check).
- [ ] Mention catalog in release notes / CONTRIBUTING.
- [ ] `step-definitions.schema.json` on release (moved to Phase 4 if deferred).

---

## Phase 1 — IDE agent (Cursor / VS Code Copilot)

**Status:** Next — current priority

### Goal

A developer opens a TestFlowKit test project in **Cursor** or **VS Code (Copilot + MCP)** and can prompt: *“Write a test for user registration”* and receive **valid Gherkin** that uses only sentences from the resolved catalog and references that exist in `config.yml`.

### Architecture

```text
┌─────────────────────────────────────────────────────────┐
│  Cursor / VS Code                                       │
│  ┌──────────────┐    ┌─────────────────────────────┐   │
│  │ Cursor rules │    │ MCP: testflowkit-mcp        │   │
│  │ (guardrails) │    │ read config / features      │   │
│  └──────────────┘    │ get/search step catalog     │   │
│                      └──────────────┬──────────────┘   │
└─────────────────────────────────────┼───────────────────┘
                                      │
          testflowkit.agent.yml       │  fetch if needed
          config.yml                  ▼
          *.feature              GitHub Release
                                 step-definitions.json
```

### 1.1 Dedicated agent config file

**Deliverable:** `testflowkit.agent.yml` at the **test project** root (separate from runtime `config.yml`).

```yaml
version: 1

project:
  test_config: "./config.yml"
  features_glob: "e2e/features/**/*.feature"

step_catalog:
  # Optional — omit to auto-fetch from release
  # file: "./build/step-definitions.json"
  # url: "https://..."

agent:
  default_tags_for_draft: "@wip @ai-generated"
  run_command: "tkit run -c config.yml --tags @wip"
```

**Checklist**

- [ ] Specification for all fields (paths, env interpolation, canary behavior).
- [ ] `testflowkit.agent.example.yml` in repo or init boilerplate.
- [ ] Discovery: MCP loads from workspace root (or explicit `-c` path).

### 1.2 Cursor / Copilot guardrails

**Deliverable:** Rule template (e.g. `.cursor/rules/testflowkit.mdc`) for consumer test repos.

| Rule | Detail |
|------|--------|
| Sentences | Only strings present in the resolved catalog |
| API steps | `I prepare a request to "api.operation"` — both parts must exist in `config.yml` |
| Drafts | Tag new scenarios `@wip` (and optionally `@ai-generated`) |
| Secrets | Never put tokens/passwords in features; use `env` / security schemes |
| Macros | `@macro` on definition; `${var}` placeholders; invocation via data table |
| Missing step | Emit a **sentence suggestion** block; do not invent executable steps |

**VS Code:** Mirror the same content in `github.copilot-instructions.md` (or equivalent).

**Checklist**

- [ ] Rule file in TestFlowKit repo (copy-paste template).
- [ ] “Enable the agent” section in README or QA guide.

### 1.3 MCP server (v0)

**Deliverable:** `testflowkit-mcp` package (Node or Go), stdio transport, editor-agnostic.

| Tool | Description |
|------|-------------|
| `get_step_catalog` | Resolve catalog (see resolution order below); return JSON + metadata (`version`, `source`) |
| `search_sentences` | Keyword / category filter over catalog |
| `read_test_config` | Parse linked `config.yml` (APIs, operations, endpoints, pages, elements) |
| `list_features` | Glob `features_glob` |
| `read_feature` | Read one `.feature` file |
| `write_feature` | Create/update feature (respect `features_glob` roots) |

**Catalog resolution order**

1. `step_catalog.file`
2. `step_catalog.url`
3. Installed `tkit version` → stable release URL (strip `-canary.*` to semver base if needed)
4. Cache at `.testflowkit/cache/step-definitions-<version>.json`
5. Warn if CLI version and catalog version disagree (major/minor → error; patch → warning)

**Checklist**

- [ ] MCP server repository or `packages/testflowkit-mcp/` in monorepo.
- [ ] Cursor `mcp.json` example.
- [ ] VS Code MCP configuration example.
- [ ] Reads `testflowkit.agent.yml` before tools run.
- [ ] Network fetch with sensible timeout; cache on disk.

### 1.4 Out of scope for Phase 1

- GraphQL introspection / OpenAPI import (Phase 2).
- Automatic `tkit run` and failure-driven iteration (Phase 3).
- Automatic `config.yml` patches (Phase 2–3).
- Publishing catalog on canary GitHub Release (decided: CLI-only canary).

### 1.5 Acceptance criteria

| # | Criterion | Verification |
|---|-----------|--------------|
| 1 | Agent config via dedicated file | `testflowkit.agent.yml` documented and loadable by MCP |
| 2 | Catalog without local file | Omit `step_catalog.file` → MCP fetches stable release for `tkit version` |
| 3 | Read feature files | `read_feature` / `list_features` against project glob |
| 4 | Create feature files | `write_feature` produces valid Gherkin structure |
| 5 | Sentence suggestions | Rule + tool behavior when no catalog match |
| 6 | Registration-style prompt | Manual eval: plausible feature using frontend/API steps from catalog |
| 7 | Cursor + VS Code | Same MCP server; editor-specific setup docs only |

---

## Phase 2 — Schema-aware generation

**Status:** Planned

### Goal

Agent understands **GraphQL** and **REST** contracts from **file or URL**, proposes `config.yml` snippets, and generates tests that reference real operations/endpoints.

### 2.1 Extend `testflowkit.agent.yml`

```yaml
schemas:
  graphql:
    - name: main_api
      maps_to_api: "main_graphql"
      source:
        endpoint: "https://api.example.com/graphql"   # introspection
        # file: "./schemas/schema.graphql"
        # url: "https://cdn.example.com/schema.graphql"
      auth:
        bearer: "{{ env.GRAPHQL_TOKEN }}"
      cache:
        path: ".testflowkit/cache/graphql/main_api.json"
        ttl_hours: 24

  openapi:
    - name: users_api
      maps_to_api: "users_api"
      source:
        file: "./api/openapi.yaml"
        # url: "https://api.example.com/openapi.json"
      cache:
        path: ".testflowkit/cache/openapi/users_api.json"
```

**Source rule:** exactly one of `file`, `url`, or (GraphQL only) `endpoint` per entry.

### 2.2 MCP tools

| Tool | Description |
|------|-------------|
| `resolve_graphql_schema` | file \| url \| endpoint → cached schema summary |
| `resolve_openapi_schema` | file \| url → paths, methods, models summary |
| `propose_config_snippet` | YAML fragment for `apis.definitions` (dry-run, no auto-apply by default) |
| `refresh_schema_cache` | Force re-fetch / re-introspect |

### 2.3 Checklist

- [ ] GraphQL introspection implementation (or SDL parse for `file`).
- [ ] OpenAPI 3 loader (Swagger 2 optional).
- [ ] Auth on fetch (bearer, basic, apikey).
- [ ] `.testflowkit/cache/` gitignored in consumer docs.
- [ ] Agent rules: do not invent operations not in config or resolved schema.

### 2.4 Acceptance criteria (from product framing)

| Criterion | Phase 2 scope |
|-----------|----------------|
| Registration test | GraphQL mutation + optional UI steps |
| Connection through / non-through | Two scenarios or outline; document “through” in `domain:` section of agent config |
| Task project workflow | Multi-step API + assertions using stored variables |

---

## Phase 3 — Run loop, macros, and golden tests

**Status:** Planned

### Goal

Close the loop: generate → **run** → read failures → fix → repeat. Treat macros and config updates as first-class agent capabilities.

### 3.1 MCP tools

| Tool | Description |
|------|-------------|
| `run_tests` | Wrap `tkit run` (tags, paths from agent config) |
| `validate_features` | Wrap `tkit validate` |
| `parse_last_report` | Extract failed steps / messages from HTML or JSON reporter |
| `write_macro` | Create `@macro` scenario + document invocation pattern |
| `patch_test_config` | Apply approved YAML changes to `config.yml` (with dry-run diff) |

### 3.2 Agent config

```yaml
agent:
  max_run_iterations: 5
  run_command: "tkit run -c config.yml --tags @wip"
  report_format: "json"   # or html — for parse_last_report
```

### 3.3 Golden prompts (regression suite for the agent)

| Prompt | Expected outcome |
|--------|------------------|
| “I want to test the registration” | Valid feature + config refs or explicit missing-config report |
| “I want to test the connection for through and non-through cases” | Two scenarios or scenario outline + tags |
| “Test a workflow for modifying a task project…” | Multi-step flow with variables and assertions |
| “Write a macro for login” | `@macro` feature + example invocation |

Store as fixtures under `docs/plans/golden-prompts/` or `e2e/agent-eval/` (TBD).

### 3.4 Missing-sentence workflow

When the model cannot map intent to the catalog, emit structured output:

```yaml
missing_sentence:
  intent: "assert websocket message received"
  closest_matches:
    - "the response should contain {string}"
  proposed_sentence: "the websocket should receive a message containing {string}"
  proposed_category: "backend"
```

Contributors implement new Go steps; catalog refreshes on next release.

### 3.5 Checklist

- [ ] `run_tests` + `validate_features` in MCP.
- [ ] Reporter parsing (failures → structured JSON for the model).
- [ ] Iteration limit and `@wip` isolation documented.
- [ ] Golden prompt evals (manual or semi-automated).
- [ ] Macro authoring rules in Cursor rule file.

### 3.6 Acceptance criteria

| Criterion | Phase 3 |
|-----------|---------|
| Run tests via agent | `run_tests` executes CLI, returns exit code + summary |
| Iterate on failure | At least one fix loop documented in QA guide |
| Write macros | Agent produces valid `@macro` + invocation |
| Update configuration | `patch_test_config` with human-approved diff |
| Sentence suggestions | Structured `missing_sentence` output |

---

## Phase 4 — Hardening and distribution

**Status:** Planned

### Goal

Production-grade catalog distribution, validation, and enterprise-friendly overrides.

### 4.1 Release artifacts

- [ ] `step-definitions.schema.json` on stable release.
- [ ] `checksums.txt` (SHA256 for catalog).
- [ ] Optional: `step-definitions.json.gz` for large catalogs.

### 4.2 npm / offline

- [ ] Bundle `step-definitions.json` in CLI npm package (stable + optional canary).
- [ ] MCP prefers bundled catalog when `node_modules/@testflowkit/cli` version matches.

### 4.3 Enterprise

- [ ] `step_catalog.url` for internal mirrors (GitHub Enterprise, Artifactory).
- [ ] `GITHUB_TOKEN` / custom headers for private release assets.
- [ ] Document air-gapped: commit `step_catalog.file` or internal mirror only.

### 4.4 Catalog format v2 (optional)

Wrap array with metadata:

```json
{
  "version": "1.2.4",
  "generated_at": "2026-05-23T12:00:00Z",
  "steps": [ ... ]
}
```

Breaking change — coordinate with MCP and export script.

### 4.5 Checklist

- [ ] Schema + checksum on release pipeline.
- [ ] CONTRIBUTING: how to add a step + refresh catalog.
- [ ] Version mismatch policy documented (CLI vs catalog).
- [ ] Re-evaluate canary catalog on GitHub (rolling `canary` release) if demand appears.

---

## Cross-cutting concerns

### Security

| Topic | Guidance |
|-------|----------|
| API keys in agent config | Use `{{ env.* }}`; never commit secrets |
| MCP file writes | Restrict to `features_glob` and explicit config path |
| URL fetch | Timeouts, size limits, no SSRF to internal networks without opt-in |

### Version alignment

| Situation | Behavior |
|-----------|----------|
| Stable CLI `1.2.4` | Fetch release `1.2.4` catalog |
| Canary CLI `3.6.1-canary.abc1234` | Use `3.6.1` stable catalog, local export, or `step_catalog.file` |
| PR adding new steps | `make export-step-definitions` + `step_catalog.file` in agent config |

### Repository layout (target)

```text
testflowkit/                          # framework repo
├── docs/plans/
│   ├── testflowkit-agent-roadmap.md  # this file
│   └── publish-step-definitions-on-release.md
├── scripts/step_definitions_export/
└── packages/testflowkit-mcp/         # Phase 1 (TBD)

consumer-test-project/
├── config.yml
├── testflowkit.agent.yml
├── .cursor/rules/testflowkit.mdc
├── .testflowkit/cache/
└── e2e/features/
```

---

## Progress tracker (master checklist)

### Phase 0 — Foundation

- [x] `export-step-definitions` Makefile target
- [x] Upload on stable `release.yml`
- [x] PR + post-release JSON validation
- [x] README download URLs
- [ ] Manual verify on next release tag
- [ ] CONTRIBUTING / release notes mention

### Phase 1 — IDE agent

- [ ] `testflowkit.agent.yml` spec + example
- [ ] Cursor rules template
- [ ] VS Code Copilot instructions template
- [ ] MCP v0 (`get_step_catalog`, config/features R/W)
- [ ] Editor setup documentation
- [ ] Phase 1 acceptance criteria signed off

### Phase 2 — Schema-aware

- [ ] `schemas.graphql` / `schemas.openapi` in agent config
- [ ] Introspection + OpenAPI loaders
- [ ] `propose_config_snippet`
- [ ] Product golden prompts (schema-heavy)

### Phase 3 — Run loop

- [ ] `run_tests` / `validate_features`
- [ ] Report parsing + iteration
- [ ] Macros + config patch tools
- [ ] Missing-sentence report format
- [ ] Golden prompt regression suite

### Phase 4 — Hardening

- [ ] schema.json + checksums on release
- [ ] Optional npm-bundled catalog
- [ ] Enterprise mirror documentation
- [ ] Catalog format v2 (if needed)

---

## References

| Resource | Path |
|----------|------|
| Step export | `scripts/step_definitions_export/main.go` |
| Release workflow | `.github/workflows/release.yml` |
| Canary workflow | `.github/workflows/publish-canary.yml` |
| Publish plan | [publish-step-definitions-on-release.md](./publish-step-definitions-on-release.md) |
| Step browser (human) | Documentation site `/sentences` |
| Macros | `documentation/content/docs/3.features/2.macros.md` |
| API testing | `documentation/content/docs/3.features/4.api-testing.md` |
