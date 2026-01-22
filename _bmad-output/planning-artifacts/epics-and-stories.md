# TestFlowKit V1.0 - Epics & User Stories

**Version:** 1.0  
**Date:** 2026-01-22  
**Product Manager:** John  
**Scrum Master:** TBD  
**Type de Release:** V1.0 Enterprise Edition

---

## 📋 Vue d'Ensemble

### Structure Release V1.0

**Total Duration:** 7 semaines  
**Total Epics:** 12 epics  
**Total Stories:** 40 user stories  
**Total Story Points:** ~140 SP  
**Team Velocity:** ~35 SP/sprint (2 développeurs)

### Sprint Breakdown

| Sprint | Duration | Focus | Stories | Story Points |
|--------|----------|-------|---------|--------------|
| Sprint 1 | Semaines 1-2 | Must-Have Core (Reporting) | 8 stories | 34 SP |
| Sprint 2 | Semaines 3-4 | High-Value (Debugging) | 10 stories | 38 SP |
| Sprint 3 | Semaines 5-6 | Differentiators (Visual) | 12 stories | 42 SP |
| Sprint 4 | Semaine 7 | Polish & Launch | 10 stories | 26 SP |

---

## 🎯 SPRINT 1: Must-Have Core (Semaines 1-2)

**Objectif Sprint:** Établir les fondations enterprise avec reporting professionnel et génération de données

**Story Points Total:** 34 SP  
**Features:** 5 features principales  
**Deliverable:** Reporting Allure + Cucumber JSON + Test Data Faker + Performance Metrics

---

### EPIC 1: Professional Reporting System

**Epic Goal:** Implémenter système de reporting enterprise-grade (Allure, Cucumber JSON)  
**Business Value:** Intégration CI/CD et visualisation professionnelle  
**Story Points:** 13 SP  
**Priority:** P0 - Critical

#### Story 1.1: Allure JSON Format Support
**As a** QA Engineer  
**I want** to generate test results in Allure JSON format  
**So that** I can use Allure reporting framework for professional visualizations

**Story Points:** 5 SP  
**Priority:** P0

**Acceptance Criteria:**
- [ ] AC1: Reporter interface extended with AllureReporter implementation
- [ ] AC2: Integration of `github.com/allure-framework/allure-go` library
- [ ] AC3: CLI flag `--report-format=allure` generates Allure JSON
- [ ] AC4: All test metadata mapped to Allure format (status, timing, steps)
- [ ] AC5: Screenshots attached to Allure reports
- [ ] AC6: Tags/features categorization in Allure
- [ ] AC7: Unit tests for AllureReporter (80%+ coverage)

**Technical Tasks:**
- [ ] Install allure-go dependency
- [ ] Create `pkg/reporters/allure/allure_reporter.go`
- [ ] Implement Reporter interface
- [ ] Map TestResult → AllureResult
- [ ] Attach screenshots to results
- [ ] Add CLI flag parsing
- [ ] Write unit tests
- [ ] Documentation update

**Dependencies:** None  
**Risks:** None  
**Definition of Done:**
- Code merged to main
- Tests passing (80%+ coverage)
- Documentation updated
- Peer review approved

---

#### Story 1.2: Allure Historical Trends
**As a** Test Manager  
**I want** to see historical test trends in Allure  
**So that** I can track quality improvements over time

**Story Points:** 3 SP  
**Priority:** P0

**Acceptance Criteria:**
- [ ] AC1: Test execution history stored in `history/` directory
- [ ] AC2: Allure history.json generated per run
- [ ] AC3: Trends visible in Allure report (pass/fail rates)
- [ ] AC4: Retries tracked in history
- [ ] AC5: Duration trends visible

**Technical Tasks:**
- [ ] Implement history storage logic
- [ ] Generate history.json file
- [ ] Trend calculation algorithm
- [ ] Documentation

**Dependencies:** Story 1.1  
**DoD:** Same as 1.1

---

#### Story 1.3: Cucumber JSON Export
**As a** DevOps Engineer  
**I want** to export test results in Cucumber JSON format  
**So that** I can integrate with Jenkins/GitLab CI/CD plugins

**Story Points:** 3 SP  
**Priority:** P0

**Acceptance Criteria:**
- [ ] AC1: CLI flag `--report-format=cucumber-json`
- [ ] AC2: Output file `cucumber-report.json` in standard format
- [ ] AC3: Compatible with Jenkins Cucumber Reports Plugin
- [ ] AC4: Compatible with GitLab Test Reports
- [ ] AC5: All features, scenarios, steps mapped correctly
- [ ] AC6: Tags and metadata included
- [ ] AC7: Timestamps in ISO 8601 format

**Technical Tasks:**
- [ ] Create `pkg/reporters/cucumber/cucumber_json.go`
- [ ] Implement Cucumber JSON schema
- [ ] Map TestResult → Cucumber format
- [ ] CLI flag integration
- [ ] Validation against schema
- [ ] Integration tests with Jenkins/GitLab
- [ ] Documentation

**Dependencies:** JSON reporter existant  
**DoD:** Same as 1.1 + validation with CI tools

---

#### Story 1.4: Multi-Format Report Generation
**As a** Developer  
**I want** to generate multiple report formats simultaneously  
**So that** I can use different tools for different purposes

