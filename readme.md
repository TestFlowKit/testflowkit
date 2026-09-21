# TestFlowKit

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)](https://github.com/TestFlowKit/testflowkit/actions)
[![GitHub release](https://img.shields.io/github/v/release/TestFlowKit/testflowkit)](https://github.com/TestFlowKit/testflowkit/releases)

**TestFlowKit** is an open-source test automation framework written in Go. Write web UI, REST and GraphQL tests in Gherkin (plain-English steps), run them with one command, get HTML, Cucumber JSON or JUnit XML reports.

**Full documentation: [https://testflowkit.github.io/testflowkit](https://testflowkit.github.io/testflowkit/)**

## Get started in 3 minutes

```bash
npm install -g @testflowkit/cli   # or download a binary from GitHub Releases
tkit --version

mkdir my-tests && cd my-tests
tkit init       # scaffolds testflowkit.yml and a sample feature
tkit validate   # checks config and steps
tkit run        # runs the tests and writes the report
```

Then follow the [Quick Start](https://testflowkit.github.io/testflowkit/docs/getting-started/quick-start) to point the sample at your own app.

## What do you want to do?

| I want to... | Go to |
|---|---|
| Write my first UI test | [Quick Start](https://testflowkit.github.io/testflowkit/docs/getting-started/quick-start) |
| Find the step for an action or assertion | [Step Catalog](https://testflowkit.github.io/testflowkit/sentences) |
| Test a REST or GraphQL API | [Add a REST endpoint](https://testflowkit.github.io/testflowkit/docs/how-to/add-a-rest-endpoint), [Add a GraphQL operation](https://testflowkit.github.io/testflowkit/docs/how-to/add-a-graphql-operation) |
| Log in once for all tests | [Authenticate before tests](https://testflowkit.github.io/testflowkit/docs/how-to/authenticate-before-tests) |
| Run on staging or CI | [Test multiple environments](https://testflowkit.github.io/testflowkit/docs/how-to/test-multiple-environments), [CLI Reference](https://testflowkit.github.io/testflowkit/docs/reference/cli) |
| Reuse steps and data | [Macros](https://testflowkit.github.io/testflowkit/docs/patterns/macros), [Variables](https://testflowkit.github.io/testflowkit/docs/patterns/variables) |
| Let an AI agent write tests | [IDE Agent (MCP)](https://testflowkit.github.io/testflowkit/docs/guides/ide-agent) |
| Fix an error | [Troubleshooting](https://testflowkit.github.io/testflowkit/docs/troubleshooting/common-issues) |

## Features

- **Gherkin syntax** for readable BDD scenarios
- **Frontend testing** with Rod (default, bundled) or Playwright, CSS and XPath selectors, screenshot on failure
- **REST and GraphQL testing** with JSON (GJSON) and XML (XPath) assertions
- **Variables, random data, macros, global hooks** to keep scenarios short
- **Parallel execution** (`settings.concurrency`) and tag filtering (`--tags`)
- **Reports**: `report.html`, `report/report.json` (Cucumber JSON), `report/report.xml` (JUnit XML)

## Requirements

- Chrome or Edge for the default Rod driver
- Playwright driver only: run `tkit install` once (needs Go 1.19+)
- Building from source: Go 1.23+, Git, Make

## Install from source

```bash
git clone https://github.com/TestFlowKit/testflowkit.git
cd testflowkit
go mod tidy
make build GOOS=linux GOARCH=amd64  # or your target platform
```

## 🛠️ Development

### Building from Source

```bash
# Install dependencies
go mod tidy

# Run tests
make test

# Build for all platforms
make releases

# Build for specific platform
make build GOOS=linux GOARCH=amd64
```

### Running Tests

```bash
# Run all tests
make test

# Run with coverage
go test -v -race -coverprofile=coverage.out ./...

# Run end-to-end tests
make run_e2e
```

### Code Quality

```bash
# Run linter
make lint

# Format code
go fmt ./...

# Generate documentation
make generate_doc

# Export step definitions catalog (for agents / IDE tooling)
go run ./cmd/testflowkit/*.go export-step-definitions --format json > build/step-definitions.json
```

### Step definitions catalog

Generate the catalog locally from the CLI command and redirect stdout to a file:

```bash
mkdir -p build
go run ./cmd/testflowkit/*.go export-step-definitions --format json > build/step-definitions.json
```

The command emits JSON to stdout and returns a non-zero exit code on failure.


## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feat/amazing-feature`
3. **Follow the coding standards**:
   - Use conventional commit messages
   - Add tests for new features
   - Update documentation as needed
4. **Commit your changes**: `git commit -m 'feat: add amazing feature'`
5. **Push to the branch**: `git push origin feat/amazing-feature`
6. **Open a Pull Request**

### Branch Naming Convention

- `feat/`: New features
- `fix/`: Bug fixes
- `docs/`: Documentation updates
- `style/`: Code style changes
- `refactor/`: Code refactoring
- `test/`: Test additions or updates
- `chore/`: Maintenance tasks

### Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
type(scope): description

[optional body]

[optional footer]
```


## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Marc-Henry Nanguy** - _Initial work_ - [marckent04](https://github.com/marckent04)

## 🙏 Acknowledgments

### Contributors

- **Stéphane Salou** - [stephsalou](https://github.com/stephsalou)

### Dependencies

- **[alexflint/go-arg](https://github.com/alexflint/go-arg)** - Command-line argument parsing
- **[cucumber/godog](https://github.com/cucumber/godog)** - Gherkin parsing and BDD test execution framework
- **[fatih/color](https://github.com/fatih/color)** - Colorized terminal output
- **[go-rod/rod](https://github.com/go-rod/rod)** - Chrome DevTools Protocol automation library
- **[goccy/go-yaml](https://github.com/goccy/go-yaml)** - High-performance YAML parser and emitter
- **[gofrs/uuid/v5](https://github.com/gofrs/uuid/v5)** - UUID v5 implementation
- **[stretchr/testify](https://github.com/stretchr/testify)** - Testing utilities and assertions
- **[tdewolff/parse](https://github.com/tdewolff/parse)** - HTML/CSS parsing utilities
- **[tidwall/gjson](https://github.com/tidwall/gjson)** - Fast JSON parser and getter

## 📊 Project Status

- **Status**: Active Development
- **Go Version**: 1.23+
- **License**: MIT

## 🔗 Links

- **Documentation**: [https://testflowkit.github.io/testflowkit/](https://testflowkit.github.io/testflowkit/)
- **npm Package**: [https://www.npmjs.com/package/@testflowkit/cli](https://www.npmjs.com/package/@testflowkit/cli)
- **GitHub**: [https://github.com/TestFlowKit/testflowkit](https://github.com/TestFlowKit/testflowkit)
- **Issues**: [https://github.com/TestFlowKit/testflowkit/issues](https://github.com/TestFlowKit/testflowkit/issues)
- **Releases**: [https://github.com/TestFlowKit/testflowkit/releases](https://github.com/TestFlowKit/testflowkit/releases)

---

