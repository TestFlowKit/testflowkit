# Plan: Publish step definitions on GitHub Release

This document describes how to ship the TestFlowKit **step catalog** (`step-definitions.json`) as a **GitHub Release asset** on every stable release, and how **canary** builds should expose the same catalog without creating a release per commit.

It supports the future **IDE agent** (Cursor / VS Code MCP), which can fetch the catalog when `testflowkit.agent.yml` does not specify a local file.

---

## Goals

| Goal | Description |
|------|-------------|
| **Zero-config agent** | Consumers can omit a local catalog; the agent resolves it from the release matching the installed CLI version. |
| **Version alignment** | Catalog for tag `v1.2.4` matches the step implementations in that tag. |
| **Stable channel** | One `step-definitions.json` per semver GitHub Release (alongside platform ZIPs). |
| **Canary channel** | Latest main-line catalog available without per-SHA releases. |
| **Reproducibility** | Same export command used locally, in CI, and on release. |

## Non-goals (this iteration)

- Publishing YAML instead of JSON (JSON is the export format today).
- Embedding the catalog inside every platform ZIP (optional later).
- Implementing the MCP server or `testflowkit.agent.yml` (separate work; this plan only defines the release artifact contract).

---

## Current state

### Export pipeline (already exists)

| Piece | Location |
|-------|----------|
| Export command | `go run ./scripts/step_definitions_export/main.go` |
| Default output | `./step-definitions.json` (gitignored via `**/step-definitions.json`) |
| Data source | `scripts/shared.GetAllDocs()` → all registered Go steps |
| Doc site generator | `go run ./scripts/doc_generator/main.go` (same source, different output layout) |

### JSON entry shape (per step)

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

### Release pipeline (stable)

| Workflow | File | Behavior |
|----------|------|----------|
| Release | `.github/workflows/release.yml` | `semantic-release` creates tag → `make releases` → `gh release upload "$LATEST_TAG" build/*.zip` |
| NPM publish | `.github/workflows/publish-npm-package.yml` | Downloads `*.zip` from latest release, publishes npm `latest` |

**Gap:** Release assets today are **binaries only**; `step-definitions.json` is not generated or uploaded.

### Canary pipeline

| Workflow | File | Behavior |
|----------|------|----------|
| Canary | `.github/workflows/publish-canary.yml` | Version `"{last-tag}-canary.{short-sha}"` → `make releases` → workflow artifact (1 day) → npm `canary` tag |

**Gap:** No GitHub Release and no published catalog for canary.

---

## Target artifact

| Property | Value |
|----------|--------|
| **Filename** | `step-definitions.json` |
| **Content-Type** | `application/json` |
| **Produced by** | `go run ./scripts/step_definitions_export/main.go -output-file build/step-definitions.json` |
| **Attached to** | GitHub Release (stable + rolling canary) |

### Optional v1.1 enhancements

| Asset | Purpose |
|-------|---------|
| `step-definitions.schema.json` | JSON Schema for catalog entries (validation in MCP) |
| `checksums.txt` | SHA256 of `step-definitions.json` (+ binaries if desired) |

---

## Published artifacts (implemented)

| Channel | URL pattern | Example |
|---------|-------------|---------|
| **Stable** | `https://github.com/TestFlowKit/testflowkit/releases/download/<tag>/step-definitions.json` | Tag `1.2.4` (no `v` prefix, per semantic-release) |
| **Canary** | `https://github.com/TestFlowKit/testflowkit/releases/download/canary/step-definitions.json` | Rolling pre-release, updated on each `main` push |

**Local export** (gitignored output):

```bash
make export-step-definitions
# → build/step-definitions.json
```

**Agent rule:** Resolve the catalog for the **installed CLI version** (`tkit version`), not GitHub `latest`. Use the canary URL only when `step_catalog.channel: canary` (future agent spec).

---

## Download URL contract (for IDE agent)

