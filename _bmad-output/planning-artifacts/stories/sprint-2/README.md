# Sprint 2 - High-Value Features

**Duration:** Semaines 3-4  
**Story Points:** 38 SP  
**Focus:** Enhanced debugging et testing backend avancé

---

## 📊 Sprint Overview

**Objectif:** Améliorer capacités debugging et ajouter WebSocket, storage management

**Deliverable:**
- Browser Console Logs Capture 🐛
- Network Request Logging 🌐
- WebSocket Testing 💬
- Skip/Pending Tags ⏭️
- Cookie & Storage Management 🍪
- File Download & PDF Testing 📄

---

## 📋 Stories (10 stories)

### EPIC 6: Enhanced Debugging (13 SP)

**Story 6.1: Browser Console Logs Capture** - 5 SP - P1
- Capture console.log, .warn, .error, .info
- Display in HTML report with filtering
- Source file/line numbers
- Performance overhead < 5%

**Story 6.2: Network Request Logging** - 5 SP - P1
- Log all HTTP requests/responses
- Method, URL, headers, body, timing
- Filter by status code
- HAR export (optional)

**Story 6.3: Enhanced Error Messages** - 3 SP - P1
- Detailed error context
- Suggest fixes for common errors
- Show selectors tried
- Color-coded error types

---

### EPIC 7: WebSocket Testing (8 SP)

**Story 7.1: WebSocket Connection & Messaging** - 5 SP - P1
- Steps: connect, send, receive
- Support ws:// and wss://
- Timeout configurable
- Auto cleanup

**Story 7.2: WebSocket JSON Messages** - 3 SP - P1
- Send/receive JSON messages
- Field validation
- Variable extraction
- Pretty print logs

---

### EPIC 8: Test Control & Skip/Pending (5 SP)

**Story 8.1: Skip & Pending Tag Support** - 3 SP - P1
- Tags: @skip, @pending, @wip
- CLI: --skip-pending flag
- Reports show status separately
- Metrics tracking

**Story 8.2: Conditional Execution Tags** - 2 SP - P2
- @requires(condition)
- @env(staging,prod)
- Condition evaluation

---

### EPIC 9: Browser Storage Management (6 SP)

**Story 9.1: Cookie Management** - 3 SP - P2
- Steps: set, get, delete cookies
- Support all cookie attributes
- Variable storage
- Cross-step persistence

**Story 9.2: Local/Session Storage** - 3 SP - P2
- localStorage and sessionStorage steps
- Support JSON values
- Clear storage
- Variable integration

---

### EPIC 10: File Operations (6 SP)

**Story 10.1: File Download Validation** - 3 SP - P2
- Download and validate files
- Check name, size, content
- Auto cleanup temp files
- Support multiple formats

**Story 10.2: PDF Content Validation** - 3 SP - P2
- Extract PDF text
- Page count, metadata
- Content validation
- Support encrypted PDFs

---

## ✅ Sprint Goals

- [ ] Console logs visible in all reports
- [ ] Network requests fully logged
- [ ] WebSocket testing functional
- [ ] Skip/Pending tags working
- [ ] Cookie/storage management complete
- [ ] File download & PDF testing ready
- [ ] All P1 stories completed
- [ ] Sprint demo prepared

---

## 📈 Velocity Target

**Planned:** 38 SP  
**Team Capacity:** 2 developers × 2 weeks × 9.5 SP/dev/week = 38 SP  
**Confidence:** Medium-High

---

## 🎯 Success Criteria

- All P1 stories completed
- Enhanced debugging working end-to-end
- WebSocket examples documented
- Code coverage 80%+
- Ready for Sprint 3

---

**Dependencies:**
- Sprint 1 must be completed
- Allure reporter needed for enhanced debugging display

---

**Created:** 2026-01-22
