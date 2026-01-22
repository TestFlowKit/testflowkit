# Story 1.3: Cucumber JSON Export

**Epic:** EPIC 1 - Professional Reporting System  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 3 SP  
**Priority:** P0 - Critical  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** DevOps Engineer  
**I want** to export test results in Cucumber JSON format  
**So that** I can integrate with Jenkins/GitLab CI/CD plugins

---

## 🎯 Business Value

- **Impact:** Élevé - CI/CD integration essentielle
- **User Benefit:** Compatibilité outils standard (Jenkins, GitLab)
- **Market Differentiation:** Intégration seamless CI/CD

---

## ✅ Acceptance Criteria

- [ ] **AC1:** CLI flag `--report-format=cucumber-json`
  - Parse flag correctly
  - Generate Cucumber JSON output

- [ ] **AC2:** Output file `cucumber-report.json` in standard format
  - Standard Cucumber JSON schema
  - Proper structure: features → scenarios → steps

- [ ] **AC3:** Compatible with Jenkins Cucumber Reports Plugin
  - Tested with Jenkins plugin
  - All fields mapped correctly
  - Screenshots included

- [ ] **AC4:** Compatible with GitLab Test Reports
  - GitLab CI can parse report
  - Test summary displayed in MR
  - Failed tests highlighted

- [ ] **AC5:** All features, scenarios, steps mapped correctly
  - Feature metadata preserved
  - Scenario tags included
  - Step definitions with status

- [ ] **AC6:** Tags and metadata included
  - Gherkin tags in JSON
  - Custom metadata fields
  - Environment info

- [ ] **AC7:** Timestamps in ISO 8601 format
  - Start time, end time per scenario
  - Duration calculated
  - Timezone UTC

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Create `pkg/reporters/cucumber/cucumber_json.go`
  - CucumberJSONReporter struct
  - Implement Reporter interface

- [ ] **Task 2:** Implement Cucumber JSON schema
  - Feature struct
  - Scenario struct
  - Step struct
  - Proper JSON tags

- [ ] **Task 3:** Map TestResult → Cucumber format
  - Mapper functions
  - Handle all test states
  - Convert timestamps

- [ ] **Task 4:** CLI flag integration
  - Extend args parser
  - Support in config file

- [ ] **Task 5:** Validation against schema
  - JSON schema validation
  - Test with real CI tools

- [ ] **Task 6:** Integration tests with Jenkins/GitLab
  - Sample Jenkins pipeline
  - Sample GitLab CI config
  - Verify plugin compatibility

- [ ] **Task 7:** Documentation
  - CI/CD integration guide
  - Examples per platform
  - Troubleshooting

---

## 📋 Dependencies

**Upstream Dependencies:**
- JSON reporter existant (réutiliser structure)

**Downstream Dependencies:**
- Story 1.4 (Multi-Format Reporting)

---

## ✓ Definition of Done

- [ ] All acceptance criteria met
- [ ] Code merged to main
- [ ] Tests passing (80%+ coverage)
- [ ] Validated with Jenkins plugin
- [ ] Validated with GitLab CI
- [ ] Documentation updated
- [ ] CI/CD examples provided

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S1.3
