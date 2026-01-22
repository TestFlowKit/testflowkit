# Story 1.2: Allure Historical Trends

**Epic:** EPIC 1 - Professional Reporting System  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 3 SP  
**Priority:** P0 - Critical  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** Test Manager  
**I want** to see historical test trends in Allure  
**So that** I can track quality improvements over time

---

## 🎯 Business Value

- **Impact:** Élevé - Visibilité long-terme sur qualité
- **User Benefit:** Trends passé/échec, performance dans le temps
- **Market Differentiation:** Historique professionnel

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Test execution history stored in `history/` directory
  - Create `allure-results/history/` on first run
  - Preserve history between runs
  - Clean old history (configurable retention: default 30 days)

- [ ] **AC2:** Allure history.json generated per run
  - Generate `history.json` with previous results
  - Include: test name, status, duration, timestamp
  - Proper JSON schema for Allure consumption

- [ ] **AC3:** Trends visible in Allure report (pass/fail rates)
  - Pass rate trend over time
  - Duration trend per test
  - Flakiness detection (intermittent failures)

- [ ] **AC4:** Retries tracked in history
  - Retry count per test
  - Success after retry indication
  - Flaky test identification

- [ ] **AC5:** Duration trends visible
  - Performance regression detection
  - Slowest tests highlighted
  - Duration comparison with previous runs

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Implement history storage logic
  - Create HistoryManager component
  - Load previous history on startup
  - Merge with current results

- [ ] **Task 2:** Generate history.json file
  - Parse previous results
  - Extract historical data
  - Format for Allure

- [ ] **Task 3:** Trend calculation algorithm
  - Calculate pass rate trends
  - Identify performance regressions
  - Detect flaky tests (success rate < 95%)

- [ ] **Task 4:** Documentation
  - History configuration options
  - Retention policy
  - Examples

---

## 📋 Dependencies

**Upstream Dependencies:**
- Story 1.1 (Allure JSON Format) - MUST be completed first

**Downstream Dependencies:**
- None

---

## ⚠️ Risks & Mitigation

**Risk 1: History file growth**
- Probability: Medium
- Impact: Low
- Mitigation: Configurable retention, cleanup old data

---

## ✓ Definition of Done

- [ ] All acceptance criteria met
- [ ] Code merged to main
- [ ] Tests passing (80%+ coverage)
- [ ] Documentation updated
- [ ] Peer review approved
- [ ] History visible in Allure UI

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S1.2