**Story Points:** 2 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: CLI flag `--report-format=html,allure,cucumber-json`
- [ ] AC2: Multiple reporters run in parallel
- [ ] AC3: All formats generated in < 3s
- [ ] AC4: Error handling per reporter
- [ ] AC5: Config file support for default formats

**Technical Tasks:**
- [ ] Reporter factory pattern
- [ ] Parallel report generation
- [ ] Config integration
- [ ] Error handling
- [ ] Performance testing

**Dependencies:** Story 1.1, 1.3  
**DoD:** Same as 1.1

---

### EPIC 2: Test Data Generation

**Epic Goal:** Implémenter système de génération données de test réalistes  
**Business Value:** Réduire hardcoding, améliorer robustesse tests  
**Story Points:** 8 SP  
**Priority:** P0

#### Story 2.1: Faker Library Integration
**As a** QA Engineer  
**I want** to generate random test data (emails, names, phones)  
**So that** I can avoid hardcoding values and improve test coverage

**Story Points:** 5 SP  
**Priority:** P0

**Acceptance Criteria:**
- [ ] AC1: Integration `github.com/brianvoe/gofakeit`
- [ ] AC2: New steps in `internal/step_definitions/variables/faker.go`
- [ ] AC3: Step: `I set "field" to random email`
- [ ] AC4: Support types: email, name, phone, address, UUID, date, number
- [ ] AC5: Values stored in scenario context variables
- [ ] AC6: Reproducible with seed: `--faker-seed=12345`
- [ ] AC7: Documentation with 20+ examples

**Technical Tasks:**
- [ ] Install gofakeit dependency
- [ ] Create faker.go step definitions
- [ ] Implement 20+ faker types
- [ ] Variable storage integration
- [ ] Seed support CLI flag
- [ ] Unit tests (90%+ coverage)
- [ ] Documentation + examples

**Dependencies:** Variable system existant  
**DoD:** Same as 1.1 + 20+ types supported

---

#### Story 2.2: Custom Faker Formats
**As a** Developer  
**I want** to specify custom formats for generated data  
**So that** I can match specific validation rules

**Story Points:** 3 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Step: `I set "field" to random phone "US format"`
- [ ] AC2: Step: `I set "field" to random date "YYYY-MM-DD"`
- [ ] AC3: Support regex patterns
- [ ] AC4: Custom templates: `random "###-###-####"`
- [ ] AC5: Documentation formats disponibles

**Technical Tasks:**
- [ ] Format parsing logic
- [ ] Template engine integration
- [ ] Regex support
- [ ] Examples + tests
- [ ] Documentation

**Dependencies:** Story 2.1  
**DoD:** Same as 1.1

---

### EPIC 3: Performance Monitoring

**Epic Goal:** Capturer et reporter métriques de performance  
**Business Value:** Visibilité performance, détection régressions  
**Story Points:** 5 SP  
**Priority:** P0

#### Story 3.1: Basic Performance Metrics Capture
**As a** Performance Engineer  
**I want** to capture response times and page load times  
**So that** I can monitor performance regressions

**Story Points:** 3 SP  
**Priority:** P0

**Acceptance Criteria:**
- [ ] AC1: API response time logged per request
- [ ] AC2: Page load time captured per navigation
- [ ] AC3: Step execution time in reports
- [ ] AC4: Performance struct in Report model
- [ ] AC5: Metrics in HTML/JSON reports
- [ ] AC6: Overhead < 1% execution time

**Technical Tasks:**
- [ ] Add Performance struct to models
- [ ] Timer integration in HTTP client
- [ ] Page load time capture (Rod)
- [ ] Report rendering
- [ ] Performance testing
- [ ] Documentation

**Dependencies:** Reporters existants  
**DoD:** Same as 1.1 + performance overhead verified

---

#### Story 3.2: Performance Thresholds & Warnings
**As a** QA Engineer  
**I want** to be warned when operations exceed thresholds  
**So that** I can identify slow tests early

**Story Points:** 2 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Config: `performance.thresholds.api_response: 500ms`
- [ ] AC2: Warning in logs si > threshold
- [ ] AC3: Highlighted in reports
- [ ] AC4: Fail test option: `fail_on_slow: true`
- [ ] AC5: Aggregate metrics per scenario

**Technical Tasks:**
- [ ] Config schema update
- [ ] Threshold checking logic
- [ ] Warning/fail mechanisms
- [ ] Report highlights
- [ ] Tests + docs

**Dependencies:** Story 3.1  
**DoD:** Same as 1.1

---

### EPIC 4: Environment Configuration

**Epic Goal:** Améliorer flexibilité configuration environnements  
**Business Value:** Faciliter tests multi-environnements  
**Story Points:** 5 SP  
**Priority:** P0

#### Story 4.1: CLI Environment Variable Override
**As a** Developer  
**I want** to override env variables via CLI  
**So that** I can test different environments without changing config files

**Story Points:** 3 SP  
**Priority:** P0

