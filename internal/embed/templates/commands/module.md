# /module: Create or Work on a Deep Module

## Purpose

Unified command for module lifecycle:
- If module doesn't exist → Create new module
- If module exists → Work on existing module

## Process

### Step 1: Detect Module State

Ask user for module name, then:
1. Check if `services/[module-name]/` exists
2. If NOT exists → Go to Step 2 (Create)
3. If exists → Go to Step 4 (Work)

### Step 2: Create Module (New)

#### 2a. Module Definition

Ask user for:
- Module name (e.g., "auth", "database")
- Brief purpose
- Key responsibilities

#### 2b. Interface Design

Generate `INTERFACE.md` template with:
- Purpose
- Public Interface (Types, Functions)
- Constraints & Invariants

**Get human approval before proceeding.**

#### 2c. Scaffold Structure

Create:
```
services/[module-name]/
├── [module entry point]   # Public types + exports
├── INTERFACE.md            # Approved interface
└── src/
    ├── core impl
    └── private helpers

tests/services/
└── [boundary test file]
```

#### 2d. Write Boundary Tests

- Test all public functions
- Test edge cases
- Lock in expected behavior

#### 2e. Initial Implementation

Implement to satisfy both INTERFACE.md and boundary tests

### Step 3: Work on Module (Existing)

#### 3a. Read Contract

Read and understand:
- `services/[module-name]/INTERFACE.md`
- `tests/services/[boundary test file]`
- `services/[module-name]/[entry point]`

**Do not modify INTERFACE.md unless evolving contract.**

#### 3b: Identify Changes

Ask user what needs changing:
- Bug fixes
- Performance improvements
- Internal refactoring
- Feature additions (within scope)

#### 3c: Modify Implementation

Only modify files in `services/[module-name]/src/`

#### 3d: Run Boundary Tests

Use appropriate test command for detected language.

#### 3e: Verify

- Boundary tests pass
- Interface contract maintained
- No breaking changes

## Success Criteria

Create mode:
- [ ] Interface approved
- [ ] Structure created
- [ ] Boundary tests written
- [ ] Implementation passes tests

Work mode:
- [ ] Boundary tests pass
- [ ] Contract unchanged (or evolved with approval)
- [ ] Implementation improved

## AI Inference Guidelines

**Language Detection:**
- Examine file extensions, package management files, build config
- Adapt file patterns (.go, .py, .ts, .java, .rs)
- Use language-appropriate test commands

**Structure Preservation:**
- Keep INTERFACE.md as contract
- Keep src/ for implementation
- Keep tests/ separate
