# Auth & Security – Implementation Traceability

> **Purpose**: Living document tracking what is implemented, what is pending, and key design decisions for the enterprise auth/security framework introduced via the Tech Framing document (March 2026).
> Keep this file up-to-date whenever a phase is completed or a decision changes.

---

## Baseline (before this feature)

The only auth mechanism in the codebase was manual header injection via `default_headers` at the API-definition level:

```yaml
apis:
  definitions:
    my_api:
      base_url: "{{ env.API_URL }}"
      default_headers:
        Authorization: "Bearer {{ env.TOKEN }}"
```

No security scheme registry, no token persistence, no inheritance hierarchy, no proxy-aware transport.

---

## Architecture Overview

```
config.yml
├── security_schemes:          ← global registry (SecurityScheme objects)
│     my_idp: { type, credentials, proxy_url, persist, … }
├── default_security: my_idp  ← project-level fallback
└── apis:
      definitions:
        my_api:
          security_ref: my_idp        ← API-level override (reference)
          security_overrides:         ← API-level partial overrides
            scopes: ["inventory.all"]
          endpoints:
            get_item:
              security: none          ← endpoint-level: disables all inheritance
              # or:
              security_ref: other_idp ← endpoint-level: full override
```

**Inheritance resolution order (highest priority first):**

1. Endpoint / Operation level (`security` or `security_ref`)
2. API definition level (`security` / `security_ref` + `security_overrides`)
3. Project root `default_security`

`security: none` at any level disables inheritance from all lower-priority levels.

**Token persistence key**: SHA256 of the fully-resolved (env-substituted) SecurityScheme + any active overrides.

---

## Implementation Status

### ✅ Phase 1 – Config Schema & Validation

| Item | Status | Location |
|------|--------|----------|
| `SecuritySchemeType` enum (bearer, basic, apikey, oauth2, oidc, certificate, none) | ✅ Done | `internal/config/types.go` |
| `SecurityScheme` struct (type, token/credentials, scopes, proxy_url, persist, duration, retry_on_401) | ✅ Done | `internal/config/types.go` |
| `SecurityRef` struct (name + inline SecurityScheme) | ✅ Done | `internal/config/types.go` |
| `SecurityOverrides` struct (scopes, audience, proxy_url) | ✅ Done | `internal/config/types.go` |
| `APIKeyPlacement` enum (header, query, cookie) | ✅ Done | `internal/config/types.go` |
| `OAuth2TokenAuthMethod` enum (`client_secret_post`, `client_secret_basic`) | ✅ Done | `internal/config/security_types.go` |
| `SecurityScheme.TokenEndpointAuthMethod` (mandatory for oauth2) | ✅ Done | `internal/config/security_types.go` |
| `Config.SecuritySchemes` map (root registry) | ✅ Done | `internal/config/main_type.go` |
| `Config.DefaultSecurity` string reference | ✅ Done | `internal/config/main_type.go` |
| `APIDefinition.SecurityRef` | ✅ Done | `internal/config/types.go` |
| `APIDefinition.SecurityOverrides` | ✅ Done | `internal/config/types.go` |
| `Endpoint.SecurityRef` | ✅ Done | `internal/config/types.go` |
| `GraphQLOperation.SecurityRef` | ✅ Done | `internal/config/types.go` |
| Validation: scheme type required fields | ✅ Done | `internal/config/main_type.go` |
| Validation: dangling `security_ref` references | ✅ Done | `internal/config/main_type.go` |
| Validation: `security_overrides` shape | ✅ Done | `internal/config/main_type.go` |
| `Config.GetSecurityScheme()` helper | ✅ Done | `internal/config/main_type.go` |
| Backward compat: existing `default_headers` still works | ✅ Done | no change to DefaultHeaders |

### ✅ Phase 2 – Resolver & Canonical Hash

