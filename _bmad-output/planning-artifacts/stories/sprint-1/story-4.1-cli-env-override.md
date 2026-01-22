# Story 4.1: CLI Environment Variable Override

**Epic:** EPIC 4 - Environment Configuration  
**Sprint:** Sprint 1 (Semaines 1-2)  
**Story Points:** 3 SP  
**Priority:** P0 - Critical  
**Assignee:** TBD  
**Status:** 📋 To Do

---

## 📖 User Story

**As a** Developer  
**I want** to override env variables via CLI  
**So that** I can test different environments without changing config files

---

## 🎯 Business Value

- **Impact:** Moyen - Flexibilité testing multi-env
- **User Benefit:** Switch env facilement
- **Efficiency:** Pas besoin modifier configs

---

## ✅ Acceptance Criteria

- [ ] **AC1:** Syntax `tkit run --env KEY=VALUE`
  - Parse correctly
  - Support KEY=VALUE format

- [ ] **AC2:** Multiple vars: `--env KEY1=VAL1 --env KEY2=VAL2`
  - Multiple --env flags
  - All applied correctly

- [ ] **AC3:** Override order: CLI > .env file > config.yaml
  - Priority resolution
  - CLI highest priority

- [ ] **AC4:** Variables accessible in Gherkin: `{{ env.KEY }}`
  - Variable substitution
  - Work in all steps

- [ ] **AC5:** Validation required vars
  - Check required variables present
  - Clear error messages

- [ ] **AC6:** Documentation + examples
  - Usage guide
  - Common scenarios

---

## 🔧 Technical Tasks

- [ ] **Task 1:** Extend CLI args parser (alexflint/go-arg)
  - Add EnvVars field (slice of strings)
  - Parse KEY=VALUE format

- [ ] **Task 2:** Environment merge logic
  - Load from config.yaml
  - Load from .env file
  - Override with CLI
  - Proper precedence

- [ ] **Task 3:** Priority resolution
  - Implement merge strategy
  - Test all scenarios

- [ ] **Task 4:** Variable substitution
  - Extend template system
  - Support {{ env.KEY }}

- [ ] **Task 5:** Validation
  - Required vars check
  - Error messages

- [ ] **Task 6:** Tests + docs
  - Unit tests
  - Integration tests
  - Documentation

---

## 📋 Dependencies

**Upstream:** None

---

## ✓ Definition of Done

- [ ] All AC met
- [ ] Code merged
- [ ] Tests passing
- [ ] Documentation with examples

---

## 📝 Example Usage

```bash
# Override base URL for staging
tkit run --env BASE_URL=https://staging.example.com

# Multiple overrides
tkit run --env API_URL=https://api.staging.com --env API_KEY=test123

# Use in Gherkin
Given I go to "{{ env.BASE_URL }}/login"
```

---

**Created:** 2026-01-22  
**Story ID:** TKIT-S4.1
