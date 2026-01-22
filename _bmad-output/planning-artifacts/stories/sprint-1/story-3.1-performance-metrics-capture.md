# Story 3.1: Basic Performance Metrics Capture

**Epic:** EPIC 3 - Performance Monitoring  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 3 SP  
**Priority:** P0 - Critical  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** Performance Engineer  
**I want** to capture response times and page load times  
**So that** I can monitor performance regressions

---

## 🎯 Business Value

- **Impact:** Moyen-Élevé - Performance visibility
- **User Benefit:** Détection régressions performance
- **Market Differentiation:** Performance testing built-in BDD

---

## ✅ Acceptance Criteria

- [ ] **AC1:** API response time logged per request
  - Capture timing for all HTTP requests
  - Precision millisecond
  - Include in logs

- [ ] **AC2:** Page load time captured per navigation
  - DOMContentLoaded timing
  - window.load timing
  - First contentful paint (optional)

- [ ] **AC3:** Step execution time in reports
  - Each step shows duration
  - Slowest steps highlighted
  - Aggregate metrics per scenario

- [ ] **AC4:** Performance struct in Report model
  - Add Performance field to models
  - Store all metrics
  - Serialize to JSON

- [ ] **AC5:** Metrics in HTML/JSON reports
  - Display in HTML template
  - Include in JSON output
  - Summary statistics

- [ ] **AC6:** Overhead < 1% execution time
  - Benchmarks to verify
  - Minimal impact on tests
  - Optimized timing code

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Add Performance struct to models
  ```go
  type Performance struct {
      APIResponseTimes    []ResponseTime
      PageLoadTimes       []PageLoadTime
      StepDurations       []StepDuration
      TotalDuration       time.Duration
  }
  ```

- [ ] **Task 2:** Timer integration in HTTP client
  - Wrap HTTP client
  - Start timer before request
  - Stop timer after response
  - Store timing

- [ ] **Task 3:** Page load time capture (Rod)
  - Use Rod timing API
  - Capture navigation events
  - Store metrics

- [ ] **Task 4:** Report rendering
  - Update HTML template
  - Add performance section
  - Charts/graphs (optional)

- [ ] **Task 5:** Performance testing
  - Benchmark overhead
  - Verify < 1% impact
  - Load testing

- [ ] **Task 6:** Documentation
  - Performance metrics guide
  - Interpretation guide
  - Examples

---

## 📋 Dependencies

**Upstream Dependencies:**
- Reporters existants

**Downstream Dependencies:**
- Story 3.2 (Performance Thresholds)

---

## ✓ Definition of Done

- [ ] All acceptance criteria met
- [ ] Overhead verified < 1%
- [ ] Code merged to main
- [ ] Tests passing
- [ ] Documentation updated
- [ ] Performance visible in reports

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S3.1