**Acceptance Criteria:**
- [ ] AC1: Syntax `tkit run --env KEY=VALUE`
- [ ] AC2: Multiple vars: `--env KEY1=VAL1 --env KEY2=VAL2`
- [ ] AC3: Override order: CLI > .env file > config.yaml
- [ ] AC4: Variables accessible in Gherkin: `{{ env.KEY }}`
- [ ] AC5: Validation required vars
- [ ] AC6: Documentation + examples

**Technical Tasks:**
- [ ] Extend CLI args parser (alexflint/go-arg)
- [ ] Environment merge logic
- [ ] Priority resolution
- [ ] Variable substitution
- [ ] Validation
- [ ] Tests + docs

**Dependencies:** None  
**DoD:** Same as 1.1

---

#### Story 4.2: Environment Templates
**As a** QA Lead  
**I want** to define environment templates  
**So that** team members can quickly switch environments

**Story Points:** 2 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Config: `tkit run --env-template=staging`
- [ ] AC2: Templates in `config.yaml` or separate files
- [ ] AC3: Template inheritance
- [ ] AC4: List templates: `tkit list-env-templates`
- [ ] AC5: Documentation

**Technical Tasks:**
- [ ] Template system design
- [ ] Config loading
- [ ] List command
- [ ] Inheritance logic
- [ ] Tests + docs

**Dependencies:** Story 4.1  
**DoD:** Same as 1.1

---

### EPIC 5: CSV/Excel Data-Driven Testing

**Epic Goal:** Support data-driven testing avec sources externes  
**Business Value:** Éviter duplication scénarios, tester multiples cas  
**Story Points:** 8 SP  
**Priority:** P1

#### Story 5.1: CSV Data Source for Scenario Outline
**As a** QA Engineer  
**I want** to use CSV files as data source for Scenario Outline  
**So that** I can test multiple cases without duplicating Gherkin

**Story Points:** 5 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Syntax: `Examples: @data(users.csv)`
- [ ] AC2: CSV with headers mapped to variables
- [ ] AC3: Support large files (10,000+ rows)
- [ ] AC4: Skip invalid rows with warnings
- [ ] AC5: Report shows data source
- [ ] AC6: Performance: load CSV < 1s
- [ ] AC7: Documentation + examples

**Technical Tasks:**
- [ ] CSV parser (stdlib encoding/csv)
- [ ] Examples tag parser extension
- [ ] Godog integration
- [ ] Row validation
- [ ] Performance optimization
- [ ] Tests (including 10K rows)
- [ ] Documentation

**Dependencies:** Macro system, Godog  
**DoD:** Same as 1.1 + performance verified

---

#### Story 5.2: Excel Data Source Support
**As a** Business Analyst  
**I want** to use Excel files as data source  
**So that** I can manage test data in familiar tools

**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Syntax: `Examples: @data(testdata.xlsx)`
- [ ] AC2: Support .xlsx format
- [ ] AC3: Multiple sheets support
- [ ] AC4: Sheet selection: `@data(file.xlsx:Sheet1)`
- [ ] AC5: Performance acceptable (< 2s for 1000 rows)
- [ ] AC6: Documentation

**Technical Tasks:**
- [ ] Install `github.com/xuri/excelize`
- [ ] Excel parser implementation
- [ ] Sheet handling
- [ ] Integration with CSV logic
- [ ] Tests + docs

**Dependencies:** Story 5.1  
**DoD:** Same as 1.1

---

## 🎯 SPRINT 2: High-Value Features (Semaines 3-4)

**Objectif Sprint:** Améliorer capacités debugging et testing backend avancé

**Story Points Total:** 38 SP  
**Features:** 5 features principales  
**Deliverable:** Console logs + Network logging + WebSocket + Skip/Pending + Storage

---

### EPIC 6: Enhanced Debugging

**Epic Goal:** Capturer logs et requêtes pour faciliter debugging  
**Business Value:** Réduire temps debug, améliorer troubleshooting  
**Story Points:** 13 SP  
**Priority:** P1

#### Story 6.1: Browser Console Logs Capture
**As a** Developer  
**I want** to see browser console logs in test reports  
**So that** I can debug JavaScript errors

**Story Points:** 5 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Capture console.log, .warn, .error, .info
- [ ] AC2: Display in HTML report with filtering
- [ ] AC3: Console logs per step
- [ ] AC4: Timestamps and source file/line
- [ ] AC5: Config: enable/disable capture
- [ ] AC6: Performance overhead < 5%
- [ ] AC7: Documentation

**Technical Tasks:**
- [ ] Rod console event listeners
- [ ] ConsoleLog model struct
- [ ] Storage in Step context
- [ ] HTML template rendering
- [ ] Filter UI (log levels)
- [ ] Config integration
- [ ] Tests + docs

**Dependencies:** Rod browser, HTML reporter  
**DoD:** Same as 1.1 + overhead verified

---

#### Story 6.2: Network Request Logging
**As a** API Tester  
**I want** to see all HTTP requests/responses in reports  
**So that** I can debug API failures

**Story Points:** 5 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Log all HTTP requests (method, URL, headers, body)
- [ ] AC2: Log all responses (status, headers, body, timing)
- [ ] AC3: Display in HTML report
- [ ] AC4: Filter by status code
- [ ] AC5: Highlight failed requests (4xx, 5xx)
- [ ] AC6: Optional HAR export format
- [ ] AC7: Performance overhead < 2%

