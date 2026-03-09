# mere

**The minimal, module-first architecture toolkit for Claude Code.**

> Just structure. AI does the rest.

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![GitHub Stars](https://img.shields.io/github/stars/Enriquefft/Mere?style=for-the-badge)](https://github.com/Enriquefft/Mere/stargazers)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)
[![CI](https://github.com/Enriquefft/Mere/workflows/Test/badge.svg)](https://github.com/Enriquefft/Mere/actions/workflows/test.yml)

---

## Quick Start

```bash
# 1. Install
curl -fsSL https://mere.enriquefft.com/install | bash

# 2. Initialize your project
mere init

# 3. Open Claude Code and create a module
/module
```

**That's it.** Mere scaffolds the structure, you define the interface, AI builds the implementation.

---

## Why Mere?

| Problem | Traditional | With Mere |
|---------|-------------|-----------|
| Context cost | Fresh agent = cold cache (3x tokens) | Session continuation = warm cache (1x tokens) |
| Consistency | Drifts over time, no contracts | Interface contracts lock behavior |
| Control | AI modifies freely | Boundaries enforced by `INTERFACE.md` |
| Setup | Complex agent configurations | 3 markdown files |

**The insight:** Your codebase structure is the most critical factor in AI-assisted development — far outweighing prompt engineering or agent orchestration.

Mere replaces bloated meta-prompting frameworks with a single Go binary that enforces John Ousterhout's "Deep Modules" pattern.

---

## What Gets Created

`mere init` scaffolds:

```
.claude/
├── commands/
│   ├── module.md        # /module command
│   └── status.md        # /status command
├── templates/
│   └── INTERFACE.md     # Interface contract template
└── context/
    ├── ARCHITECTURE.md  # Your architecture decisions (human-written)
    └── MANIFEST.md      # Module inventory (AI-maintained)
CLAUDE.md                # Entry point — auto-loaded by Claude Code
```

Three files drive every session:

| File | Contains | Owner | Changes |
|------|----------|-------|---------|
| `ARCHITECTURE.md` | Decisions, constraints, conventions | Human | Rarely |
| `MANIFEST.md` | Module inventory, dependency graph | AI | Every module op |
| `INTERFACE.md` | Module public contract | Human | When scope changes |

**Human owns the rules. AI owns the map. You define the contracts.**

---

## Architecture Principles

### Structure Over Prompts

Your codebase, not prompts or agent configs, is the biggest influence on AI output. Mere enforces deep modules: lots of implementation behind a simple interface.

### Cache Advantage

Stay in a single session to exploit Claude Code's prefix caching. Cache reads cost 10% of normal tokens. Same output, fraction of the cost.

### Taste at the Boundaries

Apply taste at module boundaries by writing `INTERFACE.md` contracts and boundary tests. Implementation inside modules is a "graybox" managed by AI.

---

## What is a Deep Module?

From John Ousterhout's *A Philosophy of Software Design*:

> "The best modules are deep: they have a simple interface but a complex implementation."

| | Interface | Implementation | Depth |
|---|-----------|----------------|-------|
| **Shallow module** | 10 lines | 10 lines | 1:1 |
| **Deep module** | 10 lines | 1,000 lines | 1:100 |

Mere enforces depth by separating `INTERFACE.md` (simple, human-written) from implementation (complex, AI-generated). You define the "what." AI figures out the "how."

---

## The /module Workflow

When you run `/module`, the AI follows a strict order:

1. **Read first** — `ARCHITECTURE.md` (constraints, decisions) and `MANIFEST.md` (existing modules)
2. **Design** — draft `INTERFACE.md`, get your approval before writing any code
3. **Build** — scaffold structure, write boundary tests, implement
4. **Update last** — `MANIFEST.md` only after work is complete and tests pass

AI never modifies `ARCHITECTURE.md`. It can suggest additions — you decide.

---

## Example

Here's what you write in `INTERFACE.md`:

```markdown
# Module: UserAuth

## Public API
- `authenticate(email, password) → Token`
- `validate(token) → User | null`
- `logout(token) → void`

## Constraints
- No external auth services
- Tokens expire in 24h
- Max 5 failed attempts before lockout

## Errors
- Invalid credentials → AuthError
- Account locked → LockoutError
- Token expired → ExpiredError
```

The AI builds the full implementation from this contract — password hashing, token generation, lockout logic, error types, and tests.

---

## Installation

**One-command install:**
```bash
curl -fsSL https://mere.enriquefft.com/install | bash
```

**From Go:**
```bash
go install github.com/Enriquefft/Mere@latest
```

**From source:**
```bash
git clone https://github.com/Enriquefft/Mere.git
cd Mere
just build && just install
```

---

## Usage

```bash
mere init        # Initialize .claude/ directory
mere version     # Show version
mere --help      # Show all commands
```

After `mere init`, use these commands in Claude Code:

| Command | Description |
|---------|-------------|
| `/module` | Create or work on a deep module |
| `/status` | Report project state, update MANIFEST.md |

---

## Development

```bash
just build   # Build to ./bin/
just test    # Run tests
just dev     # Clean + deps + build + test
just fmt     # Format code
just lint    # Run linter
```

Nix dev shell: `direnv reload`

---

## Contributing

Contributions welcome — see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE).
