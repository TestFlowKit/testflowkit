# Story 5.1: CSV Data Source for Scenario Outline

**Epic:** EPIC 5 - CSV/Excel Data-Driven Testing  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 5 SP  
**Priority:** P1 - High  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** QA Engineer  
**I want** to use CSV files as data source for Scenario Outline  
**So that** I can test multiple cases without duplicating Gherkin

---

## 🎯 Business Value

- **Impact:** Élevé - Data-driven testing essentiel
- **User Benefit:** Éviter duplication, gérer données dans CSV
- **Market Differentiation:** External data sources

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Syntax: `Examples: @data(users.csv)`
  - Parse @data tag
  - Load CSV file
  - Map to Examples

- [ ] **AC2:** CSV with headers mapped to variables
  - First row = headers
  - Headers → Gherkin variables
  - Case-insensitive matching

- [ ] **AC3:** Support large files (10,000+ rows)
  - Performance optimization
  - Streaming parsing
  - Memory efficient

- [ ] **AC4:** Skip invalid rows with warnings
  - Validate row structure
  - Log warnings for skipped rows
  - Continue execution

- [ ] **AC5:** Report shows data source
  - CSV filename in report
  - Row number for each iteration
  - Traceability

- [ ] **AC6:** Performance: load CSV < 1s
  - Benchmark with large files
  - Optimize parsing
  - Caching if needed

- [ ] **AC7:** Documentation + examples
  - CSV format guide
  - Example files
  - Common patterns

---

## 🔧 Technical Tasks

- [ ] **Task 1:** CSV parser (stdlib encoding/csv)
  - Implement CSV reader
  - Header parsing
  - Row iteration

- [ ] **Task 2:** Examples tag parser extension
  - Detect @data(filename)
  - Extract filename
  - Load file

- [ ] **Task 3:** Godog integration
  - Generate Scenario Outline table
  - Map CSV → Examples table
  - Execute scenarios

- [ ] **Task 4:** Row validation
  - Check column count
  - Validate data types
  - Skip malformed rows

- [ ] **Task 5:** Performance optimization
  - Benchmark 10K rows
  - Streaming if needed
  - Memory profiling

- [ ] **Task 6:** Tests (including 10K rows)
  - Unit tests
  - Integration tests
  - Performance tests

- [ ] **Task 7:** Documentation
  - CSV format specification
  - Examples
  - Best practices

---

## 📋 Dependencies

**Upstream:** Macro system, Godog

**Downstream:** Story 5.2 (Excel Support)

---

## ⚠️ Risks & Mitigation

**Risk 1: Large CSV files performance**
- Probability: Medium
- Impact: Medium
- Mitigation: Streaming parser, benchmarks

**Risk 2: CSV encoding issues**
- Probability: Medium
- Impact: Low
- Mitigation: UTF-8 default, encoding detection

---

## ✓ Definition of Done

- [ ] All AC met
- [ ] Performance < 1s for 10K rows
- [ ] Code merged
- [ ] Tests passing
- [ ] Documentation + examples

---

## 📝 Example Usage

**users.csv:**
```csv
email,password,expectedResult
test1@example.com,pass123,success
test2@example.com,wrong,failure
admin@example.com,admin123,success
```

**test.feature:**
```gherkin
Scenario Outline: Login with different users
  Given I go to the login page
  When I enter "<email>" into the email field
  And I enter "<password>" into the password field
  And I click the login button
  Then I should see "<expectedResult>"

  Examples: @data(users.csv)
```

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S5.1