**Technical Tasks:**
- [ ] HTTP client wrapper/interceptor
- [ ] NetworkRequest model
- [ ] Storage in context
- [ ] HTML template update
- [ ] HAR export (optional)
- [ ] Tests + docs

**Dependencies:** HTTP client existant  
**DoD:** Same as 1.1

---

#### Story 6.3: Enhanced Error Messages
**As a** QA Engineer  
**I want** detailed error messages with context  
**So that** I can quickly understand failures

**Story Points:** 3 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Error messages include step context
- [ ] AC2: Suggest fixes for common errors
- [ ] AC3: Show element selectors tried
- [ ] AC4: Include relevant logs/screenshots
- [ ] AC5: Color-coded error types
- [ ] AC6: Documentation common errors

**Technical Tasks:**
- [ ] Error message formatter
- [ ] Context extraction
- [ ] Suggestion engine (common errors)
- [ ] Template updates
- [ ] Tests + docs

**Dependencies:** None  
**DoD:** Same as 1.1

---

### EPIC 7: WebSocket Testing

**Epic Goal:** Support WebSocket connections pour real-time testing  
**Business Value:** Tester apps temps-réel (chat, notifications)  
**Story Points:** 8 SP  
**Priority:** P1

#### Story 7.1: WebSocket Connection & Messaging
**As a** QA Engineer  
**I want** to test WebSocket connections  
**So that** I can validate real-time features

**Story Points:** 5 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Step: `I connect to WebSocket "ws://localhost:8080/ws"`
- [ ] AC2: Step: `I send WebSocket message "Hello"`
- [ ] AC3: Step: `I should receive WebSocket message containing "World"`
- [ ] AC4: Support ws:// and wss:// protocols
- [ ] AC5: Timeout configurable
- [ ] AC6: Auto cleanup connections
- [ ] AC7: Documentation + examples

**Technical Tasks:**
- [ ] Install `github.com/gorilla/websocket`
- [ ] Create `internal/step_definitions/backend/websocket/`
- [ ] Connection management
- [ ] Message send/receive steps
- [ ] Context storage
- [ ] AfterScenario cleanup
- [ ] Tests + docs

**Dependencies:** Backend step pattern  
**DoD:** Same as 1.1

---

#### Story 7.2: WebSocket JSON Messages
**As a** API Tester  
**I want** to send/receive JSON over WebSocket  
**So that** I can test structured messages

**Story Points:** 3 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Step: `I send WebSocket JSON message:` (with DataTable)
- [ ] AC2: Step: `the WebSocket message field "status" should be "ok"`
- [ ] AC3: JSON validation
- [ ] AC4: Variable extraction from messages
- [ ] AC5: Pretty print in logs
- [ ] AC6: Documentation

**Technical Tasks:**
- [ ] JSON serialization/deserialization
- [ ] Field validation steps
- [ ] Variable extraction
- [ ] Tests + docs

**Dependencies:** Story 7.1  
**DoD:** Same as 1.1

---

### EPIC 8: Test Control & Skip/Pending

**Epic Goal:** Support tags @skip, @pending, @wip  
**Business Value:** Contrôle exécution flexible  
**Story Points:** 5 SP  
**Priority:** P1

#### Story 8.1: Skip & Pending Tag Support
**As a** Developer  
**I want** to mark tests as @skip or @pending  
**So that** I can control which tests run

**Story Points:** 3 SP  
**Priority:** P1

**Acceptance Criteria:**
- [ ] AC1: Tag @skip → test skipped
- [ ] AC2: Tag @pending → test runs but failure OK
- [ ] AC3: Tag @wip → work in progress
- [ ] AC4: CLI: `--skip-pending` flag
- [ ] AC5: Reports show skip/pending status separately
- [ ] AC6: Metrics: passed, failed, skipped, pending
- [ ] AC7: Documentation

**Technical Tasks:**
- [ ] Extend tag filtering (Godog)
- [ ] Status enum update (Skipped, Pending)
- [ ] CLI flag integration
- [ ] Report rendering
- [ ] Tests + docs

**Dependencies:** Tag system existant  
**DoD:** Same as 1.1

---

#### Story 8.2: Conditional Execution Tags
**As a** QA Lead  
**I want** to run tests based on conditions  
**So that** I can optimize test execution

**Story Points:** 2 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Tag @requires(auth) → skip if no auth
- [ ] AC2: Tag @env(staging,prod) → run only on envs
- [ ] AC3: Condition evaluation
- [ ] AC4: Documentation

**Technical Tasks:**
- [ ] Condition parser
- [ ] Evaluation logic
- [ ] Integration
- [ ] Tests + docs

**Dependencies:** Story 8.1  
**DoD:** Same as 1.1

---

### EPIC 9: Browser Storage Management

**Epic Goal:** Support cookies, localStorage, sessionStorage  
**Business Value:** Tester client-side state management  
**Story Points:** 6 SP  
**Priority:** P2

#### Story 9.1: Cookie Management
**As a** QA Engineer  
**I want** to manage cookies (set, get, delete)  
**So that** I can test session and authentication

