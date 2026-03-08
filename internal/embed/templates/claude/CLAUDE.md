# Project Context

## Project Overview

This project uses the Mere architecture toolkit for AI-native development.

## Architecture Philosophy

**Structure Over Prompts:** The codebase structure is the most critical factor in AI-assisted development. We use John Ousterhout's "Deep Modules" pattern: lots of implementation controlled by a simple interface.

**Cache Advantage:** Stay in a single session to exploit Claude Code's prefix caching.

**Taste at the Boundaries:** Apply taste at the boundaries by writing INTERFACE.md contracts and boundary tests. The implementation inside modules is managed by AI.

## Language-Agnostic Philosophy

Mere is designed to work with any programming language. The templates use conceptual descriptors rather than language-specific extensions, allowing the AI to:

1. **Detect the project language** from existing code structure
2. **Adapt file patterns** to language conventions
3. **Use appropriate tooling** for the detected ecosystem

The patterns (deep modules, interfaces, boundary tests) are universal. Only the implementation details vary by language.

**Examples of how this adapts:**
- Go projects: Module entry in `module.go`, tests in `module_test.go`
- Python projects: Module entry in `__init__.py`, tests in `test_module.py`
- TypeScript projects: Module entry in `index.ts`, tests in `module.test.ts`
- Rust projects: Module entry in `mod.rs` or `lib.rs`, tests in `module_test.rs`
- Java projects: Module entry in `Module.java`, tests in `ModuleTest.java`

## Module Structure

```
services/
├── [module-name]/
│   ├── [module entry point file]   # INTERFACE: types + public exports
│   ├── INTERFACE.md                # AI-readable contract (Progressive Disclosure)
│   └── src/                        # IMPLEMENTATION: AI manages
tests/
└── services/                        # Boundary tests lock behavior
    └── [boundary test file]
```

## Available Commands

Use `/module-init` to create a new deep module:
1. Specify module name
2. Generate interface first
3. Get human approval
4. Scaffold implementation/test structure

Use `/module-work` to work on an existing module:
1. Read the INTERFACE.md
2. Modify implementation only
3. Run boundary tests
4. Report results without breaking the contract

## Development Guidelines

1. **Always start with the interface** - Define what, not how
2. **Write boundary tests first** - Lock in expected behavior
3. **Let AI handle implementation** - Treat modules as grayboxes

## Project State

<!-- Update markers as work progresses -->
- Active module: [MODULE_NAME]
- Current phase: [PHASE]