Stable release (tag matches semantic-release `tagFormat`: `${version}`, e.g. `1.2.4` without a `v` prefix):

```text
https://github.com/TestFlowKit/testflowkit/releases/download/<tag>/step-definitions.json
```

Canary rolling pre-release:

```text
https://github.com/TestFlowKit/testflowkit/releases/download/canary/step-definitions.json
```

Agent resolution order (documented in future `testflowkit.agent.yml` spec):

1. `step_catalog.file` (local path)
2. `step_catalog.url` (explicit override)
3. `step_catalog.release.version` or installed `tkit version` → GitHub Release URL above
4. Cache under `.testflowkit/cache/step-definitions-<version>.json`

**Rule:** Default fetch version = **installed CLI version**, not `latest` on GitHub.

---

## Stable release — CI changes

### 1. Makefile target

Add a phony target so CI and developers share one command:

```makefile
.PHONY: export-step-definitions
export-step-definitions:
	@mkdir -p $(BUILD_DIR)
	go run ./scripts/step_definitions_export/main.go -output-file $(BUILD_DIR)/step-definitions.json
```

### 2. Update `.github/workflows/release.yml`

In the **Upload Release Assets** step (after `make releases`):

1. Setup Go (already present).
2. Run `make export-step-definitions` (or equivalent `go run`).
3. Upload JSON with the same tag as ZIPs:

```bash
gh release upload "$LATEST_TAG" build/step-definitions.json --clobber
```

Optional: append SHA256 to `build/checksums.txt` and upload that file too.

### 3. Release notes

Semantic-release / changelog should mention:

- New asset: **Step definitions catalog** for AI/agent tooling.

### 4. Verification job (recommended)

After upload, a lightweight check on the release tag:

- Download `step-definitions.json`.
- Assert valid JSON array, non-empty.
- Assert required fields on first entry: `sentence`, `description`, `categories`.
- Optional: minimum entry count threshold vs previous release (regression guard).

---

## Canary channel — recommended approach

**Do not** create a new GitHub Release per commit (`1.2.3-canary.abc1234`).

Use a **single rolling pre-release** named `canary`:

| Property | Value |
|----------|--------|
| Tag / release name | `canary` (or `v0.0.0-canary` if tags must be semver-shaped) |
| Pre-release flag | `true` |
| Update cadence | Every successful `publish-canary.yml` run on `main` |
| Assets | `step-definitions.json` (+ optionally latest canary ZIPs) |
| Upload | `gh release upload canary build/step-definitions.json --clobber` |

Agent config (future):

```yaml
step_catalog:
  channel: canary   # resolves to rolling release "canary"
```

### Alternative (if avoiding GitHub releases for canary)

Bundle `step-definitions.json` in the **npm canary package** next to CLI binaries (`publish-canary.yml` / `_publish-packages.yml`). MCP reads from `node_modules/...` when `tkit version` contains `-canary`.

| Approach | Pros | Cons |
|----------|------|------|
| Rolling `canary` release | Same URL model as stable; works without npm | Needs `contents: write` on canary workflow |
| npm-bundled only | Matches existing canary install path | No direct URL; npm required |

**Recommendation:** Rolling **`canary` pre-release** for catalog; keep binary distribution as today (artifact → npm). Catalog is small and benefits from a stable HTTPS URL.

### Canary workflow changes (`.github/workflows/publish-canary.yml`)

1. Grant `contents: write` (only if creating/updating releases).
2. After build, run `export-step-definitions`.
3. Ensure release `canary` exists (`gh release create canary --prerelease` if missing).
4. `gh release upload canary build/step-definitions.json --clobber`.

---

## NPM publish workflow impact

`.github/workflows/publish-npm-package.yml` currently downloads only `*.zip`.

| Option | Action |
|--------|--------|
| **A — No change** | Catalog consumed only via GitHub Release URL or agent cache. |
| **B — Include in npm** | Also download `step-definitions.json` and pack into CLI npm package (better offline). |