**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `I set cookie "name" to "value"`
- [ ] AC2: Step: `the cookie "name" should be "value"`
- [ ] AC3: Step: `I delete cookie "name"`
- [ ] AC4: Step: `I clear all cookies`
- [ ] AC5: Support domain, path, expiry attributes
- [ ] AC6: Cookie storage in variables
- [ ] AC7: Documentation

**Technical Tasks:**
- [ ] Create `frontend/cookies/` steps
- [ ] Rod cookie API integration
- [ ] Attribute support
- [ ] Variable storage
- [ ] Tests + docs

**Dependencies:** Browser interface  
**DoD:** Same as 1.1

---

#### Story 9.2: Local/Session Storage
**As a** Frontend Developer  
**I want** to manipulate localStorage and sessionStorage  
**So that** I can test client-side state

**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `I set localStorage "key" to "value"`
- [ ] AC2: Step: `localStorage "key" should be "value"`
- [ ] AC3: Step: `I clear localStorage`
- [ ] AC4: Support sessionStorage
- [ ] AC5: Support JSON values
- [ ] AC6: Cross-step variable storage
- [ ] AC7: Documentation

**Technical Tasks:**
- [ ] JavaScript execution helpers
- [ ] Storage manipulation steps
- [ ] JSON serialization
- [ ] Tests + docs

**Dependencies:** Browser JS execution  
**DoD:** Same as 1.1

---

### EPIC 10: File Operations

**Epic Goal:** Support file download validation et PDF testing  
**Business Value:** Tester téléchargements et documents  
**Story Points:** 6 SP  
**Priority:** P2

#### Story 10.1: File Download Validation
**As a** QA Engineer  
**I want** to validate downloaded files  
**So that** I can test export features

**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `I download file from "button"`
- [ ] AC2: Step: `the downloaded file name should be "report.pdf"`
- [ ] AC3: Step: `the downloaded file size should be > 1MB`
- [ ] AC4: Step: `the downloaded file should contain "text"`
- [ ] AC5: Auto cleanup temp files
- [ ] AC6: Documentation

**Technical Tasks:**
- [ ] Rod download interceptor
- [ ] File validation helpers
- [ ] Temp directory management
- [ ] AfterScenario cleanup
- [ ] Tests + docs

**Dependencies:** File system utils  
**DoD:** Same as 1.1

---

#### Story 10.2: PDF Content Validation
**As a** Tester  
**I want** to extract and validate PDF content  
**So that** I can test report generation

**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `the PDF "file" should contain "text"`
- [ ] AC2: Extract page count
- [ ] AC3: Extract metadata
- [ ] AC4: Support encrypted PDFs
- [ ] AC5: Documentation

**Technical Tasks:**
- [ ] Install `github.com/ledongthuc/pdf`
- [ ] PDF parser
- [ ] Text extraction
- [ ] Tests + docs

**Dependencies:** Story 10.1  
**DoD:** Same as 1.1

---

## 🎯 SPRINT 3: Differentiators (Semaines 5-6)

**Objectif Sprint:** Implémenter features de différenciation marché

**Story Points Total:** 42 SP  
**Features:** Visual regression, GraphQL advanced, Accessibility  
**Deliverable:** Competitive advantages

---

### EPIC 11: Visual Regression Testing

**Epic Goal:** Comparer screenshots pour détecter régressions visuelles  
**Business Value:** Différenciateur marché, qualité UI  
**Story Points:** 13 SP  
**Priority:** P3

#### Story 11.1: Screenshot Baseline Management
**As a** QA Engineer  
**I want** to save screenshot baselines  
**So that** I can compare against future runs

**Story Points:** 5 SP  
**Priority:** P3

**Acceptance Criteria:**
- [ ] AC1: Step: `I save screenshot baseline "homepage"`
- [ ] AC2: Baselines stored in `baselines/` directory
- [ ] AC3: Organized by test name
- [ ] AC4: Versioning support
- [ ] AC5: Update baseline step
- [ ] AC6: Baseline management CLI
- [ ] AC7: Documentation

**Technical Tasks:**
- [ ] Baseline storage structure
- [ ] Save/load logic
- [ ] Versioning system
- [ ] CLI commands
- [ ] Tests + docs

**Dependencies:** Screenshot system existant  
**DoD:** Same as 1.1

---

#### Story 11.2: Visual Comparison Algorithm
**As a** Visual Tester  
**I want** to compare screenshots with tolerance  
**So that** I can detect visual regressions

**Story Points:** 5 SP  
**Priority:** P3

**Acceptance Criteria:**
- [ ] AC1: Step: `the page should match baseline "name"`
- [ ] AC2: Pixel-by-pixel comparison
- [ ] AC3: Configurable tolerance (0-100%)
- [ ] AC4: Generate diff image
- [ ] AC5: Highlight differences
- [ ] AC6: Report similarity percentage
- [ ] AC7: Performance: < 500ms per comparison

**Technical Tasks:**
- [ ] Image diff algorithm (`golang.org/x/image`)
- [ ] Tolerance calculation
- [ ] Diff image generation
- [ ] Performance optimization
- [ ] Tests + docs

**Dependencies:** Story 11.1  
**DoD:** Same as 1.1 + performance verified

---

