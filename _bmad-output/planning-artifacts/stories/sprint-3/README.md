# Sprint 3 - Differentiators

**Duration:** Semaines 5-6  
**Story Points:** 42 SP  
**Focus:** Features de différenciation marché

---

## 📊 Sprint Overview

**Objectif:** Implémenter les features qui différencient TestFlowKit du marché

**Deliverable:**
- Visual Screenshot Comparison 📸
- GraphQL Schema Validation 🔍
- Accessibility Testing ♿
- iFrame & Shadow DOM Support 🎭
- Network Throttling & Geolocation 🌍

---

## 📋 Stories (12 stories)

### EPIC 11: Visual Regression Testing (13 SP)

**Story 11.1: Screenshot Baseline Management** - 5 SP - P3
- Save screenshot baselines
- Baselines stored in `baselines/` directory
- Versioning support
- Update baseline capability
- CLI management commands

**Story 11.2: Visual Comparison Algorithm** - 5 SP - P3
- Pixel-by-pixel comparison
- Configurable tolerance (0-100%)
- Generate diff images
- Highlight differences
- Performance < 500ms per comparison
- Similarity percentage reporting

**Story 11.3: Ignore Regions for Dynamic Content** - 3 SP - P3
- Define ignore regions by selector or coordinates
- Mask regions in comparison
- Visual indication in diff
- Support multiple ignore regions

---

### EPIC 12: GraphQL Advanced Testing (8 SP)

**Story 12.1: GraphQL Schema Introspection** - 3 SP - P3
- Step: `the GraphQL schema should be valid`
- Introspect schema automatically
- Validate types and fields
- Cache schema for performance
- Query optimization

**Story 12.2: GraphQL Schema Diff & Breaking Changes** - 5 SP - P3
- Save baseline schema
- Compare schemas (baseline vs current)
- Detect breaking changes (field removal, type change)
- Detect non-breaking changes (new fields)
- Report with severity levels
- Fail on breaking changes option

---

### EPIC 13: Accessibility Testing (8 SP)

**Story 13.1: Basic WCAG AA Checks** - 5 SP - P2
- Step: `the page should be accessible`
- Inject axe-core library
- Check: alt text, labels, contrast, ARIA
- Detailed violation report
- Configurable severity (A, AA, AAA)
- 80%+ violation detection rate

**Story 13.2: Accessibility Report Integration** - 3 SP - P3
- Violations in HTML report
- Violations in Allure report
- Aggregate metrics
- Trends over time
- Executive summary

---

### EPIC 14: Advanced Browser Features (13 SP)

**Story 14.1: iFrame Support** - 3 SP - P2
- Steps: switch to iframe, switch to parent
- Support by name, index, selector
- Nested iframes (3+ levels)
- Auto-detect iframe context
- Proper cleanup

**Story 14.2: Shadow DOM Support** - 3 SP - P2
- Selector syntax: `shadow::#element`
- Nested shadow roots
- Auto-detect shadow DOM
- Fallback to light DOM
- Performance optimized

**Story 14.3: Enhanced Multi-Tab Management** - 3 SP - P2
- Switch tab by title, index, URL
- Close tabs (current, all except main)
- Wait for new tab
- Support 10+ tabs
- Tab tracking in context

**Story 14.4: Network Throttling & Geolocation** - 4 SP - P2
- Network throttling: 3G, 4G, Fast3G, Slow3G, Offline
- Custom download/upload speeds
- Geolocation mocking
- 10+ city presets (Paris, NYC, Tokyo, etc.)
- Latitude/longitude precision
- Permission API integration

---

## ✅ Sprint Goals

- [ ] Visual regression testing working with baselines
- [ ] GraphQL schema validation functional
- [ ] Accessibility checks integrated
- [ ] iFrame & Shadow DOM fully supported
- [ ] Network throttling & geolocation working
- [ ] All differentiator features demoed
- [ ] Documentation complete with examples
- [ ] Sprint demo prepared

---

## 📈 Velocity Target

**Planned:** 42 SP  
**Team Capacity:** 2 developers × 2 weeks × 10.5 SP/dev/week = 42 SP  
**Confidence:** Medium

---

## 🎯 Success Criteria

- Visual regression reduces UI bugs 50%+
- GraphQL schema validation working
- Accessibility detection 80%+ WCAG violations
- All advanced browser features functional
- Code coverage 80%+
- Market differentiators clearly visible

---

## ⚠️ Risks

**Risk 1: Visual regression complexity**
- Mitigation: Early POC, iterative development

**Risk 2: Accessibility false positives**
- Mitigation: axe-core mature library, configurable rules

---

**Dependencies:**
- Sprint 1 & 2 must be completed
- Allure reporter for accessibility integration
- Performance metrics for comparison benchmarks

---

**Created:** 2026-01-22