**Recommendation for v1:** Option A (release asset only). Option B as follow-up if npm installs should work offline without GitHub.

---

## Version mismatch handling (agent / MCP)

When CLI reports `1.2.3` and catalog is fetched for `1.2.4` (or vice versa):

| Severity | Condition | Action |
|----------|-----------|--------|
| Warning | Patch mismatch, same minor | Log warning, continue |
| Error | Major/minor mismatch | Block auto-fetch; require `step_catalog.file` or pin `version` |

Document that **contributors on a branch** should run local export:

```bash
make export-step-definitions
```

and set `step_catalog.file: "./build/step-definitions.json"` in `testflowkit.agent.yml`.

---

## Local developer workflow

```bash
# Generate catalog locally (gitignored)
make export-step-definitions
# or
go run ./scripts/step_definitions_export/main.go -output-file ./step-definitions.json
```

Use for:

- Agent development before a release exists.
- PRs that add new steps (catalog newer than last release).

---

## Implementation checklist

### Phase 1 — Stable release (MVP)

- [x] Add `export-step-definitions` Makefile target.
- [x] Extend `release.yml` upload step with `build/step-definitions.json`.
- [ ] Manually verify asset on next semantic-release tag.
- [x] Document download URL in agent plan / README (link to this doc).

### Phase 2 — Quality & discoverability

- [x] Post-upload JSON validation in CI.
- [ ] Add `step-definitions.schema.json` (optional).
- [ ] Add `checksums.txt` for catalog SHA256 (optional).
- [ ] Mention new asset in release notes template / CONTRIBUTING.

### Phase 3 — Canary

- [x] Update `publish-canary.yml` to export and upload to rolling `canary` pre-release.
- [x] Document `channel: canary` URL in agent spec (see Published artifacts above).
- [x] Decide permissions (`contents: write`) and release creation idempotency.

### Phase 2 (continued) — PR CI

- [x] Validate export on `pull-request.yml` (array, non-empty, required fields).

### Phase 4 — Consumer integration (out of band)

- [ ] MCP `get_step_catalog` with resolution order (file → url → release).
- [ ] `testflowkit.agent.yml` `step_catalog` section.
- [ ] Version check against `tkit version`.

---

## Risks and mitigations

| Risk | Mitigation |
|------|------------|
| Catalog out of sync with binary | Generate in same job, same commit, same tag before upload |
| Agent uses wrong version | Pin to CLI version; never default to GitHub `latest` |
| Canary URL moves every push | Fixed release name `canary` + `--clobber` |
| Large JSON over time | Compress optional (`step-definitions.json.gz`); MCP accepts both |
| Private GitHub org | `step_catalog.url` + token via env in agent config |

---

## Open decisions

| # | Question | Proposal |
|---|----------|----------|
| 1 | Tag format in URL: `v1.2.4` vs `1.2.4`? | Match `git describe` / semantic-release output exactly (verify in first release). |
| 2 | Include catalog version inside JSON? | Add top-level wrapper in v2: `{ "version": "1.2.4", "steps": [...] }` (breaking change; defer or version file). |
| 3 | Canary: GitHub rolling release vs npm bundle? | Rolling `canary` pre-release for catalog. |
| 4 | Commit catalog to repo? | No — release + local export only (stays gitignored). |

---

## References

| Resource | Path |
|----------|------|
| Export script | `scripts/step_definitions_export/main.go` |
| Shared docs source | `scripts/shared/main.go` |
| Stable release workflow | `.github/workflows/release.yml` |
| Canary workflow | `.github/workflows/publish-canary.yml` |
| NPM publish | `.github/workflows/publish-npm-package.yml` |
| Gitignore entry | `**/step-definitions.json` |

---

## Related documents (to create later)

- `testflowkit.agent.yml` specification (step catalog resolution, GraphQL/OpenAPI sources).
- IDE agent architecture (MCP tools + Cursor rules).
- MCP server design (`get_step_catalog`, cache paths).