#### Story 11.3: Ignore Regions for Dynamic Content
**As a** Frontend Tester  
**I want** to ignore dynamic regions in comparisons  
**So that** timestamps/ads don't cause false failures

**Story Points:** 3 SP  
**Priority:** P3

**Acceptance Criteria:**
- [ ] AC1: Define ignore regions by selector
- [ ] AC2: Define ignore regions by coordinates
- [ ] AC3: Mask regions in comparison
- [ ] AC4: Visual indication in diff
- [ ] AC5: Documentation

**Technical Tasks:**
- [ ] Region masking algorithm
- [ ] Selector → coordinates conversion
- [ ] Integration with comparison
- [ ] Tests + docs

**Dependencies:** Story 11.2  
**DoD:** Same as 1.1

---

### EPIC 12: GraphQL Advanced Testing

**Epic Goal:** Schema validation et introspection GraphQL  
**Business Value:** Différenciateur GraphQL expertise  
**Story Points:** 8 SP  
**Priority:** P3

#### Story 12.1: GraphQL Schema Introspection
**As a** API Developer  
**I want** to introspect GraphQL schema  
**So that** I can validate schema structure

**Story Points:** 3 SP  
**Priority:** P3

**Acceptance Criteria:**
- [ ] AC1: Step: `the GraphQL schema should be valid`
- [ ] AC2: Step: `the schema should have type "User"`
- [ ] AC3: Step: `the schema should have field "User.email"`
- [ ] AC4: Auto introspection query
- [ ] AC5: Cache schema for performance
- [ ] AC6: Documentation

**Technical Tasks:**
- [ ] Introspection query implementation
- [ ] Schema parser
- [ ] Validation logic
- [ ] Caching
- [ ] Tests + docs

**Dependencies:** GraphQL client existant  
**DoD:** Same as 1.1

---

#### Story 12.2: GraphQL Schema Diff & Breaking Changes
**As a** GraphQL Maintainer  
**I want** to detect breaking schema changes  
**So that** I can prevent API regressions

**Story Points:** 5 SP  
**Priority:** P3

**Acceptance Criteria:**
- [ ] AC1: Save baseline schema
- [ ] AC2: Compare schemas (baseline vs current)
- [ ] AC3: Detect breaking changes (field removal, type change)
- [ ] AC4: Detect non-breaking changes (new fields)
- [ ] AC5: Report with severity levels
- [ ] AC6: Fail on breaking changes option
- [ ] AC7: Documentation

**Technical Tasks:**
- [ ] Schema diff algorithm
- [ ] Breaking change rules
- [ ] Severity classification
- [ ] Report generation
- [ ] Tests + docs

**Dependencies:** Story 12.1  
**DoD:** Same as 1.1

---

### EPIC 13: Accessibility Testing (Basic)

**Epic Goal:** Checks accessibilité WCAG basiques  
**Business Value:** Qualité inclusive  
**Story Points:** 8 SP  
**Priority:** P2-P3

#### Story 13.1: Basic WCAG AA Checks
**As a** Accessibility Tester  
**I want** to run basic accessibility checks  
**So that** I can ensure WCAG AA compliance

**Story Points:** 5 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `the page should be accessible`
- [ ] AC2: Check: images have alt text
- [ ] AC3: Check: form inputs have labels
- [ ] AC4: Check: color contrast ratios
- [ ] AC5: Check: ARIA attributes valid
- [ ] AC6: Detailed violation report
- [ ] AC7: Configurable severity (A, AA, AAA)

**Technical Tasks:**
- [ ] Inject axe-core JavaScript library
- [ ] Execute axe.run() via Rod
- [ ] Parse violations
- [ ] Report formatting
- [ ] Severity filtering
- [ ] Tests + docs

**Dependencies:** Browser JS execution  
**DoD:** Same as 1.1

---

#### Story 13.2: Accessibility Report Integration
**As a** QA Lead  
**I want** accessibility results in test reports  
**So that** I can track compliance

**Story Points:** 3 SP  
**Priority:** P3

**Acceptance Criteria:**
- [ ] AC1: Violations in HTML report
- [ ] AC2: Violations in Allure report
- [ ] AC3: Aggregate metrics
- [ ] AC4: Trends over time
- [ ] AC5: Documentation

**Technical Tasks:**
- [ ] Report template updates
- [ ] Allure integration
- [ ] Metrics aggregation
- [ ] Tests + docs

**Dependencies:** Story 13.1, Allure reporter  
**DoD:** Same as 1.1

---

### EPIC 14: Advanced Browser Features

**Epic Goal:** iFrame, Shadow DOM, Multi-tab, Network throttling  
**Business Value:** Support composants modernes  
**Story Points:** 13 SP  
**Priority:** P2

#### Story 14.1: iFrame Support
**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `I switch to iframe "name"`
- [ ] AC2: Step: `I switch to iframe by selector`
- [ ] AC3: Step: `I switch to parent frame`
- [ ] AC4: Nested iframes support (3+ levels)
- [ ] AC5: Auto-detect iframe context
- [ ] AC6: Documentation

**Technical Tasks:**
- [ ] Rod frame switching
- [ ] Context management
- [ ] Frame stack for nested
- [ ] Tests + docs

**Dependencies:** Browser context  
**DoD:** Same as 1.1

