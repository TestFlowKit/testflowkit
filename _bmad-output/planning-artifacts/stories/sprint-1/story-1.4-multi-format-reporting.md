# Story 1.4: Multi-Format Report Generation

**Epic:** EPIC 1 - Professional Reporting System  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 2 SP  
**Priority:** P1 - High  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** Developer  
**I want** to generate multiple report formats simultaneously  
**So that** I can use different tools for different purposes

---

## 🎯 Business Value

- **Impact:** Moyen - Flexibilité pour équipes
- **User Benefit:** Un run → plusieurs rapports
- **Efficiency:** Éviter multiple exécutions

---

## ✅ Acceptance Criteria

- [ ] **AC1:** CLI flag `--report-format=html,allure,cucumber-json`
  - Comma-separated format list
  - All formats generated

- [ ] **AC2:** Multiple reporters run in parallel
  - Goroutines for each reporter
  - No blocking between reporters

- [ ] **AC3:** All formats generated in < 3s
  - Performance benchmarks
  - Parallel execution optimized

- [ ] **AC4:** Error handling per reporter
  - One reporter failure doesn't stop others
  - Errors logged clearly
  - Partial success acceptable

- [ ] **AC5:** Config file support for default formats
  - `report_formats: [html, allure]` in YAML
  - CLI overrides config
  - Documentation

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Reporter factory pattern
  - Factory creates multiple reporters
  - Registry of available formats

- [ ] **Task 2:** Parallel report generation
  - WaitGroup for goroutines
  - Error collection
  - Timeout handling

- [ ] **Task 3:** Config integration
  - Add report_formats to config schema
  - Load from config file
  - Merge with CLI flags

- [ ] **Task 4:** Error handling
  - Collect errors from all reporters
  - Log warnings for failures
  - Success if at least one succeeds

- [ ] **Task 5:** Performance testing
  - Benchmark parallel vs sequential
  - Optimize resource usage

---

## 📋 Dependencies

**Upstream Dependencies:**
- Story 1.1 (Allure) - MUST be completed
- Story 1.3 (Cucumber JSON) - MUST be completed

---

## ✓ Definition of Done

- [ ] All acceptance criteria met
- [ ] Performance < 3s verified
- [ ] Code merged to main
- [ ] Tests passing
- [ ] Documentation updated

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S1.4
