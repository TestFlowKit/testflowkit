# TestFlowKit VS Code Extension

Autocomplete, inline documentation, and macro data-table generation for Gherkin `.feature` files in TestFlowKit projects.

## Features

### 1. Step sentence autocomplete
Type a Gherkin step keyword (`Given`, `When`, `Then`, `And`, `But`) followed by a space and part of a sentence — suggestions appear immediately.

Each suggestion shows:
- **Sentence text** as the label
- **Category** (restapi, form, graphql, …) in the detail field
- **Description**, variable table, and Gherkin example in the documentation panel

### 2. Tab-navigable wildcard placeholders
When you accept a sentence completion, every `{string}` / `{int}` / `{float}` wildcard is converted into a tab-stop using the variable's name. Press **Tab** to jump from one placeholder to the next.

```gherkin
# Accepted completion inserts this snippet:
When the user enters ${1:text} into the ${2:name} field
# Tab stops let you fill "text" then "name" without touching the mouse
```

### 3. Macro completion with auto-generated data table
Macros discovered across all `@macro`-tagged scenarios in your workspace are offered as completion items. Accepting a macro that has variables automatically inserts the call line **plus a two-column data table** with one row per variable:

```gherkin
When user login with credentials
  | username | ${1:username_value} |
  | password | ${2:password_value} |
```

- Macro completions are also tab-navigable (value column = tab-stop)
- Macros with no variables insert only the call line

### 4. Hover documentation
Hover over any step in a `.feature` file to see the sentence's description, variables, and example — or the macro's variable list if it is a macro call.

### 5. Refresh commands
| Command | What it does |
|---------|-------------|
| **TestFlowKit: Refresh Sentence Index** | Reloads all sentence JSON files |
| **TestFlowKit: Refresh Macro Index** | Re-scans workspace `.feature` files for `@macro` scenarios |

## Requirements

- The sentence catalog must be generated first: run `make generate_doc` (or the equivalent for your setup) in the TestFlowKit repository root. Sentences are read from `documentation/content/sentences/**/*.json`.
- No external Gherkin extension is required, but if you have one installed the extension will also activate under its language id.

## Configuration

| Setting | Default | Description |
|---------|---------|-------------|
| `testflowkit.sentencesPath` | `""` (auto-detect) | Custom path to the sentences directory, relative to the workspace root |
| `testflowkit.enableMacroCompletion` | `true` | Insert a data table when completing a macro call |

## Development

```bash
# Install dependencies
npm install

# Build (development)
npm run compile:dev

# Build (production)
npm run compile

# Run tests
npm test

# Watch for changes
npm run watch
```

To launch the extension in the VS Code Extension Development Host, press **F5** after opening this folder.