---

#### Story 14.2: Shadow DOM Support
**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Selector: `shadow::#element`
- [ ] AC2: Nested shadow roots support
- [ ] AC3: Auto-detect shadow DOM
- [ ] AC4: Fallback if no shadow
- [ ] AC5: Documentation

**Technical Tasks:**
- [ ] Shadow root traversal
- [ ] Selector strategy extension
- [ ] Tests + docs

**Dependencies:** Selector engine  
**DoD:** Same as 1.1

---

#### Story 14.3: Enhanced Multi-Tab Management
**Story Points:** 3 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `I switch to tab by title "..."`
- [ ] AC2: Step: `I close current tab`
- [ ] AC3: Step: `I close all tabs except main`
- [ ] AC4: Step: `I wait for new tab`
- [ ] AC5: Support 10+ tabs
- [ ] AC6: Documentation

**Technical Tasks:**
- [ ] Tab tracking in context
- [ ] Switch/close logic
- [ ] Wait for tab mechanism
- [ ] Tests + docs

**Dependencies:** Browser page management  
**DoD:** Same as 1.1

---

#### Story 14.4: Network Throttling & Geolocation
**Story Points:** 4 SP  
**Priority:** P2

**Acceptance Criteria:**
- [ ] AC1: Step: `I set network to "3G"`
- [ ] AC2: Presets: Fast3G, Slow3G, 4G, Offline
- [ ] AC3: Custom download/upload speeds
- [ ] AC4: Step: `I set geolocation to "Paris"`
- [ ] AC5: 10+ city presets
- [ ] AC6: Documentation

**Technical Tasks:**
- [ ] Rod/CDP network conditions
- [ ] Geolocation override
- [ ] Presets configuration
- [ ] Tests + docs

**Dependencies:** Browser CDP  
**DoD:** Same as 1.1

---

## 🎯 SPRINT 4: Polish & Launch (Semaine 7)

**Objectif Sprint:** Finalisation, polish, documentation, launch prep

**Story Points Total:** 26 SP  
**Features:** 10 bonus features + documentation + launch  
**Deliverable:** V1.0 Production-Ready

---

### EPIC 15: Quick Polish Features

**Epic Goal:** Améliorer UX et DX  
**Business Value:** Professional polish  
**Story Points:** 10 SP  
**Priority:** P2-P3

#### Story 15.1: Enhanced CLI UX
**Story Points:** 3 SP

**Features:**
- Color-coded test results enhanced
- Progress bar CLI
- Emoji support in reports ✅❌⏭️
- Interactive mode (choose tests)

---

#### Story 15.2: Auto-Retry Failed Tests
**Story Points:** 3 SP

**Acceptance Criteria:**
- [ ] Config: `retry_failed_tests: 3`
- [ ] CLI: `--retry-failed=3`
- [ ] Report shows retry count
- [ ] Stop on success
- [ ] Documentation

---

#### Story 15.3: Developer Tools
**Story Points:** 4 SP

**Features:**
- Config validation: `tkit validate-config`
- List steps: `tkit list-steps`
- Dry run: `tkit run --dry-run`
- Watch mode: `tkit watch`
- Duration warnings

---

### EPIC 16: Documentation & Examples

**Epic Goal:** Documentation complète pour V1 launch  
**Business Value:** Adoption facilitée  
**Story Points:** 8 SP  
**Priority:** P0

#### Story 16.1: Getting Started Guide
**Story Points:** 2 SP

**Content:**
- Installation (5 min)
- Quick start tutorial
- First test in 10 min
- Video walkthrough

---

#### Story 16.2: Feature Documentation
**Story Points:** 3 SP

**Content:**
- Allure reporting setup
- Data-driven testing guide
- Visual regression guide
- GraphQL advanced guide
- WebSocket testing guide
- All new features documented

---

#### Story 16.3: Examples Repository
**Story Points:** 3 SP

**Content:**
- 20+ example projects
- E-commerce example
- API testing example
- GraphQL example
- Visual regression example
- CI/CD integration examples

---

### EPIC 17: Launch Preparation

**Epic Goal:** Préparer V1.0 release  
**Business Value:** Go-to-market success  
**Story Points:** 8 SP  
**Priority:** P0

#### Story 17.1: Beta Testing
**Story Points:** 3 SP

**Tasks:**
- Recruit 50 beta testers
- Collect feedback
- Fix critical bugs
- Polish UX based on feedback

---

#### Story 17.2: Marketing Materials
**Story Points:** 3 SP

**Tasks:**
- Press release
- Blog post series (5 posts)
- Video demos (3 videos)
- Social media campaign
- Product Hunt preparation
- Comparison guides (vs Cypress, Playwright)

---

#### Story 17.3: Launch Execution
**Story Points:** 2 SP

**Tasks:**
- Product Hunt launch
- Hacker News post
- Reddit announcements
- Twitter/LinkedIn campaign
- Email existing users
- Monitor feedback

---

## 📊 Sprint Planning Summary

### Resource Allocation

**Team Composition:**
- 2 Senior Go Developers (full-time)
- 1 QA Engineer (part-time, testing)
- 1 Technical Writer (part-time, docs)
- 1 Product Manager (oversight)

