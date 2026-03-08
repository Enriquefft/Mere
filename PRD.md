# Product Requirements Document: Mere

**The minimal, module-first architecture toolkit for Claude Code.**

## 1. Executive Summary

Mere is a foundational developer tool designed to optimize codebases for autonomous AI agents (specifically Claude Code).

The core finding behind Mere is revolutionary in its simplicity: **the codebase structure itself is the most critical factor in AI-assisted development success, far outweighing prompt engineering or agent orchestration systems**. Mere replaces bloated, 15,000-line meta-prompting frameworks with a single, lightweight Go binary that scaffolds exactly what AI needs: deep modules, clean interfaces, and strict boundary tests.

**Slogan:** Just the architecture. AI does the rest.

## 2. Core Philosophy & Economics

Mere is built on three unshakeable pillars:

### 2.1 Structure Over Prompts

Your codebase, far more than your prompts or agent configurations, is the biggest influence on AI output. AI enters a codebase with no memory or context. Mere enforces John Ousterhout's "Deep Modules" pattern: lots of implementation controlled by a simple interface. The map of your codebase needs to be easily navigable and enforced by using modules.

### 2.2 Claude Code Economics (The Cache Advantage)

Mere's workflow is explicitly designed to exploit Claude Code's prefix caching model.

* Cache matches by prefix from the start of the context.


* Cache reads cost 10% of normal token cost, while writes cost 100%.


* Spawning fresh sub-agents (as competitor frameworks do) forces a cold cache every time, multiplying costs by 3-5x.


*
**The Mere Way:** Stay in a single session, use `/compact` to compress state into the cached prefix (`CLAUDE.md`), and maintain a warm cache for 10-20% relative costs .



### 2.3 Taste at the Boundaries

Developers using Mere stop writing implementation code. Instead, they apply taste at the boundaries by writing `INTERFACE.md` contracts and boundary tests. The implementation inside the module is treated as a "graybox" managed entirely by the AI.

## 3. Product Architecture

### 3.1 The Delivery Mechanism (The CLI)

Mere is distributed as a single, compiled **Go binary**.

* **Zero Dependencies:** Users do not need Node.js, Python, or external dependencies.
* **`go:embed`:** The CLI utilizes Go's `go:embed` feature to pack all Markdown templates directly into the binary.
* **Single Repository:** The entire tool, including the CLI parser and the prompt templates, lives in one unified GitHub repository to minimize surface area and bloat.

### 3.2 The Embedded Toolkit (~200 Lines)

When a user runs `mere init`, the CLI injects the following minimal file structure into their project :

```text
.claude/
├── CLAUDE.md              # ~50 lines: project context + boundary rules
└── commands/
    ├── module-init.md     # Prompt: create deep module structure [cite: 121]
    ├── module-work.md     # Prompt: work on ONE module with tests [cite: 122]
    └── compact.md         # Prompt: compress and sync state [cite: 123]
```

### 3.3 The Target Project Structure

The `mere init` and `mere work` commands guide the AI to structure the user's codebase into isolated, highly testable "seams" :

```text
services/
├── [module-name]/
│   ├── [module entry point file]   # INTERFACE: types + public exports
│   ├── INTERFACE.md                # AI-readable contract (Progressive Disclosure)
│   └── src/                        # IMPLEMENTATION: AI manages
tests/
└── services/                        # Boundary tests lock behavior
    └── [boundary test file]

```

**Note:** File patterns are conceptual and adapt to any programming language. The AI detects the project's language and creates appropriate files (e.g., .go, .py, .ts, .java, .rs) following ecosystem conventions.

## 4. Technical Specifications

### 4.1 CLI Commands

1. **`mere init`**: Initializes the `.claude/` directory in the current project, unpacking the embedded Markdown templates.
2. **`mere version`**: Outputs the current version of the CLI.
3. *(Future)* **`mere upgrade`**: Self-updates the binary to the latest GitHub release.

### 4.2 Prompt Specifications

*
**`/module-init`**: Instructs the AI to ask for a module name, generate an interface first, get human approval, and scaffold the implementation/test structure .


*
**`/module-work`**: Instructs the AI to read the `INTERFACE.md`, modify implementation *only*, run boundary tests, and report results without breaking the contract.


*
**`/compact`**: Instructs the AI to summarize work, update state markers in `CLAUDE.md`, compress context, and prepare for continued work.



## 5. Go-to-Market & Distribution

### 5.1 Installation Strategy

Mere prioritizes frictionless, zero-bloat installation using standard native distribution methods.

**The "Hero" Install (Viral/Default):**

```bash
curl -fsSL https://mere.run/install | bash

```

*(Script detects OS/Arch and pulls the standalone Go binary from GitHub Releases into local path).*

**Secondary Package Managers:**

* Go Developers: `go install github.com/[org]/mere@latest`
* macOS/Linux: `brew install [org]/tap/mere`