| Item | Status | Location |
|------|--------|----------|
| `internal/security/resolver.go` – `Resolve()` with full precedence chain | ✅ Done | `internal/security/resolver.go` |
| `ResolvedSecurity` result type (effective scheme + overrides merged) | ✅ Done | `internal/security/resolver.go` |
| `internal/security/hash.go` – canonical JSON serialisation + SHA256 | ✅ Done | `internal/security/hash.go` |
| Sorted-key deterministic serialisation (env-invalidation guarantee) | ✅ Done | `internal/security/hash.go` |

### ✅ Phase 3 – testflowkit.lock Manager

| Item | Status | Location |
|------|--------|----------|
| `internal/state/manager.go` – Load / Save / Get / Put / Invalidate | ✅ Done | `internal/state/manager.go` |
| JSON schema with `version`, `updated_at`, per-entry metadata | ✅ Done | `internal/state/manager.go` |
| 30-second safety buffer on expiry checks | ✅ Done | `internal/state/manager.go` |
| In-process `sync.RWMutex` for goroutine safety | ✅ Done | `internal/state/manager.go` |
| Cross-process file locking (O_EXCL sidecar `.lock`) | ✅ Done | `internal/state/manager.go` |
| `Invalidate(hash)` for retry_on_401 flow | ✅ Done | `internal/state/manager.go` |
| Lock file path defaults to `testflowkit.lock` next to `config.yml` | ✅ Done | `internal/state/manager.go` |

### ✅ Phase 4 – Auth Providers

| Item | Status | Location |
|------|--------|----------|
| `Provider` interface (`Authenticate(ctx, scheme) → TokenResult`) | ✅ Done | `internal/security/providers/provider.go` |
| `ProviderFactory` (keyed by SecuritySchemeType) | ✅ Done | `internal/security/providers/factory.go` |
| Bearer static provider | ✅ Done | `internal/security/providers/bearer.go` |
| Basic auth provider | ✅ Done | `internal/security/providers/basic.go` |
| API key provider (header / query / cookie placement) | ✅ Done | `internal/security/providers/apikey.go` |
| OAuth2 client_credentials provider (with proxy support) | ✅ Done | `internal/security/providers/oauth2.go` |
| OAuth2 strategy pattern: `tokenEndpointAuthHandler` interface (`client_secret_post` / `client_secret_basic`) | ✅ Done | `internal/security/providers/oauth2.go` |
| OIDC provider stub (returns `ErrNotImplemented`, reserves interface) | ✅ Done | `internal/security/providers/oidc.go` |
| Certificate / mTLS provider stub (returns `ErrNotImplemented`) | ✅ Done | `internal/security/providers/certificate.go` |

### ✅ Phase 5 – Transport, Proxy & Retry-on-401

| Item | Status | Location |
|------|--------|----------|
| `internal/httpauth/transport.go` – `AuthTransport` (RoundTripper) | ✅ Done | `internal/httpauth/transport.go` |
| Auth header/token injection via `AuthTransport` | ✅ Done | `internal/httpauth/transport.go` |
| Per-scheme `proxy_url` support inside transport | ✅ Done | `internal/httpauth/transport.go` |
| Opt-in `retry_on_401`: disabled by default; one invalidate + one retry | ✅ Done | `internal/httpauth/transport.go` |
| REST adapter refactored to use `AuthTransport` | ✅ Done | `internal/step_definitions/api/protocol/restapi_adapter.go` |
| GraphQL adapter wired to same transport via `WithHTTPClient` option | ✅ Done | `internal/step_definitions/api/protocol/graphql_adapter.go` |
| Resolver called during `PrepareRequest` and result stored in scenario ctx | ✅ Done | `internal/step_definitions/backend/commonbackendsteps/prepare_request.go` |
| `BackendContext.ResolvedSecurity` field | ✅ Done | `internal/step_definitions/core/scenario/backend_context_types.go` |

### ✅ Phase 6 – CLI Lifecycle & Boilerplate