**Capacity:**
- Sprint 1: 34 SP (2 devs × 2 weeks × 8.5 SP/dev/week)
- Sprint 2: 38 SP
- Sprint 3: 42 SP  
- Sprint 4: 26 SP (1 week)

**Total:** 140 SP over 7 weeks

---

### Risk Management

#### High Risks

**Risk 1: Visual Regression Performance**
- Impact: High
- Probability: Medium
- Mitigation: Early POC, optimize algorithm, parallel processing

**Risk 2: Allure Integration Complexity**
- Impact: Medium
- Probability: Low
- Mitigation: Good documentation, community support

**Risk 3: Scope Creep**
- Impact: High
- Probability: Medium
- Mitigation: Strict prioritization, defer to V1.1/V2.0

#### Medium Risks

**Risk 4: Beta Tester Availability**
- Impact: Medium
- Probability: Medium
- Mitigation: Start recruitment early, incentivize participation

**Risk 5: Documentation Time Underestimated**
- Impact: Medium
- Probability: Medium
- Mitigation: Start docs early, parallel with dev

---

### Dependencies Graph

```
Sprint 1 (Foundation)
├─ Allure Reporting (EPIC 1)
├─ Test Data Faker (EPIC 2)
├─ Performance Metrics (EPIC 3)
├─ Env Config (EPIC 4)
└─ CSV Data-Driven (EPIC 5)

Sprint 2 (Enhancement) - depends on Sprint 1
├─ Enhanced Debugging (EPIC 6) → Allure
├─ WebSocket Testing (EPIC 7)
├─ Skip/Pending (EPIC 8)
├─ Storage Management (EPIC 9)
└─ File Operations (EPIC 10)

Sprint 3 (Differentiation) - depends on Sprint 1 & 2
├─ Visual Regression (EPIC 11) → Performance, Allure
├─ GraphQL Advanced (EPIC 12)
├─ Accessibility (EPIC 13) → Allure
└─ Browser Features (EPIC 14)

Sprint 4 (Polish) - depends on Sprint 1-3
├─ Quick Polish (EPIC 15)
├─ Documentation (EPIC 16) → All features
└─ Launch (EPIC 17) → Beta testing
```

---

## ✅ Definition of Done (DoD)

### Code DoD
- [ ] Code written and reviewed (2 reviewers)
- [ ] Unit tests written (80%+ coverage)
- [ ] Integration tests written
- [ ] No linting errors
- [ ] Performance verified (benchmarks)
- [ ] Security scan passed
- [ ] Merged to main branch

### Feature DoD
- [ ] All acceptance criteria met
- [ ] Code DoD completed
- [ ] Documentation updated
- [ ] Examples created
- [ ] Changelog updated
- [ ] Release notes drafted

### Sprint DoD
- [ ] All stories completed (DoD)
- [ ] Sprint demo prepared
- [ ] Sprint retrospective held
- [ ] Next sprint planned
- [ ] Deployable increment ready

---

## 📈 Metrics & Tracking

### Sprint Metrics

**Velocity Tracking:**
- Target: 35 SP/sprint (2-week)
- Measure: Completed SP per sprint
- Adjust: Based on actual velocity

**Burndown:**
- Daily burndown chart
- Track remaining SP
- Identify blockers early

**Quality Metrics:**
- Code coverage: 80%+ target
- Bug count: < 5 critical bugs per sprint
- Test pass rate: 95%+

### Release Metrics

**Go-to-Market:**
- GitHub stars: 1,000+ in month 1
- Installations: 5,000+ in month 1
- Beta testers: 50+
- NPS score: > 50

**Quality:**
- Critical bugs: 0 at launch
- Test coverage: 85%+
- Documentation completeness: 100%

---

## 🎯 Next Steps

### Immediate Actions (This Week)

1. **Team Confirmation**
   - [ ] Confirm 2 Go developers
   - [ ] Confirm QA engineer availability
   - [ ] Confirm technical writer

2. **Sprint 1 Kickoff**
   - [ ] Sprint planning meeting
   - [ ] Story grooming session
   - [ ] Technical design review
   - [ ] Setup development environment

3. **Beta Program**
   - [ ] Create beta signup form
   - [ ] Reach out to potential testers
   - [ ] Prepare beta communication plan

4. **Launch Prep**
   - [ ] Draft press release
   - [ ] Start blog post series
   - [ ] Create demo videos outline

---

## 📝 Appendix

### Story Point Reference

**Story Points Mapping:**
- 1 SP: 2-4 hours (trivial)
- 2 SP: 4-8 hours (simple)
- 3 SP: 1 day (straightforward)
- 5 SP: 2-3 days (moderate)
- 8 SP: 3-5 days (complex)
- 13 SP: 1-2 weeks (very complex, consider splitting)

### Estimation Methodology

**Planning Poker:**
- Team estimates together
- Fibonacci sequence (1, 2, 3, 5, 8, 13)
- Discuss outliers
- Re-estimate if needed

---

**Document prepared by:** Product Manager (John) & Scrum Master  
**Date:** 2026-01-22  
**Version:** 1.0  
**Next Review:** Sprint 1 Planning Meeting  
**Status:** Ready for Team Review & Sprint Kickoff
