---
title: Configure Authentication
description: Apply Bearer, Basic, API key or OAuth2 authentication to your API calls automatically
navigation:
  title: Configure Authentication
---

# Configure Authentication

## Goal

Have TestFlowKit attach authentication to every request for an API, instead of setting an `Authorization` header manually in each scenario.

Use this page when credentials are static (a fixed token, client id/secret). If you need to log in first and reuse a token obtained at runtime, see [Authenticate Before Tests](/docs/how-to/authenticate-before-tests).

## Configuration

Define a scheme once in `security_schemes`, then reference it from an API with `security_ref`:

```yaml
security_schemes:
  my_bearer:
    type: bearer
    token: "{{ env.api_token }}"

  my_oauth2:
    type: oauth2
    token_url: "{{ env.auth_url }}"
    client_id: "{{ env.client_id }}"
    client_secret: "{{ env.client_secret }}"
    token_endpoint_auth_method: client_secret_post
    scopes: ["read"]

default_security: "my_bearer"

apis:
  definitions:
    my_api:
      type: rest
      base_url: "{{ env.api_base_url }}"
      security_ref:
        name: my_oauth2
      endpoints:
        get_data:
          method: GET
          path: "/data"
```

Supported types: `bearer`, `basic`, `apikey`, `oauth2`. Use `security_ref.name: none` on a specific API to disable the default scheme.

## Scenario

No extra step needed — the header is added automatically:

```gherkin
Scenario: Read protected data
  Given I prepare a request to "my_api.get_data"
  When I send the request
  Then the response status code should be 200
```

## Verify

Use `I display the request cURL` in the scenario to confirm the `Authorization` header is present.

## Common issues

- **`oidc` and `certificate` are not yet implemented.**
- A request-level header set with `I set the header "Authorization" to "..."` overrides `security_ref`.

## See also

- [testflowkit.yml](/docs/config/overview) — full authentication section
- [Authenticate Before Tests](/docs/how-to/authenticate-before-tests) — login flow + token reuse
