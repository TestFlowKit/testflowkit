---
title: Migration Guide
description: Guide for migrating from older TestFlowKit versions
navigation:
  title: Migration Guide
---

Breaking changes from older TestFlowKit versions. Run `tkit validate` after migrating.

## 1. Environment variables (replaces `environments`)

**Before:**

```yaml
active_environment: "local"
environments:
  local:
    frontend_base_url: "http://localhost:3000"
  staging:
    frontend_base_url: "https://staging.example.com"
```

```bash
tkit run --env staging
```

**After:**

```yaml
settings:
  env_file: ".env.local.yml"

env:
  base_url: "http://localhost:3000"

frontend:
  base_url: "{{ env.base_url }}"
```

```yaml
# .env.staging.yml
base_url: "https://staging.example.com"
```

```bash
tkit run --env-file .env.staging.yml
```

In Gherkin and config, replace environment-specific keys with `{{ env.variable_name }}`.

## 2. Unified APIs (replaces `backend` + `graphql` blocks)

**Before:**

```yaml
backend:
  baseUrl: "https://api.example.com"
  endpoints:
    get_users:
      method: GET
      path: "/users"

graphql:
  url: "https://api.example.com/graphql"
  queries:
    get_user: "queries/get_user.graphql"
```

**After:**

```yaml
apis:
  definitions:
    my_api:
      type: rest
      base_url: "https://api.example.com"
      endpoints:
        get_users:
          method: GET
          path: "/users"
    my_graphql:
      type: graphql
      endpoint: "https://api.example.com/graphql"
      operations:
        get_user:
          type: query
          operation: "queries/get_user.graphql"
```

**Step syntax change:**

| Old | New |
|-----|-----|
| `I prepare a REST request to "get_users"` | `I prepare a request to "my_api.get_users"` |
| `I prepare a GraphQL request for the "get_user" query` | `I prepare a request to "my_graphql.get_user"` |
| `I send the GraphQL request` | `I send the request` |

**Field renames:**

| Old | New |
|-----|-----|
| `baseUrl` | `base_url` |
| `url` (GraphQL) | `endpoint` |
| `headers` | `default_headers` |
| `queries` | `operations` |

## 3. `json` report is now Cucumber JSON

`report_format: "json"` still writes `report/report.json`, but the content changed. It used to be a single object listing scenarios with a start date and a duration; it is now the standard Cucumber JSON layout (features, then scenarios, then steps, with tags and durations in nanoseconds). Failure screenshots are now `image/png` embeddings on the failed step.

Tools that read the old file must be updated to read Cucumber JSON, which most reporting tools already understand.

A new `report_format: "junit"` writes `report/report.xml`. See [Reporters](/docs/reference/reporters).

## Checklist

- [ ] Replace `environments` with `env:` + env files
- [ ] Replace `--env` with `--env-file`
- [ ] Move `frontend.base_url` into `frontend` section with `{{ env.base_url }}`
- [ ] Merge `backend` and `graphql` into `apis.definitions`
- [ ] Update all feature files to `"api_name.endpoint_name"` syntax
- [ ] If a tool reads `report/report.json`, update it to Cucumber JSON
- [ ] Run `tkit validate`

## Next Steps

- [testflowkit.yml](/docs/config/overview) — Current config reference
- [API Testing](/docs/guides/api-testing) — Updated API steps
