# Story 2.2: Custom Faker Formats

**Epic:** EPIC 2 - Test Data Generation  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 3 SP  
**Priority:** P1 - High  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** Developer  
**I want** to specify custom formats for generated data  
**So that** I can match specific validation rules

---

## 🎯 Business Value

- **Impact:** Moyen - Flexibilité pour cas spécifiques
- **User Benefit:** Données conformes aux règles métier
- **Examples:** Phone formats par pays, date formats spécifiques

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Step: `I set "field" to random phone "US format"`
  - Parse format parameter
  - Generate matching data

- [ ] **AC2:** Step: `I set "field" to random date "YYYY-MM-DD"`
  - Support date format strings
  - Common formats: ISO, US, EU

- [ ] **AC3:** Support regex patterns
  - Step: `I set "field" to random pattern "[A-Z]{3}-[0-9]{4}"`
  - Regex validation
  - Generate matching strings

- [ ] **AC4:** Custom templates: `random "###-###-####"`
  - # → digit, @ → letter, * → alphanumeric
  - Easy template syntax

- [ ] **AC5:** Documentation formats disponibles
  - All supported formats listed
  - Examples per format
  - Common patterns library

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Format parsing logic
  - Parse format parameter from step
  - Validate format string

- [ ] **Task 2:** Template engine integration
  - Implement template parser
  - ### → digits, @@@ → letters

- [ ] **Task 3:** Regex support
  - Use gofakeit regex generator
  - Validate regex patterns

- [ ] **Task 4:** Examples + tests
  - Test common formats
  - Edge cases

- [ ] **Task 5:** Documentation
  - Format reference guide
  - Examples library

---

## 📋 Dependencies

**Upstream Dependencies:**
- Story 2.1 (Faker Library) - MUST be completed first

---

## ✓ Definition of Done

- [ ] All acceptance criteria met
- [ ] Code merged to main
- [ ] Tests passing
- [ ] Documentation with 10+ format examples

---

## 📝 Example Usage

```gherkin
# US Phone format
I set "phone" to random phone "US format"
# Result: (555) 123-4567

# Date format
I set "startDate" to random date "YYYY-MM-DD"
# Result: 2026-03-15

# Custom pattern
I set "orderNumber" to random pattern "ORD-[0-9]{6}"
# Result: ORD-123456

# Template
I set "productCode" to random "ABC-###-@@@@"
# Result: ABC-742-XKPL
```

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S2.2
