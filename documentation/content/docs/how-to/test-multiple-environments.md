---
title: Test Multiple Environments
description: Run the same tests against local, staging, or production with different config
navigation:
  title: Test Multiple Environments
---

# Test Multiple Environments

## Goal

Run the same `.feature` files against different environments (local, staging, production) without changing any test.

## Configuration

Keep environment-specific values (URLs, secrets) out of the main config, in a separate env file:

```yaml
# .env.staging.yml
base_url: "https://staging.example.com"
api_base_url: "https://api-staging.example.com"
```

Reference `{{ env.* }}` anywhere in `testflowkit.yml`:

```yaml
env:
  base_url: "http://localhost:3000"
  api_base_url: "http://localhost:3001"

frontend:
  base_url: "{{ env.base_url }}"

apis:
  definitions:
    my_api:
      base_url: "{{ env.api_base_url }}"
```

## Run against an environment

```bash
tkit run --env-file .env.staging.yml
```

Priority order: CLI `--env-file` → `settings.env_file` in `testflowkit.yml` → inline `env:` block.

You can also point to an entirely different config file per environment:

```bash
tkit run --config staging.yml
```

## Verify

Run the same scenario against two env files and confirm requests hit the right base URLs (use `I display the request cURL` to check).

## Common issues

- Secrets committed in `testflowkit.yml` — keep them in an untracked `.env.*.yml` file instead.
- Nested env values use dot notation: `{{ env.database.host }}`.

## See also

- [testflowkit.yml](/docs/config/overview) — environment files reference
- [CLI Reference](/docs/reference/cli) — `--env-file` / `--config` flags
