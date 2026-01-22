# Story 1.1: Allure JSON Format Support

**Epic:** EPIC 1 - Professional Reporting System  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 5 SP  
**Priority:** P0 - Critical  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** QA Engineer  
**I want** to generate test results in Allure JSON format  
**So that** I can use Allure reporting framework for professional visualizations

---

## 🎯 Business Value

- **Impact:** Très élevé - Reporting professionnel enterprise-grade
- **User Benefit:** Visualisations riches, historique trends, catégorisation
- **Market Differentiation:** Standard industrie pour reporting

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Reporter interface extended with AllureReporter implementation
  - AllureReporter implements `pkg/reporters/Reporter` interface
  - Clean integration with existing reporter factory

- [ ] **AC2:** Integration of `github.com/allure-framework/allure-go` library
  - Dependency added to go.mod
  - Library properly imported and initialized

- [ ] **AC3:** CLI flag `--report-format=allure` generates Allure JSON
  - New CLI flag parsed correctly
  - Flag triggers AllureReporter execution
  - Default output directory: `allure-results/`

- [ ] **AC4:** All test metadata mapped to Allure format (status, timing, steps)
  - Test status: passed, failed, skipped, broken
  - Timing: start time, duration
  - Steps with individual status and timing
  - Features and scenarios properly mapped

- [ ] **AC5:** Screenshots attached to Allure reports
  - Screenshots from failed tests attached
  - Proper MIME type and encoding
  - Displayed correctly in Allure UI

- [ ] **AC6:** Tags/features categorization in Allure
  - Gherkin tags → Allure labels
  - Feature categorization
  - Severity levels support

- [ ] **AC7:** Unit tests for AllureReporter (80%+ coverage)
  - Test happy path
  - Test error scenarios
  - Test edge cases
  - Coverage verified with go test -cover

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Install allure-go dependency
  ```bash
  go get github.com/allure-framework/allure-go
  ```

- [ ] **Task 2:** Create `pkg/reporters/allure/allure_reporter.go`
  - Define AllureReporter struct
  - Implement Reporter interface methods
  - Handle file I/O for JSON output

- [ ] **Task 3:** Implement Reporter interface
  - `GenerateReport(result *TestResult) error`
  - `GetFormat() string`
  - `SetOutputDir(dir string)`

- [ ] **Task 4:** Map TestResult → AllureResult
  - Create mapper functions
  - Handle all test states
  - Convert timestamps
  - Map Gherkin features to Allure structure

- [ ] **Task 5:** Attach screenshots to results
  - Read screenshot files
  - Base64 encoding
  - Attach to step/test
  - Proper MIME type

- [ ] **Task 6:** Add CLI flag parsing
  - Extend `cmd/testflowkit/args.config.go`
  - Add `ReportFormat` field
  - Support multiple formats (comma-separated)

- [ ] **Task 7:** Write unit tests
  - Test AllureReporter.GenerateReport()
  - Test mapping functions
  - Test screenshot attachment
  - Mock filesystem for tests

- [ ] **Task 8:** Documentation update
  - Update README.md with Allure instructions
  - Add examples to documentation site
  - CLI help text update

---

## 📋 Dependencies

**Upstream Dependencies:**
- None - Foundation feature

**Downstream Dependencies:**
- Story 1.2 (Allure Historical Trends) depends on this
- Story 1.4 (Multi-Format Reporting) depends on this

**External Dependencies:**
- `github.com/allure-framework/allure-go` library
- Allure CLI for report generation (user-installed)

---

## ⚠️ Risks & Mitigation

**Risk 1: Library API changes**
- Probability: Low
- Impact: Medium
- Mitigation: Pin specific version, monitor releases

**Risk 2: Complex test result mapping**
- Probability: Medium
- Impact: Medium
- Mitigation: Incremental development, comprehensive tests

---

## ✓ Definition of Done

- [ ] All acceptance criteria met and verified
- [ ] Code written and peer reviewed (2 reviewers)
- [ ] Unit tests written with 80%+ coverage
- [ ] Integration tests written
- [ ] No linting errors (`golangci-lint run`)
- [ ] Performance verified (report generation < 2s)
- [ ] Security scan passed
- [ ] Documentation updated (README, CLI help, docs site)
- [ ] Code merged to main branch
- [ ] Demo prepared for sprint review

---

## 📝 Notes

**Implementation Notes:**
- Use allure-go's ResultWriter for JSON generation
- Store results in `allure-results/` by default
- Each test execution creates UUID-named JSON files
- Support both `--report-format=allure` and config file setting

**Testing Notes:**
- Use golden files for expected JSON output
- Test with real Godog test results
- Verify Allure UI can consume generated JSON

**Performance Considerations:**
- JSON serialization should be fast (< 100ms)
- Batch write operations
- No blocking on report generation

---

## 🔗 Related Links

- [Allure Framework Documentation](https://docs.qameta.io/allure/)
- [allure-go Library](https://github.com/allure-framework/allure-go)
- [EPIC 1: Professional Reporting System](../../../epics-and-stories.md#epic-1-professional-reporting-system)

---

**Created:** 2026-01-22  
**Last Updated:** 2026-01-22  
**Story ID:** TKIT-S1.1