| Item | Status | Location |
|------|--------|----------|
| `actionrun` initialises lock manager before suite start | ✅ Done | `internal/actions/actionrun/main.go` |
| `actionrun` saves lock manager after suite end | ✅ Done | `internal/actions/actionrun/main.go` |
| `actioninit` creates empty `testflowkit.lock` alongside `config.yml` | ✅ Done | `internal/actions/actioninit/main.go` |
| `config.boilerplate.yml` security section comments | ✅ Done | `internal/actions/actioninit/boilerplate/config.boilerplate.yml` |

### ✅ Phase 7 – Test Coverage

| Item | Status | Location |
|------|--------|----------|
| `internal/security/resolver_test.go` – precedence, none sentinel, env invalidation | ✅ Done | `internal/security/resolver_test.go` |
| `internal/security/hash_test.go` – determinism, env-driven key change | ✅ Done | `internal/security/hash_test.go` |
| `internal/state/manager_test.go` – expiry buffer, corruption recovery, concurrency | ✅ Done | `internal/state/manager_test.go` |
| `internal/config/` – security validation tests added | ✅ Done | `internal/config/types_test.go` |
| `internal/security/providers/*_test.go` – each provider | ✅ Done | `internal/security/providers/` |
| `internal/httpauth/transport_test.go` – inject, proxy, retry_on_401 | ✅ Done | `internal/httpauth/transport_test.go` |

---

## Deferred (out of scope for this implementation pass)

| Feature | Reason | Expected phase |
|---------|--------|---------------|
| OIDC / PKCE runtime flow | High scope, browser-less or device-code auth flow needed | Future phase |
| Certificate / mTLS provider (`.p12`, `.pem`) | Low priority for current users | Future phase |
| BDD Gherkin steps for role-switch / force-refresh | Scope decision (runtime/config only for now) | Future phase |
| Token encryption at rest in `testflowkit.lock` | Plain JSON sufficient for current needs; docs warning present | Future phase |
| `security: none` on a `security_overrides`-only entry | Edge case; not validated yet | Future phase |

---

## Scope Decisions Log

| # | Decision | Date |
|---|----------|------|
| 1 | Canonical API override key is `security_overrides` (not `scope_overrides`) | 2026-03-16 |
| 2 | Tokens persisted as plain JSON in `testflowkit.lock` with CI cache guidance in docs | 2026-03-16 |
| 3 | OIDC/PKCE deferred; provider interface reserved with explicit not-implemented error | 2026-03-16 |
| 4 | Certificate provider deferred; same interface stub strategy | 2026-03-16 |
| 5 | BDD auth steps (role switch, force refresh) deferred to next implementation pass | 2026-03-16 |
| 6 | `default_headers` backward compatibility preserved with zero changes to DefaultHeaders behaviour | 2026-03-16 |
| 7 | Safety buffer for token expiry is 30 seconds | 2026-03-16 |
| 8 | retry_on_401 is opt-in per scheme, disabled by default, max 1 re-auth + 1 retry | 2026-03-16 |

---

## Key Config Reference

```yaml
# Project-level security registry
security_schemes:
  enterprise_idp:
    type: oauth2
    grant_type: client_credentials
    token_url: "{{ env.AUTH_URL }}"
    token_endpoint_auth_method: client_secret_post   # required: client_secret_post | client_secret_basic
    client_id: "{{ env.CLIENT_ID }}"
    client_secret: "{{ env.CLIENT_SECRET }}"
    scopes: ["read", "write"]
    proxy_url: "http://proxy.internal:8080"   # optional per-scheme proxy
    persist: true
    duration: "1h"
    retry_on_401: true

  admin_key:
    type: apikey
    key: "{{ env.ADMIN_KEY }}"
    placement: header         # header | query | cookie
    header_name: "X-Api-Key"

default_security: "enterprise_idp"   # project-level fallback

apis:
  definitions:
    inventory_api:
      base_url: "{{ env.API_URL }}"
      security_ref: enterprise_idp        # API-level – inherits + allows overrides
      security_overrides:
        scopes: ["inventory.all"]         # override scopes for this API only

    public_api:
      base_url: "{{ env.PUBLIC_URL }}"
      # no security_ref → falls back to default_security

    admin_api:
      base_url: "{{ env.ADMIN_URL }}"
      security_ref: admin_key

      endpoints:
        health_check:
          method: GET
          path: /health
          description: Public health endpoint
          security: none                  # disables all inherited security
```

