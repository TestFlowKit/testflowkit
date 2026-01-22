# Story 2.1: Faker Library Integration

**Epic:** EPIC 2 - Test Data Generation  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 5 SP  
**Priority:** P0 - Critical  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** QA Engineer  
**I want** to generate random test data (emails, names, phones)  
**So that** I can avoid hardcoding values and improve test coverage

---

## 🎯 Business Value

- **Impact:** Élevé - Réutilisabilité tests, données réalistes
- **User Benefit:** Éviter hardcoding, tests plus robustes
- **Market Differentiation:** Data generation built-in

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Integration `github.com/brianvoe/gofakeit`
  - Dependency installed
  - Library initialized properly

- [ ] **AC2:** New steps in `internal/step_definitions/variables/faker.go`
  - Clean file structure
  - Proper package organization

- [ ] **AC3:** Step: `I set "field" to random email`
  - Generates valid email format
  - Stores in scenario context

- [ ] **AC4:** Support types: email, name, phone, address, UUID, date, number
  - Minimum 20 data types supported
  - All types documented
  - Examples for each type

**Supported Types:**
- email, name, firstName, lastName
- phone, address, city, country, zipCode
- uuid, guid
- date, time, datetime
- number, integer, float
- url, domain, ipAddress
- company, jobTitle
- password, username
- creditCard, SSN
- color, hexColor

- [ ] **AC5:** Values stored in scenario context variables
  - Integration avec variable system existant
  - Accessible dans steps suivants
  - Proper scoping (scenario-level)

- [ ] **AC6:** Reproducible with seed: `--faker-seed=12345`
  - CLI flag for seed
  - Same seed → same data
  - Useful for debugging

- [ ] **AC7:** Documentation with 20+ examples
  - Step library updated
  - Examples in docs site
  - Common use cases covered

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Install gofakeit dependency
  ```bash
  go get github.com/brianvoe/gofakeit/v6
  ```

- [ ] **Task 2:** Create faker.go step definitions
  - Create `internal/step_definitions/variables/faker.go`
  - Initialize gofakeit
  - Register steps with Godog

- [ ] **Task 3:** Implement 20+ faker types
  - Create step for each type
  - Use gofakeit generators
  - Handle edge cases

- [ ] **Task 4:** Variable storage integration
  - Store generated values in context
  - Follow existing variable patterns
  - Support variable substitution

- [ ] **Task 5:** Seed support CLI flag
  - Add `--faker-seed` to args
  - Initialize gofakeit with seed
  - Document reproducibility

- [ ] **Task 6:** Unit tests (90%+ coverage)
  - Test each faker type
  - Test seed reproducibility
  - Test variable storage

- [ ] **Task 7:** Documentation + examples
  - Update step library docs
  - Add faker section to docs site
  - Create example feature files

---

## 📋 Dependencies

**Upstream Dependencies:**
- Variable system existant

**Downstream Dependencies:**
- Story 2.2 (Custom Faker Formats)

---

## ⚠️ Risks & Mitigation

**Risk 1: Generated data not realistic enough**
- Probability: Low
- Impact: Low
- Mitigation: gofakeit is mature library with quality data

**Risk 2: Seed not working across platforms**
- Probability: Low
- Impact: Low
- Mitigation: Test on Mac/Linux/Windows

---

## ✓ Definition of Done

- [ ] All acceptance criteria met
- [ ] 20+ data types supported
- [ ] Code merged to main
- [ ] Tests passing (90%+ coverage)
- [ ] Documentation complete with examples
- [ ] Peer review approved
- [ ] Demo with real feature files

---

## 📝 Example Usage

```gherkin
Scenario: Register new user with random data
  Given I open the registration page
  When I set "email" to random email
  And I set "firstName" to random firstName
  And I set "lastName" to random lastName
  And I set "phone" to random phone
  And I set "password" to random password
  And I enter "{{ email }}" into the "email" field
  And I enter "{{ firstName }}" into the "firstName" field
  And I click the "submit" button
  Then I should see "Registration successful"
```

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S2.1
