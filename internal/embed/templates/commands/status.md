# /status: Report Project State

## Purpose

Provide AI with quick understanding of current project state.

## Process

### Step 1: Read Project Context

Read `.claude/CLAUDE.md` for:
- Project overview
- Architecture philosophy
- Development guidelines

### Step 2: Scan Services Directory

For each module in `services/`:
1. Read `INTERFACE.md`
2. Check if `src/` exists and has implementation
3. Check if tests exist in `tests/services/`

### Step 3: Report Status

Format as:
```
## Project Status

**Modules Found:**
- [module-name]: [CREATED|PARTIAL|TESTS_PENDING]
  Purpose: [from INTERFACE.md]
  Interface: [complete/incomplete]
  Tests: [present/missing]

**Recent Work:**
[Summarize from project state markers or git if needed]

**Active Context:**
- Language: [detected language]
- Active module: [from CLAUDE.md if specified]
```

### Step 4: Test Status Summary

Check if tests pass:
- Identify appropriate test command
- Run tests and report results
- Note any failing tests