---

## Files Created / Modified

| Action | Path |
|--------|------|
| Created | `internal/config/security_types.go` |
| Modified | `internal/config/types.go` |
| Modified | `internal/config/main_type.go` |
| Modified | `internal/config/loader.go` |
| Created | `internal/security/resolver.go` |
| Created | `internal/security/resolver_test.go` |
| Created | `internal/security/hash.go` |
| Created | `internal/security/hash_test.go` |
| Created | `internal/state/manager.go` |
| Created | `internal/state/manager_test.go` |
| Created | `internal/security/providers/provider.go` |
| Created | `internal/security/providers/factory.go` |
| Created | `internal/security/providers/bearer.go` |
| Created | `internal/security/providers/basic.go` |
| Created | `internal/security/providers/apikey.go` |
| Created | `internal/security/providers/oauth2.go` |
| Created | `internal/security/providers/oidc.go` |
| Created | `internal/security/providers/certificate.go` |
| Created | `internal/httpauth/transport.go` |
| Created | `internal/httpauth/transport_test.go` |
| Modified | `internal/step_definitions/api/protocol/restapi_adapter.go` |
| Modified | `internal/step_definitions/api/protocol/graphql_adapter.go` |
| Modified | `internal/step_definitions/backend/commonbackendsteps/prepare_request.go` |
| Modified | `internal/step_definitions/core/scenario/backend_context_types.go` |
| Modified | `internal/actions/actionrun/main.go` |
| Modified | `internal/actions/actioninit/main.go` |
| Modified | `internal/actions/actioninit/boilerplate/config.boilerplate.yml` |
| Created | `auth-security-traceability.md` (this file) |

---

## CI/CD Cache Guidance

To share the `testflowkit.lock` across parallel matrix jobs and avoid hitting auth rate limits (Auth0/Okta 429):

**GitHub Actions**
```yaml
- uses: actions/cache@v4
  with:
    path: testflowkit.lock
    key: tfk-lock-${{ vars.ENVIRONMENT }}-v1-${{ steps.week.outputs.value }}
    restore-keys: tfk-lock-${{ vars.ENVIRONMENT }}-v1-
```

**GitLab CI**
```yaml
cache:
  key: tfk-lock-${CI_ENVIRONMENT_NAME}
  paths:
    - testflowkit.lock
```

> ⚠️ **Security note**: `testflowkit.lock` contains authentication tokens in plain text. Do not commit it to source control. Add it to `.gitignore`. In CI, treat it as a sensitive cache artifact and restrict cache access to the appropriate pipeline/project.

---

## Verification Checklist

- [ ] `go test ./internal/config/...` — no regressions, new validation tests pass
- [ ] `go test ./internal/security/...` — resolver precedence, hash determinism
- [ ] `go test ./internal/state/...` — expiry buffer, concurrency, corruption recovery
- [ ] `go test ./internal/httpauth/...` — auth injection, proxy, retry_on_401
- [ ] `go test ./...` — full suite green
- [ ] Existing config with only `default_headers` runs without change
- [ ] `security: none` on endpoint disables project and API auth
- [ ] Changing `env.AUTH_URL` produces a different lock key
- [ ] `retry_on_401: false` by default; when `true`, one re-auth + one retry
- [ ] Two parallel test runners do not corrupt `testflowkit.lock`
