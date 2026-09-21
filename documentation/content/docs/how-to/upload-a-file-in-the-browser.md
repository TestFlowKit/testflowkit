---
title: Upload a File in the Browser
description: Attach one or several files to a file input
navigation:
  title: Upload a File (Browser)
---

## Goal

Attach files to an `<input type="file">` in a UI scenario. For an API upload, see [Send a JSON body or file](/docs/how-to/send-a-json-body-or-file).

## 1. Give your files logical names

```yaml
files:
  base_directory: "./test-files"
  definitions:
    avatar: "images/avatar.png"
    photo_1: "images/photo1.jpg"
    photo_2: "images/photo2.jpg"
```

Paths are relative to `base_directory`.

## 2. Declare the file input

```yaml
frontend:
  elements:
    profile:
      avatar_input: "input[type=file]"
```

## 3. Upload

```gherkin
When the user uploads the "avatar" file into the "avatar_input" field
```

Several files into one input, names separated by commas:

```gherkin
When the user uploads the "photo_1, photo_2" files into the "gallery_input" field
```

## Verify

Assert on what your app shows after the upload, for example a preview or a file name:

```gherkin
Then the "avatar_preview" should be visible
```

If the step fails because a file is not found, check the name against `files.definitions` and the path against `base_directory`.

## See also

- [testflowkit.yml](/docs/config/overview): the `files` block
- [Fill in a form](/docs/how-to/fill-in-a-form)
