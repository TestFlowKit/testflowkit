# Story 3.2: Performance Thresholds & Warnings

**Epic:** EPIC 3 - Performance Monitoring  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 2 SP  
**Priority:** P1 - High  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** QA Engineer  
**I want** to be warned when operations exceed thresholds  
**So that** I can identify slow tests early

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Config: `performance.thresholds.api_response: 500ms`
  - YAML configuration
  - Multiple thresholds supported

- [ ] **AC2:** Warning in logs si > threshold
  - Colored warning message
  - Show actual vs threshold

- [ ] **AC3:** Highlighted in reports
  - Slow operations highlighted
  - Visual indication (red/yellow)

- [ ] **AC4:** Fail test option: `fail_on_slow: true`
  - Optional strict mode
  - Test fails if threshold exceeded

- [ ] **AC5:** Aggregate metrics per scenario
  - Total scenario duration
  - Average step duration
  - Slowest step identification

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Config schema update
- [ ] **Task 2:** Threshold checking logic
- [ ] **Task 3:** Warning/fail mechanisms
- [ ] **Task 4:** Report highlights
- [ ] **Task 5:** Tests + docs

---

## 📋 Dependencies

**Upstream:** Story 3.1 (Performance Metrics)

---

## ✓ Definition of Done

- [ ] All AC met
- [ ] Code merged
- [ ] Tests passing
- [ ] Documentation updated

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S3.2
