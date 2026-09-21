---
title: Send a JSON Body or File
description: Send an inline JSON body, a body from a file, or upload a file with a request
navigation:
  title: Send a JSON Body or File
---

## Goal

Send a request body as inline JSON, from a file, or upload a file as part of the request.

## Inline JSON body

```gherkin
Given I prepare a request to "my_api.create_user"
And I set the JSON request body to:
  """
  {
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30,
    "tags": ["developer", "golang"]
  }
  """
When I send the request
Then the response status code should be 201
```

`I set the JSON request body to:` validates the body is well-formed JSON before sending.

## Body from a file

Declare the file in `testflowkit.yml`:

```yaml
files:
  base_directory: "./test-files"
  definitions:
    sample_json: "data/post.json"
```

Reference it in a scenario:

```gherkin
Given I prepare a request to "my_api.create_user"
And I set the request body from file "sample_json"
When I send the request
```

## Verify

Assert the response reflects the sent body:

```gherkin
Then the response field "name" should be "John Doe"
```

## Common issues

- **Invalid JSON** — `I set the JSON request body to:` fails fast with a validation error; use `I set the request body to:` for raw/non-JSON payloads.
- **File not found** — check `files.base_directory` and the relative path in `definitions`.

## See also

- [API Testing](/docs/guides/api-testing) — full REST/GraphQL guide
- [testflowkit.yml](/docs/config/overview) — `files` section reference
