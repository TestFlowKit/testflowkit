# Story 5.2: Excel Data Source Support

**Epic:** EPIC 5 - CSV/Excel Data-Driven Testing  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 3 SP  
**Priority:** P2 - Medium  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** Business Analyst  
**I want** to use Excel files as data source  
**So that** I can manage test data in familiar tools

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Syntax: `Examples: @data(testdata.xlsx)`
  - Parse .xlsx files
  - Load correctly

- [ ] **AC2:** Support .xlsx format
  - Modern Excel format
  - Read cells correctly

- [ ] **AC3:** Multiple sheets support
  - Detect sheets
  - Default to first sheet

- [ ] **AC4:** Sheet selection: `@data(file.xlsx:Sheet1)`
  - Specify sheet by name
  - Error if sheet not found

- [ ] **AC5:** Performance acceptable (< 2s for 1000 rows)
  - Benchmark
  - Optimize if needed

- [ ] **AC6:** Documentation
  - Excel format guide
  - Examples
  - Sheet naming conventions

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Install `github.com/xuri/excelize`
- [ ] **Task 2:** Excel parser implementation
- [ ] **Task 3:** Sheet handling
- [ ] **Task 4:** Integration with CSV logic
- [ ] **Task 5:** Tests + docs

---

## 📋 Dependencies

**Upstream:** Story 5.1 (CSV Data-Driven)

---

## ✓ Definition of Done

- [ ] All AC met
- [ ] Performance verified
- [ ] Code merged
- [ ] Tests passing
- [ ] Documentation

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S5.2
